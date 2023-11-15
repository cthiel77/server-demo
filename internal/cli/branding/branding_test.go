package branding

import (
	"fmt"
	"testing"
)

func TestPrintClaim(t *testing.T) {
	PrintClaim("hello world!")
}

func TestPrintCompanyName(t *testing.T) {
	PrintCompanyName()
}

func TestPrintIntro(t *testing.T) {
	fmt.Print(GetIntro())
}
