import logo from './logo.svg';
import './App.css';
import Cookies from 'js-cookie';

async function init() {
  console.log("init")
  const response = await fetch('/incr');
  const json_rep = await response.json();
  console.log(json_rep.count)
  console.log(json_rep.count)
}
init()

function handler() {
  console.log("hello")
  // fetch('/incr', ).then(response=>response.text()) // 只有这样写，才能在chrome的开发者工具中看到返回值
  // console.log(Cookies.get('mysession')) // 对于gin， key为mysession， value为gin中session序列化的值
  // console.log(Cookies)
   //init()
}
function App() {
  return (
    <button onClick={handler}>
    Activate Lasers
  </button>
  );
}

export default App;
