<template>
  <FeedLayout>
    <div class="py-14 px-10 grid grid-cols-3 grid-rows-4 gap-5">
      <div
        class="flex flex-col justify-center items-center gap-5 dark:shadow-slate-700/[.7] shadow-sm rounded-xl pb-5"
      >
        <img
          class="w-full h-96 rounded-t-xl object-cover overflow-hidden"
          :src="`/src/assets/manga/${Serie.Cover_Image}`"
        />
        <h3 dir="ltr" class="text-xl font-bold text-gray-800 dark:text-white truncate">
          {{ Serie.Title }}
        </h3>
      </div>
      <div class="dark:shadow-slate-700/[.7] shadow-sm rounded-xl px-10 text-white col-span-2">
        <h3 class="text-gray-800 font-semibold text-lg dark:text-white">القصة :</h3>
            {{ Serie.Description }}
      </div>
    </div>
  </FeedLayout>
</template>

<script>
import axios from "axios"
export default {
  data() {
    return {
      Serie: [],
      Chapters: []
    }
  },
  methods: {
    fetchManga() {
      let id = this.$route.params.id
      axios.get(`http://localhost:3000/manga/${id}`).then((response) => {
        let { data } = response
        this.Serie = data.manga
      })
    },
    getChapters(){
      let id = this.$route.params.id
      axios.get(`http://localhost:3000/manga/chapters/${id}`).then((response) => {
        let { data } = response
        this.Chapters = data.chapters
      })
    }
  },
  mounted() {
    this.fetchManga()
    this.getChapters()
  }
}
</script>

<style scoped></style>
