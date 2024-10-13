package handlers

import (
	"encoding/json"
	"log"
	"multitenant-hosting/domain"
	"multitenant-hosting/errors"
	"multitenant-hosting/service"
	"net/http"
)

const (
	HeaderErrormessage = "Error-Message"
)

func CreateAppHandler(serviceRegistry *service.Registry) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req domain.AppRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Printf("Error decoding the request %s", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		resp, err := serviceRegistry.CreateAppSvc.CreateApp(r.Context(), req.Name, req.DeployType)
		if err != nil {
			switch err {
			case errors.ErrAppAlreadyExists:
				w.Header().Set(HeaderErrormessage, err.Error())
				w.WriteHeader(http.StatusBadRequest)
			default:
				w.Header().Set(HeaderErrormessage, err.Error())
				w.WriteHeader(http.StatusServiceUnavailable)
			}
			return
		}
		w.WriteHeader(http.StatusCreated)
		body, _ := json.Marshal(resp)
		w.Write(body)
		return
	}
}
