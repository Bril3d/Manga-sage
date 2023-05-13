<template>
    <div
      class="flex flex-col bg-white border shadow-sm rounded-xl dark:bg-gray-800 dark:border-gray-700 dark:shadow-slate-700/[.7] "
    ><a class="overflow-hidden" :href="`/series/${manga.ID}`">
      <img
        class="w-full h-64 rounded-t-xl object-cover hover:scale-150 transition-all ease-in overflow-hidden"
        :src="`/src/assets/manga/${manga.Cover_Image}`"
        :alt="manga.Title"
      />
    </a>
      <div class="p-4 md:p-5">
        <h3 dir="ltr" class="text-lg font-bold text-gray-800 dark:text-white truncate">{{ manga.Title }}</h3>
        <div class="flex flex-col gap-2 mt-3">
        <a v-for="chapter in manga.Chapters"
          class="py-3 px-4 inline-flex justify-center items-center rounded-md border border-transparent font-semibold bg-blue-500 text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all text-sm dark:focus:ring-offset-gray-800"
          href="#"
        >
          {{ chapter.Number }}
        </a>
      </div>
      </div>
    </div>
</template>

<script>
export default {
  props: {
    manga: Object
  }
}
</script>

<style scoped></style>
