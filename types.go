package GoClans

const (
	WarFrequencyAlways              = "ALWAYS"
	WarFrequencyUnknown             = "UNKNOWN"
	WarFrequencyMoreThanOncePerWeek = "MORE_THAN_ONCE_PER_WEEK"
	WarFrequencyOncePerWeek         = "ONCE_PER_WEEK"
	WarFrequencyLessThanOncePerWeek = "LESS_THAN_ONCE_PER_WEEK"
	WarFrequencyNever               = "NEVER"
	WarFrequencyAny                 = "ANY"

	LegueWarStateGroupNotFound = "GROUP_NOT_FOUND"
	LegueWarStateNotInWar      = "NOT_IN_WAR"
	LegueWarStatePreparation   = "PREPARATION"
	LegueWarStateWar           = "WAR"
	LegueWarStateEnded         = "ENDED"

	WarStateClanNotFound  = "CLAN_NOT_FOUND"
	WarStateAccessDenied  = "ACCESS_DENIED"
	WarStateNotInWar      = "NOT_IN_WAR"
	WarStateInMatchmaking = "IN_MATCHMAKING"
	WarStateEnterWar      = "ENTER_WAR"
	WarStateMatched       = "MATCHED"
	WarStatePreparation   = "PREPARATION"
	WarStateWar           = "WAR"
	WarStateInWar         = "IN_WAR"
	WarStateInEnded       = "ENDED"

	ClanWarResultLose = "LOSE"
	ClanWarResultWin  = "WIN"
	ClanWarResultTie  = "TIE"
)

type Clan struct {
	Tag       string    `json:"Tag"`
	ClanLevel int       `json:"clanLevel"`
	Name      string    `json:"name"`
	BadgeUrls BadgeUrls `json:"badgeUrls"`
}

type ClanList struct {
	Items []Clan `json:"items"`
}

type ClanVersusRankingList struct {
	Items []ClanVersusRanking `json:"items"`
}

type BadgeUrls struct {
	Small  string `json:"small"`
	Large  string `json:"large"`
	Medium string `json:"medium"`
}

type ClanVersusRanking struct {
	Tag              string    `json:"tag"`
	Name             string    `json:"name"`
	Location         Location  `json:"location"`
	BadgeUrls        BadgeUrls `json:"badgeUrls"`
	ClanLevel        int       `json:"clanLevel"`
	Members          int       `json:"members"`
	Rank             int       `json:"rank"`
	PreviousRank     int       `json:"previousRank"`
	ClanVersusPoints int       `json:"clanVersusPoints"`
}

type ClanWarLeagueGroup struct {
	Tag    string               `json:"tag"`
	State  string               `json:"state"`
	Season string               `json:"season"`
	Clans  []ClanWarLeagueClan  `json:"clans"`
	Rounds []ClanWarLeagueRound `json:"rounds"`
}

type ClanWarLeagueClanList struct {
	Items []ClanWarLeagueClan `json:"items"`
}
type LeagueSeason struct {
	Id string `json:"id"`
}

type LeagueSeasonList struct {
	Items []LeagueSeason `json:"items"`
}

type GetLeagueSeasonsRequest struct {
	LeagueId string
	Limit    int
	After    string
	Before   string
}

type ClanWarLeagueClan struct {
	Tag       string `json:"tag"`
	ClanLevel int    `json:"clanLevel"`
	Name      string `json:"name"`
}

type ClanWarLog struct {
	Items []ClanWarLogEntry `json:"items"`
}

type ClanMemberList struct {
	Items []ClanMember `json:"items"`
}

type ClanMember struct {
	League            League `json:"league"`
	Tag               string `json:"tag"`
	Name              string `json:"name"`
	Role              string `json:"role"`
	ExpLevel          int    `json:"expLevel"`
	ClanRank          int    `json:"clanRank"`
	PreviousClanRank  int    `json:"previousClanRank"`
	Donations         int    `json:"donations"`
	DonationsReceived int    `json:"donationsReceived"`
	Trophies          int    `json:"trophies"`
	VersusTrophies    int    `json:"versusTrophies"`
}

type ClanWarLogEntry struct {
	Clan             WarClan `json:"clan"`
	AttacksPerMember int     `json:"attacksPerMember"`
	TeamSize         int     `json:"teamSize"`
	Opponent         WarClan `json:"opponent"`
	EndTime          string  `json:"endTime"`
	Result           string  `json:"result"`
}

