var Api = require('api');

Api.ApiClient.instance.basePath = "http://localhost:3000";

var api = new Api.AccountsServiceApi();
var body = new Api.AccountsLogInRequest(); // {AccountsLogInRequest} 

var callback = function (error, data, response) {
    if (error) {
        console.error(error);
    } else {
        console.log('API called successfully. Returned data: ' + data);
    }
};
api.logIn(body, callback);
