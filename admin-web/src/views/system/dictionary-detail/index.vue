<template>
    <n-layout has-sider sider-placement="right">
<!--        {{ $route.params.id }}-->
        <n-layout-content>
            <list v-show="activeName === 'data'" ref="dataRef" @activeMethod="activeTab"></list>
            <edit v-show="activeName === 'form'" ref="formRef" @activeMethod="activeTab"></edit>
        </n-layout-content>
        <n-layout-sider bordered style="height: calc(100vh - 113px);"
                        width="46">
            <n-tabs v-model:value="activeName" placement="right" type="line" size="small"
                    tab-style="writing-mode: tb-rl;">
                <n-tab-pane name="data" tab="数 据"></n-tab-pane>
                <n-tab-pane name="form" tab="表 单"></n-tab-pane>
            </n-tabs>
        </n-layout-sider>
    </n-layout>
</template>
<script setup>
import list from "./list.vue";
import edit from "./edit.vue";

const activeName = ref('data')
const dataRef = ref(null);
const formRef = ref(null)
const activeTab = (param,id) => {
    if(param == 'data'){
        dataRef.value.tableProps.fetchData()
    }else{
        formRef.value.toEdit(id)
    }
    activeName.value = param;
};
</script>