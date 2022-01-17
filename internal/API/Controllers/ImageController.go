package Controllers

import (
	"main/internal/API/Services"
	"main/internal/Structures"
)

func GetWidgets() *Structures.WidgetScreenshots {
	return Services.GetWidgets()
}

func UpdWidgets() string {
	return Services.UpdWidgets()
}
