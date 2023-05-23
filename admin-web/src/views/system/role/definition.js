export const columnsMetaData = [
    {
        field: "id",
        fieldName: "编号",
        fieldType: "uint",
    },
    {
        field: "role_name",
        fieldName: "角色名",
        fieldType: "string",
    },
    {
        field: "default_router",
        fieldName: "默认菜单",
        fieldType: "string",
    },
    {
        field: "created_at",
        fieldName: "创建时间",
        fieldType: "time.Time",
    },
    {
        field: "updated_at",
        fieldName: "修改时间",
        fieldType: "time.Time",
    },
]

export const caller = {
    edit: (id) => {
    },
    remove: (id) => {
    },
    auth: (id) => {
    },
}

export const columns = [
    {
        title: "编号",
        key: "id",
        ellipsis: {
            tooltip: true
        }
    },
    {
        title: "角色名",
        key: "roleName",
        ellipsis: {
            tooltip: true
        }
    },
    {
        title: "默认菜单",
        key: "defaultRouter",
        ellipsis: {
            tooltip: true
        }
    },
    {
        title: "创建时间",
        key: "createdAt",
        ellipsis: {
            tooltip: true
        }
    },
    {
        title: "修改时间",
        key: "updatedAt",
        ellipsis: {
            tooltip: true
        }
    },
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
                        onClick: () => caller.auth(row.id)
                    },
                    {default: () => '授权'}
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
                    {default: () => '编辑'}
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
                    {default: () => '删除'}
                )
            ]
        }
    }
]