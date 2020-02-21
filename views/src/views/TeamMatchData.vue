<template>
	<el-main>
		<div id="team_info" v-if="teamData!=null" class="wrapper">
			<div id="base_info">
				<div class="logo"><img :src="teamData.team_base_info.team_icon" width="120"></div>
				<div class="player_info">
					<div>选手英雄池</div>
					<div>
						<div class="hero_pool"  v-for="item in teamData.team_hero_pool" :key="item.player_register_string_id">
							<div id="position">{{item.player_register_string_id}}({{item.position}})：</div>
							<div id="hero_pool_dets" v-for="item2 in item.match_player_heroes.hero_play_count">
								{{item2.hero}}*{{item2.count}}
							</div>
						</div>
					</div>
				</div>
				<div style="clear:both"></div>
			</div>
			<Page :pageData="teamData.team_bp"></Page> 
		</div>
	</el-main>
</template>
<script>
import {mapState} from 'vuex'
const Page=()=>import('@/components/Page.vue')
	export default{
		data(){
			return{

			}
		},

		computed:{
			...mapState(['teamData'])
		},
		components:{
			Page
		},
		created(){
			var param={team_name:this.$route.query.team,version:this.$store.state.version}
			this.$store.dispatch('getTeamData',param)
		}
	}
</script>
<style>
.logo{
	width:200px;
	float:left
}
.player_title{
	float:left
}
.player_info{
	float:left;
	width:1200px
}
.hero_pool{
widht:100%;
height:30px
}
#position{
	float:left
}
#base_info{
	width:100%;
	height:auto;
}
#hero_pool_dets{
	float:left;
	margin-left:10px;
}
</style>