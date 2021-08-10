// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package finspacedata

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
)

const opCreateChangeset = "CreateChangeset"

// CreateChangesetRequest generates a "aws/request.Request" representing the
// client's request for the CreateChangeset operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateChangeset for more information on using the CreateChangeset
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateChangesetRequest method.
//    req, resp := client.CreateChangesetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/CreateChangeset
func (c *FinSpaceData) CreateChangesetRequest(input *CreateChangesetInput) (req *request.Request, output *CreateChangesetOutput) {
	op := &request.Operation{
		Name:       opCreateChangeset,
		HTTPMethod: "POST",
		HTTPPath:   "/datasets/{datasetId}/changesets",
	}

	if input == nil {
		input = &CreateChangesetInput{}
	}

	output = &CreateChangesetOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateChangeset API operation for FinSpace Public API.
//
// Creates a new changeset in a FinSpace dataset.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for FinSpace Public API's
// API operation CreateChangeset for usage and error information.
//
// Returned Error Types:
//   * ResourceNotFoundException
//   One or more resources can't be found.
//
//   * InternalServerException
//   The request processing has failed because of an unknown error, exception
//   or failure.
//
//   * ValidationException
//   The input fails to satisfy the constraints specified by an AWS service.
//
//   * ThrottlingException
//   The request was denied due to request throttling.
//
//   * AccessDeniedException
//   You do not have sufficient access to perform this action.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/CreateChangeset
func (c *FinSpaceData) CreateChangeset(input *CreateChangesetInput) (*CreateChangesetOutput, error) {
	req, out := c.CreateChangesetRequest(input)
	return out, req.Send()
}

// CreateChangesetWithContext is the same as CreateChangeset with the addition of
// the ability to pass a context and additional request options.
//
// See CreateChangeset for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FinSpaceData) CreateChangesetWithContext(ctx aws.Context, input *CreateChangesetInput, opts ...request.Option) (*CreateChangesetOutput, error) {
	req, out := c.CreateChangesetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetProgrammaticAccessCredentials = "GetProgrammaticAccessCredentials"

// GetProgrammaticAccessCredentialsRequest generates a "aws/request.Request" representing the
// client's request for the GetProgrammaticAccessCredentials operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetProgrammaticAccessCredentials for more information on using the GetProgrammaticAccessCredentials
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetProgrammaticAccessCredentialsRequest method.
//    req, resp := client.GetProgrammaticAccessCredentialsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/GetProgrammaticAccessCredentials
func (c *FinSpaceData) GetProgrammaticAccessCredentialsRequest(input *GetProgrammaticAccessCredentialsInput) (req *request.Request, output *GetProgrammaticAccessCredentialsOutput) {
	op := &request.Operation{
		Name:       opGetProgrammaticAccessCredentials,
		HTTPMethod: "GET",
		HTTPPath:   "/credentials/programmatic",
	}

	if input == nil {
		input = &GetProgrammaticAccessCredentialsInput{}
	}

	output = &GetProgrammaticAccessCredentialsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetProgrammaticAccessCredentials API operation for FinSpace Public API.
//
// Request programmatic credentials to use with Habanero SDK.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for FinSpace Public API's
// API operation GetProgrammaticAccessCredentials for usage and error information.
//
// Returned Error Types:
//   * InternalServerException
//   The request processing has failed because of an unknown error, exception
//   or failure.
//
//   * ThrottlingException
//   The request was denied due to request throttling.
//
//   * AccessDeniedException
//   You do not have sufficient access to perform this action.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/GetProgrammaticAccessCredentials
func (c *FinSpaceData) GetProgrammaticAccessCredentials(input *GetProgrammaticAccessCredentialsInput) (*GetProgrammaticAccessCredentialsOutput, error) {
	req, out := c.GetProgrammaticAccessCredentialsRequest(input)
	return out, req.Send()
}

// GetProgrammaticAccessCredentialsWithContext is the same as GetProgrammaticAccessCredentials with the addition of
// the ability to pass a context and additional request options.
//
// See GetProgrammaticAccessCredentials for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FinSpaceData) GetProgrammaticAccessCredentialsWithContext(ctx aws.Context, input *GetProgrammaticAccessCredentialsInput, opts ...request.Option) (*GetProgrammaticAccessCredentialsOutput, error) {
	req, out := c.GetProgrammaticAccessCredentialsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetWorkingLocation = "GetWorkingLocation"

// GetWorkingLocationRequest generates a "aws/request.Request" representing the
// client's request for the GetWorkingLocation operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetWorkingLocation for more information on using the GetWorkingLocation
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetWorkingLocationRequest method.
//    req, resp := client.GetWorkingLocationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/GetWorkingLocation
func (c *FinSpaceData) GetWorkingLocationRequest(input *GetWorkingLocationInput) (req *request.Request, output *GetWorkingLocationOutput) {
	op := &request.Operation{
		Name:       opGetWorkingLocation,
		HTTPMethod: "POST",
		HTTPPath:   "/workingLocationV1",
	}

	if input == nil {
		input = &GetWorkingLocationInput{}
	}

	output = &GetWorkingLocationOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetWorkingLocation API operation for FinSpace Public API.
//
// A temporary Amazon S3 location to copy your files from a source location
// to stage or use as a scratch space in Habanero notebook.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for FinSpace Public API's
// API operation GetWorkingLocation for usage and error information.
//
// Returned Error Types:
//   * InternalServerException
//   The request processing has failed because of an unknown error, exception
//   or failure.
//
//   * AccessDeniedException
//   You do not have sufficient access to perform this action.
//
//   * ThrottlingException
//   The request was denied due to request throttling.
//
//   * ValidationException
//   The input fails to satisfy the constraints specified by an AWS service.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/finspace-2020-07-13/GetWorkingLocation
func (c *FinSpaceData) GetWorkingLocation(input *GetWorkingLocationInput) (*GetWorkingLocationOutput, error) {
	req, out := c.GetWorkingLocationRequest(input)
	return out, req.Send()
}

// GetWorkingLocationWithContext is the same as GetWorkingLocation with the addition of
// the ability to pass a context and additional request options.
//
// See GetWorkingLocation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FinSpaceData) GetWorkingLocationWithContext(ctx aws.Context, input *GetWorkingLocationInput, opts ...request.Option) (*GetWorkingLocationOutput, error) {
	req, out := c.GetWorkingLocationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// You do not have sufficient access to perform this action.
type AccessDeniedException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s AccessDeniedException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessDeniedException) GoString() string {
	return s.String()
}

func newErrorAccessDeniedException(v protocol.ResponseMetadata) error {
	return &AccessDeniedException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *AccessDeniedException) Code() string {
	return "AccessDeniedException"
}

// Message returns the exception's message.
func (s *AccessDeniedException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *AccessDeniedException) OrigErr() error {
	return nil
}

func (s *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *AccessDeniedException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *AccessDeniedException) RequestID() string {
	return s.RespMetadata.RequestID
}

// A changeset is unit of data in a dataset.
type ChangesetInfo struct {
	_ struct{} `type:"structure"`

	// Change type indicates how a changeset is applied to a dataset.
	//
	//    * REPLACE - Changeset is considered as a replacement to all prior loaded
	//    changesets.
	//
	//    * APPEND - Changeset is considered as an addition to the end of all prior
	//    loaded changesets.
	//
	//    * MODIFY - Changeset is considered as a replacement to a specific prior
	//    ingested changeset.
	ChangeType *string `locationName:"changeType" type:"string" enum:"ChangeType"`

	// The ARN identifier of the changeset.
	ChangesetArn *string `locationName:"changesetArn" min:"20" type:"string"`

	// Tags associated with the changeset.
	ChangesetLabels map[string]*string `locationName:"changesetLabels" type:"map"`

	// The timestamp at which the changeset was created in FinSpace.
	CreateTimestamp *time.Time `locationName:"createTimestamp" type:"timestamp"`

	// The unique identifier for the FinSpace dataset in which the changeset is
	// created.
	DatasetId *string `locationName:"datasetId" min:"1" type:"string"`

	// The structure with error messages.
	ErrorInfo *ErrorInfo `locationName:"errorInfo" type:"structure"`

	// Structure of the source file(s).
	FormatParams map[string]*string `locationName:"formatParams" type:"map"`

	// Format type of the input files loaded into the changeset.
	FormatType *string `locationName:"formatType" type:"string" enum:"FormatType"`

	// Unique identifier for a changeset.
	Id *string `locationName:"id" min:"1" type:"string"`

	// Source path from which the files to create the changeset are sourced.
	SourceParams map[string]*string `locationName:"sourceParams" type:"map"`

	// Type of the data source from which the files to create the changeset are
	// sourced.
	//
	//    * S3 - Amazon S3.
	SourceType *string `locationName:"sourceType" type:"string" enum:"SourceType"`

	// The status of changeset creation operation.
	Status *string `locationName:"status" type:"string" enum:"ChangesetStatus"`

	// Unique identifier of the changeset that is updated a changeset.
	UpdatedByChangesetId *string `locationName:"updatedByChangesetId" type:"string"`

	// Unique identifier of the changeset that is updated.
	UpdatesChangesetId *string `locationName:"updatesChangesetId" type:"string"`
}

// String returns the string representation
func (s ChangesetInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangesetInfo) GoString() string {
	return s.String()
}

// SetChangeType sets the ChangeType field's value.
func (s *ChangesetInfo) SetChangeType(v string) *ChangesetInfo {
	s.ChangeType = &v
	return s
}

// SetChangesetArn sets the ChangesetArn field's value.
func (s *ChangesetInfo) SetChangesetArn(v string) *ChangesetInfo {
	s.ChangesetArn = &v
	return s
}

// SetChangesetLabels sets the ChangesetLabels field's value.
func (s *ChangesetInfo) SetChangesetLabels(v map[string]*string) *ChangesetInfo {
	s.ChangesetLabels = v
	return s
}

// SetCreateTimestamp sets the CreateTimestamp field's value.
func (s *ChangesetInfo) SetCreateTimestamp(v time.Time) *ChangesetInfo {
	s.CreateTimestamp = &v
	return s
}

// SetDatasetId sets the DatasetId field's value.
func (s *ChangesetInfo) SetDatasetId(v string) *ChangesetInfo {
	s.DatasetId = &v
	return s
}

// SetErrorInfo sets the ErrorInfo field's value.
func (s *ChangesetInfo) SetErrorInfo(v *ErrorInfo) *ChangesetInfo {
	s.ErrorInfo = v
	return s
}

// SetFormatParams sets the FormatParams field's value.
func (s *ChangesetInfo) SetFormatParams(v map[string]*string) *ChangesetInfo {
	s.FormatParams = v
	return s
}

// SetFormatType sets the FormatType field's value.
func (s *ChangesetInfo) SetFormatType(v string) *ChangesetInfo {
	s.FormatType = &v
	return s
}

// SetId sets the Id field's value.
func (s *ChangesetInfo) SetId(v string) *ChangesetInfo {
	s.Id = &v
	return s
}

// SetSourceParams sets the SourceParams field's value.
func (s *ChangesetInfo) SetSourceParams(v map[string]*string) *ChangesetInfo {
	s.SourceParams = v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *ChangesetInfo) SetSourceType(v string) *ChangesetInfo {
	s.SourceType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ChangesetInfo) SetStatus(v string) *ChangesetInfo {
	s.Status = &v
	return s
}

// SetUpdatedByChangesetId sets the UpdatedByChangesetId field's value.
func (s *ChangesetInfo) SetUpdatedByChangesetId(v string) *ChangesetInfo {
	s.UpdatedByChangesetId = &v
	return s
}

// SetUpdatesChangesetId sets the UpdatesChangesetId field's value.
func (s *ChangesetInfo) SetUpdatesChangesetId(v string) *ChangesetInfo {
	s.UpdatesChangesetId = &v
	return s
}

type CreateChangesetInput struct {
	_ struct{} `type:"structure"`

	// Option to indicate how a changeset will be applied to a dataset.
	//
	//    * REPLACE - Changeset will be considered as a replacement to all prior
	//    loaded changesets.
	//
	//    * APPEND - Changeset will be considered as an addition to the end of all
	//    prior loaded changesets.
	//
	// ChangeType is a required field
	ChangeType *string `locationName:"changeType" type:"string" required:"true" enum:"ChangeType"`

	// The unique identifier for the FinSpace dataset in which the changeset will
	// be created.
	//
	// DatasetId is a required field
	DatasetId *string `location:"uri" locationName:"datasetId" min:"1" type:"string" required:"true"`

	// Options that define the structure of the source file(s).
	FormatParams map[string]*string `locationName:"formatParams" type:"map"`

	// Format type of the input files being loaded into the changeset.
	FormatType *string `locationName:"formatType" type:"string" enum:"FormatType"`

	// Source path from which the files to create the changeset will be sourced.
	//
	// SourceParams is a required field
	SourceParams map[string]*string `locationName:"sourceParams" type:"map" required:"true"`

	// Type of the data source from which the files to create the changeset will
	// be sourced.
	//
	//    * S3 - Amazon S3.
	//
	// SourceType is a required field
	SourceType *string `locationName:"sourceType" type:"string" required:"true" enum:"SourceType"`

	// Metadata tags to apply to this changeset.
	Tags map[string]*string `locationName:"tags" type:"map"`
}

// String returns the string representation
func (s CreateChangesetInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateChangesetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateChangesetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateChangesetInput"}
	if s.ChangeType == nil {
		invalidParams.Add(request.NewErrParamRequired("ChangeType"))
	}
	if s.DatasetId == nil {
		invalidParams.Add(request.NewErrParamRequired("DatasetId"))
	}
	if s.DatasetId != nil && len(*s.DatasetId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DatasetId", 1))
	}
	if s.SourceParams == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceParams"))
	}
	if s.SourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetChangeType sets the ChangeType field's value.
func (s *CreateChangesetInput) SetChangeType(v string) *CreateChangesetInput {
	s.ChangeType = &v
	return s
}

// SetDatasetId sets the DatasetId field's value.
func (s *CreateChangesetInput) SetDatasetId(v string) *CreateChangesetInput {
	s.DatasetId = &v
	return s
}

// SetFormatParams sets the FormatParams field's value.
func (s *CreateChangesetInput) SetFormatParams(v map[string]*string) *CreateChangesetInput {
	s.FormatParams = v
	return s
}

// SetFormatType sets the FormatType field's value.
func (s *CreateChangesetInput) SetFormatType(v string) *CreateChangesetInput {
	s.FormatType = &v
	return s
}

// SetSourceParams sets the SourceParams field's value.
func (s *CreateChangesetInput) SetSourceParams(v map[string]*string) *CreateChangesetInput {
	s.SourceParams = v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *CreateChangesetInput) SetSourceType(v string) *CreateChangesetInput {
	s.SourceType = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateChangesetInput) SetTags(v map[string]*string) *CreateChangesetInput {
	s.Tags = v
	return s
}

type CreateChangesetOutput struct {
	_ struct{} `type:"structure"`

	// Returns the changeset details.
	Changeset *ChangesetInfo `locationName:"changeset" type:"structure"`
}

// String returns the string representation
func (s CreateChangesetOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateChangesetOutput) GoString() string {
	return s.String()
}

// SetChangeset sets the Changeset field's value.
func (s *CreateChangesetOutput) SetChangeset(v *ChangesetInfo) *CreateChangesetOutput {
	s.Changeset = v
	return s
}

// Set short term API credentials.
type Credentials struct {
	_ struct{} `type:"structure"`

	// The access key identifier.
	AccessKeyId *string `locationName:"accessKeyId" min:"1" type:"string"`

	// The access key.
	SecretAccessKey *string `locationName:"secretAccessKey" type:"string"`

	// The session token.
	SessionToken *string `locationName:"sessionToken" type:"string"`
}

// String returns the string representation
func (s Credentials) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Credentials) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *Credentials) SetAccessKeyId(v string) *Credentials {
	s.AccessKeyId = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *Credentials) SetSecretAccessKey(v string) *Credentials {
	s.SecretAccessKey = &v
	return s
}

