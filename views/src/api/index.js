import Vue from 'vue'
import Axios from 'axios'
import VueAxios from 'vue-axios'
Vue.use(VueAxios,Axios)

export const reqLogin=()=>{
	return Vue.axios.get('./api/login')
};
export const reqCaptcha=()=>{
	return Vue.axios.get('./api/getcaptcha')
};
export const reqPostLogin=(param)=>{
	return Vue.axios.post('./api/login',param)
};
export const reqAuth=()=>{
	return Vue.axios.get('./api/getAuth')
};
export const reqHeroData=(param)=>{
	return Vue.axios.post('./api/getherodata',param)
};
export const reqHeroIcon=()=>{
	return Vue.axios.get('./api/getheroesicon')
};
export const reqAllTeamInfo=()=>{
	return Vue.axios.get('./api/getallteaminfo/')
}
export const reqTeamHeroPool=(param)=>{
	return Vue.axios.get('./api/getteamheropool',{params:param})
}
export const reqTenMinData=(param)=>{
	return Vue.axios.post('./api/gettenmindata/',param)
}
export const reqTeamData=(param)=>{
	return Vue.axios.post('./api/getteamalldata/',param)
}
export const reqMatchDetails=(param)=>{
	return Vue.axios.get('./api/getmatchdetails',{params:param})
}

  
