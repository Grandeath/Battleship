package connection

type statusStruct struct {
	Desc     string `json:"desc"`
	Nick     string `json:"nick"`
	Opp_desc string `json:"opp_desc"`
	Opponent string `json:"opponent"`
}
