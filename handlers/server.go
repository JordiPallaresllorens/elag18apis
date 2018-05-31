package handlers

import (
	"log"

	"github.com/cmh2166/elag18apis/authorization"
	"github.com/cmh2166/elag18apis/db"
	"github.com/cmh2166/elag18apis/generated/restapi"
	"github.com/cmh2166/elag18apis/generated/restapi/operations"
	"github.com/cmh2166/elag18apis/identifier"
	"github.com/go-openapi/loads"
)

// BuildAPI create new service API
func BuildAPI(database db.Database, identifierService identifier.Service) *operations.TaquitoAPI {
	api := operations.NewTaquitoAPI(swaggerSpec())
	api.RemoteUserAuth = func(identifier string) (*authorization.Agent, error) {
		return &authorization.Agent{Identifier: identifier}, nil
	}
	api.RetrieveResourceHandler = NewRetrieveResource(database)
	api.DeleteResourceHandler = NewDeleteResource(database)
	api.DepositResourceHandler = NewDepositResource(database, identifierService)
	api.UpdateResourceHandler = NewUpdateResource(database)
	api.HealthCheckHandler = NewHealthCheck()
	return api
}

func swaggerSpec() *loads.Document {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	return swaggerSpec
}
