package dotfiles

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
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
					if err = os.Rename(d.HomePath(), d.DotfilePath()); err == nil {
						err = d.Symlink(force)
					}
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
	if matcher, err := dotfileIgnore(); err != nil {
		println(err.Error())
		return err
	} else {
		return filepath.Walk(home+"/.local/share/dotfiles/", func(path string, info os.FileInfo, err error) error {
			if matcher.Match([]string{strings.TrimPrefix(path, DotfileDir()+"/")}, info.IsDir()) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if !info.IsDir() {
				walkfunc(Dotfile{Path: strings.TrimPrefix(path, home+"/.local/share/dotfiles/")})
			}
			return nil
		})
	}
}

func DotfileDir() string {
	return home + "/.local/share/dotfiles"
}

func dotfileIgnore() (gitignore.Matcher, error) {
	if patterns, err := readIgnoreFile(osfs.New(DotfileDir()+"/"), []string{}, "/.dotfileignore"); err != nil {
		return nil, err
	} else {
		return gitignore.NewMatcher(patterns), nil
	}
}

// readIgnoreFile reads a specific git ignore file. (soruce gitignore/dir.go)
func readIgnoreFile(fs billy.Filesystem, path []string, ignoreFile string) (ps []gitignore.Pattern, err error) {
	commentPrefix := "#"
	f, err := fs.Open(fs.Join(append(path, ignoreFile)...))
	if err == nil {
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := scanner.Text()
			if !strings.HasPrefix(s, commentPrefix) && len(strings.TrimSpace(s)) > 0 {
				ps = append(ps, gitignore.ParsePattern(s, path))
			}
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return
}
