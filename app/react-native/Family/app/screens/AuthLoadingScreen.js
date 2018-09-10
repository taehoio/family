import React from 'react';
import {
    ActivityIndicator,
    StatusBar,
    StyleSheet,
    View,
} from 'react-native';

import {
    getUserTokenAsync,
    refreshAuthToken,
    getRefreshTokenAsync,
    startRefreshAuthTokenInterval,
} from "../lib/auth";


export default class AuthLoadingScreen extends React.Component {
    constructor() {
        super();
        this._bootstrapAsync();
    }

    // Fetch the token from storage then navigate to our appropriate place.
    _bootstrapAsync = async () => {
        const accountId = await getUserTokenAsync();
        if (accountId) {
            try {
                await refreshAuthToken(await getRefreshTokenAsync());
            } catch (err) {
                this.props.navigation.navigate('Auth');
                return;
            }

            this.props.navigation.navigate('App');
            startRefreshAuthTokenInterval();
            return;
        }

        this.props.navigation.navigate('Auth');
    };

    render() {
        return (
            <View style={styles.container}>
                <ActivityIndicator />
                <StatusBar barStyle="default" />
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
});