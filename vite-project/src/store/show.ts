import {defineStore} from 'pinia'

export const useShowStore=defineStore('show',{
    //函式
    actions:{
    },
    //數據儲存
    state(){
        return{
            showHeader:true
        }
    }

})