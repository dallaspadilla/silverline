package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Ping struct {
	PingID       uint32    `gorm:"primary_key;auto_increment" json: pingID`
	TimeStamp    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
	PacketLoss   string    `gorm:"size:100;not null;" json: packetLoss`
	MinRoundTrip string    `gorm:"size:100;not null;" json: minRoundTrip`
	MaxRoundTrip string    `gorm:"size100;not null;" json: maxRoundTrip`
	AvgRoundTrip string    `gorm:"size100;not null;" json: avgRoundTrip`
}

func (u *Ping) Prepare() {
	u.PingID = 0
	u.TimeStamp = time.Now()
	u.PacketLoss = html.EscapeString(strings.TrimSpace(u.PacketLoss))
	u.MinRoundTrip = html.EscapeString(strings.TrimSpace(u.MinRoundTrip))
	u.MaxRoundTrip = html.EscapeString(strings.TrimSpace(u.MaxRoundTrip))
	u.AvgRoundTrip = html.EscapeString(strings.TrimSpace(u.AvgRoundTrip))
}

func (p *Ping) Validate() error {

	if p.PacketLoss == "" {
		return errors.New("Required PacketLoss")
	}
	if p.MinRoundTrip == "" {
		return errors.New("Required MinRoundTrip")
	}
	if p.MaxRoundTrip == "" {
		return errors.New("Required MaxRoundTrip")
	}
	if p.AvgRoundTrip == "" {
		return errors.New("Required AvgRoundTrip")
	}
	return nil
}

func (u *Ping) SaveUser(db *gorm.DB) (*Ping, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Ping{}, err
	}
	return u, nil
}

func (u *Ping) FindAllUsers(db *gorm.DB) (*[]Ping, error) {
	var err error
	pings := []Ping{}
	err = db.Debug().Model(&Ping{}).Limit(100).Find(&pings).Error
	if err != nil {
		return &[]Ping{}, err
	}
	return &pings, err
}

func (u *Ping) FindUserByID(db *gorm.DB, uid uint32) (*Ping, error) {
	var err error
	err = db.Debug().Model(Ping{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Ping{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Ping{}, errors.New("User Not Found")
	}
	return u, err
}
func (u *Ping) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Ping{}).Where("id = ?", uid).Take(&Ping{}).Delete(&Ping{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
