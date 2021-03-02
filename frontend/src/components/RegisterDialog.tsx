import { DialogActions } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { DialogTitle } from '@material-ui/core';
import { DialogContent } from '@material-ui/core';
import { Dialog } from '@material-ui/core';
import { Button } from '@material-ui/core';
import React from 'react';
import { useState } from 'react';

function RegisterDialog() {
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
        Register
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
            <TextField id="email" label="E-Mail" type="email" required />
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
          <Button onClick={handleClose}>Register</Button>
        </DialogActions>
      </Dialog>
    </form>
  );
}

export default RegisterDialog;
