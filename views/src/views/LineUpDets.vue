<template>
	<el-main style="padding:0px;">
		<div v-if="tenMinData!=null">
         <div v-if="tenMinData.mid_hero_data!=null" >
            <p><h2>中路对线数据</h2></p>
            <p v-if="tenMinData.mid_hero_data!=null">
               <el-button  @click="lastHitSort(tenMinData.mid_hero_data[0].mid_hero_1.hero_name)">按照{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}的补刀高低排序</el-button>
               <el-button @click="lastHitSort(tenMinData.mid_hero_data[0].mid_hero_2.hero_name)">按照{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}的补刀高低排序</el-button>
               <el-button @click="allGoldSort(tenMinData.mid_hero_data[0].mid_hero_1.hero_name)">按照{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}总经济高低排序</el-button>
               <el-button @click="allGoldSort(tenMinData.mid_hero_data[0].mid_hero_2.hero_name)">按照{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}总经济高低排序</el-button></p>  
            <table bgcolor="#334656" id="tenmin_table">
               <tr>
                  <td>比赛ID</td>
                  <td>5分钟补刀</td>
                  <td>10分钟补刀</td>
                  <td>5分钟总经济</td>
                  <td>补刀经济</td>
                  <td>助攻及击杀经济</td>
                  <td>10分钟总经济</td>
                  <td>补刀经济</td>
                  <td>助攻及击杀经济</td>
                  <td>10分钟前等级信息</td>
                  <td>10分钟前技能加点</td>
                  <td>10分钟前消耗品信息</td>
               </tr>
               <tr v-for="item in tenMinData.mid_hero_data" :key="item.match_id">
                  <td>
                     {{item.match_id}}<br/>
                     <span style="font-size:14px">{{tenMinData.mid_hero_data[0].mid_hero_1.player_name}} {{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</span><br/>
                     <span style="font-size:14px">{{tenMinData.mid_hero_data[0].mid_hero_2.player_name}} {{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</span>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{item.mid_hero_1.pre_5min_lasthit}}</td>
                           <td>{{item.mid_hero_2.pre_5min_lasthit}}</td>
                           <td v-if="item.mid_hero_3">{{item.mid_hero_3.pre_5min_lasthit}}</td>
                           <td v-if="item.mid_hero_4">{{item.mid_hero_4.pre_5min_lasthit}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{item.mid_hero_1.pre_10min_lasthit}}</td>
                           <td>{{item.mid_hero_2.pre_10min_lasthit}}</td>
                           <td v-if="item.mid_hero_3">{{item.mid_hero_3.pre_10min_lasthit}}</td>
                           <td v-if="item.mid_hero_4">{{item.mid_hero_4.pre_10min_lasthit}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_5min_gold.all_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_5min_gold.all_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.all_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.all_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_5min_gold.last_hit_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_5min_gold.last_hit_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.last_hit_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.last_hit_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_5min_gold.combat_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_5min_gold.combat_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.combat_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.combat_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_10min_gold.all_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_10min_gold.all_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.pre_10min_gold.all_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.pre_10min_gold.all_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_10min_gold.last_hit_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_10min_gold.last_hit_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.pre_10min_gold.last_hit_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.pre_10min_gold.last_hit_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td>{{Math.ceil(item.mid_hero_1.pre_10min_gold.combat_gold)}}</td>
                           <td>{{Math.ceil(item.mid_hero_2.pre_10min_gold.combat_gold)}}</td>
                           <td v-if="item.mid_hero_3">{{Math.ceil(item.mid_hero_3.pre_5min_gold.pre_10min_gold.combat_gold)}}</td>
                           <td v-if="item.mid_hero_4">{{Math.ceil(item.mid_hero_4.pre_5min_gold.pre_10min_gold.combat_gold)}}</td>
                        </tr>
                     </table>
                  </td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>
                              {{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}  
                           </td>
                           <td v-for="(value,index) in item.mid_hero_1.pre_10min_levelup_time">
                              <table>
                                 <tr><td>{{index+2}}</td></tr>
                                 <tr><td>{{value}}</td></tr>
                              </table>
                           </td>
                        </tr>
                        <tr>
                           <td>
                              {{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}  
                           </td>
                           <td v-for="(value1,index1) in item.mid_hero_2.pre_10min_levelup_time">
                              <table>
                              	<tr><td>{{index1+1}}</td></tr>
                                <tr><td>{{value1}}</td></tr>
                              </table>
                           </td>
                        </tr>
                     </table>
                  </td>
                  <td>
                  	<table class="details_table">
                  		<tr>
                  			<td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                  			<td v-for="(value,index) in item.mid_hero_1.ability_upgrades">
                  				<table>
                  					<tr>
                  						<td>{{index}}</td>
                  					</tr>
                  					<tr>
                  						<td>
                  							<img :src="value.Icon" width="30">
                  						</td>
                  					</tr>
                  				</table>
                  			</td>
                  		</tr>
                  		<tr>
                  			<td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                  			<td v-for="(value1,index1) in item.mid_hero_2.ability_upgrades">
                  				<table>
                  					<tr>
                  						<td>{{index1}}</td>
                  					</tr>
                  					<tr>
                  						<td>
                  							<img :src="value1.Icon" width="30">
                  						</td>
                  					</tr>
                  				</table>
                  			</td>
                  		</tr>	
                  	</table>
                  </td>
                  <td>
                     <table class="details_table">
                     	<tr>
                     		<td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                     		<td v-for="consumable_value in item.mid_hero_1.consumable">
                     			<img :src="consumable_value.item_icon" width="30">*{{consumable_value.count}}
                     		</td>
                     	</tr>
                     	<tr>
                     		<td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                     		<td v-for="consumable_value2 in item.mid_hero_2.consumable">
                     			<img :src="consumable_value2.item_icon" width="30">*{{consumable_value2.count}}
                     		</td>
                     	</tr>
                     	<tr v-if="tenMinData.mid_hero_data[0].mid_hero_3">
                     		<td >{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>
                     		<td v-for="consumable_value3 in item.mid_hero_3.consumable">
                     			<img :src="consumable_value3.item_icon" width="30">*{{consumable_value3.count}}
                     		</td>
                     	</tr>
                     	<tr v-if="tenMinData.mid_hero_data[0].mid_hero_4">
                     		<td >{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>
                     		<td v-for="consumable_value4 in item.mid_hero_4.consumable">
                     			<img :src="consumable_value4.item_icon" width="30">*{{consumable_value4.count}}
                     		</td>
                     	</tr>
                     </table>   
                  </td>
               </tr>
            </table> 
         </div>
      </div> 
	</el-main>
</template>
<script>
import {mapState} from 'vuex'
	export default{
		data(){
			return {

			}
		},
		created(){
			var params=this.$store.state.lineUpRequest
			this.$store.dispatch('getTenMinData',params)
		},
		computed:{
			 ...mapState(['tenMinData'])
		},
		methods:{
			test(){
				console.log(this.lineUpRequest)
			},
         lastHitSort(heroname){
            var tenMinData=this.tenMinData
            var len =  tenMinData.length
            var tmp = null
            if (tenMinData.mid_hero_data[0].mid_hero_1.heroname==heroname){
               for(var j=0; j<len-1; j++){      //循环9次
                  for(var i=0; i<len-1-j; i++){    //每次比较10-j-1次数
                     if(tenMinData.mid_hero_data[i].mid_hero_1.pre_10min_lasthit>tenMinData.mid_hero_data[i+1].mid_hero_1.pre_10min_lasthit){
                        tmp = tenMinData.mid_hero_data[i+1];
                        tenMinData.mid_hero_data[i+1] = tenMinData.mid_hero_data[i];
                        tenMinData.mid_hero_data[i] = tmp;
                     }
                  }
               }
            }else{
               for(var j=0; j<len-1; j++){      //循环9次
                  for(var i=0; i<len-1-j; i++){    //每次比较10-j-1次数
                     if(tenMinData.mid_hero_data[i].mid_hero_2.pre_10min_lasthit>tenMinData.mid_hero_data[i+1].mid_hero_2.pre_10min_lasthit){
                        tmp = tenMinData.mid_hero_data[i+1];
                        tenMinData.mid_hero_data[i+1] = tenMinData.mid_hero_data[i];
                        tenMinData.mid_hero_data[i] = tmp;
                     }
                  }
               }
            }
            this.$store.commit('receive_ten_min_data',{tenMinData})      
		   },
         allGoldSort(){
            var tenMinData=this.tenMinData
            var len =  tenMinData.length
            var tmp = null
            if (tenMinData.mid_hero_data[0].mid_hero_1.heroname==heroname){
               for(var j=0; j<len-1; j++){      //循环9次
                  for(var i=0; i<len-1-j; i++){    //每次比较10-j-1次数
                     if(tenMinData.mid_hero_data[i].mid_hero_1.pre_10min_gold>tenMinData.mid_hero_data[i+1].mid_hero_1.pre_10min_gold){
                        tmp = tenMinData.mid_hero_data[i+1];
                        tenMinData.mid_hero_data[i+1] = tenMinData.mid_hero_data[i];
                        tenMinData.mid_hero_data[i] = tmp;
                     }
                  }
               }
            }else{
               for(var j=0; j<len-1; j++){      //循环9次
                  for(var i=0; i<len-1-j; i++){    //每次比较10-j-1次数
                     if(tenMinData.mid_hero_data[i].mid_hero_2.pre_10min_gold>tenMinData.mid_hero_data[i+1].mid_hero_2.pre_10min_gold){
                        tmp = tenMinData.mid_hero_data[i+1];
                        tenMinData.mid_hero_data[i+1] = tenMinData.mid_hero_data[i];
                        tenMinData.mid_hero_data[i] = tmp;
                     }
                  }
               }
            }
            this.$store.commit('receive_ten_min_data',{tenMinData})  
         }
	   }
}
</script>
<style >


</style>
