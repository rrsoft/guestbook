package utils

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	CHARS       = "0123456789abcdefghijklmnopqrstuvwxyz"
	TIME_FORMAT = "2006-01-02 15:04:05"
)

var (
	timeZero = time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)
	rnd      = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandString(n int) string {
	var s string
	l := len(CHARS)
	for ; n > 0; n-- {
		s += string(CHARS[rnd.Intn(l)])
	}
	return s
}

func RandNumber(min, max int) int {
	if min > max || min < 0 {
		return 0
	}
	if min == max {
		return min
	}
	return min + rnd.Intn(max-min)
}

func ToString(src interface{}) string {
	switch v := src.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	}
	return fmt.Sprintf("%v", src)
}

func SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	http.SetCookie(w, cookie)
}

func HasArrayItem(str string, arr []string) int {
	for i, c := range arr {
		if c == str {
			return i
		}
	}
	return -1
}

func HasContainItem(str string, arr []string) int {
	for i, c := range arr {
		if strings.Contains(c, str) {
			return i
		}
	}
	return -1
}

func HtmlEscape(str string) string {
	return html.EscapeString(str)
}

func HtmlUnescape(str string) string {
	return html.UnescapeString(str)
}

func UrlEscape(uri string) string {
	return url.QueryEscape(uri)
}

func UrlUnescape(uri string) string {
	if unurl, err := url.QueryUnescape(uri); err == nil {
		return unurl
	}
	return uri
}

func IsEmpty(str string) bool {
	str = strings.TrimSpace(str)
	return len(str) == 0
}

// 转换为整数 转换失败返回默认值
func ToInt(str string, def int) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		return def
	}
	return result
}

// 转换为长整数 转换失败返回默认值
func ToInt64(str string, def int64) int64 {
	result, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// 转换为无符号长整数 转换失败返回默认值
func ToUInt64(str string, def uint64) uint64 {
	result, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return def
	}
	return result
}

// 转换为浮点类型 转换失败返回默认值
func ToFloat(str string, def float64) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return def
	}
	return result
}

// 转换为布尔值 转换失败返回默认值
func ToBool(str string, def bool) bool {
	result, err := strconv.ParseBool(str)
	if err != nil {
		return def
	}
	return result
}

// 转换为时间类型
func ToTime(str string) time.Time {
	t, err := time.Parse(TIME_FORMAT, str)
	if err == nil {
		return t
	}
	return timeZero
}

// 验证是否为正整数
func IsInt(str string) (bool, error) {
	return regexp.MatchString(`^[0-9]+$`, str)
}

// 判断字符串是否是yyyy-mm-dd字符串
func IsDateString(str string) (bool, error) {
	return regexp.MatchString(`(\d{4})-(\d{1,2})-(\d{1,2})"`, str)
}

// 是否为ip
func IsIP(ip string) (bool, error) {
	return regexp.MatchString(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$`, ip)
}

// 判断是否为base64字符串 A-Z, a-z, 0-9, +, /, =
func IsBase64String(str string) (bool, error) {
	return regexp.MatchString(`[A-Za-z0-9\+\/\=]`, str)
}

// 检测是否有Sql危险字符
func IsSafeSqlString(str string) (bool, error) {
	b, err := regexp.MatchString(`[-|;|,|\/|\(|\)|\[|\]|\}|\{|%|@|\*|!|\']`, str)
	return !b, err
}

// 检测是否符合email格式
func IsEmail(email string) (bool, error) {
	return regexp.MatchString(`^([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$`, email)
}

// 检测是否是正确的Url
func IsUrl(url string) (bool, error) {
	return regexp.MatchString(`^(http|https)\://([a-zA-Z0-9\.\-]+(\:[a-zA-Z0-9\.&%\$\-]+)*@)*((25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9])\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[1-9]|0)\.(25[0-5]|2[0-4][0-9]|[0-1]{1}[0-9]{2}|[1-9]{1}[0-9]{1}|[0-9])|localhost|([a-zA-Z0-9\-]+\.)*[a-zA-Z0-9\-]+\.(com|edu|gov|int|mil|net|org|biz|arpa|info|name|pro|aero|coop|museum|[a-zA-Z]{1,10}))(\:[0-9]+)*(/($|[a-zA-Z0-9\.\,\?\'\\\+&%\$#\=~_\-]+))*$`, url)
}
