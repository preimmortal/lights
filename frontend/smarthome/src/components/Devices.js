import React, {Component} from 'react';

import '../slate/bootstrap.min.css';


class Devices extends Component {
    constructor(props) {
        super(props);
        this.state = {
            devices: [],
        }
    }
    componentDidMount() {
        fetch('http://192.168.1.203:8081/devices')
        .then(res => res.json())
        .then((data) => {
            this.setState({ devices: data })
        })
        .catch(console.log)
    }
    render() {
        const {devices} = this.state;
        return (
            <div>
                <h2>Found the following list of devices:</h2>
                <table class="table">
                    <thead>
                        <th scope="col">#</th>
                        <th scope="col">Device Name</th>
                        <th scope="col">IP Address</th>
                        <th scope="col">Status</th>
                    </thead>
                    <tbody>
                            {devices.map((device, index) =>
                                <tr>
                                    <th scope="row"><a href={'/devices/' + device.ip}>{index+1}</a></th>
                                    <td key={index+1}>
                                        <a href={'/devices/' + device.ip}>{device.name}</a>
                                    </td>
                                    <td>
                                        <a href={'/devices/' + device.ip}>{device.ip}</a>
                                    </td>
                                    <td>
                                        <a href={'/devices/' + device.ip}>OFF</a>
                                    </td>
                                </tr>
                            )}
                    </tbody>
                <ul>
                </ul>
                </table>
            </div>
        )
    }

}

export default Devices;