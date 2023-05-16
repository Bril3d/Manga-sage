<template>
  <FeedLayout>
    <div class="py-5 text-center font-bold text-3xl text-white">
      <h1>{{ mangaStore.Manga.Title }}</h1>
    </div>
    <div class="flex flex-col">
      <img class="mx-auto" v-for="page in Pages" :src="`/src/assets/manga/${page.Image}`" alt="Manga Page" />
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
      Pages:[]
    }
  },
  methods: {
    fetchPages() {
      let id = this.$route.params.id
      let chapter = this.$route.params.chapter
      axios.get(`${this.$hostname}/manga/${id}/chapter/${chapter}`).then((response) => {
        let { data } = response
        this.Pages = data
      })
    }
  },
  mounted() {
    this.fetchPages()
  }
}
</script>

<style scoped></style>
