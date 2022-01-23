import styles from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({ type, onClick, className, children }) => {
  return (
    <button
      className={`${styles.button} ${className}`}
      type={type || 'button'}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default <% .ComponentName %>;
