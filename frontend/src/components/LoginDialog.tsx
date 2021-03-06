import { DialogActions, DialogContent, DialogTitle } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { Button, Dialog } from '@material-ui/core';
import { TokenResponse } from 'models/Responses';
import React, { ChangeEvent } from 'react';
import { useState } from 'react';
import { fetchJson } from 'utils/FetchUtils';

function LoginDialog() {
  const [open, setOpen] = useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  type Input = {
    name: string,
    password: string
  }

  let input: Input = {name: "", password: ""}

  const handleInputChange = (key: keyof Input) => {
    return (event: ChangeEvent<HTMLInputElement>) => {
      input[key] = event.currentTarget.value
    }
  }

  const onSubmit = async () => {
    const responseBody = await fetchJson('http://localhost:3001/login', {
      method: 'POST',
    }, input)
    const { token } = responseBody as TokenResponse;
    let date = new Date();
    date.setTime(date.getTime() + 365 * 24 * 60 * 60 * 1000); //expires afer one year
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
              label="Username"
              type="text"
              required
              autoFocus
              value={input.name}
              onChange={handleInputChange("name")}
            />
          </div>
          <div>
            <TextField
              label="Password"
              type="password"
              required
              value={input.password}
              onChange={handleInputChange("password")}
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
