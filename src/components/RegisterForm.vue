<template>
   <div class="header">IT 02-2</div>
  <div class="register-container">
   
    <form @submit.prevent="onRegister">
      <div>
        <label>User</label>
        <input v-model="form.username" required />
      </div>
      <div>
        <label>Password</label>
        <input v-model="form.password" type="password" required />
      </div>
      <div>
        <label>Confirm Password</label>
        <input v-model="form.confirmPassword" type="password" required />
      </div>
      <button type="submit">สมัครสมาชิก</button>
      <div v-if="error" :class="['alert', success ? 'alert-success' : 'error']">{{ error }}</div>
    </form>
    <router-link to="/" class="back-link">ย้อนกลับ</router-link>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['close', 'registered'])
const form = ref({ username: '', password: '', confirmPassword: '' })
const error = ref('')
const success = ref(false)

async function onRegister() {
  error.value = ''
  success.value = false
  if (form.value.password !== form.value.confirmPassword) {
    error.value = 'Password ไม่ตรงกัน'
    return
  }
  try {
    // Call backend API for registration
    const res = await fetch('/api/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username: form.value.username, password: form.value.password })
    })
    if (!res.ok) throw new Error('Registration failed')
    emit('registered', form.value.username)
    // Show success message and redirect after short delay
    success.value = true
    error.value = 'สมัครสมาชิกสำเร็จ! กำลังกลับไปหน้าเข้าสู่ระบบ...'
    setTimeout(() => {
      window.location.href = '/'
    }, 1200)
  } catch (e) {
    error.value = 'สมัครสมาชิกไม่สำเร็จ'
    success.value = false
  }
}
</script>

<style scoped>
.register-container {
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
  background: #ffeaea;
  border: 1px solid #ffbdbd;
  padding: 8px;
  border-radius: 4px;
}
.alert {
  margin-top: 8px;
  padding: 8px;
  border-radius: 4px;
  font-weight: bold;
  text-align: center;
}
.alert-success {
  color: #155724;
  background: #d4edda;
  border: 1px solid #b7e3c2;
}
.back-link {
  display: block;
  margin-top: 12px;
  color: #22a322;
  cursor: pointer;
  text-align: right;
  text-decoration: underline;
}
</style>
