import styles from '../styles/Login.module.css'
import { useState } from 'react'
import { useRouter } from 'next/router'

export default function Login() {
  const [errors, setErrors] = useState(false)
  const router = useRouter()

  async function processLogin() {
    let errorString = 'Errors:<ul>'
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
      console.log("Making API call")
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
      console.log(respJSON)
      if (respJSON.success === false) {
        setErrors(true)
        if (respJSON.msg === 'user_not_exist') {
          errorString += '<li>User with that email does not exist</li>'
        } else if (respJSON.msg === "wrong_pass") {
          errorString += '<li>Incorrect Password</li>'
        } else {
          errorString += '<li>Server error when querying for user. Please see server or contact administrator</li>'
        }
      } else if (respJSON.success === true) {
        router.push("/home")
      }
    } else {
      setErrors(true)
    }
    errorString += '</ul>'
    document.getElementById('errors').innerHTML = errorString
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