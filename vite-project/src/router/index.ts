import {createRouter,createWebHistory} from 'vue-router'
import Main from '../pages/Main.vue'
import ReferenceBook from '../pages/ReferenceBook.vue'
import Account from '../pages/Account.vue'
import { useShowStore } from '../store/show.ts'

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
        { 
            path:'/Account',
            component: Account,
        },
    ]
})
router.beforeEach((to, from) => {
    const showStore = useShowStore()
    if(to.fullPath!='/Account'){
        showStore.showHeader=true
    }
})
export default router