package akiya

type Akiya struct {
	ID         uint `gorm:"primaryKey:autoIncrement:not null"`
	Title      string
	Link       string
	Price      string
	Desc       string
	Area       string
	Type       string
	Location   string
	LocationID string
}

type Akiyas []Akiya
