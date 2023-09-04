package main

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/mwinters-stuff/halo-one-thing/gen/models"
	"github.com/mwinters-stuff/halo-one-thing/gen/restapi"
	"github.com/mwinters-stuff/halo-one-thing/gen/restapi/operations"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewServerAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// TODO: Set Handle
	api.GetVersionHandler = operations.GetVersionHandlerFunc(
		func(gvp operations.GetVersionParams) middleware.Responder {
			payload := models.Version{
				Version: "1.0.0",
			}
			return operations.NewGetVersionOK().WithPayload(&payload)
		},
	)

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
