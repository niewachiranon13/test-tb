<template>
  <div class="welcome-container">
    <div class="header">IT 02-3</div>
    <div class="welcome">Welcome User: {{ username }}</div>
  </div>
</template>

<script setup>
import { defineProps } from 'vue'
const props = defineProps({ username: String })
</script>

<style scoped>
.welcome-container {
  max-width: 400px;
  margin: 40px auto;
  border: 1px solid #ccc;
  padding: 32px;
  border-radius: 8px;
  background: #fff;
}
.header {
  background: #22a322;
  color: #fff;
  padding: 8px;
  text-align: center;
  margin-bottom: 24px;
  font-weight: bold;
}
.welcome {
  text-align: center;
  font-size: 1.2em;
  margin-top: 32px;
}
</style>
