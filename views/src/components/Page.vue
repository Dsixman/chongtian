<template>
	<div id="bp">
		<div v-for="item in currentPageData" style="height:110px">
			<div class="teamname"><div class="left_team_name">{{item.radiant_team_tag}}</div><div class="right_team_name">{{item.dire_team_tag}}</div></div>
			<div v-for="(item2,index2) in item.picks_bans" :key="item2.hero_id" class="picks_bans" v-if="item2.team==2">
				<img :src="item2.hero_icon" width="55" :class="{ isban: !item2.is_pick }"/>
				<div class="bp_order">
					<span v-if="item2.is_pick">{{index2+1}}pick</span>
					<span v-else>{{index2+1}}ban</span>
				</div>
			</div>
			<div style="width:80px;float:left;height:50px;line-height:45px;">胜者：{{item.game_winner}}</div>
			<div v-for="(item2,index2) in item.picks_bans" :key="item2.hero_id" class="picks_bans" v-if="item2.team==3">
				<img :src="item2.hero_icon" width="55" :class="{ isban: !item2.is_pick }"/>
				<div class="bp_order">
					<span v-if="item2.is_pick">{{index2+1}}pick</span>
					<span v-else>{{index2+1}}ban</span>
				</div>
			</div>

			<div class="clear"></div>
			<div class="lane">
				<span class="lane_span">比赛ID：<router-link :to="{name:'matchdetails',query:{matchid:item.match_id}}" style="color:white">{{item.match_id}}</router-link></span>
				<span class="lane_span">上路对线：{{item.top_lane}}</span>
				<span class="lane_span">中路对线：{{item.mid_lane}}</span>
				<span class="lane_span">下路对线：{{item.bottom_lane}}</span>
			</div>
		</div>
		<!-- <div><button @click="test()">提示data</button></div> -->
		<!-- <div v-for="pageData">
			<div></div>
		</div> -->
		<div class="block">
		    <el-pagination
		      @size-change="handleSizeChange"
		      @current-change="handleCurrentChange"
		      :current-page="currentPage4"
		      :page-sizes="[2]"
		      :page-size="pageSize"
		      layout="total, sizes, prev, pager, next, jumper"
		      :total="total">
		    </el-pagination>
	    </div>
    </div>
</template>
<script>
	export default{
		data(){
			return {
				currentPage4: 4,
				pageSize:2,
				currentIndex:0
			}
		},
		computed:{
			currentPageData:function(){
				var startIndex=this.currentIndex
				var endIndex=this.currentIndex+this.pageSize
				return this.pageData.slice(startIndex,endIndex)
			},
			total(){
				return this.pageData.length
			},
		},
		props:['pageData'],
		methods: {
	      handleSizeChange(val) {
	        //console.log(`每页 ${val} 条`);
	      },
	      handleCurrentChange(val) {
	      	this.currentIndex=val
	        //console.log(`当前页: ${val}`);
	      },
	      /*test(){
	      	console.log(this.pageData)
	      }*/
   	 	},
   	 	/*watch: {
	      pageData: function(newVal,oldVal){
	        this.currentPageData = newVal; //newVal即是chartData
	        this.drawChart();
	      }
	    },*/
	}
</script>
<style>
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
	.clear{
		clear:both
	}
	.bp_order{
		width:55px;
		background:#000000;
		
		margin:-4px auto 0px auto;
		padding:0px;
		font-size:13px
	}
	.left_team_name{
		float:left;
		width:50%;
		text-align:center;
	}
	.right_team_name{
		float:left;
		width:50%;
		text-align:center;
	}
	.teamname{
		width:100%;
	}
	.lane{
		width:100%;
		line-height:25px
	}
	.lane_span{
		padding-right:15px
	}
</style>