package loader

import (
	//_ use native embed
	_ "embed"
	"encoding/json"
	"fmt"

	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	defaultError "github.com/cthiel77/server-demo/internal/error"
)

var (
	flagConfigFileName string
	//go:embed usage.txt
	flagConfigFileNameUsage string
)

func init() {
	flag.StringVarP(&flagConfigFileName, "config-file", "f", "", flagConfigFileNameUsage)

	// manually set time zone
	var tz string
	if tz = os.Getenv("TZ"); tz == "" {
		tz = "UTC"
	}
	t, err := time.LoadLocation(tz)
	if err != nil {
		panic(defaultError.NewError("1699199699", "500", fmt.Sprintf("error loading location %q: %+v\n", tz, err)))
	} else {
		time.Local = t
	}

}

// LoadConfigDataFile loads file by filename as fallback to flag
func LoadConfigDataFile(filename *string) (s *[]byte, successMsg string, err *defaultError.Error) {

	successMsg = ""
	if len(flagConfigFileName) == 0 && filename != nil || len(*filename) > 0 {
		flagConfigFileName = *filename
	}

	if len(flagConfigFileName) == 0 {
		err = defaultError.NewError("1699199710", defaultError.StatusConfigFileNameNotSet, "no config file given")
		return
	}

	s, loadErr := LoadFile(flagConfigFileName)
	if loadErr != nil {
		err = loadErr
		return
	}

	if !json.Valid(*s) {
		err = defaultError.NewError("1699199714", defaultError.StatusDataJSONError, "file content is not JSON")
		return
	}

	successMsg = fmt.Sprintf("loaded config data %q", flagConfigFileName)
	return
}

// LoadConfigData loads file by flag
func LoadConfigData() (*[]byte, *defaultError.Error) {

	if len(flagConfigFileName) == 0 {
		return nil, defaultError.NewError("1699199718", defaultError.StatusConfigFileNameNotSet, "no config file given")
	}

	s, loadErr := LoadFile(flagConfigFileName)

	if loadErr != nil {
		return nil, loadErr
	}

	if !json.Valid(*s) {
		return nil, defaultError.NewError("1699199722", defaultError.StatusDataJSONError, "file content is not JSON")
	}

	return s, nil
}

// LoadFile loads a file and returns the byte load
func LoadFile(path string) (*[]byte, *defaultError.Error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, defaultError.FromGoError("1699199726", defaultError.StatusFileReadError, err)
	}
	return &bytes, nil
}

// PrintConfig prints the config
func PrintConfig(config interface{}) {

	fmt.Printf("Local time zone config: \n")
	z, o := time.Now().Zone()
	fmt.Printf("timezone: %s, offset: %d\n", z, o)
	fmt.Printf("time: %s\n", time.Now().Format("2006-01-02T15:04:05.000 MST"))
	fmt.Printf("\n")
	if e := walkConfig(config, 0); e != nil {
		fmt.Printf("[ERRO] %+v", e)
	}
}

func walkConfig(config interface{}, level int) (err *defaultError.Error) {
	var inInterface map[string]interface{}
	inrec, e := json.Marshal(config)
	if e != nil {
		err = defaultError.FromGoError("1699194210", defaultError.StatusDataJSONEncodeError, e)
		return
	}

	if e := json.Unmarshal(inrec, &inInterface); e != nil {
		err = defaultError.FromGoError("1699194741", defaultError.StatusDataJSONDecodeError, e)
		return
	}
	keys := make([]string, 0)
	for field := range inInterface {
		keys = append(keys, field)
	}
	indent := indent(level, "  ")

	sort.Strings(keys)
	for _, k := range keys {
		val := inInterface[k]
		switch val.(type) {
		case map[string]interface{}:
			fmt.Printf("%s[%s]:\n", indent, k)
			if e := walkConfig(val, (level + 1)); e != nil {
				err = defaultError.FromGoError("1699194818", defaultError.StatusNotAcceptableData, e)
				return
			}
		default:
			checkVal := fmt.Sprintf("%v", val)
			if len(checkVal) == 0 {
				checkVal = "####"
			}
			if regexp.MustCompile(`(?i)(pass|secret)`).MatchString(k) {
				val = checkVal[:1] + "*****"
			}
			fmt.Printf("%s%s: %+v\n", indent, k, val)

		}

	}

	return
}

func indent(level int, char string) string {
	indent := make([]string, 0)
	for i := 0; i < level; i++ {
		indent = append(indent, char)
	}
	return strings.Join(indent, "")
}
