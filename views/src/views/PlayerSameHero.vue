<template>
	<div>
		<table>
			<tr>
				<td>比赛ID</td>
				<td>输赢</td>
				<td>5分钟总经济</td>
				<td>10分钟总经济</td>
				<td>15分钟总经济</td>
				<td>20分钟总经济</td>
				<td>25分钟总经济</td>
				<td>前20分钟KDA情况</td>
				<td>加点</td>
			</tr>
			<tr v-for="item in alldata">
				<td>{{item.match_id}}</td>
				<td >
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero">
						{{item2.game_result}}
						</span>
					</div>
				</td>
				<td >
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero">
						{{item2.gold.pre_5min_gold.all_gold}}
						</span>
					</div>	
				</td>
				<td >
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero">
						{{item2.gold.pre_10min_gold.all_gold}}
						</span>
					</div>		
				</td>
				<td >
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero">
						{{item2.gold.pre_15min_gold.all_gold}}
						</span>
					</div>		
				</td>
				<td >
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero&&item2.gold.pre_20min_gold!=null">
						{{item2.gold.pre_20min_gold.all_gold}}
						</span>
					</div>		
				</td>
				<td>
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero&&item2.gold.pre_25min_gold!=null">
						{{item2.gold.pre_20min_gold.all_gold}}
						</span>
					</div>
				</td>
				<td>
					{{kdadata}}
				</td>
				<td>
					<div v-for="item2 in item.players_heroes_dets">
						<span v-if="item2.hero_name==$route.params.hero" >
							<img v-for="item3 in item2.ability_data" :src="item3.Icon" alt="" width="30" style="padding-left:8px;">
						</span>
					</div>
				</td>
			</tr>
		</table>
	</div>
</template>
<script>
import {mapState} from 'vuex'
	export default{
		data(){
			return {}
		},
		computed:{
			...mapState['playerSameHero'],
			alldata(){
				var alldata	=[]
				if (this.$store.state.playerSameHero){
					if (this.$store.state.playerSameHero.win_data!=null){
						alldata=this.$store.state.playerSameHero.win_data
					}
				}
				if (this.$store.state.playerSameHero){
					if (this.$store.state.playerSameHero.lose_data!=null){
						this.$store.state.playerSameHero.lose_data.forEach(item=>{
							alldata.push(item)
						})
					}
				}
				console.log(alldata)
				return alldata				
			},
			kdadata(){
				var data=this.alldata
				var kills
				var deaths
				var assists

				data.forEach(item=>{
					item.players_heroes_dets.forEach(item2=>{
						if (item2.hero_name==this.$route.params.hero){
							item2.inventory_snap_shot.forEach(item3=>{
								if (item3.game_time>900){
									kills=item3.kills
									deaths=item3.deaths
									assists=item3.assists
								}
							})
						}
					})
				})
				var res=kills+"/"+deaths+"/"+assists
				return res
			}
		},
		created(){
			var params={playerid:this.$route.params.playerid,hero:this.$route.params.hero,version:this.$route.params.version}
			this.$store.dispatch('getPlayerSameHero',params)
		}
	}
</script>
<style>
</style>