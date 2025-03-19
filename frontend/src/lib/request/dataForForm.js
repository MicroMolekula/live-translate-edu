import axios from '@/axios'

export async function getDataForForm(token) {
    let response = await axios.get('/api/lesson/form/data', {
        headers: {
            Authorization: 'Bearer ' + token
        }
    })
    if (response.status === 200) {
        return response.data
    } else {
        throw new Error('error get form data')
    }
}