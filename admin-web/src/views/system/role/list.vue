<template>
    <n-layout>
        <n-layout-header ref="header" style="padding-left: 5px">
            <n-table size="small" :single-line="true" :single-column="true" :bordered="false">
                <tbody>
                <tr>
                    <td colspan="2" style="padding-left: 0;padding-right: 0">
                        <GcsSearch v-model:search-condition="tableProps.searchCondition"
                                   v-model:search-condition-meta-data="columnsMetaData"/>
                    </td>
                </tr>
                <tr>
                    <td style="padding: 0">
                        <n-button @click="add">新增</n-button>
                    </td>
                    <td style="padding: 0;width: 400px">
                        <n-button @click="tableProps.fetchData">查询</n-button>
                    </td>
                </tr>
                </tbody>
            </n-table>
        </n-layout-header>
        <n-layout-content style="margin-top: 0px">
            <n-card content-style="padding:5px;">
                <n-data-table
                        class="table-dynamic-height"
                        size="small"
                        :columns="columns"
                        :data="tableProps.data"
                        :pagination="tableProps.pagination"
                        flex-height
                        bordered
                        :single-line="false"
                        remote
                />
            </n-card>
        </n-layout-content>
        <n-drawer v-model:show="showDrawer" :width="502" placement="right">
            <n-drawer-content >
                <n-tabs default-value="menu" type="card">
                    <n-tab-pane name="menu" tab="菜 单">
<!--                        <MenuTree :role-id="roleId"></MenuTree>-->
                        <menu-tree :role-id="roleId" v-model:showDrawer="showDrawer"></menu-tree>
                    </n-tab-pane>
                    <n-tab-pane name="api" tab="接 口">
                        接口
                    </n-tab-pane>
                </n-tabs>
            </n-drawer-content>
        </n-drawer>
    </n-layout>
</template>
<style lang="less" scoped>
.table-dynamic-height {
  height: var(--table-dynamic-height);
}
</style>
<script setup>
import {http} from "@/core/http";
import {columnsMetaData, columns, caller} from "./definition.js";
import {useEventListener, useResizeObserver, useWindowSize} from "@vueuse/core";
import { useMessage } from 'naive-ui'
import MenuTree from "./components/MenuTree.vue"


const showDrawer = ref(false)
// const drawerPlacement = ref('right')
const roleId = ref(0)

const header = ref(null);
const scope = effectScope()
scope.run(() => {
    const dynamicHeightUpdate = () => {
        let top = parseFloat(getComputedStyle(document.documentElement).getPropertyValue('--header-height'))
        let dynamicHeight = useWindowSize().height.value - header.value.$el.offsetHeight - top - 56;
        if (dynamicHeight < 300) {
            requestAnimationFrame(()=>{
                document.documentElement.style.setProperty('--table-dynamic-height', 300 + 'px');
            })
        } else {
            requestAnimationFrame(()=>{
                document.documentElement.style.setProperty('--table-dynamic-height', dynamicHeight + 'px');
            })
        }
    }
    useResizeObserver(header, () => {
        dynamicHeightUpdate()
    })
    useEventListener('resize', () => {
        dynamicHeightUpdate()
    })
})

onMounted(() => {
    tableProps.fetchData()
})

const tableProps = reactive({
    searchCondition: [

    ],
    params: {
        expandProps: {
            "page": 1,
            "pageSize": 15,
            "searchCondition": []
        }
    },
    data: [],
    pagination: {
        displayOrder: ['quick-jumper', 'pages', 'size-picker'],
        showQuickJumper: true,
        showSizePicker: true,
        pageSizes: [
            {label: '15/页', value: 15},
            {label: '50/页', value: 50},
            {label: '100/页', value: 100},
            {label: '500/页', value: 500},
        ],
        onUpdatePage: (page) => {
            if (tableProps.params.expandProps) {
                tableProps.params.expandProps["page"] = page
                tableProps.fetchData()
            }
        },
        onUpdatePageSize: (pageSize) => {
            if (tableProps.params.expandProps) {
                tableProps.params.expandProps["page"] = 1
                tableProps.params.expandProps["pageSize"] = pageSize
                tableProps.fetchData()
            }
        }
    },
    fetchData: async () => {
        tableProps.params.expandProps["searchCondition"] = tableProps.searchCondition
        const apiRet = await http.post('/role/list', tableProps.params)
        tableProps.data = apiRet.data
        tableProps.pagination.page = apiRet.page
        tableProps.pagination.pageSize = apiRet.pageSize
        tableProps.pagination.itemCount = apiRet.totalCount
    }
})

const add = () => {
    emits('activeMethod', 'form');
};

const message = useMessage()
caller.edit = (id) => {
    emits('activeMethod', 'form',id);
};

caller.remove = async (id) => {
    const apiRet = await http.post('/role/del',{'id':id})
    if(apiRet.ok){
        tableProps.fetchData()
        message.info("删除成功!")
    }
}
caller.auth = async (id) => {
    roleId.value = id
    showDrawer.value=true
}
const emits = defineEmits(['activeMethod']);

defineExpose({
    tableProps
})
</script>
