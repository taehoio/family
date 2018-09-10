import React from 'react';
import {
    View,
    StyleSheet,
} from 'react-native';
import {
    List,
    ListItem,
} from 'react-native-elements';


const list = [
    {
        name: 'Amy Farha',
        avatar_url: 'https://s3.amazonaws.com/uifaces/faces/twitter/ladylexy/128.jpg',
        subtitle: 'Vice President'
    }, {
        name: 'Chris Jackson',
        avatar_url: 'https://s3.amazonaws.com/uifaces/faces/twitter/adhamdannaway/128.jpg',
        subtitle: 'Vice Chairman'
    },
];


export default class NoteListScreen extends React.Component {
    static navigationOptions = {
        title: 'Notes',
    };

    render() {
        return (
            <View style={styles.container}>
                <List>
                    {
                        list.map((l) => (
                            <ListItem
                                roundAvatar
                                avatar={{uri: l.avatar_url}}
                                key={l.name}
                                title={l.name}
                            />
                        ))
                    }
                </List>
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
    },
    list: {
        marginBottom: 20,
    },
});
