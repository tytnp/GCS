<template>
    <n-layout>
        <n-layout-header ref="header">
            <n-space justify="center">
                <h3>表单</h3>
            </n-space>
        </n-layout-header>
        <n-layout-content style="margin-top: 10px;margin-right: 0px">
            <n-space justify="space-evenly">
                <n-form size="small"
                        :style="{maxWidth: '640px',minWidth:'320px'}">
                    
                        <n-form-item v-show="false" placeholder="name" label="id">
                            <n-input-number placeholder="id" v-model:value="data.id" :show-button="false"/>
                        </n-form-item>
                    
                    
                        <n-form-item label="展示值">
                            <n-input placeholder="label" v-model:value="data.label"/>
                        </n-form-item>
                    
                    
                        <n-form-item label="字典值">
                            <n-input-number placeholder="value" v-model:value="data.value" :show-button="false"/>
                        </n-form-item>
                    
                    
                        <n-form-item label="状态">
                            <n-switch v-model:value="data.status"/>
                        </n-form-item>

                    
                    
                        <n-form-item label="排序">
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
        </n-layout-content>
    </n-layout>
</template>
<script setup>
import {http} from "@/core/http";

const data = reactive({
});

const toEdit = async (id) => {
    if (id) {
        const apiRet = await http.post('/dictionary-detail/one', {'id': id})
        if (apiRet.ok) {
            Object.assign(data, apiRet.data);
        }
    } else {
        data.id = null
    }
}

const router = useRouter()
const submit = async () => {
    let dictionaryID = parseInt(router.currentRoute.value.params.id)
    if(NaN != dictionaryID){
        data.dictionaryID = dictionaryID
    }
    console.log(data)
    if (data.id) {
        const apiRet = await http.post('/dictionary-detail/edit', data)
        if (apiRet.ok) {
            backList()
        }
    } else {
        const apiRet = await http.post('/dictionary-detail/add', data)
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