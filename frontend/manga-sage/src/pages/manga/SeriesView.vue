<template>
  <FeedLayout>
    <div v-if="mangaStore.Manga" class="py-14 px-10 grid grid-cols-3 grid-rows-4 gap-5">
      <div
        class="flex flex-col justify-center items-center gap-5 dark:shadow-slate-700/[.7] shadow-sm rounded-xl pb-5"
      >
        <img
          class="w-full h-96 rounded-t-xl object-cover overflow-hidden"
          :src="`/src/assets/manga/${mangaStore.Manga.Cover_Image}`"
        />
        <h3 dir="ltr" class="text-xl font-bold text-gray-800 dark:text-white truncate">
          {{ mangaStore.Manga.Title }}
        </h3>
      </div>
      <div class="dark:shadow-slate-700/[.7] shadow-sm rounded-xl px-10 text-white col-span-2 flex flex-col gap-4">
        <div>
        <h3 class="text-gray-800 font-semibold text-lg dark:text-white">القصة :</h3>
            {{ mangaStore.Manga.Description }}
          </div>
          <div class="dark:shadow-slate-700/[.7] shadow-sm rounded-xl overflow-y-scroll h-44 flex gap-3 flex-wrap px-4 py-2 scrollbar-thin scrollbar-thumb-blue-700 scrollbar-track-blue-300">
            <button @click="goToPages(chapter.Number)" v-for="chapter in Chapters" type="button" class="py-3 px-4 h-fit inline-flex justify-center items-center gap-2 rounded-md border-2 border-gray-900 font-semibold text-gray-800 hover:text-white hover:bg-gray-800 hover:border-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-800 focus:ring-offset-2 transition-all text-sm dark:hover:bg-gray-900 dark:border-gray-900 dark:hover:border-gray-900 dark:text-white dark:focus:ring-gray-900 dark:focus:ring-offset-gray-800">
  الفصل {{ chapter.Number }}
</button>
          </div>
      </div>
    </div>
  </FeedLayout>
</template>

<script>
import axios from "axios"
import { useMangaStore } from '../../stores/manga'
export default {
  data() {
    return {
      mangaStore:  useMangaStore(),
      Chapters: [],
      id: this.$route.params.id,
    }
  },
  methods: {
    fetchManga() {
      this.mangaStore.MangaByID(this.id)
      
    },
    getChapters(){
      axios.get(`http://localhost:3000/manga/chapters/${this.id}`).then((response) => {
        let { data } = response
        this.Chapters = data.chapters
      })
    },
    goToPages(ChapterNumber){
      this.$router.push({name:"chapter", params: { id: this.id, chapter:ChapterNumber}})
    }
  },
  mounted() {
    this.fetchManga()
    this.getChapters()
  }
}
</script>

<style scoped></style>
