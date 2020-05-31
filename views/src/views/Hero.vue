<template>
	<el-main style="padding:0px;">
		<div>
			<el-select v-model="value" placeholder="请选择" @change="test">
			    <el-option
			      v-for="item in options"
			      :key="item.value"
			      :label="item.label"
			      :value="item.value">
			    </el-option>
			 </el-select>
		</div>
		<div id="strengthhero" class="heroeswrapper">
			<ul class="herobox" >
				<li><img src="../static/img/hero/eldertitan.jpg" width="60"  title="上古巨神" @click="getData"/></li>
				<li><img src="../static/img/hero/undying.jpg" width="60" title="尸王" @click="getData"/></li>
				<li><img src="../static/img/hero/abaddon.jpg" width="60" title="亚巴顿" @click="getData"/></li>
				<li><img src="../static/img/hero/timbersaw.jpg" width="60" title="伐木机" @click="getData"/></li>
				<li><img src="../static/img/hero/omniknight.jpg" width="60" title="全能骑士" @click="getData"/></li>
				<li><img src="../static/img/hero/beastmaster.jpg" width="60" title="兽王" @click="getData"/></li>
				<li><img src="../static/img/hero/legioncommander.jpg" width="60" title="军团指挥官" @click="getData"/></li>
				<li><img src="../static/img/hero/wraithking.jpg" width="60" title="冥魂大帝" @click="getData"/></li>
				<li><img src="../static/img/hero/phoenix.jpg" width="60" title="凤凰" @click="getData"/></li>
				<li><img src="../static/img/hero/centaurwarrunner.jpg" width="60" title="半人马战士" @click="getData"/></li>
				<li><img src="../static/img/hero/clockwerk.jpg" width="60" title="发条技师" @click="getData"/></li>
				<li><img src="../static/img/hero/lifestealer.jpg" width="60" title="噬魂鬼" @click="getData"/></li>
				<li><img src="../static/img/hero/huskar.jpg" width="60" title="哈斯卡" @click="getData"/></li>
				<li><img src="../static/img/hero/earthspirit.jpg" width="60" title="大地之灵" @click="getData"/></li>
				<li><img src="../static/img/hero/underlord.jpg" width="60" title="孽主" @click="getData"/></li>
				<li><img src="../static/img/hero/tiny.jpg" width="60" title="小小" @click="getData"/></li>
				<li><img src="../static/img/hero/tusk.jpg" width="60" title="巨牙海民" @click="getData"/></li>
				<li><img src="../static/img/hero/pudge.jpg" width="60" title="屠夫" @click="getData"/></li>
				<li><img src="../static/img/hero/earthshaker.jpg" width="60" title="撼地者" @click="getData"/></li>
				<li><img src="../static/img/hero/axe.jpg" width="60" title="斧王" @click="getData"/></li>
				<li><img src="../static/img/hero/slardar.jpg" width="60" title="斯拉达" @click="getData"/></li>
				<li><img src="../static/img/hero/sven.jpg" width="60" title="斯温" @click="getData"/></li>
				<li><img src="../static/img/hero/kunkka.jpg" width="60" title="昆卡" @click="getData"/></li>
				<li><img src="../static/img/hero/nightstalker.jpg" width="60" title="暗夜魔王" @click="getData"/></li>
				<li><img src="../static/img/hero/doom.jpg" width="60" title="末日使者" @click="getData"/></li>
				<li><img src="../static/img/hero/treantprotector.jpg" width="60" title="树精卫士" @click="getData"/></li>
				<li><img src="../static/img/hero/sandking.jpg" width="60" title="沙王" @click="getData"/></li>
				<li><img src="../static/img/hero/chaosknight.jpg" width="60" title="混沌骑士" @click="getData"/></li>
				<li><img src="../static/img/hero/tidehunter.jpg" width="60" title="潮汐猎人" @click="getData"/></li>
				<li><img src="../static/img/hero/alchemist.jpg" width="60" title="炼金术士" @click="getData"/></li>
				<li><img src="../static/img/hero/lycan.jpg" width="60" title="狼人" @click="getData"/></li>
				<li><img src="../static/img/hero//mars.jpg" width="60" title="马尔斯" @click="getData"/></li>
				<li><img src="../static/img/hero/snapfire.jpg" width="60" title="电炎绝手" @click="getData"/></li>
				<li><img src="../static/img/hero/io.jpg" width="60" title="小精灵" @click="getData"/></li>
				<li><img src="../static/img/hero/spiritbreaker.jpg" width="60" title="裂魂人" @click="getData"/></li>
				<li><img src="../static/img/hero/brewmaster.jpg" width="60" title="酒仙" @click="getData"/></li>
				<li><img src="../static/img/hero/bristleback.jpg" width="60" title="钢背兽" @click="getData"/></li>
				<li><img src="../static/img/hero/magnus.jpg" width="60" title="马格纳斯" @click="getData"/></li>
				<li><img src="../static/img/hero/dragonknight.jpg" width="60" title="龙骑士" @click="getData"/></li>
			</ul>
		</div>
		<div id="agilityhero" class="heroeswrapper">
			<ul class="herobox" >
				<li><img src="../static/img/hero/juggernaut.jpg" width="60"  title="主宰" @click="getData"/></li>
				<li><img src="../static/img/hero/clinkz.jpg" width="60" title="克林克兹" @click="getData"/></li>
				<li><img src="../static/img/hero/viper.jpg" width="60" title="冥界亚龙" @click="getData"/></li>
				<li><img src="../static/img/hero/venomancer.jpg" width="60" title="剧毒术士" @click="getData"/></li>
				<li><img src="../static/img/hero/riki.jpg" width="60" title="力丸" @click="getData"/></li>
				<li><img src="../static/img/hero/drowranger.jpg" width="60" title="卓尔游侠" @click="getData"/></li>
				<li><img src="../static/img/hero/morphling.jpg" width="60" title="水人" @click="getData"/></li>
				<li><img src="../static/img/hero/nyx-assassin.jpg" width="60" title="司夜刺客" @click="getData"/></li>
				<li><img src="../static/img/hero/templarassassin.jpg" width="60" title="圣堂刺客" @click="getData"/></li>
				<li><img src="../static/img/hero/vengeful-spirit.jpg" width="60" title="复仇之魂" @click="getData"/></li>
				<li><img src="../static/img/hero/arcwarden.jpg" width="60" title="天穹守望者" @click="getData"/></li>
				<li><img src="../static/img/hero/nagasiren.jpg" width="60" title="娜迦海妖" @click="getData"/></li>
				<li><img src="../static/img/hero/trollwarlord.jpg" width="60" title="巨魔战将" @click="getData"/></li>
				<li><img src="../static/img/hero/phantomassassin.jpg" width="60" title="幻影刺客" @click="getData"/></li>
				<li><img src="../static/img/hero/phantomlancer.jpg" width="60" title="幻影长矛手" @click="getData"/></li>
				<li><img src="../static/img/hero/spectre.jpg" width="60" title="幽鬼" @click="getData"/></li>
				<!-- <li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="npc_dota_hero_nevermore" @click="getData"/></li> -->
				<li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="影魔" @click="getData"/></li>
				<li><img src="../static/img/hero/lone-druid.jpg" width="60" title="德鲁伊" @click="getData"/></li>
				<li><img src="../static/img/hero/terrorblade.jpg" width="60" title="恐怖利刃" @click="getData"/></li>
				<li><img src="../static/img/hero/antimage.jpg" width="60" title="敌法师" @click="getData"/></li>
				<li><img src="../static/img/hero/slark.jpg" width="60" title="斯拉克" @click="getData"/></li>
				<li><img src="../static/img/hero/emberspirit.jpg" width="60" title="灰烬之灵" @click="getData"/></li>
				<li><img src="../static/img/hero/ursa.jpg" width="60" title="熊战士" @click="getData"/></li>
				<li><img src="../static/img/hero/sniper.jpg" width="60" title="狙击手" @click="getData"/></li>
				<li><img src="../static/img/hero/gyrocopter.jpg" width="60" title="矮人直升机" @click="getData"/></li>
				<li><img src="../static/img/hero/pangolier.jpg" width="60" title="石鳞剑士" @click="getData"/></li>
				<li><img src="../static/img/hero/mirana.jpg" width="60" title="米拉娜" @click="getData"/></li>
				<li><img src="../static/img/hero/meepo.jpg" width="60" title="米波" @click="getData"/></li>
				<li><img src="../static/img/hero/weaver.jpg" width="60" title="编织者" @click="getData"/></li>
				<li><img src="../static/img/hero/medusa.jpg" width="60" title="美杜莎" @click="getData"/></li>
				<li><img src="../static/img/hero/broodmother.jpg" width="60" title="育母蜘蛛" @click="getData"/></li>
				<li><img src="../static/img/hero//gyrocopter.jpg" width="60" title="虚空假面" @click="getData"/></li>
				<li><img src="../static/img/hero/bloodseeker.jpg" width="60" title="血魔" @click="getData"/></li>
				<li><img src="../static/img/hero/bountyhunter.jpg" width="60" title="赏金猎人" @click="getData"/></li>
				<li><img src="../static/img/hero/luna.jpg" width="60" title="露娜" @click="getData"/></li>
				<li><img src="../static/img/hero/monkeyking.jpg" width="60" title="大圣" @click="getData"/></li>
			</ul>
		</div>
		<div id="agilityhero" class="heroeswrapper">
			<ul class="herobox" >
				<li><img src="../static/img/hero/tinker.jpg" width="60"  title="修补匠" @click="getData"/></li>
				<li><img src="../static/img/hero/naturesprophet.jpg" width="60" title="先知" @click="getData"/></li>
				<li><img src="../static/img/hero/keeperofthelight.jpg" width="60" title="光法" @click="getData"/></li>
				<li><img src="../static/img/hero/skywrathmage.jpg" width="60" title="天怒" @click="getData"/></li>
				<li><img src="../static/img/hero/grimstroke.jpg" width="60" title="天涯墨客" @click="getData"/></li>
				<li><img src="../static/img/hero/zeus.jpg" width="60" title="宙斯" @click="getData"/></li>
				<li><img src="../static/img/hero/winterwyvern.jpg" width="60" title="寒冬飞龙" @click="getData"/></li>
				<li><img src="../static/img/hero/techies.jpg" width="60" title="炸弹人" @click="getData"/></li>
				<li><img src="../static/img/hero/witchdoctor.jpg" width="60" title="巫医" @click="getData"/></li>
				<li><img src="../static/img/hero/lich.jpg" width="60" title="巫妖" @click="getData"/></li>
				<li><img src="../static/img/hero/puck.jpg" width="60" title="帕克" @click="getData"/></li>
				<li><img src="../static/img/hero/pugna.jpg" width="60" title="帕格纳" @click="getData"/></li>
				<li><img src="../static/img/hero/disruptor.jpg" width="60" title="干扰者" @click="getData"/></li>
				<li><img src="../static/img/hero/dazzle.jpg" width="60" title="戴泽" @click="getData"/></li>
				<li><img src="../static/img/hero/leshrac.jpg" width="60" title="拉席克" @click="getData"/></li>
				<li><img src="../static/img/hero/rubick.jpg" width="60" title="拉比克" @click="getData"/></li>
				<li><img src="../static/img/hero/shadowdemon.jpg" width="60" title="暗影恶魔" @click="getData"/></li>
				<li><img src="../static/img/hero/shadowshaman.jpg" width="60" title="暗影萨满" @click="getData"/></li>
				<li><img src="../static/img/hero/warlock.jpg" width="60" title="术士" @click="getData"/></li>
				<li><img src="../static/img/hero/jakiro.jpg" width="60" title="杰奇洛" @click="getData"/></li>
				<li><img src="../static/img/hero/deathprophet.jpg" width="60" title="死亡先知" @click="getData"/></li>
				<li><img src="../static/img/hero/outworlddevourer.jpg" width="60" title="黑鸟" @click="getData"/></li>
				<li><img src="../static/img/hero/crystalmaiden.jpg" width="60" title="冰女" @click="getData"/></li>
				<li><img src="../static/img/hero/silencer.jpg" width="60" title="沉默术士" @click="getData"/></li>
				<li><img src="../static/img/hero/queenofpain.jpg" width="60" title="痛苦女王" @click="getData"/></li>
				<li><img src="../static/img/hero/necrophos.jpg" width="60" title="瘟疫法师" @click="getData"/></li>
				<li><img src="../static/img/hero/invoker.jpg" width="60" title="祈求者" @click="getData"/></li>
				<li><img src="../static/img/hero/oracle.jpg" width="60" title="神谕者" @click="getData"/></li>
				<li><img src="../static/img/hero/bane.jpg" width="60" title="霍乱之源" @click="getData"/></li>
				<li><img src="../static/img/hero/visage.jpg" width="60" title="维萨吉" @click="getData"/></li>
				<li><img src="../static/img/hero/lina.jpg" width="60" title="莉娜" @click="getData"/></li>
				<li><img src="../static/img/hero//lion.jpg" width="60" title="莱恩" @click="getData"/></li>
				<li><img src="../static/img/hero/void-spirit.jpg" width="60" title="虚空之灵" @click="getData"/></li>
				<li><img src="../static/img/hero/batrider.jpg" width="60" title="蝙蝠骑士" @click="getData"/></li>
				<li><img src="../static/img/hero/enigma.jpg" width="60" title="谜团" @click="getData"/></li>
				<li><img src="../static/img/hero/ancient.jpg" width="60" title="远古冰魄" @click="getData"/></li>
				<li><img src="../static/img/hero/dark-willow.jpg" width="60" title="邪影芳灵" @click="getData"/></li>
				<li><img src="../static/img/hero/chen.jpg" width="60" title="陈" @click="getData"/></li>
				<li><img src="../static/img/hero/stormspirit.jpg" width="60" title="风暴之灵" @click="getData"/></li>
				<li><img src="../static/img/hero/windranger.jpg" width="60" title="风行者" @click="getData"/></li>
				<li><img src="../static/img/hero/ogremagi.jpg" width="60" title="蓝胖" @click="getData"/></li>
				<li><img src="../static/img/hero/enchantress.jpg" width="60" title="魅惑魔女" @click="getData"/></li>
				<li><img src="../static/img/hero/dark-seer.jpg" width="60" title="黑暗贤者" @click="getData"/></li>
			</ul>
		</div>
	</el-main>
