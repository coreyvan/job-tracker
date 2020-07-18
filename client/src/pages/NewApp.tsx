import React from 'react';
import Company from '../components/Company';
import NewCompanyForm from '../components/NewCompanyForm';
import NewRoleForm from '../components/NewRoleForm';
import {ICompany, IRole} from '../utils/types';

const newCompanyEndpoint = "http://localhost:3000/company";
const newRoleEndpoint = "http://localhost:3000/role";

interface INewAppProps {
    companies: Array<ICompany>
}

interface INewAppState {
    company: ICompany | null
}

class NewApp extends React.Component <INewAppProps, INewAppState>{

    constructor(props: INewAppProps) {
      super(props);
      this.state = {
        company: null
      }
    }

    onCompanySelected = (company: ICompany) => {
        this.setState({company:company})
    }

    onNewCompanySubmitted = (company: ICompany) =>{
        //This should probably refresh the companies data in home.
        let co
        for(co in this.props.companies){
            if(company.name === co.name){
                // Eventually, do this checking in form validation.
                console.log("Company name already exists!")
                return
            }
        }
        fetch(newCompanyEndpoint,{
            method: 'post',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(company)
        })
        .then(resp=>resp.json())
        .then(data=> {
            this.setState({company:data})
        }).catch(error=>{
            console.log("Request to create a new Company failed")
            console.log(error)
        })
    }

    onNewRoleSubmitted = (role: IRole) =>{
        //This should probably refresh the roles data in home.
        //Or just append the result to our larger data?
        fetch(newRoleEndpoint,{
            method: 'post',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(role)
        })
        .then(resp=>resp.json())
        .then(data=> {
            //Not sure where the workflow goes from here?  Homepage?
            console.log(data)
        }).catch(error=>{
            console.log("Request to create a new Role failed")
            console.log(error)
        })
    }
    render(){
        let companies = this.props.companies
        let company = this.state.company
        return (
            <div className="Page">
                <h1>Lets get you into a new application</h1>
                {company == null 
                ?
                    <div>
                        <h2>First select an existing company or fill out a new one</h2>
                        <div className="WidgetHolder">
                            {companies.map((company)=>
                                <Company 
                                    key={company.id}
                                    company={company} 
                                    onSelect={this.onCompanySelected}
                                    />
                            )}
                        </div>
                        <h3 className="NewAppCompanyDivider">Didn't find what you were looking for up there?</h3>
                        <NewCompanyForm onSubmit={this.onNewCompanySubmitted}/>
                    </div>
                :
                    <div>
                        <Company company={company}/> //This might end up taking us to a new container, just a stub for now.
                        <NewRoleForm onRoleSubmit={this.onNewRoleSubmitted} company={company}/>
                    </div>
                }
            </div>
        )
    }
}
export default NewApp