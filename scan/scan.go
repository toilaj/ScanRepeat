package scan

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"sync"
)

const (
	workerCount = 10
)

var scanData *ScanData

type ScanData struct {
	Path     string
	FileList []string
	Md5Map   map[string]string
	mutex    sync.Mutex
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

func init() {
	if pathExists("./output.txt") {
		os.Remove("./output.txt")
	}
	if pathExists("repeat.txt") {
		os.Remove("./repeat.txt")
	}
	scanData = new(ScanData)
	scanData.Md5Map = make(map[string]string)
}

func SetPath(path string) {
	scanData.Path = path
}

func getFileList(path string, fileList *[]string) {
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			getFileList(path+"/"+file.Name(), fileList) // 递归调用
		} else {
			*fileList = append(*fileList, path+"/"+file.Name())
		}
	}
}

func calcHash(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		fmt.Println(err, path)
		return "", err
	}
	h := md5.New()
	_, err = io.Copy(h, f)
	re := h.Sum(nil)
	md5str := fmt.Sprintf("%x", re)
	return md5str, nil
}

func compareInternal(fr *os.File, fo *os.File, filelist []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, path := range filelist {
		md5str, err := calcHash(path)
		if err != nil {
			continue
		}
		scanData.mutex.Lock()
		fo.WriteString(path + "\n")
		if v, ok := scanData.Md5Map[md5str]; ok {
			outStr := fmt.Sprintf("origin: %s\n", v)
			outStr = fmt.Sprintf("%srepeat: %s\n\n", outStr, path)
			fr.WriteString(outStr)
		} else {
			scanData.Md5Map[md5str] = path
		}
		scanData.mutex.Unlock()
	}
}

func compare() {
	fr, err := os.OpenFile("./repeat.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer fr.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fo, err := os.OpenFile("./output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer fo.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var countPerworker int
	if len(scanData.FileList) < workerCount {
		countPerworker = len(scanData.FileList)
	} else {
		countPerworker = len(scanData.FileList) / workerCount
	}
	var inputList []string = make([]string, 0)
	wg := sync.WaitGroup{}
	var batch int
	for index, path := range scanData.FileList {
		inputList = append(inputList, path)
		if index-batch+1 == countPerworker {
			wg.Add(1)
			go compareInternal(fr, fo, inputList, &wg)
			inputList = make([]string, 0)
			batch += countPerworker
		}
	}
	wg.Wait()

}

func Run() {
	getFileList(scanData.Path, &scanData.FileList)
	compare()
}
