<template>
<el-container>
	<!-- <el-aside width="200px">
		<img :src="require('../'+heroData.HeroIcon)" width="120" /> 这样图片也可以出现，不过浏览器会报错 
		<img :src="getHeroData.HeroIcon" width="120" />
		<img :src="heroData.HeroIcon" /> 
	</el-aside> -->
	<el-main style="padding:0px;">
		<div class="wrapper" style="border:1px solid #ff0000">
			<div class="hero_introduce" style="border:1px solid #ff0000">
				<div class="hero_introduce_icon">
					<img :src="heroData.HeroIcon" class="hero_icon"/>
					<p class="hero_name">{{heroData.Name}}</p>
					<p v-if="heroData.AttributePrimary==='DOTA_ATTRIBUTE_AGILITY'" class="hero_attribute">敏捷英雄</p>
					<p v-else-if="heroData.AttributePrimary==='DOTA_ATTRIBUTE_STRENGTH'" class="hero_attribute">力量英雄</p>
					<p v-else class="hero_attribute">智力英雄</p>
				</div>
				<div class="attribute">
					<div class="attribut_str">
						<div class="attribute_str_icon"><img src="../static/img/hero_str.png" width="40" /></div>
						<div class="attribut_gain">{{heroData.AttributeBaseStrength}}+{{heroData.AttributeStrengthGain}}</div>
					</div>
					<div class="attribut_agi">
						<div ><img src="../static/img/hero_agi.png" width="40" /></div>
						<div class="attribut_gain">{{heroData.AttributeBaseAgility}}+{{heroData.AttributeAgilityGain}}</div>
					</div>
					<div class="attribut_int">
						<div ><img src="../static/img/hero_int.png" width="40" /></div>
						<div class="attribut_gain">{{heroData.AttributeBaseIntelligence}}+{{heroData.AttributeIntelligenceGain}}</div>
					</div>
				</div>
				<div class="other_info" style="border:1px solid #ff0000">
					<p v-if="heroData.AttributePrimary==='DOTA_ATTRIBUTE_AGILITY'">攻击力：{{heroData.AttackDamageMax+heroData.AttributeBaseAgility}}/{{heroData.AttackDamageMin+heroData.AttributeBaseAgility}}</p>
					<p v-else-if="heroData.AttributePrimary==='DOTA_ATTRIBUTE_STRENGTH'">攻击力：{{heroData.AttackDamageMax+heroData.AttributeBaseStrength}}/{{heroData.AttackDamageMin+heroData.AttributeBaseStrength}}</p>
					<p v-else>攻击力：{{heroData.AttackDamageMax+heroData.AttributeBaseIntelligence}}/{{heroData.AttackDamageMin+heroData.AttributeBaseIntelligence}}</p>
					<p>护甲：{{armor}}</p>
					<p>回血：{{healthRegen}}</p>
				</div>
				<div class="other_info" style="border:1px solid #ff0000">
					<p>移速：{{heroData.MovementSpeed}}</p>
					<p>射程：{{heroData.AttackRange}}</p>
					<p>1级时攻击间隔：{{attackrate}}</p>
				</div>

			</div>
			<div style="clear:both"></div>
			<div class="lvl_data" style="border:1px solid #ff0000;width:100%;">
			<table class="" style="border:1px solid #ff0000;width:100%">
				<caption >每级基本数据</caption>
				<tr>
					<td>等级</td>
					<td>1</td>
					<td>2</td>
					<td>3</td>
					<td>4</td>
					<td>5</td>
					<td>6</td>
					<td>7</td>
					<td>8</td>
					<td>9</td>
					<td>10</td>
					<td>11</td>
					<td>12</td>
					<td>13</td>
					<td>14</td>
					<td>15</td>
					<td>16</td>
					<td>17</td>
					<td>18</td>
					<td>19</td>
					<td>20</td>
					<td>21</td>
					<td>22</td>
					<td>23</td>
					<td>24</td>
					<td>25</td>
				</tr>
				<tr><td>生命值</td>
					<td>{{heroData.AttributeBaseStrength*20+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*1)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*2)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*3)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*4)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*5)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*6)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*7)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*8)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*9)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*10)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*11)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*12)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*13)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*14)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*15)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*16)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*17)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*18)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*19)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*20)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*21)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*22)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*23)*2+statusHealth}}</td>
					<td>{{(heroData.AttributeBaseStrength)*20+(heroData.AttributeStrengthGain*10*24)*2+statusHealth}}</td>
				</tr>
				<tr>
					<td>护甲</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+this.$store.state.heroData.AttributeBaseAgility*0.16}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*1)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*2)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*3)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*4)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*5)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*6)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*7)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*8)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*9)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*10)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*11)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*12)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*13)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*14)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*15)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*16)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*17)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*18)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*19)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*20)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*21)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*22)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*23)*16)/1000}}</td>
					<td>{{this.$store.state.heroData.ArmorPhysical+((this.$store.state.heroData.AttributeBaseAgility*10+heroData.AttributeAgilityGain*10*24)*16)/1000}}</td>
				</tr>
				<!-- <tr><td>攻击间隔</td></tr> -->
			</table>
		
		</div> 
		</div>

		<!-- <div style="background:#fff">{{heroData}}</div> -->
		<!-- <div style="background:#fff"><p>初始力量：{{heroData.AttributeBaseStrength}}+{{heroData.AttributeStrengthGain}}</p></div> -->
	</el-main>
