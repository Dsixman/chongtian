import {
	RECEVIE_LOGIN_STATE,
	RECEVIE_CAPTCHA_KEY,
	RECEVIE_CAPTCHA_STR,
	RECEVIE_USER_NAME
} from './mutation-type'

import {
	reqLogin,
	reqCaptcha,
	reqPostLogin,
	reqAuth
} from '../api'

export default{
	
	async getLoginData ({commit}) {
		var loginData
		await reqLogin()
		.then((data)=>{
			loginData=data.data
			var loginState=loginData.loginState
			var captchaKey=loginData.capKey
			var captchaStr=loginData.capStr
			commit(RECEVIE_LOGIN_STATE,{loginState})
			commit(RECEVIE_CAPTCHA_STR,{captchaStr})
			commit(RECEVIE_CAPTCHA_KEY,{captchaKey})
		})
		.catch((err)=>{
			console.log(err)
		});
	},
	async getCaptcha ({commit}) {
		var captcha
		await reqCaptcha()
		.then((data)=>{
			captcha=data.data
			var captchaKey=captcha.capKey
			var captchaStr=captcha.captchaStr
			commit(RECEVIE_CAPTCHA_STR,{captchaStr})
			commit(RECEVIE_CAPTCHA_KEY,{captchaKey})
		})
		.catch((err)=>{
			console.log(err)
		});
	},
	getLoginAuth({commit},data){
		console.log(data)
		var userName=data.userName
		var loginState=data.loginState
		console.log(userName)
		console.log(loginState)
		commit(RECEVIE_LOGIN_STATE,{loginState})
        commit(RECEVIE_USER_NAME,{userName}) 
		/*await reqPostLogin(param)
                .then((data)=>{
                    var data=data.data
                    if (data.loginState=="true"){
                    	var loginState=data.loginState
                    	var userName=data.userName
                    	commit(RECEVIE_LOGIN_STATE,{loginState})
                    	commit(RECEVIE_USER_NAME,{userName})                        
                    }
                }).catch((err)=>{
                    alert("用户名或密码出错")
                })*/
	},
	async getAuth({commit}){
		
	await reqAuth().then((data)=>{
		var data=data.data
		var loginState=data.loginState
	commit(RECEVIE_LOGIN_STATE,{loginState})
	}).catch((err)=>{
		console.log(err)
	})
	}
/*	
	getUserInfo({commit},userInfo){
		commit(RECEVIE_USERINFO,{userInfo})
	}*/
}