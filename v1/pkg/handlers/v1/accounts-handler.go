package v1

import (
	"errors"
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
	*Base
}

// PostCreateAccount add a new account
func (h *AccountsHandler) PostCreateAccount(c *gin.Context) {
	resp := response{c: c, env: h.env}
	var e error
	b := aphv1req.CreateAccount{}

	if e = c.ShouldBindWith(&b, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e)
		return
	}
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
	go h.findAccountByUsername(b.Username, conc, &outAccountByUsername)

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
		resp.sendNotFound(aphv1resp.CodeConflictEmail, errors.New("email exist"))
		return
	}
	if len(outAccountByUsername.Username) != 0 {
		resp.sendNotFound(aphv1resp.CodeConflictUsername, errors.New("username exist"))
		return
	}
	// save on db
	item := h.toModelAccountFromRequest(&b)
	item.Password, e = apauth.GeneratePassword([]byte(b.Password))
	if e != nil {
		resp.sendInternalError(aphv1resp.CodeInternalError, e)
		return
	}
	id, e := h.DB.Accounts().Add(item)

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
	b := aphv1req.LoginUser{}

	if e = c.ShouldBindWith(&b, binding.JSON); e != nil {
		resp.sendBadRequest(aphv1resp.CodeInvalidArgument, e)
		return
	}
	usernameOrEmail := b.UsernameOrEmail
	var acc *apv1models.Account
	if strings.Contains(usernameOrEmail, "@") {
		acc, e = h.DB.Accounts().FindByEmail(usernameOrEmail)
	} else {
		acc, e = h.DB.Accounts().FindByUsername(usernameOrEmail)
	}
	if e != nil {
		// sino se encuentra el username or email
		if e == apdbabstract.ErrorNoItems {
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
}

// === conv === //

func (h *AccountsHandler) toModelAccountFromRequest(d *aphv1req.CreateAccount) *apv1models.Account {
	now := time.Now().UTC().Unix()
	return &apv1models.Account{
		ID:       "",
		Fullname: d.Fullname,
		Email:    d.Email,
		Username: d.Username,
		// Password:  d.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
