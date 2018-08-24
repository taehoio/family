//
// TodosTodo.swift
//
// Generated by swagger-codegen
// https://github.com/swagger-api/swagger-codegen
//

import Foundation



open class TodosTodo: Codable {

    public var assignedTo: String?
    public var createdAt: String?
    public var description: String?
    public var doneAt: String?
    public var dueAt: String?
    public var order: String?
    public var parentId: String?
    public var parentType: TodosParentType?
    public var priority: TodosPriority?
    public var status: TodosStatus?
    public var subTasks: [TodosTodo]?
    public var title: String?
    public var todoId: String?
    public var updatedAt: String?


    
    public init(assignedTo: String?, createdAt: String?, description: String?, doneAt: String?, dueAt: String?, order: String?, parentId: String?, parentType: TodosParentType?, priority: TodosPriority?, status: TodosStatus?, subTasks: [TodosTodo]?, title: String?, todoId: String?, updatedAt: String?) {
        self.assignedTo = assignedTo
        self.createdAt = createdAt
        self.description = description
        self.doneAt = doneAt
        self.dueAt = dueAt
        self.order = order
        self.parentId = parentId
        self.parentType = parentType
        self.priority = priority
        self.status = status
        self.subTasks = subTasks
        self.title = title
        self.todoId = todoId
        self.updatedAt = updatedAt
    }
    

    // Encodable protocol methods

    public func encode(to encoder: Encoder) throws {

        var container = encoder.container(keyedBy: String.self)

        try container.encodeIfPresent(assignedTo, forKey: "assigned_to")
        try container.encodeIfPresent(createdAt, forKey: "created_at")
        try container.encodeIfPresent(description, forKey: "description")
        try container.encodeIfPresent(doneAt, forKey: "done_at")
        try container.encodeIfPresent(dueAt, forKey: "due_at")
        try container.encodeIfPresent(order, forKey: "order")
        try container.encodeIfPresent(parentId, forKey: "parent_id")
        try container.encodeIfPresent(parentType, forKey: "parent_type")
        try container.encodeIfPresent(priority, forKey: "priority")
        try container.encodeIfPresent(status, forKey: "status")
        try container.encodeIfPresent(subTasks, forKey: "sub_tasks")
        try container.encodeIfPresent(title, forKey: "title")
        try container.encodeIfPresent(todoId, forKey: "todo_id")
        try container.encodeIfPresent(updatedAt, forKey: "updated_at")
    }

    // Decodable protocol methods

    public required init(from decoder: Decoder) throws {
        let container = try decoder.container(keyedBy: String.self)

        assignedTo = try container.decodeIfPresent(String.self, forKey: "assigned_to")
        createdAt = try container.decodeIfPresent(String.self, forKey: "created_at")
        description = try container.decodeIfPresent(String.self, forKey: "description")
        doneAt = try container.decodeIfPresent(String.self, forKey: "done_at")
        dueAt = try container.decodeIfPresent(String.self, forKey: "due_at")
        order = try container.decodeIfPresent(String.self, forKey: "order")
        parentId = try container.decodeIfPresent(String.self, forKey: "parent_id")
        parentType = try container.decodeIfPresent(TodosParentType.self, forKey: "parent_type")
        priority = try container.decodeIfPresent(TodosPriority.self, forKey: "priority")
        status = try container.decodeIfPresent(TodosStatus.self, forKey: "status")
        subTasks = try container.decodeIfPresent([TodosTodo].self, forKey: "sub_tasks")
        title = try container.decodeIfPresent(String.self, forKey: "title")
        todoId = try container.decodeIfPresent(String.self, forKey: "todo_id")
        updatedAt = try container.decodeIfPresent(String.self, forKey: "updated_at")
    }
}

