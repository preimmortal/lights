import React, { Component } from 'react';
import { StyleSheet, Text, View, Button, FlatList, ActivityIndicator } from 'react-native';


export default class Devices extends Component {
    constructor(props) {
        super(props)
        this.state = { 
            isLoading: true,
            devices: [],
        }
    }

    componentDidMount() {
        fetch("http://192.168.1.203:8081/devices")
        .then(res => res.json())
        .then((data) => {
            this.setState({
                devices: data,
                isLoading: false,
            }, function(){
            });
        })
        .catch(console.log)
    }

    render() {
        if(this.state.isLoading){
            return(
              <View style={{flex: 1, padding: 20}}>
                <ActivityIndicator/>
              </View>
            )
        }
        return (
            <View style={styles.container}>
                <FlatList
                    data={this.state.devices}
                    renderItem={({item}) => <Text style={styles.item}>{item.ip} - {item.alias} - {item.name}</Text>}
                    keyExtractor={({id}, index) => id}
                />
            </View>
        );
    }
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        backgroundColor: '#1F1B1B',
        alignItems: 'center',
        justifyContent: 'center',
    },
    item: {
        flex: 1,
        color: 'white',
        fontWeight: 'bold',
        fontSize: 30,
    },
});
