<template>
  <n-card
      style="width: 1000px;"
      :bordered="false"
      size="huge"
      role="dialog"
      aria-modal="true"
  >
    <n-layout>
      <n-layout-header>
        <n-table size="small" :single-line="true" :single-column="true" :bordered="false">
          <tbody>
          <tr>
            <td colspan="2" style="padding-left: 0;padding-right: 0">
              <GcsSearch v-model:search-condition="tableProps.searchCondition"
                         v-model:search-condition-meta-data="props.columnsMetaData"
                         :setting-width="100"
                         :url-path="props.urlPath"
              />
            </td>
          </tr>
          <tr>
            <td style="padding: 0">
            </td>
            <td style="padding: 0;width: 100px">
              <n-button @click="tableProps.fetchData">查询</n-button>
            </td>
          </tr>
          </tbody>
        </n-table>
      </n-layout-header>
      <n-layout-content style="margin-top: 10px">
        <n-card content-style="padding:10px;">
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
    </n-layout>
  </n-card>
</template>
<style lang="less" scoped>
.table-dynamic-height {
  height: 400px;
}
</style>
<script setup>
import {http} from "@/core/http";

onMounted(() => {
  tableProps.fetchData()
})
//
const tableProps = reactive({
  searchCondition: [],
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
    const apiRet = await http.post(props.queryUrl, tableProps.params)
    tableProps.data = apiRet.data
    tableProps.pagination.page = apiRet.page
    tableProps.pagination.pageSize = apiRet.pageSize
    tableProps.pagination.itemCount = apiRet.totalCount
  }
})

const emit = defineEmits(['update:retVal', 'update:showModal'])
const props = defineProps({
  operateFunc: Function,
  columnsMetaData: {
    type: Array
  },
  columns: {
    type: Array
  },
  queryUrl: {
    type: String
  },
  retVal: {
    type: Number
  },
  showModal: Boolean,
  urlPath: {
    type: String
  }
})

const columns = reactive(props.columns);
onBeforeMount(() => {
  const exist = columns.find(item => {
    if (item.title == "操作") {
      return true
    }
  })
  if (exist) {
    columns.pop();
  }
  columns.push(
      {
        title: '操作',
        key: 'actions',
        align: "center",
        render(row) {
          return [
            h(
                NButton,
                {
                  strong: true,
                  tertiary: true,
                  size: 'small',
                  onClick: () => {
                    emit('update:retVal', row.id)
                    emit('update:showModal', false)
                    if (props.operateFunc) {
                      props.operateFunc(row.id)
                    }
                  }
                },
                {default: () => '选择'}
            ),
          ]
        }
      }
  )
})


</script>
