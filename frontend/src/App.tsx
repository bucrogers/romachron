import React from 'react';
import logo from './logo.svg';
import './App.css';

import RomanTimestamps from "./components/RomanTimestamps"

function App() {

  React.useEffect(() => {
    const interval = setInterval(() => {
      //window.location.reload()
    }, 500);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>Romachron</h1>
      </header>
      <RomanTimestamps/>
    </div>
  );
}

export default App;
