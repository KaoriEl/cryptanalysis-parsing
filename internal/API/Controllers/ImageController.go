package Controllers

import "main/internal/API/Services"

func GetWidgets() []byte {
	return Services.GetWidgets()
}

func UpdWidgets() string {
	return Services.UpdWidgets()
}
