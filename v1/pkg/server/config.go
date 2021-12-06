package server

import (
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	AddressHTTP   string `envconfig:"ADDRESS_HTTP"`
	AddressHTTPS  string `envconfig:"ADDRESS_HTTPS"`
	PathCertHTTPS string `envconfig:"PATH_CERT_HTTPS"`
	PathKeyHTTPS  string `envconfig:"PATH_KEY_HTTPS"`

	ReadTimeout  int64 `envconfig:"READ_TIMEOUT"`
	WriteTimeout int64 `envconfig:"WRITE_TIMEOUT"`

	PathFileLogs string `envconfig:"PATH_FILE_LOGS"`

	StoreUploadedFilesMode StoreFileType `envconfig:"STORE_UPLOADED_FILES_MODE"`
}

// ConfigFromEnv ...
func ConfigFromEnv(prefix string) (*config, error) {
	conf := config{
		AddressHTTP:   ":80",
		AddressHTTPS:  ":443",
		PathCertHTTPS: "",
		PathKeyHTTPS:  "",

		ReadTimeout:  10,
		WriteTimeout: 10,

		PathFileLogs:           "",
		StoreUploadedFilesMode: StoreFilesSystemLocal,
	}
	e := envconfig.Process(prefix, &conf)
	return &conf, e
}
