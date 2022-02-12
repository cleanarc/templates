import classes from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({ children }) => {
  return (
    <>
      <div className={classes['background-top']} />
      <div className={classes['background-bottom']} />
      <div className={classes.content}>{children}</div>
    </>
  );
};

export default <% .ComponentName %>;
