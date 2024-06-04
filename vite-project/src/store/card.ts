import {defineStore} from 'pinia'

export const useCardStore=defineStore('card',{
    //函式
    actions:{
        // increment(value){
        //     this.sum+=value
        // }
    },
    //數據儲存
    state(){
        return{
            books:[
                {
                    title: "追尋失落的時光：一個旅人的心靈之旅",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/50IgnOZ10MZ/cf9ca0abf2a77c311fc44d9f7ad1a99e/344/1648650462.png",
                    price: 100
                },
                {
                    title: "探索無盡的宇宙：星際航行的奇妙之旅",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/TUjmSfBSBs2/3e59d5989f9eceaca8625b6bf780c2e4/344/1648665973.png",
                    price: 100
                },
                {
                    title: "破曉之前：夢幻世界的神秘啟示",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
                    price: 100
                },
                {
                    title: "靈魂的奇跡：生命中的意義與使命",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
                    price: 100
                },
                {
                    title: "時光的記憶：遺忘與回憶的交錯",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/nMzY6NJ97Xs/0d40507679ea0528a8cbc080bf0a74e6/344/1648657152.png",
                    price: 100
                },
                {
                    title: "破曉之前：夢幻世界的神秘啟示",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
                    price: 100
                },
                {
                    title: "靈魂的奇跡：生命中的意義與使命",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
                    price: 100
                },
                {
                    title: "追尋失落的時光：一個旅人的心靈之旅",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/50IgnOZ10MZ/cf9ca0abf2a77c311fc44d9f7ad1a99e/344/1648650462.png",
                    price: 100
                },
                {
                    title: "探索無盡的宇宙：星際航行的奇妙之旅",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/TUjmSfBSBs2/3e59d5989f9eceaca8625b6bf780c2e4/344/1648665973.png",
                    price: 100
                },
                {
                    title: "破曉之前：夢幻世界的神秘啟示",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/VGMSCiN4BsQ/0b08935b2765516d8c98167e6f47ab73/344/1648652486.png",
                    price: 100
                },
                {
                    title: "靈魂的奇跡：生命中的意義與使命",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/yQev8MF16sZ/01e52a0edf3bf72541ecc41f0b41fc58/344/1648666817.png",
                    price: 100
                },
                {
                    title: "時光的記憶：遺忘與回憶的交錯",
                    imgUrl: "https://s3.amazonaws.com/virginia.webrand.com/virginia/nMzY6NJ97Xs/0d40507679ea0528a8cbc080bf0a74e6/344/1648657152.png",
                    price: 100
                },
            ]
        }
    }
})