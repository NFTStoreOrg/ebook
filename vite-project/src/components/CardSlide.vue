<script setup lang="ts">
import Card from '../components/Card.vue'
defineProps<{ title: string, slideBgColor: string }>()
</script>
<template>
    <div class="row mx-auto text-2xl max-w-[95%] font-bold pl-3 pt-3 tracking-wider"><span
            class="bg-blue-100 text-blue-800 me-2 px-2.5 py-0.5 rounded dark:bg-blue-900 dark:text-blue-300">{{ title
            }}</span>
    </div>
    <div :class=slideBgColor class="row mx-auto rounded-3xl m-3 max-h-[80vh] max-w-[95%] bg-opacity-50">
        <div class="flex snap-x snap-mandatory items-center gap-x-3 overflow-x-auto p-4">
            <card></card>
        </div>
    </div>
</template>