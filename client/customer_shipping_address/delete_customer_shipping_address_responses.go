// Code generated by go-swagger; DO NOT EDIT.

package customer_shipping_address

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeleteCustomerShippingAddressReader is a Reader for the DeleteCustomerShippingAddress structure.
type DeleteCustomerShippingAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomerShippingAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCustomerShippingAddressNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomerShippingAddressBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCustomerShippingAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomerShippingAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewDeleteCustomerShippingAddressFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCustomerShippingAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCustomerShippingAddressNoContent creates a DeleteCustomerShippingAddressNoContent with default headers values
func NewDeleteCustomerShippingAddressNoContent() *DeleteCustomerShippingAddressNoContent {
	return &DeleteCustomerShippingAddressNoContent{}
}

/*DeleteCustomerShippingAddressNoContent handles this case with default header values.

The server fulfilled the request but does not need to return a body
*/
type DeleteCustomerShippingAddressNoContent struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string
}

func (o *DeleteCustomerShippingAddressNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressNoContent ", 204)
}

func (o *DeleteCustomerShippingAddressNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	return nil
}

// NewDeleteCustomerShippingAddressBadRequest creates a DeleteCustomerShippingAddressBadRequest with default headers values
func NewDeleteCustomerShippingAddressBadRequest() *DeleteCustomerShippingAddressBadRequest {
	return &DeleteCustomerShippingAddressBadRequest{}
}

/*DeleteCustomerShippingAddressBadRequest handles this case with default header values.

Bad Request: e.g. A required header value could be missing.
*/
type DeleteCustomerShippingAddressBadRequest struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string

	Payload *DeleteCustomerShippingAddressBadRequestBody
}

func (o *DeleteCustomerShippingAddressBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCustomerShippingAddressBadRequest) GetPayload() *DeleteCustomerShippingAddressBadRequestBody {
	return o.Payload
}

func (o *DeleteCustomerShippingAddressBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	o.Payload = new(DeleteCustomerShippingAddressBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerShippingAddressForbidden creates a DeleteCustomerShippingAddressForbidden with default headers values
func NewDeleteCustomerShippingAddressForbidden() *DeleteCustomerShippingAddressForbidden {
	return &DeleteCustomerShippingAddressForbidden{}
}

/*DeleteCustomerShippingAddressForbidden handles this case with default header values.

403ForbiddenResponse: e.g. The profile might not have permission to perform the operation.
*/
type DeleteCustomerShippingAddressForbidden struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string

	Payload *DeleteCustomerShippingAddressForbiddenBody
}

func (o *DeleteCustomerShippingAddressForbidden) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomerShippingAddressForbidden) GetPayload() *DeleteCustomerShippingAddressForbiddenBody {
	return o.Payload
}

