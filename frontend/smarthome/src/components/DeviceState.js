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
        const fetchUrl = process.env.REACT_APP_BACKEND_URL + "/devices/" + this.props.ip
        console.log("GET" , fetchUrl)
        fetch(fetchUrl)
        .then(res => res.json())
        .then((data) => {
            this.setState({ deviceinfo: data })
        })
        .catch(console.log)
    }

    sendToggle(ip, state) {
        const postUrl = process.env.REACT_APP_BACKEND_URL + "/devices/" + this.props.ip
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
            const fetchUrl = process.env.REACT_APP_BACKEND_URL + "/devices/" + this.props.ip
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
        const {deviceinfo} = this.state;
        const {ip} = this.props.ip
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