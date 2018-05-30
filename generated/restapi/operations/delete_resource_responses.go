// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteResourceNoContentCode is the HTTP code returned for type DeleteResourceNoContent
const DeleteResourceNoContentCode int = 204

/*DeleteResourceNoContent TAQUITO resource metadata delete.

swagger:response deleteResourceNoContent
*/
type DeleteResourceNoContent struct {
}

// NewDeleteResourceNoContent creates DeleteResourceNoContent with default headers values
func NewDeleteResourceNoContent() *DeleteResourceNoContent {
	return &DeleteResourceNoContent{}
}

// WriteResponse to the client
func (o *DeleteResourceNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteResourceUnauthorizedCode is the HTTP code returned for type DeleteResourceUnauthorized
const DeleteResourceUnauthorizedCode int = 401

/*DeleteResourceUnauthorized You are not authorized to delete a resource in TAQUITO.

swagger:response deleteResourceUnauthorized
*/
type DeleteResourceUnauthorized struct {
}

// NewDeleteResourceUnauthorized creates DeleteResourceUnauthorized with default headers values
func NewDeleteResourceUnauthorized() *DeleteResourceUnauthorized {
	return &DeleteResourceUnauthorized{}
}

// WriteResponse to the client
func (o *DeleteResourceUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// DeleteResourceNotFoundCode is the HTTP code returned for type DeleteResourceNotFound
const DeleteResourceNotFoundCode int = 404

/*DeleteResourceNotFound Invalid ID supplied

swagger:response deleteResourceNotFound
*/
type DeleteResourceNotFound struct {
}

// NewDeleteResourceNotFound creates DeleteResourceNotFound with default headers values
func NewDeleteResourceNotFound() *DeleteResourceNotFound {
	return &DeleteResourceNotFound{}
}

// WriteResponse to the client
func (o *DeleteResourceNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteResourceInternalServerErrorCode is the HTTP code returned for type DeleteResourceInternalServerError
const DeleteResourceInternalServerErrorCode int = 500

/*DeleteResourceInternalServerError This resource could not be deleted at this time by TAQUITO.

swagger:response deleteResourceInternalServerError
*/
type DeleteResourceInternalServerError struct {
}

// NewDeleteResourceInternalServerError creates DeleteResourceInternalServerError with default headers values
func NewDeleteResourceInternalServerError() *DeleteResourceInternalServerError {
	return &DeleteResourceInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteResourceInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
