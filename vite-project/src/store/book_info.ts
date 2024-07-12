import axios from 'axios';
import { defineStore } from 'pinia'
import { LocationQueryValue } from 'vue-router';

//  Define return info format
interface BookInfo {
    title: string;
    writer: string;
    publisher: string;
    publishDate: string;
    uploader: string;
    uploadTime: string;
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
    tokenId: number;
}

interface Class {
    className: string;
    grade: number;
}

export const useBookInfoStore = defineStore('book_info', {
    actions: {
        async getBookInfo(tokenId: string | LocationQueryValue[]) {
            try {
                //  Get book information from database
                let response = await axios.get("https://yisinnft.org/api/book/" + tokenId)
                //  Store book information, except for amount
                this.title = response.data.data.title;
                this.writer = response.data.data.writer;
                this.publisher = response.data.data.publisher;
                this.publishDate = response.data.data.publish_date;
                this.uploader = response.data.data.uploader;
                this.uploadTime = this.convertUnixToReadable(response.data.data.upload_time);
                this.ISBN = response.data.data.ISBN;
                this.introduction = response.data.data.introduction;
                this.chapter = response.data.data.chapter;
                this.maxRentTime = response.data.data.max_rent_time;
                this.price = response.data.data.price;
                this.class = {
                    className: response.data.data.class.class_name,
                    grade: response.data.data.class.grade,
                };
                this.edition = response.data.data.edition;
                this.pages = response.data.data.pages;
                this.live = response.data.data.live;
                this.cover_image = response.data.data.cover_image;
                this.tokenId = response.data.data.tokenId;


            } catch (err) {
                console.log(err)
            }
        },
        //  For more effiency, use another function to get on-chain information
        async getBookRemain(tokenId: string | LocationQueryValue[]) {
            let remain = await axios.get("https://yisinnft.org/api/book/remain/" + tokenId)
            this.amount = remain.data.remaining_amount;
        },
        convertUnixToReadable(unixTimestamp: number): string {
            //  Transform unix timestamp to ms
            const date = new Date(unixTimestamp * 1000)

            //  Format date and time
            const year = date.getFullYear()
            //  Month start with 0, need to add 1
            const month = ('0' + (date.getMonth() + 1)).slice(-2)
            const day = ('0' + date.getDate()).slice(-2)
            const hours = ('0' + date.getHours()).slice(-2)
            const minutes = ('0' + date.getMinutes()).slice(-2)
            const seconds = ('0' + date.getSeconds()).slice(-2)

            return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`
        }
    },
    state: (): BookInfo => ({
        title: '',
        writer: '',
        publisher: '',
        publishDate: '',
        uploader: '',
        uploadTime: '',
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
        tokenId: 0,
    })
})