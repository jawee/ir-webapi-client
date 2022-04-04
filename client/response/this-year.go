package response

type ThisYear struct {
	NumOfficialSessions int `json:"num_official_sessions"`
	NumLeagueSessions   int `json:"num_league_sessions"`
	NumOfficialWins     int `json:"num_official_wins"`
	NumLeagueWins       int `json:"num_league_wins"`
}
