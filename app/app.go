package app

import (
	"fmt"
	"log"
	"os/user"
        "os"
)

var (
	bonds  []bond
	output string

	//def
	storage    string
	defStorage    string = "/home/%s/.runForrest/"


        fileName string
)

type bond struct {
	script string
	path   string
}

func Start() error {

	err := Init()
        err =CreateNew()
	return err
}

//reads bonds and defult settings
func Init() error {
	storage = os.Getenv("FORREST_PATH")

        User, err :=  user.Current()
        if err != nil {
                return err
        }
        if storage == "" {
                storage = fmt.Sprint(defStorage, User.Username)
        }
        
	return err
}

// executes if there's an existing script bond to the path.
func Execute(path string, script string) error {

	return nil
}

// Create new script and bonds it to path
func CreateNew() error {
        var b bond

        b.path = CurrPath()
        b.script = func()(string){
                if fileName == "" {
                        fileName = CurrDir()+".sh"
                }
                s := fmt.Sprintf("%s%s", storage, fileName)
                return s
        }()

        err := cerateFile(b.script)

        bonds = append(bonds, b)
	return err
}

// Copy script to other path
func Copy() error {
	return nil
}

// deletes bond/script or only one off them.
func Delete() error {
	return nil
}

// Opens tui view.
func Tui() error {
	return nil
}
