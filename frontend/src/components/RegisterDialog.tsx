import { DialogActions } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { DialogTitle } from '@material-ui/core';
import { DialogContent } from '@material-ui/core';
import { Dialog } from '@material-ui/core';
import { Button } from '@material-ui/core';
import React, { ChangeEvent } from 'react';
import { useState } from 'react';
import { fetchPublic } from 'utils/FetchUtils';

function RegisterDialog() {

  const [open, setOpen] = useState(false);
  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  const [input, setInput] = useState({
    name: "",
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
    headers: {'Content-Type': 'application/json'},
    body: JSON.stringify(input)
  }
  const onSubmit = async() => {
    const responseBody = await fetchPublic('http://localhost:3001/register', requestOptions);
    setOpen(false)
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
        <DialogTitle id="login-dialog-title">Register</DialogTitle>
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
