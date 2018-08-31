/**
 * API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 0.0.1
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/TodogroupsTodoGroup'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./TodogroupsTodoGroup'));
  } else {
    // Browser globals (root is window)
    if (!root.Api) {
      root.Api = {};
    }
    root.Api.TodogroupsCreateTodoGroupResponse = factory(root.Api.ApiClient, root.Api.TodogroupsTodoGroup);
  }
}(this, function(ApiClient, TodogroupsTodoGroup) {
  'use strict';




  /**
   * The TodogroupsCreateTodoGroupResponse model module.
   * @module model/TodogroupsCreateTodoGroupResponse
   * @version 0.0.1
   */

  /**
   * Constructs a new <code>TodogroupsCreateTodoGroupResponse</code>.
   * @alias module:model/TodogroupsCreateTodoGroupResponse
   * @class
   */
  var exports = function() {
    var _this = this;


  };

  /**
   * Constructs a <code>TodogroupsCreateTodoGroupResponse</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/TodogroupsCreateTodoGroupResponse} obj Optional instance to populate.
   * @return {module:model/TodogroupsCreateTodoGroupResponse} The populated <code>TodogroupsCreateTodoGroupResponse</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('todo_group')) {
        obj['todo_group'] = TodogroupsTodoGroup.constructFromObject(data['todo_group']);
      }
    }
    return obj;
  }

  /**
   * @member {module:model/TodogroupsTodoGroup} todo_group
   */
  exports.prototype['todo_group'] = undefined;



  return exports;
}));


