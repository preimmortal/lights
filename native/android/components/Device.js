import React, { Component } from 'react';
import { 
    StyleSheet, 
    Text, 
    View, 
    FlatList, 
    ActivityIndicator 
} from 'react-native';

import DeviceState from './DeviceState'


export default class Devices extends Component {
    constructor(props) {
        super(props)
        this.state = { 
            isLoading: true,
            devices: [],
        }
    }

    componentDidMount() {
        const fetchUrl = "http://192.168.1.203:8081/devices"
        fetch(fetchUrl)
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
        function Item({data}) {
            return (
                <View style={styles.item}>
                    <Text style={styles.title}>{data.alias}</Text>
                    <DeviceState ip={data.ip} />
                </View>
            )
        }
        console.log(this.state.devices)

        if(this.state.isLoading){
            return (
              <View style={{flex: 1, padding: 20}}>
                <ActivityIndicator/>
              </View>
            )
        }
        return (
            <View style={styles.container}>
                <FlatList
                    data={this.state.devices}
                    renderItem={({item}) => <Item data={item}></Item>}
                    keyExtractor={({ip}, index) => ip}
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
    header: {
        flex: 1,
        color: 'white',
        fontWeight: 'bold',
        fontSize: 30,
    },
    title: {
        flex: 1,
        color: 'white',
        fontWeight: 'bold',
        fontSize: 30,
    },
});
