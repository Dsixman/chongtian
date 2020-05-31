package obj
import(
	//"gopkg.in/mgo.v2/bson"
	"github.com/dotabuff/manta/dota"
	"strconv"
	"strings"
	"time"
	//"fmt"
)

type Player struct{
	MatchPlayerId string `bson:"player_register_string_id" json:"player_register_string_id"`
	PlayerDota2NumId uint32 `bson:"player_dota2_register_num_id" json:"player_dota2_register_num_id"`
	PlayerSteamId uint64 `bson:"player_steam_id,omitempty" json:"player_steam_id,omitempty"`
	SecondNumId uint32 `bson:"second_num_id" json:"second_num_id"`
	ThirdNumId  uint32 `bson:"third_num_id" json:"third_num_id"`
	FourthNumId uint32 `bson:"fourth_num_id" json:"forth_num_id"`
	TeamName string `bson:"team_name" json:"team_name"`
	TeamNameTag string `bson:"team_name_tag" json:"team_name_tag"`
	Position string `bson:"position" json:"position"`
	PlayerState string `bson:"player_state" json:"player_state"`
	MatchPlayerHeroes *PlayerCommonHero `bson:"match_player_heroes" json:"match_player_heroes"`
	AlternatePlayerHeroes *PlayerCommonHero `bson:"alternate_player_heroes" json:"alternate_player_heroes"`
	OldClubPlayerHeroes *PlayerCommonHero`bson:"old_club_player_heroes" json:"old_club_player_heroes"`
	RankPlayerHeroes []*HeroCount `bson:"rank_player_heroes" json:"rank_player_heroes"`
}
type MatchHeroCount struct{
	AllCount int `bson:"all_count" json:"all_count"`
	WinCount int `bson:"win_count" json:"win_count"`
	LoseCount int `bson:"lose_count" json:"lose_count"`
}
type HeroCount struct{
	Hero string `bson:"hero" json:"hero"`
	HeroIcon string `bson:"hero_icon" json:"hero_icon"`
	Version string `bson:"version" json:"version"`
	RankHeroCount *MatchHeroCount `bson:"rank" json:"rank"`
	LeagueHeroCount *MatchHeroCount `bson:"league" json:"league"`
}
type PlayerCommonHero struct{
/*	ClubTeam string `bson:"club_team" json:"club_team"`
	ClubTeamTag string `bson:"club_team_tag" json:"club_team_tag"`*/
	HeroPlayCount []*HeroCount `bson:"hero_play_count" json:"hero_play_count"`
	//RankPosition string `bson:"rank_position" json:"rank_position"`
} 

type CDotaGameInfo struct{
		//Id_       bson.ObjectId `bson:"_id" json:"_id,omitempty"`
	Version string `bson:"version" json:"version"`
	MatchId uint64 `bson:"match_id" json:"match_id"`
	GameMode int32 `protobuf:"varint,2,opt,name=game_mode,json=gameMode" json:"game_mode,omitempty" bson:"game_mode,omitempty"`
	GameWinner string `json:"game_winner,omitempty" bson:"game_winner"`
	GameLoser string `json:"game_loser,omitempty" bson:game_loser`
	Leagueid             uint32                                     `protobuf:"varint,5,opt,name=leagueid" json:"leagueid,omitempty" bson:"leagueid,omitempty"`
	PicksBans            []*BPInfo `protobuf:"bytes,6,rep,name=picks_bans,json=picksBans" json:"picks_bans,omitempty" bson:"picks_bans,omitempty"`
	RadiantTeamTag       string                                     `protobuf:"bytes,9,opt,name=radiant_team_tag,json=radiantTeamTag" json:"radiant_team_tag,omitempty" bson:"radiant_team_tag,omitempty"`
	DireTeamTag          string                                     `protobuf:"bytes,10,opt,name=dire_team_tag,json=direTeamTag" json:"dire_team_tag,omitempty" bson:"dire_team_tag,omitempty"`
	RadiantTeamName      string                             `protobuf:"bytes,23,opt,name=radiant_team_name,json=radiantTeamName" json:"radiant_team_name,omitempty" bson:"radiant_team_name,omitempty"`
	DireTeamName         string                             `protobuf:"bytes,24,opt,name=dire_team_name,json=direTeamName" json:"dire_team_name,omitempty" bson:"dire_team_name,omitempty"`
	EndTime              uint32                                     `protobuf:"varint,11,opt,name=end_time,json=endTime" json:"end_time,omitempty" bson:"end_time,omitempty"`
	GameWinnerBp []*BPInfo `json:"game_winer_bp,omitempty" bson:"game_winer_bp,omitempty"`
	GameLoserBp []*BPInfo `json:"game_loser_bp,omitempty" bson:"game_loser_bp,omitempty"` 
	TopLane string `bson:"top_lane" json:"top_lane"`
	MidLane string `bson:"mid_lane" json:"mid_lane"`
	BottomLane string `bson:"bottom_lane" json:"bottom_lane"`
}
type BPInfo struct{
	IsPick bool `json:"is_pick" bson:"is_pick"`
	Team uint32 `json:"team" bson:"team"`
	TeamName string `json:"team_name" bson:"team_name"`
	HeroId uint32 `json:"hero_id" bson:"hero_id"`
	HeroName string `json:"hero_name" bson:"heor_name"`
	HeroIcon string `json:"hero_icon" bson:"hero_icon"`
	Position string `json:"position" bson:"position"`
	PlayerName string `json:"player_name" bson:"player_name"`
}

