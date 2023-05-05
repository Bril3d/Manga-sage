<template>
  <FeedLayout>
    <HeroSlider class="drop-shadow-xl" />
    <BestOfView />
  </FeedLayout>
</template>

<script>
import HeroSlider from '../components/HeroSlider.vue'
import BestOfView from '../components/BestOfView.vue'
export default {
  components: {
    HeroSlider,
    BestOfView,
  }
}
</script>

<style scoped></style>