func (o *DeleteCustomerShippingAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	o.Payload = new(DeleteCustomerShippingAddressForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerShippingAddressNotFound creates a DeleteCustomerShippingAddressNotFound with default headers values
func NewDeleteCustomerShippingAddressNotFound() *DeleteCustomerShippingAddressNotFound {
	return &DeleteCustomerShippingAddressNotFound{}
}

/*DeleteCustomerShippingAddressNotFound handles this case with default header values.

Token Not Found. The `tokenid` may not exist or was entered incorrectly.
*/
type DeleteCustomerShippingAddressNotFound struct {
	/*A globally unique ID associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string

	Payload *DeleteCustomerShippingAddressNotFoundBody
}

func (o *DeleteCustomerShippingAddressNotFound) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomerShippingAddressNotFound) GetPayload() *DeleteCustomerShippingAddressNotFoundBody {
	return o.Payload
}

func (o *DeleteCustomerShippingAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	o.Payload = new(DeleteCustomerShippingAddressNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerShippingAddressFailedDependency creates a DeleteCustomerShippingAddressFailedDependency with default headers values
func NewDeleteCustomerShippingAddressFailedDependency() *DeleteCustomerShippingAddressFailedDependency {
	return &DeleteCustomerShippingAddressFailedDependency{}
}

/*DeleteCustomerShippingAddressFailedDependency handles this case with default header values.

Failed Dependency: e.g. The profile represented by the profile-id may not exist or the profile-id was entered incorrectly.
*/
type DeleteCustomerShippingAddressFailedDependency struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string

	Payload *DeleteCustomerShippingAddressFailedDependencyBody
}

func (o *DeleteCustomerShippingAddressFailedDependency) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressFailedDependency  %+v", 424, o.Payload)
}

func (o *DeleteCustomerShippingAddressFailedDependency) GetPayload() *DeleteCustomerShippingAddressFailedDependencyBody {
	return o.Payload
}

func (o *DeleteCustomerShippingAddressFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	o.Payload = new(DeleteCustomerShippingAddressFailedDependencyBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomerShippingAddressInternalServerError creates a DeleteCustomerShippingAddressInternalServerError with default headers values
func NewDeleteCustomerShippingAddressInternalServerError() *DeleteCustomerShippingAddressInternalServerError {
	return &DeleteCustomerShippingAddressInternalServerError{}
}

/*DeleteCustomerShippingAddressInternalServerError handles this case with default header values.

Unexpected error.
*/
type DeleteCustomerShippingAddressInternalServerError struct {
	/*A globally unique id associated with your request.
	 */
	UniqueTransactionID string
	/*The mandatory correlation id passed by upstream (calling) system.
	 */
	VcCorrelationID string

	Payload *DeleteCustomerShippingAddressInternalServerErrorBody
}

func (o *DeleteCustomerShippingAddressInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /tms/v2/customers/{customerTokenId}/shipping-addresses/{shippingAddressTokenId}][%d] deleteCustomerShippingAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCustomerShippingAddressInternalServerError) GetPayload() *DeleteCustomerShippingAddressInternalServerErrorBody {
	return o.Payload
}

func (o *DeleteCustomerShippingAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header uniqueTransactionID
	o.UniqueTransactionID = response.GetHeader("uniqueTransactionID")

	// response header v-c-correlation-id
	o.VcCorrelationID = response.GetHeader("v-c-correlation-id")

	o.Payload = new(DeleteCustomerShippingAddressInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DeleteCustomerShippingAddressBadRequestBody delete customer shipping address bad request body
swagger:model DeleteCustomerShippingAddressBadRequestBody
*/
type DeleteCustomerShippingAddressBadRequestBody struct {

	// errors
	// Read Only: true
	Errors []*DeleteCustomerShippingAddressBadRequestBodyErrorsItems0 `json:"errors"`
}

// Validate validates this delete customer shipping address bad request body
func (o *DeleteCustomerShippingAddressBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressBadRequestBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteCustomerShippingAddressBadRequest" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressBadRequestBodyErrorsItems0 delete customer shipping address bad request body errors items0
swagger:model DeleteCustomerShippingAddressBadRequestBodyErrorsItems0
*/
type DeleteCustomerShippingAddressBadRequestBodyErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this delete customer shipping address bad request body errors items0
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressBadRequestBodyErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0 delete customer shipping address bad request body errors items0 details items0
swagger:model DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0
*/
type DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this delete customer shipping address bad request body errors items0 details items0
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressBadRequestBodyErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressFailedDependencyBody delete customer shipping address failed dependency body
swagger:model DeleteCustomerShippingAddressFailedDependencyBody
*/
type DeleteCustomerShippingAddressFailedDependencyBody struct {

	// errors
	// Read Only: true
	Errors []*DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0 `json:"errors"`
}

// Validate validates this delete customer shipping address failed dependency body
func (o *DeleteCustomerShippingAddressFailedDependencyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressFailedDependencyBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteCustomerShippingAddressFailedDependency" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressFailedDependencyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0 delete customer shipping address failed dependency body errors items0
swagger:model DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0
*/
type DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this delete customer shipping address failed dependency body errors items0
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0 delete customer shipping address failed dependency body errors items0 details items0
swagger:model DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0
*/
type DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this delete customer shipping address failed dependency body errors items0 details items0
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressFailedDependencyBodyErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressForbiddenBody delete customer shipping address forbidden body
swagger:model DeleteCustomerShippingAddressForbiddenBody
*/
type DeleteCustomerShippingAddressForbiddenBody struct {

	// errors
	// Read Only: true
	Errors []*DeleteCustomerShippingAddressForbiddenBodyErrorsItems0 `json:"errors"`
}

// Validate validates this delete customer shipping address forbidden body
func (o *DeleteCustomerShippingAddressForbiddenBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressForbiddenBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteCustomerShippingAddressForbidden" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressForbiddenBodyErrorsItems0 delete customer shipping address forbidden body errors items0
swagger:model DeleteCustomerShippingAddressForbiddenBodyErrorsItems0
*/
type DeleteCustomerShippingAddressForbiddenBodyErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this delete customer shipping address forbidden body errors items0
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressForbiddenBodyErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0 delete customer shipping address forbidden body errors items0 details items0
swagger:model DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0
*/
type DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this delete customer shipping address forbidden body errors items0 details items0
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressForbiddenBodyErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressInternalServerErrorBody delete customer shipping address internal server error body
swagger:model DeleteCustomerShippingAddressInternalServerErrorBody
*/
type DeleteCustomerShippingAddressInternalServerErrorBody struct {

	// errors
	// Read Only: true
	Errors []*DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0 `json:"errors"`
}

// Validate validates this delete customer shipping address internal server error body
func (o *DeleteCustomerShippingAddressInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressInternalServerErrorBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteCustomerShippingAddressInternalServerError" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0 delete customer shipping address internal server error body errors items0
swagger:model DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0
*/
type DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this delete customer shipping address internal server error body errors items0
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0 delete customer shipping address internal server error body errors items0 details items0
swagger:model DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0
*/
type DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this delete customer shipping address internal server error body errors items0 details items0
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressInternalServerErrorBodyErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressNotFoundBody delete customer shipping address not found body
swagger:model DeleteCustomerShippingAddressNotFoundBody
*/
type DeleteCustomerShippingAddressNotFoundBody struct {

	// errors
	// Read Only: true
	Errors []*DeleteCustomerShippingAddressNotFoundBodyErrorsItems0 `json:"errors"`
}

// Validate validates this delete customer shipping address not found body
func (o *DeleteCustomerShippingAddressNotFoundBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressNotFoundBody) validateErrors(formats strfmt.Registry) error {

	if swag.IsZero(o.Errors) { // not required
		return nil
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deleteCustomerShippingAddressNotFound" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressNotFoundBodyErrorsItems0 delete customer shipping address not found body errors items0
swagger:model DeleteCustomerShippingAddressNotFoundBodyErrorsItems0
*/
type DeleteCustomerShippingAddressNotFoundBodyErrorsItems0 struct {

	// details
	// Read Only: true
	Details []*DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0 `json:"details"`

	// The detailed message related to the type stated above.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The type of error.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this delete customer shipping address not found body errors items0
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressNotFoundBodyErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0 delete customer shipping address not found body errors items0 details items0
swagger:model DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0
*/
type DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0 struct {

	// The location of the field that caused the error.
	// Read Only: true
	Location string `json:"location,omitempty"`

	// The name of the field that caused the error.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this delete customer shipping address not found body errors items0 details items0
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteCustomerShippingAddressNotFoundBodyErrorsItems0DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}