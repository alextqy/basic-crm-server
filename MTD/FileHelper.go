package mtd

import (
	mod "basic-crm-server/MOD"
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

type FileHelper struct{}

// 文件检查
func (f *FileHelper) FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

// 新建文件
func (f *FileHelper) FileMake(filePath string) (bool, string) {
	r, err := os.Create(filePath)

	defer func(r io.Closer) {
		if err := r.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(r)

	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// 文件删除
func (f *FileHelper) FileRemove(filePath string) (bool, string) {
	err := os.Remove(filePath)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// 文件重命名
func (f *FileHelper) FileRename(filePath, newName string) (bool, string) {
	err := os.Rename(filePath, newName)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// 文件信息
func (f *FileHelper) Filespec(filePath string) (bool, os.FileInfo) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, nil
	}
	return true, fileInfo
}

// 文件读取
func (f *FileHelper) FileRead(filePath string) (bool, string) {
	contentByte, readErr := os.ReadFile(filePath)
	if readErr != nil {
		return false, readErr.Error()
	}
	return true, string(contentByte)
}

// 文件分块读取(二进制)
// buffer 偏移量
// start 开始读取的位置
func (f *FileHelper) FileReadBlock(filePath string, buffer int, start int) (bool, string, []byte) {
	r, err := os.Open(filePath)
	if err != nil {
		return false, err.Error(), nil
	}
	defer r.Close()

	b := make([]byte, buffer)
	n, err := r.ReadAt(b, int64(start))
	if err != nil && err != io.EOF {
		return false, err.Error(), nil
	}
	return true, "", b[:n]
}

// 文件写入
func (f *FileHelper) FileWrite(filePath, content string) (bool, error) {
	r, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)

	defer func(r io.Closer) {
		if err := r.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(r)

	if err != nil {
		return false, err
	} else {
		_, writeErr := r.Write([]byte(content))
		if writeErr != nil {
			return false, err
		}
		return true, nil
	}
}

// 文件写入追加
func (f *FileHelper) FileWriteAppend(filePath, content string) (bool, string) {
	r, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)

	defer func(r io.Closer) {
		if err := r.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(r)

	if err != nil {
		return false, err.Error()
	} else {
		write := bufio.NewWriter(r)
		write.WriteString(content)
		write.Flush()
		return true, ""
	}
}

// 文件二进制写入
func (f *FileHelper) FileWriteByte(filePath string, content []byte) (bool, string) {
	r, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	defer func(r io.Closer) {
		if err := r.Close(); err != nil {
			fmt.Printf("defer close file err: %v", err.Error())
		}
	}(r)
	if err != nil {
		return false, err.Error()
	}

	var bytesBuffer bytes.Buffer
	binary.Write(&bytesBuffer, binary.LittleEndian, content)
	_, err = r.Write(bytesBuffer.Bytes())
	if err != nil {
		return false, err.Error()
	}

	return true, ""
}

// 新建文件夹
func (f *FileHelper) DirMake(dirPath string) (bool, string) {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return false, err.Error()
	} else {
		return true, ""
	}
}

// 文件夹信息
func (f *FileHelper) DirCheck(dirPath string) (bool, string, []fs.DirEntry) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err.Error(), nil
	}
	return true, "", files
}

// 文件删除
func (f *FileHelper) DirDel(dirPath string) (bool, string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}

// 遍历文件夹
func (f *FileHelper) DirTraverse(dirPath string) (bool, string, []string, []string) {
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err.Error(), nil, nil
	}

	var dirs []string
	var files []string
	PathSep := string(os.PathSeparator)
	for _, obj := range dir {
		if obj.IsDir() {
			dirs = append(dirs, dirPath+PathSep+obj.Name())
		} else {
			files = append(files, dirPath+PathSep+obj.Name())
		}
	}

	return true, "", dirs, files
}

func (f *FileHelper) LogDir() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	return "./Log/" + strings.Split(t, " ")[0] + "/"
}

func (f *FileHelper) WriteLog(fileName, content, directory string) (bool, string) {
	if !f.FileExist(f.LogDir()) {
		b, s := f.DirMake(f.LogDir())
		if !b {
			return false, s
		}
	}
	dir := ""
	if directory != "" {
		dir = "/" + directory + "/"
		b, e := f.DirMake(f.LogDir() + dir)
		if !b {
			return b, e
		}
	}
	logFile := f.LogDir() + dir + fileName + ".log"
	if !f.FileExist(logFile) {
		b, s := f.FileMake(logFile)
		if !b {
			return false, s
		}
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	b, c := f.FileWriteAppend(logFile, t+" "+content+""+"\n")
	if !b {
		return false, c
	}
	return true, ""
}

func (f *FileHelper) CheckConf() mod.Conf {
	var conf mod.Conf
	if !f.FileExist("./Conf.json") {
		log.Panic("The system configuration failed")
		os.Exit(0)
	}
	_, byteData := f.FileRead("./Conf.json")
	json.Unmarshal([]byte(byteData), &conf)
	return conf
}
