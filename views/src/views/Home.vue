<template>
	<el-main style="padding:0px;">
      <div style="width:1400px;margin:0 auto;border:1px solid #ff0000">
         <div style="width:1050px;border:1px solid #ff0000;float:left;margin:0 0;padding:0 0">
            <div id="imitate_target" style="height:70px;">
               <div style="float:left;width:40%;border:1px solid #ff0000;margin-top:15px">
                  <el-select v-model="value1" filterable placeholder="己方战队" @change="getHeroPool1(value1)">
                     <el-option v-for="item in allTeamInfo" :key="item.team_tag_name" :label="item.team_tag_name" :value="item.team_tag_name">
                     </el-option>
                  </el-select>
               </div>
               <div style="float:left;width:40%;border:1px solid #ff0000;margin-top:15px">
                  <el-select v-model="value" filterable placeholder="敌方战队" @change="getHeroPool(value)">
                     <el-option v-for="item in allTeamInfo" :key="item.team_tag_name" :label="item.team_tag_name" :value="item.team_tag_name">
                     </el-option>
                  </el-select>
               </div>
               <div style="float:left;width:19%;border:1px solid #ff0000;height:40px;line-height:40px;margin-top:15px">
                  <el-radio v-model="radio" label="1" style="color:white">己方先选</el-radio>
                  <el-radio v-model="radio" label="2" style="color:white">己方后选</el-radio>
               </div>      
            </div>
            <div style="float:left;border:1px solid #ff0000;width:1050px;" >
               <div style="width:100%;">
                  <div class="hero_pool">
                   <!-- <div v-for="(item,index) in teamHeroPool">
                       <div v-for="item2 in item.match_player_heroes.hero_play_count" :key="item2.hero">
                           <span><img :src="item2.hero_icon" width="50"/>  </span>{{item2.count}}<span></span>
                       </div>// item2 所在div会被循环两次 如果item写在table上也一样
                   </div> -->
                   <table >
                       <tr v-for="(item,index) in teamHeroPool" :key="item.player_register_string_id">
                           <td v-if="index===0">一号位英雄池</td>
                           <td v-if="index===1">二号位英雄池</td>
                           <td v-if="index===2">三号位英雄池</td>
                           <td v-if="index===3">四号位英雄池</td>
                           <td v-if="index===4">五号位英雄池</td>
                           <td v-for="item2 in item.match_player_heroes.hero_play_count" :key="item2.hero" width="80"><img :src="item2.hero_icon" width="50"/>  {{item2.count}}</td>
                       </tr>
                      
                   </table>
                  </div>
               </div>
               <div id="strengthhero" style="margin-top:50px;margin:0 0;padding:0 0">
             <ul class="herobox" >
               <li><img src="../static/img/hero/eldertitan.jpg" width="60"  title="上古巨神" @click="bpClick($event)"/></li>
               <li><img src="../static/img/hero/undying.jpg" width="60" title="尸王" /></li>
               <li><img src="../static/img/hero/abaddon.jpg" width="60" title="亚巴顿" /></li>
               <li><img src="../static/img/hero/timbersaw.jpg" width="60" title="伐木机" /></li>
               <li><img src="../static/img/hero/omniknight.jpg" width="60" title="全能骑士" /></li>
               <li><img src="../static/img/hero/beastmaster.jpg" width="60" title="兽王" /></li>
               <li><img src="../static/img/hero/legioncommander.jpg" width="60" title="军团指挥官" /></li>
               <li><img src="../static/img/hero/wraithking.jpg" width="60" title="冥魂大帝" /></li>
               <li><img src="../static/img/hero/phoenix.jpg" width="60" title="凤凰" /></li>
               <li><img src="../static/img/hero/centaurwarrunner.jpg" width="60" title="半人马战士" /></li>
               <li><img src="../static/img/hero/clockwerk.jpg" width="60" title="发条技师" /></li>
               <li><img src="../static/img/hero/lifestealer.jpg" width="60" title="噬魂鬼" /></li>
               <li><img src="../static/img/hero/huskar.jpg" width="60" title="哈斯卡" /></li>
               <li><img src="../static/img/hero/earthspirit.jpg" width="60" title="大地之灵" /></li>
               <li><img src="../static/img/hero/underlord.jpg" width="60" title="孽主" /></li>
               <li><img src="../static/img/hero/tiny.jpg" width="60" title="小小" /></li>
               <li><img src="../static/img/hero/tusk.jpg" width="60" title="巨牙海民" /></li>
               <li><img src="../static/img/hero/pudge.jpg" width="60" title="屠夫" /></li>
               <li><img src="../static/img/hero/earthshaker.jpg" width="60" title="撼地者" /></li>
               <li><img src="../static/img/hero/axe.jpg" width="60" title="斧王" /></li>
               <li><img src="../static/img/hero/slardar.jpg" width="60" title="斯拉达" /></li>
               <li><img src="../static/img/hero/sven.jpg" width="60" title="斯温" /></li>
               <li><img src="../static/img/hero/kunkka.jpg" width="60" title="昆卡" /></li>
               <li><img src="../static/img/hero/nightstalker.jpg" width="60" title="暗夜魔王" /></li>
               <li><img src="../static/img/hero/doom.jpg" width="60" title="末日使者" /></li>
               <li><img src="../static/img/hero/treantprotector.jpg" width="60" title="树精卫士" /></li>
               <li><img src="../static/img/hero/sandking.jpg" width="60" title="沙王" /></li>
               <li><img src="../static/img/hero/chaosknight.jpg" width="60" title="混沌骑士" /></li>
               <li><img src="../static/img/hero/tidehunter.jpg" width="60" title="潮汐猎人" /></li>
               <li><img src="../static/img/hero/alchemist.jpg" width="60" title="炼金术士" /></li>
               <li><img src="../static/img/hero/lycan.jpg" width="60" title="狼人" /></li>
               <li><img src="../static/img/hero//mars.jpg" width="60" title="马尔斯" /></li>
               <li><img src="../static/img/hero/snapfire.jpg" width="60" title="电炎绝手" /></li>
               <li><img src="../static/img/hero/io.jpg" width="60" title="小精灵" /></li>
               <li><img src="../static/img/hero/spiritbreaker.jpg" width="60" title="裂魂人" /></li>
               <li><img src="../static/img/hero/brewmaster.jpg" width="60" title="酒仙" /></li>
               <li><img src="../static/img/hero/bristleback.jpg" width="60" title="钢背兽" /></li>
               <li><img src="../static/img/hero/magnus.jpg" width="60" title="马格纳斯" /></li>
               <li><img src="../static/img/hero/dragonknight.jpg" width="60" title="龙骑士" /></li>
             </ul>
               </div>
               <div id="agilityhero">
             <ul class="herobox" >
               <li><img src="../static/img/hero/juggernaut.jpg" width="60"  title="主宰" /></li>
               <li><img src="../static/img/hero/clinkz.jpg" width="60" title="克林克兹" /></li>
               <li><img src="../static/img/hero/viper.jpg" width="60" title="冥界亚龙" /></li>
               <li><img src="../static/img/hero/venomancer.jpg" width="60" title="剧毒术士" /></li>
               <li><img src="../static/img/hero/riki.jpg" width="60" title="力丸" /></li>
               <li><img src="../static/img/hero/drowranger.jpg" width="60" title="卓尔游侠" /></li>
               <li><img src="../static/img/hero/morphling.jpg" width="60" title="水人" /></li>
               <li><img src="../static/img/hero/nyx-assassin.jpg" width="60" title="司夜刺客" /></li>
               <li><img src="../static/img/hero/templarassassin.jpg" width="60" title="圣堂刺客" /></li>
               <li><img src="../static/img/hero/vengeful-spirit.jpg" width="60" title="复仇之魂" /></li>
               <li><img src="../static/img/hero/arcwarden.jpg" width="60" title="天穹守望者" /></li>
               <li><img src="../static/img/hero/nagasiren.jpg" width="60" title="娜迦海妖" /></li>
               <li><img src="../static/img/hero/trollwarlord.jpg" width="60" title="巨魔战将" /></li>
               <li><img src="../static/img/hero/phantomassassin.jpg" width="60" title="幻影刺客" /></li>
               <li><img src="../static/img/hero/phantomlancer.jpg" width="60" title="幻影长矛手" /></li>
               <li><img src="../static/img/hero/spectre.jpg" width="60" title="幽鬼" /></li>
               <!-- <li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="npc_dota_hero_nevermore" /></li> -->
               <li><img src="../static/img/hero/shadowfiend.jpg" width="60" title="影魔" /></li>
               <li><img src="../static/img/hero/lone-druid.jpg" width="60" title="德鲁伊" /></li>
               <li><img src="../static/img/hero/terrorblade.jpg" width="60" title="恐怖利刃" /></li>
               <li><img src="../static/img/hero/antimage.jpg" width="60" title="敌法师" /></li>
               <li><img src="../static/img/hero/slark.jpg" width="60" title="斯拉克" /></li>
               <li><img src="../static/img/hero/emberspirit.jpg" width="60" title="灰烬之灵" /></li>
               <li><img src="../static/img/hero/ursa.jpg" width="60" title="熊战士" /></li>
               <li><img src="../static/img/hero/sniper.jpg" width="60" title="狙击手" /></li>
               <li><img src="../static/img/hero/gyrocopter.jpg" width="60" title="矮人直升机" /></li>
               <li><img src="../static/img/hero/pangolier.jpg" width="60" title="石鳞剑士" /></li>
               <li><img src="../static/img/hero/mirana.jpg" width="60" title="米拉娜" /></li>
               <li><img src="../static/img/hero/meepo.jpg" width="60" title="米波" /></li>
               <li><img src="../static/img/hero/weaver.jpg" width="60" title="编织者" /></li>
               <li><img src="../static/img/hero/medusa.jpg" width="60" title="美杜莎" /></li>
               <li><img src="../static/img/hero/broodmother.jpg" width="60" title="育母蜘蛛" /></li>
               <li><img src="../static/img/hero//gyrocopter.jpg" width="60" title="虚空假面" /></li>
               <li><img src="../static/img/hero/bloodseeker.jpg" width="60" title="血魔" /></li>
               <li><img src="../static/img/hero/bountyhunter.jpg" width="60" title="赏金猎人" /></li>
               <li><img src="../static/img/hero/luna.jpg" width="60" title="露娜" /></li>
               <li><img src="../static/img/hero/monkeyking.jpg" width="60" title="大圣" /></li>
             </ul>
               </div>
               <div id="agilityhero">
             <ul class="herobox" >
               <li><img src="../static/img/hero/tinker.jpg" width="60"  title="修补匠" /></li>
               <li><img src="../static/img/hero/naturesprophet.jpg" width="60" title="先知" /></li>
               <li><img src="../static/img/hero/keeperofthelight.jpg" width="60" title="光法" /></li>
               <li><img src="../static/img/hero/skywrathmage.jpg" width="60" title="天怒" /></li>
               <li><img src="../static/img/hero/grimstroke.jpg" width="60" title="天涯墨客" /></li>
               <li><img src="../static/img/hero/zeus.jpg" width="60" title="宙斯" /></li>
               <li><img src="../static/img/hero/winterwyvern.jpg" width="60" title="寒冬飞龙" /></li>
               <li><img src="../static/img/hero/techies.jpg" width="60" title="炸弹人" /></li>
               <li><img src="../static/img/hero/witchdoctor.jpg" width="60" title="巫医" /></li>
               <li><img src="../static/img/hero/lich.jpg" width="60" title="巫妖" /></li>
               <li><img src="../static/img/hero/puck.jpg" width="60" title="帕克" /></li>
               <li><img src="../static/img/hero/pugna.jpg" width="60" title="帕格纳" /></li>
               <li><img src="../static/img/hero/disruptor.jpg" width="60" title="干扰者" /></li>
               <li><img src="../static/img/hero/dazzle.jpg" width="60" title="戴泽" /></li>
               <li><img src="../static/img/hero/leshrac.jpg" width="60" title="拉席克" /></li>
               <li><img src="../static/img/hero/rubick.jpg" width="60" title="拉比克" /></li>
               <li><img src="../static/img/hero/shadowdemon.jpg" width="60" title="暗影恶魔" /></li>
               <li><img src="../static/img/hero/shadowshaman.jpg" width="60" title="暗影萨满" /></li>
               <li><img src="../static/img/hero/warlock.jpg" width="60" title="术士" /></li>
               <li><img src="../static/img/hero/jakiro.jpg" width="60" title="杰奇洛" /></li>
               <li><img src="../static/img/hero/deathprophet.jpg" width="60" title="死亡先知" /></li>
               <li><img src="../static/img/hero/outworlddevourer.jpg" width="60" title="黑鸟" /></li>
               <li><img src="../static/img/hero/crystalmaiden.jpg" width="60" title="冰女" /></li>
               <li><img src="../static/img/hero/silencer.jpg" width="60" title="沉默术士" /></li>
               <li><img src="../static/img/hero/queenofpain.jpg" width="60" title="痛苦女王" /></li>
               <li><img src="../static/img/hero/necrophos.jpg" width="60" title="瘟疫法师" /></li>
               <li><img src="../static/img/hero/invoker.jpg" width="60" title="祈求者" /></li>
               <li><img src="../static/img/hero/oracle.jpg" width="60" title="神谕者" /></li>
               <li><img src="../static/img/hero/bane.jpg" width="60" title="霍乱之源" /></li>
               <li><img src="../static/img/hero/visage.jpg" width="60" title="维萨吉" /></li>
               <li><img src="../static/img/hero/lina.jpg" width="60" title="莉娜" /></li>
               <li><img src="../static/img/hero//lion.jpg" width="60" title="莱恩" /></li>
               <li><img src="../static/img/hero/void-spirit.jpg" width="60" title="虚空之灵" /></li>
               <li><img src="../static/img/hero/batrider.jpg" width="60" title="蝙蝠骑士" /></li>
               <li><img src="../static/img/hero/enigma.jpg" width="60" title="谜团" /></li>
               <li><img src="../static/img/hero/ancient.jpg" width="60" title="远古冰魄" /></li>
               <li><img src="../static/img/hero/dark-willow.jpg" width="60" title="邪影芳灵" /></li>
               <li><img src="../static/img/hero/chen.jpg" width="60" title="陈" /></li>
               <li><img src="../static/img/hero/stormspirit.jpg" width="60" title="风暴之灵" /></li>
               <li><img src="../static/img/hero/windranger.jpg" width="60" title="风行者" /></li>
               <li><img src="../static/img/hero/ogremagi.jpg" width="60" title="蓝胖" /></li>
               <li><img src="../static/img/hero/enchantress.jpg" width="60" title="魅惑魔女" /></li>
               <li><img src="../static/img/hero/dark-seer.jpg" width="60" title="黑暗贤者" /></li>
             </ul>
               </div>
               <div>
                  <el-button size="small" @click="savebpinfo">阵容数据分析</el-button>
               </div>
            </div>
      </div>
      <div id="team_bp" class="team_bp">
        <div id="left_team" style="float:left">
            <ul style="margin-top:0px;margin-left:0px;padding-left:15px">
               <li style="width:65px;height:65px;border:1px solid #676767;" class="shadow_hover"  data-bp-id=1><div>1</div></li>
               <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px" data-bp-id=3><div>3</div></li>
               <li style="width:65px;height:65px;border:1px solid #676767;margin-top:15px" data-bp-id=5><div>5</div></li>
               <li style="" data-bp-id=7>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div style="">7</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team1_operator1" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team1_operator1_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
               </li>
               <li data-bp-id=10> 
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>10</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team1_operator2" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team1_operator2_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
               </li>
               <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px" data-bp-id=11> <div>11</div></li>
               <li style="width:65px;height:65px;border:1px solid #676767;margin-top:15px" data-bp-id=13> <div>13</div></li>
               <li data-bp-id=16>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>16</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team1_operator3" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team1_operator3_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
               </li>
               <li data-bp-id=18>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>18</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team1_operator4" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team1_operator4_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
               </li>
               <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px"data-bp-id=20> <div>20</div></li>
               <li data-bp-id=21>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>21</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team1_operator5" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team1_operator5_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
               </li>
            </ul>
        </div>
        <div id="right_team" style="float:left">
            <ul style="margin-top:0px;padding-left:15px">
                <li style="width:65px;height:65px;border:1px solid #676767;" data-bp-id=2>
                  <div>2</div>
                </li>
                <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px" data-bp-id=4> <div>4</div></li>
                <li style="width:65px;height:65px;border:1px solid #676767;margin-top:15px" data-bp-id=6> <div>6</div></li>
                <li data-bp-id=8>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>8</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team2_operator1" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team2_operator1_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
                </li>
                <li data-bp-id=9>
                   <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>9</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team2_operator2" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team2_operator2_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
                </li>
                <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px"data-bp-id=12> <div>12</div></li>
                <li style="width:65px;height:65px;border:1px solid #676767;margin-top:15px" data-bp-id=14> <div>14</div></li>
                <li data-bp-id=15>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>15</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team2_operator3" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team2_operator3_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
                </li>
                <li data-bp-id=17>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>17</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team2_operator4" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team2_operator4_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
                </li>
                <li style="width:65px;height:65px; border:1px solid #676767;margin-top:15px"data-bp-id=19> <div>19</div></li>
                <li data-bp-id=22>
                  <div style="width:65px;height:65px; border:1px solid #676767;margin-top:15px;float:left">
                     <div>8</div>
                  </div>
                  <div style="float:left;margin-top:15px">
                     <div style="margin-top:10px;">
                        <select id="team2_operator15" >
                          <option v-for="item in teamHeroPool" :key="item.match_player_id" value ="item.match_player_id">{{item.position}}</option>
                        </select>
                     </div>
                      <div style="margin-top:10px">
                       <select id="team2_operator5_lane">
                       <option value ="上路">上路</option>
                       <option value =">中路">中路</option>
                       <option value="下路">下路</option>
                     </select>
                     </div>  
                  </div>
                  <div style="clear:both"></div>
                </li>
            </ul>
        </div>
      </div>
      </div>
      
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
    computed:{
        ...mapState(['allTeamInfo']),//获取所有的战队名称
        ...mapState(['teamHeroPool'])
    },
    created(){
      this.$store.dispatch('getAllTeamInfo')
    },
    methods: {
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
               /*liele[i2].appendChild(node)*/
            }
         }
         this.bpOrder=this.bpOrder+1
    	},
        getHeroPool(teamtagName){
            var param={version:this.$store.state.version,teamtagname:teamtagName}
            this.$store.dispatch('getTeamHeroPool',param)
        },
        getHeroPool1(teamtagName){//己方战队
            var param={version:this.$store.state.version,teamtagname:teamtagName}
            /*Vue.axios.get('./api/getteamheropool',{params:param}).then((data)=>{
               
            })*/
        },
        test(){
         console.log(this.radio)
        },
        savebpinfo(){
         /* team1={name:"",player1:{hero:}}
         if (this.radio==1){

         }*/
        },

    }
  }
</script>
<style>
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
  width:100%;border:1px solid #ff0000
}
ul.hero_pool li{
margin-bottom:7px;
}
.team_bp{
float:left;border:1px solid #ff0000
}
.shadow_hover{
   box-shadow:0px 0px 10px #888888;
}
#left_team{
   width:170px;border:1px solid #ff0000
}
#right_team{
   width:170px;
}
#left_team ul li div,#right_team ul li div{
  
   font-size:13px
}
</style>
