import {caller} from "@/views/system/role/definition.js";

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
    }
]