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


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.logIn(body, callback);
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


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.register(body, callback);
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

