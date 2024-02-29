import React from "react";
import { List } from "@mui/material";
import Message, { MessageProps } from "@/components/chat/Message";

interface MessageListProps {
    messages: MessageProps[];
}

const MessageList: React.FC<MessageListProps> = ({ messages }) => {
    return (
        <List>
            {messages.map((message, index) => (
                <Message key={index} role={message.role} content={message.content} timestamp={message.timestamp} />
            ))}
        </List>
    )
}

export default MessageList;

export type { MessageListProps };