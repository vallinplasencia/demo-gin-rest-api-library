package v1

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	apauth "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth"
	apdbabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/abstract"
	aphv1req "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/req"
	aphv1resp "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers/v1/models/resp"
	apv1models "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/models/v1"
	aputil "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/util"
)

// AccountsHandler incoming request for accounts
type AccountsHandler struct {
	*base
}

// PostCreateAccount add a new account
func (h *AccountsHandler) PostCreateAccount(c *gin.Context) {
	resp := response{c: c, env: h.env}
	var e error
	b := aphv1req.CreateAccount{}

	if e = c.ShouldBindWith(&b, binding.FormMultipart); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e)
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
				resp.sendInternalError(aphv1resp.CodeInternalError, e)
				return
			}
		}
	}
	if len(outAccountByEmail.Email) != 0 {
		resp.send(http.StatusConflict, aphv1resp.CodeConflictEmail, errors.New("email exist"))
		return
	}
	if len(outAccountByUsername.Username) != 0 {
		resp.send(http.StatusConflict, aphv1resp.CodeConflictUsername, errors.New("username exist"))
		return
	}
	// save avatar
	pathInAvatar, e := h.saveUploadFile(avatarDirectoryIn, b.Avatar)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	// generate password
	hashPass, e := apauth.GeneratePassword([]byte(b.Password))
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	item := h.toModelAccountFromRequest(&b, username, hashPass, pathInAvatar)
	// save on db
	id, e := h.db.Accounts().Add(item)

	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	resp.sendOK(&aphv1resp.ResponseID{ID: id})
}

// PostLogin check usernameOrEmail and password
func (h *AccountsHandler) PostLogin(c *gin.Context) {
	resp := response{c: c, env: h.env}
	var e error
	b := aphv1req.Login{}

	if e = c.ShouldBindWith(&b, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e)
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
			resp.sendNotFound(aphv1resp.CodeNotFoundUser, e)
			return
		} else {
			resp.sendInternalError(aphv1resp.CodeInternalError, e)
			return
		}
	}
	if e = apauth.ComparePasswords(acc.Password, []byte(b.Password)); e != nil {
		resp.send(http.StatusUnauthorized, aphv1resp.CodeIncorrectPassword, e)
		return
	}
	deviceID := aputil.RandString() // identifica al dispositvo desde el q se hizo login
	// generating access-token(jwt) adn refresh-token
	token, e := h.token.Create(acc)
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
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
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	resp.sendOK(&aphv1resp.Login{
		AuthTwoFactor: false,
		Token:         &aphv1resp.Token{AccessToken: token.AccessToken, RefreshToken: token.RefreshToken},
		Fullname:      acc.Fullname,
		Avatar:        acc.Avatar,
		DeviceID:      deviceID,
	})
}

// === conv === //

func (h *AccountsHandler) toModelAccountFromRequest(d *aphv1req.CreateAccount, username, hashPasword, pathInAvatar string) *apv1models.Account {
	now := time.Now().UTC().Unix()
	roles := []apv1models.RoleType{apv1models.RoleUser}
	if d.Email == "vallin.plasencia@gmail.com" { // simular un usuario de admin
		roles = []apv1models.RoleType{apv1models.RoleUser, apv1models.RoleAdmin}
	}
	return &apv1models.Account{
		ID:        "",
		Fullname:  d.Fullname,
		Email:     d.Email,
		Username:  username,
		Password:  hashPasword,
		Roles:     roles,
		Avatar:    pathInAvatar,
		CreatedAt: now,
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