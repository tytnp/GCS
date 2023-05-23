export const columnsMetaData = [
    {
        field:"id",
        fieldName:"编号",
        fieldType:"uint",
    },
    {
        field:"name",
        fieldName:"字典名(中)",
        fieldType:"string",
    },
    {
        field:"type",
        fieldName:"字典名(英)",
        fieldType:"string",
    },
    {
        field:"status",
        fieldName:"状态",
        fieldType:"bool",
    },
    {
        field:"desc",
        fieldName:"描叙",
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
    manage:  (id) =>{},
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
            title:"字典名(中)",
            key:"name",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"字典名(英)",
            key:"type",
            ellipsis: {
                tooltip: true
            }
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
            title:"描叙",
            key:"desc",
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
                        onClick: () => caller.manage(row.id)
                    },
                    {default: () => '管理'}
                ),
                ' ',
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