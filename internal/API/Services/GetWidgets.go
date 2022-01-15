package Services

import (
	"encoding/json"
	"fmt"
	"main/internal/Structures"
	"main/pkg"
)

func GetWidgets() []byte {
	files := pkg.OnlyReadDir("/var/www/investments-cryptanalysis-parsing/assets/img")
	screenshots := &Structures.WidgetScreenshots{
		FilesUrl: files.Files,
	}
	b, err := json.Marshal(screenshots)
	if err != nil {
		fmt.Println(err)
	}
	return b

}
