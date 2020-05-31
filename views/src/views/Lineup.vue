<template>
	<el-main style="padding:0px;">
      <div class="wrapper">
         <div style="width:1050px;margin:0 auto;padding:0 0">
            <div id="imitate_target" style="height:70px;">
               <el-button size="small">中路</el-button><el-button size="small">边路</el-button>
            </div>
            <div style="float:left;width:1050px;" >
               <div id="strengthhero" style="margin-top:50px;margin:0 0;padding:0 0">
             <ul class="herobox" >
               <li><img src="../static/img/hero/eldertitan.jpg" width="60"  title="上古巨神" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/undying.jpg" width="60" title="尸王"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/abaddon.jpg" width="60" title="亚巴顿" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/timbersaw.jpg" width="60" title="伐木机" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/omniknight.jpg" width="60" title="全能骑士" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/beastmaster.jpg" width="60" title="兽王" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/legioncommander.jpg" width="60" title="军团指挥官" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/wraithking.jpg" width="60" title="冥魂大帝" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/phoenix.jpg" width="60" title="凤凰" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/centaurwarrunner.jpg" width="60" title="半人马战士" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/clockwerk.jpg" width="60" title="发条技师" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/lifestealer.jpg" width="60" title="噬魂鬼" @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/huskar.jpg" width="60" title="哈斯卡" @click="cloneImg($event)"/></li>
               <li><img src="../static/img/hero/earthspirit.jpg" width="60" title="大地之灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/underlord.jpg" width="60" title="孽主"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/tiny.jpg" width="60" title="小小"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/tusk.jpg" width="60" title="巨牙海民"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/pudge.jpg" width="60" title="屠夫"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/earthshaker.jpg" width="60" title="撼地者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/axe.jpg" width="60" title="斧王"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/slardar.jpg" width="60" title="斯拉达"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/sven.jpg" width="60" title="斯温"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/kunkka.jpg" width="60" title="昆卡"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/nightstalker.jpg" width="60" title="暗夜魔王"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/doom.jpg" width="60" title="末日使者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/treantprotector.jpg" width="60" title="树精卫士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/sandking.jpg" width="60" title="沙王"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/chaosknight.jpg" width="60" title="混沌骑士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/tidehunter.jpg" width="60" title="潮汐猎人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/alchemist.jpg" width="60" title="炼金术士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/lycan.jpg" width="60" title="狼人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero//mars.jpg" width="60" title="马尔斯"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/snapfire.jpg" width="60" title="电炎绝手"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/io.jpg" width="60" title="小精灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/spiritbreaker.jpg" width="60" title="裂魂人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/brewmaster.jpg" width="60" title="酒仙"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/bristleback.jpg" width="60" title="钢背兽"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/magnus.jpg" width="60" title="马格纳斯"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/dragonknight.jpg" width="60" title="龙骑士"  @click="cloneImg($event)" /></li>
             </ul>
               </div>
               <div id="agilityhero">
             <ul class="herobox" >
               <li><img src="../static/img/hero/juggernaut.jpg" width="60"  title="主宰"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/clinkz.jpg" width="60" title="克林克兹"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/viper.jpg" width="60" title="冥界亚龙"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/venomancer.jpg" width="60" title="剧毒术士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/riki.jpg" width="60" title="力丸"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/drowranger.jpg" width="60" title="卓尔游侠"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/morphling.jpg" width="60" title="水人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/nyx-assassin.jpg" width="60" title="司夜刺客"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/templarassassin.jpg" width="60" title="圣堂刺客"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/vengeful-spirit.jpg" width="60" title="复仇之魂"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/arcwarden.jpg" width="60" title="天穹守望者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/nagasiren.jpg" width="60" title="娜迦海妖"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/trollwarlord.jpg" width="60" title="巨魔战将"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/phantomassassin.jpg" width="60" title="幻影刺客"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/phantomlancer.jpg" width="60" title="幻影长矛手"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/spectre.jpg" width="60" title="幽鬼"  @click="cloneImg($event)" /></li>
               <!-- <li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="npc_dota_hero_nevermore"  @click="cloneImg($event)" /></li> -->
               <li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="影魔"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/lone-druid.jpg" width="60" title="德鲁伊"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/terrorblade.jpg" width="60" title="恐怖利刃"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/antimage.jpg" width="60" title="敌法师"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/slark.jpg" width="60" title="斯拉克"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/emberspirit.jpg" width="60" title="灰烬之灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/ursa.jpg" width="60" title="熊战士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/sniper.jpg" width="60" title="狙击手"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/gyrocopter.jpg" width="60" title="矮人直升机"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/pangolier.jpg" width="60" title="石鳞剑士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/mirana.jpg" width="60" title="米拉娜"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/meepo.jpg" width="60" title="米波"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/weaver.jpg" width="60" title="编织者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/medusa.jpg" width="60" title="美杜莎"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/broodmother.jpg" width="60" title="育母蜘蛛"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero//gyrocopter.jpg" width="60" title="虚空假面"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/bloodseeker.jpg" width="60" title="血魔"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/bountyhunter.jpg" width="60" title="赏金猎人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/luna.jpg" width="60" title="露娜"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/monkeyking.jpg" width="60" title="大圣"  @click="cloneImg($event)" /></li>
             </ul>
               </div>
               <div id="agilityhero">
             <ul class="herobox" >
               <li><img src="../static/img/hero/tinker.jpg" width="60"  title="修补匠"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/naturesprophet.jpg" width="60" title="先知"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/keeperofthelight.jpg" width="60" title="光法"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/skywrathmage.jpg" width="60" title="天怒"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/grimstroke.jpg" width="60" title="天涯墨客"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/zeus.jpg" width="60" title="宙斯"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/winterwyvern.jpg" width="60" title="寒冬飞龙"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/techies.jpg" width="60" title="炸弹人"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/witchdoctor.jpg" width="60" title="巫医"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/lich.jpg" width="60" title="巫妖"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/puck.jpg" width="60" title="帕克"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/pugna.jpg" width="60" title="帕格纳"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/disruptor.jpg" width="60" title="干扰者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/dazzle.jpg" width="60" title="戴泽"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/leshrac.jpg" width="60" title="拉席克"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/rubick.jpg" width="60" title="拉比克"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/shadowdemon.jpg" width="60" title="暗影恶魔"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/shadowshaman.jpg" width="60" title="暗影萨满"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/warlock.jpg" width="60" title="术士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/jakiro.jpg" width="60" title="杰奇洛"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/deathprophet.jpg" width="60" title="死亡先知"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/outworlddevourer.jpg" width="60" title="黑鸟"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/crystalmaiden.jpg" width="60" title="冰女"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/silencer.jpg" width="60" title="沉默术士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/queenofpain.jpg" width="60" title="痛苦女王"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/necrophos.jpg" width="60" title="瘟疫法师"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/invoker.jpg" width="60" title="祈求者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/oracle.jpg" width="60" title="神谕者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/bane.jpg" width="60" title="霍乱之源"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/visage.jpg" width="60" title="维萨吉"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/lina.jpg" width="60" title="莉娜"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero//lion.jpg" width="60" title="莱恩"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/void-spirit.jpg" width="60" title="虚空之灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/batrider.jpg" width="60" title="蝙蝠骑士"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/enigma.jpg" width="60" title="谜团"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/ancient.jpg" width="60" title="远古冰魄"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/dark-willow.jpg" width="60" title="邪影芳灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/chen.jpg" width="60" title="陈"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/stormspirit.jpg" width="60" title="风暴之灵"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/windranger.jpg" width="60" title="风行者"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/ogremagi.jpg" width="60" title="蓝胖"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/enchantress.jpg" width="60" title="魅惑魔女"  @click="cloneImg($event)" /></li>
               <li><img src="../static/img/hero/dark-seer.jpg" width="60" title="黑暗贤者"  @click="cloneImg($event)" /></li>
             </ul>
               </div>
               <div>
                  <h2>对线情况</h2>
                  <div id="against_lane">
                     <div id="toplane">  
                        <div style="float:left;width:45px;height:45px;line-height:45px">边路</div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div style="width:45px;height:45px;float:left;line-height:45px" >vs</div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                     </div> 
                     <div id="midlane" >  
                        <div style="float:left;width:45px;height:45px;line-height:45px">中路</div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div style="width:45px;height:45px;float:left;line-height:45px" >vs</div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                        <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                     </div>
                    <!--  <div id="bottomlane" >  
                      <div style="float:left;width:45px;height:45px;line-height:45px">下路</div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                       <div style="width:45px;height:45px;float:left;line-height:45px" >vs</div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                       <div class="select_hero" @click="addhover($event)" @contextmenu.prevent="delImg($event)"></div>
                    </div> -->
                  </div>
               </div>
               <div>
                  <el-button size="small" @click="getAgainstData()">对线分析</el-button>
               </div>
            </div>
      </div>
      </div>
      <!-- <div v-if="tenMinData!=null">
         <div v-if="tenMinData.mid_hero_data!=null" >
            <p><h2>中路对线数据</h2></p>  
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
                  <td>{{item.match_id}}</td>
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
                                 <tr><td>{{index+1}}</td></tr>
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
                                 <tr><td>{{value1}}</td></tr>
                              </table>
                           </td>
                        </tr>
                     </table>
                  </td>
                  <td>11111</td>
                  <td>
                     <table class="details_table">
                        <tr>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_1.hero_name}}</td>
                           <td>{{tenMinData.mid_hero_data[0].mid_hero_2.hero_name}}</td>
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_3">{{tenMinData.mid_hero_data[0].mid_hero_3.hero_name}}</td>  
                           <td v-if="tenMinData.mid_hero_data[0].mid_hero_4">{{tenMinData.mid_hero_data[0].mid_hero_4.hero_name}}</td>
                        </tr>
                        <tr>
                           <td v-for="consumable_value in item.mid_hero_1.consumable">
                              {{consumable_value.item_icon}}*{{consumable_value.count}}
                           </td>
                        </tr>
                     </table>   
                  </td>
               </tr>
            </table> 
         </div>
      </div> -->
	</el-main>
