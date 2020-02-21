<template>
	<el-main style="padding:0px;">
		<div v-if="tenMinData.mid_hero_data!=null" >
			<!-- {{tenMinData.mid_hero_data}} -->
			<p><h2>中路对线数据</h2></p>	
			<table bgcolor="#334656" >
				<tr>
					<td>比赛ID</td>
					<td>5分钟补刀</td>
					<td>10分钟补刀</td>
					<td>5分钟总经济</td>
					<td>5分钟补刀经济</td>
					<td>5分钟助攻及击杀经济</td>
					<td>10分钟总经济</td>
					<td>10分钟补刀经济</td>
					<td>10分钟助攻及击杀经济</td>
					<td>10分钟前等级信息</td>
					<td>10分钟前技能点信息</td>
					<td>10分钟前消耗品信息</td>
				</tr>
				<tr v-for="item in tenMinData.mid_hero_data" :key="item.match_id">
					<td>{{item.match_id}}</td>
					<td>
						<table class="details_table">
							<tr>
								<td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
								<td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>	
							</tr>
						</table>
					</td>
				</tr>

			<!-- 	<tr>
				<td>比赛ID</td>
				<td colspan="9">{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
				<td colspan="9">{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
			</tr>
			<tr>
				<td></td>
				<td>5分钟补刀数</td>
				<td>10分钟补刀数</td>
				<td>5分钟总金钱</td>
				<td>5分钟补刀金钱</td>
				<td>5分钟助攻及击杀金钱</td>
				<td>10分钟总金钱</td>
				<td>10分钟补刀金钱</td>
				<td>10分钟助攻及击杀金钱</td>
			</tr>
			<tr v-for="item in tenMinData.mid_hero_data" :key="item.match_id">
				<td>{{item.match_id}}</td>
				<td>{{item.mid_hero_1.pre_5min_lasthit}}</td>
				<td>{{item.mid_hero_1.pre_10min_lasthit}}</td>
				<td>{{Math.ceil(item.mid_hero_1.pre_5min_gold.all_gold)}}</td>
				<td>{{item.mid_hero_1.pre_5min_gold.last_hit_gold}}</td>
				<td>{{item.mid_hero_1.pre_5min_gold.combat_gold}}</td>
			
				<td>{{Math.ceil(item.mid_hero_1.pre_10min_gold.all_gold)}}</td>
				<td>{{item.mid_hero_1.pre_10min_gold.last_hit_gold}}</td>
				<td style="border-right:1px soild #7d8f9d">{{item.mid_hero_1.pre_10min_gold.combat_gold}}</td>
			
				<td>{{item.mid_hero_2.pre_5min_lasthit}}</td>
				<td>{{item.mid_hero_2.pre_10min_lasthit}}</td>
				<td>{{Math.ceil(item.mid_hero_2.pre_5min_gold.all_gold)}}</td>
				<td>{{item.mid_hero_2.pre_5min_gold.last_hit_gold}}</td>
				<td>{{item.mid_hero_2.pre_5min_gold.combat_gold}}</td>
				<td>{{Math.ceil(item.mid_hero_2.pre_10min_gold.all_gold)}}</td>
				<td>{{item.mid_hero_2.pre_10min_gold.last_hit_gold}}</td>
				<td>{{item.mid_hero_2.pre_10min_gold.combat_gold}}</td>
			</tr> -->
			</table>	
		</div>
	</el-main>
</template>
<script>
import {mapState} from 'vuex'
	export default{
		data(){
			return {

			}
		},
		computed:{
			...mapState(['tenMinData'])
		}
	}
</script>
<style scoped>
#midlane{border:1px solid #ff0000}
.midlaneinfo{height:auto;border:1px solid #ff0000;}
table{
	border-spacing:0px;
	border-collapse:collapse;
	width:100%
}
table tr{
	border-bottom: 1px solid #7d8f9d;

	height:30px;
}
td{width:100px;}
tr:hover{background-color: #729cbf;} 
.details_table{
	border:1px solid #ff0000;
	font-size:0.9rem
}
</style>
