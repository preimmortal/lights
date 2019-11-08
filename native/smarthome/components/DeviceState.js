import React, { Component } from 'react';


import { 
    StyleSheet, 
    Text, 
    View, 
    Button, 
    SafeAreaView,
    FlatList, 
    ActivityIndicator,
    Alert
} from 'react-native';


export default class DeviceState extends Component {
    constructor(props) {
        super(props)
        this.state = { 
            isLoading: true,
            deviceinfo: null,
        }
    }

    componentDidMount() {
        const fetchUrl = "http://192.168.1.203:8081/devices/" + this.props.ip
        fetch(fetchUrl)
        .then(res => res.json())
        .then((data) => {
            this.setState({
                deviceinfo: data,
                isLoading: false,
            }, function(){
            });
        })
        .catch(console.log)
    }

    sendToggle(ip, state) {
        const postUrl = "http://192.168.1.203:8081/devices/" + this.props.ip
        console.log("POST", {postUrl})
        const newState = state === "off" ? "on" : "off"
        fetch(postUrl, {
            crossDomain: true,
            method: 'POST',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                state: `${newState}`,
            })
        })
        .then(res => res.json())
        .then((data) => {
            console.log(data)
            // Send another request to update info
            const fetchUrl = "http://192.168.1.203:8081/devices/" + this.props.ip
            console.log("GET" , fetchUrl)
            fetch(fetchUrl)
            .then(res => res.json())
            .then((data) => {
                this.setState({ deviceinfo: data })
            })
            .catch(console.log)
        })
        .catch(console.log)
    }

    render() {
        console.log(this.state.deviceinfo)
        const {deviceinfo} = this.state
        const {ip} = this.props.ip

        if (deviceinfo) {
            const state = deviceinfo.system.get_sysinfo.relay_state === 1 ? "on" : "off"
            return (
                <View style={styles.container}>
                    <Button 
                        title={state}
                        color="#14347F"
                        onPress={() => this.sendToggle(ip, state)}
                    />
                </View>
            );
        } else {
            return (
              <View style={{flex: 1, padding: 20}}>
                <ActivityIndicator/>
              </View>
            )
        }
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

