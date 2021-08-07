// Code generated by go-swagger; DO NOT EDIT.

package item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"go-server-server/lostfound/models"
)

// NewCreateItemParams creates a new CreateItemParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateItemParams() *CreateItemParams {
	return &CreateItemParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateItemParamsWithTimeout creates a new CreateItemParams object
// with the ability to set a timeout on a request.
func NewCreateItemParamsWithTimeout(timeout time.Duration) *CreateItemParams {
	return &CreateItemParams{
		timeout: timeout,
	}
}

// NewCreateItemParamsWithContext creates a new CreateItemParams object
// with the ability to set a context for a request.
func NewCreateItemParamsWithContext(ctx context.Context) *CreateItemParams {
	return &CreateItemParams{
		Context: ctx,
	}
}

// NewCreateItemParamsWithHTTPClient creates a new CreateItemParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateItemParamsWithHTTPClient(client *http.Client) *CreateItemParams {
	return &CreateItemParams{
		HTTPClient: client,
	}
}

/* CreateItemParams contains all the parameters to send to the API endpoint
   for the create item operation.

   Typically these are written to a http.Request.
*/
type CreateItemParams struct {

	// Item.
	Item *models.Item

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateItemParams) WithDefaults() *CreateItemParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create item params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateItemParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create item params
func (o *CreateItemParams) WithTimeout(timeout time.Duration) *CreateItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create item params
func (o *CreateItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create item params
func (o *CreateItemParams) WithContext(ctx context.Context) *CreateItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create item params
func (o *CreateItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create item params
func (o *CreateItemParams) WithHTTPClient(client *http.Client) *CreateItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create item params
func (o *CreateItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItem adds the item to the create item params
func (o *CreateItemParams) WithItem(item *models.Item) *CreateItemParams {
	o.SetItem(item)
	return o
}

// SetItem adds the item to the create item params
func (o *CreateItemParams) SetItem(item *models.Item) {
	o.Item = item
}

// WriteToRequest writes these params to a swagger request
func (o *CreateItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Item != nil {
		if err := r.SetBodyParam(o.Item); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
