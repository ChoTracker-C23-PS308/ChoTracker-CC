package history

import "time"

type (
	history struct {
		ID             string
		UID            string
		TotalKolestrol string
		Tingkat        string
		ImageUrl       string

		CreatedAt time.Time
		UpdatedAt time.Time
	}

	addHistory struct {
		ID             string
		UID            string
		TotalKolestrol string
		Tingkat        string
		ImageUrl       string
	}

	getHistory struct {
		history []history
	}
)
