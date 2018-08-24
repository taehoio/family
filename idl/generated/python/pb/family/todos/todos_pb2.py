# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pb/family/todos/todos.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='pb/family/todos/todos.proto',
  package='pb.family.todos',
  syntax='proto3',
  serialized_pb=_b('\n\x1bpb/family/todos/todos.proto\x12\x0fpb.family.todos\x1a\x1cgoogle/api/annotations.proto\"\xed\x02\n\x04Todo\x12\x0f\n\x07todo_id\x18\x01 \x01(\t\x12\x30\n\x0bparent_type\x18\x02 \x01(\x0e\x32\x1b.pb.family.todos.ParentType\x12\x11\n\tparent_id\x18\x03 \x01(\t\x12\r\n\x05title\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\'\n\x06status\x18\x06 \x01(\x0e\x32\x17.pb.family.todos.Status\x12\r\n\x05order\x18\x07 \x01(\t\x12\x13\n\x0b\x61ssigned_to\x18\x08 \x01(\t\x12+\n\x08priority\x18\t \x01(\x0e\x32\x19.pb.family.todos.Priority\x12(\n\tsub_tasks\x18\n \x03(\x0b\x32\x15.pb.family.todos.Todo\x12\x12\n\ncreated_at\x18\x0b \x01(\x03\x12\x12\n\nupdated_at\x18\x0c \x01(\x03\x12\x0f\n\x07\x64one_at\x18\r \x01(\x03\x12\x0e\n\x06\x64ue_at\x18\x0e \x01(\x03\"]\n\x11\x43reateTodoRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x0f\n\x07todo_id\x18\x02 \x01(\t\x12#\n\x04todo\x18\x03 \x01(\x0b\x32\x15.pb.family.todos.Todo\"9\n\x12\x43reateTodoResponse\x12#\n\x04todo\x18\x01 \x01(\x0b\x32\x15.pb.family.todos.Todo\"5\n\x0eGetTodoRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x0f\n\x07todo_id\x18\x02 \x01(\t\"6\n\x0fGetTodoResponse\x12#\n\x04todo\x18\x01 \x01(\x0b\x32\x15.pb.family.todos.Todo\"k\n\x10ListTodosRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x30\n\x0bparent_type\x18\x02 \x01(\x0e\x32\x1b.pb.family.todos.ParentType\x12\x11\n\tparent_id\x18\x03 \x01(\t\"9\n\x11ListTodosResponse\x12$\n\x05todos\x18\x01 \x03(\x0b\x32\x15.pb.family.todos.Todo\"]\n\x11UpdateTodoRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x0f\n\x07todo_id\x18\x02 \x01(\t\x12#\n\x04todo\x18\x03 \x01(\x0b\x32\x15.pb.family.todos.Todo\"9\n\x12UpdateTodoResponse\x12#\n\x04todo\x18\x01 \x01(\x0b\x32\x15.pb.family.todos.Todo\"8\n\x11\x44\x65leteTodoRequest\x12\x12\n\naccount_id\x18\x01 \x01(\t\x12\x0f\n\x07todo_id\x18\x02 \x01(\t\"\x14\n\x12\x44\x65leteTodoResponse*>\n\x06Status\x12\x0f\n\x0bSTATUS_TODO\x10\x00\x12\x0f\n\x0bSTATUS_DONE\x10\x01\x12\x12\n\x0eSTATUS_PENDING\x10\x02*l\n\x08Priority\x12\x11\n\rPRIORITY_NONE\x10\x00\x12\x10\n\x0cPRIORITY_LOW\x10\n\x12\x13\n\x0fPRIORITY_MEDIUM\x10\x1e\x12\x11\n\rPRIORITY_HIGH\x10\x32\x12\x13\n\x0fPRIORITY_URGENT\x10\x64*>\n\nParentType\x12\x1a\n\x16PARENT_TYPE_TODO_GROUP\x10\x00\x12\x14\n\x10PARENT_TYPE_TODO\x10\x01\x32\xb8\x04\n\x0cTodosService\x12k\n\nCreateTodo\x12\".pb.family.todos.CreateTodoRequest\x1a#.pb.family.todos.CreateTodoResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\"\t/v1/todos:\x01*\x12i\n\x07GetTodo\x12\x1f.pb.family.todos.GetTodoRequest\x1a .pb.family.todos.GetTodoResponse\"\x1b\x82\xd3\xe4\x93\x02\x15\x12\x13/v1/todos/{todo_id}\x12\x65\n\tListTodos\x12!.pb.family.todos.ListTodosRequest\x1a\".pb.family.todos.ListTodosResponse\"\x11\x82\xd3\xe4\x93\x02\x0b\x12\t/v1/todos\x12u\n\nUpdateTodo\x12\".pb.family.todos.UpdateTodoRequest\x1a#.pb.family.todos.UpdateTodoResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x1a\x13/v1/todos/{todo_id}:\x01*\x12r\n\nDeleteTodo\x12\".pb.family.todos.DeleteTodoRequest\x1a#.pb.family.todos.DeleteTodoResponse\"\x1b\x82\xd3\xe4\x93\x02\x15*\x13/v1/todos/{todo_id}BCZAgithub.com/taeho-io/family/idl/generated/go/pb/family/todos;todosb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,])

_STATUS = _descriptor.EnumDescriptor(
  name='Status',
  full_name='pb.family.todos.Status',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STATUS_TODO', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STATUS_DONE', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STATUS_PENDING', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1113,
  serialized_end=1175,
)
_sym_db.RegisterEnumDescriptor(_STATUS)

Status = enum_type_wrapper.EnumTypeWrapper(_STATUS)
_PRIORITY = _descriptor.EnumDescriptor(
  name='Priority',
  full_name='pb.family.todos.Priority',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PRIORITY_NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIORITY_LOW', index=1, number=10,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIORITY_MEDIUM', index=2, number=30,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIORITY_HIGH', index=3, number=50,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PRIORITY_URGENT', index=4, number=100,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1177,
  serialized_end=1285,
)
_sym_db.RegisterEnumDescriptor(_PRIORITY)

Priority = enum_type_wrapper.EnumTypeWrapper(_PRIORITY)
_PARENTTYPE = _descriptor.EnumDescriptor(
  name='ParentType',
  full_name='pb.family.todos.ParentType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='PARENT_TYPE_TODO_GROUP', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PARENT_TYPE_TODO', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=1287,
  serialized_end=1349,
)
_sym_db.RegisterEnumDescriptor(_PARENTTYPE)

ParentType = enum_type_wrapper.EnumTypeWrapper(_PARENTTYPE)
STATUS_TODO = 0
STATUS_DONE = 1
STATUS_PENDING = 2
PRIORITY_NONE = 0
PRIORITY_LOW = 10
PRIORITY_MEDIUM = 30
PRIORITY_HIGH = 50
PRIORITY_URGENT = 100
PARENT_TYPE_TODO_GROUP = 0
PARENT_TYPE_TODO = 1



_TODO = _descriptor.Descriptor(
  name='Todo',
  full_name='pb.family.todos.Todo',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='todo_id', full_name='pb.family.todos.Todo.todo_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_type', full_name='pb.family.todos.Todo.parent_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_id', full_name='pb.family.todos.Todo.parent_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='pb.family.todos.Todo.title', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='pb.family.todos.Todo.description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='status', full_name='pb.family.todos.Todo.status', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order', full_name='pb.family.todos.Todo.order', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='assigned_to', full_name='pb.family.todos.Todo.assigned_to', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='priority', full_name='pb.family.todos.Todo.priority', index=8,
      number=9, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sub_tasks', full_name='pb.family.todos.Todo.sub_tasks', index=9,
      number=10, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='pb.family.todos.Todo.created_at', index=10,
      number=11, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='pb.family.todos.Todo.updated_at', index=11,
      number=12, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='done_at', full_name='pb.family.todos.Todo.done_at', index=12,
      number=13, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='due_at', full_name='pb.family.todos.Todo.due_at', index=13,
      number=14, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=79,
  serialized_end=444,
)


_CREATETODOREQUEST = _descriptor.Descriptor(
  name='CreateTodoRequest',
  full_name='pb.family.todos.CreateTodoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='pb.family.todos.CreateTodoRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo_id', full_name='pb.family.todos.CreateTodoRequest.todo_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo', full_name='pb.family.todos.CreateTodoRequest.todo', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=446,
  serialized_end=539,
)


_CREATETODORESPONSE = _descriptor.Descriptor(
  name='CreateTodoResponse',
  full_name='pb.family.todos.CreateTodoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='todo', full_name='pb.family.todos.CreateTodoResponse.todo', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=541,
  serialized_end=598,
)


_GETTODOREQUEST = _descriptor.Descriptor(
  name='GetTodoRequest',
  full_name='pb.family.todos.GetTodoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='pb.family.todos.GetTodoRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo_id', full_name='pb.family.todos.GetTodoRequest.todo_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=600,
  serialized_end=653,
)


_GETTODORESPONSE = _descriptor.Descriptor(
  name='GetTodoResponse',
  full_name='pb.family.todos.GetTodoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='todo', full_name='pb.family.todos.GetTodoResponse.todo', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=655,
  serialized_end=709,
)


_LISTTODOSREQUEST = _descriptor.Descriptor(
  name='ListTodosRequest',
  full_name='pb.family.todos.ListTodosRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='pb.family.todos.ListTodosRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_type', full_name='pb.family.todos.ListTodosRequest.parent_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_id', full_name='pb.family.todos.ListTodosRequest.parent_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=711,
  serialized_end=818,
)


_LISTTODOSRESPONSE = _descriptor.Descriptor(
  name='ListTodosResponse',
  full_name='pb.family.todos.ListTodosResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='todos', full_name='pb.family.todos.ListTodosResponse.todos', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=820,
  serialized_end=877,
)


_UPDATETODOREQUEST = _descriptor.Descriptor(
  name='UpdateTodoRequest',
  full_name='pb.family.todos.UpdateTodoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='pb.family.todos.UpdateTodoRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo_id', full_name='pb.family.todos.UpdateTodoRequest.todo_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo', full_name='pb.family.todos.UpdateTodoRequest.todo', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=879,
  serialized_end=972,
)


_UPDATETODORESPONSE = _descriptor.Descriptor(
  name='UpdateTodoResponse',
  full_name='pb.family.todos.UpdateTodoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='todo', full_name='pb.family.todos.UpdateTodoResponse.todo', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=974,
  serialized_end=1031,
)


_DELETETODOREQUEST = _descriptor.Descriptor(
  name='DeleteTodoRequest',
  full_name='pb.family.todos.DeleteTodoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='account_id', full_name='pb.family.todos.DeleteTodoRequest.account_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='todo_id', full_name='pb.family.todos.DeleteTodoRequest.todo_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1033,
  serialized_end=1089,
)


_DELETETODORESPONSE = _descriptor.Descriptor(
  name='DeleteTodoResponse',
  full_name='pb.family.todos.DeleteTodoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1091,
  serialized_end=1111,
)

_TODO.fields_by_name['parent_type'].enum_type = _PARENTTYPE
_TODO.fields_by_name['status'].enum_type = _STATUS
_TODO.fields_by_name['priority'].enum_type = _PRIORITY
_TODO.fields_by_name['sub_tasks'].message_type = _TODO
_CREATETODOREQUEST.fields_by_name['todo'].message_type = _TODO
_CREATETODORESPONSE.fields_by_name['todo'].message_type = _TODO
_GETTODORESPONSE.fields_by_name['todo'].message_type = _TODO
_LISTTODOSREQUEST.fields_by_name['parent_type'].enum_type = _PARENTTYPE
_LISTTODOSRESPONSE.fields_by_name['todos'].message_type = _TODO
_UPDATETODOREQUEST.fields_by_name['todo'].message_type = _TODO
_UPDATETODORESPONSE.fields_by_name['todo'].message_type = _TODO
DESCRIPTOR.message_types_by_name['Todo'] = _TODO
DESCRIPTOR.message_types_by_name['CreateTodoRequest'] = _CREATETODOREQUEST
DESCRIPTOR.message_types_by_name['CreateTodoResponse'] = _CREATETODORESPONSE
DESCRIPTOR.message_types_by_name['GetTodoRequest'] = _GETTODOREQUEST
DESCRIPTOR.message_types_by_name['GetTodoResponse'] = _GETTODORESPONSE
DESCRIPTOR.message_types_by_name['ListTodosRequest'] = _LISTTODOSREQUEST
DESCRIPTOR.message_types_by_name['ListTodosResponse'] = _LISTTODOSRESPONSE
DESCRIPTOR.message_types_by_name['UpdateTodoRequest'] = _UPDATETODOREQUEST
DESCRIPTOR.message_types_by_name['UpdateTodoResponse'] = _UPDATETODORESPONSE
DESCRIPTOR.message_types_by_name['DeleteTodoRequest'] = _DELETETODOREQUEST
DESCRIPTOR.message_types_by_name['DeleteTodoResponse'] = _DELETETODORESPONSE
DESCRIPTOR.enum_types_by_name['Status'] = _STATUS
DESCRIPTOR.enum_types_by_name['Priority'] = _PRIORITY
DESCRIPTOR.enum_types_by_name['ParentType'] = _PARENTTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Todo = _reflection.GeneratedProtocolMessageType('Todo', (_message.Message,), dict(
  DESCRIPTOR = _TODO,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.Todo)
  ))
