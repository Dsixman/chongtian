<template>
    <div>
    	<form action="./api/login">
    		<input type="hidden" name="captchakey" v-bind:value="loginData">
			<input type="text" name="username" value="">
			<input type="text" name="password" value="">
			<img src="">
    	</form>
    </div>
</template>

<script type="ecmascript-6">
import {mapState} from 'vuex'
    export default{
    	computed:{
    		...mapState(['loginData']),
    	}
    
    }
</script>

<style rel="stylesheet">
    
</style>