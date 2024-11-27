import { createRouter, createWebHistory } from 'vue-router'
import TwoPlayerModeView from '@/views/TwoPlayerModeView.vue'
import MainMenuView from '@/views/MainMenuView.vue'
import { ref } from 'vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'MainMenu',
      component: MainMenuView,
      meta: {}
    },    
    {
      path: '/2player',
      name: 'TwoPlayerMode',
      component: TwoPlayerModeView,
      meta: {}
    },

  ]
  
})


export default router
