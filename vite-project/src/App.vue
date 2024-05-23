<script setup lang="ts" name="App">
  import HelloWorld from './components/HelloWorld.vue'
  import {RouterView,RouterLink} from 'vue-router'
</script>

<template>
  <div>
  <b-dropdown id="dropdown-1" text="Dropdown Button" class="m-md-2">
    <b-dropdown-item>First Action</b-dropdown-item>
    <b-dropdown-item>Second Action</b-dropdown-item>
    <b-dropdown-item>Third Action</b-dropdown-item>
    <b-dropdown-divider></b-dropdown-divider>
    <b-dropdown-item active>Active action</b-dropdown-item>
    <b-dropdown-item disabled>Disabled action</b-dropdown-item>
  </b-dropdown>
</div>
  <!-- header -->
  <div class="row sticky top-0 m-3 mx-auto items-center z-50">
        <!-- 導覽 -->
        <div class="col">
            <ul class="nav justify-center ">
                <li class="nav-item">
                    <a class="nav-link" href="#" tabindex="-1" aria-disabled="true" style="color: rgb(48, 48, 48);">
                        <b>EBOOK STORE</b></a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="">參考書</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="">其他</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="../index.html">NFT</a>
                </li>

            </ul>
        </div>
        <!-- account -->
        <div class="col">
            <ul class="nav justify-end w-5/6">
                <li class="nav-item">
                    <button type="button" class="btn btn-outline-primary">登入
                    </button>
                </li>
            </ul>
        </div>
    </div>
  <RouterLink to="/ReferenceBook" active-class="">ReferenceBook</RouterLink>
  <RouterView></RouterView>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
    <a href="https://vuejs.org/" target="_blank">
      <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
    </a>
  </div>
  <HelloWorld msg="Vite + Vue" />
  
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
