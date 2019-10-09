import React, {Component} from 'react';

import '../slate/bootstrap.min.css';


class Devices extends Component {
    componentDidMount() {
        fetch('192.168.1.203:8081/devices')
        .then(res => res.json())
        .then((data) => {
            this.setState({ contacts: data })
        })
        .catch(console.log)
    }
    render() {
        return (
            <div>
              <p1>Hello from devices</p1>
            </div>
        )
    }

}

export default Devices;