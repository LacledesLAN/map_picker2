package models

type Games struct {
	Games []Game `json:"games"`
}

type Game struct {
	Id 			int  `json:"id"`
	Name 			string `json:"name"`
	Backgroud_image 	string `json:"background_image"`
	Logo			string `json:"logo"`
	Url			string `json:"url"`
	GameMapList		[]GameMapList `json:"gameMapList"`
	Columns			[]Columns `json:"columns"`
}

type GameMapList struct {
	Id            string `json:"id"`
	Image         string `json:"image"`
	Action        string `json:"action"`
	SelectedOrder string `json:"selectedOrder"`
	ColumnID      int `json:"columnID"`
}

type Columns struct {
	Id		int `json:"id"`
	Name		string `json:"name"`
	MaxSelection 	int `json:"maxSelection"`
	MaxBan		int `json:"maxBan"`
}


