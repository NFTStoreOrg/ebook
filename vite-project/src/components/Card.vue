<script setup lang="ts">
// import { defineProps } from 'vue'
import { useCardStore } from '../store/card'
const cardStore = useCardStore()
// defineProps<{ books: Object }>()
</script>
<template>
    <div v-for="book in cardStore.books"
        class="max-w-[226px] p-2 group hover:cursor-pointer hover:shadow-lg transition duration-500 ease hover:-translate-y-2">
        <div
            class="hidden group-hover:block absolute start-[71px] top-[160px] hover:bg-blue-700 hover:border-0 rounded-full border-2 p-2 px-3 z-20">
            <div class="flex text-white">
                <p class="mb-0 text-sm font-semibold tracking-wide">Detail <span aria-hidden="true">&rarr;</span></p>
            </div>
        </div>
        <img class="min-w-[210px] max-h-[315px] group-hover:blur-[1px]" :src="book.imgUrl" alt="">
        <p class="text-lg font-bold mt-2 line-clamp-2">{{ book.title }}</p>
        <p class="text-sm font-bold text-end mb-1">優惠價 NT$
            <span class="text-sm font-bold text-sky-600">{{ book.price }}</span>
        </p>
    </div>
</template>