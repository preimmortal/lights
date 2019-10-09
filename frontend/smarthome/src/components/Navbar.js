import React, {Component} from 'react';
import { NavLink } from "react-router-dom"
import '../slate/bootstrap.min.css';

class Navbar extends Component {
    render() {
      return (
          <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
            <a class="navbar-brand" href="/">Smarthome</a>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
              <span class="navbar-toggler-icon"></span>
            </button>
          
            <div class="collapse navbar-collapse" id="navbarColor01">
              <ul class="navbar-nav mr-auto">
                <li class="nav-item active">
                  <NavLink to="/" class="nav-link">Home <span class="sr-only">(current)</span></NavLink>
                </li>
                <li class="nav-item">
                  <NavLink to="/devices" class="nav-link">Devices</NavLink>
                </li>
                <li class="nav-item">
                  <NavLink to="/about" class="nav-link">About</NavLink>
                </li>
              </ul>
              <form class="form-inline my-2 my-lg-0">
                <input class="form-control mr-sm-2" type="text" placeholder="Search"></input>
                <button class="btn btn-secondary my-2 my-sm-0" type="submit">Search</button>
              </form>
            </div>
          </nav>
      )
    }
}

export default Navbar;
