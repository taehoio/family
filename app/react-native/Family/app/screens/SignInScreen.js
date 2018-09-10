import React from 'react';
import {
    ActivityIndicator,
    StyleSheet,
    View,
} from 'react-native';
import {
    Button,
    FormLabel,
    FormInput,
    Text,
} from 'react-native-elements';

import { signInAsync } from "../lib/auth";

export default class SignInScreen extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            email: '',
            password: '',
            isLoading: false,
        };
    }

    static navigationOptions = {
        title: 'Sign In',
    };

    render() {
        const {
            isLoading,
        } = this.state;

        return (
            <View style={styles.container}>
                <View style={styles.logo}>
                    <Text h2 style={styles.logo}>Family</Text>
                </View>

                <View style={styles.form}>
                    <FormLabel>Email</FormLabel>
                    <FormInput
                        textInputRef={'email'}
                        autoCapitalize='none'
                        autoCorrect={false}
                        keyboardType={'email-address'}
                        autoFocus={true}
                        keyboardAppearance="light"
                        returnKeyType="next"
                        onChangeText={(email) => this.setState({email})}
                        onSubmitEditing={() => this.passwordFormInput.focus()}
                    />

                    <FormLabel>Password</FormLabel>
                    <FormInput
                        ref={ref => this.passwordFormInput = ref}
                        secureTextEntry={true}
                        returnKeyType="go"
                        onChangeText={(password) => this.setState({password})}
                        onSubmitEditing={() => this._signInAsync()}
                    />

                    { isLoading &&
                        <ActivityIndicator
                            style={styles.loadingIndicator}
                        />
                    }

                    <Button
                        style={styles.submitButton}
                        title="Sign In"
                        onPress={this._signInAsync}
                    />
                    <Button
                        style={styles.submitButton}
                        title="Sign Up"
                        onPress={this._signUp}
                    />
                </View>
            </View>
        );
    }

    _signInAsync = async () => {
        if (!this.state.email) {
            alert('Invalid Email');
            return;
        }
        if (!this.state.password) {
            alert('Invalid Password');
            return;
        }

        try {
            this.setState({
                isLoading: true,
            });
            await signInAsync(this.state.email, this.state.password);
            this.setState({
                isLoading: false,
            });

            this.props.navigation.navigate('App');
        } catch (err) {
            this.setState({
                isLoading: false,
            });

            if (err && err.response && err.response.status === 401) {
                alert('Unauthorized.\nEmail or Password is wrong.');
                return;
            }
            if (err && err.response && err.response.status) {
                alert(err.response.status);
                return;
            }
            if (err && err.response) {
                alert(err.response);
                return;
            }
            alert(err);
        }
    };

    _signUp = () => {
        this.props.navigation.navigate('SignUp');
    };
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
    loadingIndicator: {
        marginTop: 15,
    },
    submitButton: {
        marginTop: 20,
    }
});