</template>
<script>
import {mapState} from 'vuex'
  export default {
    data() {
      return {
        radio: '1',
        value: '',
        value1: '',
        bpOrder:1,
        ownTeam:null,
        againstTeam:null,
      };
    },
    props:['version'],
    computed:{
/*        ...mapState(['allTeamInfo']),//获取所有的战队名称
        ...mapState(['teamHeroPool']),
        ...mapState(['tenMinData'])*/
    },
    created(){
      //this.$store.dispatch('getAllTeamInfo')
     // console.log(this.version)
    },
   /* beforeRouteEnter(to,from,next){
      console.log(to)
      if (to.path=="/lineupdets"){
         next(vm=>{
            console.log(vm.page)
            vm.test()
         })
      }else{
         next()
      }
    }*/
    methods: {
      test(){
         consloe.log("lineuptest")
      },
    	bpClick(e){
         var target=e.target
         //var databpid=target.dataset.bpId
         var liele=document.getElementById("team_bp").getElementsByTagName("li")
         var order=this.bpOrder
         var node=target.cloneNode()   
         for (var i=1;i<=22;i++){
            var i2=i-1
            if (liele[i2].dataset.bpId==order){
               if (liele[i2].dataset.bpId==1||liele[i2].dataset.bpId==2||liele[i2].dataset.bpId==3||liele[i2].dataset.bpId==4||liele[i2].dataset.bpId==5|liele[i2].dataset.bpId==6||liele[i2].dataset.bpId==11||liele[i2].dataset.bpId==12||liele[i2].dataset.bpId==13||liele[i2].dataset.bpId==14||liele[i2].dataset.bpId==19||liele[i2].dataset.bpId==20){
                  node.style="filter: grayscale(1);" 
               }
               liele[i2].insertBefore(node,liele[i2].childNodes[1])
            }
         }
         this.bpOrder=this.bpOrder+1
    	},
       /* getHeroPool(teamtagName){
            var param={version:this.version,teamtagname:teamtagName}
            this.$store.dispatch('getTeamHeroPool',param)
        },*/
        addhover(e){
          //console.log(e)
         if (e.button==0){
            if (e.target.tagName=="IMG"){
               console.log(e.target.parentNode)
               e.target.parentNode.className=e.target.parentNode.className+" shadow_hover"
            }else{
               var parent=e.target.parentNode.parentNode
               var divele=parent.getElementsByClassName("select_hero")
               for (var i=0;i<divele.length;i++){
                  divele[i].className="select_hero"
               }
               e.target.className+=" shadow_hover"
            }
         }   
        },
        cloneImg(e){
         var imgele=e.target
         var cloneimg=imgele.cloneNode()
         var targetdiv=document.getElementById("against_lane").getElementsByClassName("select_hero")
          for (var i=0;i<targetdiv.length;i++){
            if (targetdiv[i].className=="select_hero shadow_hover"){   
               if (targetdiv[i].getElementsByTagName("img").length!=0){
                  //console.log(targetdiv[i].getElementsByTagName("img"))
                  targetdiv[i].removeChild(targetdiv[i].getElementsByTagName("img")[0])
               }
               targetdiv[i].appendChild(cloneimg)
               targetdiv[i].className="select_hero"
            }
          }
        },
        delImg(e){
          document.oncontextmenu = function(e){
               e.preventDefault();
           };
         if (e.target.tagName=="IMG"&&e.button ==2){
            e.target.parentNode.removeChild(e.target.parentNode.getElementsByTagName("img")[0])
         }
        },
        getAgainstData(){
         var toplane=""
         var midlane=""
         var midlane2=""
         var bottomlane=""
         var toplanesel1=""
         var toplanesel2=""
         var bottomlanesel1=""
         var bottomlanesel2=""
         var midfirstname=""
         var midsecondname=""
         var toplaneselhero=""
         var midlaneselhero=""
         var bottomlaneselhero=""
         var topselectdiv=document.getElementById("toplane").getElementsByClassName("select_hero")
         var midselectdiv=document.getElementById("midlane").getElementsByClassName("select_hero")
         var bottomselectdiv=document.getElementById("bottomlane").getElementsByClassName("select_hero")

         for (var i=0;i<topselectdiv.length;i++){
            if (i<=2){
               if (topselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=topselectdiv[i].getElementsByTagName("img")[0].title
                  toplanesel1+=heroname
                  toplaneselhero+=heroname+" "
               }
            }else{
               if (topselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=topselectdiv[i].getElementsByTagName("img")[0].title
                  toplanesel2+=heroname
                  toplaneselhero+=heroname+" "
               }
            }  
         }

         for (var i=0;i<midselectdiv.length;i++){
            if (i<=1){
               if (midselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=midselectdiv[i].getElementsByTagName("img")[0].title
                  midfirstname+=heroname
                  midlaneselhero+=heroname+" "
               }
            }else{
               if (midselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=midselectdiv[i].getElementsByTagName("img")[0].title
                  midsecondname+=heroname
                  midlaneselhero+=heroname+" "
               }
            }  
         }

         for (var i=0;i<bottomselectdiv.length;i++){
            if (i<=2){
               if (bottomselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=bottomselectdiv[i].getElementsByTagName("img")[0].title
                  bottomlanesel1+=heroname
                  bottomlaneselhero+=heroname+" "
               }
            }else{
               if (bottomselectdiv[i].getElementsByTagName("img").length>0){
                  var heroname=bottomselectdiv[i].getElementsByTagName("img")[0].title
                  bottomlanesel2+=heroname
                  bottomlaneselhero+=heroname+" "
               }
            }  
         }
         midlane=midfirstname+"vs"+midsecondname
         midlane2=midsecondname+"vs"+midfirstname
         toplane=toplanesel1+"vs"+toplanesel2
         bottomlane=bottomlanesel1+"vs"+bottomlanesel2

         toplaneselhero=toplaneselhero.replace(/(^\s*)|(\s*$)/g,'')
         midlaneselhero=midlaneselhero.replace(/(^\s*)|(\s*$)/g,'')
         bottomlaneselhero=bottomlaneselhero.replace(/(^\s*)|(\s*$)/g,'')

         var request={
            toplane:toplane,
            midlane:midlane,
            midlane2:midlane2,
            bottomlane:bottomlane,

            toplaneselhero:toplaneselhero,
            midlaneselhero:midlaneselhero,
            bottomlaneselhero:bottomlaneselhero,

            sidelane:"",
            softlane:"",
            hardlane:"",
            version:this.version
         }
         console.log(request)
         this.$store.dispatch('getLineUpRequest',request)
         this.$router.push('/lineupdets')
         //this.$store.dispatch('getTenMinData',request)
        },

    }
  }
</script>
<style>


.wrapper{
   width:1400px;margin:0 auto;
}
li{
  list-style:none;
}
ul.herobox{
	width:1050px;
   margin:0 0;
   padding:0 0;
	height:120px
}

ul.herobox li{
	
	width:65px;
	float:left
}
ul.hero_pool{
  width:100%;
}
ul.hero_pool li{
margin-bottom:7px;
}
.team_bp{
float:left;
}
.shadow_hover{
   box-shadow:0px 0px 10px #888888;
}
#left_team{
   width:170px;
}
#right_team{
   width:170px;
}
#left_team ul li div,#right_team ul li div{ 
   font-size:0.9rem
}
.select_hero{
   width:60px;height:34px;border:1px solid #676767;float:left
}
.ban_pick_order{
   background:#000000;
}
#left_team_ban_pick li,#right_team_ban_pick li{
   width:60px;height:50px; border:1px solid #676767;margin-top:15px
}
#toplane,#bottomlane,#midlane{
   width:480px;height:65px;margin:0 auto;
}


</style>
