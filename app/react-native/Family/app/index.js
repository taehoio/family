import { createStackNavigator, createSwitchNavigator } from 'react-navigation';

import AuthLoadingScreen from './screens/AuthLoadingScreen';
import SignInScreen from './screens/SignInScreen';
import SignUpScreen from './screens/SignUpScreen';
import HomeScreen from './screens/HomeScreen';
import OtherScreen from './screens/OtherScreen';
import NoteListScreen from './screens/NoteListScreen';


const AppStack = createStackNavigator({
    Home: HomeScreen,
    Other: OtherScreen,
    NoteList: NoteListScreen,
});

const AuthStack = createStackNavigator({
    SignIn: SignInScreen,
    SignUp: SignUpScreen,
});

export default createSwitchNavigator({
    AuthLoading: AuthLoadingScreen,
    App: AppStack,
    Auth: AuthStack,
}, {
    initialRouteName: 'AuthLoading',
});
