package connection

type DescriptionStruct struct {
	Desc     string `json:"desc"`
	Nick     string `json:"nick"`
	Opp_desc string `json:"opp_desc"`
	Opponent string `json:"opponent"`
}

type StatusStruct struct {
	Game_status      string   `json:"game_status"`
	Should_fire      bool     `json:"should_fire"`
	Opp_shots        []string `json:"opp_shots"`
	Last_game_status string   `json:"last_game_status"`
}
