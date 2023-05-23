<template>
    <n-space vertical size="large">
        <n-layout>
            <GcsHeader/>
            <n-layout has-sider>
                <GcsMenu/>
                <n-layout-content style="background-color: #f5f7f9">
                    <n-layout position="absolute" style="background-color: #f5f7f9">
<!--                        <n-layout-header position="absolute" bordered style="padding: 0px 0px 0px 0px">-->
<!--                            <GcsTabs></GcsTabs>-->
<!--                        </n-layout-header>-->
                        <n-layout has-sider
                                  position="absolute"
                                  :native-scrollbar="false"
                                  content-style="padding: 10px"
                              >
                            <router-view></router-view>
                        </n-layout>
                    </n-layout>
                </n-layout-content>
            </n-layout>
        </n-layout>
    </n-space>
</template>
<script lang="ts" setup>
</script>

