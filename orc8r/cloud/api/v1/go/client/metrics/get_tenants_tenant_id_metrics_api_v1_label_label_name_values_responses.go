// Code generated by go-swagger; DO NOT EDIT.

package metrics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesReader is a Reader for the GetTenantsTenantIDMetricsAPIV1LabelLabelNameValues structure.
type GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK creates a GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK with default headers values
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK() *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK {
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK{}
}

/*GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK handles this case with default header values.

List of label values
*/
type GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK struct {
	Payload *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/api/v1/label/{label_name}/values][%d] getTenantsTenantIdMetricsApiV1LabelLabelNameValuesOK  %+v", 200, o.Payload)
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK) GetPayload() *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault creates a GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault with default headers values
func NewGetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault(code int) *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault {
	return &GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault{
		_statusCode: code,
	}
}

/*GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault handles this case with default header values.

Unexpected Error
*/
type GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get tenants tenant ID metrics API v1 label label name values default response
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault) Error() string {
	return fmt.Sprintf("[GET /tenants/{tenant_id}/metrics/api/v1/label/{label_name}/values][%d] GetTenantsTenantIDMetricsAPIV1LabelLabelNameValues default  %+v", o._statusCode, o.Payload)
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody get tenants tenant ID metrics API v1 label label name values o k body
swagger:model GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody
*/
type GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody struct {

	// data
	// Required: true
	Data []string `json:"data"`

	// status
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this get tenants tenant ID metrics API v1 label label name values o k body
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getTenantsTenantIdMetricsApiV1LabelLabelNameValuesOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	return nil
}

func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("getTenantsTenantIdMetricsApiV1LabelLabelNameValuesOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody) UnmarshalBinary(b []byte) error {
	var res GetTenantsTenantIDMetricsAPIV1LabelLabelNameValuesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}