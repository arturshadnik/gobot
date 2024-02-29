import React from "react";
import { Card, CardContent, Typography } from '@mui/material';

interface MessageProps {
    role: string;
    content: string;
    timestamp: string;
}

const Message: React.FC<MessageProps> = ({ role, content, timestamp }) => {
    return (
        <Card variant="outlined">
            <CardContent>
                <Typography color={"textSecondary"} gutterBottom>
                    {role}
                </Typography>
                <Typography variant="body2" component={"p"}>
                    {content}
                </Typography>
                <Typography color={"textSecondary"} variant="caption">
                    {timestamp}
                </Typography>
            </CardContent>
        </Card>
    )
}

export default Message;
export type { MessageProps };