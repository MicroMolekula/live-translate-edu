<script setup>


const props = defineProps({
  text: {
    type: String,
  }
})
</script>

<template>
<div class="w-[35rem] p-4 bg-white rounded text-black">
  {{ props.text ? props.text : "Место для субтитров..." }}
</div>
</template>

<style scoped>

</style>