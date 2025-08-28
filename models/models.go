package models

import "time"

type Client struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"size:100;not null"`
	Code      string    `gorm:"size:50;unique;not null"`
	CreatedBy string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Warehouse struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	ClientID  uint64    `gorm:"not null;index"`
	Name      string    `gorm:"size:100;not null"`
	Location  string    `gorm:"size:200"`
	IsActive  bool      `gorm:"not null;default:false"`
	CreatedBy string    `gorm:"size:100;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Client    Client    `gorm:"foreignKey:ClientID"`
}

type Unit struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Code        string    `gorm:"size:20;unique;not null"` // PCS, KG, L, M
	Name        string    `gorm:"size:50;unique;not null"` // Piece, Kilogram, Liter
	Description string    `gorm:"type:text"`
	CreatedBy   string    `gorm:"size:100;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type Item struct {
	ID          uint64      `gorm:"primaryKey;autoIncrement"`
	Number      string      `gorm:"size:100;default:'DRAFT'"`
	ClientID    uint64      `gorm:"not null;index"`
	Name        string      `gorm:"size:100;not null"`
	UnitID      uint64      `gorm:"not null"`
	UnitCode    string      `gorm:"size:20;not null"`
	UnitName    string      `gorm:"size:50;not null"`
	DocumentURL string      `gorm:"type:text"`
	Description string      `gorm:"type:text"`
	IsActive    bool        `gorm:"not null;default:false"`
	CreatedBy   string      `gorm:"size:100;not null"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	Client      Client      `gorm:"foreignKey:ClientID"`
	StockCards  []StockCard `gorm:"foreignKey:ItemID"`
}

type StockCard struct {
	ID              uint64    `gorm:"primaryKey;autoIncrement"`
	Number          string    `gorm:"size:100;default:'DRAFT'"`
	ClientID        uint64    `gorm:"not null;index"`
	WarehouseID     uint64    `gorm:"not null;index"`
	ItemID          uint64    `gorm:"not null;index"`
	ItemNumber      string    `gorm:"size:100;default:'DRAFT'"`
	ItemDocumentURL string    `gorm:"type:text"`
	ItemName        string    `gorm:"size:100;not null"`
	ItemDescription string    `gorm:"type:text"`
	ItemUnit        string    `gorm:"size:50;not null"`
	ItemPrice       float64   `gorm:"type:decimal(12,2)"`
	Quantity        float64   `gorm:"type:decimal(12,2);not null"`
	TotalPrice      float64   `gorm:"->;type:decimal(14,2)"` // generated column
	Type            string    `gorm:"type:enum('IN','OUT','LOSS');not null"`
	ReferenceNo     string    `gorm:"size:50"`
	Note            string    `gorm:"type:text"`
	Status          string    `gorm:"type:enum('DRAFT','POSTED','CANCELLED');default:'DRAFT'"`
	CreatedBy       string    `gorm:"size:100;not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	PostedBy        string    `gorm:"size:100"`
	PostedAt        *time.Time
	CancelledBy     string `gorm:"size:100"`
	CancelledAt     *time.Time
	Client          Client    `gorm:"foreignKey:ClientID"`
	Warehouse       Warehouse `gorm:"foreignKey:WarehouseID"`
	Item            Item      `gorm:"foreignKey:ItemID"`
}

type AuditLog struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement"`
	ClientID  uint64    `gorm:"not null;index"`    // link to client
	TableName string    `gorm:"size:100;not null"` // nama tabel asal (item, warehouse, stock_card)
	RecordID  uint64    `gorm:"not null"`          // ID record yang berubah
	Action    string    `gorm:"type:enum('INSERT','UPDATE','DELETE');not null"`
	OldData   string    `gorm:"type:json"`         // snapshot sebelum perubahan
	NewData   string    `gorm:"type:json"`         // snapshot setelah perubahan
	ChangedBy string    `gorm:"size:100;not null"` // user / system yang ubah
	ChangedAt time.Time `gorm:"autoCreateTime"`    // timestamp perubahan

	Client Client `gorm:"foreignKey:ClientID"` // relationship
}
