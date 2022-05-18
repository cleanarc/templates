import { NavLink, useLocation } from 'react-router-dom';

import classes from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({ navRoutes }) => {
  const location = useLocation();

  const isActive = (path) => {
    return path === location.pathname;
  };

  return (
    <>
      <header className={classes.header}>
        <h3>WCBR</h3>
        <nav
          style={{
            width: '30rem',
          }}
        >
          {navRoutes.map((route) => (
            <NavLink
              key={route.name}
              to={route.path}
              className={isActive(route.path) ? classes.active : classes.link}
            >
              {route.name}
            </NavLink>
          ))}
        </nav>
      </header>
    </>
  );
};

export default <% .ComponentName %>;
