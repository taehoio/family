import {
    SecureStore,
} from 'expo';

import {
    refreshAuthToken,
    signIn,
    signOut,
    startRefreshAuthTokenInterval,
} from './api';


const accountIdKey = 'account_id';
const refreshTokenKey = 'refresh_token';

export async function signInAsync(email, password) {
    res = await signIn(email, password);
    await SecureStore.setItemAsync(accountIdKey, res.account_id);
    await SecureStore.setItemAsync(refreshTokenKey, res.refresh_token);
}

export async function getUserTokenAsync() {
    return await SecureStore.getItemAsync(accountIdKey);
}

export async function getRefreshTokenAsync() {
    return await SecureStore.getItemAsync(refreshTokenKey);
}

export async function signOutAsync() {
    signOut();
    await SecureStore.deleteItemAsync(accountIdKey);
}

export {
    refreshAuthToken,
    startRefreshAuthTokenInterval,
};
