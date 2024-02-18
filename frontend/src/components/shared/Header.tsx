import Image from 'next/image'
import { AppBar, Container } from '@mui/material'

function Header() {
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
                <Container maxWidth='xl'>
                    <Image src="/goot.png" alt='goot' width={70} height={70} priority={true}/>
                    
                </Container>

            </AppBar>
        </>


    );
}

export default Header;


