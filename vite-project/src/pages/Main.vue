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
            <Card :cardStore=cardStore.newBooks cardWidth2="max-w-[226px]"></Card>
        </div>
    </div>

    <!-- section3 -->
    <CardSlide title="教科書" slideBgColor="bg-orange-200" cardWidth1="w-[150px]" :cardStore=cardStore.textbook></CardSlide>
    <CardSlide title="參考書" slideBgColor="bg-red-100" cardWidth1="w-[150px]" :cardStore=cardStore.referenceBook></CardSlide>

    <div class="container min-w-[95%] mb-8 mt-14">

        <ul class="flex flex-wrap text-sm font-medium text-center text-lg" id="default-tab"
            data-tabs-active-classes="bg-purple-600 text-white border-purple-600 rounded"
            data-tabs-inactive-classes="text-gray-500 hover:text-gray-600 border-gray-100 hover:border-gray-300"
            data-tabs-toggle="#default-tab-content" role="tab">
            <li>
                <button id="profile-tab" data-tabs-target="#profile" type="button" role="tab">兒童書籍</button>
            </li>
            <li>
                <button id="dashboard-tab" data-tabs-target="#dashboard" type="button" role="tab">互動書</button>
            </li>
            <li>
                <button id="dashboard-tab" data-tabs-target="#dashboar" type="button" role="tab">影片</button>
            </li>
            <li>
                <button id="dashboard-tab" data-tabs-target="#dashboa" type="button" role="tab">其他</button>
            </li>
        </ul>

        <div id="default-tab-content" class="h-[80vh] bg-violet-300/25 rounded-3xl content-center">
            <div class="hidden p-4 rounded-lg" id="profile" role="tabpanel">
                <div class="grid grid-cols-5 overflow-hidden">
                    <div
                        class="[text-shadow:_5px_5px_rgb(255_0_0_/_40%)] font-['YouSheBiaoTiHei'] self-center text-cyan-500 text-9xl flex-col mx-auto">
                        童<br>趣<br>
                        <button type="button"
                            class="[box-shadow:_5px_5px_rgb(255_0_0_/_40%)] py-2.5 px-3 me-2 mb-2 text-sm font-medium bg-white hover:!bg-gray-100 text-gray-900 focus:outline-none rounded-full border border-gray-200 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Discover
                            <span aria-hidden="true">&rarr;</span></button>
                    </div>

                    <div class="col-span-4 self-center">
                        <CardSlide title="" slideBgColor="" cardWidth1="" :cardStore=cardStore.childrenBook></CardSlide>
                    </div>
                </div>
            </div>
            <div class="hidden p-4 rounded-lg" id="dashboard" role="tabpanel">
                <div class="grid grid-cols-5 overflow-hidden">
                    <div
                        class="font-['MaokenAssortedSans'] self-center text-blue-600 [text-shadow:_5px_5px_rgb(129_140_248_/_40%)] [-webkit-text-stroke:_1px_white] text-8xl flex-col mx-auto">
                        互動書<br>
                        <button type="button"
                            class="font-['YouSheBiaoTiHei'] [box-shadow:_5px_5px_rgb(129_140_248_/_40%)] [-webkit-text-stroke:0px] py-2.5 px-3 me-2 mb-2 text-sm font-medium bg-white hover:!bg-gray-100 text-gray-900 focus:outline-none rounded-full border border-gray-200 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Discover
                            <span aria-hidden="true">&rarr;</span></button>
                    </div>

                    <div class="col-span-4 self-center">
                        <CardSlide title="" slideBgColor="" cardWidth1="" :cardStore=cardStore.textbook></CardSlide>
                    </div>
                </div>

            </div>
            <div class="hidden p-4 rounded-lg" id="dashboar" role="tabpanel">
                <div class="grid grid-cols-5 overflow-hidden">
                    <div
                        class="font-['YouSheBiaoTiHei'] self-center text-gray-950 [text-shadow:_5px_5px_rgb(129_140_248_/_40%)] [-webkit-text-stroke:_1px_white] text-9xl flex-col mx-auto">
                        影<br>片<br>
                        <button type="button"
                            class="font-['YouSheBiaoTiHei'] [box-shadow:_5px_5px_rgb(129_140_248_/_40%)] [-webkit-text-stroke:0px] py-2.5 px-3 me-2 mb-2 text-sm font-medium bg-white hover:!bg-gray-100 text-gray-900 focus:outline-none rounded-full border border-gray-200 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Discover
                            <span aria-hidden="true">&rarr;</span></button>
                    </div>

                    <div class="col-span-4 self-center">
                        <CardSlide title="" slideBgColor="" cardWidth1="" :cardStore=cardStore.video></CardSlide>
                    </div>
                </div>
            </div>
            <div class="hidden p-4 rounded-lg" id="dashboa" role="tabpanel">
                <div class="grid grid-cols-5 overflow-hidden">
                    <div
                        class="font-['YouSheBiaoTiHei'] self-center text-blue-700 [text-shadow:_5px_5px_rgb(129_140_248_/_40%)] text-9xl flex-col mx-auto">
                        其<br>他<br>
                        <button type="button"
                            class="font-['YouSheBiaoTiHei'] [box-shadow:_5px_5px_rgb(129_140_248_/_40%)] [-webkit-text-stroke:0px] py-2.5 px-3 me-2 mb-2 text-sm font-medium bg-white hover:!bg-gray-100 text-gray-900 focus:outline-none rounded-full border border-gray-200 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Discover
                            <span aria-hidden="true">&rarr;</span></button>
                    </div>

                    <div class="col-span-4 self-center">
                        <CardSlide title="" slideBgColor="" cardWidth1="" :cardStore=cardStore.otherBook></CardSlide>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts" name="Main">
import { ref, onMounted } from 'vue'
import Card from '../components/Card.vue'
import CardSlide from '../components/CardSlide.vue'
import { useCardStore } from '../store/card'
window.addEventListener("scroll", handleScroll);
function handleScroll() {
    // console.log(window.scrollY)

    const ebookElement = document.getElementById('ebook') as HTMLElement;

    if (window.scrollY >= 3516) {
        // ebook.value = "transform: translate(50px, -100px);"
    }

    else {
    }

}
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

#ebook {
    animation: text-reveal 1s linear infinite;
    animation-play-state: paused;
    animation-delay: calc(var(--scroll) * -1s);
    animation-fill-mode: both;
}

@keyframes text-reveal {
    /* 0% {
       transform: translate(0px, 0px); 
    } */

    to {
        transform: translate(50px, -100px);
    }
}
</style>