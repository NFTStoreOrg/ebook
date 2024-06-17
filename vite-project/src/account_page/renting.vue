<template>
    <div class="p-6 bg-gray-50 text-medium text-gray-500 dark:text-gray-400 dark:bg-gray-800 rounded-lg w-full">
        已借閱
        <div class="grid grid-cols-1 gap-8 mt-3 lg:grid-cols-4 md:grid-cols-3 xl:grid-cols-6 mx-auto">
            <card></card>
        </div>
    </div>
</template>
<script setup lang="ts" name="renting">
import card from '../components/Card.vue'

</script>
