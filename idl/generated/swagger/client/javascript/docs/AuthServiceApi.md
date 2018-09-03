# Api.AuthServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**refresh**](AuthServiceApi.md#refresh) | **POST** /v1/auth/refreshToken | 


<a name="refresh"></a>
# **refresh**
> AuthRefreshResponse refresh(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.AuthServiceApi();

var body = new Api.AuthRefreshRequest(); // AuthRefreshRequest | 

apiInstance.refresh(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AuthRefreshRequest**](AuthRefreshRequest.md)|  | 

### Return type

[**AuthRefreshResponse**](AuthRefreshResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

