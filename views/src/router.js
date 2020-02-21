import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
const TeamMatchData=()=> import('./views/TeamMatchData.vue')
const Hero=()=> import('./views/Hero.vue')
const HeroData=()=> import('./views/Herodata.vue')
const Login=()=>import('./views/Login.vue')
const BaseEchart=()=>import('./views/BaseEchart.vue')
const Lineup=()=>import('./views/Lineup.vue')
const Team=()=>import('./views/Team.vue')
const MatchDetails=()=>import('./views/MatchDetails.vue')
/*
import Hero from './views/Hero.vue'
import HeroData from './views/Herodata.vue'
import Login from './views/Login.vue'*/


Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home,
      meta:{
        "requireAuth": true
      }
    },
    {
      path: '/teammatchdata',
      name: 'teammatchdata',
      component: TeamMatchData,
       meta:{
        "requireAuth": true
      }
    },
    {
      path: '/team',
      name: 'team',
      component: Team,
       meta:{
        "requireAuth": true
      }
    },
    {
      path: '/hero',
      name: 'hero',
      component: Hero,
      meta:{
        "requireAuth": true
      }
    },
    {
      path:'/matchdetails',
      name:'matchdetails',
      component:MatchDetails,
    },
    {
      path: '/herodata',
      name: 'herodata',
      component: HeroData,
       meta:{
        "requireAuth": true
      }
    },
    {
      path: '/BaseEchart',
      name: 'baseechart',
      component: BaseEchart,
      /*meta:{
        "requireAuth": true
      }*/
    },
    {
      path: '/login',
      name: 'login',
      component:Login
    },
    {
      path: '/lineup',
      name: 'lineup',
      component: Lineup,
       meta:{
        "requireAuth": true
      }
    }
  ]
})
