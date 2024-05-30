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
            <div v-for="book in books"
                class="max-w-[226px] p-2 group hover:cursor-pointer hover:shadow-lg transition duration-500 ease hover:-translate-y-2">
                <div class="group-hover:blur-[1px]">
                    <img class="min-w-[210px] max-h-[315px]" :src="book.imgUrl" alt="">
                </div>
                <p class="text-lg font-bold mt-2 line-clamp-2">{{ book.title }}</p>
                <p class="text-sm font-bold text-end mb-1">優惠價 NT$
                    <span class="text-sm font-bold text-sky-600">{{ book.price }}</span>
                </p>

            </div>
        </div>
        <!-- <summary class="card mb-3" style="border-color: transparent;background-color: transparent">
            <div class="card-group row" id="card_group" style="display: grid;gap:12px;">
                <div v-for="n in 10" class="col p-0 hover:shadow-lg transition duration-500 ease hover:-translate-y-2">
                    <div class="hover:border-indigo-600 hover:border-3 card w-100 h-100"
                        onclick="location.href = '/pages/product_info.html?tokenId={{card.tokenId}}'">
                        <img src="https://images.pexels.com/photos/762687/pexels-photo-762687.jpeg?auto=compress&cs=tinysrgb"
                            class="card-img-top">
                        <div class="card-body row">
                            <h6 class="card-title">
                                電子書
                            </h6>
                            <div class="col">
                                <h7 class="card-text">
                                    電子書簡介
                                </h7>
                            </div>

                            <div class="col text-end">
                                價格
                            </div>

                        </div>
                    </div>
                </div>
            </div>
        </summary> -->
    </div>

    <!-- section3 -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] bg-[#fff6ee]/[.4]">

    </div>
</template>
<script setup lang="ts" name="Main">
import { ref } from 'vue'
let books = ref([
    {
        title: "追尋失落的時光：一個旅人的心靈之旅",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/50IgnOZ10MZ/cf9ca0abf2a77c311fc44d9f7ad1a99e/344/1648650462.png",
        price: 100
    },
    {
        title: "探索無盡的宇宙：星際航行的奇妙之旅",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/TUjmSfBSBs2/3e59d5989f9eceaca8625b6bf780c2e4/344/1648665973.png",
        price: 100
    },
    {
        title: "破曉之前：夢幻世界的神秘啟示",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
        price: 100
    },
    {
        title: "靈魂的奇跡：生命中的意義與使命",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
        price: 100
    },
    {
        title: "時光的記憶：遺忘與回憶的交錯",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/nMzY6NJ97Xs/0d40507679ea0528a8cbc080bf0a74e6/344/1648657152.png",
        price: 100
    },
    {
        title: "破曉之前：夢幻世界的神秘啟示",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
        price: 100
    },
    {
        title: "靈魂的奇跡：生命中的意義與使命",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
        price: 100
    },
    {
        title: "追尋失落的時光：一個旅人的心靈之旅",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/50IgnOZ10MZ/cf9ca0abf2a77c311fc44d9f7ad1a99e/344/1648650462.png",
        price: 100
    },
    {
        title: "探索無盡的宇宙：星際航行的奇妙之旅",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/TUjmSfBSBs2/3e59d5989f9eceaca8625b6bf780c2e4/344/1648665973.png",
        price: 100
    },
    {
        title: "破曉之前：夢幻世界的神秘啟示",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
        price: 100
    },
    {
        title: "靈魂的奇跡：生命中的意義與使命",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
        price: 100
    },
    {
        title: "時光的記憶：遺忘與回憶的交錯",
        imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/nMzY6NJ97Xs/0d40507679ea0528a8cbc080bf0a74e6/344/1648657152.png",
        price: 100
    }
])

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