<script setup>


const props = defineProps({
  text: {
    type: String,
  }
})
</script>

<template>
<div class="w-[40rem] p-4 text-2xl bg-white rounded text-black">
  {{ props.text ? props.text : "Место для субтитров..." }}
</div>
</template>

<style scoped>

</style>