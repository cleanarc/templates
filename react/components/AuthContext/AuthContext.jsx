import React, { useState, useEffect } from 'react';

const <% .ComponentName %> = React.createContext({
  isLoggedIn: false,
  onLogin: () => {},
  onLogout: () => {},
});

export const <% .ComponentName %>Provider = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const loggedInIndicator = localStorage.getItem('loggedIn');

    if (loggedInIndicator === 'true') {
      setIsLoggedIn(true);
    }
  }, []);

  const loginHandler = (email, password) => {

    // -------------------------------------------
    // TODD implement your authentication here...
    // -------------------------------------------

    localStorage.setItem('loggedIn', 'true');

    setIsLoggedIn(true);
  };

  const logoutHandler = () => {

    // -------------------------------------------
    // TODO invalidate and auth tokens here...
    // -------------------------------------------

    localStorage.setItem('loggedIn', 'false');
    setIsLoggedIn(false);
  };

  return (
    <<% .ComponentName %>.Provider
      value={{
        isLoggedIn: isLoggedIn,
        onLogin: loginHandler,
        onLogout: logoutHandler,
      }}
    >
      {children}
    </<% .ComponentName %>.Provider>
  );
};

export default <% .ComponentName %>;
