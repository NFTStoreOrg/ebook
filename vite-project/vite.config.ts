import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { createVuePlugin as vueCompat } from 'vite-plugin-vue2';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(),
    vueCompat()
  ],
  resolve:{alias:{'vue':'@vue/compat'}}
})
