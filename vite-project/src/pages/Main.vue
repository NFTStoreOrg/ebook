<template>
    <!-- section1 -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[85vh] max-w-[95%] bg-[#ECF0FC] items-center">
        <div
            class="col rounded-3xl m-3 max-h-[80vh] max-w-[50%] bg-white flex items-center justify-center overflow-hidden">

            <h1 class="marquee font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                EBOOK STORE
            </h1>
            <h1 class="marquee font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                &nbsp;EBOOK STORE
            </h1>
            <h1 class="marquee font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                &nbsp;EBOOK STORE
            </h1>


        </div>
        <div class="col rounded-3xl m-3 max-w-[50%] flex items-center justify-center overflow-hidden">
            <div class="rounded-3xl max-h-[65vh] flex items-center justify-center overflow-hidden">
                <video loop autoplay muted>
                    <source src="../assets/book.mp4" type="video/mp4">
                </video>
            </div>
        </div>
    </div>
    <!-- section2 -->
    <div class="flex outline outline-1 outline-[#16161a14] row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] p-5">
        <p class="block mb-0 text-3xl font-extrabold text-blue-700 text-start">New Book Releases
        </p>
        <hr class="block mb-0 w-full" />
        <div class="grid grid-cols-1 gap-8 mt-3 lg:grid-cols-4 md:grid-cols-3 xl:grid-cols-6 mx-auto">
            <Card></Card>
            <!-- <div v-for="book in books"
                class="max-w-[226px] p-2 group hover:cursor-pointer hover:shadow-lg transition duration-500 ease hover:-translate-y-2">
                <div
                    class="hidden group-hover:block absolute start-[71px] top-[160px] hover:bg-blue-700 hover:border-0 rounded-full border-2 p-2 px-3 z-20">
                    <div class="flex text-white">
                        <p class="mb-0 text-sm font-semibold tracking-wide">Detail <span
                                aria-hidden="true">&rarr;</span></p>
                    </div>
                </div>
                <img class="min-w-[210px] max-h-[315px] group-hover:blur-[1px]" :src="book.imgUrl" alt="">
                <p class="text-lg font-bold mt-2 line-clamp-2">{{ book.title }}</p>
                <p class="text-sm font-bold text-end mb-1">優惠價 NT$
                    <span class="text-sm font-bold text-sky-600">{{ book.price }}</span>
                </p>
            </div> -->
        </div>
    </div>

    <!-- section3 -->
    <div class="row mx-auto text-3xl max-w-[95%]">參考書</div>
    <div class="row mx-auto rounded-3xl m-3 max-h-[80vh] max-w-[95%] bg-orange-200 bg-opacity-50">
        <div class="flex snap-x snap-mandatory items-center gap-x-3 overflow-x-auto p-4">
            <card></card>
        </div>
    </div>
    <!-- <div class="row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] bg-rose-100 bg-opacity-50">

    </div> -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] bg-sky-200 bg-opacity-50">

    </div>

</template>
<script setup lang="ts" name="Main">
// import { ref } from 'vue'
import Card from '../components/Card.vue'

</script>
<style scoped>
.marquee {
    animation: slide 10s linear infinite;
    font-size: 500px;
    -webkit-text-stroke: 2px #5988ff;
    color: transparent;
}

@keyframes slide {
    0% {
        transform: translatex(0)
    }

    100% {
        transform: translatex(-100%)
    }
}
</style>