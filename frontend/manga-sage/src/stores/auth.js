import { defineStore } from 'pinia'
import axios from 'axios'
export const useAuthStore = defineStore('auth', {
  state: () => ({
    error: null
  }),
  actions: {
    async login(email, password) {
      
      try {
        const response = await axios.post(`${this.$hostname}/login`, {
          email: email,
          password: password
        })
        const { data } = response
        localStorage.setItem('token', data.token)
      } catch (error) {
        if (error?.response?.status == 401) {
          this.error = error.response.data.error
        } else {
          this.error = "Try again later!"
        }
      }
    }
  }
})