type CMsgDOTAMatch struct{
	Duration         uint32                             `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty" bson:"duration,omitempty"`
	DurationStr string `json:"duration_str" bson:"duration_str"`
	Version  string `json:"version" bson:"version"`
	StartTime            string                             `protobuf:"fixed32,4,opt,name=startTime" json:"start_time,omitempty" bson:"start_time,omitempty"`
	Players              []*DOTAMatch_Player       `protobuf:"bytes,5,rep,name=players" json:"players,omitempty" bson:"players,omitempty"`
	MatchId              uint64                             `protobuf:"varint,6,opt,name=match_id,json=match_id" json:"match_id,omitempty" bson:"match_id,omitempty"`
	TowerStatus          []uint32 /*防御塔状态*/             `protobuf:"varint,8,rep,name=tower_status,json=towerStatus" json:"tower_status,omitempty" bson:"tower_status,omitempty"`
	BarracksStatus       []uint32 /* 兵营状态*/              `protobuf:"varint,9,rep,name=barracks_status,json=barracksStatus" json:"barracks_status,omitempty" bson:"barracks_status"`
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
	RadiantTeamScore     uint32                             `protobuf:"varint,48,opt,name=radiant_team_score,json=radiantTeamScore" json:"radiant_team_score,omitempty" bson:"radiant_team_score"`
	DireTeamScore        uint32                             `protobuf:"varint,49,opt,name=dire_team_score,json=direTeamScore" json:"dire_team_score,omitempty" bson:"dire_team_score"`
}

type DOTAMatch_Player struct{
	HeroIcon string `json:"hero_icon" bson:"hero_icon"`
	Player *dota.CMsgDOTAMatch_Player `json:"player" bson:"player"`
	Item []string `json:"item" bson:"item"`
	Buffs []*PlayerPermanentBuff `json:"buffs" bson:"buffs"`
}
type PlayerPermanentBuff struct{
	BuffId uint32 `json:"buff_id" bson:"buff_id"`
	PermanentBuff string `bson:"buffs" json:"buffs"`
	Count uint32 `bson:"count" json:"count"`
}
/*type ItemPurchase struct{
	 Item uint32
	 time string
}*/

