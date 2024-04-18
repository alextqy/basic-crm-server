package mtd

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SysHelper struct{}

func (s *SysHelper) SysEnvs() map[string]string {
	envMap := map[string]string{}
	envs := os.Environ()
	for _, e := range envs {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) != 2 {
			continue
		} else {
			envMap[string(parts[0])] = string(parts[1])
		}
	}
	return envMap
}

func (s *SysHelper) GetEnv(key string) string {
	return os.Getenv(key)
}

func (s *SysHelper) LocalIP() (bool, string, []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, err.Error(), nil
	} else {
		var ips []string
		for _, ads := range addrs {
			if ipnet, ok := ads.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ips = append(ips, string(ipnet.IP.String()))
				}
			}
		}
		return true, "", ips
	}
}

func (s *SysHelper) StringToByte(data string) []byte {
	return []byte(data)
}

func (s *SysHelper) ByteToString(data []byte) string {
	return string(data)
}

func (s *SysHelper) StringToInt(data string) (bool, string, int) {
	res, err := strconv.Atoi(data)
	if err != nil {
		return false, err.Error(), 0
	} else {
		return true, "", res
	}
}

func (s *SysHelper) StringToInt64(data string) (bool, string, int64) {
	res, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return false, err.Error(), 0
	} else {
		return true, "", res
	}
}

func (s *SysHelper) IntToString(data int) string {
	return strconv.Itoa(data)
}

func (s *SysHelper) Int64ToString(data int64) string {
	return strconv.FormatInt(data, 10)
}

func (s *SysHelper) StringToFloat32(data string) (bool, string, float64) {
	r, err := strconv.ParseFloat(data, 32)
	if err != nil {
		return false, err.Error(), 0
	} else {
		return true, "", r
	}
}

func (s *SysHelper) StringToFloat64(data string) (bool, string, float64) {
	r, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return false, err.Error(), 0
	} else {
		return true, "", r
	}
}

func (s *SysHelper) Float64ToString(data float64) string {
	return strconv.FormatFloat(data, 'E', -1, 32)
}

func (s *SysHelper) IntToBytes(data int) []byte {
	dataInt := int32(data)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, dataInt)
	return bytesBuffer.Bytes()
}

func (s *SysHelper) BytesToInt(data []byte) int {
	bytesBuffer := bytes.NewBuffer(data)
	var dataInt int32
	binary.Read(bytesBuffer, binary.BigEndian, &dataInt)
	return int(dataInt)
}

func (s *SysHelper) TimeNow() time.Time {
	return time.Now()
}

func (s *SysHelper) TimeNowStr() string {
	return time.Now().Format("2006-01-02 15:04:05") // 2006-01-02 15:04:05 golang立项时间
}

func (s *SysHelper) TimeStamp() int64 {
	return time.Now().Unix()
}

func (s *SysHelper) TimeStampMS() int64 {
	return time.Now().UnixNano()
}

func (s *SysHelper) TimeStampToStr(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

func (s *SysHelper) MD5(p string) string {
	hasher := md5.New()
	hasher.Write([]byte(p))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (s *SysHelper) EnBase64(p string) string {
	return base64.StdEncoding.EncodeToString([]byte(p))
}

func (s *SysHelper) DeBase64(s64 string) (bool, string, string) {
	decoded, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		return false, err.Error(), ""
	}
	return true, "", string(decoded)
}

func (s *SysHelper) StringContains(data, subs string) bool {
	return strings.Contains(data, subs)
}

// 随机字符串
func (s *SysHelper) RandStr(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/*
CBC 加密
data 待加密的明文
key 秘钥
vi 向量
*/
func (s *SysHelper) AesEncrypterCBC(data_s, key_s, iv_s string) (bool, string, string) {
	data := []byte(data_s)
	key := []byte(key_s)
	iv := []byte(iv_s)
	block, err := aes.NewCipher(key)
	if err != nil {
		return false, err.Error(), ""
	}
	padding := block.BlockSize() - len(data)%block.BlockSize()
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(block.BlockSize())}, block.BlockSize())
	} else {
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	paddText := append(data, paddingText...)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	result := make([]byte, len(paddText))
	blockMode.CryptBlocks(result, paddText)
	return true, "", string(result)
}

/*
CBC 解密
data 待解密的密文
key 秘钥
vi 向量
*/
func (s *SysHelper) AesDecrypterCBC(data_s, key_s, iv_s string) (bool, string, string) {
	data := []byte(data_s)
	key := []byte(key_s)
	iv := []byte(iv_s)
	block, err := aes.NewCipher(key)
	if err != nil {
		return false, err.Error(), ""
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	result := make([]byte, len(data))
	blockMode.CryptBlocks(result, data)
	unPadding := int(result[len(result)-1])
	return true, "", string(result[:(len(result) - unPadding)])
}

// 大小写英文字母
func (s *SysHelper) RegEn(p string) bool {
	r, err := regexp.Compile("^[a-zA-Z]+$")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

// 数字
func (s *SysHelper) RegNum(p string) bool {
	r, err := regexp.Compile("^[0-9]*$")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

// 中文
func (s *SysHelper) RegZh(p string) bool {
	r, err := regexp.Compile("[\u4e00-\u9fa5]")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

// 英文 数字
func (s *SysHelper) RegEnNum(p string) bool {
	r, err := regexp.Compile("^[a-zA-Z0-9]+$")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

// 中英文 数字 下划线 短横线 中英文(逗号 句号 分号 感叹号 换行符 任何空白字符)
func (s *SysHelper) RegWriting(p string) bool {
	r, err := regexp.Compile("^[\u4e00-\u9fa5_a-zA-Z0-9-,.;!，。；！\\n\\s]+$")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

// 中英文 数字 下划线 短横线
func (s *SysHelper) RegAll(p string) bool {
	r, err := regexp.Compile("^[\u4e00-\u9fa5_a-zA-Z0-9-]+$")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

func (s *SysHelper) RegEmail(p string) bool {
	r, err := regexp.Compile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return r.MatchString(p)
}

/*
发送邮件(腾讯邮箱三方设备授权码: qfjhhammjflgbjcc)
account 邮箱(tqyalex@qq.com)
password 密码
sender 发送者名称
host 邮箱服务器(smtp.qq.com:465)
to 客户邮箱
subject 标题(Reset Password)
body 内容(验证码)
*/
func (s *SysHelper) SendEmail(account, password, sender, host, to, subject, body string) (bool, string) {
	if account == "" {
		return false, "incorrect account"
	}
	if password == "" {
		return false, "incorrect password"
	}
	if sender == "" {
		return false, "incorrect sender"
	}
	if host == "" {
		return false, "incorrect host"
	}
	if to == "" {
		return false, "incorrect email address"
	}
	if subject == "" {
		return false, "incorrect title"
	}
	if body == "" {
		return false, "incorrect content"
	}
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", account, password, hp[0])
	msg := []byte("To: " + to + "\r\nFrom:" + sender + "<" + account + ">" + "\r\nSubject:" + subject + "\r\nContent-Type:text/plain;charset=UTF-8\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, account, sendTo, msg)
	if err != nil {
		return false, err.Error()
	}
	return true, ""
}
