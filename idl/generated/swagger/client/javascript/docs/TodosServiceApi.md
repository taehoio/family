# Api.TodosServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createTodo**](TodosServiceApi.md#createTodo) | **POST** /v1/todos | 
[**deleteTodo**](TodosServiceApi.md#deleteTodo) | **DELETE** /v1/todos/{todo_id} | 
[**getTodo**](TodosServiceApi.md#getTodo) | **GET** /v1/todos/{todo_id} | 
[**listTodos**](TodosServiceApi.md#listTodos) | **GET** /v1/todos | 
[**updateTodo**](TodosServiceApi.md#updateTodo) | **PUT** /v1/todos/{todo_id} | 


<a name="createTodo"></a>
# **createTodo**
> TodosCreateTodoResponse createTodo(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var body = new Api.TodosCreateTodoRequest(); // TodosCreateTodoRequest | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.createTodo(body, callback);
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
> TodosDeleteTodoResponse deleteTodo(todoId, opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoId = "todoId_example"; // String | 

var opts = { 
  'accountId': "accountId_example" // String | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.deleteTodo(todoId, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoId** | **String**|  | 
 **accountId** | **String**|  | [optional] 

### Return type

[**TodosDeleteTodoResponse**](TodosDeleteTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getTodo"></a>
# **getTodo**
> TodosGetTodoResponse getTodo(todoId, opts)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoId = "todoId_example"; // String | 

var opts = { 
  'accountId': "accountId_example" // String | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getTodo(todoId, opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoId** | **String**|  | 
 **accountId** | **String**|  | [optional] 

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
  'accountId': "accountId_example", // String | 
  'parentType': "PARENT_TYPE_TODO_GROUP", // String | 
  'parentId': "parentId_example" // String | 
};

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.listTodos(opts, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **String**|  | [optional] 
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
> TodosUpdateTodoResponse updateTodo(todoId, body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.TodosServiceApi();

var todoId = "todoId_example"; // String | 

var body = new Api.TodosUpdateTodoRequest(); // TodosUpdateTodoRequest | 


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.updateTodo(todoId, body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todoId** | **String**|  | 
 **body** | [**TodosUpdateTodoRequest**](TodosUpdateTodoRequest.md)|  | 

### Return type

[**TodosUpdateTodoResponse**](TodosUpdateTodoResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

