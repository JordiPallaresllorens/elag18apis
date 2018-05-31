package handlers

import (
	"log"
	"time"

	"github.com/cmh2166/elag18apis/authorization"
	"github.com/cmh2166/elag18apis/datautils"
	"github.com/cmh2166/elag18apis/db"
	"github.com/cmh2166/elag18apis/generated/restapi/operations"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cmh2166/elag18apis/identifier"
)

// NewDepositResource -- Accepts requests to create resource and pushes them to Kinesis.
func NewDepositResource(database db.Database, identifierService identifier.Service) operations.DepositResourceHandler {
	return &depositResource{
		database:          database,
		identifierService: identifierService,
	}
}

type depositResource struct {
	database          db.Database
	identifierService identifier.Service
}

// Handle the create resource request
func (d *depositResource) Handle(params operations.DepositResourceParams, agent *authorization.Agent) middleware.Responder {
	resource := params.Payload

	authService := authorization.NewService(agent)
	if !authService.CanCreateResourceOfType(*resource.AtType) {
		log.Printf("Agent %s is not permitted to create a resource of type %s", agent, *resource.AtType)
		return operations.NewDepositResourceUnauthorized()
	}

	id, err := identifier.NewUUIDService().Mint()
	if err != nil {
		panic(err)
	}

	resource.Identifier = id
	WithID(id).
		WithVersion(1)
	(*resource.Administrative())["created"] = time.Now().UTC().Format(time.RFC3339)

	if err := d.database.Insert(resource); err != nil {
		panic(err)
	}

	response := datautils.JSONObject{"id": externalID}
	return operations.NewDepositResourceCreated().WithPayload(response)
}
