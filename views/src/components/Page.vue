<template>
	<div id="bp">
		<div class="bp_title title-bar">战队BP数据</div>
		<div>
			<p class="sencond-title">英雄前三手被ban比率</p>
			<div >
				<span v-for="item in prethirdban" style="padding-right:10px;"><img  :src="item.icon" alt="" width="50">{{item.ban_rate_str}}</span>
			</div>
		</div>
		<div>
			<p class="sencond-title">英雄前三手pick比率</p>
			<div>
				<span v-for="item in prethirdpick" style="padding-right:10px;"><img  :src="item.icon" alt="" width="50">{{item.pick_rate_str}}</span>
			</div>
		</div>
		<div class="bp_title title-bar">比赛详细BP数据</div>
		<div v-for="item in currentPageData" class="bp-details" >
			<div class="teamname">
				<div class="left_team_name">{{item.radiant_team_tag}}</div>
				<div style="width:120px;float:left;font-size:15px; text-align:left">
					<span>胜者:{{item.game_winner}}</span>
				</div>
				<div class="right_team_name">{{item.dire_team_tag}}</div>
				<div class="clear"></div>
			</div>
			<div v-for="(item2,index2) in item.picks_bans" :key="item2.hero_id" class="picks_bans" v-if="item2.team==2">
				<img :src="item2.hero_icon" width="55" :class="{ isban: !item2.is_pick }"/>
				<div class="bp_order">
					<span v-if="item2.is_pick">{{index2+1}}pick</span>
					<span v-else>{{index2+1}}ban</span>
				</div>
			</div>
			<div style="width:50px;float:left;height:50px"></div>
			<div v-for="(item2,index2) in item.picks_bans" :key="item2.hero_id" class="picks_bans" v-if="item2.team==3">
				<img :src="item2.hero_icon" width="55" :class="{ isban: !item2.is_pick }"/>
				<div class="bp_order">
					<span v-if="item2.is_pick">{{index2+1}}pick</span>
					<span v-else>{{index2+1}}ban</span>
				</div>
			</div>

			<div class="clear"></div>
			<div class="lane">
				<span class="lane_span">比赛ID：<router-link :to="{name:'matchdetails',query:{matchid:item.match_id}}">{{item.match_id}}</router-link></span>
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
			<!-- <el-pagination
			    layout="prev, pager, next"
			    :total="total">
					  	</el-pagination> -->
		    <el-pagination
		      @size-change="handleSizeChange"
		      @current-change="handleCurrentChange"
		      :current-page="currentPage4"
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
				currentPage4: 1,
				pageSize:10,
				currentIndex:0
			}
		},
		computed:{
			currentPageData:function(){
				var lastpage=Math.ceil(this.pageData.length/this.pageSize)
				var lastpagecount=0-(this.pageData.length%this.pageSize)
				if (this.currentIndex!=lastpage){
					var startIndex=this.currentIndex
					var endIndex=this.currentIndex+this.pageSize
					return this.pageData.slice(startIndex,endIndex)
				}else{
					return this.pageData.slice(lastpagecount)
				}
				
				
			},
			total(){
				return this.pageData.length
			},
						prethirdban(){
				var teambp=this.$store.state.teamData.team_bp
				var teambpnamearr=[]
				var teamname=this.$route.params.team
				teambp.forEach((item)=>{
					var gameteam
					if  (item.radiant_team_tag==teamname){
						gameteam=2
					}else{
						gameteam=3
					}
					item.picks_bans.forEach((item2,index2)=>{
						if (item2.is_pick==false&&gameteam==gameteam&&index2<=5){

							var teambpnamearrlen=teambpnamearr.length		
							if (teambpnamearrlen==0){
								var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,ban_rate:null,ban_rate_str:null}

								teambpnamearr.push(obj)
							}else{

								teambpnamearrlen=teambpnamearrlen-1
								for (var i=0;i<=teambpnamearrlen;i++){
									if (teambpnamearr[i].id==item2.hero_id){

										teambpnamearr[i].count=teambpnamearr[i].count+1
										break;
									}
									if (teambpnamearr[i].id!=item2.hero_id&&i==teambpnamearrlen){
										var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,ban_rate:null,ban_rate_str:null}
										teambpnamearr.push(obj)
										break;
									}
								}
							}
							
						}
					})
				})
				teambpnamearr.forEach((item)=>{
					var length=teambpnamearr.length
					item.ban_rate=(item.count/length).toFixed(2)
					item.ban_rate_str=(item.ban_rate*100)+"%"
					if (item.ban_rate_str.indexOf(".")>-1){
						item.ban_rate_str=item.ban_rate_str.substring(0, 2)+"%"
					}
				})
				return teambpnamearr
			},
			prethirdpick(){
				var teambp=this.$store.state.teamData.team_bp
				var teambpnamearr=[]
				var teamname=this.$route.params.team
				teambp.forEach((item)=>{
					var gameteam
					if  (item.radiant_team_tag==teamname){
						gameteam=2
					}else{
						gameteam=3
					}
					item.picks_bans.forEach((item2,index2)=>{
						if (item2.is_pick==true&&gameteam==gameteam&&index2<=9){
							var teambpnamearrlen=teambpnamearr.length		
							if (teambpnamearrlen==0){
								var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,pick_rate:null,pick_rate_str:null}
								teambpnamearr.push(obj)
							}else{

								teambpnamearrlen=teambpnamearrlen-1
								for (var i=0;i<=teambpnamearrlen;i++){
									if (teambpnamearr[i].id==item2.hero_id){
										teambpnamearr[i].count=teambpnamearr[i].count+1
										break;
									}
									if (teambpnamearr[i].id!=item2.hero_id&&i==teambpnamearrlen){
										var obj={id:item2.hero_id,count:1,hero:item2.hero_name,icon:item2.hero_icon,pick_rate:null,pick_rate_str:null}
										teambpnamearr.push(obj)
										break;
									}
								}
							}
							
						}
					})
				})
				teambpnamearr.forEach((item)=>{
					var length=teambpnamearr.length
					item.pick_rate=(item.count/length).toFixed(2)
					item.pick_rate_str=(item.pick_rate*100)+"%"
					if (item.pick_rate_str.indexOf(".")>-1){
						item.pick_rate_str=item.pick_rate_str.substring(0, 2)+"%"
					}
				})
				return teambpnamearr
			}
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
<style >
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
		width:46%;
		text-align:center;
	}
	.right_team_name{
		float:left;
		width:44%;
		text-align:center;
	}
	.teamname{
		height:30px;
		width:100%;
		line-height:30px;
	}
	.lane{
		width:100%;
		line-height:25px
	}
	.lane_span{
		padding-right:15px
	}
	.bp_title{
		
		margin-top:8px;
		margin-bottom:15px;
	}
	.bp-details{
		box-shadow: 0 0 5px rgba(10,10,10,0.4);
    	border: 1px solid #14212d;
    	height:110px;
    	margin-bottom:10px;
	}
	#bp{
		margin-top:20px;
	}
</style>