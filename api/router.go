/*
Copyright Pascal Limeux. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"github.com/gorilla/mux"
	"github.com/pascallimeux/ocms/utils/log"
	"net/http"
)

type AppContext struct {
	HttpServer *http.Server
}

// Initialize API
func (appContext *AppContext) CreateRoutes() *mux.Router {
	log.Trace(log.Here(), "createRoutes() : calling method -")
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/ocms/v1/version", appContext.getVersion).Methods("GET")
	return router
}
