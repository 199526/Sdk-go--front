package videoanalyzer

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// OperationResultsGroupClient is the azure Video Analyzer provides a platform for you to build intelligent video
// applications that span the edge and the cloud
type OperationResultsGroupClient struct {
	BaseClient
}

// NewOperationResultsGroupClient creates an instance of the OperationResultsGroupClient client.
func NewOperationResultsGroupClient(subscriptionID string) OperationResultsGroupClient {
	return NewOperationResultsGroupClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOperationResultsGroupClientWithBaseURI creates an instance of the OperationResultsGroupClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewOperationResultsGroupClientWithBaseURI(baseURI string, subscriptionID string) OperationResultsGroupClient {
	return OperationResultsGroupClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get video analyzer operation result.
// Parameters:
// locationName - location name.
// operationID - operation Id.
func (client OperationResultsGroupClient) Get(ctx context.Context, locationName string, operationID string) (result Model, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationResultsGroupClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("videoanalyzer.OperationResultsGroupClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, locationName, operationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationResultsGroupClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationResultsGroupClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "videoanalyzer.OperationResultsGroupClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client OperationResultsGroupClient) GetPreparer(ctx context.Context, locationName string, operationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"operationId":    autorest.Encode("path", operationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Media/locations/{locationName}/videoAnalyzerOperationResults/{operationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OperationResultsGroupClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OperationResultsGroupClient) GetResponder(resp *http.Response) (result Model, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
