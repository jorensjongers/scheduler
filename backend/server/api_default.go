/*
 * Scheduler API
 *
 * Basic API for retrieving schedules, based on a ScheduleInput object.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/jorensjongers/scheduler/backend/model"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{service: s}
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{
		{
			"ModelParametersGetGet",
			strings.ToUpper("Get"),
			"/model-parameters/get",
			c.ModelParametersGetGet,
		},
		{
			"ModelParametersSetPost",
			strings.ToUpper("Post"),
			"/model-parameters/set",
			c.ModelParametersSetPost,
		},
		{
			"ModelParametersSetOptions",
			strings.ToUpper("Options"),
			"/model-parameters/set",
			c.ModelParametersSetOptions,
		},
		{
			"InstanceDataGetGet",
			strings.ToUpper("Get"),
			"/instance-data/get",
			c.InstanceDataGetGet,
		},
		{
			"InstanceDataSetPost",
			strings.ToUpper("Post"),
			"/instance-data/set",
			c.InstanceDataSetPost,
		},
		{
			"InstanceDataSetOptions",
			strings.ToUpper("Options"),
			"/instance-data/set",
			c.InstanceDataSetOptions,
		},
		{
			"ScheduleGet",
			strings.ToUpper("Get"),
			"/schedule",
			c.ScheduleGet,
		},
		{
			"DbScheduleGet",
			strings.ToUpper("Get"),
			"/db-schedule",
			c.DbScheduleGet,
		},
	}
}

// ModelParametersGetGet - Returns the current model parameters.
func (c *DefaultApiController) ModelParametersGetGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	result, err := c.service.ModelParametersGetGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}

}

// ModelParametersSetPost - Sets the model paramters in the backend.
func (c *DefaultApiController) ModelParametersSetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Content-Type", "application/json")

	modelParameters := &model.ModelParameters{}
	if err := json.NewDecoder(r.Body).Decode(&modelParameters); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.ModelParametersSetPost(r.Context(), *modelParameters)
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}

}

// ModelParametersSetOptions
func (c *DefaultApiController) ModelParametersSetOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// ModelParametersSetOptions
func (c *DefaultApiController) InstanceDataSetOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, POST")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

// InstanceDataGetGet - Returns the current instance data.
func (c *DefaultApiController) InstanceDataGetGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	result, err := c.service.InstanceDataGetGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}

}

// InstanceDataSetPost - Sets the insatnce data in the backend.
func (c *DefaultApiController) InstanceDataSetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Content-Type", "application/json")

	instanceData := &model.InstanceData{}
	if err := json.NewDecoder(r.Body).Decode(&instanceData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.InstanceDataSetPost(r.Context(), *instanceData)
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}

}

// ScheduleGet - Returns a generated schedule.
func (c *DefaultApiController) ScheduleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")

	result, err := c.service.ScheduleGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}
}

// DbScheduleGet - Returns the schedule as found in the db.
func (c *DefaultApiController) DbScheduleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	result, err := c.service.DbScheduleGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		if err := EncodeJSONResponse(err.Error(), &result.Code, w); err != nil {
			panic(err)
		}
		return
	}
	//If no error, encode the body and the result code
	if err := EncodeJSONResponse(result.Body, &result.Code, w); err != nil {
		panic(err)
	}

}
