// Copyright 2021 Tencent Inc. All rights reserved.
//
// 营销加价购对外API
//
// 指定服务商可通过该接口报名加价购活动、查询某个区域内的加价购活动列表、锁定加价活动购资格以及解锁加价购活动资格。
//
// API version: 1.3.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package retailstore

import (
	"context"
	"fmt"
	nethttp "net/http"
	neturl "net/url"
	"strings"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/services"
)

type RetailStoreActApiService services.Service

// AddRepresentative 添加零售小店活动业务代理
//
// 该接口为服务商或商户给零售小店活动添加业务代理的专用接口。 使用对象：活动创建方商户号、活动归属品牌的品牌主商户号或品牌经营商户号。
func (a *RetailStoreActApiService) AddRepresentative(ctx context.Context, req AddRepresentativeRequest) (resp *Representatives, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPut
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.ActivityId == nil {
		return nil, nil, fmt.Errorf("field `ActivityId` is required and must be specified in AddRepresentativeRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representative"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"activity_id"+"}", neturl.PathEscape(core.ParameterToString(*req.ActivityId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &AddRepresentativesRequest{
		RepresentativeInfoList: req.RepresentativeInfoList,
		OutRequestNo:           req.OutRequestNo,
		AddTime:                req.AddTime,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Representatives from Http Response
	resp = new(Representatives)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// CreateMaterials 生成小店活动物料码
//
// 该接口为服务商或商户给零售小店活动申请物料码专用接口。 使用对象：品牌的品牌主商户号或品牌服务商。
func (a *RetailStoreActApiService) CreateMaterials(ctx context.Context, req CreateMaterialsRequest) (resp *Materials, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodPost
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.BrandId == nil {
		return nil, nil, fmt.Errorf("field `BrandId` is required and must be specified in CreateMaterialsRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/retail-store-act/{brand_id}/materials"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"brand_id"+"}", neturl.PathEscape(core.ParameterToString(*req.BrandId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &CreateMaterialsBody{
		OutRequestNo: req.OutRequestNo,
		MaterialNum:  req.MaterialNum,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract Materials from Http Response
	resp = new(Materials)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// DeleteRepresentative 删除零售小店活动业务代理
//
// 该接口为服务商或商户给零售小店活动删除业务代理的专用接口。 使用对象：活动创建方商户号、活动归属品牌的品牌主商户号或品牌经营商户号。
func (a *RetailStoreActApiService) DeleteRepresentative(ctx context.Context, req DeleteRepresentativeRequest) (resp *DeleteRepresentativeResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.ActivityId == nil {
		return nil, nil, fmt.Errorf("field `ActivityId` is required and must be specified in DeleteRepresentativeRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representative"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"activity_id"+"}", neturl.PathEscape(core.ParameterToString(*req.ActivityId, "")), -1)

	// Make sure All Required Params are properly set

	// Setup Body Params
	localVarPostBody = &DeleteRepresentativeBody{
		RepresentativeInfoList: req.RepresentativeInfoList,
		OutRequestNo:           req.OutRequestNo,
		DeleteTime:             req.DeleteTime,
	}

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{"application/json"}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract DeleteRepresentativeResponse from Http Response
	resp = new(DeleteRepresentativeResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}

// ListRepresentative 查询零售小店活动业务代理
//
// 该接口为服务商或商户给零售小店活动查询业务代理的专用接口。 使用对象：活动创建方商户号、活动归属品牌的品牌主商户号或品牌经营商户号。
func (a *RetailStoreActApiService) ListRepresentative(ctx context.Context, req ListRepresentativeRequest) (resp *ListRepresentativeResponse, result *core.APIResult, err error) {
	var (
		localVarHTTPMethod   = nethttp.MethodGet
		localVarPostBody     interface{}
		localVarQueryParams  neturl.Values
		localVarHeaderParams = nethttp.Header{}
	)

	// Make sure Path Params are properly set
	if req.ActivityId == nil {
		return nil, nil, fmt.Errorf("field `ActivityId` is required and must be specified in ListRepresentativeRequest")
	}

	localVarPath := consts.WechatPayAPIServer + "/v3/marketing/goods-subsidy-activity/retail-store-act/{activity_id}/representatives"
	// Build Path with Path Params
	localVarPath = strings.Replace(localVarPath, "{"+"activity_id"+"}", neturl.PathEscape(core.ParameterToString(*req.ActivityId, "")), -1)

	// Make sure All Required Params are properly set
	if req.Offset == nil {
		return nil, nil, fmt.Errorf("field `Offset` is required and must be specified in ListRepresentativeRequest")
	}
	if req.Limit == nil {
		return nil, nil, fmt.Errorf("field `Limit` is required and must be specified in ListRepresentativeRequest")
	}

	// Setup Query Params
	localVarQueryParams = neturl.Values{}
	localVarQueryParams.Add("offset", core.ParameterToString(*req.Offset, ""))
	localVarQueryParams.Add("limit", core.ParameterToString(*req.Limit, ""))

	// Determine the Content-Type Header
	localVarHTTPContentTypes := []string{}
	// Setup Content-Type
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)

	// Perform Http Request
	result, err = a.Client.Request(ctx, localVarHTTPMethod, localVarPath, localVarHeaderParams, localVarQueryParams, localVarPostBody, localVarHTTPContentType)
	if err != nil {
		return nil, result, err
	}

	// Extract ListRepresentativeResponse from Http Response
	resp = new(ListRepresentativeResponse)
	err = core.UnMarshalResponse(result.Response, resp)
	if err != nil {
		return nil, result, err
	}
	return resp, result, nil
}
