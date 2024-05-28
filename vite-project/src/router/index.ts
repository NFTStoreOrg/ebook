import {createRouter,createWebHistory} from 'vue-router'
import Main from '../pages/Main.vue'
import ReferenceBook from '../pages/ReferenceBook.vue'

const router=createRouter({
    history:createWebHistory(),
    routes:[
        { 
            path:'/',
            component: Main
        },
        { 
            path:'/ReferenceBook',
            component: ReferenceBook
        },
        { 
            path:'/NFT',
            component: {
                beforeCreate() {
                    window.location.href = 'https://yisinnft.org/';
                }
            }
        },
    ]
})
export default router