package main

import (
	"log"

	browser "github.com/EDDYCJY/fake-useragent"
)

func Browser() {
	// 推荐使用
	random := browser.Random()
	log.Printf("Random: %s", random)

	chrome := browser.Chrome()
	log.Printf("Chrome: %s", chrome)

	internetExplorer := browser.InternetExplorer()
	log.Printf("IE: %s", internetExplorer)

	firefox := browser.Firefox()
	log.Printf("Firefox: %s", firefox)

	safari := browser.Safari()
	log.Printf("Safari: %s", safari)

	android := browser.Android()
	log.Printf("Android: %s", android)

	macOSX := browser.MacOSX()
	log.Printf("MacOSX: %s", macOSX)

	ios := browser.IOS()
	log.Printf("IOS: %s", ios)

	linux := browser.Linux()
	log.Printf("Linux: %s", linux)

	iphone := browser.IPhone()
	log.Printf("IPhone: %s", iphone)

	ipad := browser.IPad()
	log.Printf("IPad: %s", ipad)

	computer := browser.Computer()
	log.Printf("Computer: %s", computer)

	mobile := browser.Mobile()
	log.Printf("Mobile: %s", mobile)
}
