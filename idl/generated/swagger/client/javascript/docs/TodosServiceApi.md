# Api.TodosServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTodo**](TodosServiceApi.md#createTodo) | **POST** /v1/todos | 
[**deleteTodo**](TodosServiceApi.md#deleteTodo) | **DELETE** /v1/todos/{todo_id} | 
[**getTodo**](TodosServiceApi.md#getTodo) | **GET** /v1/todos/{todo_id} | 
[**listTodos**](TodosServiceApi.md#listTodos) | **GET** /v1/todos | 
[**updateTodo**](TodosServiceApi.md#updateTodo) | **PUT** /v1/todos/{todo.todo_id} | 


<a name="createTodo"></a>
# **createTodo**
> TodosCreateTodoResponse createTodo(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var body = new Api.TodosCreateTodoRequest(); // TodosCreateTodoRequest | 

apiInstance.createTodo(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TodosCreateTodoRequest**](TodosCreateTodoRequest.md)|  | 

### Return type

[**TodosCreateTodoResponse**](TodosCreateTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteTodo"></a>
# **deleteTodo**
> TodosDeleteTodoResponse deleteTodo(todoId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoId = "todoId_example"; // String | 

apiInstance.deleteTodo(todoId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoId** | **String**|  | 

### Return type

[**TodosDeleteTodoResponse**](TodosDeleteTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getTodo"></a>
# **getTodo**
> TodosGetTodoResponse getTodo(todoId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoId = "todoId_example"; // String | 

apiInstance.getTodo(todoId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoId** | **String**|  | 

### Return type

[**TodosGetTodoResponse**](TodosGetTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listTodos"></a>
# **listTodos**
> TodosListTodosResponse listTodos(opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var opts = { 
  'parentType': "PARENT_TYPE_TODO_GROUP", // String | 
  'parentId': "parentId_example" // String | 
};
apiInstance.listTodos(opts).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **parentType** | **String**|  | [optional] [default to PARENT_TYPE_TODO_GROUP]
 **parentId** | **String**|  | [optional] 

### Return type

[**TodosListTodosResponse**](TodosListTodosResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateTodo"></a>
# **updateTodo**
> TodosUpdateTodoResponse updateTodo(todoTodoId, body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoTodoId = "todoTodoId_example"; // String | 

var body = new Api.TodosUpdateTodoRequest(); // TodosUpdateTodoRequest | 

apiInstance.updateTodo(todoTodoId, body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoTodoId** | **String**|  | 
 **body** | [**TodosUpdateTodoRequest**](TodosUpdateTodoRequest.md)|  | 

### Return type

[**TodosUpdateTodoResponse**](TodosUpdateTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

