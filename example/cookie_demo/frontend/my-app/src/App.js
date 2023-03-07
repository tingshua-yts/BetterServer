import logo from './logo.svg';
import './App.css';
import Cookies from 'js-cookie';
import React from 'react';

async function init() {
  console.log("init")
  const response = await fetch('/api/v1/pipelineCheckInstall')
  const response_json = await response.text()
  console.log(response_json)

}
init()
// function sleepFor(sleepDuration){
//   var now = new Date().getTime();
//   while(new Date().getTime() < now + sleepDuration){ 
//       /* Do nothing */ 
//   }
// }

var count = 1
class Clock extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      date: new Date(),
      select: false,
      dump: "cccc"};
  }

  componentDidMount() {
    fetch('http://www.baidu.com', {mode: 'no-cors'}).then((result)=>{
      this.setState({select: SVGComponentTransferFunctionElement})
      count = count + 1
      console.log(this.state.dump)
    })
  }

  componentWillUnmount() {
    
  }


  render() {
    return (
      <div>
        <h1>Hello, world!</h1>
        <h2>It is {this.state.date.toLocaleTimeString()}.</h2>
        <h2>It is {count}</h2>
      </div>
    );
  }
}
function handler() {
  console.log("button hello")   
}
function App() {
  return (
    <div>
    <button onClick={handler}>
    Activate Lasers
    </button>
    <Clock/>
    </div>

  );
}

export default App;
