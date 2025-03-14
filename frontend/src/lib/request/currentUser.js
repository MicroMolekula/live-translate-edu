import axios from '@/axios'

export default async function currentUser(token) {
    let response = await axios.get('/api/me', {
        headers: {
            'Authorization': 'Bearer ' + token
        }
    })
    if (response.status !== 200) {
        throw new Error('unauthorized')
    } else {
        return response.data
    }
}