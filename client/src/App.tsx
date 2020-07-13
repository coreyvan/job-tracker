import React from 'react';
import './App.css';
import Company from './components/Company';
import NavBar from './components/NavBar';
import Role from './components/Role';
import {ICompany, IRole} from './utils/types';

const companiesEndpoint = "http://localhost:3000/companies?limit=5"
const rolesEndpoint = "http://localhost:3000/roles?limit=5"

interface IAppProps {
}

interface IAppState {
  companies: Array<ICompany>
  roles: Array<IRole>
}

class App extends React.Component <IAppProps, IAppState>{
  constructor(props: IAppProps) {
    super(props);
    this.state = {
      companies: [],
      roles: []
    }
  }

  componentDidMount() {
    fetch(companiesEndpoint,{
      method: 'get',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(resp=>resp.json())
    .then(data=> {
      this.setState({companies: data})
    }).catch(error=>{
      console.log(error)
    })

    fetch(rolesEndpoint,{
      method: 'get',
      headers: {
        'Content-Type': 'application/json'
      }
    })
    .then(resp=>resp.json())
    .then(data=> {
      this.setState({roles: data})
    }).catch(error=>{
      console.log(error)
    })
  }

  render() {
    var companies = this.state.companies
    var roles = this.state.roles
    return (
      <div className="App">
        <NavBar />
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
    );
  }
}

export default App;
