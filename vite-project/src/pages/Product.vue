<template>
    <div class="relative -top-20 blur-lg bg-repeat-x bg-center min-w-[90vh] min-h-[50vh] 
        " :style="{ 'background-image': 'url(' + showStore.bookInfo.cover_image + ')' }">
    </div>
    <div class="container max-w-[90vh] grid grid-rows-2">
        <div
            class="container absolute min-w-[100vh] max-h-max grid grid-cols-3 gap-6 top-60 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
            <div class="justify-self-start">
                <img :src=showStore.bookInfo.cover_image
                    class="max-h-[40vh] mx-auto outline outline-offset-2 outline-1 rounded outline-slate-400/50" />
            </div>
            <div class="rounded-md">
                <div class="text-3xl tracking-wide font-bold mb-4 text-gray-800">{{ showStore.bookInfo.title }}</div>
                <div class="text-lg tracking-wider font-medium">ISBN:{{ showStore.bookInfo.ISBN }}
                </div>
                <div class="text-lg tracking-wider font-medium">EISBN:{{ showStore.bookInfo.ISBN }}
                </div>
                <div class="text-lg tracking-wider font-medium">作者:{{ showStore.bookInfo.writer }}
                </div>
                <div class="text-lg tracking-wider font-medium">出版社:{{
                    showStore.bookInfo.publisher }}</div>
                <div class="text-lg tracking-wider font-medium">
                    出版日期:{{ showStore.bookInfo.publishDate }}</div>


            </div>
            <div class="rounded-md bg-blue-50/50 flex justify-center min-w-[70%] justify-self-end">
                <div class="content-center">
                    <!-- <div class="text-lg tracking-wider font-medium">最大借閱時間:{{
                        showStore.bookInfo.maxRentTime }}
                    </div> -->
                    <div class="text-lg font-medium">剩餘數量:{{ showStore.bookInfo.amount }}
                    </div>
                    <div class="text-2xl font-medium text-sky-600 mb-4"><span class="text-lg">優惠價</span>${{
                        showStore.bookInfo.price }}</div>
                    <button type="button"
                        class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">借閱
                        / 閱讀</button>
                </div>
            </div>
        </div>

        <div
            class="text-sm font-medium text-gray-500 border-b border-gray-200 relative top-20 bg-blue-50/75 rounded-lg">
            <ul class="flex flex-wrap -mb-px" data-tabs-toggle="#default-tab-content" role="tab">
                <li class="me-2">
                    <a href="#" role="tab" data-tabs-target="#detail"
                        class="inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg">詳細內容</a>
                </li>
                <li class="me-2">
                    <a href="#" role="tab" data-tabs-target="#introduction"
                        class="inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg">書本簡介</a>
                </li>
                <li class="me-2">
                    <a href="#" role="tab" data-tabs-target="#chapter"
                        class="inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg">章節</a>
                </li>
            </ul>
            <div id="default-tab-content">
                <div class="hidden p-4 rounded-lg min-h-96" id="detail" role="tabpanel">
                    <div class="rounded-md">
                        <div class="text-lg tracking-wider font-medium">ISBN:{{ showStore.bookInfo.ISBN }}
                        </div>
                        <div class="text-lg tracking-wider font-medium">EISBN:{{ showStore.bookInfo.ISBN }}
                        </div>
                        <div class="text-lg tracking-wider font-medium">作者:{{ showStore.bookInfo.writer }}
                        </div>
                        <div class="text-lg tracking-wider font-medium">出版社:{{
                            showStore.bookInfo.publisher }}</div>
                        <div class="text-lg tracking-wider font-medium">
                            出版日期:{{ showStore.bookInfo.publishDate }}</div>

                        <div class="text-lg tracking-wider font-medium">上傳者:{{ showStore.bookInfo.uploader }}</div>

                        <div class="text-lg tracking-wider font-medium">種類:{{ showStore.bookInfo.class.className }}
                        </div>
                        <div class="text-lg tracking-wider font-medium">年級:{{ showStore.bookInfo.class.grade }}</div>
                        <div class="text-lg tracking-wider font-medium">版本:{{ showStore.bookInfo.edition }}</div>
                        <div class="text-lg tracking-wider font-medium">互動書:{{ showStore.bookInfo.live }}</div>

                    </div>
                </div>
                <div class="hidden p-4 rounded-lg min-h-96" id="introduction" role="tabpanel">
                    <p class="text-lg tracking-wider font-medium">{{
                        showStore.bookInfo.introduction }}</p>
                </div>
                <div class="hidden p-4 rounded-lg min-h-96" id="chapter" role="tabpanel">
                    <p class="text-lg tracking-wider font-medium">{{ showStore.bookInfo.chapter }}
                    </p>
                </div>
            </div>
        </div>
    </div>
</template>
<script setup lang="ts" name="Product">
import { useBookInfoStore } from '../store/book_info.ts'
const showStore = useBookInfoStore()
</script>