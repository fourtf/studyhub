import { DialogActions } from '@material-ui/core';
import { TextField } from '@material-ui/core';
import { DialogTitle } from '@material-ui/core';
import { DialogContent } from '@material-ui/core';
import { Dialog } from '@material-ui/core';
import { Button } from '@material-ui/core';
import { TokenResponse } from 'models/Responses';
import React, { ChangeEvent, useMemo } from 'react';
import { useState } from 'react';
import { fetchJson } from 'utils/FetchUtils';

function RegisterDialog() {
  const [open, setOpen] = useState(false);
  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  type Input = {
    name: string;
    email: string;
    password: string;
  };
  const [input, setInput] = useState({ name: '', email: '', password: '' });
  const handleInputChange = (key: keyof Input) => {
    return (event: ChangeEvent<HTMLInputElement>) => {
      setInput({ ...input, [key]: event.currentTarget.value });
    };
  };

  const onSubmit = async () => {
    if (isValid.isInputValid) {
      const responseBody = await fetchJson(
        'http://localhost:3001/register',
        {
          method: 'POST',
        },
        input
      );
      const { token } = responseBody as TokenResponse;
      let date = new Date();
      date.setTime(date.getTime() + 365 * 24 * 60 * 60 * 1000); //expires afer one year
      document.cookie = `studyhub_token=${token}; expires=${date.toUTCString()}; secure`;
      setOpen(false);
    }
  };

  //minimum eight characters and one special character
  const validatePassword = (password: string): boolean => {
    const includesSpecialCharacter = /[@$!%*#?&]/.test(password)
    const isLongEnough = password.length > 7;
    return isLongEnough && includesSpecialCharacter;
  };

  const isValid = useMemo(() => {
    //ASCII letters and digits, hyphens, underscores and spaces as internal seperators
    const nameRegEx = /^[A-Za-z0-9]+(?:[ _-][A-Za-z0-9]+)*$/;
    //anystring@anystring.anystring
    const emailRegEx = /^\S+@\S+\.\S+$/;

    const isNameValid = nameRegEx.test(input.name);
    const isEmailValid = emailRegEx.test(input.email);
    const isPasswordValid = validatePassword(input.password);
    return {
      isNameValid,
      isEmailValid,
      isPasswordValid,
      isInputValid: isNameValid && isEmailValid && isPasswordValid,
    };
  }, [input]);

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
              label="Username"
              type="text"
              required
              autoFocus
              value={input.name}
              onChange={handleInputChange('name')}
              error={!isValid.isNameValid}
            />
          </div>
          <div>
            <TextField
              label="E-Mail"
              type="email"
              required
              value={input.email}
              onChange={handleInputChange('email')}
              error={!isValid.isEmailValid}
            />
          </div>
          <div>
            <TextField
              label="Password"
              type="password"
              required
              value={input.password}
              onChange={handleInputChange('password')}
              error={!isValid.isPasswordValid}
              helperText="At least eight characters, including at least one letter and one special character (eg. * or !)"
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
