import {defineStore} from "pinia"

export const useTabsStore = defineStore('tabsStore', {
    state:()=>(
        {
            tabs:[
                // {
                //     name:"主页"
                // },
            ],
            currentName:""
        }
    ),
    getters:{
        closable:(state)=>{
            return state.tabs.length > 1;
        },
        getTab:(state)=>{
            return (tabName) => {
                let nameIndex = state.tabs.findIndex((value) => value.name === tabName)
                if (~nameIndex){
                    return state.tabs[nameIndex]
                }
            }
        }
    },
    actions: {
        activateTab(tab) {
            let nameIndex = this.tabs.findIndex((value) => value.name === tab.name)
            if (!~nameIndex){
                this.tabs.push(tab)
            }
            this.currentName = tab.name;
        },
        removeTab(tabName){
            let nameIndex = this.tabs.findIndex((tab) => tab.name === tabName)
            if (~nameIndex){
                this.tabs.splice(nameIndex, 1);
            }
            // if (name === valueRef.value) {
            //     valueRef.value = panels[Math.min(nameIndex, panels.length - 1)];
            // }
        }
    },
    persist:false

})