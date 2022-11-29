package app

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func CurrPath() string {
	cmd := exec.Command("pwd")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
        
	return string(out)
}

func CurrDir() string {
        p := CurrPath()
        s := strings.Split(p, "")
        return s[len(s)-2]
}

func cerateFile(b string) error {

        file := os.NewFile(077 ,b)
        defer file.Close()

        return nil
}
