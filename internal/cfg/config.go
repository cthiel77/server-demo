package cfg

const (
	//ModeKeyDev mode key dev
	ModeKeyDev = "dev"
	//ModeKeyProd mode key prod
	ModeKeyProd = "prod"
)

var (
	//ModeDefault default mode is dev
	ModeDefault = ModeKeyDev
	//Mode mode
	Mode = ModeKeyDev
	//Code Application Code
	Code = "app"
)

// GetAppModeKey returns the current application mode key (pssible: dev, prod)
func GetAppModeKey() string {
	if Mode != ModeKeyDev {
		return ModeKeyProd
	}
	return ModeKeyDev
}

// DevModeEnabled returns true, if app is in dev mode
func DevModeEnabled() bool {
	return GetAppModeKey() == ModeKeyDev
}

// ProdModeEnabled returns true, if app is in prod mode
func ProdModeEnabled() bool {
	return GetAppModeKey() == ModeKeyProd
}
