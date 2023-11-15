package branding

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiel77/server-demo/internal/meta"
	"github.com/mbndr/figlet4go"
)

const (
	//CopyrightSign the copyright sign
	CopyrightSign = "\u00A9"
)

// GetASCIIClaim godoc
func GetASCIIClaim(claim string) string {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()

	renderStr, _ := ascii.RenderOpts(claim, options)
	return renderStr
}

// PrintClaim godoc
func PrintClaim(claim string) {
	fmt.Print(GetASCIIClaim(claim))
}

// PrintCompanyName godoc
func PrintCompanyName() {
	fmt.Print(GetASCIICompanyName())
}

// GetAppName godoc
func GetAppName() string {
	return fmt.Sprintf("---[ %s ]---", meta.GetAppName())
}

// GetIntro godoc
func GetIntro() string {

	var b strings.Builder
	b.WriteString(GetASCIICompanyName())
	b.WriteString(GetAppName())
	b.WriteString("\n\n")
	b.WriteString(GetCopyright())
	b.WriteString("\n\n")
	return b.String()
}

// GetASCIICompanyName godoc
func GetASCIICompanyName() string {
	return GetASCIIClaim(meta.GetAuthorCompany())
}

// GetCopyright godoc
func GetCopyright() string {
	year := time.Now().Format("2006")
	return fmt.Sprintf("%s%s %s", CopyrightSign, year, meta.GetAuthorCompany())
}
