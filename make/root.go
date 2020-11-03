package make

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"syscall"

	"github.com/cbroglie/mustache"
	"github.com/quadrabee/maak/config"
)

func EnsureDotMake() (string, error) {
	project, err := config.Load()
	if err != nil {
		return "", err
	}
	dirpath := path.Join(project.RootPath(), ".make")
	err = os.Mkdir(dirpath, os.ModeDir|0744)

	if err == nil || os.IsExist(err) {
		return dirpath, nil
	}
	return "", err
}

func BuildComponents(cmpNames []string, force bool) error {
	var args []string
	for _, comp := range cmpNames {
		if (force) {
			args = append(args, comp+".clean")
		}
		args = append(args, comp+".build")
	}
	return Execute(args)
}

func Execute(args []string) error {
	project, err := config.Load()
	if err != nil {
		return err
	}
	binary, lookErr := exec.LookPath("make")
	if lookErr != nil {
		panic(lookErr)
	}
	arguments := append([]string{"make", "-f", ".make/maak.mk"}, args...)
	env := os.Environ()
	os.Chdir(project.RootPath())
	execErr := syscall.Exec(binary, arguments, env)
	if execErr != nil {
		return execErr
	}
	return nil
}

func Generate() error {
	dotmake, err := EnsureDotMake()
	if err != nil {
		return err
	}
	data, err := Asset("make/templates/Makefile.tpl")
	if err != nil {
		return err
	}
	tpl := string(data)

	project, err := config.Load()
	if err != nil {
		return err
	}

	content, err := mustache.Render(tpl, &project)
	fpath := path.Join(dotmake, "maak.mk")
	err = ioutil.WriteFile(fpath, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
