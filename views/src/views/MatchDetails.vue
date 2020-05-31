<template>
	<div style="width:1400px;margin:0 auto" v-if="matchDetails">
		<div id="match_top" >
			<div class="winner_name"  v-if="matchDetails">{{matchDetails.game_info.game_winner}}胜利</div>
			<div class="match-victory-subtitle" v-if="matchDetails">
				<span >{{matchDetails.result_data.radiant_team_name}}</span>
				<span class="radiant_score">{{matchDetails.result_data.radiant_team_score}}</span>
				<span class="duration">{{matchDetails.result_data.duration_str}}</span>
				<span class="dire_score">{{matchDetails.result_data.dire_team_score}}</span>
				<span>{{matchDetails.result_data.dire_team_name}}</span>
			</div>
		</div>
		<el-tabs v-model="activeName" @tab-click="handleClick">
	    <el-tab-pane label="概观" name="first">
			<OverView :OverViewData="matchDetails"></OverView>	
	    </el-tab-pane>
	    <el-tab-pane label="补刀" name="second">
			<last-hit :details="matchDetails.details" ></last-hit>
	    </el-tab-pane>

	    <el-tab-pane label="经济" name="third">
			<Gold :details="matchDetails.details"></Gold>
	    </el-tab-pane>
	    <el-tab-pane label="等级" name="fourth">
	    	<Level :details="matchDetails.details"></Level>
	    </el-tab-pane>
	    <el-tab-pane label="KDA" name="five">
	    	<KDA :details="matchDetails.details"></KDA>
	    </el-tab-pane>
	    <el-tab-pane label="装备更新" name="six">
			<Item :details="matchDetails.details"></Item>
		</el-tab-pane>
	  	</el-tabs>
		
	</div>
	
</template>
<script>
import {mapState} from 'vuex'
const OverView=()=>import('@/components/OverView.vue')
const LastHit=()=>import('@/components/MatchLastHit.vue')
const Gold=()=>import('@/components/MatchGold.vue')
const Level=()=>import('@/components/MatchLevel.vue')
const KDA=()=>import('@/components/MatchKDA.vue')
const Item=()=>import('@/components/MatchItem.vue')
	export default{
		data() {
	      return {
	        activeName: 'first'
	      };
    	},
    	computed:{
    		...mapState(['matchDetails'])
    	},
    	components:{
    		OverView,
    		LastHit,
    		Gold,
    		Level,
    		Item,
    		KDA
    	},
	    methods: {
	      handleClick(tab, event) {
	       // console.log(tab, event);
	      }
	    },
		created(){
			var param={matchid:this.$route.query.matchid}
			this.$store.dispatch('getMatchDetails',param)
		}
	}
</script>
<style>
.el-tabs__item{
	color:#989898;
	width:232px;
}
.el-tabs__item:hover{
	color:#ffffff
}
.el-tabs__item.is-active {
	color:#ffffff
}
.el-tabs__nav{
	
}
.el-tabs__active-bar{
	background-color: #ffffff
}
.winner_name{
	font-size:25px
}
.match-victory-subtitle{
	margin-top:10px;
}
.radiant_score{
	padding-left:10px;
	color: rgb(102, 187, 106);
	font-size:20px;
}
.dire_score{
	color: rgb(255, 76, 76);
	font-size:20px;
	padding-right:10px;
}
.duration{
	padding-left:15px;
	padding-right:15px;
}

</style>