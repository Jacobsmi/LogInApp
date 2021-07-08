import styles from '../styles/Signup.module.css'

export default function Signup() {
  return(
    <div className={styles.container}>
      <input type='text' placeholder='First Name'></input>
      <input type='text' placeholder='Last Name'></input>
      <input type='text' placeholder='Email'></input>
      <input type='password' placeholder='Password'></input>
      <input type='password' placeholder='Confirm Password'></input>
      <button className={styles.submitButton}>Sign Up</button>
    </div>
  )
}