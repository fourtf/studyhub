import { AppBar, Button, Toolbar, Typography } from '@material-ui/core';
import LoginDialog from 'components/LoginDialog';
import RegisterDialog from 'components/RegisterDialog';
import React from 'react';
import { Link } from 'react-router-dom';

export default function Nav() {
  return (
    <div>
      <AppBar>
        <Toolbar>
          <Link to="/">
            <Button>
              <Typography variant="h6">StudyHub</Typography>
            </Button>
          </Link>
          <LoginDialog />
          <RegisterDialog />
        </Toolbar>
      </AppBar>
      <div style={{ height: 64 }}></div>
    </div>
  );
}
