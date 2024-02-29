import React, { useState, useEffect } from "react";
import { Container } from "@mui/material";
import MessageList from "@/components/chat/MessageList";
import MessageInput from "@/components/chat/MessageInput";
import { MessageProps } from "@/components/chat/Message"

import { useAuth } from "@/lib/auth/authContext"
import { sendMessage, fetchMessages } from "@/lib/httpHandlers"

const ChatWindow: React.FC = () => {
    const [messages, setMessages] = useState<Array<MessageProps>>([])
    const { user } = useAuth()
    const level = 'easy'
    // const messages: MessageProps[] = [
    //     {
    //       role: "user",
    //       content: "sup",
    //       timestamp: new Date().toString(),
    //     },
    //     {
    //       role: "assistant",
    //       content: "Hello!",
    //       timestamp: new Date().toString(),
    //     },    
    //   ]

    useEffect(() => {
        fetchMessages(user!, level)
        .then(
            (messages) => {
                setMessages([messages.data])
            },
            (error: any) => {
                console.log("Failed to fetch old messages: ", error)
            }
        )
        //setMessages(fetchedMessages)
    });

    const handleSendMessage = (newMessage: string) => {
        const newMessageProp: MessageProps = { role: "user", content: newMessage, timestamp: Date.toString() }
        setMessages([...messages, newMessageProp])
        sendMessage(user!, newMessage, level)
        .then(
            (resp) => {
                setMessages([...messages, resp.data])
            }, 
            (error: any) => {
                console.error("Failer to send: ", error)
            }
        )
        
    }

    return (
        <Container>
            <MessageList messages={messages} />
            <MessageInput onSendMessage={handleSendMessage} />
        </Container>
    )
}

export default ChatWindow