import React from 'react';
import './App.css';
import {
  Card, CardText,
  CardTitle, CardSubtitle
} from 'reactstrap';

const companyTwoEnpoint = "http://localhost:3000/company/2"
// Copy of output from companyTwoEnpoint, while we work out some CORS issues.
const testData = {"uid":"0x2","created_at":"0001-01-01T00:00:00Z","Company.name":"Fallback data","Company.description":"Makes corndogs","Company.website":"corndog.com","Company.industries":["lifestyles","corndogs"],"Company.months":12,"Company.location":"Nashville, TN","Company.remote_possible":true}

class App extends React.Component {
  constructor() {
    super();
    this.state = {}
  }

  componentDidMount() {
    fetch(companyTwoEnpoint,{
      method: 'get',
      headers: {
        'ContentType': 'application/json'
      }
    })
    .then(resp=>resp.json())
    .then(data=> {
      // This should work but...I haven't tested it.
      this.setState({data: data})
    }).catch(error=>{
      console.log(error)
      this.setState({data: testData})
    })
  }

  render() {
    var data = this.state.data
    return (
      <div className="App">
        <header className="App-header">
          Hey look ma we made it.
        </header>
        {data != null &&
          <Card className="Company">
            <CardTitle>{data["Company.name"]}</CardTitle>
            <CardSubtitle>{data["Company.location"]}</CardSubtitle>
            <CardText>{data["Company.description"]}</CardText>
          </Card>
        }
      </div>
    );
  }
}

export default App;
