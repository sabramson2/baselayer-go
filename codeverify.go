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
	code := "0"
	fileContents := fmt.Sprintf(template, expireSecs, code)

	e := WriteFile(path, fileContents)
	if e != nil { return "", e }
	return code, nil
}

//----------------------------------------
func CodeVerify(codeAttempt string) (bool, error) {
	json, e := ReadFileToJson(filePath())
	if e != nil { return false, e }

	timeoutSecs := json["expireSecs"].(int64)
	nowSecs := time.Now().Unix()
	if timeoutSecs > nowSecs {
		return false, nil
	}
	code := json["code"].(string)
	if code != codeAttempt {
		return false, nil
	}

	return true, nil
}

func filePath() string {
	return filepath.Join(os.TempDir(), codeFile)
}