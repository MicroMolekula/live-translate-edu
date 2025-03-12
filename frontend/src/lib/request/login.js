import axios from '@/axios'

async function login(email, password) {
    let response = await axios.post('/api/auth', {
        login: email,
        password: password
    })

    if (response.status === 200) {
        return response.data.token
    } else {
        throw new Error(response.status.toString() + ' ' + response.statusText)
    }
}

export default login