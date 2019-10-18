import Vue from 'vue'
import Axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios,Axios)

export const reqLogin=()=>{
	return Vue.axios.get('./api/login')
};
export const reqCaptcha=()=>{
	return Vue.axios.get('./api/getCaptcha')
}

export const reqPostLogin=(param)=>{
	return Vue.axios.post('./api/login',param)
};
export const reqAuth=()=>{
	return Vue.axios.get('./api/getAuth')
};