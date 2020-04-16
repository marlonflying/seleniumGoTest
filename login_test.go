package login

import (
	"fmt"
	"seleniumGoTest/support"
	"time"

	"github.com/cucumber/godog"
	messages "github.com/cucumber/messages-go/v10"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func clickOnLogIn() (err error) {
	btnLogin, err := Driver.FindElement(selenium.ByID, "wp-submit")
	if err != nil {
		return
	}
	btnLogin.Click()

	return nil
}

func pageTitleShouldBe(title string) (err error) {
	pageTitle, err := Driver.Title()
	if err != nil {
		return
	}
	if pageTitle != title {
		return fmt.Errorf("Page title mismatch, expected: %v, obteined: %v", title, pageTitle)
	}

	return nil
}

func userClickOnLogOut() (err error) {

	element, _ := Driver.FindElement(selenium.ByXPATH, `//*[@id="wp-admin-bar-my-account"]/a`)
	element.MoveTo(0, 0)
	//	coordinates, _ := element.
	//	robotgo.Move(coordinates.X, coordinates.Y)

	linkLogOut, _ := Driver.FindElement(selenium.ByXPATH, `//*[@id="wp-admin-bar-logout"]/a`)
	for {
		isdisplayed, _ := linkLogOut.IsDisplayed()
		if isdisplayed {
			break
		} else {
			time.Sleep(time.Millisecond * 100)
		}
	}
	linkLogOut.Click()

	return nil
}

func userEntersEmailAsAndPasswordAs(email, pass string) (err error) {
	txtEmail, err := Driver.FindElement(selenium.ByID, "user_login")
	if err != nil {
		return
	}
	txtEmail.SendKeys(email)

	txtPass, err := Driver.FindElement(selenium.ByID, "user_pass")
	if err != nil {
		return
	}
	txtPass.SendKeys(pass)
	return nil
}

func userLaunchChromeBrowser() error {
	Driver = support.WDinit()
	return nil
}

func userOpensURL(url string) error {
	Driver.Get(url)
	return nil
}

func closeBrowser() error {
	Driver.Quit()
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^Click on Log In$`, clickOnLogIn)
	s.Step(`^Page Title should be "([^"]*)"$`, pageTitleShouldBe)
	s.Step(`^User click on Log Out$`, userClickOnLogOut)
	s.Step(`^User enters Email as "([^"]*)" and Password as "([^"]*)"$`, userEntersEmailAsAndPasswordAs)
	s.Step(`^User Launch Chrome browser$`, userLaunchChromeBrowser)
	s.Step(`^User opens URL "([^"]*)"$`, userOpensURL)
	s.Step(`^close browser$`, closeBrowser)

	s.BeforeScenario(func(arg1 *messages.Pickle) {
	})
}
