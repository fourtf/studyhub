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

  const [input, setInput] = useState({name: "", password: ""})

  const handleInputChange = (key: keyof Input) => {
    return (event: ChangeEvent<HTMLInputElement>) => {
      setInput({...input, [key]: event.currentTarget.value})
    }
  }

  const onSubmit = async () => {
    const {token} = await fetchJson<TokenResponse>('http://localhost:3001/login', {
      method: 'POST',
    }, input)
    let date = new Date();
    date.setTime(date.getTime() + 365 * 24 * 60 * 60 * 1000); //expires afer one year
    document.cookie = `studyhub_token=${token}; expires=${date.toUTCString()}`;
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
        <DialogContent style={{width: 350}}>
          <div>
            <TextField
              label="Username"
              type="text"
              required
              autoFocus
              value={input.name}
              onChange={handleInputChange("name")}
              fullWidth
            />
          </div>
          <div>
            <TextField
              label="Password"
              type="password"
              required
              value={input.password}
              onChange={handleInputChange("password")}
              fullWidth
            />
          </div>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} variant="contained">Cancel</Button>
          <Button onClick={onSubmit} variant="contained" color="primary">Login</Button>
        </DialogActions>
      </Dialog>
    </form>
  );
}

export default LoginDialog;
