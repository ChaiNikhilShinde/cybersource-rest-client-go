// Code generated by go-swagger; DO NOT EDIT.

package payment_instrument

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCreatePaymentInstrumentParams creates a new CreatePaymentInstrumentParams object
// with the default values initialized.
func NewCreatePaymentInstrumentParams() *CreatePaymentInstrumentParams {
	var ()
	return &CreatePaymentInstrumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentInstrumentParamsWithTimeout creates a new CreatePaymentInstrumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePaymentInstrumentParamsWithTimeout(timeout time.Duration) *CreatePaymentInstrumentParams {
	var ()
	return &CreatePaymentInstrumentParams{

		timeout: timeout,
	}
}

// NewCreatePaymentInstrumentParamsWithContext creates a new CreatePaymentInstrumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePaymentInstrumentParamsWithContext(ctx context.Context) *CreatePaymentInstrumentParams {
	var ()
	return &CreatePaymentInstrumentParams{

		Context: ctx,
	}
}

// NewCreatePaymentInstrumentParamsWithHTTPClient creates a new CreatePaymentInstrumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePaymentInstrumentParamsWithHTTPClient(client *http.Client) *CreatePaymentInstrumentParams {
	var ()
	return &CreatePaymentInstrumentParams{
		HTTPClient: client,
	}
}

/*CreatePaymentInstrumentParams contains all the parameters to send to the API endpoint
for the create payment instrument operation typically these are written to a http.Request
*/
type CreatePaymentInstrumentParams struct {

	/*CreatePaymentInstrumentRequest
	  Specify the customer's payment details for card or bank account.

	*/
	CreatePaymentInstrumentRequest CreatePaymentInstrumentBody
	/*ProfileID
	  The id of a profile containing user specific TMS configuration.

	*/
	ProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create payment instrument params
func (o *CreatePaymentInstrumentParams) WithTimeout(timeout time.Duration) *CreatePaymentInstrumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment instrument params
func (o *CreatePaymentInstrumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment instrument params
func (o *CreatePaymentInstrumentParams) WithContext(ctx context.Context) *CreatePaymentInstrumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment instrument params
func (o *CreatePaymentInstrumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment instrument params
func (o *CreatePaymentInstrumentParams) WithHTTPClient(client *http.Client) *CreatePaymentInstrumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment instrument params
func (o *CreatePaymentInstrumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreatePaymentInstrumentRequest adds the createPaymentInstrumentRequest to the create payment instrument params
func (o *CreatePaymentInstrumentParams) WithCreatePaymentInstrumentRequest(createPaymentInstrumentRequest CreatePaymentInstrumentBody) *CreatePaymentInstrumentParams {
	o.SetCreatePaymentInstrumentRequest(createPaymentInstrumentRequest)
	return o
}

// SetCreatePaymentInstrumentRequest adds the createPaymentInstrumentRequest to the create payment instrument params
func (o *CreatePaymentInstrumentParams) SetCreatePaymentInstrumentRequest(createPaymentInstrumentRequest CreatePaymentInstrumentBody) {
	o.CreatePaymentInstrumentRequest = createPaymentInstrumentRequest
}

// WithProfileID adds the profileID to the create payment instrument params
func (o *CreatePaymentInstrumentParams) WithProfileID(profileID string) *CreatePaymentInstrumentParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the create payment instrument params
func (o *CreatePaymentInstrumentParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentInstrumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.CreatePaymentInstrumentRequest); err != nil {
		return err
	}

	// header param profile-id
	if err := r.SetHeaderParam("profile-id", o.ProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}