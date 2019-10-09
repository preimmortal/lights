import React, {Component} from 'react';
import Navbar from './Navbar';
import './slate/bootstrap.min.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <div>
          <Navbar />
        </div>
        <header className="App-header">
          <p>
            Hello
          </p>
        </header>
      </div>
    );
  }
}

export default App;
