<template>
	<div class="match_content">
		<div id="radiant" class="radiant">
			<div id="radiant_team_name" class="team_name">
				{{OverViewData.result_data.radiant_team_name}}
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th><th>玩家</th><th>路线</th><th>等级</th><th>K</th><th>D</th><th>A</th><th>总资产</th><th>补刀</th><th></th><th>反补</th>
						<th>GPM</th><th></th><th>XPM</th><th>英雄伤害</th><th>防御塔伤害</th><th>英雄治疗</th><th>物品</th><th>增益状态</th>
					</tr>
					<tr v-for="(item,index) in radiantPlayers" border="1">
						<td  height="50"><img :src="item.hero_icon" width="50" ></td>
						<td>{{item.player.player_name}}</td>
						<td>{{OverViewData.details.players_heroes_dets[index].init_lane}}</td>
						<td>{{item.player.level}}</td>
						<td>{{item.player.kills}}</td>
						<td>{{item.player.deaths}}</td>
						<td>{{item.player.assists}}</td>
						<td>{{item.player.gold_spent}}</td>
						<td>{{item.player.last_hits}}</td>
						<td>/</td>
						<td>{{item.player.denies}}</td>
						<td>{{item.player.gold_per_min}}</td>
						<td>/</td>
						<td>{{item.player.XP_per_min}}</td>
						<td>{{item.player.hero_damage}}</td>
						<td>{{item.player.tower_damage}}</td>
						<td>{{item.player.hero_healing}}</td>
						<td width="240"><span v-for="item2 in item.item"><img :src="item2" width="38" class="item_img"/></span></td>
						<td><span v-for="item3 in item.buffs"><img :src="item3.buffs" width="38" class="item_img" /></span></td>
					</tr>
				</table>
			</div>
			<div id="radiant_bp">
				<div v-for="(item,index) in OverViewData.game_info.picks_bans" :key="item.hero_id" class="picks_bans" v-if="item.team==2">
					<img :src="item.hero_icon" width="55" :class="{ isban: !item.is_pick }"/>
					<div class="bp_order">
						<span v-if="item.is_pick">{{index+1}}pick</span>
						<span v-else>{{index+1}}ban</span>
					</div>
				</div>
				<div class="clear"></div>
			</div>
		</div>
		<div id="dire" class="dire">
			<div id="radiant_team_name" class="team_name">
				{{OverViewData.result_data.dire_team_name}}
			</div>
			<div class="player_data">
				<table>
					<tr>
						<th>英雄</th><th>玩家</th><th>路线</th><th>等级</th><th>K</th><th>D</th><th>A</th><th>总资产</th><th>补刀</th><th></th><th>反补</th>
						<th>GPM</th><th></th><th>XPM</th><th>英雄伤害</th><th>防御塔伤害</th><th>英雄治疗</th><th>物品</th><th>增益状态</th>
					</tr>
					<tr v-for="(item,index) in direPlayers" border="1">
						<td  height="50"><img :src="item.hero_icon" width="50" ></td>
						<td>{{item.player.player_name}}</td>
						<td>{{OverViewData.details.players_heroes_dets[index+5].init_lane}}</td>
						<td>{{item.player.level}}</td>
						<td>{{item.player.kills}}</td>
						<td>{{item.player.deaths}}</td>
						<td>{{item.player.assists}}</td>
						<td>{{item.player.gold_spent}}</td>
						<td>{{item.player.last_hits}}</td>
						<td>/</td>
						<td>{{item.player.denies}}</td>
						<td>{{item.player.gold_per_min}}</td>
						<td>/</td>
						<td>{{item.player.XP_per_min}}</td>
						<td>{{item.player.hero_damage}}</td>
						<td>{{item.player.tower_damage}}</td>
						<td>{{item.player.hero_healing}}</td>
						<td width="240"><span v-for="item2 in item.item"><img :src="item2" width="38" class="item_img"/></span></td>
						<td></td>
					</tr>
				</table>
			</div>
			<div id="radiant_bp">
				<div v-for="(item,index) in OverViewData.game_info.picks_bans" :key="item.hero_id" class="picks_bans" v-if="item.team==3">
					<img :src="item.hero_icon" width="55" :class="{ isban: !item.is_pick }"/>
					<div class="bp_order">
						<span v-if="item.is_pick">{{index+1}}pick</span>
						<span v-else>{{index+1}}ban</span>
					</div>
				</div>
				<div class="clear"></div>
			</div>
		</div>
		<div id="ability_builds" class="player_data">
			<table>
				<tr>
					<th width="100">英雄</th>
					<th>1</th>
					<th>2</th>
					<th>3</th>
					<th>4</th>
					<th>5</th>
					<th>6</th>
					<th>7</th>
					<th>8</th>
					<th>9</th>
					<th>10</th>
					<th>11</th>
					<th>12</th>
					<th>13</th>
					<th>14</th>
					<th>15</th>
					<th>16</th>
					<th>17</th>
					<th>18</th>
					<th>19</th>
					<th>20</th>
					<th>21</th>
					<th>22</th>
					<th>23</th>
					<th>24</th>
					<th>25</th>
				</tr>
				<tr v-for="(item,index) in OverViewData.details.players_heroes_dets">
					<td><img :src="item.hero_icon" width="50"/></td>
					<td><img v-if="item.ability_data[0]" :src="item.ability_data[0].Icon" width="35" alt="item.ability_data[0].Name"></td>
					<td><img :src="item.ability_data[1].Icon" width="35" alt="item.ability_data[1].Name"></td>
					<td><img :src="item.ability_data[2].Icon" width="35" alt="item.ability_data[2].Name"></td>
					<td><img :src="item.ability_data[3].Icon" width="35" alt="item.ability_data[3].Name"></td>
					<td><img :src="item.ability_data[4].Icon" width="35" alt="item.ability_data[4].Name"></td>
					<td><img :src="item.ability_data[5].Icon" width="35" alt="item.ability_data[5].Name"></td>
					<td><img v-if="item.ability_data[6]":src="item.ability_data[6].Icon" width="35" alt="item.ability_data[6].Name"></td>
					<td><img v-if="item.ability_data[7]":src="item.ability_data[7].Icon" width="35" alt="item.ability_data[7].Name"></td>
					<td><img v-if="item.ability_data[8]" :src="item.ability_data[8].Icon" width="35" alt="item.ability_data[8].Name"></td>
					<td><img v-if="item.ability_data[9]" :src="item.ability_data[9].Icon" width="35" alt="item.ability_data[9].Name"></td>
					<td ><img v-if="item.ability_data[10]" :src="item.ability_data[10].Icon" width="35" alt="item.ability_data[10].Name"></td>
					<td ><img v-if="item.ability_data[11]" :src="item.ability_data[11].Icon" width="35" alt="item.ability_data[11].Name"></td>
					<td ><img v-if="item.ability_data[12]" :src="item.ability_data[12].Icon" width="35" alt="item.ability_data[12].Name"></td>
					<td ><img v-if="item.ability_data[13]" :src="item.ability_data[13].Icon" width="35" alt="item.ability_data[13].Name"></td>
					<td ><img v-if="item.ability_data[14]"  :src="item.ability_data[14].Icon" width="35" alt="item.ability_data[14].Name"></td>
					<td ><img v-if="item.ability_data[15]" :src="item.ability_data[15].Icon" width="35" alt="item.ability_data[15].Name"></td>
					<td></td>
					<td ><img v-if="item.ability_data[16]" :src="item.ability_data[16].Icon" width="35" alt="item.ability_data[16].Name"></td>
					<td></td>
					<td><img  v-if="item.ability_data[17]" :src="item.ability_data[17].Icon" width="35" alt="item.ability_data[17].Name"></td>
					<td></td>
					<td></td>
					<td></td>
					<td></td>
					<td><img v-if="item.ability_data[18]" :src="item.ability_data[18].Icon" width="35" alt="item.ability_data[18].Name"></td>

					<!-- <td v-for="(item2,index2) in item.ability_data"><img :src="item2.Icon" width="35" /></td> -->
				</tr>
			</table>
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
				return this.OverViewData.result_data.players.slice(0,5)
			},
			direPlayers(){
				return this.OverViewData.result_data.players.slice(5)
			},
/*			isempty(){

			}*/

		},
		props:['OverViewData'],
	}
</script>
<style scoped>
	.match_content{
		width:100%;
	}
	.team_name{
		width:1200px;
		margin:0 auto;
		text-align:left;
		height:25px;
	}
	.item_img{
		padding-right:2px
	}

	.picks_bans{
		float:left;
		width:60px;
		height:50px;
	}
	.isban{
		filter: grayscale(1);
		padding:0px;
		margin:0px;
	}
	.bp_order{
		width:55px;
		background:#000000;
		
		margin:-4px auto 0px auto;
		padding:0px;
		font-size:13px
	}
	#radiant_bp{
		width:1200px;
		margin:8px auto 0px auto;

	}
	#ability_builds{
		width:1200px;
		margin:15px auto;
	}
	.radiant{
		margin-bottom:15px;
	}
</style>