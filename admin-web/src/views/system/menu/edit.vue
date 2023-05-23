<template>
    <n-layout>
        <n-layout-header ref="header">
            <n-space justify="center">
                <h3>表单</h3>
            </n-space>
<!--            <pre>{{ JSON.stringify(data, null, 2) }}</pre>-->

        </n-layout-header>
        <n-layout-content style="margin-top: 10px;margin-right: 0px">
            <n-space justify="space-evenly">
                <n-form size="small"
                        :style="{maxWidth: '640px',minWidth:'320px'}">

                    <n-form-item label="父菜单ID">
                        <n-input-number  placeholder="parentId" v-model:value="data.parentId" :show-button="false"/>
                        &nbsp;<n-button @click="showModal = true" >选择</n-button>
                    </n-form-item>


                    <n-form-item label="菜单名">
                        <n-input placeholder="title" v-model:value="data.title"/>
                    </n-form-item>


                    <n-form-item label="菜单图标">
                        <gcs-icon-select v-model:value="data.icon"></gcs-icon-select>
                    </n-form-item>


                    <n-form-item label="路由名">
                        <n-input placeholder="name" v-model:value="data.name"/>
                    </n-form-item>


                    <n-form-item label="路由路径">
                        <n-input placeholder="path" v-model:value="data.path"/>
                    </n-form-item>


                    <n-form-item label="组件路径">
                        <n-input placeholder="component" v-model:value="data.component"/>
                    </n-form-item>


                    <n-form-item label="是否隐藏">
                        <n-switch v-model:value="data.hidden"/>
                    </n-form-item>


                    <n-form-item label="排序标记">
                        <n-input-number placeholder="sort" v-model:value="data.sort" :show-button="false"/>
                    </n-form-item>

                    <n-form-item>
                        <n-space justify="center">
                            <n-button attr-type="button" @click="submit">
                                保 存
                            </n-button>
                            <n-button attr-type="button" @click="backList">
                                返 回
                            </n-button>
                        </n-space>
                    </n-form-item>
                </n-form>
            </n-space>
            <n-modal v-model:show="showModal">
                <gcs-data-select :columnsMetaData="columnsMetaData" :columns="columns" :query-url="'/menu/list'"
                                 v-model:retVal="data.parentId"
                                 v-model:showModal="showModal"
                >
                </gcs-data-select>
            </n-modal>
        </n-layout-content>
    </n-layout>
</template>
<script setup>
import {http} from "@/core/http";
import GcsIconSelect from "@/components/icon/GcsIconSelect.vue";
import {columnsMetaData} from "./definition.js";
import {columns} from "./custom.js";

const showModal = ref(false)

const data = reactive({});

const toEdit = async (id) => {
    if (id) {
        const apiRet = await http.post('/menu/one', {'id': id})
        if (apiRet.ok) {
            Object.assign(data, apiRet.data);
        }
    } else {
        data.id = null
    }
}

const submit = async () => {
    if (data.id) {
        const apiRet = await http.post('/menu/edit', data)
        if (apiRet.ok) {
            backList()
        }
    } else {
        const apiRet = await http.post('/menu/add', data)
        if (apiRet.ok) {
            backList()
        }
    }
}

const backList = () => {
    emits('activeMethod', 'data');
};

const emits = defineEmits(['activeMethod']);

defineExpose({
    toEdit
})
</script>