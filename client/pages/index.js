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

      </main>
    </div>
  )
}
