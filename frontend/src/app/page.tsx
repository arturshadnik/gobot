'use client'

import React from "react"
import Image from "next/image";
import styles from "./page.module.css";
import { Box } from "@mui/material"
import withAuthProtection from "./hoc/withAuthProtection";
import ChatWindow from "@/components/chat/ChatWIndow"

function Home() {
  
  return (
    <main className={styles.main}>
      <Box>
        <ChatWindow />
      </Box>
    </main>
  );
}

export default withAuthProtection(Home)
