package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
)

var headers = map[string]string{}

// Cors handle cors
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		for k, v := range headers {
			c.Writer.Header().Set(k, v)
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func init() {
	// === headers cors === //
	c, e := configCorsFromEnv(projectName)
	if e != nil {
		log.Fatalln(e)
	}
	// default headers cors allow
	headers["Access-Control-Allow-Origin"] = "*"
	headers["Access-Control-Allow-Methods"] = "POST, GET, PUT, DELETE, UPDATE"
	headers["Access-Control-Allow-Headers"] = "Content-Type, Content-Length, Accept-Encoding, Authorization, Accept"

	if v := c.AccessControlAllowOrigin; len(v) > 0 {
		headers["Access-Control-Allow-Origin"] = v
	}
	if v := c.AccessControlAllowMethods; len(v) > 0 {
		headers["Access-Control-Allow-Methods"] = v
	}
	if v := c.AccessControlAllowHeaders; len(v) > 0 {
		headers["Access-Control-Allow-Headers"] = v
	}
}

// config ...
type config struct {
	AccessControlAllowOrigin  string `envconfig:"HEADER_ACCESS_CONTROL_ALLOW_ORIGIN"`
	AccessControlAllowMethods string `envconfig:"HEADER_ACCESS_CONTROL_ALLOW_METHODS"`
	AccessControlAllowHeaders string `envconfig:"HEADER_ACCESS_CONTROL_ALLOW_HEADERS"`
}

// configCorsFromEnv ...
func configCorsFromEnv(prefix string) (*config, error) {
	conf := config{}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
