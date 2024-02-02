package techpalace

import (
	"fmt"
	"strings"
)

const (
	WelcomeText  = "Welcome to the Tech Palace, %s"
	BorderText   = "%s\n%s\n%s"
	BorderFiller = "*"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf(WelcomeText, strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return fmt.Sprintf(BorderText,
		strings.Repeat(BorderFiller, numStarsPerLine),
		welcomeMsg,
		strings.Repeat(BorderFiller, numStarsPerLine),
	)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(
		strings.ReplaceAll(oldMsg, BorderFiller, ""),
	)
}
