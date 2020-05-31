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
						<th>5:00</th>
						<th>10:00</th>
						<th>15:00</th>
						<th>20:00</th>
					</tr>
					<tr v-for="item in details.players_heroes_dets">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>
							<ul v-for="item2 in item.inventory_snap_shot" class="kda">
								<li v-if="item2.game_time==300">击杀：{{item2.kills}}</li>
								<li v-if="item2.game_time==300">死亡：{{item2.deaths}}</li>
								<li v-if="item2.game_time==300">助攻：{{item2.assists}}</li>
							</ul> 
						</td>
						<td>
							<ul v-for="item2 in item.inventory_snap_shot" class="kda">
								<li v-if="item2.game_time==600">击杀：{{item2.kills}}</li>
								<li v-if="item2.game_time==600">死亡：{{item2.deaths}}</li>
								<li v-if="item2.game_time==600">助攻：{{item2.assists}}</li>
							</ul>
						</td>
						<td>
							<ul v-for="item2 in item.inventory_snap_shot" class="kda" >
								<li v-if="item2.game_time==1200">击杀：{{item2.kills}}</li>
								<li v-if="item2.game_time==1200">死亡：{{item2.deaths}}</li>
								<li v-if="item2.game_time==1200">助攻：{{item2.assists}}</li>
							</ul>
						</td>
						<td>
							<ul v-for="item2 in item.inventory_snap_shot" class="kda">
								<li v-if="item2.game_time==1500">击杀：{{item2.kills}}</li>
								<li v-if="item2.game_time==1500">死亡：{{item2.deaths}}</li>
								<li v-if="item2.game_time==1500">助攻：{{item2.assists}}</li>
							</ul>
						</td>
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
<style scoped>
ul.kda{
	margin:0px;padding:0px
}
ul.kda li{margin:10px 0px 10px 0px;padding:0px}
</style>