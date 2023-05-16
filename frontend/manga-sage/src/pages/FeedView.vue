<template>
  <FeedLayout>
    <HeroSlider class="drop-shadow-xl" />
    <BestOfView />
    <div
      class="grid grid-flow-row gap-8 text-neutral-600 grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 px-10"
    >
      <div v-if="Manga.length > 0" v-for="m in Manga">
        <MangaCard :manga="m" />
      </div>
      <div v-for="_ in 16" v-else>
        <CardSkeleton />
      </div>
    </div>
    <nav class="flex justify-center items-center space-x-2 mt-5">
      <button
      :disabled="currentPage === 1"
        @click="loadPreviousPage()"
        class="text-gray-500 hover:text-blue-600 p-4 inline-flex items-center gap-2 rounded-md"
      >
        <span aria-hidden="true">«</span>
        <span class="sr-only">Previous</span>
    </button>
      <button
      v-for="page in totalPages"
        @click="loadPage(page)"
        :class="currentPage == page ? 'bg-blue-500' : ''"
        class="w-10 h-10 text-white p-4 inline-flex items-center text-sm font-medium rounded-full cursor-pointer"
        aria-current="page"
        >{{ page }}</button
      >
      <button
      @click="loadNextPage()"
      :disabled="currentPage == totalPages"
      class="text-gray-500 hover:text-blue-600 p-4 inline-flex items-center gap-2 rounded-md"
        >
        <span class="sr-only">Next</span>
        <span aria-hidden="true">»</span>
    </button>
    </nav>
  </FeedLayout>
</template>

<script>
import axios from 'axios'
import HeroSlider from '../components/HeroSlider.vue'
import BestOfView from '../components/BestOfView.vue'
import MangaCard from '../components/MangaCard.vue'
import CardSkeleton from '../components/skeleton/CardSkeleton.vue'
export default {
  components: {
    HeroSlider,
    BestOfView,
    MangaCard,
    CardSkeleton
  },
  data() {
    return {
      Manga: [],
      currentPage: 1,
      totalPages: 0,
      perPage: 16
    }
  },
  methods: {
    fetchManga() {
      axios
        .get(`${this.$hostname}/manga?page=${this.currentPage}`)
        .then((response) => {
          let { data } = response
          this.Manga = data.manga
          this.totalPages = data.totalPages
        })
        .catch((error) => {
          console.error(error)
        })
    },
    loadPage(page) {
      this.currentPage = page
      this.fetchManga()
    },
    loadNextPage() {
      this.currentPage++
      this.fetchManga()
    },
    loadPreviousPage() {
      this.currentPage--
      this.fetchManga()
    }
  },
  created() {
    this.fetchManga()
  }
}
</script>

<style scoped></style>
