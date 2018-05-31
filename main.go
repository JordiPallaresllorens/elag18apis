package main

import (
	"log"
	"net/http"

	"github.com/cmh2166/elag18apis/aws_session"
	"github.com/cmh2166/elag18apis/config"
	"github.com/cmh2166/elag18apis/db"
	"github.com/cmh2166/elag18apis/generated/restapi"
	"github.com/cmh2166/elag18apis/generated/restapi/operations"
	"github.com/cmh2166/elag18apis/handlers"
	"github.com/cmh2166/elag18apis/identifier"
	"github.com/cmh2166/elag18apis/middleware"
	"github.com/justinas/alice"
)

func main() {
	// Initialize our global struct
	config := config.NewConfig()
	awsSession := aws_session.Connect(config.AwsDisableSSL)
	database := &db.DynamodbDatabase{
		Connection: db.Connect(awsSession, config.DynamodbEndpoint),
		Table:      config.ResourceTableName,
	}

	identifierService := identifier.NewService(config)
	server := createServer(database, identifierService, config.Port)
	defer server.Shutdown()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func createServer(database db.Database, identifierService identifier.Service, port int) *restapi.Server {
	api := handlers.BuildAPI(database, identifierService)
	server := restapi.NewServer(api)
	server.SetHandler(BuildHandler(api))

	// set the port this service will be run on
	server.Port = port
	return server
}

// BuildHandler sets up the middleware that wraps the API
func BuildHandler(api *operations.TaquitoAPI) http.Handler {
	return alice.New(
		middleware.NewRecoveryMW(),
		middleware.NewRequestLoggerMW(),
	).Then(api.Serve(nil))
}
