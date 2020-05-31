<template>
	<div style="width:100%">
		<div id="line_up_over_view">
				<table class="details_table" width="700" style="margin:auto">
					<caption style="width:700px;">{{sideLineUpData.over_view.line_up_heroes}} 对线概览 (数据量{{sideLineUpData.over_view.counts}})</caption>
					<tr><td>英雄</td><td>占比最多的经济区间</td><td>占比第二的经济区间</td><td>死亡区间</td></tr>
					<tr>
						<td>{{sideLineUpData.over_view.hero1_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero1_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero1_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero1_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero1_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero1_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero1_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
					<tr v-if="sideLineUpData.over_view.hero2_gold_dets.hero_name!=''">
						<td>{{sideLineUpData.over_view.hero2_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero2_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero2_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero2_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero2_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero2_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero2_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
					<tr v-if="sideLineUpData.over_view.hero3_gold_dets.hero_name!=''">
						<td>{{sideLineUpData.over_view.hero3_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero3_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero3_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero3_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero3_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero3_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero3_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
					<tr>
						<td>{{sideLineUpData.over_view.hero4_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero4_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero4_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero4_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero4_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero4_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero4_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
					<tr v-if="sideLineUpData.over_view.hero5_gold_dets.hero_name!=''">
						<td>{{sideLineUpData.over_view.hero5_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero5_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero5_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero5_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero5_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero5_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero5_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
					<tr v-if="sideLineUpData.over_view.hero6_gold_dets.hero_name!=''">
						<td>{{sideLineUpData.over_view.hero6_gold_dets.hero_name}}</td>
						<td>{{sideLineUpData.over_view.hero6_gold_dets.gold_maximum_proportion}}</td>
						<td>{{sideLineUpData.over_view.hero6_gold_dets.gold_second_max_proportion}}</td>
						<td>	
							<ul>
								<li>死亡一次的次数：{{sideLineUpData.over_view.hero6_gold_dets.death_1counts}}</li>
								<li>死亡两次的次数：{{sideLineUpData.over_view.hero6_gold_dets.death_2counts}}</li>
								<li>死亡三次的次数：{{sideLineUpData.over_view.hero6_gold_dets.death_3counts}}</li>
								<li>死亡超过三次的次数：{{sideLineUpData.over_view.hero6_gold_dets.death_more_than_3counts}}</li>
							</ul>
						</td>
					</tr>
				</table>
		</div>
		<div>
			<table bgcolor="#334656" id="tenmin_table">
               <tr>
                  <td>比赛ID</td>

                  <td>补刀经济</td>
                  <td>助攻及击杀经济</td>
                  <td>10分钟总经济</td>

                  <td>10分钟前等级信息</td>
                  <td>10分钟前技能加点</td>
                  <td>10分钟前消耗品信息</td>
               </tr>
               <tr v-for="item in sideLineUpData.line_up_dets">
               		<td v-for="item2 in item.side_line_up_dets">
               			{{item2.match_id}}
               		</td>
               		<td v-for="item2 in item.side_line_up_dets">
               			<table class="details_table">
               			 	<tr>
               			 		<td>{{item2.against_heroes[0].hero_name}}</td>
               			 		<td>{{item2.against_heroes[1].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[2]">{{item2.against_heroes[2].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[3]">{{item2.against_heroes[3].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[4]" >{{item2.against_heroes[4].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[5]">{{item2.against_heroes[5].hero_name}}</td>
               			 	</tr>
               			 	<tr>
               			 		<td>{{item2.against_heroes[0].pre_10min_gold.last_hit_gold}}</td>
               			 		<td>{{item2.against_heroes[1].pre_10min_gold.last_hit_gold}}</td>
               			 		<td v-if="item2.against_heroes[2]">{{item2.against_heroes[2].pre_10min_gold.last_hit_gold}}</td>
               			 		<td v-if="item2.against_heroes[3]">{{item2.against_heroes[3].pre_10min_gold.last_hit_gold}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{item2.against_heroes[4].pre_10min_gold.last_hit_gold}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{item2.against_heroes[5].pre_10min_gold.last_hit_gold}}</td>            			 		
               			 	</tr>
               			 </table>
               		</td>
               		<td v-for="item2 in item.side_line_up_dets">
               			<table class="details_table">
               			 	<tr>
               			 		<td>{{item2.against_heroes[0].hero_name}}</td>
               			 		<td>{{item2.against_heroes[1].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[2]">{{item2.against_heroes[2].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[3]">{{item2.against_heroes[3].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[4]" >{{item2.against_heroes[4].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[5]">{{item2.against_heroes[5].hero_name}}</td>
               			 	</tr>
               			 	<tr>
               			 		<td>{{item2.against_heroes[0].pre_10min_gold.combat_gold}}</td>
               			 		<td>{{item2.against_heroes[1].pre_10min_gold.combat_gold}}</td>
               			 		<td v-if="item2.against_heroes[2]">{{item2.against_heroes[2].pre_10min_gold.combat_gold}}</td>
               			 		<td v-if="item2.against_heroes[3]">{{item2.against_heroes[3].pre_10min_gold.combat_gold}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{item2.against_heroes[4].pre_10min_gold.combat_gold}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{item2.against_heroes[5].pre_10min_gold.combat_gold}}</td>            			 		
               			 	</tr>
               			 </table>
               		</td>
               		<td v-for="item2 in item.side_line_up_dets">
               			 <table class="details_table">
               			 	<tr>
               			 		<td>{{item2.against_heroes[0].hero_name}}</td><td>{{item2.against_heroes[1].hero_name}}</td><td>{{item2.against_heroes[2].hero_name}}</td><td>{{item2.against_heroes[3].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[4]" >{{item2.against_heroes[4].hero_name}}</td>
               			 		<td v-if="item2.against_heroes[5]">{{item2.against_heroes[5].hero_name}}</td>
               			 	</tr>
               			 	<tr>
               			 		<td>{{(item2.against_heroes[0].pre_10min_gold.last_hit_gold)+(item2.against_heroes[0].pre_10min_gold.creeps_gold)}}</td>
               			 		<td>{{(item2.against_heroes[1].pre_10min_gold.last_hit_gold)+(item2.against_heroes[1].pre_10min_gold.creeps_gold)}}</td>
               			 		<td>{{(item2.against_heroes[2].pre_10min_gold.last_hit_gold)+(item2.against_heroes[2].pre_10min_gold.creeps_gold)}}</td>
               			 		<td>{{(item2.against_heroes[3].pre_10min_gold.last_hit_gold)+(item2.against_heroes[3].pre_10min_gold.creeps_gold)}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{(item2.against_heroes[4].pre_10min_gold.last_hit_gold)+(item2.against_heroes[4].pre_10min_gold.creeps_gold)}}</td>
               			 		<td  v-if="item2.against_heroes[4]">{{(item2.against_heroes[5].pre_10min_gold.last_hit_gold)+(item2.against_heroes[5].pre_10min_gold.creeps_gold)}}</td>            			 		
               			 	</tr>
               			 </table>
               		</td>
					<td >
						<table v-for="item2 in item.side_line_up_dets">
							<tr >
								<td>{{item2.against_heroes[0].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[0].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
							<tr >
								<td>{{item2.against_heroes[1].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[1].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
							<tr v-if="item2.against_heroes[2]">
								<td>{{item2.against_heroes[2].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[2].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
							<tr v-if="item2.against_heroes[3]">
								<td>{{item2.against_heroes[3].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[3].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
							<tr v-if="item2.against_heroes[4]">
								<td>{{item2.against_heroes[4].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[4].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
							<tr v-if="item2.against_heroes[5]">
								<td>{{item2.against_heroes[5].hero_name}}</td>
								<td v-for="(item3,index3) in item2.against_heroes[5].pre_10min_lvl_str">
									<table>
										<tr><td>{{index3+2}}</td></tr>
										<tr><td>{{item3}}</td></tr>
									</table>
								</td>
							</tr>
						</table>
					</td>
               </tr>
           </table>
		</div>
	</div>
</template>
<script>
import {mapState} from 'vuex'
	export default{
		data(){
			return{

			}
		},
		computed:{
			...mapState(['sideLineUpData']),

			/*against_heroes:function(){
				var side_line_up_dets []
				var against_heroes []
				sideLineUpData.line_up_dets.forEach((item)=>{
					side_line_up_dets.append(item)
				})
				side_line_up_dets.forEach((item)=>{
					against_heroes.append(item.)
				})
			}*/
		},
		created(){

			var params=this.$store.state.lineUpRequest
			this.$store.dispatch('getSidelineUp',params)
			window.console.log(this.$store.state.sideLineUpData)
			
		},
	}
</script>
<style>
	#line_up_over_view{
		width:80%;
		text-align:center;
		margin:0 auto
	}
</style>