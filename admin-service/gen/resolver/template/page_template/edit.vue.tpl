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
                {{- range .Columns }}
                    {{if eq .JsonName "id"}}
                        <n-form-item v-show="false" placeholder="name" label="id">
                            <n-input-number placeholder="id" v-model:value="data.id" :show-button="false"/>
                        </n-form-item>
                    {{else if eq .Type "int"}}
                        <n-form-item label="{{ .AliasName }}">
                            <n-input-number placeholder="{{ .JsonName }}" v-model:value="data.{{ .JsonName }}" :show-button="false"/>
                        </n-form-item>
                    {{else if eq .Type "bool"}}
                        <n-form-item label="{{ .AliasName }}">
                            <n-switch v-model:value="data.{{ .JsonName }}"/>
                        </n-form-item>

                    {{else}}
                        <n-form-item label="{{ .AliasName }}">
                            <n-input placeholder="{{ .JsonName }}" v-model:value="data.{{ .JsonName }}"/>
                        </n-form-item>
                    {{end}}
                {{- end }}
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
        const apiRet = await http.post('/{{.Target.Methods.ModelName}}/one', {'id': id})
        if (apiRet.ok) {
            Object.assign(data, apiRet.data);
        }
    } else {
        data.id = null
    }
}

const submit = async () => {
    if (data.id) {
        const apiRet = await http.post('/{{.Target.Methods.ModelName}}/edit', data)
        if (apiRet.ok) {
            backList()
        }
    } else {
        const apiRet = await http.post('/{{.Target.Methods.ModelName}}/add', data)
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