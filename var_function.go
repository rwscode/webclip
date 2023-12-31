// Copyright 2023 webclip Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webclip

import "net/http"

func walkHandlerFunc(handlers ...http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, h := range handlers {
			h(w, r)
		}
	}
}

func writeJSON(w http.ResponseWriter, str string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(str))
}

func writeXML(w http.ResponseWriter, str string) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Write([]byte(str))
}

func writeStream(w http.ResponseWriter, str, fileName string) {
	w.Header().Set("Content-Type", "text/xml; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="`+fileName+`"`)
	w.Write([]byte(str))
}
