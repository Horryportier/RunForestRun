package app

var (
	bonds []bond
        output string


        //def
        storage string
)

type bond struct {
	script string
	path   string
}

func Start() error {

	return nil
}

//reads bonds and defult settings
func Init() error{
        return nil
}

// executes if there's an existing script bond to the path.
func Execute(path string, script string) error {
	return nil
}

// Create new script and bonds it to path
func CreateNew() error {
	return nil
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
