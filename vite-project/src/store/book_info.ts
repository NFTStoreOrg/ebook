import { defineStore } from 'pinia'

interface BookInfo {
    title: string;
    writer: string;
    publisher: string;
    publishDate: string;
    uploader: string;
    ISBN: string;
    introduction: string;
    chapter: string;
    maxRentTime: number;
    price: number;
    class: Class;
    amount: number;
    edition: string;
    pages: number;
    live: boolean;
    cover_image: string;
}

interface Class {
    className: string;
    grade: number;
}

export const useBookInfoStore = defineStore('book_info', {
    //函式
    actions: {
    },
    //數據儲存
    state: (): BookInfo => ({
        title: '',
        writer: '',
        publisher: '',
        publishDate: '',
        uploader: '',
        ISBN: '',
        introduction: '',
        chapter: '',
        maxRentTime: 0,
        price: 0,
        class: {
            className: '',
            grade: 0,
        },
        amount: 0,
        edition: '',
        pages: 0,
        live: false,
        cover_image: '',
    })
})