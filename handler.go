package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"

	"github.com/sirupsen/logrus"
)

func restartHandler(w http.ResponseWriter, r *http.Request) {
	if config.EnableBasicAuth {
		ok, errMsg := checkBasicAuthCredential(w, r)
		if !ok {
			logrus.Error(errMsg)
			http.Error(w, errMsg, http.StatusUnauthorized)
			return
		}
	}

	// TODO
	run("touch /tmp/restart")
}

func checkBasicAuthCredential(w http.ResponseWriter, r *http.Request) (bool, string) {
	username, password, ok := r.BasicAuth()
	if !ok {
		errMsg := "failed to get basic auth credential"
		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		return false, errMsg
	}
	hashedUsername := sha256.Sum256([]byte(username))
	hashedPassword := sha256.Sum256([]byte(password))
	expectedUsernameHash := config.BasicAuthUsername
	expectedPasswordHash := config.BasicAuthPassword

	usernameMatch := (subtle.ConstantTimeCompare(hashedUsername[:], expectedUsernameHash[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(hashedPassword[:], expectedPasswordHash[:]) == 1)

	if !usernameMatch || !passwordMatch {
		errMsg := "credential mismatched"
		return false, errMsg
	}

	return true, ""
}
