import React, { useState } from "react";
import { Box, TextField, Button } from "@mui/material";
import { Send } from "@mui/icons-material"

interface MessageInputProps {
    onSendMessage: (message: string) => void;
}

const MessageInput: React.FC<MessageInputProps> = ({ onSendMessage }) => {
    const [message, setMessage] = useState<string>("");

    const handleSend = () => {
        if (message.trim() !== '') {
            onSendMessage(message);
            setMessage("");
        }
    };

    return (
        <Box sx={{ display: "flex"}}>
            <TextField
                label="Message"
                variant="outlined"
                fullWidth
                value={message}
                onChange={(e) => setMessage(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && handleSend()}
            />
            <Button variant="contained" color="primary" onClick={handleSend}>
                <Send />
            </Button>
        </Box>
    );

};

export default MessageInput;
export type { MessageInputProps };