// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Overview
//
// This is the AWS Lambda API Reference. The AWS Lambda Developer Guide provides
// additional information. For the service overview, go to What is AWS Lambda
// (http://docs.aws.amazon.com/lambda/latest/dg/welcome.html), and for information
// about how the service works, go to AWS Lambda: How it Works (http://docs.aws.amazon.com/lambda/latest/dg/lambda-introduction.html)
// in the AWS Lambda Developer Guide.
type Lambda struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new Lambda client.
func New(config *aws.Config) *Lambda {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "lambda",
			APIVersion:  "2015-03-31",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(restjson.Build)
	service.Handlers.Unmarshal.PushBack(restjson.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(restjson.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(restjson.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &Lambda{service}
}

// newRequest creates a new request for a Lambda operation and runs any
// custom request initialization.
func (c *Lambda) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
