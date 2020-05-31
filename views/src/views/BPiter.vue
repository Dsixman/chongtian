<template>
	 <div>{{allpick}}</div>
</template>
<script>
	import {mapState} from 'vuex'
	export default{
		data(){
			return{

			}
		},
		computed:{
			...mapState(['teamData']),
			allpick(){
				var heropool=this.teamData.team_hero_pool
				var hero1pool=heropool[0].match_player_heroes.hero_play_count
				var hero2pool=heropool[1].match_player_heroes.hero_play_count
				var hero3pool=heropool[2].match_player_heroes.hero_play_count
				var hero4pool=heropool[3].match_player_heroes.hero_play_count
				var hero5pool=heropool[4].match_player_heroes.hero_play_count
				var allresult=[]
				for (var i=0;i<hero1pool.length;i++){		
					var v=hero1pool[i].hero
					var vicon=hero1pool[i].hero_icon
					for(var i2=0;i2<hero2pool.length;i2++){
						var v2=hero2pool[i2].hero
						if (v!=v2){
							var v2icon=hero2pool[i2].hero_icon
							for(var i3=0;i3<hero3pool.length;i3++){
								var v3=hero3pool[i3].hero
								if  (v!=v2&&v!=v3&&v2!=v3){
									var v3icon=hero3pool[i3].hero_icon
									for (var i4=0;i4<hero4pool.length;i4++){
										var v4=hero4pool[i4].hero
										if (v!=v2&&v!=v3&&v!=v4&&v2!=v3&&v2!=v4&&v3!=v4){
											var v4icon=hero4pool[i4].hero_icon
											for (var i5=0;i5<hero5pool.length;i5++){
												var v5=hero5pool[i5].hero
												if (v!=v2&&v!=v3&&v!=v4&&v!=v5&&v2!=v3&&v2!=v4&&v2!=v5&&v3!=v4&&v3!=v5&&v4!=v5){
													var v5icon=hero5pool[i5].hero_icon
													var finalarr=[]
													finalarr.push(vicon,v2icon,v3icon,v4icon,v5icon)
													allresult.push(finalarr)
												}else{
													continue
												}
											}
										}else{
											continue
										}
									}
								}else{
									continue
								}
							}
						}else{
							continue
						}
					}
					/*for(var i2=0;i2<hero2pool.length;i2++){
						var v2=hero2pool[i2]
						if (v!=v2){
							firstarr.push(hero2pool[i2].hero_icon)
							for(var i3=0;i3<hero3pool.length;i3++){
								var v3=hero3pool[i3]
								if  (v!=v2&&v!=v3&&v2!=v3){
									firstarr.push(hero3pool[i3].hero_icon)
									for (var i4=0;i4<hero4pool.length;i4++){
										var v4=hero4pool[i4]
										if  (v!=v2&&v!=v3&&v!=v4&&v2!=v3&&v2!=v4&&v3!=v4){
											firstarr.push(hero4pool[i4].hero_icon)
											for (var i5=0;i5<hero5pool.length;i5++){
												var v5=hero5pool[i5]
												if  (v!=v2&&v!=v3&&v!=v4&&v!=v5&&v2!=v3&&v2!=v4&&v2!=v5&&v3!=v4&&v3!=v5&&v4!=v5){
													firstarr.push(hero5pool[i5].hero_icon)
													allresult.push(firstarr)
												}else{
													continue
												}
											}
										}else{
											continue
										}
									}
								}else{
									continue
								}
								
							}
						}else{
							continue
						}
						
						
					}*/
				}
				console.log(allresult)
				return allresult
			},
			
		},

		
	}
</script>
<style>
	
</style>