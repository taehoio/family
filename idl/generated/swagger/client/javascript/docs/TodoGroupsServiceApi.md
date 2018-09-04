# Api.TodoGroupsServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTodoGroup**](TodoGroupsServiceApi.md#createTodoGroup) | **POST** /v1/todo-groups | 
[**deleteTodoGroup**](TodoGroupsServiceApi.md#deleteTodoGroup) | **DELETE** /v1/todo-groups/{todo_group_id} | 
[**getTodoGroup**](TodoGroupsServiceApi.md#getTodoGroup) | **GET** /v1/todo-groups/{todo_group_id} | 
[**listTodoGroups**](TodoGroupsServiceApi.md#listTodoGroups) | **GET** /v1/todo-groups | 
[**updateTodoGroup**](TodoGroupsServiceApi.md#updateTodoGroup) | **PUT** /v1/todo-groups/{todo_group.todo_group_id} | 


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
> TodogroupsDeleteTodoGroupResponse deleteTodoGroup(todoGroupId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupId = "todoGroupId_example"; // String | 

apiInstance.deleteTodoGroup(todoGroupId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupId** | **String**|  | 

### Return type

[**TodogroupsDeleteTodoGroupResponse**](TodogroupsDeleteTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getTodoGroup"></a>
# **getTodoGroup**
> TodogroupsGetTodoGroupResponse getTodoGroup(todoGroupId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupId = "todoGroupId_example"; // String | 

apiInstance.getTodoGroup(todoGroupId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupId** | **String**|  | 

### Return type

[**TodogroupsGetTodoGroupResponse**](TodogroupsGetTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listTodoGroups"></a>
# **listTodoGroups**
> TodogroupsListTodoGroupsResponse listTodoGroups()



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();
apiInstance.listTodoGroups().then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**TodogroupsListTodoGroupsResponse**](TodogroupsListTodoGroupsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateTodoGroup"></a>
# **updateTodoGroup**
> TodogroupsUpdateTodoGroupResponse updateTodoGroup(todoGroupTodoGroupId, body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodoGroupsServiceApi();

var todoGroupTodoGroupId = "todoGroupTodoGroupId_example"; // String | 

var body = new Api.TodogroupsUpdateTodoGroupRequest(); // TodogroupsUpdateTodoGroupRequest | 

apiInstance.updateTodoGroup(todoGroupTodoGroupId, body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoGroupTodoGroupId** | **String**|  | 
 **body** | [**TodogroupsUpdateTodoGroupRequest**](TodogroupsUpdateTodoGroupRequest.md)|  | 

### Return type

[**TodogroupsUpdateTodoGroupResponse**](TodogroupsUpdateTodoGroupResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

