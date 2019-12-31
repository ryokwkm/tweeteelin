package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//対象文字が含まれる文字が存在するか
func Contains(array []string, target string) bool {
	for _, s := range array {
		//対象文字列が含まれているか
		if strings.Index(s, target) != -1 {
			return true
		}
	}
	return false
}

//対象文字と一致する文字が存在するか
func ContainEqual(array []string, target string) bool {
	for _, s := range array {
		if s == target {
			return true
		}
	}
	return false
}

func RandomChoice(length int) int {
	if length == 0 {
		return 0
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(length)
}

func CutString(word string, length int) string {
	for {
		if len(word) > length {
			sc := []rune(word)
			word = fmt.Sprintf(string(sc[:(len(sc) - 1)]))
		} else {
			break
		}
	}
	return word
}

/**
 *	リアルタイムデータからタグを抽出
 */
func ReplaceTweetHashTag(tagStr string) []string {
	var returnStr []string
	tags := strings.Split(tagStr, ",")
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		tag = strings.Replace(tag, " ", "_", -1)
		tag = strings.Replace(tag, "'", "", -1)
		tag = strings.Replace(tag, ".", "", -1)
		tag = strings.Replace(tag, "-", "_", -1)
		returnStr = append(returnStr, tag)
	}
	return returnStr
}
