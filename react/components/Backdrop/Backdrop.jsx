import styles from './Backdrop.module.css';

const Backdrop = ({ onClose }) => (
  <div className={styles.backdrop} onClick={onClose} />
);

export default Backdrop;
