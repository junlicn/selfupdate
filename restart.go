package selfupdate

import (
	"os"
	"syscall"
	"path/filepath"
	"log"

	//"github.com/fynelabs/selfupdate/internal/osext"
)

// Restart will attempt to restar the current application, any error will be returned.
// If the exiter function is passed in it will be responsible for terminating the old processes.
// If exiter is passed an error it can assume the restart failed and handle appropriately.
func Restart(exiter func(error),targetpath string) error {
	/*
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	exe, err := osext.Executable()
	if err != nil {
		return err
	}
	*/
	//osext.Executable()
	wd := filepath.Dir(targetpath)
	exe := targetpath

	_, err := os.StartProcess(exe, os.Args, &os.ProcAttr{
		Dir:   wd,
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   &syscall.SysProcAttr{},
	})

	log.Println("restart:",exe,wd)
	if exiter != nil {
		exiter(err)
	} else if err == nil {
		os.Exit(0)
	}
	return err
}
