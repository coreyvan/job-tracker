import React, { Component } from 'react';
import {
  Badge, 
  Button, 
  Card, 
  CardText,
  CardTitle, 
  CardSubtitle
} from 'reactstrap';

class Company extends Component {

    render() {
        const company = this.props.data
        let availableText = ""
        if (company["Company.remote_possible"]){
            availableText = "Remote"
        }
        return (
            <Card className="Company">
                <CardTitle>{company["Company.name"]}</CardTitle>
                <CardSubtitle>{company["Company.location"]}&nbsp;<Badge color="dark">{availableText}</Badge></CardSubtitle>
                <CardText>{company["Company.description"]}</CardText>
                <Button className={"CompanyButton"} color="primary">Apply</Button>
            </Card>
        )
    }
}

export default Company