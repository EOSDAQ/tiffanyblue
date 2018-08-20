package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// SymbolInfo ...
type SymbolInfo struct {
	gorm.Model  `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Exchange    string `json:"exchange"`
	Type        string `json:"type"`
}

func (s *SymbolInfo) String() string {
	return fmt.Sprintf("Name[%s] Description[%s] Exchange[%s] Type[%s]", s.Name, s.Description, s.Exchange, s.Type)
}
