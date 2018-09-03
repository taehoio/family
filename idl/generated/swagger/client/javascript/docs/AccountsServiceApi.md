# Api.AccountsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**logIn**](AccountsServiceApi.md#logIn) | **POST** /v1/accounts/logIn | 
[**register**](AccountsServiceApi.md#register) | **POST** /v1/accounts/register | 


<a name="logIn"></a>
# **logIn**
> AccountsLogInResponse logIn(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.AccountsServiceApi();

var body = new Api.AccountsLogInRequest(); // AccountsLogInRequest | 

apiInstance.logIn(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountsLogInRequest**](AccountsLogInRequest.md)|  | 

### Return type

[**AccountsLogInResponse**](AccountsLogInResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="register"></a>
# **register**
> AccountsRegisterResponse register(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.AccountsServiceApi();

var body = new Api.AccountsRegisterRequest(); // AccountsRegisterRequest | 

apiInstance.register(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccountsRegisterRequest**](AccountsRegisterRequest.md)|  | 

### Return type

[**AccountsRegisterResponse**](AccountsRegisterResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

