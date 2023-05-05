<template>
  <div>
    <TheHeader/>
    <router-view></router-view>
    <TheFooter/>
  </div>
</template>

<script>
import TheFooter from './components/TheFooter.vue';
import TheHeader from './components/TheHeader.vue';
export default {
    components:{
    TheHeader,
    TheFooter
}
  }
</script>
