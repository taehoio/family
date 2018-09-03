import React from 'react';
import { StyleSheet, Text, View } from 'react-native';

import Api, { logIn, isLoggedIn } from './api';
logIn('taeho@taeho.io', '1234')
    .then(console.log)
    .catch(console.log);

console.log(isLoggedIn());
setTimeout(() => {
    console.log(isLoggedIn());

    const api = new Api.TodoGroupsServiceApi();
    const body = Api.TodogroupsCreateTodoGroupRequest.constructFromObject({
        account_id: 'be35al3oo3solaig0nb0',
        todo_group: Api.TodogroupsTodoGroup.constructFromObject({
            title: 'test_title',
        }),
    });

    api.createTodoGroup(body)


}, 5000);

export default class App extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <Text>Open up App.js to start working on your app!</Text>
        <Text>Yamma! Yo!</Text>
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
