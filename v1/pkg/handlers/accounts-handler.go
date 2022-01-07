package v1

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apauth "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth"
	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/models/req"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/models/resp"
	apmodels "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// AccountsHandler incoming request for accounts
type AccountsHandler struct {
	*base
}

// PostCreateAccount add a new account
func (h *AccountsHandler) PostCreateAccount(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionCreateAccount) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	b := aphv1req.CreateAccount{}

	if e = c.ShouldBindWith(&b, binding.FormMultipart); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	// SOLO simula la generacion del username. NO usar en production
	username := h.generateUsername(b.Email)
	var (
		outAccountByEmail    apv1models.Account
		outAccountByUsername apv1models.Account
	)
	conc := make(chan *aputil.ConcurrencyData)
	totalReqConc := 0
	// account email
	totalReqConc++
	go h.findAccountByEmail(b.Email, conc, &outAccountByEmail)
	// account username
	totalReqConc++
	go h.findAccountByUsername(username, conc, &outAccountByUsername)

	for i := 0; i < totalReqConc; i++ {
		if cData := <-conc; cData != nil && cData.Err != nil {
			// check error email y/o username available
			if e := cData.Err; e != apdbabstract.ErrorNoItems {
				log.Printf("ms id throws error: %s", cData.Name)
				resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
				return
			}
		}
	}
	if len(outAccountByEmail.Email) != 0 {
		resp.send(http.StatusConflict, aphv1resp.CodeConflictEmail, nil, nil, true)
		return
	}
	if len(outAccountByUsername.Username) != 0 {
		resp.send(http.StatusConflict, aphv1resp.CodeConflictUsername, nil, nil, true)
		return
	}
	// save avatar
	pathInAvatar, e := h.saveUploadFile(avatarDirectoryIn, b.Avatar)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	// generate password
	hashPass, e := apauth.GeneratePassword([]byte(b.Password))
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	item := h.toModelAccountFromRequest(&b, username, hashPass, pathInAvatar)
	// save on db
	id, e := h.db.Accounts().Add(item)

	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(&aphv1resp.ResponseID{ID: id}, nil, false)
}

// PostLogin check usernameOrEmail and password
func (h *AccountsHandler) PostLogin(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionLogin) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	b := aphv1req.Login{}

	if e = c.ShouldBindWith(&b, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	usernameOrEmail := b.UsernameOrEmail
	var acc *apv1models.Account
	if strings.Contains(usernameOrEmail, "@") {
		acc, e = h.db.Accounts().FindByEmail(usernameOrEmail)
	} else {
		acc, e = h.db.Accounts().FindByUsername(usernameOrEmail)
	}
	if e != nil {
		if e == apdbabstract.ErrorNoItems { // not found username or email
			resp.sendNotFound(aphv1resp.CodeNotFoundUser, e, nil, true)
			return
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
			return
		}
	}
	if e = apauth.ComparePasswords(acc.Password, []byte(b.Password)); e != nil {
		resp.send(http.StatusForbidden, aphv1resp.CodeIncorrectPassword, e, nil, true)
		return
	}
	deviceID := aputil.RandString() // identifica al dispositvo desde el q se hizo login
	// generating access-token(jwt) adn refresh-token
	token, e := h.token.Create(&apmodels.AuthUser{
		UserID:   acc.ID,
		Fullname: acc.Fullname,
		Username: acc.Username,
		Roles:    acc.Roles,
		Avatar:   acc.Avatar,
	})
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	ua := parseUserAgent(c.Request.UserAgent())
	ip := c.ClientIP()
	now := time.Now().UTC().Unix()
	// sessiones creadas por el usuario al hacer login.
	// util para seguridad al generar un nuevo access-token a partir del refresh-token
	sess := apv1models.Session{
		ID:                         "",
		UserID:                     acc.ID,
		RefreshToken:               token.RefreshToken,
		DeviceID:                   deviceID,
		UserAgentStr:               ua.String,
		UserAgent:                  ua.Name,
		Platform:                   ua.OS,
		IP:                         ip,
		Location:                   aputil.GetLocationFromIP(ip),
		LastAccessTokenGeneratedAt: now,
		CreatedAt:                  now,
	}
	_, e = h.db.Sessions().Add(&sess)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(&aphv1resp.Login{
		ID:            acc.ID,
		AuthTwoFactor: false,
		Token:         &aphv1resp.Token{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken},
		Fullname:      acc.Fullname,
		Avatar:        h.fullURLMedia(acc.Avatar),
		DeviceID:      deviceID,
	}, nil, false)
}

