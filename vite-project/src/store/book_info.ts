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
                    "title": "追尋失落的時光：一個旅人的心靈之旅",
                    "writer": "Frances Hodgson Burnett",
                    "publisher": "Garden Publishing House",
                    "publishDate": "2023-07-15",
                    "uploader": "Bookworm98",
                    "ISBN": "9781234567890",
                    "introduction": "The Secret Garden 是一部經典的兒童小說，故事主要講述了一個名叫瑪莉·倫諾克斯的孤兒女孩，她偶然發現了一個隱藏的花園，這個發現徹底改變了她以及周圍人的生活。在這個故事中，瑪莉不僅找到了一個神秘的花園，也重新找到了生活的意義和真正的家。這個花園原本被忘記了許多年，但隨著瑪莉的努力和她與其他孩子們的友誼，這片花園逐漸重現了昔日的光彩。這部小說通過描寫孩子們的友情、對大自然的探索以及生命力的奇跡，深深觸動了讀者的心靈，讓人不禁沉浸在這個充滿希望和奇蹟的世界中。",
                    "chapter": "12 章節 1.《秘密花園》是一部經典兒童小說，講述了一個名叫瑪莉·倫諾克斯的孤兒女孩，她偶然發現了一個隱藏的花園，這個發現不僅改變了她的生活，也深遠影響了周圍人的命運。這個神秘花園原本長年被遺忘，但隨著瑪莉的到來和她與新朋友們的努力，花園逐漸恢復了昔日的繁華和生機，成為孩子們探索、成長和治癒的場所。這部小說通過瑪莉的冒險故事，描繪了友情、家庭和希望的重要性，深深觸動讀者的心靈，讓人不禁沉浸在這個充滿魔力與奇蹟的世界中。《秘密花園》是一部經典的兒童小說，故事以瑪莉·倫諾克斯這位孤兒女孩為中心，她偶然發現了一個隱藏的花園，這個發現不僅改變了她的生命軌跡，也在無意間改變了身邊人的命運。這個神秘的花園從瑪莉的眼中重新焕發出生命的光彩，成為一個孩子們探索、學習和成長的理想場所。小說通過生動的筆觸和深刻的情感描寫，呈現了友誼和家庭的重要性，以及自然力量帶來的奇蹟，無疑是一部不可錯過的文學經典。",
                    "maxRentTime": 300000,
                    "price": 299,
                    "class":{
                        "className": "textbook",
                        "grade": 13,
                    },
                    "amount": 50,
                    "edition": "First edition",
                    "pages": 250,
                    "live": true,
                    "cover_image": "https://s3.amazonaws.com/virginia.webrand.com/virginia/nMzY6NJ97Xs/0d40507679ea0528a8cbc080bf0a74e6/344/1648657152.png",
                },
            
        }
    }
})