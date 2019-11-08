package obj
import(
	//"gopkg.in/mgo.v2/bson"
	"github.com/dotabuff/manta/dota"
	"strconv"
	"strings"
	"time"
)
type CDotaGameInfo struct{
		//Id_       bson.ObjectId `bson:"_id" json:"_id,omitempty"`
		Version string `bson:"version" json:"version"`
		MatchId uint64 `bson:"match_id" json:"match_id"`
		GameMode int32 `protobuf:"varint,2,opt,name=game_mode,json=gameMode" json:"game_mode,omitempty" bson:"game_mode,omitempty"`
		GameWinner string `json:"game_winner,omitempty" bson:"game_winner"`
		GameLoser string `json:"game_loser,omitempty" bson:game_loser`
		Leagueid             uint32                                     `protobuf:"varint,5,opt,name=leagueid" json:"leagueid,omitempty" bson:"leagueid,omitempty"`
		PicksBans            []*dota.CGameInfo_CDotaGameInfo_CHeroSelectEvent `protobuf:"bytes,6,rep,name=picks_bans,json=picksBans" json:"picks_bans,omitempty" bson:"picks_bans,omitempty"`
		RadiantTeamTag       string                                     `protobuf:"bytes,9,opt,name=radiant_team_tag,json=radiantTeamTag" json:"radiant_team_tag,omitempty" bson:"radiant_team_tag,omitempty"`
		DireTeamTag          string                                     `protobuf:"bytes,10,opt,name=dire_team_tag,json=direTeamTag" json:"dire_team_tag,omitempty" bson:"dire_team_tag,omitempty"`
		EndTime              uint32                                     `protobuf:"varint,11,opt,name=end_time,json=endTime" json:"end_time,omitempty" bson:"end_time,omitempty"`
}


type CMsgDOTAMatch struct{
	Duration         uint32                             `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty" bson:"duration,omitempty"`
	StartTime            string                             `protobuf:"fixed32,4,opt,name=startTime" json:"start_time,omitempty" bson:"start_time,omitempty"`
	Players              []*dota.CMsgDOTAMatch_Player        `protobuf:"bytes,5,rep,name=players" json:"players,omitempty" bson:"players,omitempty"`
	MatchId              uint64                             `protobuf:"varint,6,opt,name=match_id,json=match_id" json:"match_id,omitempty" bson:"match_id,omitempty"`
	TowerStatus          []uint32 /*防御塔状态*/                           `protobuf:"varint,8,rep,name=tower_status,json=towerStatus" json:"tower_status,omitempty" bson:"tower_status,omitempty"`
	BarracksStatus       []uint32 /* 兵营状态*/                            `protobuf:"varint,9,rep,name=barracks_status,json=barracksStatus" json:"barracks_status,omitempty"`
	FirstBloodTime       uint32                             `protobuf:"varint,12,opt,name=first_blood_time,json=firstBloodTime" json:"first_blood_time,omitempty" bson:"first_blood_time,omitempty"`
	AverageSkill         uint32                             `protobuf:"varint,18,opt,name=average_skill,json=averageSkill" json:"average_skill,omitempty" bson:"average_skill,omitempty"`
	RadiantTeamId        uint32                             `protobuf:"varint,20,opt,name=radiant_team_id,json=radiantTeamId" json:"radiant_team_id,omitempty" bson:"radiant_team_id,omitempty"`
	DireTeamId           uint32                             `protobuf:"varint,21,opt,name=dire_team_id,json=direTeamId" json:"dire_team_id,omitempty" bson:"dire_team_id,omitempty"`
	Leagueid             uint32                             `protobuf:"varint,22,opt,name=leagueid" json:"leagueid,omitempty" bson:"leagueid,omitempty"`
	RadiantTeamName      string                             `protobuf:"bytes,23,opt,name=radiant_team_name,json=radiantTeamName" json:"radiant_team_name,omitempty" bson:"radiant_team_name,omitempty"`
	DireTeamName         string                             `protobuf:"bytes,24,opt,name=dire_team_name,json=direTeamName" json:"dire_team_name,omitempty" bson:"dire_team_name,omitempty"`
	MatchSeqNum          uint64                             `protobuf:"varint,33,opt,name=match_seq_num,json=matchSeqNum" json:"match_seq_num,omitempty" bson:"match_seq_num,omitempty"`
	RadiantTeamTag       string                             `protobuf:"bytes,37,opt,name=radiant_team_tag,json=radiantTeamTag" json:"radiant_team_tag,omitempty" bson:"radiant_team_tag,omitempty"`
	DireTeamTag          string                             `protobuf:"bytes,38,opt,name=dire_team_tag,json=direTeamTag" json:"dire_team_tag,omitempty" bson:"dire_team_tag,omitempty"`
	SeriesId             uint32                             `protobuf:"varint,39,opt,name=series_id,json=seriesId" json:"series_id,omitempty" bson:"series_id,omitempty"`
	SeriesType           uint32                             `protobuf:"varint,40,opt,name=series_type,json=seriesType" json:"series_type,omitempty" bson:"series_type,omitempty"`
	Engine               uint32                             `protobuf:"varint,44,opt,name=engine" json:"engine,omitempty"`
	PrivateMetadataKey   uint32                             `protobuf:"fixed32,47,opt,name=private_metadata_key,json=privateMetadataKey" json:"private_metadata_key,omitempty" bson:"private_metadata_key,omitempty"`
	MatchOutcome         dota.EMatchOutcome                 `protobuf:"varint,50,opt,name=match_outcome,json=matchOutcome,enum=EMatchOutcome,def=0" json:"match_outcome,omitempty" bson:"match_outcome,omitempty"`
	TournamentId         uint32                             `protobuf:"varint,51,opt,name=tournament_id,json=tournamentId" json:"tournament_id,omitempty" bson:"tournament_id,omitempty"`
	TournamentRound      uint32                             `protobuf:"varint,52,opt,name=tournament_round,json=tournamentRound" json:"tournament_round,omitempty" bson:"tournament_round,omitempty"`
	PreGameDuration      uint32                             `protobuf:"varint,53,opt,name=pre_game_duration,json=preGameDuration" json:"pre_game_duration,omitempty" bson:"pre_game_duration,omitempty"`
	
}
type ItemPurchase struct{
	 Item uint32
	 time string
}