type HeroLastHit struct{
	Pre5Count uint32 `bson:"pre_5_count" bson:"pre_5_count"`
	Pre10Count uint32 `bson:"pre_10_count" json:"pre_10_count"`
	Pre15Count uint32 `bson:"pre_15_count" json:"pre_15_count"`
	Pre20Count uint32 `bson:"pre_20_count" json:"pre_20_count"`
	Pre25Count uint32 `bson:"pre_25_count" json:"pre_25_count"`
}
//击杀英雄的数量跟输赢关系不太大。关系最大的是阵容（大哥输出）是否被克制（在一定的经济差范围内）,前期被克制看后期团战（经济差不大的情况下）,
type GoldType struct{
	LastHitGold float64 `bson:"last_hit_gold" json:"last_hit_gold"`
	CombatGold float64 `bson:"combat_gold" json:"combat_gold"`
	CreepsGold float64 `bson:"creeps_gold" json:"creeps_gold"`
	AllGold float64 `bson:"all_gold" json:"all_gold"`
	TestGold float64 `bson:"test_gold" json:"test_gold"` 
}
type HeroGold struct{
	Pre5MinGold *GoldType `bson:"pre_5min_gold" json:"pre_5min_gold"`
	Pre10MinGold *GoldType `bson:"pre_10min_gold" json:"pre_10min_gold"`
	Pre15MinGold *GoldType `bson:"pre_15min_gold" json:"pre_15min_gold"`
	Pre20MinGold *GoldType `bson:"pre_20min_gold" json:"pre_20min_gold"`
	Pre25MinGold *GoldType `bson:"pre_25min_gold" json:"pre_25min_gold"`
}
type TowerState struct{
	Pre10MinTower uint32 `bson:"pre_10min_tower" json:"pre_10min_tower"`
	Pre20MinTower uint32 `bson:"pre_20min_tower" json:"pre_20min_tower"`
	Pre30MinTower uint32 `bson:"pre_30min_tower" json:"pre_30min_tower"`
}

