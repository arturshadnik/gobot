'use client'

import Image from "next/image";
import styles from "./page.module.css";
import withAuthProtection from "./hoc/withAuthProtection";

function Home() {
  return (
    <main className={styles.main}>
      <div>Hello</div>
    </main>
  );
}

export default withAuthProtection(Home)
