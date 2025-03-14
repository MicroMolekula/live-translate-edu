import {defineStore} from "pinia";

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