</template>
<script>
	
	export default{
		data(){
			return{
				version:"",
				options: [
				{
		          value: '',
		          label: '无'
		        },
				{
		          value: '一号位',
		          label: '一号位'
		        }, {
		          value: '二号位',
		          label: '二号位'
		        }, {
		          value: '三号位',
		          label: '三号位'
		        }, {
		          value: '四号位',
		          label: '四号位'
		        }, {
		          value: '五号位',
		          label: '五号位'
		        }],
		        value: '',
			}
		},
		computed:{
			getVersion:{
				get: function() {
					if (this.$store.state.version!=null){
						this.version=this.$store.state.version
					}
					return this.$store.state.version
				},
				set:function(value){
					this.version = value
					this.$store.dispatch('getVersion',this.version)

				}
			}
		},
		methods:{
			getData(e){
				var heroname=e.target.title
				this.$router.push({name:'herodata',query:{heroname:heroname,version:this.$store.state.version,position:this.value}})
			},
			test(value){
				console.log(value)
			}

		}
	}
</script>
<style>
#version-input{
	margin-top:15px;
background-color: #FFF;
    background-image: none;
    border-radius: 4px;
    border: 1px solid #DCDFE6;
    box-sizing: border-box;
    color: #606266;
    display: inline-block;
    font-size: inherit;
    height: 40px;
    line-height: 40px;
    outline: 0;
    padding: 0 15px;
    font-size:14px
}
.heroeswrapper{
	width:1200px;
	margin:10px auto 0 auto;
}
</style>