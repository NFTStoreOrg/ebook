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
            <Card :cardStore=cardStore.newBooks></Card>
        </div>
    </div>

    <!-- section3 -->
    <CardSlide title="教科書" slideBgColor="bg-orange-200" :cardStore=cardStore.textbook></CardSlide>
    <CardSlide title="參考書" slideBgColor="bg-red-100" :cardStore=cardStore.referenceBook></CardSlide>
    <CardSlide title="兒童書籍" slideBgColor="bg-sky-200" :cardStore=cardStore.childrenBook></CardSlide>
    <CardSlide title="互動書" slideBgColor="bg-violet-100" :cardStore=cardStore.live></CardSlide>
    <CardSlide title="影片" slideBgColor="bg-violet-100" :cardStore=cardStore.video></CardSlide>
    <CardSlide title="其他" slideBgColor="bg-violet-100" :cardStore=cardStore.otherBook></CardSlide>

</template>
<script setup lang="ts" name="Main">
// import { ref } from 'vue'
import Card from '../components/Card.vue'
import CardSlide from '../components/CardSlide.vue'
import { useCardStore } from '../store/card'
const cardStore = useCardStore()
cardStore.getNewestBook()
cardStore.getClassBook("textbook")
cardStore.getClassBook("children")
cardStore.getClassBook("video")
cardStore.getClassBook("reference")
cardStore.getClassBook("other")
cardStore.getLiveBook()
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