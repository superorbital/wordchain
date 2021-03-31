/*
Copyright © 2021 SuperOrbital, LLC <info@superorbital.io>

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
package listen

import (
	"log"

	types "github.com/superorbital/wordchain/types"

	"github.com/go-openapi/loads"

	"github.com/superorbital/wordchain/restapi"
	"github.com/superorbital/wordchain/restapi/operations"
)

func Listen(prefs types.Preferences, settings types.Listener) error {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewWordChainsAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	server.Port = settings.Port
	server.Host = "0.0.0.0"
	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
