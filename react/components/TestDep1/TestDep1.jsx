import classes from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({ children, className }) => {
  return (
    <div className={`${classes.card} ${className}`}>
      {children}
    </div>
  )
};

export default <% .ComponentName %>;
