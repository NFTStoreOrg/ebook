import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// https://vitejs.dev/config/
export default defineConfig({
  build:{
    //  output directory
    outDir:'dist',
    //  compress or not
    minify: false,
  },
  plugins: [vue()
  ],
  resolve:{alias:{'vue':'@vue/compat'}}
})
