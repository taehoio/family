//
// TodosGetTodoGroupResponse.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class TodosGetTodoGroupResponse: Codable {

    public var todoGroup: TodosTodoGroup?


    
    public init(todoGroup: TodosTodoGroup?) {
        self.todoGroup = todoGroup
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(todoGroup, forKey: "todo_group")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        todoGroup = try container.decodeIfPresent(TodosTodoGroup.self, forKey: "todo_group")
    }
}

