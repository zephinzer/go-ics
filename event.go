package ics

const ICalEventHeader = "BEGIN:VEVENT"
const ICalEventFooter = "END:VEVENT"

type Event struct {
	UID         string    `json:"uid"`
	ProductID   string    `json:"product_id"`
	CalScale    string    `json:"calscale"`
	Method      string    `json:"method"`
	Title       string    `json:"title"`
	Summary     string    `json:"summary"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Categories  []string  `json:"categories"`
	Status      string    `json:"status"`
	Organizer   Person    `json:"organizer"`
	Location    Location  `json:"geo"`
	Timestamp   Timestamp `json:"timestamp"`
	Sequence    int       `json:"sequence"`
}

func (event *Event) Serialize() string {
	return ICalEventHeader + `
		
	`
}
