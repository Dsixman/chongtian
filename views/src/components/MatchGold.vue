<template>
	<div class="match_content">
		<div class="radiant">
			<div id="radiant_team_name" class="team_name">
				天辉
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th>
						<th>5分钟经经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>10分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>15分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>20分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>25分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
					</tr>
					<tr v-for="item in radiantPlayers">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>{{item.gold.pre_5min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_5min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_10min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_10min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_15min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_15min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_20min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_20min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_25min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_25min_gold.combat_gold}}</td>
					</tr>
				</table>
			</div>
		</div>
		<div class="dire">
			<div class="team_name">
				夜魇
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th>
						<th>5分钟经经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>10分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>15分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>20分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
						<th>25分钟经济</th>
						<th>/</th>
						<th>战斗经济</th>
					</tr>
					<tr v-for="item in direPlayers">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>{{item.gold.pre_5min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_5min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_10min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_10min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_15min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_15min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_20min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_20min_gold.combat_gold}}</td>
						<td>{{item.gold.pre_25min_gold.all_gold}}</td>
						<td>/</td>
						<td>{{item.gold.pre_25min_gold.combat_gold}}</td>
					</tr>
				</table>
			</div>
		</div>
	</div>
</template>
<script>
	export default{
		data(){
			return {

			}
		},
		computed:{
			radiantPlayers(){
				return this.details.players_heroes_dets.slice(0,5)
			},
			direPlayers(){
				return this.details.players_heroes_dets.slice(5)
			},
		},
		props:['details']
	}
</script>
<style>
</style>