// SetSessionToken sets the SessionToken field's value.
func (s *Credentials) SetSessionToken(v string) *Credentials {
	s.SessionToken = &v
	return s
}

// Error message.
type ErrorInfo struct {
	_ struct{} `type:"structure"`

	// The category of the error.
	//
	//    * VALIDATION -The inputs to this request are invalid.
	//
	//    * SERVICE_QUOTA_EXCEEDED - Service quotas have been exceeded. Please contact
	//    AWS support to increase quotas.
	//
	//    * ACCESS_DENIED - Missing required permission to perform this request.
	//
	//    * RESOURCE_NOT_FOUND - One or more inputs to this request were not found.
	//
	//    * THROTTLING - The system temporarily lacks sufficient resources to process
	//    the request.
	//
	//    * INTERNAL_SERVICE_EXCEPTION - An internal service error has occurred.
	//
	//    * CANCELLED - A user recoverable error has occurred.
	ErrorCategory *string `locationName:"errorCategory" type:"string" enum:"ErrorCategory"`

	// The text of the error message.
	ErrorMessage *string `locationName:"errorMessage" type:"string"`
}

// String returns the string representation
func (s ErrorInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorInfo) GoString() string {
	return s.String()
}

// SetErrorCategory sets the ErrorCategory field's value.
func (s *ErrorInfo) SetErrorCategory(v string) *ErrorInfo {
	s.ErrorCategory = &v
	return s
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *ErrorInfo) SetErrorMessage(v string) *ErrorInfo {
	s.ErrorMessage = &v
	return s
}

type GetProgrammaticAccessCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The time duration in which the credentials remain valid.
	DurationInMinutes *int64 `location:"querystring" locationName:"durationInMinutes" min:"60" type:"long"`

	// The habanero environment identifier.
	//
	// EnvironmentId is a required field
	EnvironmentId *string `location:"querystring" locationName:"environmentId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetProgrammaticAccessCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProgrammaticAccessCredentialsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProgrammaticAccessCredentialsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetProgrammaticAccessCredentialsInput"}
	if s.DurationInMinutes != nil && *s.DurationInMinutes < 60 {
		invalidParams.Add(request.NewErrParamMinValue("DurationInMinutes", 60))
	}
	if s.EnvironmentId == nil {
		invalidParams.Add(request.NewErrParamRequired("EnvironmentId"))
	}
	if s.EnvironmentId != nil && len(*s.EnvironmentId) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("EnvironmentId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationInMinutes sets the DurationInMinutes field's value.
func (s *GetProgrammaticAccessCredentialsInput) SetDurationInMinutes(v int64) *GetProgrammaticAccessCredentialsInput {
	s.DurationInMinutes = &v
	return s
}

// SetEnvironmentId sets the EnvironmentId field's value.
func (s *GetProgrammaticAccessCredentialsInput) SetEnvironmentId(v string) *GetProgrammaticAccessCredentialsInput {
	s.EnvironmentId = &v
	return s
}

type GetProgrammaticAccessCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// Returns the programmatic credentials.
	Credentials *Credentials `locationName:"credentials" type:"structure"`

	// Returns the duration in which the credentials will remain valid.
	DurationInMinutes *int64 `locationName:"durationInMinutes" min:"60" type:"long"`
}

// String returns the string representation
func (s GetProgrammaticAccessCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProgrammaticAccessCredentialsOutput) GoString() string {
	return s.String()
}

// SetCredentials sets the Credentials field's value.
func (s *GetProgrammaticAccessCredentialsOutput) SetCredentials(v *Credentials) *GetProgrammaticAccessCredentialsOutput {
	s.Credentials = v
	return s
}

// SetDurationInMinutes sets the DurationInMinutes field's value.
func (s *GetProgrammaticAccessCredentialsOutput) SetDurationInMinutes(v int64) *GetProgrammaticAccessCredentialsOutput {
	s.DurationInMinutes = &v
	return s
}

type GetWorkingLocationInput struct {
	_ struct{} `type:"structure"`

	// Specify the type of the working location.
	//
	//    * SAGEMAKER - Use the Amazon S3 location as a temporary location to store
	//    data content when working with FinSpace Notebooks that run on SageMaker
	//    studio.
	//
	//    * INGESTION - Use the Amazon S3 location as a staging location to copy
	//    your data content and then use the location with the changeset creation
	//    operation.
	LocationType *string `locationName:"locationType" type:"string" enum:"LocationType"`
}

// String returns the string representation
func (s GetWorkingLocationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetWorkingLocationInput) GoString() string {
	return s.String()
}

// SetLocationType sets the LocationType field's value.
func (s *GetWorkingLocationInput) SetLocationType(v string) *GetWorkingLocationInput {
	s.LocationType = &v
	return s
}

type GetWorkingLocationOutput struct {
	_ struct{} `type:"structure"`

	// Returns the Amazon S3 bucket name for the working location.
	S3Bucket *string `locationName:"s3Bucket" min:"1" type:"string"`

	// Returns the Amazon S3 Path for the working location.
	S3Path *string `locationName:"s3Path" min:"1" type:"string"`

	// Returns the Amazon S3 URI for the working location.
	S3Uri *string `locationName:"s3Uri" min:"1" type:"string"`
}

// String returns the string representation
func (s GetWorkingLocationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetWorkingLocationOutput) GoString() string {
	return s.String()
}

// SetS3Bucket sets the S3Bucket field's value.
func (s *GetWorkingLocationOutput) SetS3Bucket(v string) *GetWorkingLocationOutput {
	s.S3Bucket = &v
	return s
}

// SetS3Path sets the S3Path field's value.
func (s *GetWorkingLocationOutput) SetS3Path(v string) *GetWorkingLocationOutput {
	s.S3Path = &v
	return s
}

// SetS3Uri sets the S3Uri field's value.
func (s *GetWorkingLocationOutput) SetS3Uri(v string) *GetWorkingLocationOutput {
	s.S3Uri = &v
	return s
}

// The request processing has failed because of an unknown error, exception
// or failure.
type InternalServerException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InternalServerException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InternalServerException) GoString() string {
	return s.String()
}

func newErrorInternalServerException(v protocol.ResponseMetadata) error {
	return &InternalServerException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *InternalServerException) Code() string {
	return "InternalServerException"
}

// Message returns the exception's message.
func (s *InternalServerException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *InternalServerException) OrigErr() error {
	return nil
}

func (s *InternalServerException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *InternalServerException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *InternalServerException) RequestID() string {
	return s.RespMetadata.RequestID
}

// One or more resources can't be found.
type ResourceNotFoundException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s *ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ResourceNotFoundException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ResourceNotFoundException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The request was denied due to request throttling.
type ThrottlingException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ThrottlingException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ThrottlingException) GoString() string {
	return s.String()
}

func newErrorThrottlingException(v protocol.ResponseMetadata) error {
	return &ThrottlingException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ThrottlingException) Code() string {
	return "ThrottlingException"
}

// Message returns the exception's message.
func (s *ThrottlingException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ThrottlingException) OrigErr() error {
	return nil
}

func (s *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ThrottlingException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ThrottlingException) RequestID() string {
	return s.RespMetadata.RequestID
}

// The input fails to satisfy the constraints specified by an AWS service.
type ValidationException struct {
	_            struct{}                  `type:"structure"`
	RespMetadata protocol.ResponseMetadata `json:"-" xml:"-"`

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ValidationException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ValidationException) GoString() string {
	return s.String()
}

func newErrorValidationException(v protocol.ResponseMetadata) error {
	return &ValidationException{
		RespMetadata: v,
	}
}

// Code returns the exception type name.
func (s *ValidationException) Code() string {
	return "ValidationException"
}

// Message returns the exception's message.
func (s *ValidationException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s *ValidationException) OrigErr() error {
	return nil
}

func (s *ValidationException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s *ValidationException) StatusCode() int {
	return s.RespMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s *ValidationException) RequestID() string {
	return s.RespMetadata.RequestID
}

const (
	// ChangeTypeReplace is a ChangeType enum value
	ChangeTypeReplace = "REPLACE"

	// ChangeTypeAppend is a ChangeType enum value
	ChangeTypeAppend = "APPEND"

	// ChangeTypeModify is a ChangeType enum value
	ChangeTypeModify = "MODIFY"
)

// ChangeType_Values returns all elements of the ChangeType enum
func ChangeType_Values() []string {
	return []string{
		ChangeTypeReplace,
		ChangeTypeAppend,
		ChangeTypeModify,
	}
}

const (
	// ChangesetStatusPending is a ChangesetStatus enum value
	ChangesetStatusPending = "PENDING"

	// ChangesetStatusFailed is a ChangesetStatus enum value
	ChangesetStatusFailed = "FAILED"

	// ChangesetStatusSuccess is a ChangesetStatus enum value
	ChangesetStatusSuccess = "SUCCESS"

	// ChangesetStatusRunning is a ChangesetStatus enum value
	ChangesetStatusRunning = "RUNNING"

	// ChangesetStatusStopRequested is a ChangesetStatus enum value
	ChangesetStatusStopRequested = "STOP_REQUESTED"
)

// ChangesetStatus_Values returns all elements of the ChangesetStatus enum
func ChangesetStatus_Values() []string {
	return []string{
		ChangesetStatusPending,
		ChangesetStatusFailed,
		ChangesetStatusSuccess,
		ChangesetStatusRunning,
		ChangesetStatusStopRequested,
	}
}

const (
	// ErrorCategoryTheInputsToThisRequestAreInvalid is a ErrorCategory enum value
	ErrorCategoryTheInputsToThisRequestAreInvalid = "The_inputs_to_this_request_are_invalid"

	// ErrorCategoryServiceLimitsHaveBeenExceeded is a ErrorCategory enum value
	ErrorCategoryServiceLimitsHaveBeenExceeded = "Service_limits_have_been_exceeded"

	// ErrorCategoryMissingRequiredPermissionToPerformThisRequest is a ErrorCategory enum value
	ErrorCategoryMissingRequiredPermissionToPerformThisRequest = "Missing_required_permission_to_perform_this_request"

	// ErrorCategoryOneOrMoreInputsToThisRequestWereNotFound is a ErrorCategory enum value
	ErrorCategoryOneOrMoreInputsToThisRequestWereNotFound = "One_or_more_inputs_to_this_request_were_not_found"

	// ErrorCategoryTheSystemTemporarilyLacksSufficientResourcesToProcessTheRequest is a ErrorCategory enum value
	ErrorCategoryTheSystemTemporarilyLacksSufficientResourcesToProcessTheRequest = "The_system_temporarily_lacks_sufficient_resources_to_process_the_request"

	// ErrorCategoryAnInternalErrorHasOccurred is a ErrorCategory enum value
	ErrorCategoryAnInternalErrorHasOccurred = "An_internal_error_has_occurred"

	// ErrorCategoryCancelled is a ErrorCategory enum value
	ErrorCategoryCancelled = "Cancelled"

	// ErrorCategoryAUserRecoverableErrorHasOccurred is a ErrorCategory enum value
	ErrorCategoryAUserRecoverableErrorHasOccurred = "A_user_recoverable_error_has_occurred"
)

// ErrorCategory_Values returns all elements of the ErrorCategory enum
func ErrorCategory_Values() []string {
	return []string{
		ErrorCategoryTheInputsToThisRequestAreInvalid,
		ErrorCategoryServiceLimitsHaveBeenExceeded,
		ErrorCategoryMissingRequiredPermissionToPerformThisRequest,
		ErrorCategoryOneOrMoreInputsToThisRequestWereNotFound,
		ErrorCategoryTheSystemTemporarilyLacksSufficientResourcesToProcessTheRequest,
		ErrorCategoryAnInternalErrorHasOccurred,
		ErrorCategoryCancelled,
		ErrorCategoryAUserRecoverableErrorHasOccurred,
	}
}

const (
	// FormatTypeCsv is a FormatType enum value
	FormatTypeCsv = "CSV"

	// FormatTypeJson is a FormatType enum value
	FormatTypeJson = "JSON"

	// FormatTypeParquet is a FormatType enum value
	FormatTypeParquet = "PARQUET"

	// FormatTypeXml is a FormatType enum value
	FormatTypeXml = "XML"
)

// FormatType_Values returns all elements of the FormatType enum
func FormatType_Values() []string {
	return []string{
		FormatTypeCsv,
		FormatTypeJson,
		FormatTypeParquet,
		FormatTypeXml,
	}
}

const (
	// LocationTypeIngestion is a LocationType enum value
	LocationTypeIngestion = "INGESTION"

	// LocationTypeSagemaker is a LocationType enum value
	LocationTypeSagemaker = "SAGEMAKER"
)

// LocationType_Values returns all elements of the LocationType enum
func LocationType_Values() []string {
	return []string{
		LocationTypeIngestion,
		LocationTypeSagemaker,
	}
}

const (
	// SourceTypeS3 is a SourceType enum value
	SourceTypeS3 = "S3"
)

// SourceType_Values returns all elements of the SourceType enum
func SourceType_Values() []string {
	return []string{
		SourceTypeS3,
	}
}