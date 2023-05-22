<template>
  <FeedLayout>
    <div class="py-5 text-center font-bold text-3xl text-white">
      <h1>{{ getMangaName }}</h1>
    </div>
    <div class="flex flex-col">
      <img
        class="mx-auto"
        v-for="page in Pages"
        :src="`/src/assets/manga/${page.Image}`"
        alt="Manga Page"
      />
    </div>
  </FeedLayout>
</template>

<script>
import axios from 'axios'
export default {
  data() {
    return {
      Pages: null
    }
  },
  computed: {
    getMangaName() {
      if(this.Pages){
      const regex = /\/([^/]+)\//
      const string = this.Pages[0]?.Image
      const match = string.match(regex)
      const result = match ? match[1] : null
      return result
    }
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
  created() {
    this.fetchPages()
  }
}
</script>

<style scoped></style>
