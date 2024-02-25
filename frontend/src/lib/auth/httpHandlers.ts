import axios from 'axios'
import { User } from 'firebase/auth'

type Level = 'easy' | 'medium' | 'hard'

function formatHeaders(user: User) {
    return {
        headers: {
            Authorization: `Bearer ${user?.getIdToken}`
        }
    }
}

export async function sendMessage(user: User, message: string, level: Level) {
    const headers = formatHeaders(user)
    return axios.post(`/chat/${user.uid}?message=${message}&level=${level}`, {}, headers)
}

