// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/query"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// This is the API Reference for Amazon Simple Email Service (Amazon SES). This
// documentation is intended to be used in conjunction with the Amazon SES Developer
// Guide (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/Welcome.html).
//
// For a list of Amazon SES endpoints to use in service requests, see Regions
// and Amazon SES (http://docs.aws.amazon.com/ses/latest/DeveloperGuide/regions.html)
// in the Amazon SES Developer Guide.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type SES struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "email"

// New creates a new instance of the SES client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a SES client from just a session.
//     svc := ses.New(mySession)
//
//     // Create a SES client with additional configuration
//     svc := ses.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *SES {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *SES {
	svc := &SES{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningName:   "ses",
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2010-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a SES operation and runs any
// custom request initialization.
func (c *SES) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
