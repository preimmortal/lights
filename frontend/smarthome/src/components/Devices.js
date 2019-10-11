import React, {Component} from 'react';

import '../slate/bootstrap.min.css';


class Devices extends Component {
    constructor(props) {
        super(props);
        this.state = {
            devices: [],
            // devices {
            //     name  : ""
            //     ip    : ""
            //     alias : ""
            //     state : ""
            // }
        }
    }
    componentDidMount() {
        fetch('http://localhost:8081/devices')
        .then(res => res.json())
        .then((data) => {
            this.setState({ devices: data })
        })
        .catch(console.log)
    }

    sendToggle(ip, state) {
        console.log(`POST http://localhost:8081/devices/${ip}`)
        console.log({ip})
        console.log(`Toggle: {state}`)
        console.log({state})
        const newState = state === "off" ? "on" : "off"
        console.log({newState})
        fetch(`http://localhost:8081/devices/${ip}`, {
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
        })
        .catch(console.log)
    }

    render() {
        const {devices} = this.state;
        return (
            <div>
                <h2>Found the following list of devices:</h2>
                <table className="table">
                    <thead>
                        <tr>
                            <th scope="col">#</th>
                            <th scope="col">Device Name</th>
                            <th scope="col">Device Type</th>
                            <th scope="col">IP Address</th>
                            <th scope="col">State</th>
                        </tr>
                    </thead>
                    <tbody>
                            {devices.map((device, index) =>
                                <tr  key={index+1}>
                                    <th scope="row"><a href={'/devices/' + device.ip}>{index+1}</a></th>
                                    <td>
                                        <a href={'/devices/' + device.ip}>{device.alias}</a>
                                    </td>
                                    <td>
                                        <a href={'/devices/' + device.ip}>{device.name}</a>
                                    </td>
                                    <td>
                                        <a href={'/devices/' + device.ip}>{device.ip}</a>
                                    </td>
                                    <td>
                                        <button className="btn btn-primary btn-sm" onClick={this.sendToggle.bind(this, device.ip, device.state)}>{device.state}</button>
                                    </td>
                                </tr>
                            )}
                    </tbody>
                </table>
            </div>
        )
    }

}

export default Devices;