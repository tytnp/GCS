<template>
    <n-layout-sider
            class="gcs-menu"
            :width="menuWidth"
            collapse-mode="width"
            :collapsed-width="64"
            :collapsed="collapsed"
            @collapse="collapsed = true"
            @expand="collapsed = false"
            :inverted="inverted"
            show-trigger="arrow-circle"
            bordered>
        <n-menu class="gcs-menu-content"
                ref="menuRef"
                :options="menuTree"
                :inverted="inverted"
                :collapsed-width="64"
                :collapsed-icon-size="20"
                :collapsed="collapsed"
                :root-indent="22"
                :indent="15"
                accordion
                v-model:value="selectedKey"
        />
    </n-layout-sider>


</template>

<script setup>
import {useWindowSize} from "@vueuse/core";
import {useUserStore} from "@/stores/userStore.js"

const menuWidth = computed(() => {
    return parseFloat(getComputedStyle(document.documentElement).getPropertyValue('--left-width'));
})

const inverted = ref(false)
const collapsed = ref(false)
const menuRef = ref(null);
const menuTree = ref([])
const selectedKey = ref(null);

const {width} = useWindowSize();
const router = useRouter()
const store = useUserStore()
onMounted(async () => {
    if (width.value <= 500) {
        collapsed.value = true;
    }
    let data = await store.getMenuTree
    menuTree.value = data;
    selectedKey.value = router.currentRoute.value.name;
    await nextTick(()=>{
        menuRef.value?.showOption();
    })
})



// const menuData = reactive({
//     menuTree: []
// })

// const menuOptions = [
//     {
//         label: '系统管理',
//         key: 'system',
//         icon: renderIcon(HomeOutline),
//         children: [
//             {
//                 label: () => h(
//                     RouterLink,
//                     {
//                         to: {
//                             path: '/system/home'
//                         }
//                     },
//                     {
//                         default: () => "主页"
//                     }),
//                 key: 'dictionary',
//                 icon: renderIcon(HomeOutline),
//             },
//             {
//                 label: () => h(
//                     RouterLink,
//                     {
//                         to: {
//                             path: '/system/dictionary'
//                         }
//                     },
//                     {
//                         default: () => "字典管理"
//                     }),
//                 key: 'home',
//                 icon: renderIcon(HomeOutline),
//             },
//             {
//                 label: () => h(
//                     RouterLink,
//                     {
//                         to: {
//                             path: '/system/role'
//                         }
//                     },
//                     {
//                         default: () => "角色管理"
//                     }),
//                 key: 'role',
//                 icon: renderIcon(HomeOutline),
//             },
//             {
//                 label: () => h(
//                     RouterLink,
//                     {
//                         to: {
//                             path: '/system/menu'
//                         }
//                     },
//                     {
//                         default: () => "角色管理"
//                     }),
//                 key: 'role',
//                 icon: renderIcon(HomeOutline),
//             }
//         ]
//     },
//
// ]
</script>

<style lang="less" scoped>
@menuWidth: v-bind(menuWidth);
.gcs-menu {
  &-content {
    width: @menuWidth;
  }
}
</style>