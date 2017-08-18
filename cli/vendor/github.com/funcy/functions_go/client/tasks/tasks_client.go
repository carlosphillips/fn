// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetTasks gets next task

Gets the next task in the queue, ready for processing. Consumers should start processing tasks in order. No other consumer can retrieve this task.
*/
func (a *Client) GetTasks(params *GetTasksParams) (*GetTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTasks",
		Method:             "GET",
		PathPattern:        "/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTasksOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