_sym_db.RegisterMessage(Todo)

CreateTodoRequest = _reflection.GeneratedProtocolMessageType('CreateTodoRequest', (_message.Message,), dict(
  DESCRIPTOR = _CREATETODOREQUEST,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.CreateTodoRequest)
  ))
_sym_db.RegisterMessage(CreateTodoRequest)

CreateTodoResponse = _reflection.GeneratedProtocolMessageType('CreateTodoResponse', (_message.Message,), dict(
  DESCRIPTOR = _CREATETODORESPONSE,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.CreateTodoResponse)
  ))
_sym_db.RegisterMessage(CreateTodoResponse)

GetTodoRequest = _reflection.GeneratedProtocolMessageType('GetTodoRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETTODOREQUEST,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.GetTodoRequest)
  ))
_sym_db.RegisterMessage(GetTodoRequest)

GetTodoResponse = _reflection.GeneratedProtocolMessageType('GetTodoResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETTODORESPONSE,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.GetTodoResponse)
  ))
_sym_db.RegisterMessage(GetTodoResponse)

ListTodosRequest = _reflection.GeneratedProtocolMessageType('ListTodosRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTTODOSREQUEST,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.ListTodosRequest)
  ))
_sym_db.RegisterMessage(ListTodosRequest)

