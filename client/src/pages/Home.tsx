import React from 'react';
import Company from '../components/Company';
import Role from '../components/Role';
import {ICompany, IRole} from '../utils/types';

interface IHomeProps {
    companies: Array<ICompany>
    roles: Array<IRole>
}

class Home extends React.Component <IHomeProps, {}>{

    render(){
        let companies = this.props.companies
        let roles = this.props.roles
        return (
            <div className="Page">
                <div className="WidgetHolder">
                    {companies.map((company)=>
                        <Company 
                            key={company.id}
                            company={company} 
                            />
                    )}
                </div>
                <div className="WidgetHolder">
                    {roles.map((role)=>
                    <Role 
                        key={role.id}
                        role={role} 
                        />
                    )}
                </div>
            </div>
        )
    }
}
export default Home