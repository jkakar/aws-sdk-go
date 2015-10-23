// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package configservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// AWS Config provides a way to keep track of the configurations of all the
// AWS resources associated with your AWS account. You can use AWS Config to
// get the current and historical configurations of each AWS resource and also
// to get information about the relationship between the resources. An AWS resource
// can be an Amazon Compute Cloud (Amazon EC2) instance, an Elastic Block Store
// (EBS) volume, an Elastic network Interface (ENI), or a security group. For
// a complete list of resources currently supported by AWS Config, see Supported
// AWS Resources (http://docs.aws.amazon.com/config/latest/developerguide/resource-config-reference.html#supported-resources).
//
// You can access and manage AWS Config through the AWS Management Console,
// the AWS Command Line Interface (AWS CLI), the AWS Config API, or the AWS
// SDKs for AWS Config
//
// This reference guide contains documentation for the AWS Config API and the
// AWS CLI commands that you can use to manage AWS Config.
//
// The AWS Config API uses the Signature Version 4 protocol for signing requests.
// For more information about how to sign a request with this protocol, see
// Signature Version 4 Signing Process (http://docs.aws.amazon.com/general/latest/gr/signature-version-4.html).
//
// For detailed information about AWS Config features and their associated
// actions or commands, as well as how to work with AWS Management Console,
// see What Is AWS Config? (http://docs.aws.amazon.com/config/latest/developerguide/WhatIsConfig.html)
// in the AWS Config Developer Guide.
type ConfigService struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new ConfigService client.
func New(config *aws.Config) *ConfigService {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:       defaults.DefaultConfig.Merge(config),
			ServiceName:  "config",
			APIVersion:   "2014-11-12",
			JSONVersion:  "1.1",
			TargetPrefix: "StarlingDoveService",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &ConfigService{service}
}

// newRequest creates a new request for a ConfigService operation and runs any
// custom request initialization.
func (c *ConfigService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
