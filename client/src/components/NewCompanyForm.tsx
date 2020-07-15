import React from 'react';
import {
    Button,
    Form,
    FormGroup,
    Label,
    Input
} from 'reactstrap';
import {ICompany} from '../utils/types'

interface INewCompanyFormProps {
    onSubmit(ICompany): void;
}

class NewCompanyForm extends React.Component <INewCompanyFormProps, {}> { 
    submitForm = (event) => {
        //TODO: Form validation, tests
        //In validation, we can probably store these values into state, and not be ugly upon submission.
        event.preventDefault()
        const newCompany = {
            name: event.target[0].value,
            description: event.target[1].value,
            website: event.target[2].value,
            industries: event.target[3].value,
            months: event.target[4].value,
            location: event.target[5].value,
            remote_possible: event.target[6].value
        } as ICompany
        this.props.onSubmit(newCompany)
    }
    render(){
        return (
            <Form className="NewCompanyForm" onSubmit={this.submitForm}>
                <FormGroup>
                    <Label for="name">Company Name</Label>
                    <Input type="text" name="name" id="companyName" />
                </FormGroup>
                <FormGroup>
                    <Label for="description">Description</Label>
                    <Input type="textarea" name="description" id="companyDescription" />
                </FormGroup>
                <FormGroup>
                    <Label for="website">Website</Label>
                    <Input type="url" name="website" id="companyWebsite" />
                </FormGroup>
                <FormGroup>
                    <Label for="industry">Industry</Label>
                    <Input type="text" name="industry" id="companyIndustry" placeholder="Just one for now thanks"/>
                </FormGroup>
                <FormGroup>
                    <Label for="months">Months</Label>
                    <Input type="number" name="months" id="companyMonths"/>
                </FormGroup>
                <FormGroup>
                    <Label for="location">Location</Label>
                    <Input type="text" name="location" id="companyLocation"/>
                </FormGroup>
                <FormGroup check>
                    <Label check>
                    <Input type="checkbox" />{' '}
                    Remote
                    </Label>
                </FormGroup>
                <Button submit="true" onSubmit={this.submitForm}>Submit</Button>
            </Form>
        )
    }
}

export default NewCompanyForm;