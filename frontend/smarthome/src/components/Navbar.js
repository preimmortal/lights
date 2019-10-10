import React, {Component} from 'react';
import { NavLink } from "react-router-dom"
import '../slate/bootstrap.min.css';

class Navbar extends Component {
    render() {
        return (
            <nav class="navbar navbar-expand-lg navbar-dark bg-primary">
                <a class="navbar-brand" href="/">Smarthome</a>
                <div class="collapse navbar-collapse" id="navbarColor01">
                    <ul class="navbar-nav mr-auto">
                        <li class="nav-item active">
                            <NavLink to="/" className="nav-link">Home <span class="sr-only">(current)</span></NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink to="/devices" className="nav-link">Devices</NavLink>
                        </li>
                        <li class="nav-item">
                            <NavLink to="/about" className="nav-link">About</NavLink>
                        </li>
                    </ul>
                </div>
            </nav>
        )
    }
}

export default Navbar;
