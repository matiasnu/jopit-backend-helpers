/*
   Copyright 2021 JOPIT
   Contact: matiasne45@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   Language: go
   Path: jsonHelpers.go
*/
package helpers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lithammer/shortuuid"
)

var ApiUuid string

func ValidatorJsonDecode(err error, w http.ResponseWriter, r *http.Request) bool {

	isValid := err == nil

	if !isValid {
		ApiError("JSON", fmt.Sprintf("fail Decode Json to model [%s] [%t]", err.Error(), isValid))
		ApiErr := ErrUnmatched{
			Code:      100,
			Arguments: "ERR_API_DECODE_JSON",
			Details:   "Fail Json Decode",
			Message:   fmt.Sprintf("ERROR: Fail Json Decode [%s]", err),
		}
		w.WriteHeader(http.StatusBadRequest)
		ApiErr.ApiErrorResponse(w, r)
	}
	return isValid
}

func ValidatorJsonEncode(err error, w http.ResponseWriter, r *http.Request) bool {

	isValid := err == nil

	if !isValid {
		ApiError("JSON", fmt.Sprintf("fail encode Json to model [%s]", err.Error()))
		ApiErr := ErrUnmatched{
			Code:      101,
			Arguments: "ERR_API_ENCODE_JSON",
			Details:   "Fail Json Encode",
			Message:   fmt.Sprintf("ERROR: Fail Json Encode [%s]", err),
		}
		w.WriteHeader(http.StatusBadRequest)
		ApiErr.ApiErrorResponse(w, r)
	}
	return isValid
}

func ValidatorEmptyId(w http.ResponseWriter, r *http.Request) bool {

	ApiError("ID_NOT_EMPTY", "Record ID is mandatory")
	ApiErr := ErrUnmatched{
		Code:      109,
		Arguments: "ID_NOT_EMPTY",
		Details:   "Record ID is mandatory",
		Message:   "ERROR: Record ID is mandatory",
	}
	w.WriteHeader(http.StatusBadRequest)
	ApiErr.ApiErrorResponse(w, r)
	return false
}

func SetCookieInHTTP(sessionToken string, w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(120 * time.Second),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
	}
	http.SetCookie(w, c)
}

func GenShortUUID() string {
	id := shortuuid.New()
	return id
}
