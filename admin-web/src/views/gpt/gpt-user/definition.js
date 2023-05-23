export const columnsMetaData = [
    {
        field:"id",
        fieldName:"编号",
        fieldType:"uint",
    },
    {
        field:"devId",
        fieldName:"设备编号",
        fieldType:"string",
    },
    {
        field:"username",
        fieldName:"登录名",
        fieldType:"string",
    },
    {
        field:"password",
        fieldName:"密码",
        fieldType:"string",
    },
    {
        field:"nick_name",
        fieldName:"昵称",
        fieldType:"string",
    },
    {
        field:"invite_code",
        fieldName:"邀请码",
        fieldType:"string",
    },
    {
        field:"enable",
        fieldName:"启用状态",
        fieldType:"bool",
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
        title:"设备编号",
        key:"devId",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"登录名",
        key:"userName",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"密码",
        key:"password",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"昵称",
        key:"nickName",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"邀请码",
        key:"inviteCode",
        ellipsis: {
            tooltip: true
        },
    },
    {
        title:"启用状态",
        key:"enable",
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