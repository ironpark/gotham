package git

import (
	//	"fmt"
	"github.com/IronPark/gotham/module/util"
	"github.com/speedata/gogit"
	//	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	//	"strings"
	"log"
)

type Repository struct {
	user string
	name string
	path string
}

func NewRepo(user, name string) *Repository {
	path := filepath.Join(util.WorkingDir(), "repo", user, name+".git")
	return &Repository{
		user: user,
		name: name,
		path: path,
	}
}

func (repo *Repository) CreateRepo() {
	cmd := exec.Command("git", "init", "--bare", repo.name+".git")
	os.Mkdir("./repo/"+repo.user, 0777)
	cmd.Dir = filepath.Join(util.WorkingDir(), "repo", repo.user)
	cmd.Run()
}

func (repo *Repository) FileList() ([]string, error) {
	repository, err := gogit.OpenRepository(repo.path)
	if err != nil {
		log.Fatal(err)
	}
	ref, err := repository.LookupReference("HEAD")
	if err != nil {
		log.Fatal(err)
	}
	ci, err := repository.LookupCommit(ref.Oid)
	if err != nil {
		log.Fatal(err)
	}
	var list []string

	ci.Tree.Walk(func(dirname string, te *gogit.TreeEntry) int {
		log.Println(path.Join(dirname, te.Name))
		list = append(list, path.Join(dirname, te.Name))
		return 0
	})
	return list, err
}