// PostGenerateAccessToken genera un nuevo access token valido
func (h *AccountsHandler) PostGenerateAccessToken(c *gin.Context) {
	resp, u := response{c: c, env: h.env}, h.getUser(c)
	if !h.authorize(u, apmodels.PermissionGenerateAccessToken) {
		resp.send(http.StatusForbidden, aphv1resp.CodeUnauthorized, errorUnauthorized, nil, true)
		return
	}
	var e error
	b := aphv1req.GenerateAccessToken{}

	if e = c.ShouldBindWith(&b, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	uClaimToken, e := h.token.DecodeYetInvalid(b.OldAccessToken) // parser old access token
	if e != nil || uClaimToken == nil || len(uClaimToken.Id) == 0 {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	ses, e := h.db.Sessions().FindByRefreshToken(b.RefreshToken)
	if e != nil {
		if e == apdbabstract.ErrorNoItems { // not found session by refresh token
			resp.sendNotFound(aphv1resp.CodeNotFoundUser, e, nil, true)
			return
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
			return
		}
	}
	ua := parseUserAgent(c.Request.UserAgent())
	now := time.Now().UTC().Unix()

	// check data session associated with the refresh token

	if now-ses.CreatedAt > h.token.GetLiveRefreshToken() { // Comprobando el refresh token sea valido
		_ = h.db.Sessions().Remove(ses.ID)
		resp.send(http.StatusUnauthorized, aphv1resp.CodeExpiredRefreshToken, errorExpiredRefreshToken, nil, true)
		return
	}
	if ses.DeviceID != b.DeviceID {
		_ = h.db.Sessions().Remove(ses.ID)
		resp.send(http.StatusUnauthorized, aphv1resp.CodeInvalidSessionDataOfRefreshToken, errorInvalidSessionDataOfRefreshToken, nil, true)
		return
	}
	if ses.Platform != ua.OS {
		_ = h.db.Sessions().Remove(ses.ID)
		resp.send(http.StatusUnauthorized, aphv1resp.CodeInvalidSessionDataOfRefreshToken, errorInvalidSessionDataOfRefreshToken, nil, true)
		return
	}
	if ses.UserAgent != ua.Name {
		_ = h.db.Sessions().Remove(ses.ID)
		resp.send(http.StatusUnauthorized, aphv1resp.CodeInvalidSessionDataOfRefreshToken, errorInvalidSessionDataOfRefreshToken, nil, true)
		return
	}
	uToken, e := h.db.Accounts().Find(uClaimToken.UserID)
	if e != nil {
		if e == apdbabstract.ErrorNoItems { // not found user of old access token
			resp.sendNotFound(aphv1resp.CodeNotFoundUser, e, nil, true)
			return
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
			return
		}
	}
	if uToken.ID != ses.UserID { // check session is of user on old access token
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e, nil, true)
		return
	}
	// generate new access token
	t, e := h.token.Create(&apmodels.AuthUser{
		UserID:      uToken.ID,
		Fullname:    uToken.Fullname,
		Username:    uToken.Username,
		Roles:       uToken.Roles,
		Permissions: apauth.GetPermissions(uToken.Roles),
		Avatar:      uToken.Avatar,
	})
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	e = h.db.Sessions().EditLastAccessTokenGenerated(ses.ID, now)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e, nil, true)
		return
	}
	resp.sendOK(&aphv1resp.GenerateAccessToken{
		AccessToken: t.AccessToken,
	}, nil, false)
}

// === conv form request === //

func (h *AccountsHandler) toModelAccountFromRequest(d *aphv1req.CreateAccount, username, hashPasword, pathInAvatar string) *apv1models.Account {
	roles := []apmodels.RoleType{apmodels.RoleUser}
	// simular un usuario de admin
	if d.Email == "vallin.plasencia@gmail.com" {
		roles = []apmodels.RoleType{apmodels.RoleUser, apmodels.RoleAdmin}
	}
	return &apv1models.Account{
		ID:        "",
		Fullname:  d.Fullname,
		Email:     d.Email,
		Username:  username,
		Password:  hashPasword,
		Roles:     roles,
		Avatar:    pathInAvatar,
		CreatedAt: time.Now().UTC().Unix(),
		UpdatedAt: 0,
	}
}

// generateUsername simula la generacion de un username "UNICO" a partir del email.
//
// NO usar en produccion
func (h *AccountsHandler) generateUsername(email string) (username string) {
	a := email[:strings.Index(email, "@")]
	return fmt.Sprintf("%s_%s", a, strings.ToLower(aputil.RandStringn(5)))
}
