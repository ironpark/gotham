package git

import (
	"os"
	"os/exec"
	"path/filepath"
)

func NewRepo(user, name string) {
	cmd := exec.Command("git", "init", "--bare", name+".git")
	os.Mkdir("./repo/"+user, 0777)
	cmd.Dir = filepath.Join(util.WorkingDir(), "repo", user)
	cmd.Run()
}
