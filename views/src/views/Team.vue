<template>
	<div v-if="allTeamInfo!=null" id="teamwrapper">
		<input v-model="getVersion" placeholder="输入要搜索数据的版本" id="version-input"></input> 
		<!-- <button @click="test()">提示版本</button> -->
		<ul id="team_icon_list" >
			<li v-for="item in allTeamInfo" :key="item.team_id"><img :src="item.team_icon" width="80" 
				@click="getTeamInfo(item.team_tag_name)"></li>
		</ul>
	</div>
</template>

<script>
import {mapState} from 'vuex'
export default{
	data(){
		return{
			version:"",
		}
	},
	computed:{
		...mapState(['allTeamInfo']),
		getVersion:{
			get: function() {
				if (this.$store.state.version!=null){
					this.version=this.$store.state.version
				}
				return this.$store.state.version
			},
			set:function(value){
				this.version = value
				this.$store.dispatch('getVersion',this.version)

			}
		}
	},
	created(){
		this.$store.dispatch('getAllTeamInfo')
	},
	methods:{
		/*test(){
			console.log(this.version)
		}*/
		getTeamInfo(teamname){
			this.$router.push('/teammatchdata?team='+teamname/*{name:'teammatchdata',params:{team:teamname}}*/)
		}
	}
}	
</script>
<style scoped>
li{
	list-style:none;
}
#teamwrapper{
	width:1000px;
	margin:0 auto;
}
#team_icon_list li{
	float:left;
	width:90px;
	height:110px;
}
#version-input{
	margin-top:15px;
background-color: #FFF;
    background-image: none;
    border-radius: 4px;
    border: 1px solid #DCDFE6;
    box-sizing: border-box;
    color: #606266;
    display: inline-block;
    font-size: inherit;
    height: 40px;
    line-height: 40px;
    outline: 0;
    padding: 0 15px;
    font-size:14px
}
#team_icon_list li img:hover{
	box-shadow:#cccccc 0px 0px 5px ;
	cursor:pointer
}
</style>