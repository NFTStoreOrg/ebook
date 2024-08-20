<script setup lang="ts">
import Card from '../components/Card.vue'
import { onMounted, ref } from 'vue'
const content = ref()
const leftArrow = ref(false)
const rightArrow = ref(true)

onMounted(() => {
    if (content.value instanceof HTMLElement) {
        const clientWidth = content.value.clientWidth
        const maxScrollWidth = content.value.scrollWidth
        console.log(clientWidth,maxScrollWidth)
        if (maxScrollWidth < clientWidth) {
            rightArrow.value = false
        }
    }
});

function slideR() {
    leftArrow.value = true
    const clientWidth = content.value.clientWidth
    const maxScrollWidth = content.value.scrollWidth
    console.log(clientWidth, maxScrollWidth, content.value.scrollLeft)
    if ((maxScrollWidth - content.value.scrollLeft) - clientWidth <= clientWidth) {
        content.value.scrollLeft = (maxScrollWidth - clientWidth)
        rightArrow.value = false

    } else {
        content.value.scrollLeft += clientWidth
    }
}
function slideL() {
    rightArrow.value = true
    const clientWidth = content.value.clientWidth
    console.log(clientWidth, content.value.scrollLeft)
    if (content.value.scrollLeft < clientWidth) {
        leftArrow.value = false
        content.value.scrollLeft = 0
    } else {
        content.value.scrollLeft -= clientWidth
    }
}

defineProps<{ title: string, slideBgColor: string, cardStore: { title: string; imgUrl: string; price: number; bookId: number; }[], cardWidth1: string }>()
</script>
<template>
    <div class="row mx-auto text-2xl max-w-[95%] font-bold pl-3 pt-3 tracking-wider">
        <span class="bg-blue-100 text-blue-800 me-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300">
            {{ title }}
        </span>
    </div>
    <div :class=slideBgColor class="relative row mx-auto rounded-3xl m-3 max-h-[80vh] max-w-[95%] bg-opacity-50">
        <button type="button" @click="slideL" v-show="leftArrow" class="absolute z-50 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 
            focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm 
            p-2.5 text-center inline-flex items-center me-2 top-[50%] -translate-y-[50%]">
            <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M1 5h23M1 5l4-4m-4 4 4 4" />
            </svg>
        </button>
        <div ref="content"
            class="scroll-smooth flex snap-x snap-mandatory items-center gap-x-3 overflow-x-auto overflow-x-hidden p-4"
            style="mask-image: linear-gradient(to left, transparent, black 80px, black calc(100% - 80px), transparent);">
            <Card :cardStore="cardStore" :card-width2="cardWidth1"></Card>
        </div>
        <button type="button" @click="slideR" v-show="rightArrow" class="absolute end-0 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 
            focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm p-2.5 
            text-center inline-flex items-center me-2 top-[50%] -translate-y-[50%]">
            <svg class="w-4 h-4" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                    d="M1 5h12m0 0L9 1m4 4L9 9" />
            </svg>
        </button>
    </div>
</template>
