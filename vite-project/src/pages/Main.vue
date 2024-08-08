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
            <Card cardWidth2="max-w-[226px]"></Card>
        </div>
    </div>

    <!-- section3 -->
    <CardSlide title="教科書" slideBgColor="bg-orange-200" cardWidth1="w-[150px]"></CardSlide>
    <CardSlide title="參考書" slideBgColor="bg-red-100" cardWidth1="w-[150px]"></CardSlide>
    <!-- <div class="container h-[80vh] min-w-[95%] bg-violet-300/25 rounded-3xl text-9xl content-center fontfont-bold">
        <div class="flex">
            <div :style="ebook" id="ebook" style="">E</div>
            <div>B</div>
            <div>O</div>
            <div>O</div>
            <div>K</div>
        </div>
        <div>STORE</div>
    </div> -->
    <!-- <CardSlide title="兒童書籍" slideBgColor="bg-sky-200" cardWidth1="w-[150px]"></CardSlide>
    <CardSlide title="互動書" slideBgColor="bg-violet-100" cardWidth1="w-[150px]"></CardSlide>
    <CardSlide title="影片" slideBgColor="bg-violet-100" cardWidth1="w-[150px]"></CardSlide>
    <CardSlide title="其他" slideBgColor="bg-violet-100" cardWidth1="w-[150px]"></CardSlide> -->


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

        <div id="default-tab-content" class="h-[80vh] bg-violet-300/25 fontfont-bold rounded-3xl content-center">
            <div class="hidden p-4 rounded-lg h-full" id="profile" role="tabpanel">
                <div class="grid grid-cols-5 h-full overflow-hidden">
                    <div class="font-['YouSheBiaoTiHei'] self-center text-cyan-700 text-9xl flex-col mx-auto">
                        童<br>趣<br>
                        <button type="button"
                            class="py-2.5 px-3 me-2 mb-2 text-sm font-medium bg-white hover:!bg-gray-100 text-gray-900 focus:outline-none rounded-full border border-gray-200 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Discover
                            <span aria-hidden="true">&rarr;</span></button>
                    </div>

                    <div class="col-span-4 self-center">
                        <CardSlide title="" slideBgColor="" cardWidth1="h-[300px]"></CardSlide>
                    </div>
                </div>
            </div>
            <div class="hidden p-4 rounded-lg" id="dashboard" role="tabpanel">
                互動書
            </div>
            <div class="hidden p-4 rounded-lg" id="dashboar" role="tabpanel">
                影片
            </div>
            <div class="hidden p-4 rounded-lg" id="dashboa" role="tabpanel">
                其他
            </div>
        </div>
    </div>
</template>

<script setup lang="ts" name="Main">
import { ref } from 'vue'
import Card from '../components/Card.vue'
import CardSlide from '../components/CardSlide.vue'
const ebook = ref()
import("https://chinese-fonts-cdn.deno.dev/packages/ysbth/dist/优设标题黑/result.css")
window.addEventListener("scroll", handleScroll);
function handleScroll() {
    // console.log(window.scrollY)

    const ebookElement = document.getElementById('ebook') as HTMLElement;

    if (window.scrollY >= 3516) {
        // ebook.value = "transform: translate(50px, -100px);"
    }

    else {
        // ebook.value = "transform: translate(0px, 0px);"
    }
    const scrollValue = window.scrollY;

    // ebookElement.style.setProperty('--scroll', scrollValue.toString());

    // 设置动画播放状态
    // if (window.scrollY >= 3516) {
    //     ebookElement.style.animationPlayState = 'paused';
    // } else {
    //     ebookElement.style.animationPlayState = 'paused';
    // }

}
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