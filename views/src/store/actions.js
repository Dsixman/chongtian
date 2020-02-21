import {
	RECEVIE_LOGIN_STATE,
	RECEVIE_CAPTCHA_KEY,
	RECEVIE_CAPTCHA_STR,
	RECEVIE_USER_NAME,
	RECEVIE_HERO_DATA,
	RECEVIE_All_TEAM_INFO,
	RECEVIE_TEAM_HERO_POOL,
	RECEVIE_TEN_MIN_DATA,
	RECEVIE_TEAM_DATA,
	RECEVIE_VERSION,
	RECEVIE_MATCH_DETAILS
} from './mutation-type'

import {
	reqLogin,
	reqCaptcha,
	reqPostLogin,
	reqAuth,
	reqHeroData,
	reqAllTeamInfo,
	reqTeamHeroPool,
	reqTenMinData,
	reqTeamData,
	reqMatchDetails,
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
			console.log(allTeamInfo)
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
			commit(RECEVIE_TEAM_HERO_POOL,{teamHeroPool})
		}).catch((err)=>{
			console.log(err)
		})
	},

	async getTenMinData({commit},param){
		await reqTenMinData(param).then((data)=>{
			var tenMinData=data.data
			/*tenMinData.mid_hero_data.forEach((item)=>{
				if (item!=null){
					item.pre_10min_item.forEach((item2)=>{
						item2.item_icon=require("../static/img/item/"+item2.item_icon)
					})
					item.consumable.forEach((item3)=>{
						item3.item_icon=require("../static/img/item/"+item3.item_icon)
					})
				}
			})

			tenMinData.edge_heroes_data.forEach((item)=>{
				if (item!=null){
					item.pre_10min_item.forEach((item2)=>{
						item2.item_icon=require("../static/img/item/"+item2.item_icon)
					})
					item.consumable.forEach((item3)=>{
						item3.item_icon=require("../static/img/item/"+item3.item_icon)
					})
				}
			})
			tenMinData.other_edge_heroes_data.forEach((item)=>{
				if (item!=null){
					item.pre_10min_item.forEach((item2)=>{
						item2.item_icon=require("../static/img/item/"+item2.item_icon)
					})
					item.consumable.forEach((item3)=>{
						item3.item_icon=require("../static/img/item/"+item3.item_icon)
					})
				}
			})*/
			console.log(tenMinData)
			commit(RECEVIE_TEN_MIN_DATA,{tenMinData})
		})
	},
	async getTeamData({commit},param){
		//console.log("111")
		await reqTeamData(param)
		.then((data)=>{
			var teamData=data.data
			teamData.team_base_info.team_icon=require("../static/img/team/"+teamData.team_base_info.team_icon)
			teamData.team_bp.forEach((item)=>{
				item.picks_bans.forEach((item2)=>{
					item2.hero_icon=require("../static/img/hero/"+item2.hero_icon)
				})
			})
			console.log(teamData)
			commit(RECEVIE_TEAM_DATA,{teamData})		
		})
		.catch((err)=>{
			console.log(err)
		})
	},
	getVersion({commit},param){
		var version=param

		commit(RECEVIE_VERSION,{version})
	},
	async getMatchDetails({commit},param){
		await reqMatchDetails(param).then((data)=>{
			var matchDetails=data.data
			matchDetails.game_info.picks_bans.forEach((item)=>{
				item.hero_icon=require("../static/img/hero/"+item.hero_icon)	
			})
			matchDetails.details.players_heroes_dets.forEach((item)=>{
				//console.log(item.hero_name)
				item.hero_icon=require("../static/img/hero/"+item.hero_icon)
				item.ability_data.forEach((item2)=>{
					if (item2.Icon!=null){
						if (item2.Icon!=""){
							item2.Icon=require("../static/img/abilities/"+item2.Icon)
						}else{
							item2.Icon=require("../static/img/abilities/talent.jpg")
						}
						
					}
				})

				item.items.forEach((item2,index)=>{
					if (item2.item_icon!=""){
						//console.log(index,item2.item_id)
						item2.item_icon=require("../static/img/item/"+item2.item_icon)
					}
					/*if (item2.item_icon==""){
					
						console.log(index,item2.item_id)
					}*/
				})	
			})
			matchDetails.result_data.players.forEach((item)=>{
				item.item.forEach((item2,index)=>{
					if (item2!=""){
						item.item[index]=require("../static/img/item/"+item2)		
						/*item2=require("../static/img/item/"+item2)//这样写是无效的
						console.log(item2)*/	
					}
				})
				//console.log(item.hero_icon)
				item.hero_icon=require("../static/img/hero/"+item.hero_icon)
				if(item.buffs.length>0){
					item.buffs.forEach((buffs,index)=>{
						if (buffs.buff_id==6){
							item.buffs[index].buffs=require("../static/img/item/tome-of-knowledge.jpg")
						}
					})
				}
			})
			console.log(matchDetails)
			commit(RECEVIE_MATCH_DETAILS,{matchDetails})
		})
	}
}