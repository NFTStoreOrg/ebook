<template>
    <!-- section1 -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[85vh] max-w-[95%] bg-[#ECF0FC] items-center">
        <div
            class="col rounded-3xl m-3 max-h-[80vh] max-w-[50%] bg-white flex items-center justify-center overflow-hidden">

            <h1 class="marquee font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                EBOOK STORE
            </h1>
            <h1 class="marquee2 font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                &nbsp;EBOOK STORE
            </h1>
            <h1 class="marquee3 font-bold text-[#5988ff] font-mono whitespace-nowrap inline-block">
                &nbsp;EBOOK STORE
            </h1>


        </div>
        <div
            class="col rounded-3xl m-3 max-w-[50%] flex items-center justify-center overflow-hidden">
            <div class="rounded-3xl max-h-[65vh] flex items-center justify-center overflow-hidden">
                <video loop autoplay muted>
                    <source src="../assets/book.mp4" type="video/mp4">
                </video>
            </div>
        </div>
    </div>
    <!-- section2 -->
    <div
        class="flex outline outline-1 outline-[#16161a14] row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] p-5">
        <p class="block mb-0 ml-3 text-3xl font-extrabold text-blue-700 text-start block">New Book Releases</p>
        <hr class="block mb-0 ml-0 w-full"/>
         <summary class="card mb-3" style="border-color: transparent;background-color: transparent">
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
        </summary>
    </div>

    <!-- section3 -->
    <div class="row mx-auto rounded-3xl m-3 min-h-[80vh] max-w-[95%] bg-[#fff6ee]/[.4]">

    </div>
</template>
<script setup lang="ts">
</script>
<style scoped>
.marquee,
    .marquee2,
    .marquee3 {
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
/* 卡片 */
#card_group {
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
}
@include media-breakpoint-up(xxl) { 
    #card_group {
        grid-template-columns: repeat(4, minmax(200px, 1fr));
    }
}
</style>