package main

import "fmt"

type Project struct {
	Name  string
	Stars int
}

func main() {
	projects := []Project{
		{
			Name:  "https://github.com/docker/compose",
			Stars: 27600,
		},
		{
			Name:  "https://github.com/docker/build-push-action",
			Stars: 3500,
		},
		{
			Name:  "https://github.com/docker/compose-cli",
			Stars: 950,
		},
		{
			Name:  "https://github.com/lightvector/KataGo",
			Stars: 2200,
		},
		{
			Name:  "https://github.com/xuyulong54/xuyulong",
			Stars: 51,
		},
		{
			Name:  "https://github.com/superdudnikteam/gavno",
			Stars: 4,
		},
		{
			Name:  "https://github.com/TheEnderOfficial/CSGO-ZalupaHook",
			Stars: 3,
		},
		{
			Name:  "https://github.com/coder2gwy/coder2gwy",
			Stars: 22600,
		},
		{
			Name:  "https://github.com/P4nda0s/CryptoFucker",
			Stars: 230,
		},
		{
			Name:  "https://github.com/0voice/from_coder_to_expert",
			Stars: 10700,
		},
		{
			Name:  "https://github.com/microsoft/calculator",
			Stars: 26600,
		},
		{
			Name:  "https://github.com/google/ExoPlayer",
			Stars: 20400,
		},
		{
			Name:  "https://github.com/CarGuo/GSYVideoPlayer",
			Stars: 18200,
		},
	}
	m1 := make(map[string]Project)
	for i := 0; i < 13; i++ {
		m1[projects[i].Name] = projects[i]
	}

	for _, value := range m1 {
		fmt.Println(value)
	}
	// в цикле запишите в map

	// в цикле пройдитесь по мапе и выведите значения в консоль
}
