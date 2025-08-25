package main

import "time"

type Client struct {
	ID         uint64      `gorm:"primaryKey;autoIncrement"`
	Name       string      `gorm:"size:100;not null"`
	Code       string      `gorm:"size:50;unique;not null"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
	Warehouses []Warehouse `gorm:"foreignKey:ClientID"`
	Items      []Item      `gorm:"foreignKey:ClientID"`
}

type Warehouse struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	ClientID  uint64    `gorm:"not null;index"`
	Name      string    `gorm:"size:100;not null"`
	Location  string    `gorm:"size:200"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Client    Client    `gorm:"foreignKey:ClientID"`
}

type Item struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	ClientID    uint64    `gorm:"not null;index"`
	Code        string    `gorm:"size:50;not null"`
	Name        string    `gorm:"size:100;not null"`
	Unit        string    `gorm:"size:20;not null"`
	DocumentURL string    `gorm:"type:text"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Client      Client    `gorm:"foreignKey:ClientID"`
}

type StockCard struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement"`
	ClientID        uint64    `gorm:"not null;index"`
	WarehouseID     uint64    `gorm:"not null;index"`
	ItemID          uint64    `gorm:"not null;index"`
	ItemDocumentURL string    `gorm:"type:text"`
	ItemName        string    `gorm:"size:100;not null"`
	ItemUnit        string    `gorm:"size:20;not null"`
	ItemDescription string    `gorm:"type:text"`
	Type            string    `gorm:"type:enum('IN','OUT','LOSS');not null"`
	ReferenceNo     string    `gorm:"size:50"`
	Quantity        float64   `gorm:"type:decimal(12,2);not null"`
	UnitPrice       float64   `gorm:"type:decimal(12,2)"`
	TotalPrice      float64   `gorm:"->;type:decimal(14,2)"` // read-only (generated in DB)
	Note            string    `gorm:"type:text"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	Client          Client    `gorm:"foreignKey:ClientID"`
	Warehouse       Warehouse `gorm:"foreignKey:WarehouseID"`
	Item            Item      `gorm:"foreignKey:ItemID"`
}
