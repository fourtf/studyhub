import { DialogActions } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { DialogTitle } from '@material-ui/core';
import { DialogContent } from '@material-ui/core';
import { Dialog } from '@material-ui/core';
import { Button } from '@material-ui/core';
import { SettingsInputAntennaTwoTone } from '@material-ui/icons';
import React, { ChangeEvent, FormEvent } from 'react';
import { useState } from 'react';

function RegisterDialog() {

  const [open, setOpen] = useState(false);
  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  const [input, setInput] = useState({
    username: "",
    email: "",
    password: ""
  })
  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    const value = event.currentTarget.value
    setInput({
      ...input, [event.currentTarget.id]: value
    });
  }

  const requestOptions = {
    method: "POST",
    headers: {'Content-Type': 'application/json',
              'Access-Control-Allow-Origin': 'http://localhost:3001'},
    body: JSON.stringify(input)
  }
  const onSubmit = async() => {
    const response = await fetch('http://localhost:3001/register', requestOptions);
    console.log("Response:\n" + response)
  }


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
              value={input.username}
              onChange={handleInputChange}
            />
          </div>
          <div>
            <TextField id="email" label="E-Mail" type="email" required value={input.email} onChange={handleInputChange}/>
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
          <Button onClick={onSubmit}>Register</Button>
        </DialogActions>
      </Dialog>
    </form>
  );
}

export default RegisterDialog;
