import styles from '../styles/Signup.module.css'
import {useState} from 'react'

export default function Signup() {
  
  const [errors, setErrors] = useState(false)

  async function processSignUp(){
    let errorString = 'Error(s):<ul>'
    setErrors(false)

    let validFirstName = true
    const firstName = document.getElementById('signup-fname').value
    if (!/^[A-Za-z\'\-]+$/.test(firstName)) {
      errorString += '<li>Invalid First Name</li>'
      validFirstName = false
    }

    let validLastName = true
    const lastName = document.getElementById('signup-lname').value
    if (!/^[A-Za-z\'\-]+$/.test(lastName)) {
      errorString += '<li>Invalid Last Name</li>'
      validLastName = false
    }

    let validEmail = true
    const email = document.getElementById('signup-email').value
    if(!/^(([^<>()[\]\\.,;:\s@"]+(\.[^<ul>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(email)){
      errorString += '<li>Invalid E-Mail Name</li>'
      validEmail = false
    }

    let validPass = true
    const pass = document.getElementById('signup-pass').value
    if(!/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%^&*()])[0-9a-zA-Z!@#$%^&*()]{8,}$/.test(pass)){
      errorString += '<li>Invalid Password</li>'
      validPass = false
    }

    let passMatch = true
    const cpass = document.getElementById('signup-cpass').value
    if (cpass !== pass) {
      errorString += '<li>Passwords do not match</li>'
      passMatch = false
    }
    
    errorString += '</ul>'
    if ( validFirstName && validLastName && validEmail && validPass && passMatch){
      const resp = await fetch("http://localhost:5000/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          "fname": firstName,
          "lname": lastName,
          "email": email,
          "pass": pass
        })
      })
      const respJSON = await resp.json()
      console.log(respJSON)
    }else{
      setErrors(true)
      document.getElementById('errors').innerHTML = errorString
    }
  }

  return(
    <div className={styles.container}>
      <div className={styles.error} id='errors' style={ errors? {display: 'block'}: {display: 'none'}}>
        
      </div>
      <input type='text' placeholder='First Name' id='signup-fname'></input>
      <input type='text' placeholder='Last Name' id='signup-lname'></input>
      <input type='text' placeholder='Email' id='signup-email'></input>
      <input type='password' placeholder='Password' id='signup-pass'></input>
      <input type='password' placeholder='Confirm Password' id='signup-cpass'></input>
      <button className={styles.submitButton} onClick={processSignUp}>Sign Up</button>
    </div>
  )
}