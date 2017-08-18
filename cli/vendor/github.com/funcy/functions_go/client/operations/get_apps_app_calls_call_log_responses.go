// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/models"
)

// GetAppsAppCallsCallLogReader is a Reader for the GetAppsAppCallsCallLog structure.
type GetAppsAppCallsCallLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAppsAppCallsCallLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAppsAppCallsCallLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetAppsAppCallsCallLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAppsAppCallsCallLogOK creates a GetAppsAppCallsCallLogOK with default headers values
func NewGetAppsAppCallsCallLogOK() *GetAppsAppCallsCallLogOK {
	return &GetAppsAppCallsCallLogOK{}
}

/*GetAppsAppCallsCallLogOK handles this case with default header values.

Log found
*/
type GetAppsAppCallsCallLogOK struct {
	Payload *models.LogWrapper
}

func (o *GetAppsAppCallsCallLogOK) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/{call}/log][%d] getAppsAppCallsCallLogOK  %+v", 200, o.Payload)
}

func (o *GetAppsAppCallsCallLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LogWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAppsAppCallsCallLogNotFound creates a GetAppsAppCallsCallLogNotFound with default headers values
func NewGetAppsAppCallsCallLogNotFound() *GetAppsAppCallsCallLogNotFound {
	return &GetAppsAppCallsCallLogNotFound{}
}

/*GetAppsAppCallsCallLogNotFound handles this case with default header values.

Log not found.
*/
type GetAppsAppCallsCallLogNotFound struct {
	Payload *models.Error
}

func (o *GetAppsAppCallsCallLogNotFound) Error() string {
	return fmt.Sprintf("[GET /apps/{app}/calls/{call}/log][%d] getAppsAppCallsCallLogNotFound  %+v", 404, o.Payload)
}

func (o *GetAppsAppCallsCallLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
