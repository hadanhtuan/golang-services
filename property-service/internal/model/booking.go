package model

import (
	"property-service/internal/model/enum"
	"time"

	orm "github.com/hadanhtuan/go-sdk/db/orm"
	"gorm.io/gorm"
)

type Booking struct {
	ID        string     `json:"id" gorm:"default:gen_random_uuid()"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`

	//foreign key
	PropertyId string    `json:"propertyId"  gorm:"column:property_id"`
	Property   *Property `json:"property,omitempty"`
	UserId     string    `json:"userId"  gorm:"column:user_id"`
	Username   string    `json:"username"  gorm:"column:username"`
	HostId    string    `json:"hostId"  gorm:"column:host_id"`
	HostName   string    `json:"hostName"  gorm:"column:host_name"`

	Status *enum.BookingStatusValue `json:"status,omitempty" gorm:"column:status"`

	CheckInDate  int64 `json:"checkInDate,omitempty" gorm:"column:checkin_date"`
	CheckoutDate int64 `json:"checkoutDate,omitempty" gorm:"column:checkout_date"`
	PaymentDate  int64 `json:"paymentDate,omitempty" gorm:"column:payment_date"`
	GuestNumber  int32 `json:"guestNumber,omitempty" gorm:"column:guest_number"`
	ChildNumber  int32 `json:"childNumber,omitempty" gorm:"column:child_number"`
	BabyNumber   int32 `json:"babyNumber,omitempty" gorm:"column:baby_number"`
	PetNumber    int32 `json:"petNumber,omitempty" gorm:"column:pet_number"`
	NightNumber  int32 `json:"nightNumber,omitempty" gorm:"column:nightNum"`

	NightPrice float64 `json:"nightPrice,omitempty" gorm:"column:night_price"`
	ServiceFee float64 `json:"serviceFee,omitempty" gorm:"column:service_fee"`
	TaxPercent float64 `json:"taxPercent,omitempty" gorm:"column:tax_percent"`

	TotalPriceBeforeTax float64 `json:"totalPriceBeforeTax,omitempty" gorm:"column:total_price_before_tax"`
	TotalPrice          float64 `json:"totalPrice,omitempty" gorm:"column:total_price"`
	TaxFee              float64 `json:"taxFee,omitempty" gorm:"column:tax_fee"`
}

func (Booking) TableName() string {
	return "booking"
}

var BookingDB = &orm.Instance{
	TableName: "booking",
	Model:     &Booking{},
}

func InitTableBooking(db *gorm.DB) {
	db.AutoMigrate(&Booking{})
	BookingDB.ApplyDatabase(db)
}
