import React, { useState, useEffect } from "react";
import { Container, Box, TextField, Button } from "@mui/material";
import { Save } from "@mui/icons-material"
import { usePathname, useSearchParams, useRouter } from "next/navigation";
import MessageList from "@/components/chat/MessageList";
import MessageInput from "@/components/chat/MessageInput";
import { MessageProps } from "@/components/chat/Message"
import LevelDropdown from "../LevelDropdown";
import { sendMessage, fetchMessages } from "@/lib/httpHandlers"
import { AxiosError } from "axios";

const ChatWindow: React.FC = () => {
    const [messages, setMessages] = useState<Array<MessageProps>>([])
    const [apiKey, setApiKey] = useState<string>("")
    const [errorMessage, setErrorMessage] = useState<string>("")
    const [userText, setUserText] = useState<string>("")
    const [user, setUser] = useState<string>("")
    const [userError, setUserError] = useState<string>("")

    const searchParams = useSearchParams()
    const router = useRouter()
    const pathname = usePathname()

    const level: string = searchParams.get('level') || 'easy'

    useEffect(() => {
        fetchMessages(user, level)
            .then(
                (messages) => {
                    setMessages(messages.data)
                },
                (error: AxiosError) => {
                    setMessages([])
                    console.log("Failed to fetch old messages: ", error)
                }
            ).catch(
                (error) => {
                    console.log(error)
                }
            )
    }, [level, user]);


    useEffect(() => {

    })

    const handleSendMessage = (newMessage: string) => {
        if (level !== 'easy' && apiKey === "") {
            setErrorMessage("Please enter an API key to access the final level")
        } else if (user === "") {
            setUserError("Please enter your name")
        } else {
            const newMessageProp: MessageProps = { role: "user", content: newMessage, timestamp: Date.toString(), user: user }
            setMessages([...messages, newMessageProp])
            sendMessage(user, newMessage, level, apiKey)
                .then(
                    (resp) => {
                        if (resp.status === 204) {
                            setMessages([])
                        } else {
                            const remappedData: MessageProps = {
                                role: resp.data.Role,
                                content: resp.data.Content,
                                timestamp: resp.data.Timestamp,
                                user: ""
                            };
                            setMessages((prevMessages) => [...prevMessages, remappedData])
                        }

                    },
                    (error: any) => {
                        console.error("Failer to send: ", error)
                    }
                )
        }
    }

    const handleLevelChange = (newLevel: string) => {
        setErrorMessage("")
        router.push(pathname + '?level=' + newLevel);
    }

    const handleUserChange = (newUser: string) => {
        setUserError("")
        const almostUniqueUser = `${newUser}-${Date.now()}`
        setUser(almostUniqueUser)
    }

    return (
        <Container sx={{
            width: "100%",
            border: "1px solid #ccc",
            padding: "10px",
            margin: "auto",
            boxSizing: "border-box",

        }}>

            <Box sx={{ display: "flex", padding: "5px" }}>
                <TextField
                    label="Name"
                    variant="outlined"
                    fullWidth
                    value={userText}
                    onChange={(e) => { setUserText(e.target.value), setUserError("") }}
                    error={!!userError}
                    helperText={userError}
                    onKeyDown={(e) => e.key === 'Enter' && setUser(userText)}
                    disabled={user !== ''}
                />
                <Button variant="contained" color="primary" onClick={() => handleUserChange(userText)} disabled={user !== ''}>
                    <Save />
                </Button>
            </Box>
            <Box sx={{ display: "flex" }}>
                <LevelDropdown level={level} levels_list={["easy", "medium", "hard"]} handleLevelChange={handleLevelChange} />
                <TextField
                    label="OpenAI Key"
                    variant="outlined"
                    fullWidth
                    value={apiKey}
                    onChange={(e) => { setApiKey(e.target.value), setErrorMessage("") }}
                    error={!!errorMessage}
                    helperText={errorMessage}
                />
            </Box>
            <MessageList messages={messages} user={user} />
            <MessageInput onSendMessage={handleSendMessage} />
        </Container>
    )
}

export default ChatWindow