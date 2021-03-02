import { AppBar, Toolbar, Typography } from '@material-ui/core';
import React from 'react';
import './App.css';
import LoginDialog from './components/LoginDialog';
import RegisterDialog from './components/RegisterDialog';

function App() {
  return (
    <div className="App">
      <AppBar>
        <Toolbar>
          <Typography variant="h6">StudyHub</Typography>
          <LoginDialog />
          <RegisterDialog />
        </Toolbar>
      </AppBar>
    </div>
  );
}

export default App;
