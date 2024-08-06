<template>
  <form class="max-w-md mx-auto">
    <label for="default-search" class="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">Search</label>
    <div class="relative">
      <div class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
        <svg class="w-4 h-4 text-gray-500 dark:text-gray-400" aria-hidden="true" xmlns="http://www.w3.org/2000/svg"
          fill="none" viewBox="0 0 20 20">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
            d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z" />
        </svg>
      </div>
      <input type="search" v-model="inputValue" id="default-search" class="search-input" placeholder="搜尋書名..." />
    </div>
  </form>
  <div style="display: flex;">
    <div ref="content"
      class="overflow-x-auto p-4 cardContainer justify-start" >
      <Card :cardStore=result cardWidth2="w-[150px]"></Card>
    </div>

  </div>
  <div style="display: flex; justify-content: center; align-items: center;">
    <span>共 {{ amount }} 筆結果</span>
  </div>
</template>

<script setup lang="ts" name="Search">
import axios from 'axios';
import { ref, watchEffect } from 'vue';
import Card from '../components/Card.vue';

interface Book {
  title: string;
  imgUrl: string;
  price: number;
  bookId: number;
}
const result = ref<Book[]>([])
const inputValue = ref('')
const amount = ref(0)

watchEffect(async () => {
  result.value.length = 0
  amount.value = 0
  let res = await axios.get("https://yisinnft.org/api/es/" + inputValue.value)

  res.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => {
    let item: Book = {
      title: book.title,
      imgUrl: book.cover_image,
      price: book.price,
      bookId: book.tokenId
    }

    result.value.push(item)
  })

  amount.value = result.value.length
})

</script>
<style>
/*   Search input style */ 
.search-input {
  width: 100%;
  padding: 1rem;
  padding-left: 2.5rem;
  font-size: 0.875rem;
  color: #1f2937;
  border: 1px solid #d1d5db;
  border-radius: 5rem;
  background-color: rgba(255, 255, 255, .2);
  backdrop-filter: blur(50px);
  -webkit-backdrop-filter: blur(50px);
  box-shadow: 0 4px 6px rgba(0, 0, 0, .1);
}

.search-input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, .5);
}

/* Card container style */
.cardContainer{
  display: flex;
  flex-wrap: wrap;
}

</style>