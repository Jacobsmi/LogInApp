import styles from '../styles/Login.module.css'
import { useState } from 'react'

export default function Login() {
  const [errors, setErrors] = useState(false)

  async function processLogin() {
    errorString = 'Errors:<ul>'
    setErrors(false)

    let validEmail = true
    const email = document.getElementById('login-email').value
    if (!/^(([^<>()[\]\\.,;:\s@"]+(\.[^<ul>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email)) {
      errorString += '<li>Invalid E-Mail Name</li>'
      validEmail = false
    }

    let validPass = true
    const pass = document.getElementById('login-pass').value
    if (!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()])[0-9a-zA-Z!@#$%^&*()]{8,}$/.test(pass)) {
      errorString += '<li>Invalid Password</li>'
      validPass = false
    }

    if (validEmail && validPass) {
      const resp = await fetch("http://localhost:5000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({
          "email": email,
          "pass": pass
        })
      })
      const respJSON = await resp.json()
    }
  }

  return (
    <div className={styles.container}>
      <h1>Log In</h1>
      <div className={styles.error} id='errors' style={errors ? { display: 'block' } : { display: 'none' }}>

      </div>
      <input type='text' placeholder='E-Mail' id='login-email'></input>
      <input type='password' placeholder='Password' id='login-pass'></input>
      <button className={styles.submitButton} onClick={processLogin}>Log In</button>
    </div>
  )
}