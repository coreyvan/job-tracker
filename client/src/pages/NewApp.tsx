import React from 'react';
import Company from '../components/Company';
import NewCompanyForm from '../components/NewCompanyForm'
import {ICompany} from '../utils/types';


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
        //This doesn't have an ID yet!  Isn't on the server!
        //Def should be making a request before displaying this.
        console.log(company)
        this.setState({company:company})
    }

    render(){
        let companies = this.props.companies
        let company = this.state.company
        return (
            <div className="Page">
                <h1>Lets get you into a new application</h1>
                {company == null 
                ?<div>
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
                <Company 
                    company={company} //This might end up taking us to a new container, just a stub for now.
                    />
                        }
            </div>
        )
    }
}
export default NewApp