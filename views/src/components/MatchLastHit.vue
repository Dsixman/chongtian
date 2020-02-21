<template>
	<div class="match_content">
		<div id="radiant" class="radiant">
			<div id="radiant_team_name" class="team_name">
				天辉
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th>
						<th>5分钟补刀</th>
						<th>10分钟补刀</th>
						<th>15分钟补刀</th>
						<th>20分钟补刀</th>
						<th>25分钟补刀</th>
					</tr>
					<tr v-for="item in radiantPlayers" border="1">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>{{item.last_hit.Pre5Count}}</td>
						<td>{{item.last_hit.pre_10_count}}</td>
						<td>{{item.last_hit.pre_15_count}}</td>
						<td>{{item.last_hit.pre_20_count}}</td>
						<td>{{item.last_hit.pre_25_count}}</td>
					</tr>
				</table>
			</div>
		</div>
		<div id="dire" class="dire">
			<div id="dire_team_name" class="team_name">
				夜魇
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th>
						<th>5分钟补刀</th>
						<th>10分钟补刀</th>
						<th>15分钟补刀</th>
						<th>20分钟补刀</th>
						<th>25分钟补刀</th>
					</tr>
					<tr v-for="item in direPlayers" border="1">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>{{item.last_hit.Pre5Count}}</td>
						<td>{{item.last_hit.pre_10_count}}</td>
						<td>{{item.last_hit.pre_15_count}}</td>
						<td>{{item.last_hit.pre_20_count}}</td>
						<td>{{item.last_hit.pre_25_count}}</td>
					</tr>
				</table>
			</div>
		</div>
	</div>
</template>
<script>
	export default{
		data(){
			return{

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