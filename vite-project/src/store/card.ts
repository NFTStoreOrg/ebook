import { defineStore } from 'pinia'
import axios from 'axios'

export const useCardStore = defineStore('card', {
    //函式
    actions: {
        async getNewestBook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/' + 'index')
                this.newBooks = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getReferenceBook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/reference')
                this.referenceBook = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getChildBook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/children')
                this.childrenBook = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getTextbook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/textbook')
                this.textbook = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getOtherBook() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/other')
                this.otherBook = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
        async getVideo() {
            try {
                let response = await axios.get('https://yisinnft.org/api/book/index/video')
                this.video = response.data.data.map((book: { title: string; cover_image: string; price: number }) => ({
                    title: book.title,
                    imgUrl: book.cover_image,
                    price: book.price,
                }))
            } catch (error) {
                console.error("API request failure:", error)
            }
        },
    },
    //數據儲存
    state() {
        return {
            newBooks: [] as Array<{ title: string, imgUrl: string, price: number }>,
            referenceBook: [] as Array<{ title: string, imgUrl: string, price: number }>,
            textbook: [] as Array<{ title: string, imgUrl: string, price: number }>,
            childrenBook: [] as Array<{ title: string, imgUrl: string, price: number }>,
            otherBook: [] as Array<{ title: string, imgUrl: string, price: number }>,
            video: [] as Array<{ title: string, imgUrl: string, price: number }>,
        }
    }
})