'use client'

import React from "react";
import { Box, Typography, Button } from '@mui/material';
import { useAuth } from "@/lib/auth/authContext"
import handleGoogleLogin from "@/lib/auth/authHandlers"
import GoogleSignIn from "@/components/shared/GoogleSignIn";

function withAuthProtection(Component: React.ComponentType<any>) {
    return function ProtectedRoute(props: any) {
        const { user } = useAuth();

        if (!user) {
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
                    <Typography variant="h6" sx={{ fontWeight: "bold" }}>Log in to continue</Typography>
                        <Button variant="contained" color="primary" onClick={() => handleGoogleLogin()}>
                            <GoogleSignIn className="oauth-provider"/>
                        </Button>
                </Box>
            );
        } else {
            return <Component {...props} />;
        }
    };
}

export default withAuthProtection;