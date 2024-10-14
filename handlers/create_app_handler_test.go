package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"multitenant-hosting/domain"
	error2 "multitenant-hosting/errors"
	"multitenant-hosting/handlers"
	"multitenant-hosting/service"
	"multitenant-hosting/service/mocks"
)

func TestCreateAppHandler(t *testing.T) {
	mockCreateAppSvc := new(mocks.CreateAppServiceMock)
	serviceRegistry := &service.Registry{
		CreateAppSvc: mockCreateAppSvc,
	}
	handler := handlers.CreateAppHandler(serviceRegistry)

	t.Run("app created successfully", func(t *testing.T) {
		mockResponse := &domain.AppResponse{
			Name: "test-app",
			Port: 1234,
			DNS:  "test-app.abc.xyz",
		}
		mockCreateAppSvc.On("CreateApp", mock.Anything, "test-app", "cloud").Return(mockResponse, nil)

		reqBody := domain.AppRequest{Name: "test-app", DeployType: "cloud"}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/app", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		handler(w, req)

		var response domain.AppResponse
		assert.Equal(t, http.StatusCreated, w.Code)
		json.NewDecoder(w.Body).Decode(&response)
		assert.Equal(t, *mockResponse, response)
		mockCreateAppSvc.AssertExpectations(t)
	})

	t.Run("invalid JSON request", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/app", bytes.NewReader([]byte("invalid json")))
		w := httptest.NewRecorder()

		handler(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("app already exists error", func(t *testing.T) {
		mockCreateAppSvc.On("CreateApp", mock.Anything, "test-2", "local").Return(nil, error2.ErrAppAlreadyExists)

		reqBody := domain.AppRequest{Name: "test-2", DeployType: "local"}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/app", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		handler(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, error2.ErrAppAlreadyExists.Error(), w.Header().Get(handlers.HeaderErrormessage))

		mockCreateAppSvc.AssertExpectations(t)
	})

	t.Run("service unavailable error", func(t *testing.T) {
		mockCreateAppSvc.On("CreateApp", mock.Anything, "test-3", "local").Return(nil, errors.New("unexpected error"))

		reqBody := domain.AppRequest{Name: "test-3", DeployType: "local"}
		reqBodyBytes, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/app", bytes.NewReader(reqBodyBytes))
		w := httptest.NewRecorder()

		handler(w, req)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
		assert.Equal(t, "unexpected error", w.Header().Get(handlers.HeaderErrormessage))

		mockCreateAppSvc.AssertExpectations(t)
	})
}
