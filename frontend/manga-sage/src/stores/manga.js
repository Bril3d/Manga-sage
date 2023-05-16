import { defineStore } from 'pinia'
import { getCurrentInstance } from 'vue'
import axios from 'axios'
export const useMangaStore = defineStore('manga', {
  state: () => ({
    error: null,
    Manga: null
  }),
  actions: {
    async MangaByID(id) {
      try {
        const response = await axios.get(`${this.$hostname}/manga/${id}`)
        const { data } = response
        this.Manga = data.manga
      } catch (error) {
        if (error?.response?.status == 401) {
          this.error = error.response.data.error
        } else {
          this.error = 'Try again later!'
        }
      }
    }
  },
  getters: {
    getManga: (state) => state.getManga
  }
})
