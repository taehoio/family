//
// TodosServiceAPI.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation
import Alamofire
import RxSwift



open class TodosServiceAPI {
    /**
     Todos
     
     - parameter todoGroupId: (path)  
     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func createTodo(todoGroupId: String, body: TodosCreateTodoRequest, completion: @escaping ((_ data: TodosCreateTodoResponse?,_ error: Error?) -> Void)) {
        createTodoWithRequestBuilder(todoGroupId: todoGroupId, body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**
     Todos
     
     - parameter todoGroupId: (path)  
     - parameter body: (body)  
     - returns: Observable<TodosCreateTodoResponse>
     */
    open class func createTodo(todoGroupId: String, body: TodosCreateTodoRequest) -> Observable<TodosCreateTodoResponse> {
        return Observable.create { observer -> Disposable in
            createTodo(todoGroupId: todoGroupId, body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     Todos
     - POST /v1/todogroups/{todo_group_id}/todos
     - examples: [{contentType=application/json, example={
  "todo" : {
    "todo_group_id" : "todo_group_id",
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "title" : "title",
    "type" : { }
  }
}}]
     
     - parameter todoGroupId: (path)  
     - parameter body: (body)  

     - returns: RequestBuilder<TodosCreateTodoResponse> 
     */
    open class func createTodoWithRequestBuilder(todoGroupId: String, body: TodosCreateTodoRequest) -> RequestBuilder<TodosCreateTodoResponse> {
        var path = "/v1/todogroups/{todo_group_id}/todos"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosCreateTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

    /**
     TodoGruops
     
     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func createTodoGroup(body: TodosCreateTodoGroupRequest, completion: @escaping ((_ data: TodosCreateTodoGroupResponse?,_ error: Error?) -> Void)) {
        createTodoGroupWithRequestBuilder(body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**
     TodoGruops
     
     - parameter body: (body)  
     - returns: Observable<TodosCreateTodoGroupResponse>
     */
    open class func createTodoGroup(body: TodosCreateTodoGroupRequest) -> Observable<TodosCreateTodoGroupResponse> {
        return Observable.create { observer -> Disposable in
            createTodoGroup(body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     TodoGruops
     - POST /v1/todosgroups
     - examples: [{contentType=application/json, example={
  "todo_group" : {
    "editor" : [ "editor", "editor" ],
    "viewer" : [ "viewer", "viewer" ],
    "todo_group_id" : "todo_group_id",
    "updated_at" : "updated_at",
    "created_at" : "created_at",
    "description" : "description",
    "owners" : [ "owners", "owners" ],
    "todos" : [ {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    }, {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    } ],
    "title" : "title",
    "created_by" : "created_by"
  }
}}]
     
     - parameter body: (body)  

     - returns: RequestBuilder<TodosCreateTodoGroupResponse> 
     */
    open class func createTodoGroupWithRequestBuilder(body: TodosCreateTodoGroupRequest) -> RequestBuilder<TodosCreateTodoGroupResponse> {
        let path = "/v1/todosgroups"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosCreateTodoGroupResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "POST", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter accountId: (query)  (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func deleteTodo(todoGroupId: String, todoId: String, accountId: String? = nil, completion: @escaping ((_ data: TodosDeleteTodoResponse?,_ error: Error?) -> Void)) {
        deleteTodoWithRequestBuilder(todoGroupId: todoGroupId, todoId: todoId, accountId: accountId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter accountId: (query)  (optional)
     - returns: Observable<TodosDeleteTodoResponse>
     */
    open class func deleteTodo(todoGroupId: String, todoId: String, accountId: String? = nil) -> Observable<TodosDeleteTodoResponse> {
        return Observable.create { observer -> Disposable in
            deleteTodo(todoGroupId: todoGroupId, todoId: todoId, accountId: accountId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - DELETE /v1/todogroups/{todo_group_id}/todos/{todo_id}
     - examples: [{contentType=application/json, example={ }}]
     
     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter accountId: (query)  (optional)

     - returns: RequestBuilder<TodosDeleteTodoResponse> 
     */
    open class func deleteTodoWithRequestBuilder(todoGroupId: String, todoId: String, accountId: String? = nil) -> RequestBuilder<TodosDeleteTodoResponse> {
        var path = "/v1/todogroups/{todo_group_id}/todos/{todo_id}"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        path = path.replacingOccurrences(of: "{todo_id}", with: "\(todoId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)
        url?.queryItems = APIHelper.mapValuesToQueryItems(values:[
            "account_id": accountId
        ])
        

        let requestBuilder: RequestBuilder<TodosDeleteTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "DELETE", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func deleteTodoGroup(todoGroupId: String, accountId: String? = nil, completion: @escaping ((_ data: TodosDeleteTodoGroupResponse?,_ error: Error?) -> Void)) {
        deleteTodoGroupWithRequestBuilder(todoGroupId: todoGroupId, accountId: accountId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)
     - returns: Observable<TodosDeleteTodoGroupResponse>
     */
    open class func deleteTodoGroup(todoGroupId: String, accountId: String? = nil) -> Observable<TodosDeleteTodoGroupResponse> {
        return Observable.create { observer -> Disposable in
            deleteTodoGroup(todoGroupId: todoGroupId, accountId: accountId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - DELETE /v1/todosgroups/{todo_group_id}
     - examples: [{contentType=application/json, example={ }}]
     
     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)

     - returns: RequestBuilder<TodosDeleteTodoGroupResponse> 
     */
    open class func deleteTodoGroupWithRequestBuilder(todoGroupId: String, accountId: String? = nil) -> RequestBuilder<TodosDeleteTodoGroupResponse> {
        var path = "/v1/todosgroups/{todo_group_id}"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)
        url?.queryItems = APIHelper.mapValuesToQueryItems(values:[
            "account_id": accountId
        ])
        

        let requestBuilder: RequestBuilder<TodosDeleteTodoGroupResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "DELETE", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func getTodoGroup(todoGroupId: String, accountId: String? = nil, completion: @escaping ((_ data: TodosGetTodoGroupResponse?,_ error: Error?) -> Void)) {
        getTodoGroupWithRequestBuilder(todoGroupId: todoGroupId, accountId: accountId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)
     - returns: Observable<TodosGetTodoGroupResponse>
     */
    open class func getTodoGroup(todoGroupId: String, accountId: String? = nil) -> Observable<TodosGetTodoGroupResponse> {
        return Observable.create { observer -> Disposable in
            getTodoGroup(todoGroupId: todoGroupId, accountId: accountId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - GET /v1/todogroups/{todo_group_id}
     - examples: [{contentType=application/json, example={
  "todo_group" : {
    "editor" : [ "editor", "editor" ],
    "viewer" : [ "viewer", "viewer" ],
    "todo_group_id" : "todo_group_id",
    "updated_at" : "updated_at",
    "created_at" : "created_at",
    "description" : "description",
    "owners" : [ "owners", "owners" ],
    "todos" : [ {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    }, {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    } ],
    "title" : "title",
    "created_by" : "created_by"
  }
}}]
     
     - parameter todoGroupId: (path)  
     - parameter accountId: (query)  (optional)

     - returns: RequestBuilder<TodosGetTodoGroupResponse> 
     */
    open class func getTodoGroupWithRequestBuilder(todoGroupId: String, accountId: String? = nil) -> RequestBuilder<TodosGetTodoGroupResponse> {
        var path = "/v1/todogroups/{todo_group_id}"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)
        url?.queryItems = APIHelper.mapValuesToQueryItems(values:[
            "account_id": accountId
        ])
        

        let requestBuilder: RequestBuilder<TodosGetTodoGroupResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter accountId: (query)  (optional)
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func listTodoGroups(accountId: String? = nil, completion: @escaping ((_ data: TodosListTodoGroupsResponse?,_ error: Error?) -> Void)) {
        listTodoGroupsWithRequestBuilder(accountId: accountId).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter accountId: (query)  (optional)
     - returns: Observable<TodosListTodoGroupsResponse>
     */
    open class func listTodoGroups(accountId: String? = nil) -> Observable<TodosListTodoGroupsResponse> {
        return Observable.create { observer -> Disposable in
            listTodoGroups(accountId: accountId) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - GET /v1/todogroups
     - examples: [{contentType=application/json, example={
  "todo_groups" : [ {
    "editor" : [ "editor", "editor" ],
    "viewer" : [ "viewer", "viewer" ],
    "todo_group_id" : "todo_group_id",
    "updated_at" : "updated_at",
    "created_at" : "created_at",
    "description" : "description",
    "owners" : [ "owners", "owners" ],
    "todos" : [ {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    }, {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    } ],
    "title" : "title",
    "created_by" : "created_by"
  }, {
    "editor" : [ "editor", "editor" ],
    "viewer" : [ "viewer", "viewer" ],
    "todo_group_id" : "todo_group_id",
    "updated_at" : "updated_at",
    "created_at" : "created_at",
    "description" : "description",
    "owners" : [ "owners", "owners" ],
    "todos" : [ {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    }, {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    } ],
    "title" : "title",
    "created_by" : "created_by"
  } ]
}}]
     
     - parameter accountId: (query)  (optional)

     - returns: RequestBuilder<TodosListTodoGroupsResponse> 
     */
    open class func listTodoGroupsWithRequestBuilder(accountId: String? = nil) -> RequestBuilder<TodosListTodoGroupsResponse> {
        let path = "/v1/todogroups"
        let URLString = SwaggerClientAPI.basePath + path
        let parameters: [String:Any]? = nil

        let url = NSURLComponents(string: URLString)
        url?.queryItems = APIHelper.mapValuesToQueryItems(values:[
            "account_id": accountId
        ])
        

        let requestBuilder: RequestBuilder<TodosListTodoGroupsResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "GET", URLString: (url?.string ?? URLString), parameters: parameters, isBody: false)
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func updateTodo(todoGroupId: String, todoId: String, body: TodosUpdateTodoRequest, completion: @escaping ((_ data: TodosUpdateTodoResponse?,_ error: Error?) -> Void)) {
        updateTodoWithRequestBuilder(todoGroupId: todoGroupId, todoId: todoId, body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter body: (body)  
     - returns: Observable<TodosUpdateTodoResponse>
     */
    open class func updateTodo(todoGroupId: String, todoId: String, body: TodosUpdateTodoRequest) -> Observable<TodosUpdateTodoResponse> {
        return Observable.create { observer -> Disposable in
            updateTodo(todoGroupId: todoGroupId, todoId: todoId, body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - PUT /v1/todogroups/{todo_group_id}/todos/{todo_id}
     - examples: [{contentType=application/json, example={
  "todo" : {
    "todo_group_id" : "todo_group_id",
    "todo_id" : "todo_id",
    "done_at" : "done_at",
    "created_at" : "created_at",
    "description" : "description",
    "title" : "title",
    "type" : { }
  }
}}]
     
     - parameter todoGroupId: (path)  
     - parameter todoId: (path)  
     - parameter body: (body)  

     - returns: RequestBuilder<TodosUpdateTodoResponse> 
     */
    open class func updateTodoWithRequestBuilder(todoGroupId: String, todoId: String, body: TodosUpdateTodoRequest) -> RequestBuilder<TodosUpdateTodoResponse> {
        var path = "/v1/todogroups/{todo_group_id}/todos/{todo_id}"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        path = path.replacingOccurrences(of: "{todo_id}", with: "\(todoId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosUpdateTodoResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "PUT", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter body: (body)  
     - parameter completion: completion handler to receive the data and the error objects
     */
    open class func updateTodoGroup(todoGroupId: String, body: TodosUpdateTodoGroupRequest, completion: @escaping ((_ data: TodosUpdateTodoGroupResponse?,_ error: Error?) -> Void)) {
        updateTodoGroupWithRequestBuilder(todoGroupId: todoGroupId, body: body).execute { (response, error) -> Void in
            completion(response?.body, error);
        }
    }

    /**

     - parameter todoGroupId: (path)  
     - parameter body: (body)  
     - returns: Observable<TodosUpdateTodoGroupResponse>
     */
    open class func updateTodoGroup(todoGroupId: String, body: TodosUpdateTodoGroupRequest) -> Observable<TodosUpdateTodoGroupResponse> {
        return Observable.create { observer -> Disposable in
            updateTodoGroup(todoGroupId: todoGroupId, body: body) { data, error in
                if let error = error {
                    observer.on(.error(error))
                } else {
                    observer.on(.next(data!))
                }
                observer.on(.completed)
            }
            return Disposables.create()
        }
    }

    /**
     - PUT /v1/todosgroups/{todo_group_id}
     - examples: [{contentType=application/json, example={
  "todo_group" : {
    "editor" : [ "editor", "editor" ],
    "viewer" : [ "viewer", "viewer" ],
    "todo_group_id" : "todo_group_id",
    "updated_at" : "updated_at",
    "created_at" : "created_at",
    "description" : "description",
    "owners" : [ "owners", "owners" ],
    "todos" : [ {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    }, {
      "todo_group_id" : "todo_group_id",
      "todo_id" : "todo_id",
      "done_at" : "done_at",
      "created_at" : "created_at",
      "description" : "description",
      "title" : "title",
      "type" : { }
    } ],
    "title" : "title",
    "created_by" : "created_by"
  }
}}]
     
     - parameter todoGroupId: (path)  
     - parameter body: (body)  

     - returns: RequestBuilder<TodosUpdateTodoGroupResponse> 
     */
    open class func updateTodoGroupWithRequestBuilder(todoGroupId: String, body: TodosUpdateTodoGroupRequest) -> RequestBuilder<TodosUpdateTodoGroupResponse> {
        var path = "/v1/todosgroups/{todo_group_id}"
        path = path.replacingOccurrences(of: "{todo_group_id}", with: "\(todoGroupId)", options: .literal, range: nil)
        let URLString = SwaggerClientAPI.basePath + path
        let parameters = JSONEncodingHelper.encodingParameters(forEncodableObject: body)

        let url = NSURLComponents(string: URLString)


        let requestBuilder: RequestBuilder<TodosUpdateTodoGroupResponse>.Type = SwaggerClientAPI.requestBuilderFactory.getBuilder()

        return requestBuilder.init(method: "PUT", URLString: (url?.string ?? URLString), parameters: parameters, isBody: true)
    }

}
