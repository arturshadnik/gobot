import React, { useLayoutEffect, useState } from "react";
import { Card, CardContent, Typography } from '@mui/material';
import { useAuth } from "@/lib/auth/authContext";

interface MessageProps {
    role: string;
    content: string;
    timestamp: string;
}

const Message: React.FC<MessageProps> = ({ role, content, timestamp }) => {
    const { user } = useAuth()
    const getBackgroundColor = (role: string) => {
        switch (role) {
            case 'user':
                return '#f0f0f0';
            case 'assistant':
                return '#ffffff';
        }
    }

    const formatRole = (role: string) => {
        switch (role) {
            case 'user':
                return user?.displayName;
            case 'assistant':
                return "Goot"
        }
    }
    
    return (
        <Card variant="outlined" sx={{ backgroundColor: getBackgroundColor(role)}}>
            <CardContent>
                <Typography color={"textSecondary"} gutterBottom>
                    {formatRole(role)}
                </Typography>
                <Typography variant="body2" component={"p"}>
                    {content}
                </Typography>
                {/* <Typography color={"textSecondary"} variant="caption">
                    {timestamp}
                </Typography> */}
            </CardContent>
        </Card>
    )
}

export default Message;
export type { MessageProps };