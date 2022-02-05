import classes from './<% .ComponentName %>.module.css';

const <% .ComponentName %> = ({
  textArea,
  rows,
  className,
  valid,
  label,
  type,
  id,
  value,
  onChange,
  onBlur,
}) => {
  const field = textArea ? (
    <textarea
      id={id}
      rows={rows}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
    />
  ) : (
    <input
      id={id}
      type={type}
      value={value}
      onChange={onChange}
      onBlur={onBlur}
    />
  );

  return (
    <div
      className={`${classes.field} ${
        valid === false ? classes.invalid : ''
      } ${className}`}
    >
      <label htmlFor={id}>{label}</label>
      {field}
    </div>
  );
};

export default <% .ComponentName %>;
