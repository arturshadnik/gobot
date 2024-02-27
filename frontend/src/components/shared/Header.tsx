'use client'

import Image from 'next/image'
import { AppBar, Container, Button, Box } from '@mui/material'
import { useAuth } from "@/lib/auth/authContext"

function Header() {
    const { user, signOut } = useAuth()
    return (
        <>
            <AppBar 
                position="static" 
                sx={{ 
                    paddingTop:'0.5rem', 
                    paddingBottom: '0.5rem', 
                    width:'100vw'
                }}
            >
                <Container maxWidth='xl' sx={{ display: 'flex', justifyContent: 'space-between'}}>
                    <Image src="/goot.png" alt='goot' width={70} height={70} priority={true}/>
                    <Box>
                        <Button onClick={signOut} sx={{ color: "black"}}>
                            Sign Out
                        </Button>
                    </Box>
                </Container>
            </AppBar>
        </>


    );
}

export default Header;


