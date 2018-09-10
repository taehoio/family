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
export function startRefreshAuthTokenInterval() {
   refreshAuthTokenIntervalId = setInterval(refreshAuthToken, 10000);
}

function clearRefreshAuthTokenInterval() {
    if (refreshAuthTokenIntervalId) {
        clearInterval(refreshAuthTokenIntervalId);
        refreshAuthTokenIntervalId = null;
    }
}

export function refreshAuthToken(refreshToken) {
    if (refreshToken) {
        loggedInAccount.refresh_token = refreshToken;
    } else {
        refreshToken = loggedInAccount.refresh_token;
    }

    const api = new Api.AuthServiceApi();
    const body = Api.AuthRefreshRequest.constructFromObject({
        refresh_token: refreshToken,
    });

    return new Promise((resolve, reject) => {
        api.refresh(body)
            .then(res => {
                loggedInAccount.access_token = res.access_token;
                loggedInAccount.expires_in = res.expires_in;
                updateAuthentication();

                resolve();
            })
            .catch(err => {
                clearRefreshAuthTokenInterval();

                loggedInAccount = new Api.AccountsLogInResponse();

                reject(err);
            });
    });
}

export function signIn(email, password) {
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

export function signOut() {
    if (isLoggedIn()) {
        clearRefreshAuthTokenInterval();
    }
}