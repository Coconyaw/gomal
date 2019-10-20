package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets6561c035e3a6e741e934b73bc2b818ffbc855690 = "This is a heads up text for advanced perspective attack with email.\nPlease do not open suspicious attached file.\nThank you.\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"head_up.txt"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1571458483, 1571458483973289498),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1571460888, 1571460888992516846),
		Data:     nil,
	}, "/assets/head_up.txt": &assets.File{
		Path:     "/assets/head_up.txt",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1571460888, 1571460888952516422),
		Data:     []byte(_Assets6561c035e3a6e741e934b73bc2b818ffbc855690),
	}}, "")

// URL : Send reqeust with UserInfo.
// For testing. requestbin url.
const URL = "https://enapt4e6utnzq.x.pipedream.net"

// UserInfo : Wanted information of mail training.
type UserInfo struct {
	Hostname string
	Username string
	Time     time.Time
}

func copyHeadUpFile(copyto string) {
	// Read assets file data for copying.
	file, err := Assets.Open("/assets/head_up.txt")
	data, _ := ioutil.ReadAll(file)
	file.Close()

	// Copy file
	err = ioutil.WriteFile(copyto, data, 0655)
	if err != nil {
		panic(err)
	}
}

func openHeadUpFile() {
	// Prepare open headup txt.
	// To begin with, copy embeded file to local user's home directory.
	// Then, open copied file with cmd.

	// copy headup file user's home directory
	home, _ := os.UserHomeDir()
	copyto := filepath.Join(home, "head_up.txt")

	// Open copied file with os command.
	var cmd *exec.Cmd
	switch goos := runtime.GOOS; goos {
	case "windows":
		cmd = exec.Command("start", copyto)
	case "darwin":
		cmd = exec.Command("open", copyto)
	case "linux":
		cmd = exec.Command("xdg-open", copyto)
	}

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
}

func sendUserInfo(ui *UserInfo) (*http.Response, error) {
	data := url.Values{"username": {ui.Username}, "hostname": {ui.Hostname}, "time": {ui.Time.Format(time.RFC3339)}}
	resp, err := http.PostForm(URL, data)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getUserInfo() (*UserInfo, error) {
	ui := &UserInfo{}
	hn, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	ui.Hostname = hn

	un, err := user.Current()
	if err != nil {
		return nil, err
	}
	ui.Username = un.Username
	ui.Time = time.Now()

	return ui, nil
}

func main() {
	openHeadUpFile()
	ui, err := getUserInfo()
	if err != nil {
		panic(err)
	}

	_, err = sendUserInfo(ui)
	if err != nil {
		panic(err)
	}
}
