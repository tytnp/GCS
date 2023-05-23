<template>
    <n-tabs
            type="card"
            size="small"
            :closable="store.closable"
            @close="store.removeTab"
            @update:value="handleUpdateValue"
            v-model:value="store.currentName"
            closable>
        <!--        addable="addable"-->
        <n-tab v-for="panel in store.tabs" :name="panel.name">
            {{ panel.name }}
        </n-tab>
    </n-tabs>
</template>
<script setup>
import {useTabsStore} from "@/stores/tabsStore.js"

const router = useRouter()
const store = useTabsStore()

const handleUpdateValue = (value) => {
    const tabInfo = store.getTab(value)
    router.push(tabInfo.path)
}

onMounted(() => {
    let route = router.getRoutes().find((r) => r.name === router.currentRoute.value.name);
    if(route){
        store.activateTab(
            {
                name: route.meta.title,
                path: route.path,
                routeName: route.name
            }
        )
    }else{
        store.activateTab(
            {
                name: "主页",
                path: "/root",
                routeName: "root"
            }
        )
    }
})
</script>
