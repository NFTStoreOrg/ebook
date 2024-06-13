<script setup lang="ts">
import { useCardStore } from '../store/card'
const cardStore = useCardStore()
cardStore.getNewestBook()
</script>
<template>
    <div v-for="book in cardStore.newBooks"
        class="max-w-[226px] min-w-[150px] p-2 group hover:cursor-pointer hover:shadow-lg transition duration-500 ease hover:-translate-y-2 snap-center">
        <div class="relative">
            <img class="group-hover:blur-[3px]" :src="book.imgUrl" alt="">
            <div
                class="hidden group-hover:block absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 hover:bg-gradient-to-r hover:from-cyan-500 hover:to-blue-700 hover:border-0 rounded-full border-2 p-2 px-3 z-20 whitespace-nowrap">
                <p class="mb-0 text-sm font-semibold tracking-wide text-white">
                    Detail
                    <span aria-hidden="true">&rarr;</span>
                </p>
            </div>
        </div>
        <p class="text-lg font-bold mt-2 line-clamp-2">{{ book.title }}</p>
        <p class="text-sm font-bold text-end mb-1">優惠價 NT$
            <span class="text-sm font-bold text-sky-600">{{ book.price }}</span>
        </p>
    </div>
</template>