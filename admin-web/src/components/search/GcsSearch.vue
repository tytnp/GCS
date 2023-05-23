<template>
    <div style="display: flex">
        <div style="flex-grow: 1;padding-right: 100px">
            <n-grid cols="3" x-gap="50" y-gap="8">
                <n-gi v-for="(searchCondition,index) in searchCondition">
                    <n-input-group>
                        <!--<n-input-group-label>{{ searchCondition.fieldName }}</n-input-group-label>-->
                        <!--<n-input v-model:value="searchCondition.value" :placeholder="searchCondition.field"></n-input>-->
                        <n-input v-model:value="searchCondition.value"
                                 :placeholder="searchCondition.fieldName"></n-input>
                        <n-tooltip trigger="hover">
                            <template #trigger>
                                <n-input-group-label>
                                    <component :is="dynamicCondition.renderIcon(searchCondition.condition)"/>
                                </n-input-group-label>
                            </template>
                            {{
                            dynamicCondition.conditionOptions.find(option => option.value === searchCondition.condition).label
                            }}
                        </n-tooltip>
                        <n-tooltip trigger="hover">
                            <template #trigger>
                                <n-button @click="dynamicCondition.removeCondition(index)" circle>
                                    <n-icon>
                                        <Close></Close>
                                    </n-icon>
                                </n-button>
                            </template>
                            移除
                        </n-tooltip>
                    </n-input-group>
                </n-gi>
            </n-grid>
            <!--            <pre>{{ JSON.stringify(dynamicCondition, null, 2) }}</pre>-->
        </div>
        <div style="flex-shrink: 0" :style="{'width':props.settingWidth+'px'}">
            <n-popover trigger="click" placement="bottom-end">
                <template #trigger>
                    <n-button>
                        <template #icon>
                            <n-icon>
                                <VirusSearch/>
                            </n-icon>
                        </template>
                    </n-button>
                </template>
                <div>
                    <n-grid cols="3" x-gap="10">
                        <n-gi>
                            <n-popselect v-model:value="dynamicCondition.fieldValue"
                                         :options="dynamicCondition.fieldOptions" size="large" scrollable>
                                <n-button>
                                    {{ dynamicCondition.fieldLabel() }}
                                </n-button>
                            </n-popselect>
                        </n-gi>
                        <n-gi>
                            <n-popselect v-model:value="dynamicCondition.conditionValue"
                                         :options="dynamicCondition.conditionOptions" size="large">
                                <n-button>
                                    {{ dynamicCondition.conditionLabel() }}
                                </n-button>
                            </n-popselect>
                        </n-gi>
                        <n-gi>
                            <n-button @click="dynamicCondition.addCondition">
                                <template #icon>
                                    <n-icon>
                                        <SearchAdvanced/>
                                    </n-icon>
                                </template>
                                添加
                            </n-button>
                        </n-gi>
                    </n-grid>
                </div>
            </n-popover>
        </div>
    </div>
</template>
<script setup>
import {conditionOptions} from "@/core/entity/search_condition.js";
import {useMessage} from "naive-ui";
import {useSearchStore} from '@/stores/searchStore.js'

const store = useSearchStore()

const dynamicCondition = ref({
    fieldValue: '',
    fieldLabel: () => {
        let result = dynamicCondition.value.fieldOptions.find(option => option.value === dynamicCondition.value.fieldValue)
        return result ? result.label : '字段'
    },
    fieldOptions: [],
    conditionValue: 0,
    conditionLabel: () => {
        let result = dynamicCondition.value.conditionOptions.find(option => option.value === dynamicCondition.value.conditionValue)
        return result ? result.label : '条件'
    },
    renderIcon: (conditionValue) => {
        switch (conditionValue) {
            case 1:
                return h(NIcon, null, {default: () => h(EqualsOutlined)})
            case 2:
                return h(NIcon, null, {default: () => h(NotEqualOutlined)})
            case 3:
                return h(NIcon, null, {default: () => h(GreaterThanOutlined)})
            case 4:
                return h(NIcon, null, {default: () => h(GreaterThanEqualOutlined)})
            case 5:
                return h(NIcon, null, {default: () => h(LessThanOutlined)})
            case 6:
                return h(NIcon, null, {default: () => h(LessThanEqualOutlined)})
            case 7:
                return h(NIcon, null, {default: () => h(PercentOutlined)})
        }
    },
    conditionOptions: conditionOptions,
    addCondition: () => {
        if (dynamicCondition.value.fieldValue != '' && dynamicCondition.value.conditionValue != '' && props.searchCondition) {
            const exists = props.searchCondition.find((value) => {
                if (value.field === dynamicCondition.value.fieldValue &&
                    value.condition === dynamicCondition.value.conditionValue
                ) {
                    return true;
                }
            })
            if (!exists) {
                let condition = {
                    field: dynamicCondition.value.fieldValue,
                    fieldName: dynamicCondition.value.fieldLabel(),
                    fieldType: dynamicCondition.value.fieldOptions.find(option => option.value === dynamicCondition.value.fieldValue).type,
                    condition: dynamicCondition.value.conditionValue,
                    value: ''
                }
                props.searchCondition.push(
                    condition
                )
                store.setSearchCondition(urlPath, condition)
            } else {
                message.info("条件重复!")
            }
        }
    },
    removeCondition: (index) => {
        if (props.searchCondition) {
            store.removeSearchCondition(urlPath, props.searchCondition.at(index), index)
            props.searchCondition.splice(index, 1)
        }
    }
})
const message = useMessage()
const props = defineProps(
    {
        searchCondition: {
            type: Array,
            required: true,
        },
        searchConditionMetaData: {
            type: Array,
            required: true,
        },
        settingWidth: {
            type: Number,
            default: () => {
                return 400
            }
        },
        urlPath: {
            type: String
        }
    }
)

let urlPath = undefined;
onMounted(() => {
    urlPath = useRoute().path
    if(props.urlPath){
        urlPath = props.urlPath
    }
    let storeSearch = store.getSearchCondition
    let result = storeSearch.find(item => {
        if (item && item[urlPath]) {
            return true
        }
    })
    if (result) {
        // console.log(JSON.stringify(result[urlPath]))
        props.searchCondition.push(...result[urlPath])
    }
    dynamicCondition.value.fieldOptions = props.searchConditionMetaData?.map(item => ({
        label: item.fieldName,
        value: item.field,
        type: item.fieldType
    }));
})

</script>