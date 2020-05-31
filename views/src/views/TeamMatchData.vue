<template>
	<el-main> 
		<div>
			<el-tabs v-model="activeName" >
			   <el-tab-pane label="选手使用英雄" name="first">
				<div id="team_info" v-if="teamData!=null" class="wrapper">
					<div id="base_info">
						<div class="logo"><img :src="teamData.team_base_info.team_icon" width="120"></div>
						<div class="player_info">
							<div>
								<div class="hero_pool player_data" v-for="item in teamData.team_hero_pool" :key="item.player_register_string_id">
									<table>
										<tr>
											<td>选手</td><td>位置</td><td>英雄池</td>
										</tr>
										<tr>
											<td>{{item.player_register_string_id}}</td><td>{{item.position}}</td>
											<td>
												<table >
													<tr>
														<td>英雄</td><td>比赛使用总次数</td><td>比赛胜率</td><td>天梯使用次数</td>
													</tr>
													<tr v-for="item2 in item.match_player_heroes.hero_play_count">
														<td><router-link :to="{name:'playersamehero',params:{playerid:item.player_dota2_register_num_id,hero:item2.hero,version:version}}" :data-player="item.player_register_string_id" :data-hero="item2.hero">{{item2.hero}}</router-link></td>
														<td>{{item2.league.all_count}}</td>
														<td>{{(item2.league.win_count/item2.league.all_count).toFixed(2)*100}}%</td>
														<td>{{item2.rank.all_count}}</td>
													</tr>
												</table>
											</td>
										</tr>
									</table>
								</div>
							</div>
						</div>
						<div class="clear"></div>
					</div>				
				</div>
			   </el-tab-pane>
			   <el-tab-pane label="战队BP数据" name="second">
			   	<Page :pageData="teamData.team_bp"></Page>
			   </el-tab-pane>
			    <el-tab-pane label="选手对线" name="third">
					<div class="mid_against ">
						<p class="title-bar">中单选手英雄对线数据（左边为本战队中单使用的英雄）</p>
						<ul>
							<li v-for="item in teamData.mid_against_data" @click="getMidAgainstData($event)">
								{{item}}
							</li>
							<div class="clear"></div>
						</ul>
					</div>
					<div class="side_against">
						<p class="title-bar">三号位边路对线</p>
						<ul>
							<li v-for="item in teamData.hardlineup" @click="getSideAgainstData($event)">
								{{item}}
							</li>
							<div class="clear"></div>
						</ul>
					</div>	
					<div class="side_against">
						<p class="title-bar">一号位边路对线</p>
						<ul>
							<li v-for="item in teamData.softlineup" @click="getSideAgainstData($event)">
								{{item}}
							</li>
							<div class="clear"></div>
						</ul>
						
					</div>
			    </el-tab-pane>
			</el-tabs>
		</div>

	</el-main>
