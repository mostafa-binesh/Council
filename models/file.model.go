package models

import (
	U "docker/utils"
	"time"
)

var FileTypes = map[string]uint16{
	"plan":        1,
	"certificate": 2,
	"attachment":  3,
}
var IntFileTypes = map[uint16]string{
	1: "plan",
	2: "certificate",
	3: "attachment",
}

type File struct {
	ID   uint   `gorm:"primaryKey"`
	Type uint16 `json:"type"` // like attachment, certificate and etc...
	Name string `json:"name"`
	// relations
	LawID     uint
	Law       Law        `json:"law"`
	CreatedAt *time.Time `json:"createdAt" gorm:"not null;default:now()"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type FileMinimal struct {
	ID uint `json:"id"`
	// Type string `json:"type"` // like attachment, certificate and etc...
	Type uint16 `json:"type"` // like attachment, certificate and etc...
	URL  string `json:"name"`
	// relations
	// CreatedAt *time.Time `json:"createdAt" gorm:"not null;default:now()"`
	// UpdatedAt *time.Time `json:"updatedAt" gorm:"not null;default:now()"`
}
type UploadFile struct {
	LawId uint   `json:"lawId" validate:"required"`
	Type  string `json:"type" validate:"required"`
}

// tips: if convert is array, no need to call be reference
// but if convert is about one object, it would be better if we call them as
// -- call by refrence
func FileToFileMinimal(files []File) []FileMinimal {
	var minimalFiles []FileMinimal
	for i := 0; i < len(files); i++ {
		minimalFile := files[i].ToFileMinimal()
		minimalFiles = append(minimalFiles, minimalFile)
	}
	return minimalFiles
}
func (file File) ToFileMinimal() FileMinimal {
	return FileMinimal{
		ID:   file.ID,
		Type: file.Type,
		URL:  U.BaseURL + "/public/uploads/" + file.Name,
		// CreatedAt: file.CreatedAt,
		// UpdatedAt: file.UpdatedAt,
	}
}
func (f File) isAttachment() bool {
	return f.Type == FileTypes["attachment"]
}
func (f File) isPlan() bool {
	return f.Type == FileTypes["plan"]
}
func (f File) isCertificate() bool {
	return f.Type == FileTypes["certificate"]
}
