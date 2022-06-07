import Card from './Card';
import NavContainer from './NavContainer';

const appRoutes = [
  { name: 'Home', path: '/', element: <Card>Welcome Home!</Card> },
  // { name: 'YourPage', path: '/Pathname', element: <YourComponent /> },
];

const <% .ComponentName %> = () => {
  return <NavContainer routes={appRoutes} />;
};

export default <% .ComponentName %>;
