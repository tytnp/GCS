<template>
    <n-form-item size="large">
        <n-input v-model:value="filterText" placeholder="筛选" style="margin-right: 20px" size="medium"></n-input>
        <n-button @click="saveMenus">确 定</n-button>
    </n-form-item>
    <pre></pre>
    <!--    {{JSON.stringify(tree,null,2)}}-->
    <!--    {{ checkedKeys }}-->
    <!--    {{props.roleId}}-->
    <n-tree
            :show-irrelevant-nodes="false"
            block-line
            check-on-click
            default-expand-all
            checkable
            ref="treeRef"
            :pattern="filterText"
            :data="tree"
            :checked-keys="checkedKeys"
            @updateCheckedKeys="updateCheckedKeys"
    >
    </n-tree>
</template>

<script setup>
import {http} from "@/core/http/index.js";
import {useMessage} from "naive-ui";

const filterText = ref('')
const tree = ref([])
const treeRef = ref(null)
const checkedKeys = ref([])

const updateCheckedKeys = (keys, option, meta, action) => {
    //console.log("keys", keys, "option", option, "meta", meta, "action", action)
    if(meta.action==='check'){
        checkedKeys.value.push(meta.node.key)
        if(meta.node.children){
            meta.node.children.map(node=>{
                checkedKeys.value.push(node.key)
            })
        }
    }else{
        checkedKeys.value = checkedKeys.value.filter(function(item) {
            return item != meta.node.key;
        });
        if(meta.node.children){
            meta.node.children.map(node=>{
                checkedKeys.value = checkedKeys.value.filter(function(item) {
                    return item != node.key;
                });
            })
        }
    }
}

const message = useMessage()
const saveMenus = async () => {
    const apiRet = await http.post("/role/grantMenu", {roleId:props.roleId,menuIds:checkedKeys.value})
    // console.log(JSON.stringify(apiRet,null,2))
    // console.log("CheckedData", treeRef.value.getCheckedData())
    if(apiRet.ok){
        message.info("授权完成")
        emit('update:showDrawer', false)
    }

}

onMounted(async () => {
    let m = await http.post('/menu/GetMenuTreeByRole',{'id':props.roleId})
    if(m.ok){
        m.data.map(item=>{
            checkedKeys.value.push(item.id)
        })
    }
    let res = await http.post('/menu/GetMenuTree')
    tree.value = formatTree(res).children
})

const formatTree = (data) => {
    let treeRoot = {
        key: data.id,
        label: data.title,
        prefix: data.id,
    }
    if (data.children && data.children.length > 0) {
        treeRoot.children = []
        for (const dataKey in data.children) {
            treeRoot.children.push(formatTree(data.children[dataKey]))
        }
    }
    return treeRoot;
}

const emit = defineEmits(['update:showDrawer'])
const props = defineProps({
    roleId:{
        required:true,
        type:Number
    },
    showDrawer:Boolean
})
</script>

