//const Api = require("api");
//
//Api.ApiClient.instance.basePath = "http://localhost:3000";
//
//const api = new Api.AccountsServiceApi();
//const body = new Api.AccountsLogInRequest(); // {AccountsLogInRequest}
//
//const callback = function (error, data, response) {
//    if (error) {
//        console.error(error);
//    } else {
//        console.log('API called successfully. Returned data: ' + data);
//    }
//};
//api.logIn(body, callback);

//import Swagger from 'swagger-client';
const Swagger = require('swagger-client');

Swagger('http://localhost:3000/swagger.json')
    .then( client => {
        //console.log(client);
        client.url
        client.apis.AccountsService.LogIn({body:{}}).then(console.log).catch(console.log);
    });
