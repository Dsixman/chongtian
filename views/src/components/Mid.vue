<template>
	<div v-if="midUseHeroes">
		<div v-for="(value, name) in midUseHeroes" >
			<div class="against_hero_name">{{name}}</div>
			<div>
				<table bgcolor="#334656">
	               <tr>
	                  <td>比赛ID</td>
	                  <td>5分钟补刀</td>
	                  <td>10分钟补刀</td>
	                  <td>10分钟前等级信息</td>
	                  <td>10分钟前技能加点</td>
	                  <td>10分钟前消耗品信息</td>
	               </tr>
	               <tr v-for="item in value">
	               	<td>{{item.match_id}}</td>
	               	<td>
	               		<table class="details_table">
	               			<tr><td>{{item.mid_hero_1.hero_name}}</td><td>{{item.mid_hero_2.hero_name}}</td></tr>
	               			<tr><td>{{item.mid_hero_1.pre_5min_lasthit}}</td><td>{{item.mid_hero_2.pre_5min_lasthit}}</td></tr>
	               		</table>
	               	</td>
	               	<td>
	               		<table class="details_table">
	               			<tr><td>{{item.mid_hero_1.hero_name}}</td><td>{{item.mid_hero_2.hero_name}}</td></tr>
	               			<tr><td>{{item.mid_hero_1.pre_10min_lasthit}}</td><td>{{item.mid_hero_2.pre_10min_lasthit}}</td></tr>
	               		</table>
	               	</td>
	               	<td>
	               		<table class="details_table">
	               			<tr>
	               				<td>
	               					{{item.mid_hero_1.hero_name}}
	               				</td>
	               				<td v-for="(value2,index) in item.mid_hero_1.pre_10min_levelup_time">
	               					<table>
	               						<tr><td>{{index+2}}</td></tr>
                                 		<tr><td>{{value2}}</td></tr>
	               					</table>
	               				</td>
	               			</tr>
	               		</table>
	               		<table class="details_table">
	               			<tr>
	               				<td>
	               					{{item.mid_hero_2.hero_name}}
	               				</td>
	               				<td v-for="(value3,index3) in item.mid_hero_2.pre_10min_levelup_time">
	               					<table>
	               						<tr><td>{{index3+2}}</td></tr>
                                 		<tr><td>{{value3}}</td></tr>
	               					</table>
	               				</td>
	               			</tr>
	               		</table>
	               	</td>
	               	<td>
	                  	<table class="details_table">
	                  		<tr>
	                  			<td>{{item.mid_hero_1.hero_name}}</td>
	                  			<td v-for="(value,index) in item.mid_hero_1.ability_upgrades">
	                  				<table>
	                  					
	                  					<tr>
	                  						<td>
	                  							<img :src="value.Icon" width="30">
	                  						</td>
	                  					</tr>
	                  				</table>
	                  			</td>
	                  		</tr>
	                  		<tr>
	                  			<td>{{item.mid_hero_2.hero_name}}</td>
	                  			<td v-for="(value1,index1) in item.mid_hero_2.ability_upgrades">
	                  				<table>
	                  					
	                  					<tr>
	                  						<td>
	                  							<img :src="value1.Icon" width="30">
	                  						</td>
	                  					</tr>
	                  				</table>
	                  			</td>
	                  		</tr>	
	                  	</table>
                  	</td>
                  <td>
                     <table class="details_table">
                     	<tr>
                     		<td>{{item.mid_hero_1.hero_name}}</td>
                     		<td v-for="consumable_value in item.mid_hero_1.consumable">
                     			<img :src="consumable_value.item_icon" width="30">*{{consumable_value.count}}
                     		</td>
                     	</tr>
                     	<tr>
                     		<td>{{item.mid_hero_2.hero_name}}</td>
                     		<td v-for="consumable_value2 in item.mid_hero_2.consumable">
                     			<img :src="consumable_value2.item_icon" width="30">*{{consumable_value2.count}}
                     		</td>
                     	</tr>
                     </table>   
                  </td>
	               </tr>
	               <tr>
	               		<td>数据总结</td><td colspan="5"><div></div></td>
	               </tr>
	            </table> 
			</div>
		</div>	
	</div>
</template>
<script>
import {mapState} from 'vuex'
	export default{
		data(){
			return {
				
			}
		},
		computed:{
			...mapState(['midUseHeroes'])

		},
		created(){
			var params={playerid:this.$route.params.playerid,version:this.$route.params.version}
			this.$store.dispatch('getMidUseHeroes',params)
		}
	}
</script>
<style>
	.against_hero_name{
		text-align:left
	}
</style>