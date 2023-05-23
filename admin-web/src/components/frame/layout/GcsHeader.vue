<template>
    <n-layout has-sider>
      <n-layout-sider class="gcs-logo" width="var(--left-width)" content-style="padding-left: 30px;">
        <n-space align="center" class="gcs-logo-content">
          <img src="@/assets/images/logo.png" width="32"  height="32">
          <h3>Gcs Platform</h3>
        </n-space>
      </n-layout-sider>
      <n-layout-content>
        <n-grid x-gap="0" cols="2" class="gcs-header-content">
          <n-gi></n-gi>
          <n-gi style="display: flex;align-items: center;justify-content: end;padding-right: 30px">
              <n-button @click="switchTheme" dashed>
                <template #icon>
                  <n-icon>
                    <SettingsAutomation/>
                  </n-icon>
                </template>
                {{ globalStore.themeName }}
              </n-button>
            <n-divider vertical/>
            <n-dropdown :options="options" @select="handleSelect">
              <div style="display: flex;align-items: center;justify-content: end">
                <h4 style="margin-right: 10px;margin-left: 10px">{{user.nickName}}</h4>
                <n-avatar round size="medium">A</n-avatar>
              </div>
            </n-dropdown>
          </n-gi>
        </n-grid>
      </n-layout-content>
    </n-layout>
</template>
<script setup>
import {useUserStore} from "@/stores/userStore.js";
import {useGlobalStore} from "@/stores/globalStore.js"

const globalStore = useGlobalStore()
import {useRouter} from "vue-router";

const switchTheme = () =>{
  globalStore.setTheme()
}

const store = useUserStore()
const router = useRouter()

const user = reactive({
    nickName:null
})

onMounted(async ()=>{
    await store.asyncUserInfo()
    user.nickName =  store.user.nickName;
})

const options = [
  // {
  //   label: '个人资料',
  //   key: 'user-info',
  // },
  // {
  //   label: '修改密码',
  //   key: "update-pwd"
  // },
  {
    label: '退出',
    key: 'logout'
  }
]

const handleSelect = (key) =>{
    if(key == 'logout'){
        store.setToken(null)
        router.push({ path: '/login' })
    }
}

</script>
<style lang="less" scoped>

</style>