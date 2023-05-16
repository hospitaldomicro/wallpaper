package main

import (
	"fmt"

	//"github.com/reujab/wallpaper"
	"github.com/hospitaldomicro/wallpaper"
)

func main() {
	background, err := wallpaper.Get()

	if err != nil {
		panic(err)
	}

	fmt.Println("Papel de parede atual:", background)

	err = wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	if err != nil {
		panic(err)
	}

	err = wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
	if err != nil {
		panic(err)
	}
}
