<template>
    <div class="container max-w-[90vh] max-h-[85vh]">
        <div class="grid grid-cols-2">
            <div class=""><img :src=showStore.bookInfo.cover_image
                    class=" max-h-[40vh] mx-auto outline outline-offset-2 outline-1 rounded outline-slate-400/50" />
            </div>
            <div
                class="rounded-md px-4 py-3 max-h-[40vh] outline outline-offset-2 outline-1 rounded outline-slate-400/50">
                <div class="text-xl font-bold">{{ showStore.bookInfo.title }}</div>
                <div><span class="text-gray-600 font-bold me-1 mt-3">作者:</span>{{ showStore.bookInfo.writer }}
                </div>
                <div><span class="text-gray-600 font-bold me-1">出版社:</span>{{ showStore.bookInfo.publisher }}</div>
                <div><span class="text-gray-600 font-bold me-1">
                        出版日期:</span>{{ showStore.bookInfo.publishDate }}</div>
                <div><span class="text-gray-600 font-bold me-1">ISBN:</span>{{ showStore.bookInfo.ISBN }}
                </div>
                <div><span class="text-gray-600 font-bold me-1">上傳者:</span>{{ showStore.bookInfo.uploader }}</div>
                <div><span class="text-gray-600 font-bold me-1">最大借閱時間:</span>{{ showStore.bookInfo.maxRentTime }}
                </div>
                <div><span class="text-gray-600 font-bold me-1">數量:</span>{{ showStore.bookInfo.uploader }}</div>
                <div><span class="text-gray-600 font-bold me-1">種類:</span>{{ showStore.bookInfo.class.className }}</div>
                <div><span class="text-gray-600 font-bold me-1">年級:</span>{{ showStore.bookInfo.class.grade }}</div>
                <div><span class="text-gray-600 font-bold me-1">版本:</span>{{ showStore.bookInfo.edition }}</div>
                <div><span class="text-gray-600 font-bold me-1">互動書:</span>{{ showStore.bookInfo.live }}</div>
                <div><span class="text-gray-600 font-bold me-1">$</span>{{ showStore.bookInfo.price }}</div>
                <button type="button"
                    class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">借閱
                    / 閱讀</button>
            </div>
        </div>

        <div class="text-sm font-medium text-center text-gray-500 border-b border-gray-200">
            <ul class="flex flex-wrap -mb-px" data-tabs-toggle="#default-tab-content" role="tab">
                <li class="me-2">
                    <a href="#" role="tab" data-tabs-target="#profile"
                        class="inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg">簡介</a>
                </li>
                <li class="me-2">
                    <a href="#" role="tab" data-tabs-target="#dashboard"
                        class="inline-block p-4 text-blue-600 border-b-2 border-blue-600 rounded-t-lg">章節</a>
                </li>
            </ul>
            <div id="default-tab-content">
                <div class="hidden p-4 rounded-lg bg-gray-50 min-h-96" id="profile" role="tabpanel">
                    <p class="text-sm text-gray-500">{{ showStore.bookInfo.introduction }}</p>
                </div>
                <div class="hidden p-4 rounded-lg bg-gray-50 min-h-96" id="dashboard" role="tabpanel">
                    <p class="text-sm text-gray-500">{{ showStore.bookInfo.chapter }}</p>
                </div>

            </div>
        </div>
    </div>
</template>
<script setup lang="ts" name="Product">
import { useBookInfoStore } from '../store/book_info.ts'
const showStore = useBookInfoStore()
</script>