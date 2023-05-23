<template>
    <div id="userLayout">
        <div class="login_panel">
            <div class="login_panel_form">
                <div class="login_panel_form_title">
                    <img
                            class="login_panel_form_title_logo"
                            src="@/assets/images/logo.png"
                            alt
                    >
                    <p class="login_panel_form_title_p">GCS-Admin</p>
                </div>
                <n-form size="small" ref="loginFormRef" :model="loginData" :rules="rules"
                        @keyup.enter.native="submitLogin">
                    <n-form-item label-placement="left" path="loginName">
                        <n-input
                                v-model:value="loginData.loginName"
                                size="large"
                                placeholder="请输入用户名"
                        />
                    </n-form-item>
                    <n-form-item label-placement="left" path="loginPwd">
                        <n-input
                                v-model:value="loginData.loginPwd"
                                type="password"
                                size="large"
                                show-password-on="click"
                                placeholder="请输入密码"
                        />
                    </n-form-item>
                    <n-form-item>
                        <n-button
                                type="primary"
                                size="large"
                                style="width: 50%; margin-left: 25%"
                                @click="submitLogin"
                        >登 录
                        </n-button>
                    </n-form-item>
                </n-form>
            </div>
            <div class="login_panel_right"/>
            <div class="login_panel_foot">
                <div class="links">
                    <img src="@/assets/docs.png" class="link-icon" alt="文档">
                    <img src="@/assets/kefu.png" class="link-icon" alt="客服">
                    <img src="@/assets/github.png" class="link-icon" alt="github">
                    <img src="@/assets/video.png" class="link-icon" alt="视频站">
                </div>
                <div class="copyright">
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import {useMessage} from "naive-ui";
import {http} from "@/core/http/index.js";
import {useUserStore} from "@/stores/userStore.js"
import {useRouter} from "vue-router";

const loginFormRef = ref(null);
const loginData = reactive({
    loginName: 'admin',
    loginPwd: 'admin',
    userType: 1,
})
const rules = {
    loginName: {
        required: true,
        message: "请输入用户名",
        trigger: ['input']
    },
    loginPwd: {
        required: true,
        message: "请输入密码",
        trigger: ['input']
    }
}
const store = useUserStore()
const message = useMessage()
const router = useRouter()
const submitLogin = () => {
    loginFormRef.value?.validate(async (errors) => {
        if (!errors) {
            const apiRet = await http.post('/login', loginData)
            if(apiRet.ok){
                store.setToken(apiRet.data)
                router.push({ path: '/' })
            }else{
                message.error('登录失败,信息错误')
            }
        } else {
            message.warning('请完成输入')
        }
    });
}
</script>
<style lang="less" scoped>
@import "@/config/style/newLogin.less";
</style>
