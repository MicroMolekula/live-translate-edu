import {defineStore} from "pinia";
import {getDataForForm} from "@/lib/request/dataForForm.js";
import {getRoomToken} from "@/lib/request/lesson.js";

export const userStore = defineStore("userStore", {
    state: () => ({
        user: {
          id: -1,
          name: '',
          surname: '',
          email: '',
          role: ''
        },
        token: '',
        roomToken: ''
    })
})

export const formStore = defineStore('formStore', {
    state: () => ({
        formData: {
            languages: null,
            groups: null
        }
    }),
    actions: {
        async getFormData() {
            try {
                const data = await getDataForForm(userStore().token)
                console.log(data)
                this.formData = {
                    languages: data.languages,
                    groups: data.groups
                }
            } catch {
                console.log("Ошибка получения данных")
            }
        }
    }
})

export const roomStore = defineStore('roomStore', {
    state: () => ({
        mapRoomToken: {
            roomName: '',
            roomToken: ''
        }
    }),
    actions: {
        async getRoomToken() {
            try {
                let result = await getRoomToken(userStore().token, this.mapRoomToken.roomName)
                if (result.status === 200) {
                    this.mapRoomToken.roomToken = result.data.token
                }
            } catch {
                throw new Error('ошибка получения токена комнаты')
            }
        }
    }
})