type PlayersHeroInfo struct{
	GameTeam uint32  `bson:"game_team" json:"game_team"`//天辉夜魇
	HeroName string  `bson:"hero_name" json:"hero_name"`
	HeroIcon string  `bson:"hero_icon" json:"hero_icon"`
	GameResult string `bson:"game_result" json:"game_result"`
	Player string `bson:"player" json:"player"`
	Position string `bson:"position" json:"position"`
	AccountId            uint32   `protobuf:"varint,1,opt,name=account_id,json=accountId" bson:"account_id,omitempty"`
	PlayerSlot           uint32 `protobuf:"varint,3,opt,name=player_slot,json=playerSlot" bson:"player_slot,omitempty"`
	LastHit *HeroLastHit `bson:"last_hit,omitempty" json:"last_hit"`//每5分钟补刀数
	Denis *HeroLastHit `bson:"denis,omitempty" json:"denis"`//每5分钟反补数
	LevelUpTimes         []uint32  `protobuf:"varint,22,rep,name=level_up_times,json=levelUpTimes" json:"level_up_times,omitempty" bson:"level_up_times,omitempty"`
	LevelUpTimeStr []string `json:"level_up_time_str" bson:"level_up_time_str"`
	Gold *HeroGold `bson:"gold,omitempty" json:"gold,omitempty"`//每5分钟金钱信息 
	//Consumables []*ConsumablesInfo `bson:"Consumables" json:"Consumables"`//中单对线消耗品使用
	//PurchaseInfo []*dota.CDOTAMatchMetadata_Team_ItemPurchase  `bson:"purchase_Info" json:"purchase_Info"`//装备购买信息 
	AbilityUpgrades []uint32  `bson:"ability_upgrades" json:"ability_upgrades"`//技能加点升级信息
	AbilityData []*Ability  `bson:"ability_data" json:"ability_data"`//技能加点升级信息
	InitLane string `bson:"init_lane" json:"init_lane"`
	ChangeLaneTime string `bson:"change_lane_time,omitempty" json:"change_lane_time,omitempty"`
	//PureLastHitPec float32 `bson:"pure_last_hit_pec" json:"pure_last_hit"`
	InventorySnapshot []*PlayerInventorySnapshot `bson:"inventory_snap_shot" json:"inventory_snap_shot"`
	DeathDets []*KDADetails `bson:"death_dets" json:"death_dets"`
	Item []*ItemDetails `protobuf:"bytes,6,rep,name=items" json:"items,omitempty" bson:"items"`
	CourierDeathTime []uint32 `bson:"courier_death_time" json:"courier_death_time" ` 
	KillCourierCount int `bson:"kill_courier_count" json:"kill_courier_count"`
	Consumable []*PlayerConsumable `bson:"consumable" json:"consumable"`
	//Lane AgainstLaneInfo `bson:"against_lane_info,omitempty" json:"against_lane_info,omitempty"`
}
//对线期消耗品
type PlayerConsumable struct{
	ItemId uint32 `json:"item_id" bson:"item_id"`
	ItemIcon string `json:"item_icon" bson:"item_icon"`
	Count uint32 `json:"count" bson:"count"`
	TimeSlice []int32 `json:"time_slice" bson:"time_slice"`
}
type PlayerInventorySnapshot struct{
	Items []string `json:"items" bson:"items"`
	GameTime int32 `json:"game_time" bson:"game_time"`
	GameTimeStr string `json:"game_time_str" bson:"game_time_str"`
	Kills    uint32 `json:"kills" bson:"kills"`
	Deaths uint32 `json:"deaths" bson:"deaths"`
	Assists   uint32 `json:"assists"  bson:"assists"`
	Level     uint32 `json:"level" bson:"level"`
}
type ItemDetails struct{
	ItemId uint32 `json:"item_id" bson:"item_id"`
	ItemIcon string `json:"item_icon" bson:"item_icon"`
	PurchaseTime string `json:"purchase_time" bson:"purchase_time"`
	PurchaseTimeNum int32 `json:"purchase_time_Num" bson:"purchase_time_Num"`
}
type KDADetails struct{
	Attacker string `bson:"attacker" json:"attacker"`
	Assistant []string `bson:"assistant" json:"assistant"`
	DeathTarget string `bson:"death_target" json:"death_target"`
	Time int64 `bson:"time" json:"time"`
	TimeStr string `bson:"time_str" json:"time_str"`
}
type CMsgMatchDetails struct{
	MatchId              uint64   `protobuf:"varint,6,opt,name=match_id,json=match_Id" json:"match_id,omitempty" bson:"match_id,omitempty"`
	Version  string `json:"version" bson:"version"`
	Duration         uint32  `protobuf:"varint,3,opt,name=duration" json:"duration,omitempty" bson:"duration,omitempty"`
	DurationStr string `json:"duration_str" bson:"duration_str"`
	RadiantTeamTag       string   `protobuf:"bytes,37,opt,name=radiant_team_tag,json=radiantTeamTag" json:"radiant_team_tag,omitempty" bson:"radiant_team_tag,omitempty"`
	DireTeamTag          string   `protobuf:"bytes,38,opt,name=dire_team_tag,json=direTeamTag" json:"dire_team_tag,omitempty" bson:"dire_team_tag,omitempty"`
	RadiantTeamName      string   `protobuf:"bytes,23,opt,name=radiant_team_name,json=radiantTeamName" json:"radiant_team_name,omitempty" bson:"radiant_team_name,omitempty"`
	DireTeamName         string   `protobuf:"bytes,24,opt,name=dire_team_name,json=direTeamName" json:"dire_team_name,omitempty" bson:"dire_team_name,omitempty"`
	PlayersHeroesDets  []*PlayersHeroInfo `bson:"players_heroes_dets" json:"players_heroes_dets"`
	AgainstHeroes 		 string `bson:"AgainstHeroes" json:"AgainstHeroes"`	
	TopLane string `bson:"top_lane" json:"top_lane"`
	MidLane string `bson:"mid_lane" json:"mid_lane"`
	BottomLane string `bson:"bottom_lane" json:"bottom_lane"`	
	RadiantDestoryTower TowerState `bson:"radiant_destory_tower" json:"radiant_destory_tower"`
	DireDestoryTower TowerState `bson:"dire_destory_tower" json:"dire_destory_tower"`
	KDADets []*KDADetails `bson:"kda_dets" json:"kda_dets"`
	IsMidSolo int `bson:"is_mid_solo" json:"is_mid_solo"`
}

type League struct{
	SeriesId uint32 `bson:"series_id" json:"series_id"`
	LeagueId uint32 `bson:"league_id" json:"league_id"`
	LeagueIcon string `bson:"league_icon" json:"league_icon"`
	LeagueName string `bson:"league_name" json:"league_name"`
}

