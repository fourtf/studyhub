import { DialogActions, DialogContent, DialogTitle } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { Button, Dialog } from '@material-ui/core';
import { TokenResponse } from 'models/Token';
import React, { ChangeEvent } from 'react';
import { useState } from 'react';
import { fetchPublic } from 'utils/FetchUtils';

function LoginDialog() {
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const [input, setInput] = useState({
    name: '',
    password: '',
  });

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value;
    setInput({
      ...input,
      [event.currentTarget.id]: value,
    });
  };

  const requestOptions: RequestInit = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(input),
  };

  const onSubmit = async () => {
    const responseBody = await fetchPublic('http://localhost:3001/login', requestOptions)
    const { token } = responseBody as TokenResponse;
    let date = new Date();
    date.setTime(date.getTime() + 365 * 24 * 60 * 60 * 1000); //expires in one year
    document.cookie = `studyhub_token=${token}; expires=${date.toUTCString()}; secure`;
    setOpen(false)
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
              id="name"
              label="Username"
              type="text"
              required
              autoFocus
              value={input.name}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <TextField
              id="password"
              label="Password"
              type="password"
              required
              value={input.password}
              onChange={handleInputChange}
            />
          </div>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose}>Cancel</Button>
          <Button onClick={onSubmit}>Login</Button>
        </DialogActions>
      </Dialog>
    </form>
  );
}

export default LoginDialog;
