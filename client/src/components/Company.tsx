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

class Company extends Component<{company: ICompany, onSelect?:Function}> {
    render() {
        let company = this.props.company
        let availableText = ""
        if (company.remote_possible){
            availableText = "Remote"
        }
        let onClickFunc = () =>{
            this.props.onSelect?.(company)
        }
        return (
            <Card className="Company">
                <CardTitle>{company.name}</CardTitle>
                <CardSubtitle>{company.location}&nbsp;<Badge color="dark">{availableText}</Badge></CardSubtitle>
                <CardText>{company.description}</CardText>
                <Button className={"CompanyButton"} color="primary" onClick={onClickFunc}>Apply</Button>
            </Card>
        )
    }
}

export default Company