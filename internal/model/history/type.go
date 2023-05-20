package history

import "time"

type (
	AuthHistory struct {
		ID   string
		Role uint32
	}

	History struct {
		ID             string
		Uid            string
		TotalKolestrol string
		Tingkat        string
		ImageUrl       string
		CreatedAt      time.Time
		UpdatedAt      time.Time
		ID_2           string
		Name           string
		Email          string
		BirthDate      string
		Gender         string
		ImageUrl_2     string
		CreatedAt_2    time.Time
		UpdatedAt_2    time.Time
	}

	AddHistory struct {
		ID             string
		Uid            string
		TotalKolestrol string
		Tingkat        string
		ImageUrl       string
	}

	GetHistory struct {
		history []History
	}
)
