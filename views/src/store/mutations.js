import {
	RECEVIE_LOGIN_STATE,
	RECEVIE_CAPTCHA_KEY,
	RECEVIE_CAPTCHA_STR,
	RECEVIE_USER_NAME,
	RECEVIE_HERO_DATA,
	RECEVIE_All_TEAM_INFO,
	RECEVIE_TEAM_HERO_POOL
} from './mutation-type'

export default {
	[RECEVIE_LOGIN_STATE] (state,{loginState}){
		state.loginState=loginState;
	},
	[RECEVIE_CAPTCHA_STR] (state,{captchaStr}){
		state.captchaStr=captchaStr;
	},
	[RECEVIE_CAPTCHA_KEY] (state,{captchaKey}){
		state.captchaKey=captchaKey;
	},
	[RECEVIE_USER_NAME] (state,{userName}){
		state.userName=userName;
	},
	[RECEVIE_HERO_DATA] (state,{heroData}){
		state.heroData=heroData;
	},
	[RECEVIE_All_TEAM_INFO] (state,{allTeamInfo}){
		state.allTeamInfo=allTeamInfo;
	},
	[RECEVIE_TEAM_HERO_POOL] (state,{teamHeroPool}){
		state.teamHeroPool=teamHeroPool;
	}
}