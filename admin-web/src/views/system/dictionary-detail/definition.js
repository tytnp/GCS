export const columnsMetaData = [
    {
        field:"id",
        fieldName:"编号",
        fieldType:"uint",
    },
    {
        field:"label",
        fieldName:"展示值",
        fieldType:"string",
    },
    {
        field:"value",
        fieldName:"字典值",
        fieldType:"int",
    },
    {
        field:"status",
        fieldName:"状态",
        fieldType:"bool",
    },
    {
        field:"sort",
        fieldName:"排序",
        fieldType:"int",
    },
    {
        field:"created_at",
        fieldName:"创建时间",
        fieldType:"time.Time",
    },
    {
        field:"updated_at",
        fieldName:"修改时间",
        fieldType:"time.Time",
    },
]

export const caller = {
    edit: (id) =>{},
    remove: (id) =>{},
}

export const columns  = [
    {
        title:"编号",
        key:"id",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"展示值",
        key:"label",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"字典值",
        key:"value",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"状态",
        key:"status",
        ellipsis: {
            tooltip: true
        },
        render(row) {
            if(row.status){
                return '是'
            }
            return '否'
        }
    },
    {
        title:"排序",
        key:"sort",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"创建时间",
        key:"createdAt",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"修改时间",
        key:"updatedAt",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title: '操作',
        key: 'actions',
        align: "center",
        render (row) {
            return [
                h(
                    NButton,
                    {
                        strong: true,
                        tertiary: true,
                        size: 'small',
                        onClick: () => caller.edit(row.id)
                    },
                    { default: () => '编辑' }
                ),
                ' ',
                h(
                    NButton,
                    {
                        strong: true,
                        tertiary: true,
                        size: 'small',
                        onClick: () => caller.remove(row.id)
                    },
                    { default: () => '删除' }
                )
            ]
        }
    }
]