import {ComponentOptions} from 'vue'
declare module '@vue/runtime-core'{
    interface AppConfig{
        compatConfig?:Record<string,any>
    }
}