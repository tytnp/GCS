export const columnsMetaData = [
    {
        field: "id",
        fieldName: "编号",
        fieldType: "uint",
    },
    {
        field: "username",
        fieldName: "登录名",
        fieldType: "string",
    },
    {
        field: "password",
        fieldName: "密码",
        fieldType: "string",
    },
    {
        field: "nick_name",
        fieldName: "昵称",
        fieldType: "string",
    },
    {
        field: "phone",
        fieldName: "手机",
        fieldType: "string",
    },
    {
        field: "email",
        fieldName: "邮箱",
        fieldType: "string",
    },
    {
        field: "has_super_admin",
        fieldName: "是否超管",
        fieldType: "bool",
    },
    {
        field: "enable",
        fieldName: "启用状态",
        fieldType: "bool",
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
        title: "登录名",
        key: "userName",
        ellipsis: {
            tooltip: true
        }
    },
    // {
    //     title: "密码",
    //     key: "password",
    //     ellipsis: {
    //         tooltip: true
    //     }
    // },
    {
        title: "昵称",
        key: "nickName",
        ellipsis: {
            tooltip: true
        }
    },
    // {
    //     title: "手机",
    //     key: "phone",
    //     ellipsis: {
    //         tooltip: true
    //     }
    // },
    // {
    //     title: "邮箱",
    //     key: "email",
    //     ellipsis: {
    //         tooltip: true
    //     }
    // },
    // {
    //     title: "是否超管",
    //     key: "hasSuperAdmin",
    //     ellipsis: {
    //         tooltip: true
    //     }
    // },
    {
        title: "启用状态",
        key: "enable",
        ellipsis: {
            tooltip: true
        },
        render(row) {
            if (row.enable) {
                return '是'
            }
            return '否'
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
        title: "角色",
        key: "role",
        ellipsis: {
            tooltip: true
        },
        render(row) {
            if (row.roles.length > 0) {
                return row.roles[0].roleName
            }
            return '无'
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