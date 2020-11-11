package dotfiles

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

var home string

func init() {
	var err error
	if home, err = homedir.Dir(); err != nil {
		panic("cannot determine user home")
	}
}

type Dotfile struct {
	Path string
}

func (d *Dotfile) HomePath() string {
	return fmt.Sprintf("%v/%v", home, d.Path)
}

func (d *Dotfile) DotfilePath() string {
	return fmt.Sprintf("%v/%v", DotfileDir(), d.Path)
}

func (d *Dotfile) Add(force bool) (err error) {
	if !d.Verify() {
		var stat os.FileInfo
		if _, err = os.Stat(d.DotfilePath()); os.IsNotExist(err) || force {
			if stat, err = os.Stat(d.HomePath()); err == nil {
				if stat.IsDir() {
					err = errors.New(fmt.Sprintf("%v is a directory", d.HomePath()))
				} else {
					_ = os.MkdirAll(filepath.Dir(d.DotfilePath()), os.ModePerm)
					err = os.Rename(d.HomePath(), d.DotfilePath())
				}
			}
		}
	}
	return
}

func (d *Dotfile) Symlink(force bool) (err error) {
	_ = os.MkdirAll(filepath.Dir(d.HomePath()), os.ModePerm)
	err = os.Symlink(d.DotfilePath(), d.HomePath())
	return
}

func (d *Dotfile) Verify() (ok bool) {
	if stat, err := os.Lstat(d.HomePath()); err == nil {
		if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
			if destination, err := os.Readlink(d.HomePath()); err == nil {
				ok = destination == d.DotfilePath()
			}
		}
	}
	return
}

func Walk(walkfunc func(dotfile Dotfile) error) error {
	return filepath.Walk(home+"/.local/share/dotfiles/", func(path string, info os.FileInfo, err error) error {
		if info.Name() == ".git" {
			return filepath.SkipDir
		}

		if !info.IsDir() {
			walkfunc(Dotfile{Path: strings.TrimPrefix(path, home+"/.local/share/dotfiles/")})
		}
		return nil
	})
}

func DotfileDir() string {
	return home + "/.local/share/dotfiles"
}
