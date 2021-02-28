import React from 'react';
import './App.css';
import LoginDialog from './components/LoginDialog';
import RegisterDialog from './components/RegisterDialog';

function App() {
  return (
    <div className="App">
      <LoginDialog/>
      <RegisterDialog/>
    </div>
  );
}

export default App;
