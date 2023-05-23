import {defineStore} from "pinia"
import { darkTheme } from "naive-ui";

export const useGlobalStore = defineStore('globalStore', {
    state:()=>(
        {
            theme:undefined,
            themeName:"深色"
        }
    ),
    getters:{
        getTheme:(state)=>{
            if(state.themeName === "深色"){
                state.theme = darkTheme
            }else{
                state.theme = undefined
            }
            return state.theme
        }
    },
    actions: {
        setTheme(){
            if(this.themeName === "深色"){
                this.themeName = "浅色"
                this.theme = undefined
            }else{
                this.themeName = "深色"
                this.theme = darkTheme
            }
        },
    },

    persist: {
        key: 'themeName',
        paths: ['themeName'],
        storage: localStorage,
    }
})