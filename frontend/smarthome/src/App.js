import React, {Component} from 'react';
import { BrowserRouter, Route, Switch} from 'react-router-dom'

import './slate/bootstrap.min.css';

import Home from './components/Home';
import Devices from './components/Devices';
import DeviceInfo from './components/DeviceInfo';
import About from './components/About';
import Error from './components/Error'
import Navbar from './components/Navbar';

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Navbar />
          <Switch>
            <Route exact path="/" component={Home} />
            <Route exact path="/devices" component={Devices} />
            <Route exact path="/devices/:ip" component={DeviceInfo} />
            <Route exact path="/about" component={About} />
            <Route component={Error} />
          </Switch>
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
