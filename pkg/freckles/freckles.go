package freckles

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
	"github.com/go-git/go-git/v5/plumbing/format/gitignore"
)

var home string

func init() {
	var err error
	if home, err = os.UserHomeDir(); err != nil {
		panic("cannot determine user home")
	}
}

type Freckle struct {
	Path string
}

func (d *Freckle) HomePath() string {
	return fmt.Sprintf("%v/%v", home, d.Path)
}

func (d *Freckle) FrecklePath() string {
	return fmt.Sprintf("%v/%v", FreckleDir(), d.Path)
}

func (d *Freckle) Add(force bool) (err error) {
	if !d.Verify() {
		var stat os.FileInfo
		if _, err = os.Stat(d.FrecklePath()); os.IsNotExist(err) || force {
			if stat, err = os.Stat(d.HomePath()); err == nil {
				if stat.IsDir() {
					err = fmt.Errorf("%v is a directory", d.HomePath())
				} else {
					_ = os.MkdirAll(filepath.Dir(d.FrecklePath()), os.ModePerm)
					if err = os.Rename(d.HomePath(), d.FrecklePath()); err == nil {
						err = d.Symlink(force)
					}
				}
			}
		}
	}
	return
}

func (d *Freckle) Symlink(force bool) (err error) {
	_ = os.MkdirAll(filepath.Dir(d.HomePath()), os.ModePerm)
	err = os.Symlink(d.FrecklePath(), d.HomePath())
	return
}

func (d *Freckle) Verify() (ok bool) {
	if stat, err := os.Lstat(d.HomePath()); err == nil {
		if stat.Mode()&os.ModeSymlink == os.ModeSymlink {
			if destination, err := os.Readlink(d.HomePath()); err == nil {
				ok = destination == d.FrecklePath()
			}
		}
	}
	return
}

func Walk(walkfunc func(freckle Freckle) error) error {
	if matcher, err := frecklesIgnore(); err != nil {
		println(err.Error())
		return err
	} else {
		return filepath.Walk(home+"/.local/share/freckles/", func(path string, info os.FileInfo, err error) error {
			if matcher.Match([]string{strings.TrimPrefix(path, FreckleDir())}, info.IsDir()) {
				if info.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}

			if !info.IsDir() {
				walkfunc(Freckle{Path: strings.TrimPrefix(path, home+"/.local/share/freckles/")})
			}
			return nil
		})
	}
}

func FreckleDir() string {
	return home + "/.local/share/freckles/"
}

func frecklesIgnore() (gitignore.Matcher, error) {
	if patterns, err := readIgnoreFile(osfs.New(FreckleDir()), []string{}, "/.frecklesignore"); err != nil {
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
