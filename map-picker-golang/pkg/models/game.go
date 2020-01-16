package models

type Games struct {
	Games []Game `json:"games"`
}

type Game struct {
	Id 					int
	Name 				string
	Backgroud_image 	string
	Logo				string
	Url					string
	GameMapList		[]GameMapList
	Columns				[]Columns
}

type GameMapList struct {
	Id            string
	Image         string
	Action        string
	SelectedOrder string
	ColumnID      int
}

type Columns struct {
	Id				int
	Name			string
	MaxSelection 	int
	MaxBan			int
}


