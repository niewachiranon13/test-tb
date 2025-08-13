<template>
  <div class="header">IT 02-1 </div>
  <div class="login-container">
    <form @submit.prevent="onLogin">
      <div>
        <label>User</label>
        <input v-model="loginForm.username" required />
      </div>
      <div>
        <label>Password</label>
        <input v-model="loginForm.password" type="password" required />
      </div>
      <button type="submit">ลงชื่อเข้าใช้งาน</button>
      <div v-if="loginError" class="error">{{ loginError }}</div>
    </form>
    <router-link to="/register" class="register-link">สมัครสมาชิก</router-link>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'


const emit = defineEmits(['login-success'])
const loginForm = ref({ username: '', password: '' })
const loginError = ref('')
const router = useRouter()

async function onLogin() {
  loginError.value = ''
  try {
 
    const payload = {
      username: loginForm.value.username,
  password: btoa(loginForm.value.password)
    }
    const res = await fetch('/api/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(payload)
    })
    if (!res.ok) throw new Error('Login failed')
    const data = await res.json()
    localStorage.setItem('token', data.token)
    emit('login-success', { username: data.username, token: data.token })
  } catch (e) {
    loginError.value = 'ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง'
  }
}
</script>

<style scoped>
.login-container {
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
.login-container {
  max-width: 400px;
  margin: 80px auto 40px auto;
  border: 1px solid #ccc;
  padding: 32px;
  border-radius: 8px;
  background: #fff;
}
form > div {
  margin-bottom: 16px;
}
input {
  width: 100%;
  padding: 8px;
  box-sizing: border-box;
}
button {
  background: #22a322;
  color: #fff;
  border: none;
  padding: 8px 16px;
  border-radius: 4px;
  cursor: pointer;
}
.error {
  color: red;
  margin-top: 8px;
}
.register-link {
  display: block;
  margin-top: 12px;
  color: #22a322;
  cursor: pointer;
  text-align: right;
  text-decoration: underline;
}
</style>
