<template>
	<div v-if="allTeamInfo!=null" id="teamwrapper">

		<!-- <input v-model="getVersion" placeholder="输入要搜索数据的版本" id="version-input"></input> -->
		 <el-tabs v-model="activeName" @tab-click="handleClick" id="select">
		    <el-tab-pane label="战队面板" name="first" id="select_team">
				<ul id="team_icon_list">
					<li v-for="item in allTeamInfo" :key="item.team_id">
						<!-- <img :src="item.team_icon" width="80" 
						@click="getTeamInfo(item.team_tag_name)"> -->
						<router-link :to="{name:'teambp',params:{team:item.team_tag_name,version:$store.state.version}}">
							<img :src="item.team_icon" width="80" />
						</router-link>
					</li>
					
				</ul>
		    </el-tab-pane>
		  </el-tabs>
		<!-- <div><router-view></router-view></div> -->
	</div>
</template>

<script>
import {mapState} from 'vuex'
export default{
	data(){
		return{
			version:"",
			/*activeNames: ['1'],*/
			activeName: 'first'
		}
	},
	computed:{
		...mapState(['allTeamInfo']),
		/*activeName:function(){
			var activeName1 
			if (this.$route.path!="/team"){
				activeName1='first'
			}
			return activeName1
		},*/
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
		 handleChange(val) {
        console.log(val);

      },
      handleClick(tab, event) {
        //console.log(tab, event);
  		var content=document.getElementById("select_team")
  		if (content.style.display==""){
  			content.style.display="none"
  		}else{
  			content.style.display=""
  		}
      },
		getTeamInfo(teamname){
			//this.$router.push('/teammatchdata?team='+teamname/*{name:'teammatchdata',params:{team:teamname}}*/)
			var content=document.getElementById("select_team")
			//content.style.display="none"
			//this.activeName=''
			//console.log()
			this.$router.push({name:'teambp',params:{team:teamname,version:this.$store.state.version}})
		}
	}
}	
</script>
<style>
li{
	list-style:none;
}
#teamwrapper{
	width:1400px;
	margin:0 auto;
}
#team_icon_list li{
	float:left;
	width:90px;
	height:150px;
}

#team_icon_list li img:hover{
	box-shadow:#cccccc 0px 0px 5px ;
	cursor:pointer
}
#select{
	margin-top:15px
}
#select .el-tabs__nav{
	background:#091723;
	width:100%;
	opacity:0.5
}
#select .el-tabs__active-bar{
	width:0px;height:0px;
}
</style>