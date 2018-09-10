import React from 'react';
import {
    StyleSheet,
    View,
} from 'react-native';
import {
    Button,
    FormLabel,
    FormInput,
    FormValidationMessage,
    Text,
} from 'react-native-elements';

import { signInAsync } from "../lib/auth";

export default class SignInScreen extends React.Component {
    static navigationOptions = {
        title: 'Sign Up',
    };

    render() {
        return (
            <View style={styles.container}>
                <View style={styles.form}>
                    <FormLabel>Email</FormLabel>
                    <FormInput
                        textInputRef={'email'}
                        autoCorrect={false}
                        keyboardType={'email-address'}
                        autoFocus={true}
                        keyboardAppearance="light"
                        returnKeyType="next"
                    />
                    <FormValidationMessage>Required</FormValidationMessage>

                    <FormLabel>Password</FormLabel>
                    <FormInput secureTextEntry={true} />
                    <FormValidationMessage>Required</FormValidationMessage>

                    <Button style={styles.submitButton} title="Sign Up" onPress={this._signUpAsync} />
                </View>
            </View>
        );
    }

    _signUpAsync = () => {
        this.props.navigation.navigate('SignUp');
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
    },
    logo: {
        alignItems: 'center',
        justifyContent: 'center',
        margin: 10,
    },
    form: {
    },
    submitButton: {
        marginTop: 20,
    }
});
