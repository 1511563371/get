package utils

import (
	"fmt"
	"github.com/pterm/pterm"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var debug = false

// SetDebug set env to debug
func SetDebug() {
	debug = true
}

// IsDebug Current env is debug?
func IsDebug() bool {
	return debug
}

// EncodeToUrl encode val to urlEncode
func EncodeToUrl(val string) string {
	return url.QueryEscape(val)
}

// ReadBody request body and close it.
func ReadBody(readerCloser io.ReadCloser) []byte {
	defer readerCloser.Close()
	bytes, err := io.ReadAll(readerCloser)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func MappingArtName(name string) string {
	if name == "" {
		return "未知"
	}
	return name
}

func NormalizeFileName(name string) string {
	return strings.ReplaceAll(name, ":", "")
}

const (
	BYTE = 1.0 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

func FormatBytes(bytes int64) string {
	unit := ""
	value := float32(bytes)

	switch {

	case bytes >= TERABYTE:
		unit = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = value / KILOBYTE
	case bytes == 0:
		return "0"

	}

	result := fmt.Sprintf("%.2f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}

func ToInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}

var maxWidth int

func init() {
	w, _, err := pterm.GetTerminalSize()
	if err != nil {
		panic(err)
	}
	maxWidth = w
}

func DefaultTruncate(s string) string {
	return Truncate(s, maxWidth-5)
}

func Truncate(s string, size int) string {
	slen := len(s)
	if slen < size {
		return s
	}
	sprintf := fmt.Sprintf("%s", s[:size])
	return sprintf
	//return fmt.Sprintf("%s%s", s[0:size], "...")
}

func GetSize(url string) int64 {
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}

	return resp.ContentLength
}
