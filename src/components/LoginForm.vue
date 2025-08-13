<template>
  <div class="login-container">
    <div class="header">IT 02-1</div>
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
      password: loginForm.value.password
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
  background: #22a322;
  color: #fff;
  padding: 8px;
  text-align: center;
  margin-bottom: 24px;
  font-weight: bold;
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
