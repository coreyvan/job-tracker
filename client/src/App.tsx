import React from 'react';
import './App.css';
import NavBar from './components/NavBar';
import Home from './pages/Home'
import NewApp from './pages/NewApp'
import {ICompany, IRole} from './utils/types';

const companiesEndpoint = "http://localhost:3000/companies?limit=5"
const rolesEndpoint = "http://localhost:3000/roles?limit=5"

interface IAppProps {
}

interface IAppState {
  companies: Array<ICompany>
  roles: Array<IRole>
  route: string
}

class App extends React.Component <IAppProps, IAppState>{
  constructor(props: IAppProps) {
    super(props);
    this.state = {
      companies: [],
      roles: [],
      route: "Home"
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

  onRouteChanged = (route: string) => {
    console.log(route)
    this.setState({route: route});
  }

  render() {
    var companies = this.state.companies
    var roles = this.state.roles
    let route = this.state.route
    return (
      <div className="App">
        <NavBar routeChange={this.onRouteChanged}/>
        { route === 'Home'
          ? <Home
              companies={companies}
              roles={roles}/>
          : <NewApp
              companies={companies}/>}
      </div>
    );
  }
}

export default App;
