const Api = require("api");



let basePath = 'http://localhost:3000';
if (process.env.NODE_ENV === 'staging') {
    basePath = 'https://family-staging.taeho.io'
} else if (process.env.NODE_ENV === 'production') {
    basePath = 'https://family.taeho.io'
}
Api.ApiClient.instance.basePath = basePath;

let loggedInAccount = new Api.AccountsLogInResponse();
function updateAuthentication() {
    Api.ApiClient.instance.defaultHeaders = {
        'Authorization': 'Bearer ' + loggedInAccount.access_token,
    };
}

export default Api;

export { Api };

export function isLoggedIn() {
    if (loggedInAccount.access_token) {
        return true
    }
    return false
}

let refreshAuthTokenIntervalId = null;
function startRefreshAuthTokenInterval() {
   refreshAuthTokenIntervalId = setInterval(refreshAuthToken, 10000);
}

function clearRefreshAuthTokenInterval() {
    clearInterval(refreshAuthTokenIntervalId);
}

function refreshAuthToken() {
    const api = new Api.AuthServiceApi();
    const body = Api.AuthRefreshRequest.constructFromObject({
        refresh_token: loggedInAccount.refresh_token,
    });

    api.refresh(body)
        .then(res => {
            loggedInAccount.access_token = res.access_token;
            loggedInAccount.expires_in = res.expires_in;
            updateAuthentication();

            console.log(refreshAuthTokenIntervalId);
            console.log(res);
            console.log(loggedInAccount);
        })
        .catch(err => {
            //if (refreshAuthTokenIntervalId !== null) {
            //    clearRefreshAuthTokenInterval();
            //    refreshAuthTokenIntervalId = null;
            //}

            loggedInAccount = new Api.AccountsLogInResponse();

            console.log(refreshAuthTokenIntervalId);
            console.log(err);
            console.log(loggedInAccount);
        });
}

export function logIn(email, password) {
    const api = new Api.AccountsServiceApi();
    const body = Api.AccountsLogInRequest.constructFromObject({
        auth_type: Api.AccountsAuthType.EMAIL,
        email,
        password,
    });

    return new Promise((resolve, reject) => {
        api.logIn(body)
            .then(res => {
                loggedInAccount = res;
                updateAuthentication();

                startRefreshAuthTokenInterval();

                resolve(res);
            })
            .catch(err => {
                reject(err);
            });
    });
}