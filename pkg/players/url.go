package players

func GET_ATHLETES_URL(seasonId string, teamId string) string {
	return "http://core.espnuk.org/v2/sports/cricket/leagues/" + seasonId + "/teams/" + teamId + "/athletes?limit=50"
}
