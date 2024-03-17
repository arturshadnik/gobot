import React, { useLayoutEffect, useState } from "react";
import { Card, CardContent, Typography } from '@mui/material';
import { User } from "firebase/auth"

interface MessageProps {
    role: string;
    content: string;
    timestamp: string;
    user: User;
}

const Message: React.FC<MessageProps> = ({ role, content, timestamp, user }) => {
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
                const name = user.displayName ? user.displayName : user.email
                return name;
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