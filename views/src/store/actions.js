import {
	RECEVIE_LOGIN_STATE,
	RECEVIE_CAPTCHA_KEY,
	RECEVIE_CAPTCHA_STR,
	RECEVIE_USER_NAME,
	RECEVIE_HERO_DATA,
	RECEVIE_All_TEAM_INFO,
	RECEVIE_TEAM_HERO_POOL
} from './mutation-type'

import {
	reqLogin,
	reqCaptcha,
	reqPostLogin,
	reqAuth,
	reqHeroData,
	reqAllTeamInfo,
	reqTeamHeroPool
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
		//console.log(data)
		var userName=data.userName
		var loginState=data.loginState
		commit(RECEVIE_LOGIN_STATE,{loginState})
        commit(RECEVIE_USER_NAME,{userName}) 
	},
	async getAuth({commit},router){	
	await reqAuth().then((data)=>{
		var data=data.data
		var loginState=data.loginState
		var userName=data.userName
		commit(RECEVIE_LOGIN_STATE,{loginState})
		commit(RECEVIE_USER_NAME,{userName}) 
		/*if (loginState!=false) {
			router.push('/hero')
		}	*/
	}).catch((err)=>{
		console.log(err)
	})
	},

	async getHeroData({commit},param){
		await reqHeroData(param).then((data)=>{
		var heroData=data.data
		for (var key in heroData){
			var str=heroData[key]
			if (typeof(str)=="string"){
				var result=str.indexOf(".jpg")
				if (result!=-1){
					var heroIcon=require("../static/img/hero/"+str)
					heroData[key]=heroIcon	
				}
			}
		}
		for (var key in heroData.Ability){
			var str=heroData.Ability[key]
			if (typeof(str)=="string"){
				var result=str.indexOf(".jpg")
				if (result!=-1){
					var abilityIcon=require("../static/img/abilities/"+heroData.Ability[key])
					heroData.Ability[key]=abilityIcon
				}
			}
		}
		commit(RECEVIE_HERO_DATA,{heroData})
		//console.log(heroData)
		}).catch((err)=>{
			console.log(err)
		})
	},
	async getAllTeamInfo({commit}){
		await reqAllTeamInfo().then((data)=>{
			var allTeamInfo=data.data
			for (var key in allTeamInfo){
				allTeamInfo[key].team_icon=require("../static/img/team/"+allTeamInfo[key].team_icon)
			}
			commit(RECEVIE_All_TEAM_INFO,{allTeamInfo})
			//console.log(allTeamInfo)
		}).catch((err)=>{
			console.log(err)
		})
	},
	async getTeamHeroPool({commit},param){
		await reqTeamHeroPool(param).then((data)=>{
			var teamHeroPool=data.data
			/*hero_play_count
			match_player_heroes*/
			teamHeroPool.forEach((item)=>{
				item.match_player_heroes.hero_play_count.forEach((item2)=>{
					item2.hero_icon=require("../static/img/hero/"+item2.hero_icon)
				})
				if (item.alternate_player_heroes!=null){
					item.alternate_player_heroes.hero_play_count.forEach((item2)=>{
						item2.hero_icon=require("../static/img/hero/"+item2.hero_icon)
					})	
				}
					
				if (item.old_club_player_heroes!=null){
					item.old_club_player_heroes.hero_play_count.forEach((item2)=>{
						item2.hero_icon=require("../static/img/hero/"+item2.hero_icon)
					})
				}
				if (item.RankPlayerHeroes!=null){
					item.RankPlayerHeroes.forEach((item2)=>{
						item2.hero_icon=require("../static/img/hero/"+item2.hero_icon)
					})
				}
			})

			/*teamHeroPool.match_player_heroes.hero_play_count.forEach((item)=>{
				item.hero_icon=require("../static/img/hero/"+item.hero_icon)
			})*/
			/*if teamHeroPool.alternate_player_heroes.HeroPlayCount.forEach((item)=>{
				item.hero_icon=require("../static/img/hero/"+item.hero_icon)
			})*/
			commit(RECEVIE_TEAM_HERO_POOL,{teamHeroPool})
			//console.log(teamHeroPool)
		}).catch((err)=>{
			console.log(err)
		})
	}
/*	
	getUserInfo({commit},userInfo){
		commit(RECEVIE_USERINFO,{userInfo})
	}*/
}