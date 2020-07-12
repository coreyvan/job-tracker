import React, { Component } from 'react';
import {
  Badge, 
  Button, 
  Card, 
  CardText,
  CardTitle, 
  CardSubtitle
} from 'reactstrap';
import {ICompany} from '../utils/types';

class Company extends Component<{company: ICompany;}> {
    render() {
        let company = this.props.company
        let availableText = ""
        if (company.remote_possible){
            availableText = "Remote"
        }
        return (
            <Card className="Company">
                <CardTitle>{company.name}</CardTitle>
                <CardSubtitle>{company.location}&nbsp;<Badge color="dark">{availableText}</Badge></CardSubtitle>
                <CardText>{company.description}</CardText>
                <Button className={"CompanyButton"} color="primary">Apply</Button>
            </Card>
        )
    }
}

export default Company