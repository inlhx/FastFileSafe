/**
*  文件压缩加密工具,用于Windows/Linux/MacOS等,加密后的文件尾部名称为:dooxb
*  使用方法分为两种
*  1.双击直接打开输入文件位置进行加密解密
*  2.Windows使用注册表方式放到右键进行快速加解密
*  3.传参方式,FastFileSafe /opt/info/myfiledooxb 进行加密或解密(自动根据尾部识别)
*  create by liuxiujun    2022-03-16 22:25
**/
package main

import (
	"archive/zip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unsafe"

	"github.com/atotto/clipboard"
)

var key []byte
var enfileType = "dooxb"

// 字节转string
func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// 获取文件sha1
func GetFileSHA1(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("sum error: ", err)
		return ""
	}

	b := sha1.Sum(data)
	return fmt.Sprintf("%X", b)
}

//批量文件压缩
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

//单个文件压缩
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//文件解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

// 获取目录位置
func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

//字符串截取
func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

// 随机生成指定长度的string
func rand_str(str_size int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, str_size)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

//创建解密key
func createPrivKey() []byte {
	newkey := []byte(rand_str(32))
	key = newkey
	return newkey
}

//加密文件
func encryptFile(inputfile string, outputfile string) {
	b, err := ioutil.ReadFile(inputfile) //Read the target file
	if err != nil {
		fmt.Printf("不能打开输入文件请检查!\n" + inputfile)
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
	ciphertext := encrypt(key, b)
	//fmt.Printf("%x\n", ciphertext)
	err = ioutil.WriteFile(outputfile, ciphertext, 0644)
	if err != nil {
		fmt.Printf("不能创建加密文件,请检查是否有创建文件权限!\n" + inputfile)
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
}

//解密文件
func decryptFile(inputfile string, outputfile string) {
	z, err := ioutil.ReadFile(inputfile)
	result := decrypt(key, z)
	// fmt.Printf("请检查是否有文件创建权限或权限为 0777\n")
	err = ioutil.WriteFile(outputfile, result, 0777)
	if err != nil {
		fmt.Printf("文件解密失败!\n")
		time.Sleep(3 * time.Second)
		os.Exit(0)
	}
}

//base64转码
func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

//base64解密
func decodeBase64(b []byte) []byte {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		fmt.Printf("Error: Bad Key!\n")
		time.Sleep(10 * time.Second)
		os.Exit(0)
	}
	return data
}

// 加密方法
func encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("您的密码格式错误")
		panic(err)
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Printf("您的密码格式错误")
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext
}

// 解密方法
func decrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("\r\n\r\n您的密码错误\r\n\r\n")
		time.Sleep(5 * time.Second)
		panic(err)
	}
	if len(text) < aes.BlockSize {
		// fmt.Printf("Error!\n")
		fmt.Printf("\r\n\r\n您的密码错误\r\n\r\n")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return decodeBase64(text)
}

// func barprocess() {
// 	bar := progressbar.Default(100)
// 	for i := 0; i < 100; i++ {
// 		bar.Add(1)
// 		time.Sleep(40 * time.Millisecond)
// 	}
// }
func runMain(filePath string, isDeFile bool) {
	// barprocess()
	if isDeFile {
		password := ""
		fmt.Printf("\r\n\r\n\r\n正在执行解密请输入文件密码:")
		fmt.Scanf("%s", &password)
		key = []byte(password)
		dir, _ := os.Getwd()
		defile := rand_str(20) + "dooxc"
		fmt.Printf("defile" + defile)
		decryptFile(filePath, defile)
		DeCompress(defile, dir)
		os.Remove(defile)
		fmt.Printf("\r\n解密完成,1秒后自动关闭")
		time.Sleep(1 * time.Second)
		// fmt.Printf("\r\n\r\n\r\n\r\n按回车键关闭...")
		fmt.Scanf(filePath)
	} else {
		createPrivKey()
		f3, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("error!\n")
		}
		defer f3.Close()
		var files = []*os.File{f3}
		// currentTime := time.Now().UnixMilli()
		dest := rand_str(20) + "dooxa"
		err = Compress(files, dest)
		enfile := rand_str(20) + enfileType
		encryptFile(dest, enfile)
		os.Remove(dest)
		var skey = bytes2str(key)
		fmt.Printf("加密后文件名:" + enfile)
		fmt.Printf("\r\n\r\n\r\n文件密码是:\n" + skey + "\n请妥善保管")
		fmt.Printf("\r\n提示,密码已复制到剪切板!")
		clipboard.WriteAll(skey)
		fmt.Printf("\r\n\r\n\r\n\r\n按回车键关闭...")
		fmt.Scanf(filePath)
	}
}

func main() {
	fmt.Printf("\r\n欢迎使用FastFileSafe(保密文件加密工具),版本v1.0.0")
	fmt.Printf("\r\ncreate by lxj")
	filePath := ""
	if len(os.Args) < 2 {
		fmt.Printf("\r\n\r\n\r\n\r\n请输入要加密/解密文件位置:")
		fmt.Scanf("%s", &filePath) //注意此方法在win下会因为\r\n读取两次
		fmt.Printf("filePath: %q\r\n", filePath)
	} else if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	isDeFile := strings.HasSuffix(filePath, enfileType)
	//如果是一个需要解压解密的文件那么走解压解密流程
	if isDeFile {
		runMain(filePath, true)
	} else {
		//如果不是那么就是需要进行压缩
		runMain(filePath, false)
	}
}
