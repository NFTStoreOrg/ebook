import { defineStore } from 'pinia'
import axios from 'axios'

export const useCardStore = defineStore('card', {
    //函式
    actions: {
        async getNewestBook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index')
                this.newBooks = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                    bookId: book.tokenId,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getClassBook(type: string) {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/' + type)
                if (type == "reference") {
                    this.referenceBook = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                        title: book.title,
                        imgUrl: book.cover_image,
                        price: book.price,
                        bookId: book.tokenId,
                    }))
                } else if (type == "children") {
                    this.childrenBook = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                        title: book.title,
                        imgUrl: book.cover_image,
                        price: book.price,
                        bookId: book.tokenId,
                    }))
                } else if (type == "textbook") {
                    this.textbook = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                        title: book.title,
                        imgUrl: book.cover_image,
                        price: book.price,
                        bookId: book.tokenId,
                    }))
                } else if (type == "other") {
                    this.otherBook = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                        title: book.title,
                        imgUrl: book.cover_image,
                        price: book.price,
                        bookId: book.tokenId,
                    }))
                } else if (type == "video") {
                    this.video = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                        title: book.title,
                        imgUrl: book.cover_image,
                        price: book.price,
                        bookId: book.tokenId,
                    }))
                }

            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getLiveBook() {
            let response = await axios.get("https://yisinnft.org/api/book/live")
            this.live = response.data.data.map((book: { title: string; cover_image: string; price: number; tokenId: number; }) => ({
                title: book.title,
                imgUrl: book.cover_image,
                price: book.price,
                bookId: book.tokenId,
            }))
        }
    },
    //數據儲存
    state() {
        return {
            newBooks: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            referenceBook: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            textbook: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            childrenBook: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            otherBook: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            video: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
            live: [] as Array<{ title: string, imgUrl: string, price: number, bookId: number }>,
        }
    }
})