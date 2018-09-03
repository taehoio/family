# Api.TodoGroupsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTodoGroup**](TodoGroupsServiceApi.md#createTodoGroup) | **POST** /v1/todo-groups | 
[**deleteTodoGroup**](TodoGroupsServiceApi.md#deleteTodoGroup) | **DELETE** /v1/todo-groups/{todo_group_id} | 
[**getTodoGroup**](TodoGroupsServiceApi.md#getTodoGroup) | **GET** /v1/todo-groups/{todo_group_id} | 
[**listTodoGroups**](TodoGroupsServiceApi.md#listTodoGroups) | **GET** /v1/todo-groups | 
[**updateTodoGroup**](TodoGroupsServiceApi.md#updateTodoGroup) | **PUT** /v1/todo-groups/{todo_group_id} | 


<a name="createTodoGroup"></a>
# **createTodoGroup**
> TodogroupsCreateTodoGroupResponse createTodoGroup(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var body = new Api.TodogroupsCreateTodoGroupRequest(); // TodogroupsCreateTodoGroupRequest | 

apiInstance.createTodoGroup(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TodogroupsCreateTodoGroupRequest**](TodogroupsCreateTodoGroupRequest.md)|  | 

### Return type

[**TodogroupsCreateTodoGroupResponse**](TodogroupsCreateTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteTodoGroup"></a>
# **deleteTodoGroup**
> TodogroupsDeleteTodoGroupResponse deleteTodoGroup(todoGroupId, opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupId = "todoGroupId_example"; // String | 

var opts = { 
  'accountId': "accountId_example" // String | 
};
apiInstance.deleteTodoGroup(todoGroupId, opts).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupId** | **String**|  | 
 **accountId** | **String**|  | [optional] 

### Return type

[**TodogroupsDeleteTodoGroupResponse**](TodogroupsDeleteTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getTodoGroup"></a>
# **getTodoGroup**
> TodogroupsGetTodoGroupResponse getTodoGroup(todoGroupId, opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupId = "todoGroupId_example"; // String | 

var opts = { 
  'accountId': "accountId_example" // String | 
};
apiInstance.getTodoGroup(todoGroupId, opts).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupId** | **String**|  | 
 **accountId** | **String**|  | [optional] 

### Return type

[**TodogroupsGetTodoGroupResponse**](TodogroupsGetTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listTodoGroups"></a>
# **listTodoGroups**
> TodogroupsListTodoGroupsResponse listTodoGroups(opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var opts = { 
  'accountId': "accountId_example" // String | 
};
apiInstance.listTodoGroups(opts).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **String**|  | [optional] 

### Return type

[**TodogroupsListTodoGroupsResponse**](TodogroupsListTodoGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateTodoGroup"></a>
# **updateTodoGroup**
> TodogroupsUpdateTodoGroupResponse updateTodoGroup(todoGroupId, body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupId = "todoGroupId_example"; // String | 

var body = new Api.TodogroupsUpdateTodoGroupRequest(); // TodogroupsUpdateTodoGroupRequest | 

apiInstance.updateTodoGroup(todoGroupId, body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupId** | **String**|  | 
 **body** | [**TodogroupsUpdateTodoGroupRequest**](TodogroupsUpdateTodoGroupRequest.md)|  | 

### Return type

[**TodogroupsUpdateTodoGroupResponse**](TodogroupsUpdateTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

