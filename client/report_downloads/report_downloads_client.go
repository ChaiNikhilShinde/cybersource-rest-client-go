// Code generated by go-swagger; DO NOT EDIT.

package report_downloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new report downloads API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for report downloads API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DownloadReport downloads a report

Download a report using the unique report name and date.

*/
func (a *Client) DownloadReport(params *DownloadReportParams) (*DownloadReportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadReportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadReport",
		Method:             "GET",
		PathPattern:        "/reporting/v3/report-downloads",
		ProducesMediaTypes: []string{"application/xml", "text/csv"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadReportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DownloadReportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for downloadReport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}