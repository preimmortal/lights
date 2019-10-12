import React, {Component} from 'react';

import '../slate/bootstrap.min.css';


class DeviceInfo extends Component {
    constructor(props) {
        super(props);
        this.state = {
            deviceinfo: null,
            // devices {
            //     name  : ""
            //     ip    : ""
            //     alias : ""
            //     state : ""
            // }
        }
    }
    componentDidMount() {
        console.log("GET" , `http://localhost:8081/devices/${this.props.ip}`)
        fetch(`http://localhost:8081/devices/${this.props.ip}`)
        .then(res => res.json())
        .then((data) => {
            this.setState({ deviceinfo: data })
        })
        .catch(console.log)
    }

    sendToggle(ip, state) {
        console.log(`POST http://localhost:8081/devices/${this.props.ip}`)
        const newState = state === "off" ? "on" : "off"
        fetch(`http://localhost:8081/devices/${this.props.ip}`, {
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
            console.log("GET" , `http://localhost:8081/devices/${this.props.ip}`)
            fetch(`http://localhost:8081/devices/${this.props.ip}`)
            .then(res => res.json())
            .then((data) => {
                this.setState({ deviceinfo: data })
            })
            .catch(console.log)
        })
        .catch(console.log)

    }

    render() {
        const {deviceinfo} = this.state;
        const { ip } = this.props.ip
        if (deviceinfo) {
            const state = deviceinfo.system.get_sysinfo.relay_state === 1 ? "on" : "off"
            return (
                <div>
                    <button className="btn btn-primary btn-sm" onClick={this.sendToggle.bind(this, ip, state)}>{state}</button>
                </div>
            )
        }
        return (
            <div>
                <button className="btn btn-primary btn-sm" onClick={this.sendToggle.bind(this, ip, "N/A")}>N/A</button>
            </div>
        )
    }
}

export default DeviceInfo;