import React from 'react';
import { StyleSheet, Text, View } from 'react-native';

var Api = require('api');

Api.ApiClient.instance.basePath = "http://localhost:3000";

var api = new Api.AccountsServiceApi();

var body = new Api.AccountsLogInRequest(); // {AccountsLogInRequest}


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
api.logIn(body, callback);

export default class App extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <Text>Open up App.js to start working on your app!</Text>
        <Text>Changes you make will automatically reload.</Text>
        <Text>Shake your phone to open the developer menu.</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
