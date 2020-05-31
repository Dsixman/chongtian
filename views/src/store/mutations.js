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
	RECEVIE_MATCH_DETAILS,
	RECEVIE_LINE_UP_REQUEST,
	RECEVIE_MID_USE_HEROES,
	RECEVIE_PLAYER_SAME_HERO,
	RECEVIE_SIDE_LINE_UP_DATA
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
	},
	[RECEVIE_TEN_MIN_DATA] (state,{tenMinData}){
		state.tenMinData=tenMinData;
	},
	[RECEVIE_TEAM_DATA] (state,{teamData}){
		state.teamData=teamData;
	},
	[RECEVIE_VERSION] (state,{version}){
		state.version=version;
	},
	[RECEVIE_MATCH_DETAILS] (state,{matchDetails}){
		state.matchDetails=matchDetails;
	},
	[RECEVIE_MID_USE_HEROES] (state,{midUseHeroes}){
		state.midUseHeroes=midUseHeroes;
	},
	[RECEVIE_LINE_UP_REQUEST] (state,{lineUpRequest}){
		state.lineUpRequest=lineUpRequest;
	},
	[RECEVIE_PLAYER_SAME_HERO] (state,{playerSameHero}){
		state.playerSameHero=playerSameHero;
	},
	[RECEVIE_SIDE_LINE_UP_DATA] (state,{sideLineUpData}){
		state.sideLineUpData=sideLineUpData;
	},
}