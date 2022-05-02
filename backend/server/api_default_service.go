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
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DennisMaes6/scheduler/backend/model"
	"github.com/DennisMaes6/scheduler/backend/schedule_generator"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	scheduleGenerator schedule_generator.ScheduleGenerator
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() *DefaultApiService {
	fmt.Println("NEW DEFAULT SERVICE")
	return &DefaultApiService{schedule_generator.NewScheduleGenerator("demo.db")}
}

// ScheduleGet - Returns a generated schedule.
func (s *DefaultApiService) ScheduleGet(ctx context.Context) (ImplResponse, error) {
	res, err := s.scheduleGenerator.GenerateSchedule()
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}

// ScheduleGet - Returns a generated schedule.
func (s *DefaultApiService) DbScheduleGet(ctx context.Context) (ImplResponse, error) {
	log.Printf("DbScheduleGet")
	res, err := s.scheduleGenerator.GetScheduleFromDb()
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}

// ScheduleGet - Returns a generated schedule.
func (s *DefaultApiService) FileScheduleGet(ctx context.Context) (ImplResponse, error) {
	bytes, err := ioutil.ReadFile("schedule.json")
	res := model.Schedule{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}

// Generate a new database and store in db file
func (s *DefaultApiService) GenerateScheduleGet(ctx context.Context) (ImplResponse, error) {
	res, err := s.scheduleGenerator.GenerateScheduleFromDb()
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}

// SetModelParamsPost - Sets the model paramters in the backend.
func (s *DefaultApiService) ModelParametersSetPost(ctx context.Context, modelParameters model.ModelParameters) (ImplResponse, error) {
	if err := s.scheduleGenerator.UpdateModelParameters(modelParameters); err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(204, nil), nil
}

// ScheduleGet - Returns a generated schedule.
func (s *DefaultApiService) ModelParametersGetGet(ctx context.Context) (ImplResponse, error) {
	res, err := s.scheduleGenerator.GetModelParameters()
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}

// SetModelParamsPost - Sets the model paramters in the backend.
func (s *DefaultApiService) InstanceDataSetPost(ctx context.Context, instanceData model.InstanceData) (ImplResponse, error) {
	if err := s.scheduleGenerator.UpdateInstanceData(instanceData); err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(204, nil), nil
}

// DbFilePost - Sets the db file in the backend.
func (s *DefaultApiService) DbFileSetPost(ctx context.Context, DbFile model.DbFile) (ImplResponse, error) {
	s.scheduleGenerator.UpdateDbFile(DbFile.Filename)
	return Response(204, nil), nil
}

// ScheduleGet - Returns a generated schedule.
func (s *DefaultApiService) InstanceDataGetGet(ctx context.Context) (ImplResponse, error) {
	res, err := s.scheduleGenerator.GetInstanceData()
	if err != nil {
		return Response(http.StatusInternalServerError, err.Error()), err
	}
	return Response(http.StatusOK, res), nil
}
