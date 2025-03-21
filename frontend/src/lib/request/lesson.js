import axios from '@/axios'


export async function createLesson(token, data) {
    try {
        return await axios.post('/api/lesson/create', data, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('ошибка создания занятия')
    }
}

export async function getLesson(token) {
    try {
        return await axios.get('/api/lesson', {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('ошибка получения списка занятий')
    }
}

export async function getRoomToken(token, roomName) {
    try {
        return await axios.get('/api/user/room_token?room=' + roomName, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('Ошибка получения токена подключения к комнате')
    }
}

export async function connectRoom(token, roomName) {
    try {
        return await axios.get('/api/connect?room='+roomName, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('ошибка подключения к распознаванию речи')
    }
}

export async function disconnectRoom(token, roomName) {
    try {
        return await axios.get('/api/disconnect?room='+roomName, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('ошибка отключения от распознавания речи')
    }
}

export async function getUsersInLesson(token, roomName) {
    try {
        return await axios.get(`/api/chat/${roomName}/users`, {
            headers: {
                Authorization: 'Bearer ' + token
            }
        })
    } catch {
        throw new Error('ошибка получения списка участников занятия')
    }
}