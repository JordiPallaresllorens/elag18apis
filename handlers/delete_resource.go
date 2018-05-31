package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/cmh2166/elag18apis/db"
	"github.com/cmh2166/elag18apis/generated/restapi/operations"
)

// NewDeleteResource -- Accepts requests to remove a resource.
func NewDeleteResource(repository db.Database) operations.DeleteResourceHandler {
	return &deleteResourceEntry{repository: repository}
}

type deleteResourceEntry struct {
	repository db.Database
}

// Handle the delete entry request.
func (d *deleteResourceEntry) Handle(params operations.DeleteResourceParams) middleware.Responder {
	if err := d.DeleteAllVersions(params.ID); err != nil {
		panic(err)
	}

	return operations.NewDeleteResourceNoContent()
}

// DeleteAllVersions removes all versions with the given external id
func (d *deleteResourceEntry) DeleteAllVersions(externalID string) error {
	resource, err := d.repository.RetrieveLatest(externalID)
	if err != nil {
		if _, ok := err.(*db.RecordNotFound); !ok {
			panic(err)
		}
	}
	// Delete all versions of the resource
	for resource != nil {
		if err = d.repository.DeleteByID(resource.ID()); err != nil {
			panic(err)
		}

		// retrieve the next resource
		resource, err = d.repository.RetrieveLatest(externalID)
		if err != nil {
			if _, ok := err.(*db.RecordNotFound); !ok {
				panic(err)
			}
		}

	}
	return nil
}
