package player

type Player struct {
	Tag                   string               `json:"tag,omitempty"`
	Name                  string               `json:"name,omitempty"`
	ExpLevel              int64                `json:"expLevel,omitempty"`
	Trophies              int64                `json:"trophies,omitempty"`
	BestTrophies          int64                `json:"bestTrophies,omitempty"`
	WINS                  int64                `json:"wins,omitempty"`
	Losses                int64                `json:"losses,omitempty"`
	BattleCount           int64                `json:"battleCount,omitempty"`
	ThreeCrownWINS        int64                `json:"threeCrownWins,omitempty"`
	ChallengeCardsWon     int64                `json:"challengeCardsWon,omitempty"`
	ChallengeMaxWINS      int64                `json:"challengeMaxWins,omitempty"`
	TournamentCardsWon    int64                `json:"tournamentCardsWon,omitempty"`
	TournamentBattleCount int64                `json:"tournamentBattleCount,omitempty"`
	Role                  string               `json:"role,omitempty"`
	Donations             int64                `json:"donations,omitempty"`
	DonationsReceived     int64                `json:"donationsReceived,omitempty"`
	TotalDonations        int64                `json:"totalDonations,omitempty"`
	WarDayWINS            int64                `json:"warDayWins,omitempty"`
	ClanCardsCollected    int64                `json:"clanCardsCollected,omitempty"`
	Clan                  Clan                 `json:"clan,omitempty"`
	Arena                 Arena                `json:"arena,omitempty"`
	LeagueStatistics      LeagueStatistics     `json:"leagueStatistics,omitempty"`
	Badges                []Badge               `json:"badges,omitempty"`
	Achievements          []Achievement         `json:"achievements,omitempty"`
	Cards                 []Card                `json:"cards,omitempty"`
	CurrentDeck           []Card                `json:"currentDeck,omitempty"`
	CurrentFavouriteCard  CurrentFavouriteCard `json:"currentFavouriteCard,omitempty"`
	StarPoints            int64                `json:"starPoints,omitempty"`
	ExpPoints             int64                `json:"expPoints,omitempty"`
}

type Achievement struct {
	Name           string     `json:"name,omitempty"`
	Stars          int64      `json:"stars,omitempty"`
	Value          int64      `json:"value,omitempty"`
	Target         int64      `json:"target,omitempty"`
	Info           string     `json:"info,omitempty"`
	CompletionInfo interface{} `json:"completionInfo"`
}

type Arena struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Badge struct {
	Name     string `json:"name,omitempty"`
	Progress int64  `json:"progress,omitempty"`
	Level    int64  `json:"level,omitempty"`
	MaxLevel int64  `json:"maxLevel,omitempty"`
	Target   int64  `json:"target,omitempty"`
}

type Card struct {
	Name      string   `json:"name,omitempty"`
	ID        int64    `json:"id,omitempty"`
	Level     int64    `json:"level,omitempty"`
	MaxLevel  int64    `json:"maxLevel,omitempty"`
	Count     int64    `json:"count,omitempty"`
	IconUrls  IconUrls `json:"iconUrls,omitempty"`
	StarLevel int64    `json:"starLevel,omitempty"`
}

type IconUrls struct {
	Medium string `json:"medium,omitempty"`
}

type Clan struct {
	Tag     string `json:"tag,omitempty"`
	Name    string `json:"name,omitempty"`
	BadgeID int64  `json:"badgeId,omitempty"`
}

type CurrentFavouriteCard struct {
	Name     string   `json:"name,omitempty"`
	ID       int64    `json:"id,omitempty"`
	MaxLevel int64    `json:"maxLevel,omitempty"`
	IconUrls IconUrls `json:"iconUrls,omitempty"`
}

type LeagueStatistics struct {
	CurrentSeason  CurrentSeason  `json:"currentSeason,omitempty"`
	PreviousSeason PreviousSeason `json:"previousSeason,omitempty"`
	BestSeason     BestSeason     `json:"bestSeason,omitempty"`
}

type BestSeason struct {
	ID       string `json:"id,omitempty"`
	Trophies int64  `json:"trophies,omitempty"`
}

type CurrentSeason struct {
	Trophies     int64 `json:"trophies,omitempty"`
	BestTrophies int64 `json:"bestTrophies,omitempty"`
}

type PreviousSeason struct {
	ID           string `json:"id,omitempty"`
	Trophies     int64  `json:"trophies,omitempty"`
	BestTrophies int64  `json:"bestTrophies,omitempty"`
}

type Verification struct {
	Tag    string `json:"tag,omitempty"`
	Token  string `json:"token,omitempty"`
	Status string `json:"status,omitempty"`
}
