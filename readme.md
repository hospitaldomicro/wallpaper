# wallpaper [![godoc](https://godoc.org/github.com/reujab/wallpaper?status.svg)](https://godoc.org/github.com/reujab/wallpaper)

Uma biblioteca Golang cross-platform (Linux, Windows, and macOS) para se obter e definir um papel de parede.

## Instalação

```sh
go get github.com/reujab/wallpaper
```

## Exemplo

```go
package main

import (
	"fmt"

	"github.com/reujab/wallpaper"
)

func main() {
	background, err := wallpaper.Get()
	check(err)
	fmt.Println("Current wallpaper:", background)

	err = wallpaper.SetFromFile("/usr/share/backgrounds/gnome/adwaita-day.jpg")
	check(err)

	err = wallpaper.SetFromURL("https://i.imgur.com/pIwrYeM.jpg")
	check(err)

	err = wallpaper.SetMode(wallpaper.Crop)
	check(err)
}

```

## Sistemas suportados

* Windows
* macOS
* GNOME
* KDE
* Cinnamon
* Unity
* Budgie
* XFCE
* LXDE
* MATE
* Deepin
* A maioria dos compositores Wayland (somente define, necessita de swaybg)
* i3 (somente define, necessita de feh)
