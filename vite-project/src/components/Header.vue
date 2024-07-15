<template>
  <header class="fixed inset-x-0 top-0 z-10 bg-opacity-[90%]" ref="scrollDown">
    <nav class="flex items-center justify-between p-6 lg:px-8">
      <div class="flex font-bold text-gray-400">
        EBOOK STORE
      </div>
      <div class="flex lg:gap-x-12 gap-x-7">
        <RouterLink to="/" class="text-sm font-semibold"
          :class="$route.path === '/' ? 'text-gray-900' : 'text-indigo-900'">首頁</RouterLink>
        <RouterLink to="/" class="text-sm font-semibold text-indigo-900">教科書</RouterLink>
        <RouterLink to="/ReferenceBook" class="text-sm font-semibold"
          :class="$route.path === '/ReferenceBook' ? 'text-gray-900' : 'text-indigo-900'">參考書</RouterLink>
        <RouterLink to="/" class="text-sm font-semibold text-indigo-900">兒童書籍</RouterLink>
        <RouterLink to="/" class="text-sm font-semibold text-indigo-900">互動書</RouterLink>
        <RouterLink to="/" class="text-sm font-semibold text-indigo-900">影片</RouterLink>
        <RouterLink to="/" class="text-sm font-semibold text-indigo-900">其他</RouterLink>
        <RouterLink target="_blank" to="/NFT" class="text-sm font-semibold text-indigo-900">NFT</RouterLink>
      </div>

      <div class="flex items-center">
        <RouterLink to="/Account">
          <button type="button"
            class="text-xs me-2 text-gray-900 focus:outline-none bg-white rounded-full border border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100 dark:focus:ring-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:border-gray-600 dark:hover:text-white dark:hover:bg-gray-700">
            Account
          </button>
        </RouterLink>
        <a class="text-sm font-semibold text-gray-900 cursor-pointer" ref="modalPop"
          data-modal-target="authentication-modal" data-modal-toggle="authentication-modal">Log in <span
            aria-hidden="true">&rarr;</span></a>
      </div>
    </nav>
  </header>
  <!-- Main modal -->
  <teleport to="body">
    <div id="authentication-modal" aria-hidden="true" data-modal-backdrop="static" @click="handleBackgroundClick"
      class="animate-fade-down hidden overflow-y-auto overflow-x-hidden fixed top-0 right-0 left-0 z-50 justify-center items-center w-full md:inset-0 h-[calc(100%-1rem)] max-h-full">
      <div class="relative p-4 w-full max-w-md max-h-full transition-transform duration-1000 ease-in-out"
        ref="modalPop">
        <!-- Modal content -->
        <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">
              Sign in to our platform
            </h3>
            <button type="button"
              class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white font-mono"
              data-modal-hide="authentication-modal">
              X
              <span class="sr-only">Close modal</span>
            </button>
          </div>
          <!-- Modal body -->
          <div class="p-4 md:p-5">
            <form class="space-y-4" action="#">
              <div>
                <label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your
                  email</label>
                <input type="email" name="email" id="email"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  placeholder="name@company.com" required />
              </div>
              <div>
                <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Your
                  password</label>
                <input type="password" name="password" id="password" placeholder="••••••••"
                  class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white"
                  required />
              </div>
              <div class="flex justify-between">
                <div class="flex items-start">
                  <div class="flex items-center h-5">
                    <input id="remember" type="checkbox" value=""
                      class="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-blue-300 dark:bg-gray-600 dark:border-gray-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 dark:focus:ring-offset-gray-800"
                      required />
                  </div>
                  <label for="remember" class="ms-2 text-sm font-medium text-gray-900 dark:text-gray-300">Remember
                    me</label>
                </div>
                <a href="#" class="text-sm text-blue-700 hover:underline dark:text-blue-500">Lost
                  Password?</a>
              </div>
              <button type="submit"
                class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Login
                to your account</button>
              <div class="text-sm font-medium text-gray-500 dark:text-gray-300">
                Not registered? <a href="#" class="text-blue-700 hover:underline dark:text-blue-500">Create
                  account</a>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </teleport>
</template>
<script setup lang="ts" name="Header">
import { ref } from 'vue'

var modalPop = ref()

const handleBackgroundClick = (event: any) => {
  if (event.target.id === 'authentication-modal') {
    modalPop.value.classList.add('scale-105')
    setTimeout(() => {
      modalPop.value.classList.remove('scale-105');
    }, 300);
  }
};

var scrollDown = ref()

window.addEventListener("scroll", handleScroll);
function handleScroll() {
  if (window.scrollY >= 200) {
    scrollDown.value.classList.add('bg-slate-100');
  } else {
    scrollDown.value.classList.remove('bg-slate-100');
  }

}

</script>