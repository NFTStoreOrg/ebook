import {createRouter,createWebHistory} from 'vue-router'
import Main from '../pages/Main.vue'
import ReferenceBook from '../pages/ReferenceBook.vue'
import Product from '../pages/Product.vue'

import Account from '../pages/Account.vue'
import posted from '../account_page/posted.vue'
import renting from '../account_page/renting.vue'
import Header from '../components/Header.vue'

const router=createRouter({
    history:createWebHistory(),
    routes:[
        { 
            path:'/',
            components: {
                default:Main,
                Header: Header,
            }
        },
        { 
            path:'/ReferenceBook',
            components: {
                default:ReferenceBook,
                Header: Header,
            }
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
        {
             path:'/Product',
             components:{
                default:Product,
                Header: Header,
             } 
        }
    ]
})
export default router