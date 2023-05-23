<template>
  <n-config-provider :theme="store.getTheme">
    <n-space vertical size="large">
      <n-layout>
        <n-layout-header bordered class="gcs-header">
          <GcsHeader/>
        </n-layout-header>
        <n-layout-content>
          <n-layout has-sider>
            <GcsMenu/>
            <n-layout-content style="background-color: #0a2d50;padding: 0px">
              <n-layout>
                <n-layout-header style="padding: 5px 5px 0px 5px ">
                  <GcsTabs></GcsTabs>
                </n-layout-header>
                <n-layout-content class="gcs-content">
                  <n-layout position="absolute"
                            content-style="padding: 0px 0px 0px 0px;">
                    <router-view v-slot="{Component, route}">
                      <keep-alive>
                        <component :is="Component" :key="route.path + (route.meta.keepAlive ? '' : '')"/>
                      </keep-alive>
                    </router-view>
                  </n-layout>
                </n-layout-content>
              </n-layout>
            </n-layout-content>
          </n-layout>
        </n-layout-content>
      </n-layout>
    </n-space>
  </n-config-provider>
</template>
<script setup>
import {useGlobalStore} from "@/stores/globalStore.js"

const store = useGlobalStore()

</script>

