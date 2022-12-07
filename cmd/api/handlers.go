package main

import (
	"net/http"
	"time"

	"github.com/see-air-uh/finn-inspector-gadget/data"
)

type JSONPayload struct {
	ID     string    `json:"id,omitempty"`
	User   string    `json:"user"`
	Date   time.Time `json:"timestamp"`
	Module string    `json:"module"`
	Event  string    `json:"event"`
	Action string    `json:"action"`
	Data   string    `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {

	// read the json into a var
	var requestPayload JSONPayload
	_ = app.readJSON(w, r, &requestPayload)

	// insert the data
	event := data.LogEntry{
		User:   requestPayload.User,
		Date:   requestPayload.Date,
		Module: requestPayload.Module,
		Event:  requestPayload.Event,
		Action: requestPayload.Action,
		Data:   requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "logged",
	}
	app.writeJSON(w, http.StatusAccepted, resp)
}
