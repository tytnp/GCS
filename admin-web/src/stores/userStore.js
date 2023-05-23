import {defineStore} from "pinia"
import {http} from "@/core/http/index.js";
import {RouterLink, useRouter} from "vue-router";
import {iconMap} from "@/components/icon/icons.js";
import {useTabsStore} from "@/stores/tabsStore.js"

const viewModules = import.meta.glob('@/views/**/*.vue');
export const useUserStore = defineStore('userStore', {
    state: () => (
        {
            routes: [],
            user: {},
            token: null
        }
    ),
    getters: {
        getMenuTree: async (state) => {
            if (state.routes.length == 0) {
                await state.asyncRouters()
            }
            let menuTree = formatMenuTree(state.routes);
            return menuTree.children
        },
        getRoutes: async (state) => {
            let routes = [];
            if (state.routes.length == 0) {
                await state.asyncRouters()
            }
            let result = formatRoutes(state.routes)
            if (result && result.children && result.children.length > 0) {
                routes.push(...result.children)
                delete result.children
                routes.push(result)
            }
            return routes
        },
    },
    actions: {
        async asyncRouters() {
            let ret = await http.post('/token_menu', {token: this.token})
            if (ret.ok) {
                this.routes = ret.data;
            } else {
                this.setToken(null)
            }
        },
        async asyncUserInfo() {
            let ret = await http.post('/token_user_info', {token: this.token})
            if (ret.ok) {
                this.user = ret.data;
            } else {
                this.setToken(null)
            }
        },
        setToken(t) {
            this.token = t
            if (!t) {
                this.routes = []
                this.user = {}
            }
        }
    },
    persist: {
        key: 'token',
        paths: ['token'],
        storage: localStorage,
    }

})

const formatRoutes = (data) => {
    let route = {
        path: data.path,
        name: data.name,
        meta:{
            title:data.title
        }
    }
    if (data.name === 'root') {
        route.name = data.name
        route.path = "/"
        route.redirect = data.path

    }
    if (typeof data.component === "string" && data.component.trim().length > 0) {
        route.component = viewModules[`/src/${data.component}`]
    }

    if (data.children && data.children.length > 0) {
        route.children = []
        for (const dataKey in data.children) {
            route.children.push(formatRoutes(data.children[dataKey]))
        }
    }
    return route
}

const formatMenuTree = (data) => {
    const renderIcon = (icon) => {
        return () => {
            if (iconMap[icon]) {
                return h(NIcon, null, {default: () => h(iconMap[icon])})
            }
            return null
        }
    }
    //判断是否隐藏菜单,隐藏返回空菜单
    if(data.hidden){
        return null
    }
    let menu = {
        key: data.name,
        icon: renderIcon(data.icon)
    }
    if (typeof data.path === "string" && data.path.trim().length === 0) {
        menu.label = data.title
    } else {
        if (data.children && data.children.length > 0) {
            menu.label = data.title
        } else {
            menu.label = () => h(
                RouterLink,
                {
                    to: {
                        path: data.path
                    },
                    onClick: (event) => {
                        const tabsStore = useTabsStore()
                        tabsStore.activateTab(
                            {
                                name: data.title,
                                path: data.path,
                                routeName: data.name
                            }
                        )
                    }
                },
                {
                    default: () => data.title,

                })
        }
    }

    if (data.children && data.children.length > 0) {
        menu.children = []
        for (const dataKey in data.children) {
            let childrenMenu = formatMenuTree(data.children[dataKey])
            if(childrenMenu!=null){
                menu.children.push(childrenMenu)
            }

        }
    }
    return menu
}