ListTodosResponse = _reflection.GeneratedProtocolMessageType('ListTodosResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTTODOSRESPONSE,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.ListTodosResponse)
  ))
_sym_db.RegisterMessage(ListTodosResponse)

UpdateTodoRequest = _reflection.GeneratedProtocolMessageType('UpdateTodoRequest', (_message.Message,), dict(
  DESCRIPTOR = _UPDATETODOREQUEST,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.UpdateTodoRequest)
  ))
_sym_db.RegisterMessage(UpdateTodoRequest)

UpdateTodoResponse = _reflection.GeneratedProtocolMessageType('UpdateTodoResponse', (_message.Message,), dict(
  DESCRIPTOR = _UPDATETODORESPONSE,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.UpdateTodoResponse)
  ))
_sym_db.RegisterMessage(UpdateTodoResponse)

DeleteTodoRequest = _reflection.GeneratedProtocolMessageType('DeleteTodoRequest', (_message.Message,), dict(
  DESCRIPTOR = _DELETETODOREQUEST,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.DeleteTodoRequest)
  ))
_sym_db.RegisterMessage(DeleteTodoRequest)

DeleteTodoResponse = _reflection.GeneratedProtocolMessageType('DeleteTodoResponse', (_message.Message,), dict(
  DESCRIPTOR = _DELETETODORESPONSE,
  __module__ = 'pb.family.todos.todos_pb2'
  # @@protoc_insertion_point(class_scope:pb.family.todos.DeleteTodoResponse)
  ))
