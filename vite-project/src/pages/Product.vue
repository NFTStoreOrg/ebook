<template>
    <div class="container min-w-[90vh] min-h-[85vh]">
        <div class="grid grid-cols-2">
            <img :src=showStore.bookInfo.cover_image />

            <div class="grid grid-rows-6">
                {{ showStore.bookInfo.title }}
                {{ showStore.bookInfo.writer }}
                {{ showStore.bookInfo.publisher }}
                {{ showStore.bookInfo.publishDate }}
                {{ showStore.bookInfo.ISBN }}
                {{ showStore.bookInfo.introduction }}
            </div>
        </div>


    </div>
</template>
<script setup lang="ts" name="Product">
import { useBookInfoStore } from '../store/book_info.ts'
const showStore = useBookInfoStore()

</script>