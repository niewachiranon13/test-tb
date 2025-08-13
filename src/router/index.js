import { createRouter, createWebHistory } from 'vue-router'
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from '../components/RegisterForm.vue'
import WelcomeUser from '../components/WelcomeUser.vue'

const routes = [
  { path: '/', component: LoginForm },
  { path: '/register', component: RegisterForm },
  { path: '/welcome', component: WelcomeUser, props: route => ({ username: route.query.username }) }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
