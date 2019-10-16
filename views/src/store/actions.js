import {
	RECEVIE_LOGIN_DATA
} from './mutation-type'

import {
	reqLogin
} from '../api'

export default{
	async getloginData ({commit}) {
		var loginData
		await reqLogin()
		.then((data)=>{
			loginData=data.data
			commit(RECEVIE_LOGIN_DATA,{loginData})
		})
		.catch((err)=>{
			console.log(err)
		});
	}
/*	
	getUserInfo({commit},userInfo){
		commit(RECEVIE_USERINFO,{userInfo})
	}*/
}