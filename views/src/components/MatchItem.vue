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
						<th>25:00</th>
						<th>30:00</th>
						<th>35:00</th>
					</tr>
					<tr v-for="item in details.players_heroes_dets">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>
							<span v-for="item2 in item.items" v-if="item2.purchase_time_Num<=300">
								<img :src="item2.item_icon" alt="" width="30">
							</span>
						</td>
						<td>
							<span v-for="item2 in item.items" v-if="(item2.purchase_time_Num>300) && (item2.purchase_time_Num<=600)">
								<img :src="item2.item_icon" alt="" width="30">
							</span>
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
<style>
</style>