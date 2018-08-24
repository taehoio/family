//
// TodosListTodosResponse.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class TodosListTodosResponse: Codable {

    public var todos: [TodosTodo]?


    
    public init(todos: [TodosTodo]?) {
        self.todos = todos
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(todos, forKey: "todos")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        todos = try container.decodeIfPresent([TodosTodo].self, forKey: "todos")
    }
}

