package models

type Service struct {
	ID          int64
	Name        string
	Description string
}

type Stat struct {
	UserID      int64
	ServiceID   int64
	Count       int64
	ServiceName string
	Price       int64
}
