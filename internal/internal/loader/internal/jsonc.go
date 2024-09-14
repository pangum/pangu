package internal

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pangum/pangu/internal/internal/loader/internal/internal/constant"
)

type Jsonc struct {
	index       int    // 当前处理坐标
	comment     uint   // 注释
	len         int    // 长度
	objectDepth uint   // 对象深度
	arrayDepth  uint   // 数组深度
	inString    bool   // 是否在字符串中
	inArray     bool   // 是否在数组中
	inObject    bool   // 是否在对象中
	last        string // 最后一个有效非空格字符
	stringLabel string // 字符串标识

	comma      *regexp.Regexp
	spacesPair map[string]string
}

func NewJsonc() *Jsonc {
	return &Jsonc{
		comma: regexp.MustCompile(`(?:,+)(\s*)$`),
		spacesPair: map[string]string{
			"\n": `\n`,
			"\t": `\t`,
			"\r": `\r`,
		},
	}
}

func (j *Jsonc) Strip(from string) (to string) {
	j.reset()

	j.len = len(from)
	original := ""
	prev := ""
	char := ""
	next := ""
	quote := ""
	quoted := false

	for j.index < j.len {
		original, prev, char, next = j.getSegments(from, prev)

		// 处理十六进制数据
		if j.isNonStringValue(char, "0") && ("x" == next || "X" == next) {
			to += j.hexadecimal(from)
			continue
		}
		// 果内部字符串或外部注释，则按原样附加
		if j.insideString(prev, char, next, original) || j.outsideComment(char, next) {
			to += j.compliment(prev, char, next)
			continue
		}

		j.strip(&to, char, next, &quote, &quoted)
	}

	return
}

func (j *Jsonc) strip(to *string, char string, next string, quote *string, quoted *bool) {
	*quote, *quoted = j.quoteKey(char, *quoted)
	*to += *quote

	// 修剪数组或对象末尾的尾部逗号
	if j.comment == 0 && !j.inString && ((j.inArray && char == "]") || (j.inObject && char == "}")) {
		*to = j.comma.ReplaceAllString(*to, `$1`)
	}

	j.checkArrayObject(char)

	// 清除评论周围的尾随空格
	if j.hasCommentEnded(char, next) && char == "\n" {
		*to = strings.TrimRight(*to, "\r\n\t") + char
	}
}

func (j *Jsonc) hexadecimal(from string) (to string) {
	hexadecimal := ""
	j.index++
	for j.index < j.len {
		char := from[j.index : j.index+1]
		if !j.isNumber(char, true) {
			break
		}
		hexadecimal += char
		j.index++
	}
	if dec, pie := strconv.ParseInt(hexadecimal, 16, 32); nil == pie {
		to = fmt.Sprintf("%d", dec)
	}

	return
}

func (j *Jsonc) isNumber(char string, hex bool) (number bool) {
	if hex {
		number = strings.ContainsAny(char, "0123456789abcdefABCDEF")
	} else {
		number = strings.ContainsAny(char, "0123456789")
	}

	return
}

func (j *Jsonc) outsideComment(char, next string) bool {
	if !j.inString && j.comment == 0 {
		if char+next == "//" {
			j.comment = 1
		}
		if char+next == "/*" {
			j.comment = 2
		}
	}

	return 0 == j.comment
}

func (j *Jsonc) compliment(prev string, char string, next string) (final string) {
	if j.inString && char == constant.Esc && constant.Enter == next {
		j.index++
		final = ""
	} else if cached, ok := j.spacesPair[char]; ok && j.inString {
		final = cached
	} else if j.isNonStringValue(char, "+") && j.isNumber(next, false) {
		final = ""
	} else if j.isNonStringValue(char, ".") {
		final = j.decimalPointNumber(prev, next)
	} else if constant.SingleQuote == j.stringLabel {
		final = j.singleQuotedString(prev, char, next)
	} else {
		final = char
	}

	return
}

