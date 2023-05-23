export const columnsMetaData = [
    {{- range .ColumnsMetaData }}
    {
        field:"{{ .JsonName }}",
        fieldName:"{{ .AliasName }}",
        fieldType:"{{ .Type }}",
    },
    {{- end }}
]

export const caller = {
    edit: (id) =>{},
    remove: (id) =>{},
}

export const columns  = [
    {{- range .Columns }}
    {
        title:"{{ .AliasName }}",
        key:"{{ .JsonName }}",
        ellipsis: {
            tooltip: true
        },
        {{- if eq .Type "bool"}}
        render(row) {
            if(row.status){
                return '是'
            }
            return '否'
        }
        {{- end}}
    },
    {{- end }}
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