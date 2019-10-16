import {
	RECEVIE_LOGIN_DATA
} from './mutation-type'

export default {
	[RECEVIE_LOGIN_DATA] (state,{loginData}){
		state.loginData=loginData;
	}
}