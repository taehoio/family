import {
    SecureStore,
} from 'expo';

import { signIn, signOut, startRefreshAuthTokenInterval } from './api';


const accountIdKey = 'account_id';

export async function signInAsync(email, password) {
    res = await signIn(email, password);
    await SecureStore.setItemAsync(accountIdKey, res.account_id);
}

export async function getUserTokenAsync() {
    return await SecureStore.getItemAsync(accountIdKey);
}

export async function signOutAsync() {
    signOut();
    await SecureStore.deleteItemAsync(accountIdKey);
}

export { startRefreshAuthTokenInterval };
