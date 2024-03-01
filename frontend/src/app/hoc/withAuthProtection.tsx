'use client'

import React, { useState, useEffect } from "react";
import { Box, Typography, Button, TextField } from '@mui/material';
import { Save } from "@mui/icons-material";
import { useRouter } from "next/navigation" 

function withAuthProtection(Component: React.ComponentType<any>) {
    return function ProtectedRoute(props: any) {
        // const [userName, setUserName] = useState<string | null>(localStorage.getItem('userName'));
        const [userText, setUserText] = useState<string>('');
        const [userError, setUserError] = useState<string>('');
        const [loggedIn, setLoggedIn] = useState<boolean>(false)
        const router = useRouter()
        // useEffect(() => {
        //     // Update the userName state if it changes in localStorage
        //     const handleStorageChange = () => {
        //         setUserName(localStorage.getItem('userName'));
        //     };

        //     window.addEventListener('storage', handleStorageChange);

        //     return () => {
        //         window.removeEventListener('storage', handleStorageChange);
        //     };
        // }, []);

        const handleUserChange = (name: string) => {
            if (name.trim() === '') {
                setUserError('Name cannot be empty');
                return;
            }

            const almostUniqueName = `${name}-${Date.now()}` // if anyone reads this, dont laugh. firestore auth is blocked by HCL :)))
            // localStorage.setItem('userName', almostUniqueName);
            // setUserName(name);
            setLoggedIn(true)
            router.push(`/${almostUniqueName}`);

        };

        if (!loggedIn) {
            return (
                <Box sx={{
                    display: "flex",
                    flexDirection: "column",
                    alignItems: "center",
                    justifyContent: "center",
                    gap: "1rem",
                    width: "100vw",
                    height: "100vh"                   
                }}>
                    <Typography variant="h6" sx={{ fontWeight: "bold" }}>Enter your name to continue</Typography>
                    <Box sx={{ display: "flex", padding: "10px"}}>
                        <TextField
                            label="Name"
                            variant="outlined"
                            fullWidth
                            value={userText}
                            onChange={(e) => {setUserText(e.target.value); setUserError("");}}
                            error={!!userError}
                            helperText={userError}
                            onKeyDown={(e) => e.key === 'Enter' && handleUserChange(userText)}
                        />
                        <Button variant="contained" color="primary" onClick={() => handleUserChange(userText)}>
                            <Save />
                        </Button>
                    </Box>
                </Box>
            );
        } else {
            return <Component {...props} />;
        }
    };
}

export default withAuthProtection;