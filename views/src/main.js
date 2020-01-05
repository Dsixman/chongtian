import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import Axios from 'axios'
import VueAxios from 'vue-axios'
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'normalize.css/normalize.css';
Vue.use(ElementUI);
Vue.use(VueAxios,Axios)

Vue.config.productionTip = false

router.beforeEach(function (to, from, next) {
  if (to.meta.requireAuth){
    Vue.axios.get('./api/getAuth').then((data)=>{
      var data=data.data
      var loginState=data.loginState
      var userName=data.userName
      store.commit('receive_login_state',{loginState})
      store.commit('receive_user_name',{userName})
      if (store.state.userName){
          next()
        }else {    
            next({name: "login"})   
        }
    }).catch((err)=>{
      console.log(err)
    })
  }else{
     next()
  }
});

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
