<script setup>

import {onMounted} from "vue";
import {useRouter} from "vue-router";

const router = useRouter()

async function getUser() {

}

onMounted(async function () {
  if (localStorage.getItem("jwt") === null) {
    await router.push('/auth')
  }
  let response = await fetch('http://localhost:8080/api/me', {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('jwt')
    },
  })
  if (response.status !== 200) {
    await router.push('/auth')
  }
})
</script>

<template>
<router-view></router-view>
</template>

<style scoped>

</style>
