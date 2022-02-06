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
   Path: apiLog.go
*/
package helpers

import (
	"github.com/sirupsen/logrus"
)

// TODO Add UID to log "uid": router.ApiUuid

func ApiError(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Error(msg)
}

func ApiInfo(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Info(msg)
}

func ApiWarn(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Warn(msg)
}

func ApiDebug(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Debug(msg)
}

func ApiPanic(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Panic(msg)
}

func ApiFatal(ref string, msg string) {
	Log.WithFields(logrus.Fields{"source": "API", "ref": ref}).Fatal(msg)
}
