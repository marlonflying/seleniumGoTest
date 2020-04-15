package support

import (
	"fmt"

	"github.com/tebeka/selenium"
)

var wd selenium.WebDriver

//WDinit returns a selenium WebDriver
func WDinit() selenium.WebDriver {

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	wd, err := selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Println("Error al inicializar el driver")
	}

	return wd
}
