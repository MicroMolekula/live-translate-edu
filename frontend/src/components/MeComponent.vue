<script setup>
async function getUser() {
  let response = await fetch('http://localhost:8080/api/me', {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('jwt')
    },
  })
  let result = await response.json()
  await console.log(result)
}
</script>

<template>
<button @click="getUser">Тест авторизации</button>
</template>

<style scoped>

</style>