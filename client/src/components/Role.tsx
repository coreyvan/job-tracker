import React, { Component } from 'react';
import {
  Badge, 
  Button, 
  Card, 
  CardText,
  CardTitle, 
  CardSubtitle
} from 'reactstrap';
import {IRole} from '../utils/types';

class Role extends Component<{role: IRole;}> {
    render() {
        let role = this.props.role
        let availableText = ""
        if (role.remote_possible){
            availableText = "Remote"
        }
        return (
            <Card className="Role">
                <CardTitle>{role.title}</CardTitle>
                <CardSubtitle>{role.location}&nbsp;<Badge color="dark">{availableText}</Badge></CardSubtitle>
                <CardText>Posted on {role.posted_on}</CardText>
                <Button className={"RoleButton"} color="secondary">Apply</Button>
            </Card>
        )
    }
}

export default Role