func (j *Jsonc) decimalPointNumber(prev string, next string) (final string) {
	prevNumber := j.isNumber(prev, false)
	nextNumber := j.isNumber(next, false)
	if !prevNumber && nextNumber {
		final = "0."
	} else if prevNumber && !nextNumber {
		final = ".0"
	}

	return
}

func (j *Jsonc) singleQuotedString(prev string, char string, next string) (final string) {
	if constant.Esc+constant.SingleQuote == char+next {
		j.index++
		final = constant.SingleQuote
	} else if constant.Esc != prev && constant.SingleQuote == char {
		final = constant.DoubleQuote
	} else if constant.DoubleQuote == char {
		final = `\"`
	}

	return
}

func (j *Jsonc) checkArrayObject(char string) {
	if j.isNonStringValue(char, char) {
		if !strings.ContainsAny(char, "\r\n\t /") {
			j.last = char
		}
		if char == "{" {
			j.objectDepth++
			j.inObject, j.inArray = true, false
		} else if j.objectDepth > 0 && char == "}" {
			j.objectDepth--
			j.inObject, j.inArray = j.objectDepth > 0, j.arrayDepth > 0
		} else if char == "[" {
			j.arrayDepth++
			j.inObject, j.inArray = false, true
		} else if j.arrayDepth > 0 && char == "]" {
			j.arrayDepth--
			j.inObject, j.inArray = j.objectDepth > 0, j.arrayDepth > 0
		}
	}
}

func (j *Jsonc) isNonStringValue(char, chars string) bool {
	return !j.inString && j.comment == 0 && strings.ContainsAny(char, chars)
}

func (j *Jsonc) quoteKey(char string, wasQuoted bool) (quote string, quoted bool) {
	quoted = wasQuoted
	inKey := j.inObject && j.comment == 0 && (j.last == "{" || j.last == ",")
	if !j.inString && inKey && !strings.ContainsAny(char, "[]{}'\",/*:\r\n\t ") {
		quote = constant.DoubleQuote
		j.inString, quoted, j.stringLabel = true, true, constant.DoubleQuote
	}
	if j.inString && wasQuoted && inKey && (char == ":" || char == " " || char == constant.SingleQuote) {
		quote = constant.DoubleQuote
		j.inString, quoted, j.stringLabel = false, false, ""
	}

	return
}

func (j *Jsonc) insideString(prev string, char string, next string, old string) bool {
	charNext := char + next
	maybeStr := (char == constant.DoubleQuote || char == constant.SingleQuote) && (!j.inString || j.stringLabel == char)
	if j.comment == 0 && maybeStr && prev != constant.Esc {
		if !j.inString {
			j.stringLabel = char
		}
		j.inString = !j.inString
		return j.inString
	}
	if j.inString && (charNext == `":` || charNext == `",` || charNext == `"]` || charNext == `"}`) {
		j.inString = old+prev != constant.Esc+constant.Esc
	}

	return j.inString
}

func (j *Jsonc) getSegments(json string, old string) (original, prev, char, next string) {
	original = old
	if j.index > 0 {
		prev = json[j.index-1 : j.index]
	}
	char = json[j.index : j.index+1]
	if j.index < j.len-1 {
		next = json[j.index+1 : j.index+2]
	}
	j.index++

	return
}

func (j *Jsonc) hasCommentEnded(char, next string) bool {
	singleEnded := j.comment == 1 && char == "\n"
	multiEnded := j.comment == 2 && char+next == "*/"
	if singleEnded || multiEnded {
		j.comment = 0
	}
	if multiEnded {
		j.index++
	}

	return 0 == j.comment
}

func (j *Jsonc) reset() {
	j.index, j.comment = 0, 0
	j.objectDepth, j.arrayDepth = 0, 0
	j.inString, j.inArray, j.inObject = false, false, false
	j.last, j.stringLabel = "", ""
}
