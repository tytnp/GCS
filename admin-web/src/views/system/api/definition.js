export const columnsMetaData = [
    {
        field:"id",
        fieldName:"编号",
        fieldType:"uint",
    },
    {
        field:"path",
        fieldName:"Api路径",
        fieldType:"string",
    },
    {
        field:"description",
        fieldName:"api描述",
        fieldType:"string",
    },
    {
        field:"api_group",
        fieldName:"api组",
        fieldType:"string",
    },
    {
        field:"method",
        fieldName:"api方法",
        fieldType:"string",
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
            }
        },
        {
            title:"Api路径",
            key:"path",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"api描述",
            key:"description",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"api组",
            key:"apiGroup",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"api方法",
            key:"method",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"创建时间",
            key:"createdAt",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"修改时间",
            key:"updatedAt",
            ellipsis: {
                tooltip: true
            }
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