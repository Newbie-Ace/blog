import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    { path: '/', redirect: '/login' },
    { path: '/login', component: () => import('@/views/login/LoginView.vue') },
    {
      path: '/home',
      component: () => import('@/views/home/HomeView.vue'),
      redirect: '/blog',
      children: [
        { path: '/blog', component: () => import('@/components/blog/BlogComponent.vue') },
        { path: '/search', component: () => import('@/components/search/SearchComponent.vue') },
        { path: '/about', component: () => import('@/components/about/AboutComponent.vue') }
      ]
    }
  ]
})

export default router
