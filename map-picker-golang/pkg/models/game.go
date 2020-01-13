package models

type Game struct {
	Id 					string `json:"id"`
	Name 				string `json:"name"`
	BackgroundImage 	string `json:"background_image"`
	Logo				string `json:"logo"`
	URL					string `json:"url"`
	GameMapList			string `json:"game_map_list"`
}

type GameMapList struct {
	Id 					string `json:"id"`
	Image 				string `json:"image"`
	Action 				string `json:"action"`
	SelectedOrder		string `json:"selected_order"`
	ColumnId			string `json:"column_id"`
}