type HeroLastHit struct{
	Pre5Count uint32 `bson:"pre_5_count" bson:"pre_5_count"`
	Pre10Count uint32 `bson:"pre_10_count" json:"pre_10_count"`
	Pre15Count uint32 `bson:"pre_15_count" json:"pre_15_count"`
	Pre20Count uint32 `bson:"pre_20_count" json:"pre_20_count"`
	Pre25Count uint32 `bson:"pre_25_count" json:"pre_25_count"`
}
//击杀英雄的数量跟输赢关系不太大。关系最大的是阵容（大哥输出）是否被克制（在一定的经济差范围内）,前期被克制看后期团战（经济差不大的情况下）,
type HeroGold struct{
	Pre5MinGold float64 `bson:"pre_5min_gold" json:"pre_5min_gold"`
	Pre10MinGold float64 `bson:"pre_10min_gold" json:"pre_10min_gold"`
	Pre15MinGold float64 `bson:"pre_15min_gold" json:"pre_15min_gold"`
	Pre20MinGold float64 `bson:"pre_20min_gold" json:"pre_20min_gold"`
	Pre25MinGold float64 `bson:"pre_25min_gold" json:"pre_25min_gold"`
}
type TowerState struct{
	Pre10MinTower uint32 `bson:"pre_10min_tower" json:"pre_10min_tower"`
	Pre20MinTower uint32 `bson:"pre_20min_tower" json:"pre_20min_tower"`
	Pre30MinTower uint32 `bson:"pre_30min_tower" json:"pre_30min_tower"`
}
type PlayersHeroInfo struct{
	GameTeam uint32  `bson:"game_team" json:"game_team"`//天辉夜魇
	HeroName string  `bson:"hero_name" json:"hero_name"`
	Player string `bson:"player" json:"player"`
	AccountId            uint32   `protobuf:"varint,1,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PlayerSlot           uint32 `protobuf:"varint,3,opt,name=player_slot,json=playerSlot" json:"player_slot,omitempty"`
	LastHit HeroLastHit `bson:"last_hit,omitempty" json:"last_hit"`//每5分钟补刀数
	Denis HeroLastHit `bson:"denis,omitempty" json:"denis"`//每5分钟反补数
	LevelUpTimes         []uint32  `protobuf:"varint,22,rep,name=level_up_times,json=levelUpTimes" json:"level_up_times,omitempty" bson:"level_up_times,omitempty"`
	Gold HeroGold `bson:"gold,omitempty" json:"gold,omitempty"`//每5分钟金钱信息 
	//Consumables []*ConsumablesInfo `bson:"Consumables" json:"Consumables"`//中单对线消耗品使用
	//PurchaseInfo []*dota.CDOTAMatchMetadata_Team_ItemPurchase  `bson:"purchase_Info" json:"purchase_Info"`//装备购买信息 
	AbilityUpgrades []uint32  `bson:"ability_upgrades" json:"ability_upgrades"`//技能加点升级信息
	InitLane string `bson:"init_lane" json:"init_lane"`
	ChangeLaneTime string `bson:"change_lane_time,omitempty" json:"change_lane_time,omitempty"`
	//PureLastHitPec float32 `bson:"pure_last_hit_pec" json:"pure_last_hit"`
	InventorySnapshot []*dota.CDOTAMatchMetadata_Team_InventorySnapshot `bson:"inventory_snap_shot" json:"inventory_snap_shot"`
	Item []*dota.CDOTAMatchMetadata_Team_ItemPurchase `protobuf:"bytes,6,rep,name=items" json:"items,omitempty"`

	//Lane AgainstLaneInfo `bson:"against_lane_info,omitempty" json:"against_lane_info,omitempty"`
}

type CMsgMatchDetails struct{
	MatchId              uint64                             `protobuf:"varint,6,opt,name=match_id,json=matchId" json:"match_id,omitempty" bson:"match_id,omitempty"`
	RadiantTeamTag       string                             `protobuf:"bytes,37,opt,name=radiant_team_tag,json=radiantTeamTag" json:"radiant_team_tag,omitempty" bson:"radiant_team_tag,omitempty"`
	DireTeamTag          string                             `protobuf:"bytes,38,opt,name=dire_team_tag,json=direTeamTag" json:"dire_team_tag,omitempty" bson:"dire_team_tag,omitempty"`
	PlayersHeroesDets	     []*PlayersHeroInfo `bson:"players_heroes_dets" json:"players_heroes_dets"`
	AgainstHeroes 		 string `bson:"AgainstHeroes" json:"AgainstHeroes"`	
	RadiantDestoryTower TowerState `bson:"radiant_destory_tower" json:"radiant_destory_tower"`
	DireDestoryTower TowerState `bson:"dire_destory_tower" json:"dire_destory_tower"`
}


//选手英雄操作细节信息



/*      游戏时间转换成字符串的函数                    */
func Durformat(durint int64) string {
	var durtime, h, min, sec string
	durstring := strconv.FormatInt(durint, 10) + "s" //将duration 由int64转换成字符串
	dur, _ := time.ParseDuration(durstring)          //将时间字符串转换成Time.duration类型
	dur2 := dur.String()
	dur2 = strings.Replace(dur2, "h", ":", -1)
	dur2 = strings.Replace(dur2, "m", ":", -1) //-1表示替换所有的m
	dur2 = strings.Replace(dur2, "s", ":", -1)
	strarr := strings.Split(dur2, ":")
	if durint > 3600 {
		h = strarr[0] + ":"

		if len(strarr[1]) < 2 {
			min = "0" + strarr[1] + ":"
		} else {
			min = strarr[1] + ":"
		}

		if len(strarr[2]) < 2 {
			sec = "0" + strarr[2]
		} else {
			sec = strarr[2]
		}
		durtime = h + min + sec
	} else {
		if len(strarr[0]) < 2 {
			min = "0" + strarr[1] + ":"
		} else {
			min = strarr[0] + ":"
		}

		if len(strarr[1]) < 2 {
			sec = "0" + strarr[2]
		} else {
			sec = strarr[1]
		}
		durtime = min + sec
	}
	return durtime
}