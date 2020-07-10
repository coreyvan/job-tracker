import React from 'react';
import './App.css';
import Company from './components/Company';
import NavBar from './components/NavBar';

const companyEndpoint = "http://localhost:3000/company"
// Copy of output from companyTwoEnpoint, while we work out some CORS issues.
const testData = [
  {"uid":"0x2","created_at":"0001-01-01T00:00:00Z","Company.name":"Corndog Inc","Company.description":"This is fallback data","Company.website":"corndog.com","Company.industries":["lifestyles","corndogs"],"Company.months":12,"Company.location":"Nashville, TN","Company.remote_possible":true},
  {"uid":"0x3","created_at":"0001-01-01T00:00:00Z","Company.name":"Moonbase Industries","Company.description":"Galactic voyages","Company.website":"moonbase.com","Company.industries":["lifestyles","space"],"Company.months":48,"Company.location":"Seattle, WA","Company.remote_possible":false},
  {"uid":"0x4","created_at":"0001-01-01T00:00:00Z","Company.name":"Meteor Street Studio","Company.description":"Branding design for inspire businesses","Company.website":"meteorstreetstudios.com","Company.industries":["design","entrepreneurs","business"],"Company.months":16,"Company.location":"Londom, United Kingdom","Company.remote_possible":true}
]

class App extends React.Component {
  constructor() {
    super();
    this.state = {
      data: []
    }
  }

  componentDidMount() {
    fetch(companyEndpoint,{
      method: 'get',
      headers: {
        'ContentType': 'application/json'
      }
    })
    .then(resp=>resp.json())
    .then(data=> {
      // This should work but...I haven't tested it.
      this.setState({data: [data]})
    }).catch(error=>{
      console.log(error)
      this.setState({data: testData})
    })
  }

  render() {
    var data = this.state.data
    return (
      <div className="App">
        <NavBar />
        {data.map((company)=>
          <Company 
            key={company.uid}
            data={company} 
            />
        )}
      </div>
    );
  }
}

export default App;
