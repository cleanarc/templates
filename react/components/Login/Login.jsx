import React, { useState, useEffect, useReducer, useContext } from 'react';

import Card from '<% .DepsRelativePath %>/Card';
import Button from '<% .DepsRelativePath %>/Button';

import AuthContext from '<% .DepsRelativePath %>/AuthContext';

import classes from './<% .ComponentName %>.module.css';

const isValidEmail = (email) => {
  return email.includes('@');
};

const isValidPassword = (password) => {
  return password.trim().length > 6;
};

const emailReducer = (prevState, action) => {
  if (action.type === 'EMAIL_INPUT') {
    return { value: action.value, isValid: isValidEmail(action.value) };
  }

  if (action.type === 'EMAIL_BLUR') {
    return { value: prevState.value, isValid: isValidEmail(prevState.value) };
  }

  return { value: '', isValid: false };
};

const passwordReducer = (prevState, action) => {
  if (action.type === 'PASSWORD_INPUT') {
    return { value: action.value, isValid: isValidPassword(action.value) };
  }

  if (action.type === 'PASSWORD_BLUR') {
    return {
      value: prevState.value,
      isValid: isValidPassword(prevState.value),
    };
  }
};

const <% .ComponentName %> = () => {
  const ctx = useContext(AuthContext);

  const [formIsValid, setFormIsValid] = useState(false);

  const [emailState, dispatchEmail] = useReducer(emailReducer, {
    value: '',
    isValid: undefined,
  });

  const [passwordState, dispatchPassword] = useReducer(passwordReducer, {
    value: '',
    isValid: undefined,
  });

  const { isValid: emailIsValid } = emailState;
  const { isValid: passwordIsValid } = passwordState;

  useEffect(() => {
    const timeout = setTimeout(() => {
      console.log('checking');
      setFormIsValid(emailIsValid && passwordIsValid);
    }, 500);

    return () => {
      console.log('cleanup');
      clearTimeout(timeout);
    };
  }, [emailIsValid, passwordIsValid]);

  const emailChangeHandler = (event) => {
    dispatchEmail({ type: 'EMAIL_INPUT', value: event.target.value });
  };

  const passwordChangeHandler = (event) => {
    dispatchPassword({ type: 'PASSWORD_INPUT', value: event.target.value });
  };

  const validateEmailHandler = () => {
    dispatchEmail({ type: 'EMAIL_BLUR' });
  };

  const validatePasswordHandler = () => {
    dispatchPassword({ type: 'PASSWORD_BLUR' });
  };

  const submitHandler = (event) => {
    event.preventDefault();
    ctx.onLogin(emailState.value, passwordState.value);
  };

  return (
    <Card className={classes.login}>
      <form onSubmit={submitHandler}>
        <div
          className={`${classes.control} ${
            emailState.isValid === false ? classes.invalid : ''
          }`}
        >
          <label htmlFor='email'>E-Mail</label>
          <input
            type='email'
            id='email'
            value={emailState.value}
            onChange={emailChangeHandler}
            onBlur={validateEmailHandler}
          />
        </div>
        <div
          className={`${classes.control} ${
            passwordState.isValid === false ? classes.invalid : ''
          }`}
        >
          <label htmlFor='password'>Password</label>
          <input
            type='password'
            id='password'
            value={passwordState.value}
            onChange={passwordChangeHandler}
            onBlur={validatePasswordHandler}
          />
        </div>
        <div className={classes.actions}>
          <Button type='submit' className={classes.btn} disabled={!formIsValid}>
            Login
          </Button>
        </div>
      </form>
    </Card>
  );
};

export default <% .ComponentName %>;
