package akiya

type Akiya struct {
	ID         uint `gorm:"primaryKey"`
	Title      string
	Link       string
	Price      string
	Desc       string
	Area       string
	Type       string
	Location   string
	LocationID string
}
