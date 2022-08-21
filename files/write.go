package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func SaveFile(name, data string) {
	fname := fmt.Sprintf("data/%s", name)
	os.Remove(fname)
	ioutil.WriteFile(fname, []byte(data), 0644)
}

func RmRfBang() {
	exec.Command("rm", "-rf", "data").CombinedOutput()
	os.Mkdir("data", 0755)
}
