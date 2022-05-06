package teams

func GET_TEAM_URL(seasonId string) string {
	return "http://core.espnuk.org/v2/sports/cricket/leagues/" + seasonId + "/teams"
}
