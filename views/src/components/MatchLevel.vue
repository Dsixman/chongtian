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
					<tr v-for="item in radiantPlayers">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>
							{{item.level_up_time_str[0]}}
						</td>
						<td>
							{{item.level_up_time_str[1]}}
						</td>
						<td>
							{{item.level_up_time_str[2]}}
						</td>
						<td>
							{{item.level_up_time_str[3]}}
						</td>
						<td>
							{{item.level_up_time_str[4]}}
						</td>
						<td>
							{{item.level_up_time_str[5]}}
						</td>
						<td>
							<span v-if="item.level_up_time_str[6]">{{item.level_up_time_str[6]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[7]">{{item.level_up_time_str[7]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[8]">{{item.level_up_time_str[8]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[9]">{{item.level_up_time_str[9]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[10]">{{item.level_up_time_str[10]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[11]">{{item.level_up_time_str[11]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[12]">{{item.level_up_time_str[12]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[13]">{{item.level_up_time_str[13]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[14]">{{item.level_up_time_str[14]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[15]">{{item.level_up_time_str[15]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[16]">{{item.level_up_time_str[16]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[17]">{{item.level_up_time_str[17]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[18]">{{item.level_up_time_str[18]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[19]">{{item.level_up_time_str[19]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[20]">{{item.level_up_time_str[20]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[21]">{{item.level_up_time_str[21]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[22]">{{item.level_up_time_str[22]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[23]">{{item.level_up_time_str[23]}}</span>
						</td>
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
					<tr v-for="item in direPlayers">
						<td><img :src="item.hero_icon" width="50" /></td>
						<td>
							{{item.level_up_time_str[0]}}
						</td>
						<td>
							{{item.level_up_time_str[1]}}
						</td>
						<td>
							{{item.level_up_time_str[2]}}
						</td>
						<td>
							{{item.level_up_time_str[3]}}
						</td>
						<td>
							{{item.level_up_time_str[4]}}
						</td>
						<td>
							{{item.level_up_time_str[5]}}
						</td>
						<td>
							<span v-if="item.level_up_time_str[6]">{{item.level_up_time_str[6]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[7]">{{item.level_up_time_str[7]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[8]">{{item.level_up_time_str[8]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[9]">{{item.level_up_time_str[9]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[10]">{{item.level_up_time_str[10]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[11]">{{item.level_up_time_str[11]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[12]">{{item.level_up_time_str[12]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[13]">{{item.level_up_time_str[13]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[14]">{{item.level_up_time_str[14]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[15]">{{item.level_up_time_str[15]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[16]">{{item.level_up_time_str[16]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[17]">{{item.level_up_time_str[17]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[18]">{{item.level_up_time_str[18]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[19]">{{item.level_up_time_str[19]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[20]">{{item.level_up_time_str[20]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[21]">{{item.level_up_time_str[21]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[22]">{{item.level_up_time_str[22]}}</span>
						</td>
						<td>
							<span v-if="item.level_up_time_str[23]">{{item.level_up_time_str[23]}}</span>
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