</el-container>
</template>
<script>
import {mapState} from 'vuex'
export default{
	data(){
		return {
			name:null
		}
	},
	computed:{
		...mapState(['heroData']),
		...mapState(['statusHealth']),
		...mapState(['statusMana']),
		...mapState(['statusHealthRegen']),
		attackrate:function(){
			var rate
			if(this.$store.state.heroData.hasOwnProperty('BaseAttackSpeed')){
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_AGILITY'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseAgility+this.$store.state.heroData.BaseAttackSpeed)*0.01)
					rate=rate.toFixed(2)
				}
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_STRENGTH'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseStrength+this.$store.state.heroData.BaseAttackSpeed)*0.01)
					rate=rate.toFixed(2)
				}
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_INTELLECT'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseIntelligence+this.$store.state.heroData.BaseAttackSpeed)*0.01)
					rate=rate.toFixed(2)
				}
			}else{
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_AGILITY'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseAgility+100)*0.01)
					rate=rate.toFixed(2)
				}
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_STRENGTH'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseStrength+100)*0.01)
					rate=rate.toFixed(2)
				}
				if (this.$store.state.heroData.AttributePrimary==='DOTA_ATTRIBUTE_INTELLECT'){
					rate=this.$store.state.heroData.AttackRate/((this.$store.state.heroData.AttributeBaseIntelligence+100)*0.01)
					rate=rate.toFixed(2)
				}
			}
			return rate
		},
		armor:function(){
			var armor=this.$store.state.heroData.ArmorPhysical+this.$store.state.heroData.AttributeBaseAgility*0.16
			armor=armor.toFixed(2)
			return armor
		},
		healthRegen:function(){
			var healthRegen
			if (this.$store.state.heroData.hasOwnProperty('StatusHealthRegen')){
				healthRegen=this.$store.state.heroData.StatusHealthRegen+this.$store.state.statusHealthRegen+this.$store.state.heroData.AttributeBaseStrength*0.09
			}else{
				healthRegen=this.$store.state.statusHealthRegen+this.$store.state.heroData.AttributeBaseStrength*0.09
			}
			healthRegen=healthRegen.toFixed(2)
			return healthRegen
		}

	},
	mounted(){
		var version=this.$store.state.version
		var heroname=this.$route.query.heroname
		var param=JSON.stringify({"version":version,"heroname":heroname})
		this.$store.dispatch('getHeroData',param)	
	},
	created(){
		
	},
	methods:{
			
	}
}
</script>
<style>
.wrapper{
	width:1150px;
	margin:0 auto
}
.hero_introduce{
	height:75px;
	width:900px;
	margin:15px auto 0px auto;
}
.hero_icon{
	float:left;
	padding-right:10px;
}
.hero_introduce_icon{
	float:left;
	padding:0px;
	
}
.hero_name{
	text-align:left;
	font-size:16px;
	width:300px;
	font-weight:bold;
	margin:0px;
	padding:13px
}
.hero_attribute{
	text-align:left;
	font-size:15px;
	width:300px;
	margin:0px;
}
.hero_introduce{
	
}
.attribute{
	width:200px;
float:left;
font-size:14px;

}
.attribut_icon{

}
.attribut_str{
	float:left
}
.attribut_gain{
	padding-top:3px
}
.attribut_agi{
	float:left;
	margin-left:25px;
	
}
.attribut_int{

}
.other_info{
	width:180px;
	float:left
}
.other_info p{
	font-size:14px;
	margin:8px

}
</style>