package logging

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"syscall"
	"time"
)

type DatetimeRotateWriter struct {
	lock        sync.Mutex
	filename    string
	fp          *os.File
	backupCount int
}

func New(filename string, backupCount int) (*DatetimeRotateWriter, error) {
	absFile, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalf("Invalid filename %s - %v", filename, err)
	}
	logPath := filepath.Dir(absFile)
	_, err = os.Stat(logPath)
	if err != nil {
		err = os.MkdirAll(logPath, 0777)
		if err != nil {
			return nil, err
		}
	}

	if _, err := IsWritable(logPath); err != nil {
		return nil, err
	}

	w := &DatetimeRotateWriter{filename: absFile,
		backupCount: backupCount}
	err = w.Rotate()
	if err != nil {
		return nil, err
	}
	w.Purge()
	return w, err
}

func IsWritable(path string) (bool, error) {

	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if !info.IsDir() {
		return false, fmt.Errorf("%s is not a folder", path)
	}

	// Check if the user bit is enabled in file permission
	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return false, fmt.Errorf("Write permission bit is not set on this file for user")
	}

	var stat syscall.Stat_t
	if err = syscall.Stat(path, &stat); err != nil {
		return false, fmt.Errorf("Unable to get stat")
	}

	if uint32(os.Geteuid()) != stat.Uid {
		return false, fmt.Errorf("User doesn't have permission to write to this directory")
	}

	return true, nil
}

func (w *DatetimeRotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.fp.Write(output)
}

func (w *DatetimeRotateWriter) Rotate() (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	ct, err := fileCreationTime(w.filename)
	if err == nil && ct.Before(today()) {
		if w.fp != nil {
			err = w.fp.Close()
			w.fp = nil
			if err != nil {
				return err
			}
		}
		ext := filepath.Ext(w.filename)
		basename := strings.TrimSuffix(filepath.Base(w.filename), ext)
		newFile := filepath.Join(filepath.Dir(w.filename), basename+"."+ct.Format("20060102")+ext)
		err = os.Rename(w.filename, newFile)
	}
	w.fp, err = os.OpenFile(w.filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	return err
}

func (w *DatetimeRotateWriter) Purge() (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	files, err := ioutil.ReadDir(filepath.Dir(w.filename))
	if err != nil {
		return err
	}
	filesList := []string{}
	ext := filepath.Ext(w.filename)
	basename := strings.TrimSuffix(filepath.Base(w.filename), ext)
	for _, f := range files {
		if !f.IsDir() && strings.HasPrefix(f.Name(), basename) && filepath.Base(w.filename) != f.Name() {
			filesList = append(filesList, f.Name())
		}
	}
	if len(filesList) < w.backupCount {
		return nil
	}
	sort.Strings(filesList)
	for n, f := range filesList {
		if n >= len(filesList)-w.backupCount {
			break
		}
		err = os.Remove(filepath.Join(filepath.Dir(w.filename), f))
		if err != nil {
			return err
		}

	}

	return nil
}

func today() time.Time {
	return time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)
}

func fileCreationTime(filename string) (time.Time, error) {
	fi, err := os.Stat(filename)
	if err != nil {
		return time.Unix(0, 0), err
	}
	stat := fi.Sys().(*syscall.Stat_t)
	ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	return ctime, nil
}
