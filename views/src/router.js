import Vue from 'vue'
import Router from 'vue-router'

const TeamBP=()=> import('./views/TeamMatchData.vue')
const Hero=()=> import('./views/Hero.vue')
const HeroData=()=> import('./views/Herodata.vue')
const Login=()=>import('./views/Login.vue')
const BaseEchart=()=>import('./views/BaseEchart.vue')
const LineUp=()=>import('./views/LineUp.vue')
const Team=()=>import('./views/Team.vue')
const MatchDetails=()=>import('./views/MatchDetails.vue')
const Mid=()=>import('@/components/Mid.vue')
const LineUpDets=()=>import('./views/LineUpDets.vue')
const SideLineUpDets=()=>import('./views/SideLineUpDets.vue')
const PlayerSameHero=()=>import('./views/PlayerSameHero.vue')
const BPiter=()=>import('./views/BPiter.vue')

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Hero,
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
      },
      children: [/*{
        path: 'teambp/:team/:version',
        name:'teambp',
        component: TeamBP,
      },*/
      {
       path: 'midlineup/:playerid/:version',
        name:'midlineup',
        component:Mid,
      },
      {
        path:'bpiter/:team/:version',
        name:'bpiter',
        component:BPiter
      }
      ]
    },
    {
        path: '/teambp/:team/:version',
        name:'teambp',
        component: TeamBP,
      },
    
    {
      path:'/playersamehero/:playerid/:hero/:version',
      name:'playersamehero',
      component:PlayerSameHero
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
    },
    {
      path: '/login',
      name: 'login',
      component:Login
    },
    {
      path: '/lineup',
      name: 'lineup',
      component: LineUp,
       meta:{
        "requireAuth": true
      }
    },
    {
      path: '/lineupdets',
      name: 'lineupdets',
      component: LineUpDets,
       meta:{
        "requireAuth": true
      }
    },
    {
      path: '/sidelineupdets',
      name: 'sidelineupdets',
      component: SideLineUpDets,
       meta:{
        "requireAuth": true
      }
    }
  ]
})
