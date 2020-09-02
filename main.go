package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sqweek/dialog"
	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Labtech\Service`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	i, _, err := k.GetIntegerValue("ID")
	if err != nil {
		log.Fatal(err)
	}

	s, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	result := fmt.Sprintf("Your Computer ID is %d\nYour Hostname is %s", i, s)
	dialog.Message("%s", result).Title("Your Remote Support Information").Info()

}
