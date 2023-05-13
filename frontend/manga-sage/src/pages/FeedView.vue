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
      <CardSkeleton/>
    </div>
    </div>
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
      Manga: []
    }
  },
  methods: {
    async fetchManga() {
      axios.get('http://localhost:3000/manga').then((response) => {
        let { data } = response
        this.Manga = data.manga
      })
    }
  },
  mounted() {
    this.fetchManga()
  }
}
</script>

<style scoped></style>