_sym_db.RegisterMessage(DeleteTodoResponse)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('ZAgithub.com/taeho-io/family/idl/generated/go/pb/family/todos;todos'))

_TODOSSERVICE = _descriptor.ServiceDescriptor(
  name='TodosService',
  full_name='pb.family.todos.TodosService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=1352,
  serialized_end=1920,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateTodo',
    full_name='pb.family.todos.TodosService.CreateTodo',
    index=0,
    containing_service=None,
    input_type=_CREATETODOREQUEST,
    output_type=_CREATETODORESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\016\"\t/v1/todos:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='GetTodo',
    full_name='pb.family.todos.TodosService.GetTodo',
    index=1,
    containing_service=None,
    input_type=_GETTODOREQUEST,
    output_type=_GETTODORESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\025\022\023/v1/todos/{todo_id}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListTodos',
    full_name='pb.family.todos.TodosService.ListTodos',
    index=2,
    containing_service=None,
    input_type=_LISTTODOSREQUEST,
    output_type=_LISTTODOSRESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\013\022\t/v1/todos')),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateTodo',
    full_name='pb.family.todos.TodosService.UpdateTodo',
    index=3,
    containing_service=None,
    input_type=_UPDATETODOREQUEST,
    output_type=_UPDATETODORESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\030\032\023/v1/todos/{todo_id}:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='DeleteTodo',
    full_name='pb.family.todos.TodosService.DeleteTodo',
    index=4,
    containing_service=None,
    input_type=_DELETETODOREQUEST,
    output_type=_DELETETODORESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\025*\023/v1/todos/{todo_id}')),
  ),
])
_sym_db.RegisterServiceDescriptor(_TODOSSERVICE)

DESCRIPTOR.services_by_name['TodosService'] = _TODOSSERVICE

# @@protoc_insertion_point(module_scope)
