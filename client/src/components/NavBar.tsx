import React from 'react';
import {
  Navbar,
  NavbarBrand,
  Nav,
  NavItem,
  NavLink,
  NavbarText
} from 'reactstrap';

interface INavBarProps {
    routeChange(string): void;
}
const NavBar = ({ routeChange }: INavBarProps) => {
    return (
        <Navbar color="info" light expand="md">
            <NavbarBrand onClick={()=>routeChange("Home")}>
                    Trackr
            </NavbarBrand>
            <Nav className="mr-auto" navbar>
                <NavItem onClick={()=>routeChange("NewApp")}>
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

export default NavBar;