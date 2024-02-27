import React, { useState } from "react";
import { useAuth } from '@/lib/auth/authContext';
import { Box, Typography, Button } from '@mui/material';
import GoogleSignIn from "@/components/shared/GoogleSignIn";
import handleGoogleLogin from "@/lib/auth/authHandlers";

function withAuthProtection(Component: React.ComponentType<any>) {
    return function ProtectedRoute(props: any) {
        const { user, userLoading} = useAuth()
        const [error, setError] = useState<string>("")

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
                    <Typography variant="h6" sx={{ fontWeight: "bold" }}>Sign in to continue</Typography>
                    <Button 
                        onClick={handleGoogleLogin}
                        sx={{ width: "fit-content", border: "none", padding: "0" }}>
                        <GoogleSignIn className="h-6 w-6"/>
                    </Button>
                </Box>
            )
        } else {
            return <Component {...props} />
        }
    }
}

export default withAuthProtection