type ClanWar struct {
	Clan                 WarClan `json:"clan"`
	AttacksPerMember     int     `json:"attacksPerMember"`
	TeamSize             int     `json:"teamSize"`
	Opponent             WarClan `json:"opponent"`
	StartTime            string  `json:"startTime"`
	State                string  `json:"state"`
	EndTime              string  `json:"endTime"`
	PreparationStartTime string  `json:"preparationStartTime"`
}

type WarClan struct {
	DestructionPercentage float64         `json:"destructionPercentage"`
	Tag                   string          `json:"tag"`
	Name                  string          `json:"name"`
	BadgeUrls             BadgeUrls       `json:"badgeUrls"`
	ClanLevel             int             `json:"clanLevel"`
	Attacks               int             `json:"attacks"`
	Stars                 int             `json:"stars"`
	ExpEarned             int             `json:"expEarned"`
	Members               []ClanWarMember `json:"members"`
}
type ClanWarLeagueClanMember struct {
	Tag           string `json:"tag"`
	TownHallLevel int    `json:"townHallLevel"`
	Name          string `json:"name"`
}

type ClanWarMember struct {
	Tag                string          `json:"tag"`
	Name               string          `json:"name"`
	MapPosition        int             `json:"mapPosition"`
	TownhallLevel      int             `json:"townhallLevel"`
	OpponentAttacks    int             `json:"opponentAttacks"`
	BestOpponentAttack ClanWarAttack   `json:"bestOpponentAttack"`
	Attacks            []ClanWarAttack `json:"attacks"`
}

type ClanWarAttack struct {
	Order                 int    `json:"order"`
	AttackerTag           string `json:"attackerTag"`
	DefenderTag           string `json:"defenderTag"`
	Stars                 int    `json:"stars"`
	DestructionPercentage int    `json:"destructionPercentage"`
	Duration              int    `json:"duration"`
}

type ClanWarLeagueRoundList struct {
	Items []ClanWarLeagueRound `json:"items"`
}

type ClanWarLeagueRound struct {
	WarTags []string `json:"warTags"`
}

type PlayerRankingList struct {
	Items []PlayerRanking `json:"items"`
}
type PlayerVersusRankingList struct {
	Items []PlayerVersusRanking `json:"items"`
}

type PlayerVersusRanking struct {
	Clan             PlayerRankingClan `json:"clan"`
	VersusBattleWins int               `json:"versusBattleWins"`
	Tag              string            `json:"tag"`
	Name             string            `json:"name"`
	ExpLevel         int               `json:"expLevel"`
	Rank             int               `json:"rank"`
	PreviousRank     int               `json:"previousRank"`
	VersusTrophies   int               `json:"versusTrophies"`
}

type PlayerRanking struct {
	Clan         PlayerRankingClan `json:"clan"`
	League       League            `json:"league"`
	AttackWins   int               `json:"attackWins"`
	DefenseWins  int               `json:"defenseWins"`
	Tag          string            `json:"tag"`
	Name         string            `json:"name"`
	ExpLevel     int               `json:"expLevel"`
	Rank         int               `json:"rank"`
	PreviousRank int               `json:"previousRank"`
	Trophies     int               `json:"trophies"`
}

type GoldPassSeason struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type League struct {
	Name     string   `json:"name"`
	Id       int      `json:"id"`
	IconUrls IconUrls `json:"iconUrls"`
}
type LeagueList struct {
	Items []League `json:"items"`
}

type IconUrls struct {
	Small  string `json:"small"`
	Tiny   string `json:"tiny"`
	Medium string `json:"medium"`
}

type PlayerRankingClan struct {
	Tag       string    `json:"tag"`
	Name      string    `json:"name"`
	BadgeUrls BadgeUrls `json:"badgeUrls"`
}

type ClientError struct {
	Reason  string      `json:"reason"`
	Message string      `json:"message"`
	Type    string      `json:"type"`
	Detail  interface{} `json:"detail"`
}

type ClanRankingList struct {
	Items []ClanRanking `json:"items"`
}

type ClanRanking struct {
	ClanLevel    int       `json:"clanLevel"`
	ClanPoints   int       `json:"clanPoints"`
	Location     Location  `json:"location"`
	Members      int       `json:"members"`
	Tag          string    `json:"tag"`
	Name         string    `json:"name"`
	Rank         int       `json:"rank"`
	PreviousRank int       `json:"previousRank"`
	BadgeUrls    BadgeUrls `json:"badgeUrls"`
}

