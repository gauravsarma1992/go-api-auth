package goapiauth

import (
	"encoding/json"
	"io/ioutil"
)

const (
	DefaultJwtConfigFile = "./config/jwt.json"
)

type (
	AuthConfig struct {
		SigningKey        string `json:"signing_key"`
		Issuer            string `json:"issuer"`
		ValidityInMinutes int    `json:"validity_in_minutes"`
	}
)

func (auth *Authenticator) readConfig() (err error) {
	var (
		fileB []byte
	)
	auth.Config = &AuthConfig{}
	if fileB, err = ioutil.ReadFile(DefaultJwtConfigFile); err != nil {
		return
	}
	if err = json.Unmarshal(fileB, auth.Config); err != nil {
		return
	}
	return
}
