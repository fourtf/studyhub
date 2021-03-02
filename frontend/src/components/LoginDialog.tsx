import { DialogActions, DialogContent, DialogTitle } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { Button, Dialog } from '@material-ui/core';
import React from 'react';
import { useState } from 'react';

function LoginDialog() {
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  return (
    <form>
      <Button onClick={handleClickOpen} color="inherit">
        Login
      </Button>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="login-dialog-title"
      >
        <DialogTitle id="login-dialog-title">Login</DialogTitle>
        <DialogContent>
          <div>
            <TextField
              id="username"
              label="Username"
              type="text"
              required
              autoFocus
            />
          </div>
          <div>
            <TextField
              id="password"
              label="Password"
              type="password"
              required
            />
          </div>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={handleClose}>Login</Button>
        </DialogActions>
      </Dialog>
    </form>
  );
}

export default LoginDialog;
