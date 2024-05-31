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
                <div
                    class="hover:bg-blue-700 hover:border-white hidden group-hover:block absolute start-[85px] top-[140px] rounded-full border-x-slate-200 p-2 z-20 border-4 text-white bg-blue-700 bg-opacity-50">
                    <p class="mb-0">Detail</p>
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                        class="bi bi-arrow-up-right-square ml-2.5" viewBox="0 0 16 16">
                        <path fill-rule="evenodd"
                            d="M15 2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1zM0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2zm5.854 8.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z" />
                    </svg>
                </div>
                <img class="min-w-[210px] max-h-[315px] group-hover:blur-[1px]" :src="book.imgUrl" alt="">
                <p class="text-lg font-bold mt-2 line-clamp-2">{{ book.title }}</p>
                <p class="text-sm font-bold text-end mb-1">優惠價 NT$
                    <span class="text-sm font-bold text-sky-600">{{ book.price }}</span>
                </p>
            </div>
        </div>
    </div>

    <!-- section3 -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] bg-[#fff6ee]/[.4]">

    </div>
    <div class="snap-x snap-mandatory ...">
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1604999565976-8913ad2ddb7c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1540206351-d6465b3ac5c1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1622890806166-111d7f6c7c97?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1590523277543-a94d2e4eb00b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1575424909138-46b05e5919ec?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
        <div class="snap-always snap-center ...">
            <img
                src="https://images.unsplash.com/photo-1559333086-b0a56225a93c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=320&h=160&q=80" />
        </div>
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