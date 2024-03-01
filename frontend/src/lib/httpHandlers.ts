import axios from 'axios'
import { User } from 'firebase/auth'

const backend_url = process.env.NEXT_PUBLIC_BACKEND_URL;

async function formatHeaders(user: User) {
    return {
        headers: {
            Authorization: `Bearer ${await user?.getIdToken()}`
        }
    }
}

export async function sendMessage(user: string, message: string, level: string) {
    // const headers = await formatHeaders(user);
    return await axios.post(`${backend_url}/chat/${user}?message=${message}&level=${level}`)
}

export async function fetchMessages(user: string, level: string) {
    // const headers = await formatHeaders(user);
    return await axios.get(`${backend_url}/chat/${user}?level=${level}`)
}