type Team struct{
	TeamName  string `bson:"team_name" json:"team_name"`
	TeamTagName  string `bson:"team_tag_name" json:"team_tag_name"`
	TeamId uint32 `bson:"team_id" json:"team_id"`
	TeamIcon string `bson:"team_icon" json:"team_icon"`
}
//队友支援/gank中路情况
type MidAssisant struct{
	Assisant1 string `bson:"assisant1" json:"assisant1"`
	Assisant1Icon string `bson:"assisant1_icon" json:"assisant1_icon"`
	Assisant2 string `bson:"assisant2" json:"assisant2"`
	Assisant2Icon string `bson:"assisant2_icon" json:"assisant2_icon"`
	Assisant3 string `bson:"assisant3" json:"assisant3"`
	Assisant3Icon string `bson:"assisant3_Icon" json:"assisant3_Icon"`
	Assisant4 string `bson:"assisant4" json:"assisant4"`
	Assisant4Icon string `bson:"assisant4_Icon" json:"assisant4_Icon"`
}
type MidAgainst struct{
	MatchId uint64 `bson:"match_id" json:"match_id"`
	Hero1PlayerId uint32 `bson:"hero1_player_id" json:"hero1_player_id"`
	Hero1Name string `bson:"hero1_name" json:"hero1_name"`
	Hero1Icon string `bson:"hero1_icon" json:"hero1_icon"`
	Hero1Pre5LastHit uint32 `bson:"hero1_pre_5min_lasthit" json:"hero1_pre_5min_lasthit"`
	Hero1Pre10LastHit uint32 `bson:"hero1_pre_10min_lasthit" json:"hero1_pre_10min_lasthit"`
	Hero1Pre5Denis uint32 `bson:"hero1_pre_5min_Denis" json:"hero1_pre_5min_Denis"`
	Hero1Pre10Denis uint32 `bson:"hero1_pre_10min_Denis" json:"hero1_pre_10min_Denis"`
	Hero1Pre10MinLevelUpTime []uint32 `bson:"hero1_pre_10min_lvl" json:"hero1_pre_10min_lvl"`
	Hero1Pre10MinLevelUpTimeStr []string `bson:"hero1_pre_10min_lvl_str" json:"hero1_pre_10min_lvl_str"`
	Hero1Pre10minItem []*ItemDetails `json:"hero1_pre_10min_item" bson:"hero1_pre_10min_item"`
	Hero1ConsumableItem []*PlayerConsumable `json:"hero1_consumable" bson:"hero1_consumable"`
	Hero1AbilityUpgrades []*Ability `json:"hero1_ability_upgrades" bson:"hero1_ability_upgrades"`
	Hero1DeathCount int `bson:"hero1_death_count" json:"hero1_death_count"`
	//Hero1Assisant []*MidAssisant `bson:"hero1_assisants" json:"hero1_assisants"`
	Hero2PlayerId uint32 `bson:"hero2_player_id" json:"hero2_player_id"`
	Hero2Name string `bson:"hero2_name" json:"hero2_name"`
	Hero2Icon string `bson:"hero2_icon" json:"hero2_icon"`
	Hero2Pre5LastHit uint32 `bson:"hero2_pre_5min_lasthit" json:"hero2_pre_5min_lasthit"`
	Hero2Pre10LastHit uint32 `bson:"hero2_pre_10min_lasthit" json:"hero2_pre_10min_lasthit"`
	Hero2Pre5Denis uint32 `bson:"hero2_pre_5min_Denis" json:"hero2_pre_5min_Denis"`
	Hero2Pre10Denis uint32 `bson:"hero2_pre_10min_Denis" json:"hero2_pre_10min_Denis"`
	Hero2Pre10MinLevelUpTime []uint32 `bson:"hero2_pre_10min_lvl" json:"hero2_pre_10min_lvl"`
	Hero2Pre10MinLevelUpTimeStr []string `bson:"hero2_pre_10min_lvl_str" json:"hero2_pre_10min_lvl_str"`
	Hero2Pre10minItem []*ItemDetails `json:"hero2_pre_10min_item" bson:"hero2_pre_10min_item"`
	Hero2ConsumableItem []*PlayerConsumable `json:"hero2_consumable" bson:"hero2_consumable"`
	Hero2AbilityUpgrades []*Ability `json:"hero2_ability_upgrades" bson:"hero2_ability_upgrades"`
	Hero2DeathCount int `bson:"hero2_death_count" json:"hero2_death_count"`
	//Hero2Assisant []*MidAssisant `bson:"hero2_assisants" json:"hero2_assisants"`
}
type MidLineUp struct{
	Version string `bson:"version" json:"version"`
	MidHero1 string `bson:"mid_hero1" json:"mid_hero1"`
	MidHero2 string `bson:"mid_hero2" json:"mid_hero2"`
	Hero1PlayerId uint32 `bson:"hero1_player_id" json:"hero1_player_id"`
	Hero2PlayerId uint32 `bson:"hero2_player_id" json:"hero2_player_id"`
	AgainstDets []*MidAgainst `bson:"aginst_dets" json:"against_dets"`
}

