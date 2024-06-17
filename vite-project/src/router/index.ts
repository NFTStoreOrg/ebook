import {createRouter,createWebHistory} from 'vue-router'
import Main from '../pages/Main.vue'
import ReferenceBook from '../pages/ReferenceBook.vue'
import Account from '../pages/Account.vue'
import posted from '../account_page/posted.vue'
import renting from '../account_page/renting.vue'

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
            children: [
                { path:'/posted' , component: posted},
                { path:'/renting' , component: renting}
              ],
        },
    ]
})
router.beforeEach((to, from) => {
    const showStore = useShowStore()
    if(to.fullPath!='/Account' && to.fullPath!='/posted'&&to.fullPath!='/renting'){
        showStore.showHeader=true
    }
})
export default router