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
   Path: helpers/errorResponse.go
*/
package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ErrUnmatched struct {
	Timestamp int64  `json:"timestamp"`
	Code      int    `json:"code"`
	Arguments string `json:"arguments"`
	Details   string `json:"details"`
	Message   string `json:"message"`
}

func Unmatched(w http.ResponseWriter, r *http.Request) {

	ApiErr := ErrUnmatched{
		Timestamp: time.Now().Unix(),
		Code:      100,
		Arguments: "ERR_API_URL_UNMATCHED",
		Details:   "Invalid action ",
		Message:   fmt.Sprintf("ERROR: Not Valid Action Method:%s Url:%s", r.Method, r.URL),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Status", "404 Not Found")
	w.WriteHeader(http.StatusNotFound)

	//fmt.Printf("ERR:%d MSG:%s Method:%s Url:%s\n", ApiErr.Code,ApiErr.Arguments,r.Method,r.URL)
	Log.Errorf("API ERR:%d MSG:%s Method:%s Url:%s", ApiErr.Code, ApiErr.Arguments, r.Method, r.URL)

	ApiErr.ApiErrorResponse(w, r)
}

func (Err *ErrUnmatched) ApiErrorResponse(w http.ResponseWriter, r *http.Request) {

	Err.Timestamp = time.Now().Unix()
	//w.Header().Set("Status", "400 Bad Request")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(Err); err != nil {
		Log.Errorf("ERROR %s", err.Error())
	}
}
