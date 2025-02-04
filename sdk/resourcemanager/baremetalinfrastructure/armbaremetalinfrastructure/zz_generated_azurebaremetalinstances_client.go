//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbaremetalinfrastructure

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AzureBareMetalInstancesClient contains the methods for the AzureBareMetalInstances group.
// Don't use this type directly, use NewAzureBareMetalInstancesClient() instead.
type AzureBareMetalInstancesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAzureBareMetalInstancesClient creates a new instance of AzureBareMetalInstancesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAzureBareMetalInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AzureBareMetalInstancesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AzureBareMetalInstancesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets an Azure BareMetal instance for the specified subscription, resource group, and instance name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-09
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - Name of the Azure BareMetal on Azure instance.
// options - AzureBareMetalInstancesClientGetOptions contains the optional parameters for the AzureBareMetalInstancesClient.Get
// method.
func (client *AzureBareMetalInstancesClient) Get(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientGetOptions) (AzureBareMetalInstancesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, options)
	if err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureBareMetalInstancesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AzureBareMetalInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, options *AzureBareMetalInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AzureBareMetalInstancesClient) getHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientGetResponse, error) {
	result := AzureBareMetalInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstance); err != nil {
		return AzureBareMetalInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets a list of AzureBareMetal instances in the specified subscription and resource group.
// The operations returns various properties of each Azure BareMetal instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-09
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - AzureBareMetalInstancesClientListByResourceGroupOptions contains the optional parameters for the AzureBareMetalInstancesClient.ListByResourceGroup
// method.
func (client *AzureBareMetalInstancesClient) NewListByResourceGroupPager(resourceGroupName string, options *AzureBareMetalInstancesClientListByResourceGroupOptions) *runtime.Pager[AzureBareMetalInstancesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalInstancesClientListByResourceGroupResponse]{
		More: func(page AzureBareMetalInstancesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalInstancesClientListByResourceGroupResponse) (AzureBareMetalInstancesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureBareMetalInstancesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureBareMetalInstancesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureBareMetalInstancesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AzureBareMetalInstancesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AzureBareMetalInstancesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AzureBareMetalInstancesClient) listByResourceGroupHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientListByResourceGroupResponse, error) {
	result := AzureBareMetalInstancesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstancesListResult); err != nil {
		return AzureBareMetalInstancesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets a list of AzureBareMetal instances in the specified subscription. The operations returns
// various properties of each Azure BareMetal instance.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-09
// options - AzureBareMetalInstancesClientListBySubscriptionOptions contains the optional parameters for the AzureBareMetalInstancesClient.ListBySubscription
// method.
func (client *AzureBareMetalInstancesClient) NewListBySubscriptionPager(options *AzureBareMetalInstancesClientListBySubscriptionOptions) *runtime.Pager[AzureBareMetalInstancesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AzureBareMetalInstancesClientListBySubscriptionResponse]{
		More: func(page AzureBareMetalInstancesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AzureBareMetalInstancesClientListBySubscriptionResponse) (AzureBareMetalInstancesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AzureBareMetalInstancesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AzureBareMetalInstancesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AzureBareMetalInstancesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AzureBareMetalInstancesClient) listBySubscriptionCreateRequest(ctx context.Context, options *AzureBareMetalInstancesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AzureBareMetalInstancesClient) listBySubscriptionHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientListBySubscriptionResponse, error) {
	result := AzureBareMetalInstancesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstancesListResult); err != nil {
		return AzureBareMetalInstancesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patches the Tags field of a Azure BareMetal instance for the specified subscription, resource group, and instance
// name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-08-09
// resourceGroupName - The name of the resource group. The name is case insensitive.
// azureBareMetalInstanceName - Name of the Azure BareMetal on Azure instance.
// tagsParameter - Request body that only contains the new Tags field
// options - AzureBareMetalInstancesClientUpdateOptions contains the optional parameters for the AzureBareMetalInstancesClient.Update
// method.
func (client *AzureBareMetalInstancesClient) Update(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags, options *AzureBareMetalInstancesClientUpdateOptions) (AzureBareMetalInstancesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, azureBareMetalInstanceName, tagsParameter, options)
	if err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AzureBareMetalInstancesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AzureBareMetalInstancesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, azureBareMetalInstanceName string, tagsParameter Tags, options *AzureBareMetalInstancesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.BareMetalInfrastructure/bareMetalInstances/{azureBareMetalInstanceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if azureBareMetalInstanceName == "" {
		return nil, errors.New("parameter azureBareMetalInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{azureBareMetalInstanceName}", url.PathEscape(azureBareMetalInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, tagsParameter)
}

// updateHandleResponse handles the Update response.
func (client *AzureBareMetalInstancesClient) updateHandleResponse(resp *http.Response) (AzureBareMetalInstancesClientUpdateResponse, error) {
	result := AzureBareMetalInstancesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBareMetalInstance); err != nil {
		return AzureBareMetalInstancesClientUpdateResponse{}, err
	}
	return result, nil
}
