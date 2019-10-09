import React, {Component} from 'react';
import { BrowserRouter, Route, Switch} from 'react-router-dom'

import './slate/bootstrap.min.css';

import Home from './components/Home';
import Devices from './components/Devices';
import Error from './components/Error'
import Navbar from './components/Navbar';

class App extends Component {
  render() {
    return (
      <BrowserRouter>
        <div>
          <Navbar />
          <Switch>
            <Route path="/" component={Home} exact />
            <Route path="/devices" component={Devices} />
            <Route component={Error} />
          </Switch>
        </div>
      </BrowserRouter>
    );
  }
}

export default App;
