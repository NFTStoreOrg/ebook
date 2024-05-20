import {createRouter,createWebHistory} from 'vue-router'
import reference_book from '@/components/reference_book.vue'
const router=createRouter({
    history:createWebHistory(),
    routes:[
        {
            path:'/reference_book',
            component:reference_book
        }
    ]
})
export default router