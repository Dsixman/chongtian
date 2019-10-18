import {
	RECEVIE_LOGIN_STATE,
	RECEVIE_CAPTCHA_KEY,
	RECEVIE_CAPTCHA_STR,
	RECEVIE_USER_NAME

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
	}
}