type Location struct {
	LocalizedName string `json:"localizedName"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	IsCountry     bool   `json:"isCountry"`
	CountryCode   string `json:"countryCode"`
}

type VerifyTokenResponse struct {
	Tag    string `json:"tag"`
	Token  string `json:"token"`
	Status string `json:"status"`
}

type VerifyTokenRequest struct {
	Token string `json:"token"`
}

type ClanWarLogRequest struct {
	ClanTag string
	Limit   int
	After   int
	Before  int
}

type ClanMemberRequest struct {
	ClanTag string
	Limit   int
	After   int
	Before  int
}

type SearchClansRequest struct {
	Name          string
	WarFrequency  string
	LocationId    int
	MinMembers    int
	MaxMembers    int
	MinClanPoints int
	MinClanLevel  int
	Limit         int
	After         string
	Before        string
	LabelIds      string
}

type Player struct {
	Role                 string                      `json:"role"`
	WarPreference        string                      `json:"warPreference"`
	AttackWins           int                         `json:"attackWins"`
	DefenseWins          int                         `json:"defenseWins"`
	TownHallLevel        int                         `json:"townHallLevel"`
	TownHallWeaponLevel  int                         `json:"townHallWeaponLevel"`
	Troops               []PlayerItemLevel           `json:"troops"`
	VersusBattleWins     int                         `json:"versusBattleWins"`
	LegendStatistics     PlayerLegendStatistics      `json:"legendStatistics"`
	Heroes               []PlayerItemLevel           `json:"heroes"`
	Spells               []PlayerItemLevel           `json:"spells"`
	Labels               []Label                     `json:"labels"`
	Tag                  string                      `json:"tag"`
	Name                 string                      `json:"name"`
	ExpLevel             int                         `json:"expLevel"`
	Trophies             int                         `json:"trophies"`
	BestTrophies         int                         `json:"bestTrophies"`
	Donations            int                         `json:"donations"`
	DonationsReceived    int                         `json:"donationsReceived"`
	BuilderHallLevel     int                         `json:"builderHallLevel"`
	VersusTrophies       int                         `json:"versusTrophies"`
	BestVersusTrophies   int                         `json:"bestVersusTrophies"`
	WarStars             int                         `json:"warStars"`
	Achievements         []PlayerAchievementProgress `json:"achievements"`
	VersusBattleWinCount int                         `json:"versusBattleWinCount"`
}

type WarLeague struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type WarLeagueList struct {
	Items []WarLeague `json:"items"`
}

type GetLeaguesSeasonRankingRequest struct {
	LeagueId string
	SeasonId string
	Limit    int
	After    string
	Before   string
}

type GetLeaguesRequest struct {
	Limit  int
	After  string
	Before string
}

type PlayerAchievementProgress struct {
	Stars          int    `json:"stars"`
	Value          int    `json:"value"`
	Name           string `json:"name"`
	Target         int    `json:"target"`
	Info           string `json:"info"`
	CompletionInfo string `json:"completionInfo"`
	Village        string `json:"village"`
}

type Label struct {
	Name     string   `json:"name"`
	Id       int      `json:"id"`
	IconUrls IconUrls `json:"iconUrls"`
}
type LabelList struct {
	Items []Label `json:"items"`
}

type PlayerItemLevel struct {
	Level              int    `json:"level"`
	Name               string `json:"name"`
	MaxLevel           int    `json:"maxLevel"`
	Village            string `json:"village"`
	SuperTroopIsActive bool   `json:"superTroopIsActive"`
}

type LegendLeagueTournamentSeasonResult struct {
	Trophies int    `json:"trophies"`
	Id       string `json:"id"`
	Rank     int    `json:"rank"`
}

type GetWarLeaguesRequest struct {
	Limit  int
	After  string
	Before string
}
type GetLocationRankingsRequest struct {
	LocationId string
	Limit      int
	After      string
	Before     string
}

type GetLocationsRequest struct {
	Limit  int
	After  string
	Before string
}

type PlayerLegendStatistics struct {
	PreviousVersusSeason LegendLeagueTournamentSeasonResult `json:"previousVersusSeason"`
	BestVersusSeason     LegendLeagueTournamentSeasonResult `json:"bestVersusSeason"`
	LegendTrophies       int                                `json:"legendTrophies"`
	BestSeason           LegendLeagueTournamentSeasonResult `json:"bestSeason"`
	CurrentSeason        LegendLeagueTournamentSeasonResult `json:"currentSeason"`
	PreviousSeason       LegendLeagueTournamentSeasonResult `json:"previousSeason"`
}