type SideLineUp struct{
	Version string `bson:"version" json:"version"`
	SoftLaneTeam string `bson:"soft_lane_team" bson:"soft_lane_team"`
	HardLaneTeam string `bson:"hard_lane_team" bson:"hard_lane_team"`
	SoftLaneMainPositionPlayerId uint32 `bson:"soft_main_position_player_id" json:"soft_main_position_player_id"`
	HardLaneMainPositionPlayerId uint32 `bson:"hard_main_position_player_id" json:"hard_main_position_player_id"`
	LineUpHeroes string `bson:"line_up_heroes" json:"line_up_heroes"`
	LineUpDets []*SideLineUpDets `bson:"side_line_up_dets" json:"side_line_up_dets"`
}
type SideLineUpDets struct{
	MatchId uint64 `bson:"match_id" json:"match_id"`
	AgainstHeroes []*AgainstHero `bson:"against_heroes" json:"against_heroes"`
}
type AgainstHero struct{
	PlayerId uint32 `bson:"player_id" json:"player_id"`
	PlayerName string `bson:"player_name" json:"player_name"`
	HeroName string `bson:"hero_name" json:"hero_name"`
	HeroIcon string `bson:"hero_icon" json:"hero_icon"`
	Pre5LastHit uint32 `bson:"pre_5min_lasthit" json:"pre_5min_lasthit"`
	Pre10LastHit uint32 `bson:"pre_10min_lasthit" json:"pre_10min_lasthit"`
	Pre5Denis uint32 `bson:"pre_5min_Denis" json:"pre_5min_Denis"`
	Pre10Denis uint32 `bson:"pre_10min_Denis" json:"pre_10min_Denis"`
	Pre5MinGold *GoldType `bson:"5min_gold" json:"pre_5min_gold"`
	Pre10MinGold *GoldType `bson:"10min_gold" json:"pre_10min_gold"`
	Pre10MinLevelUpTime []uint32 `bson:"pre_10min_lvl" json:"pre_10min_lvl"`
	Pre10MinLevelUpTimeStr []string `bson:"pre_10min_lvl_str" json:"pre_10min_lvl_str"`
	Pre10minItem []*ItemDetails `json:"pre_10min_item" bson:"pre_10min_item"`
	ConsumableItem []*PlayerConsumable `json:"hero_consumable" bson:"hero_consumable"`
	AbilityUpgrades []*Ability `json:"hero_ability_upgrades" bson:"hero_ability_upgrades"`
	HeroDeathCount int `bson:"hero_death_count" json:"hero_death_count"`
}
/*type SideLineUp struct{
	Version string `bson:"version" json:"version"`
	Side1Heroes string `bson:"side1_heroes" json:"side1_heroes"`
	Side2Heroes string `bson:"side2_heroes" json:"side2_heroes"`
	AgainstDets []*AgainstInfo `bson:"aginst_dets" json:"against_dets"`
}
type AgainstHero struct{
	HeroPlayerId uint32 `bson:"hero_player_id" json:"hero_player_id"`
	HeroName string `bson:"hero_name" json:"hero_name"`
	HeroIcon string `bson:"hero_icon" json:"hero_icon"`
	HeroPre5LastHit uint32 `bson:"hero_pre_5min_lasthit" json:"hero_pre_5min_lasthit"`
	HeroPre10LastHit uint32 `bson:"hero_pre_10min_lasthit" json:"hero_pre_10min_lasthit"`
	HeroPre5Denis uint32 `bson:"hero_pre_5min_Denis" json:"hero_pre_5min_Denis"`
	HeroPre10Denis uint32 `bson:"hero_pre_10min_Denis" json:"hero_pre_10min_Denis"`
	HeroPre10MinLevelUpTime []uint32 `bson:"hero_pre_10min_lvl" json:"hero_pre_10min_lvl"`
	HeroPre10MinLevelUpTimeStr []string `bson:"hero_pre_10min_lvl_str" json:"hero_pre_10min_lvl_str"`
	HeroPre10minItem []*ItemDetails `json:"hero_pre_10min_item" bson:"hero_pre_10min_item"`
	HeroConsumableItem []*PlayerConsumable `json:"hero_consumable" bson:"hero_consumable"`
	HeroAbilityUpgrades []*Ability `json:"hero_ability_upgrades" bson:"hero_ability_upgrades"`
	HeroDeathCount int `bson:"hero_death_count" json:"hero_death_count"`
}


type AgainstInfo struct{
	MatchId uint64 `bson:"match_id" json:"match_id"`
	Team1PlayerId uint32 `bson:"team1_player_id" json:"team1_player_id"`
	Team2PlayerId uint32 `bson:"team2_player_id" json:"team2_player_id"`
	Team1PlayersName string `bson:"team1_players_name" json:"team1_players_name"`
	Team2PlayersName string `bson:"team2_players_name" json:"team2_players_name"`
	Team1DeathCounts int `bson:"team1_death_counts" json:"team1_death_counts"`
	Team2DeathCounts int `bson:"team2_death_counts" json:"team2_death_counts"`
	Team1AgainstKda *[]AgainstKda `bson:"team1_kda" json:"team1_kda"`
	Team2AgainstKda *[]AgainstKda `bson:"team2_kda" json:"team2_kda"`
	Hero1Data *AgainstHero `bson:"hero1_data" json:"hero1_data"`
	Hero2Data *AgainstHero `bson:"hero2_data" json:"hero2_data"`
	Hero3Data *AgainstHero `bson:"hero3_data" json:"hero3_data"`
	Hero4Data *AgainstHero `bson:"hero4_data" json:"hero4_data"`
	Hero5Data *AgainstHero `bson:"hero5_data" json:"hero5_data"`
	Hero6Data *AgainstHero `bson:"hero6_data" json:"hero6_data"`
}*/
/*type MidPlayerLineUp struct{
	
}

type MatchMidLineUp struct{
	Version string `bson:"version" json:"version"`
	MidHero1 string `bson:"mid_hero1" json:"mid_hero1"`
	MidHero2 string `bson:"mid_hero2" json:"mid_hero2"`
	AllMidPlayerLineUp []*MidPlayerLineUp
}*/

/*      游戏时间转换成字符串的函数                    */
func Durformat(durint int64) string {
	
	durstring := strconv.FormatInt(durint, 10) + "s" //将duration 由int64转换成字符串
	dur, _ := time.ParseDuration(durstring)          //将时间字符串转换成Time.duration类型
	dur2 := dur.String()
	//fmt.Printf("dur2:%v\n", dur2)
	if strings.Contains(dur2, "h") {
		dur2 = strings.Replace(dur2, "h", ":", -1)
	}
	if strings.Contains(dur2, "0m") {
		dur2 = strings.Replace(dur2, "0 m", "0:", -1)
	}
	if strings.Contains(dur2, "m") {
		dur2 = strings.Replace(dur2, "m", ":", -1)
	}
	if strings.Contains(dur2, "s") {
		dur2 = strings.Replace(dur2, "s", "", -1)
	}
	return dur2
}