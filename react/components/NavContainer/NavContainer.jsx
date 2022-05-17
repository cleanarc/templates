import { Route, Routes } from 'react-router-dom';

import Header from '../Header';
import Background from '../Background';

const <% .ComponentName %> = ({ routes }) => {
  return (
    <>
      <Header navRoutes={routes} />
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
