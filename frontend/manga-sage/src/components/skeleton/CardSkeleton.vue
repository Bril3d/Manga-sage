<template>
<div
      class="flex flex-col bg-white border shadow-sm rounded-xl dark:bg-gray-800 dark:border-gray-700 dark:shadow-slate-700/[.7] "
    ><div class="overflow-hidden">
      <div
        class="w-full h-64 rounded-t-xl animate-pulse bg-gray-200 dark:bg-gray-700 overflow-hidden"
      ></div>
    </div>
      <div class="p-4 md:p-5">
        <h3 class="animate-pulse"></h3>
        <span class="w-full h-4 block bg-gray-200 rounded-full dark:bg-gray-700"></span>
        <div class="flex flex-col gap-2 mt-3">
        <div v-for="_ in 2"
          class="h-8 w-full inline-flex justify-center items-center gap-2 animate-pulse bg-gray-200 rounded-full dark:bg-gray-700"
        >
        </div>
      </div>
      </div>
    </div>
</template>
