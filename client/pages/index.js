import styles from '../styles/Home.module.css'
import Link from 'next/link'

export default function Home() {
  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <div className={styles.leftHeader}>
          LogIn App
        </div>
        <div className={styles.rightHeader}>
          <Link href='/signup'>
            Sign Up
          </Link>
          <Link href='/login'>
            Login
          </Link>
        </div>
      </header>
      <main>
        <h1 className={styles.mainHead}>The Web App that allows for easy JWT auth.</h1>
      </main>
    </div>
  )
}
