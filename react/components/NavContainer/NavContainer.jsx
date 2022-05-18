import { Route, Routes } from 'react-router-dom';

import NavHeader from '../NavHeader';
import Background from '../Background';

const <% .ComponentName %> = ({ routes }) => {
  return (
    <>
      <NavHeader navRoutes={routes} />
      <main>
        <Background>
          <Routes>
            {routes.map((route) => (
              <Route
                key={route.name}
                path={route.path}
                element={route.element}
              />
            ))}
          </Routes>
        </Background>
      </main>
    </>
  );
};

export default <% .ComponentName %>;
