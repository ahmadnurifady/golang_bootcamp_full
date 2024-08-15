package response

type Event struct {
	Id        string   `json:"id" valo:"notblank,sizeMin=2,sizeMax=10"`
	EventName string   `json:"event_name" valo:"notblank,sizeMin=2,sizeMax=50"`
	Location  string   `json:"location" valo:"notblank,sizeMin=2,sizeMax=50"`
	Date      string   `json:"date" valo:"notblank,sizeMin=2,sizeMax=20"`
	Ticket    []Ticket `json:"ticket,omitempty" valo"valid,notnil"`
}

type Ticket struct {
	Id        string  `json:"id" valo:"notblank,sizeMin=5,sizeMax=50"`
	Type      string  `json:"type" valo:"notblank,sizeMin=5,sizeMax=50"`
	Price     float64 `json:"price" valo:"min=50,max=10000"`
	Stock     int     `json:"stock" valo:"min=10,max=10000"`
	EventName string  `json:"event_name" valo:"notblank,sizeMin=5,sizeMax=50"`
}

type BaseResponse struct {
	StatusCode int    `json:"statuscode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}
