import {defineStore} from 'pinia'

export const useBookInfoStore=defineStore('book_info',{
    //函式
    actions:{
    },
    //數據儲存
    state(){
        return{
            bookInfo:
                { 
                    "title": "The Secret Garden",
                    "writer": "Frances Hodgson Burnett",
                    "publisher": "Garden Publishing House",
                    "publishDate": "2023-07-15",
                    "uploader": "Bookworm98",
                    "ISBN": "9781234567890",
                    "introduction": "The Secret Garden is a classic children's novel about an orphaned girl named Mary Lennox who discovers a hidden garden that transforms her life and the lives of those around her.",
                    "chapter": "12 chapters",
                    "maxRentTime": 300000,
                    "price": 12.99,
                    "class":{
                        "className": "textbook",
                        "grade": 13,
                    },
                    "amount": 50,
                    "edition": "First edition",
                    "pages": 250,
                    "live": true,
                    "cover_image": "https://s3.amazonaws.com/virginia.webrand.com/virginia/50IgnOZ10MZ/cf9ca0abf2a77c311fc44d9f7ad1a99e/344/1648650462.png",
                },
            
        }
    }
})