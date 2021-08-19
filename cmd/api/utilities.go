package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	wrapper := make(map[string]interface{})

	wrapper[wrap] = data

	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) {
	// default code: 400
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	// Status Code 400 (Bad Request if Error)
	app.writeJSON(w, statusCode, theError, "error")
}

func generateHashedPassword() string {
	password := "password"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)

	return string(hashedPassword)
}

// func generateJwtSecret() string {

// 	secret := "mysecret"
// 	data := "data"

// 	// Create a new HMAC by defining the hash type and the key (as byte array)
// 	h := hmac.New(sha256.New, []byte(secret))

// 	// Write Data to it
// 	h.Write([]byte(data))

// 	// Get result and encode as hexadecimal string
// 	sha := hex.EncodeToString(h.Sum(nil))

// 	return sha
//  // 2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160
// }
