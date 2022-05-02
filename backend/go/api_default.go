/*
 * Scheduler API
 *
 * API for getting generated schedules. Also used for getting and setting model parameters and instance data.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"DbScheduleGet",
			strings.ToUpper("Get"),
			"/backend/db-schedule",
			c.DbScheduleGet,
		},
		{
			"FileScheduleGet",
			strings.ToUpper("Get"),
			"/backend/file-schedule",
			c.FileScheduleGet,
		},
		{
			"InstanceDataGetGet",
			strings.ToUpper("Get"),
			"/backend/instance-data/get",
			c.InstanceDataGetGet,
		},
		{
			"InstanceDataSetPost",
			strings.ToUpper("Post"),
			"/backend/instance-data/set",
			c.InstanceDataSetPost,
		},
		{
			"ModelParametersGetGet",
			strings.ToUpper("Get"),
			"/backend/model-parameters/get",
			c.ModelParametersGetGet,
		},
		{
			"ModelParametersSetPost",
			strings.ToUpper("Post"),
			"/backend/model-parameters/set",
			c.ModelParametersSetPost,
		},
		{
			"ScheduleFileSetPost",
			strings.ToUpper("Post"),
			"/backend/schedule-file/set",
			c.ScheduleFileSetPost,
		},
		{
			"ScheduleGenerateGet",
			strings.ToUpper("Get"),
			"/backend/schedule-generate",
			c.ScheduleGenerateGet,
		},
		{
			"ScheduleGet",
			strings.ToUpper("Get"),
			"/backend/schedule",
			c.ScheduleGet,
		},
	}
}

// DbScheduleGet - Returns the schedule as found in the db.
func (c *DefaultApiController) DbScheduleGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.DbScheduleGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// FileScheduleGet - Returns the schedule as found in a file.
func (c *DefaultApiController) FileScheduleGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.FileScheduleGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// InstanceDataGetGet - Returns the current instance data.
func (c *DefaultApiController) InstanceDataGetGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.InstanceDataGetGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// InstanceDataSetPost - Sets the insatnce data in the backend.
func (c *DefaultApiController) InstanceDataSetPost(w http.ResponseWriter, r *http.Request) {
	instanceDataParam := InstanceData{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&instanceDataParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertInstanceDataRequired(instanceDataParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.InstanceDataSetPost(r.Context(), instanceDataParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ModelParametersGetGet - Returns the current model parameters.
func (c *DefaultApiController) ModelParametersGetGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ModelParametersGetGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ModelParametersSetPost - Sets the model paramters in the backend.
func (c *DefaultApiController) ModelParametersSetPost(w http.ResponseWriter, r *http.Request) {
	modelParametersParam := ModelParameters{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&modelParametersParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertModelParametersRequired(modelParametersParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ModelParametersSetPost(r.Context(), modelParametersParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ScheduleFileSetPost - Sets the schedule DB file in the backend.
func (c *DefaultApiController) ScheduleFileSetPost(w http.ResponseWriter, r *http.Request) {
	dbFileParam := DbFile{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dbFileParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDbFileRequired(dbFileParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ScheduleFileSetPost(r.Context(), dbFileParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ScheduleGenerateGet - Returns a schedule generated with Java.
func (c *DefaultApiController) ScheduleGenerateGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ScheduleGenerateGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ScheduleGet - Returns a schedule generated with MiniZinc.
func (c *DefaultApiController) ScheduleGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.ScheduleGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
