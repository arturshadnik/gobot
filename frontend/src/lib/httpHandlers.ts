import axios from 'axios'
import { User } from 'firebase/auth'

type Level = 'easy' | 'medium' | 'hard'
const backend_url = process.env.NEXT_PUBLIC_BACKEND_URL;

async function formatHeaders(user: User) {
    return {
        headers: {
            Authorization: `Bearer ${await user?.getIdToken()}`
        }
    }
}

export async function sendMessage(user: User, message: string, level: Level) {
    const headers = await formatHeaders(user);
    return await axios.post(`${backend_url}/chat/${user.uid}?message=${message}&level=${level}`, {}, headers)
}
export async function fetchMessages(user: User, level: Level) {
    const headers = await formatHeaders(user);
    return await axios.get(`${backend_url}/chat/${user.uid}?level=${level}`, headers)
}