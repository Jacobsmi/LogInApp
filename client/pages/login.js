import styles from '../styles/Login.module.css'
import {useState} from 'react'

export default function Login() {
  const [errors, setErrors] = useState(false)

  return (
    <div className={styles.container}>
      <h1>Log In</h1>
      <div className={styles.error} id='errors' style={errors ? { display: 'block' } : { display: 'none' }}>

      </div>
      <input type='text' placeholder='First Name' id='signup-fname'></input>
      <input type='text' placeholder='Last Name' id='signup-lname'></input>
      <button className={styles.submitButton}>Log In</button>
    </div>
  )
}