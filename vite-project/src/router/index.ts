import {createRouter,createWebHistory} from 'vue-router'
import ReferenceBook from '../pages/ReferenceBook.vue'
const router=createRouter({
    history:createWebHistory(),
    routes:[
        { 
            path:'/ReferenceBook',
            component:ReferenceBook
        },
    ]
})
export default router