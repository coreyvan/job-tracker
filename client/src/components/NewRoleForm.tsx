import React from 'react';
import {
    Button,
    Form,
    FormGroup,
    Label,
    Input
} from 'reactstrap';
import {ICompany, IRole} from '../utils/types'

interface INewRoleFormProps {
    company: ICompany,
    onRoleSubmit(IRole): void;
}

class NewRoleForm extends React.Component <INewRoleFormProps, {}> { 
    submitForm = (event) => {
        //TODO: Form validation, tests
        //In validation, we can probably store these values into state, and not be ugly upon submission.
        event.preventDefault()
        const newRole = {
            title: event.target[0].value,
            company: {
                id: this.props.company.id,
                name: this.props.company.name
            },
            url: event.target[1].value,
            technologies: [event.target[2].value],
            pay_lower: parseInt(event.target[3].value),
            pay_upper: parseInt(event.target[4].value),
            location: event.target[5].value,
            level: event.target[6].value,
            posted_on: (new Date()).toISOString(),
            remote_possible: event.target[7].checked
        } as IRole
        this.props.onRoleSubmit(newRole)
    }
    render(){
        return (
            <Form className="NewRoleForm" onSubmit={this.submitForm}>
                <FormGroup>
                    <Label for="title">Title</Label>
                    <Input type="text" name="title" id="roleName" />
                </FormGroup>
                <FormGroup>
                    <Label for="url">URL</Label>
                    <Input type="url" name="url" id="roleUrl" />
                </FormGroup>
                <FormGroup>
                    <Label for="technologies">Tech</Label>
                    <Input type="textarea" name="technologies" id="roleTech" />
                </FormGroup>
                <FormGroup>
                    <Label for="payLower">Minimum Estimated Salary</Label>
                    <Input type="number" name="payLower" id="rolePayLower"/>
                </FormGroup>
                <FormGroup>
                    <Label for="payUpper">Maximum Estimated Salary</Label>
                    <Input type="number" name="payUpper" id="rolePayLower"/>
                </FormGroup>
                <FormGroup>
                    <Label for="location">Location</Label>
                    <Input type="text" name="location" id="roleLocation"/>
                </FormGroup>
                <FormGroup>
                    <Label for="level">Level</Label>
                    <Input type="text" name="level" id="roleLevel"/>
                </FormGroup>
                {/* <FormGroup>
                    <Label for="posted">Posted On</Label>
                    <Input type="number" name="posted" id="rolePosted"/>
                </FormGroup> */}
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

export default NewRoleForm;
