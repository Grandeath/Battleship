package connection

type StartingStruct struct {
	Desc        string `json:"desc"`
	Nick        string `json:"nick"`
	Target_nick string `json:"target_nick"`
	Wpbot       bool   `json:"wpBot"`
}

type PlayerListStruct struct {
	PlayerStruct []struct {
		Game_status string `json:"game_status"`
		Nick        string `json:"nick"`
	}
}
