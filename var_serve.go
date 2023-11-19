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

import (
	"log"
	"net/http"
	"os"

	"github.com/rwscode/webclip/middleware"
)

func serve(addr string) { ServeRouter(addr, "/webclip/generate") }

func serveRouter(addr, generateRouter string) {
	logger := log.New(os.Stdout, "[webclip] ", log.LstdFlags|log.Lshortfile)
	logger.Println("served on " + addr)
	http.HandleFunc(generateRouter, WrappedGenerateHandlerFunc(middleware.Cors()))
	logger.Println(http.ListenAndServe(addr, nil))
}