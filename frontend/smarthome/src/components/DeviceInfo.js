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
        const { ip } = this.props.match.params
        console.log(`http://localhost:8081/devices/${ip}`)
        fetch(`http://localhost:8081/devices/${ip}`)
        .then(res => res.json())
        .then((data) => {
            console.log(data)
            this.setState({ deviceinfo: data })
        })
        .catch(console.log)
    }
    render() {
        const {deviceinfo} = this.state;
        console.log(this.state)
        return (
            <div>
                <h2>Device Info for: {this.props.match.params.ip}</h2>
                <div><pre>{JSON.stringify(deviceinfo, null, 2) }</pre></div>
            </div>
        )
    }

}

export default DeviceInfo;