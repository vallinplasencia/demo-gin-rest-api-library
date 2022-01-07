package server

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	apauthtoken "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/auth/access-token/dgrijalva"
	apdbclient "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/db/mysql"
	apstoreabstract "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/abstract"
	apstorelocalclient "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/local"
	apstores3client "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/external-services/store/s3"
	aphandv1 "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/handlers"
	aprouter "github.com/vallinplasencia/demo-gin-rest-api-library/v1/pkg/routers"
)

// Server ...
type Server struct {
	eng *gin.Engine

	srvHTTP  *http.Server
	srvHTTPS *http.Server

	pathCertHTTPS string
	pathKeyHTTPS  string
}

func New() *Server {
	projectName := "LIBRARY"

	// server http-https
	serverConf, e := ConfigFromEnv(projectName)
	we("unable read server environment", e)

	// db
	dbConf, e := apdbclient.ConfigFromEnv(projectName)
	we("unable read db environment", e)
	db, e := apdbclient.New(dbConf)
	we("unable configurate db client", e)

	// token: access-token(jwt) and refresh token auth
	tokenConf, e := apauthtoken.ConfigFromEnv(projectName)
	we("unable read auth token environment", e)
	token, e := apauthtoken.New(tokenConf)
	we("unable configurate auth token", e)

	uploadPath := ""
	var storeFiles apstoreabstract.Store
	if serverConf.StoreUploadedFilesMode == StoreAwsS3 {
		// store upload files on aws-s3
		storeAWSS3Conf, e := apstores3client.ConfigFromEnv(projectName)
		we("unable read store aws-s3 environment", e)
		storeFiles, e = apstores3client.New(storeAWSS3Conf)
		we("unable configurate store aws-s3 client", e)
	} else {
		// store upload files on files-system-local
		storeLocalConf, e := apstorelocalclient.ConfigFromEnv(projectName)
		we("unable read store files-system-local environment", e)
		storeFiles, e = apstorelocalclient.New(storeLocalConf)
		we("unable configurate store files-system-local client", e)
		uploadPath = storeLocalConf.DestinationTarget
	}
	// handlers incoming request
	handlersConf, e := aphandv1.ConfigFromEnv(projectName)
	we("unable read handlers environment", e)
	handlers := aphandv1.New(handlersConf, db, token, storeFiles)
	we("unable configurate handlers", e)

	if len(serverConf.PathFileLogs) > 0 { // logs on file if found path on enviroment
		gin.DisableConsoleColor()
		f, e := os.Create(serverConf.PathFileLogs)
		we("unable create file logs", e)
		gin.DefaultWriter = io.MultiWriter(f)
	}
	eng := gin.New()
	eng.Use(gin.Logger())
	eng.Use(gin.Recovery())

	// ======= init routers with yours handlers ======= //
	r := aprouter.Router{
		Token:      token,
		Eng:        eng,
		H:          handlers,
		UploadPath: uploadPath,
	}
	r.InitRouters()

	return &Server{
		eng: eng,
		srvHTTP: &http.Server{
			Addr:           serverConf.AddressHTTP,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		srvHTTPS: &http.Server{
			Addr:           serverConf.AddressHTTPS,
			Handler:        eng,
			ReadTimeout:    time.Duration(serverConf.ReadTimeout) * time.Second,
			WriteTimeout:   time.Duration(serverConf.WriteTimeout) * time.Second,
			MaxHeaderBytes: 0,
		},
		pathCertHTTPS: serverConf.PathCertHTTPS,
		pathKeyHTTPS:  serverConf.PathKeyHTTPS,
	}
}

// Run ...
func (s *Server) Run() {
	// http
	go func() {
		log.Printf("http server listen on :%s", s.srvHTTP.Addr)

		if e := s.srvHTTP.ListenAndServe(); e != nil && !errors.Is(e, http.ErrServerClosed) {
			log.Fatalf("http server not run, error: %s", e.Error())
		}
	}()
	if len(s.pathCertHTTPS) != 0 && len(s.pathKeyHTTPS) != 0 { // run https only if find path certs
		// https
		go func() {
			log.Printf("https server listen on :%s", s.srvHTTPS.Addr)

			if e := s.srvHTTPS.ListenAndServeTLS(s.pathCertHTTPS, s.pathKeyHTTPS); e != nil && !errors.Is(e, http.ErrServerClosed) {
				log.Fatalf("https server not run, error: %s", e.Error())
			}
		}()
	} else {
		log.Println("https server not run, cert and/or key no find on enviroment")
	}
}

// Run ...
func (s *Server) Shutdown(ctx context.Context) {
	if e := s.srvHTTP.Shutdown(ctx); e != nil {
		log.Printf("http server shutdown, error: %s", e.Error())
	} else {
		log.Printf("http server shutdown")
	}
	if e := s.srvHTTPS.Shutdown(ctx); e != nil {
		log.Printf("https server shutdown, error: %s", e.Error())
	} else {
		log.Printf("https server shutdown")
	}
}

// log error
func we(prefix string, e error) {
	if e != nil && len(prefix) > 0 {
		log.Fatalf("%s, %s", prefix, e.Error())
	} else if e != nil {
		log.Fatal(e)
	}
}
