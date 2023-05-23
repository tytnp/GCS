<template>
  <q-layout class=" column justify-center items-center">
    <div class="text-center text-primary q-gutter-md" style="width: 80%">
      <p class="text-h4">Welcome to GPT4.0</p>
      <!--      <p>{{ router.currentRoute.value.query }}</p>-->
      <!--      <q-input class="text-primary" label="Username" v-model="user.username"></q-input>-->
      <!--      <q-input label="Password" v-model="user.password"></q-input>-->
      <q-input style="margin-top: 100px" label="Invite Code" v-model="inviteCode"
               :rules="[val => !!val || errorMessage ]"
               ref="inputRef"
      >
      </q-input>
      <q-btn style="margin-top: 30px" @click="go" :loading="btnDisable">&nbsp; G o &nbsp;</q-btn>
    </div>
  </q-layout>
</template>
<script setup>
import {onMounted, ref, computed} from 'vue'
import {useRouter} from "vue-router";
import {api} from "boot/axios";

const router = useRouter()

const inputRef = ref(null)
const devId = router.currentRoute.value.query.devId
const inviteCode = ref("")
const errorMessage = ref("Invite Code is required")
const btnDisable = ref(false)

onMounted(async () => {
  btnDisable.value = true
  const apiRet = await api.post('/gpt-user/one', {'devId': devId})
  if (apiRet.ok && apiRet.data) {
    console.log("to chatgpt...")
    window.location.href = 'https://chat.openai.com/';
    btnDisable.value = false
  } else {
    btnDisable.value = false
  }
})

async function go() {
  btnDisable.value = true
  if (inputRef.value.validate()) {
    let syncAccountData = {
      "devId": devId,
      "agent": inviteCode.value
    }
    if (await syncAccount(syncAccountData)) {
      let localAccountData = {
        "devId": devId,
        "inviteCode": inviteCode.value
      }
      if (await addUser(localAccountData)) {
        console.log("to chatgpt...")
        window.location.href = 'https://chat.openai.com/';
      }
    }
  } else {
    console.log("to chatgpt...")
    window.location.href = 'https://chat.openai.com/';
  }
  setTimeout(() => {
    btnDisable.value = false
  }, 500);
}

async function addUser(data) {
  const apiRet = await api.post('/gpt-user/addExt', data)
  return apiRet.ok
}

async function syncAccount(data) {
  const apiRet = await api.post('/gpt-user/SyncAccount', data)
  if (apiRet.resultCode === 1) {
    return true
  } else {
    inviteCode.value = ""
    errorMessage.value = "inviteCode is invalid"
  }
  return false
}
</script>
