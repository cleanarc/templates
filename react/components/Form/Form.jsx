import Button from '<% .DepsRelativePath %>/Button';
import Card from '<% .DepsRelativePath %>/Card';

import classes from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({ className, title, onSubmit, valid, children }) => {
  return (
    <Card className={className}>
      <h2 className={classes.title}>{title}</h2>
      <form onSubmit={onSubmit}>
        {children}
        <div className={classes.actions}>
          <Button type='submit' disabled={!valid}>
            Submit
          </Button>
        </div>
      </form>
    </Card>
  );
};

export default <% .ComponentName %>;
