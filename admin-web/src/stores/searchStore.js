import {defineStore} from "pinia"
//import {parse, stringify} from 'zipson'

export const useSearchStore = defineStore('searchStore', {
    state: () => (
        {
            searchCondition: new Map(),
            scData: []
        }
    ),
    getters: {
        getSearchCondition: (state) => {
            return state.scData;
        }
    },
    actions: {
        setSearchCondition(path, data) {
            let value = this.scData.find(item => {
                   if(!item){
                       return false;
                   }else{
                       if(item[path]){
                           return true
                       }
                   }
                }
            )
            if(value){
                value[path].push(data)
            }else{
                const urlObj = { };
                urlObj[path]=[]
                urlObj[path].push(data);
                this.scData.push(urlObj)
            }
        },
        removeSearchCondition(path,data,index){
            let value = this.scData.find(item => {
                    if(!item){
                        return false;
                    }else{
                        if(item[path]){
                            return true
                        }
                    }
                }
            )
            if(value){
                value[path].splice(index,1)

            }
        }
    },
    persist: { // 自定义数据持久化方式
        key: 'search-key',
        paths: ['scData'],
        serializer: {
            deserialize: (value) => {
                return JSON.parse(value)
            },
            serialize: (value) => {
                // console.log(value, "121211212")
                return JSON.stringify(value)
            },
        },
        storage: localStorage,

    },
    //debug: true

})