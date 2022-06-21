import Card from '<% .DepsRelativePath %>/Card';
import NavContainer from '<% .DepsRelativePath %>/NavContainer';

const appRoutes = [
  { name: 'Home', path: '/', element: <Card>Welcome Home!</Card> },
  // { name: 'YourPage', path: '/Pathname', element: <YourComponent /> },
];

const <% .ComponentName %> = () => {
  return <NavContainer routes={appRoutes} />;
};

export default <% .ComponentName %>;
