package models

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// A photo can have multiple images and sidecar files
type Photo struct {
	Model
	PhotoUUID          string `gorm:"unique_index;"`
	PhotoTitle         string
	PhotoTitleChanged  bool
	PhotoDescription   string `gorm:"type:text;"`
	PhotoNotes         string `gorm:"type:text;"`
	PhotoArtist        string
	PhotoCanonicalName string
	PhotoFavorite      bool
	PhotoPrivate       bool
	PhotoSensitive     bool
	PhotoStory         bool
	PhotoLat           float64
	PhotoLong          float64
	PhotoAltitude      int
	PhotoFocalLength   int
	PhotoIso           int
	PhotoAperture      float64
	PhotoExposure      string
	PhotoViews         uint
	Camera             *Camera
	CameraID           uint
	Lens               *Lens
	LensID             uint
	Country            *Country
	CountryID          string
	CountryChanged     bool
	Location           *Location
	LocationID         uint
	LocationChanged    bool
	LocationEstimated  bool
	TakenAt            time.Time
	TakenAtChanged     bool
	TimeZone           string
	Labels             []*PhotoLabel
	Files              []*File
	Albums             []*Album `gorm:"many2many:album_photos;"`
}

func (m *Photo) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("PhotoUUID", uuid.NewV4().String())
}
