import { createRouter, createWebHistory } from 'vue-router'
import Blogs from '../views/Blogs.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Blogs,
    },
    {
      path: '/blogs/:id',
      name: 'blog',
      component: () => import('../views/BlogView.vue')
    },
    {
      path: '/blogs/add',
      name: 'add',
      component: () => import('../views/AddBlog.vue')
    }
  ]
})

export default router
