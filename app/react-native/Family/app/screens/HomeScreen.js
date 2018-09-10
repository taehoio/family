import React from 'react';
import {
    Button,
    StyleSheet,
    View,
} from 'react-native';

import { signOutAsync } from "../lib/auth";


export default class HomeScreen extends React.Component {
    static navigationOptions = {
        title: 'Home',
    };

    render() {
        return (
            <View style={styles.container}>
                <Button title="NoteListScreen" onPress={this._showNoteListScreen} />
                <Button title="Show me more of the app" onPress={this._showMoreApp} />
                <Button title="Actually, sign me out :)" onPress={this._signOutAsync} />
            </View>
        );
    }

    _showNoteListScreen = () => {
        this.props.navigation.navigate('NoteList');
    };

    _showMoreApp = () => {
        this.props.navigation.navigate('Other');
    };

    _signOutAsync = async () => {
        await signOutAsync();
        this.props.navigation.navigate('Auth');
    };
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'center',
    },
});
