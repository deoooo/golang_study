package main

import (
	"fmt"
	"golang_study/balance/balance"
	"time"
)

func main() {
	var volumes []*balance.Volume
	for i := 0; i < 8; i++ {
		host := fmt.Sprintf("192.168.1.%d", i)
		volume := balance.NewVolume(host, 8080, i)
		volumes = append(volumes, volume)
	}

	var balancerName = "weights"
	for {
		v, err := balance.SelectByName(balancerName, volumes)
		if err != nil {
			fmt.Printf("select fail:%s\n", err)
			continue
		}

		fmt.Println(v)
		time.Sleep(time.Millisecond * 500)
	}

}
