<template>
    <div class="relative -top-20 blur-lg bg-repeat-x bg-center min-w-[90vh] min-h-[50vh] 
        " :style="{ 'background-image': 'url(' + showStore.cover_image + ')' }">
    </div>
    <div class="container max-w-[90vh]">
        <div
            class="container absolute min-w-[100vh] max-h-max grid grid-cols-3 gap-6 top-60 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
            <div class="justify-self-start">
                <img :src=showStore.cover_image
                    class="max-h-[40vh] mx-auto outline outline-offset-2 outline-1 rounded outline-slate-400/50" />
            </div>
            <div class="rounded-md">
                <div class="text-3xl tracking-wide font-bold mb-4 text-gray-800">{{ showStore.title }}</div>
                <div class="text-lg tracking-wider font-medium">EISBN:{{ showStore.ISBN }}
                </div>
                <div class="text-lg tracking-wider font-medium">作者:{{ showStore.writer }}
                </div>
                <div class="text-lg tracking-wider font-medium">出版社:{{
                    showStore.publisher }}</div>
                <div class="text-lg tracking-wider font-medium">
                    出版日期:{{ showStore.publishDate }}</div>
            </div>
            <div class="rounded-md bg-blue-50/50 flex justify-center min-w-[70%] justify-self-end p-4">
                <div class="content-center">
                    <div class="text-2xl font-medium mb-2 flex justify-between">
                        <span class="text-sm">剩餘數量</span>{{ showStore.amount }}
                    </div>
                    <div class="text-2xl font-medium text-sky-600 mb-4 flex justify-between">
                        <span class="text-sm">優惠價</span>${{ showStore.price }}
                    </div>
                    <button type="button" v-if="canRead" @click="handleReadPermission"
                        class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">閱讀</button>
                    <button type="button" v-if="!canRead" @click="handleRent"
                        class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800">借閱</button>
                </div>
            </div>
        </div>

        <div class="text-sm font-medium text-gray-500 border border-gray-200 relative top-20 rounded-lg">
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
            <div id="default-tab-content" class="bg-blue-50/75">
                <div class="lg:ml-20 md:ml-8 sm:ml-0 hidden p-4 rounded-lg min-h-96 text-slate-700 text-lg tracking-wider font-medium grid grid-cols-2 gap-2 content-center"
                    id="detail" role="tabpanel">
                    <div>EISBN:{{ showStore.ISBN }}</div>
                    <div>作者:{{ showStore.writer }}</div>
                    <div>出版社:{{ showStore.publisher }}</div>
                    <div>出版日期:{{ showStore.publishDate }}</div>
                    <div>上傳者:{{ showStore.uploader }}</div>
                    <div>上傳時間:{{ showStore.uploadTime }}</div>
                    <div>種類:{{ showStore.class.className }}</div>
                    <div>年級:{{ showStore.class.grade }}</div>
                    <div>版本:{{ showStore.edition }}</div>
                    <div>互動書:{{ showStore.live }}</div>
                </div>
                <div class="hidden p-4 rounded-lg min-h-96 text-slate-700" id="introduction" role="tabpanel">
                    <p class="text-lg tracking-wider font-medium">{{
                        showStore.introduction }}</p>
                </div>
                <div class="hidden p-4 rounded-lg min-h-96 text-slate-700" id="chapter" role="tabpanel">
                    <p class="text-lg tracking-wider font-medium">{{ showStore.chapter }}
                    </p>
                </div>
            </div>
        </div>
        <div class="relative text-xl font-bold top-20 mt-6">其他可能有興趣</div>
        <div class="relative top-20 min-h-96 rounded-lg mb-6">
            <CardSlide :title=showStore.class.className slide-bg-color="bg-purple-100" :cardStore=bookArray cardWidth1="w-[150px]"></CardSlide>
        </div>
    </div>
</template>
<script setup lang="ts" name="Product">
import { useBookInfoStore } from '../store/book_info.ts'
import CardSlide from '../components/CardSlide.vue';
import { useCardStore } from '../store/card.ts';
import { useRoute } from 'vue-router';
import { onMounted, ref, watchEffect } from 'vue';
import axios from 'axios';
import { useSignatureStore } from '../store/personal_sign.ts';

//  Define bookArray's class
interface Book {
    title: string;
    imgUrl: string;
    price: number;
    bookId: number;
}

const showStore = useBookInfoStore()
const cardInfoStore = useCardStore()
const sign = useSignatureStore()
const route = useRoute()


const bookClass = ref('')
const bookArray = ref<Book[]>([]);
let canRead = ref(false)
let address: string | null | undefined;
sign.initSDK()

onMounted(async () => {
    const tokenId = route.query.bookId
    if (tokenId) {
        await showStore.getBookInfo(tokenId)
        //  Ensure that set bookClass after get book information
        bookClass.value = showStore.class.className
        cardInfoStore.getClassBook(bookClass.value)
        await showStore.getBookRemain(tokenId)
    }
    await sign.onConnect()
    address = sign.account
    await readOrRent()
})

watchEffect(() => {
    if (bookClass.value == "reference") {
        bookArray.value = cardInfoStore.referenceBook
    } else if (bookClass.value == "children") {
        bookArray.value = cardInfoStore.childrenBook
    } else if (bookClass.value == "textbook") {
        bookArray.value = cardInfoStore.textbook
    } else if (bookClass.value == "other") {
        bookArray.value = cardInfoStore.otherBook
    } else if (bookClass.value == "video") {
        bookArray.value = cardInfoStore.video
    }
})

//  Use for display button word read or rent.
async function readOrRent() {
    let tokenId = route.query.bookId
    try {
        let response = await axios.get("https://yisinnft.org/api/" + address + "/" + tokenId + "/read")
        canRead.value = response.data.rentStatus
    } catch (err) {
        console.error(err)
    }
}

async function handleReadPermission() {

    //  Get signature
    let signature: any;
    try {
        signature = await sign.getSign()
    } catch (err) {
        console.error(err)
        return
    }

    let tokenId = route.query.bookId;
    try {
        let response = await axios({
            url: "https://yisinnft.org/api/" + address + "/" + tokenId + "/" + signature,
            method: 'GET',
            responseType: 'blob',    // Ensure that response type is blob
        });

        // Extract file name in header
        const contentDisposition = response.headers['content-disposition'];
        let fileName = 'downloaded_file';
        if (contentDisposition) {
            const fileNameMatch = contentDisposition.match(/filename="?(.+)"?/);
            if (fileNameMatch && fileNameMatch.length === 2) {
                fileName = fileNameMatch[1];
            }
        } else {
            console.error("Get header fail!")
            return
        }

        // Simulation user click
        const url = window.URL.createObjectURL(new Blob([response.data]));
        console.log(response.data)
        const link = document.createElement('a');
        link.href = url;     // Set url to download
        link.setAttribute('download', fileName);
        document.body.appendChild(link);     // Add the a tag
        link.click();        // Click link
        document.body.removeChild(link);     // Remove the a tag
    } catch (err) {
        console.error(err);
    }
}

async function handleRent() {
    if (!sign.connected) {
        await sign.onConnect()
    }
}

</script>