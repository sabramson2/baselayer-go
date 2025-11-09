package baselayergo

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const codeFile = "codeVerifyFile.json"

//----------------------------------------
func CodeCreate(timeoutSecs int64) (string, error) {
	path := filePath()
	template := `{
		"expireSecs": %d,
		"code": "%s"
	}`

	expireSecs := time.Now().Unix() + timeoutSecs
	code := RandNumString(5)
	fileContents := fmt.Sprintf(template, expireSecs, code)
	e := WriteFile(path, fileContents)
	if e != nil { return "", e }
	return code, nil
}

//----------------------------------------
func CodeVerify(codeAttempt string) (bool, error) {
	json, e := ReadFileToJson(filePath())
	if e != nil { return false, e }

	timeoutSecs := int64(json["expireSecs"].(float64))
	nowSecs := time.Now().Unix()
	if nowSecs > timeoutSecs {
		return false, nil
	}
	code := json["code"].(string)
	if code != codeAttempt {
		return false, nil
	}

	return true, nil
}

//----------------------------------------
func filePath() string {
	return filepath.Join(os.TempDir(), codeFile)
}