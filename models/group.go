package models

//Group representa a estrutura de grupos do workplace
type Group struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Privacy string `json:"privacy"`
}

type GroupList struct {
	Data []struct {
		Name    string `json:"name"`
		Privacy string `json:"privacy"`
		ID      string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

//GroupMembers é a estrutura de membros de um grupo
type GroupMembers struct {
	Data []struct {
		Name          string `json:"name"`
		Administrator bool   `json:"administrator"`
		ID            string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}
