package domain

import "time"

type App struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DNS       *DNS
	Port      int
	Status    string
}

type DNS struct {
	ARecord      string
	CNAMERecords []string
}