</template>
<script>
import {mapState} from 'vuex'
const Page=()=>import('@/components/Page.vue')
const MidUseHeroes=()=>import('@/components/PlayerUseHeroes.vue')
	export default{
		data(){
			return{
				version:this.$route.params.version,
				activeName: 'second'
			}
		},

		computed:{
			...mapState(['teamData']),
			currentUrl(){
				var urlstr=this.$route.path
				return urlstr
			},

			midUrl(){
				var playerid
				if  (this.$store.state.teamData!=null){
					this.$store.state.teamData.team_hero_pool.forEach((item)=>{
						if (item.position=="二号位"){
							playerid=item.player_dota2_register_num_id
						}
					})
				}
				//var urlstr ="/team/midlineup/"+playerid+"/"+this.$route.params.version
				var urlstr="/team/midlineup/"+playerid+"/"+this.$route.params.version
				return urlstr
			},
			
			//前三ban最多次的英雄
			prethirdban(){
				var teambp=this.$store.state.teamData.team_bp
				var teambpnamearr=[]
				var teamname=this.$route.params.team
				teambp.forEach((item)=>{
					var gameteam
					if  (item.radiant_team_tag==teamname){
						gameteam=2
					}else{
						gameteam=3
					}
					item.picks_bans.forEach((item2,index2)=>{
						if (item2.is_pick==false&&gameteam==gameteam&&index2<=5){

							var teambpnamearrlen=teambpnamearr.length		
							if (teambpnamearrlen==0){
								var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,ban_rate:null,ban_rate_str:null}

								teambpnamearr.push(obj)
							}else{

								teambpnamearrlen=teambpnamearrlen-1
								for (var i=0;i<=teambpnamearrlen;i++){
									if (teambpnamearr[i].id==item2.hero_id){

										teambpnamearr[i].count=teambpnamearr[i].count+1
										break;
									}
									if (teambpnamearr[i].id!=item2.hero_id&&i==teambpnamearrlen){
										var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,ban_rate:null,ban_rate_str:null}
										teambpnamearr.push(obj)
										break;
									}
								}
							}
							
						}
					})
				})
				teambpnamearr.forEach((item)=>{
					var length=teambpnamearr.length
					item.ban_rate=(item.count/length).toFixed(2)
					item.ban_rate_str=(item.ban_rate*100)+"%"
					if (item.ban_rate_str.indexOf(".")>-1){
						item.ban_rate_str=item.ban_rate_str.substring(0, 2)+"%"
					}
				})
				return teambpnamearr
			},
			prethirdpick(){
				var teambp=this.$store.state.teamData.team_bp
				var teambpnamearr=[]
				var teamname=this.$route.params.team
				teambp.forEach((item)=>{
					var gameteam
					if  (item.radiant_team_tag==teamname){
						gameteam=2
					}else{
						gameteam=3
					}
					item.picks_bans.forEach((item2,index2)=>{
						if (item2.is_pick==true&&gameteam==gameteam&&index2<=9){
							var teambpnamearrlen=teambpnamearr.length		
							if (teambpnamearrlen==0){
								var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,pick_rate:null,pick_rate_str:null}
								teambpnamearr.push(obj)
							}else{

								teambpnamearrlen=teambpnamearrlen-1
								for (var i=0;i<=teambpnamearrlen;i++){
									if (teambpnamearr[i].id==item2.hero_id){
										teambpnamearr[i].count=teambpnamearr[i].count+1
										break;
									}
									if (teambpnamearr[i].id!=item2.hero_id&&i==teambpnamearrlen){
										var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,pick_rate:null,pick_rate_str:null}
										teambpnamearr.push(obj)
										break;
									}
								}
							}
							
						}
					})
				})
				teambpnamearr.forEach((item)=>{
					var length=teambpnamearr.length
					item.pick_rate=(item.count/length).toFixed(2)
					item.pick_rate_str=(item.pick_rate*100)+"%"
					if (item.pick_rate_str.indexOf(".")>-1){
						item.pick_rate_str=item.pick_rate_str.substring(0, 2)+"%"
					}
				})
				return teambpnamearr
			}
		},
		components:{
			Page,
			MidUseHeroes
		},
		created(){
			window.console.log("aaa")
			var param={team_name:this.$route.params.team,version:this.$route.params.version}
			
			this.$store.dispatch('getTeamData',param)
			/*if(!this.$route.query.playerid){
				this.$store.dispatch('getTeamData',param)	
			}*/
			
		},
		methods:{
			selectPosition(value){
				var param={position:value,team_name:this.$route.query.team,version:this.$route.query.version}
				//this.$store.dispatch('get')
			},
			getMidAgainstData(e){
				//console.log(e.target.innerText)
				var str=e.target.innerText
				var midlaneselhero=str.replace(/vs/," ")
				var strarr=midlaneselhero.split(" ")
				console.log(strarr)
				var midlane2=strarr[1]+"vs"+strarr[0]
				var midlane=str
				var request={
		            toplane:"",
		            midlane:midlane,
		            midlane2:midlane2,
		            bottomlane:"",

		            toplaneselhero:"",
		            midlaneselhero:midlaneselhero,
		            bottomlaneselhero:"",

		            sidelane:"",
		            softlane:"",
		            hardlane:"",
		            version:this.$route.params.version,
		            team_name:this.$route.params.team
		         }
		         console.log(request)
		         this.$store.dispatch('getLineUpRequest',request)
		         this.$router.push('/lineupdets')
			},
			getSideAgainstData(e){
				//console.log(e.target.innerText)
				var line_up_heroes=e.target.innerText
				/*var sidelaneselhero=str.replace(/vs/," ")
				var strarr=sidelaneselhero.split(" ")
				//console.log(strarr)
				var sidelane2=strarr[1]+"vs"+strarr[0]
				var sidelane=str*/
				var request={
		            line_up_heroes:line_up_heroes,
		            version:this.$route.params.version,
		            team_name_tag:this.$route.params.team
		         }
		         //console.log(request)
		         this.$store.dispatch('getLineUpRequest',request)
		         this.$router.push('/sidelineupdets')
			},
			/*handleClick(tab, event) {
        		console.log(tab, event);
      		}*/
		}
	}
</script>
<style scoped>
.el-tabs__item{
	color:#989898;
}
.logo{
	width:200px;
	float:left
}
.player_title{
	float:left
}
.player_info{
	float:left;
	width:1200px
}
.hero_pool{
width:90%;
margin-bottom:10px;
}
#position{
	float:left
}
#base_info{
	width:100%;
	height:auto;
}
.hero_pool_dets{
	float:left;
	margin-left:10px;
	margin-bottom:10px;
}
.heor_play{
	margin-top:10px;
	margin-bottom:10px
}
.heor_play_title{
	margin-bottom:5px;
}
.top-nav{
	height:50px;
}
.top-nav-bp{
	padding-right:10px;
}
.mid_against ul li{
	float:left;height:25px;width:180px;text-align:center

}
.side_against ul li{
	float:left;height:25px;width:260px;text-align:center

}
.mid_against ul li:hover,.side_against ul li:hover{
	cursor:pointer
}
.top-nav div{
	float:left
}
.title-bar{
	background-color: rgba(0, 0, 0, 0.37);
	opacity: 0.5;
	padding:5px;
}
.sencond-title{
	font-size:14px;
	text-align:left
}
</style>