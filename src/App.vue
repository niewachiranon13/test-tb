<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

function onLoginSuccess({ username, token: jwt }) {
  localStorage.setItem('token', jwt)
  router.push({ path: '/welcome', query: { username } })
}
</script>

<template>
  <router-view @login-success="onLoginSuccess" />
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
