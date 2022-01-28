import ReactDOM from 'react-dom';

import Card from './Card';
import Button from './Button';
import Backdrop from './Backdrop';

import styles from './<% .ComponentName %>.module.css';

const <% .ComponentName %>Overlay = ({ title, message, onClose }) => {
  return (
    <Card className={styles.modal}>
      <header className={styles.header}>
        <h2>{title}</h2>
      </header>
      <div className={styles.content}>
        <p>{message}</p>
      </div>
      <footer className={styles.actions}>
        <Button onClick={onClose}>Okay</Button>
      </footer>
    </Card>
  );
};

const <% .ComponentName %> = ({ title, message, onClose }) => {
  return (
    <>
      {ReactDOM.createPortal(
        <Backdrop onClose={onClose} />,
        document.getElementById('backdrop-root')
      )}
      {ReactDOM.createPortal(
        <<% .ComponentName %>Overlay title={title} message={message} onClose={onClose} />,
        document.getElementById('overlay-root')
      )}
    </>
  );
};

export default <% .ComponentName %>;
