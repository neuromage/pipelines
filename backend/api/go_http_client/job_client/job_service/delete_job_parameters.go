// Code generated by go-swagger; DO NOT EDIT.

package job_service

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

// NewDeleteJobParams creates a new DeleteJobParams object
// with the default values initialized.
func NewDeleteJobParams() *DeleteJobParams {
	var ()
	return &DeleteJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteJobParamsWithTimeout creates a new DeleteJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteJobParamsWithTimeout(timeout time.Duration) *DeleteJobParams {
	var ()
	return &DeleteJobParams{

		timeout: timeout,
	}
}

// NewDeleteJobParamsWithContext creates a new DeleteJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteJobParamsWithContext(ctx context.Context) *DeleteJobParams {
	var ()
	return &DeleteJobParams{

		Context: ctx,
	}
}

// NewDeleteJobParamsWithHTTPClient creates a new DeleteJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteJobParamsWithHTTPClient(client *http.Client) *DeleteJobParams {
	var ()
	return &DeleteJobParams{
		HTTPClient: client,
	}
}

/*DeleteJobParams contains all the parameters to send to the API endpoint
for the delete job operation typically these are written to a http.Request
*/
type DeleteJobParams struct {

	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete job params
func (o *DeleteJobParams) WithTimeout(timeout time.Duration) *DeleteJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete job params
func (o *DeleteJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete job params
func (o *DeleteJobParams) WithContext(ctx context.Context) *DeleteJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete job params
func (o *DeleteJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete job params
func (o *DeleteJobParams) WithHTTPClient(client *http.Client) *DeleteJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete job params
func (o *DeleteJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete job params
func (o *DeleteJobParams) WithID(id string) *DeleteJobParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete job params
func (o *DeleteJobParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
