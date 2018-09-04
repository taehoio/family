# Api.NotesServiceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createNote**](NotesServiceApi.md#createNote) | **POST** /v1/notes | 
[**deleteNote**](NotesServiceApi.md#deleteNote) | **DELETE** /v1/notes/{note_id} | 
[**getNote**](NotesServiceApi.md#getNote) | **GET** /v1/notes/{note_id} | 
[**listNotes**](NotesServiceApi.md#listNotes) | **GET** /v1/notes | 
[**updateNote**](NotesServiceApi.md#updateNote) | **PUT** /v1/notes/{note.note_id} | 


<a name="createNote"></a>
# **createNote**
> NotesCreateNoteResponse createNote(body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.NotesServiceApi();

var body = new Api.NotesCreateNoteRequest(); // NotesCreateNoteRequest | 

apiInstance.createNote(body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NotesCreateNoteRequest**](NotesCreateNoteRequest.md)|  | 

### Return type

[**NotesCreateNoteResponse**](NotesCreateNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteNote"></a>
# **deleteNote**
> NotesDeleteNoteResponse deleteNote(noteId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.NotesServiceApi();

var noteId = "noteId_example"; // String | 

apiInstance.deleteNote(noteId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteId** | **String**|  | 

### Return type

[**NotesDeleteNoteResponse**](NotesDeleteNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getNote"></a>
# **getNote**
> NotesGetNoteResponse getNote(noteId)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.NotesServiceApi();

var noteId = "noteId_example"; // String | 

apiInstance.getNote(noteId).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteId** | **String**|  | 

### Return type

[**NotesGetNoteResponse**](NotesGetNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="listNotes"></a>
# **listNotes**
> NotesListNotesResponse listNotes()



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.NotesServiceApi();
apiInstance.listNotes().then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**NotesListNotesResponse**](NotesListNotesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="updateNote"></a>
# **updateNote**
> NotesUpdateNoteResponse updateNote(noteNoteId, body)



### Example
```javascript
var Api = require('api');

var apiInstance = new Api.NotesServiceApi();

var noteNoteId = "noteNoteId_example"; // String | 

var body = new Api.NotesUpdateNoteRequest(); // NotesUpdateNoteRequest | 

apiInstance.updateNote(noteNoteId, body).then(function(data) {
  console.log('API called successfully. Returned data: ' + data);
}, function(error) {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteNoteId** | **String**|  | 
 **body** | [**NotesUpdateNoteRequest**](NotesUpdateNoteRequest.md)|  | 

### Return type

[**NotesUpdateNoteResponse**](NotesUpdateNoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

