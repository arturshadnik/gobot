import React, { useEffect, useRef } from "react";
import { List } from "@mui/material";
import Message, { MessageProps } from "@/components/chat/Message";
import { User } from "firebase/auth"

interface MessageListProps {
    messages: MessageProps[];
    user: User
}

const MessageList: React.FC<MessageListProps> = ({ messages, user }) => {
    const endOfMessageRef = useRef<null | HTMLDivElement>(null);

    const scrollToBottom = () => {
        endOfMessageRef.current?.scrollIntoView({ behavior: "smooth"});
    };

    useEffect(() => {
        scrollToBottom();
    }, [messages]);
    
    return (
        <List sx={{ maxHeight: "700px", overflow: "auto"}}>
            {messages.map((message, index) => (
                <Message key={index} role={message.role} content={message.content} timestamp={message.timestamp} user={user}/>
            ))}
            <div ref={endOfMessageRef} />
        </List>
    )
}

export default MessageList;

export type { MessageListProps };