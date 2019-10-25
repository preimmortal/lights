import React from 'react';
import { StyleSheet, Text, View, Button } from 'react-native';
import Devices from './components/Device'

export default function App() {
  return (
    <View style={styles.container}>
        <Text style={styles.text}>Smarthome</Text>
        <Devices />
    </View>
  );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: 'column',
        backgroundColor: '#1F1B1B',
        alignItems: 'center',
        justifyContent: 'center',
    },
    text: {
        flex: 1,
        color: 'white',
        fontWeight: 'bold',
        fontSize: 30,
    },
});
