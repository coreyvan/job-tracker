import React, { Component } from 'react';
import {
  Navbar,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  NavbarText
} from 'reactstrap';

class NavBar extends Component {
    render() {
        return (
            <Navbar color="info" light expand="md">
                <NavbarBrand href="/">Trackr</NavbarBrand>
                <Nav className="mr-auto" navbar>
                    <NavItem>
                        <NavLink>New Application</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink>In Progress</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink>Past Applications</NavLink>
                    </NavItem>
                </Nav>
                <NavbarText>Hey look ma we made it</NavbarText>
            </Navbar>
        )
    }
}

export default NavBar;