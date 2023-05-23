export const columnsMetaData = [
    {
        field:"id",
        fieldName:"编号",
        fieldType:"uint",
    },
    {
        field:"parent_id",
        fieldName:"父菜单ID",
        fieldType:"string",
    },
    {
        field:"title",
        fieldName:"菜单名",
        fieldType:"string",
    },
    {
        field:"icon",
        fieldName:"菜单图标",
        fieldType:"string",
    },
    {
        field:"name",
        fieldName:"路由名",
        fieldType:"string",
    },
    {
        field:"path",
        fieldName:"路由路径",
        fieldType:"string",
    },
    {
        field:"component",
        fieldName:"组件路径",
        fieldType:"string",
    },
    {
        field:"hidden",
        fieldName:"是否隐藏",
        fieldType:"bool",
    },
    {
        field:"sort",
        fieldName:"排序标记",
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
            }
        },
        {
            title:"父菜单ID",
            key:"parentId",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"菜单名",
            key:"title",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"菜单图标",
            key:"icon",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"路由名",
            key:"name",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"路由路径",
            key:"path",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"组件路径",
            key:"component",
            ellipsis: {
                tooltip: true
            }
        },
        {
            title:"是否隐藏",
            key:"hidden",
            ellipsis: {
                tooltip: true
            },
            render(row) {
                if (row.hidden) {
                    return '是'
                }
                return '否'
            }
        },
        {
            title:"排序标记",
            key:"sort",
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