<script setup lang="ts">
import Card from '../components/Card.vue'
import { ref } from 'vue'
defineProps<{ title: string, slideBgColor: string }>()
const content = ref()
const leftArrow = ref(false)
const rightArrow = ref(true)

function slideR() {
    leftArrow.value = true
    const viewWidth = content.value.clientWidth
    const scrollWidth = content.value.scrollWidth
    const maxScroll = viewWidth * Math.floor(scrollWidth / viewWidth - 1)
    if (content.value.scrollLeft >= maxScroll) {
        rightArrow.value = false
    }
    content.value.scrollLeft += viewWidth
}
function slideL() {
    rightArrow.value = true
    const viewWidth = content.value.clientWidth
    console.log(content.value.scrollLeft - viewWidth)
    if (content.value.scrollLeft - viewWidth < 100) {
        leftArrow.value = false
    }
    content.value.scrollLeft -= viewWidth
}
</script>
<template>
    <div class="row mx-auto text-2xl max-w-[95%] font-bold pl-3 pt-3 tracking-wider"><span
            class="bg-blue-100 text-blue-800 me-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300">{{ title
            }}</span>
    </div>
    <div :class=slideBgColor class="row mx-auto rounded-3xl m-3 max-h-[80vh] max-w-[95%] bg-opacity-50">
        <button type="button" @click="slideL" v-show="leftArrow"
            class="absolute text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center me-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M5 12h14M5 12l4-4m-4 4 4 4" />
            </svg>

            <span class="sr-only">Icon description</span>
        </button>
        <div ref="content"
            class="scroll-smooth flex snap-x snap-mandatory items-center gap-x-3 overflow-x-auto overflow-x-hidden p-4">
            <card></card>
        </div>
        <button type="button" @click="slideR" v-show="rightArrow"
            class="absolute end-0 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 text-center inline-flex items-center me-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
            <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M1 5h12m0 0L9 1m4 4L9 9" />
            </svg>
            <span class="sr-only">Icon description</span>
        </button>
    </div>

</template>