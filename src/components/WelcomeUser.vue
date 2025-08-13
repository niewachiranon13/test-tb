<template>
     <div class="header">IT 02-3</div>
  <div class="welcome-container">

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
  width: 100vw;
  background: #22a322;
  color: #fff;
  padding: 16px 0;
  text-align: center;
  font-weight: bold;
  position: fixed;
  top: 0;
  left: 0;
  z-index: 1000;
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
}
.welcome {
  text-align: center;
  font-size: 1.2em;
  margin-top: 32px;
}
</style>
