package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	a := "input.txt"
	b := "output.txt"
	c := "СКРЫТО"

	d, e := os.Open(a)
	if e != nil {
		fmt.Println("Ошибка при открытии файла:", e)
		return
	}
	defer d.Close()

	f, g := os.Create(b)
	if g != nil {
		fmt.Println("Ошибка при создании файла:", g)
		return
	}
	defer f.Close()

	h := bufio.NewScanner(d)
	i := bufio.NewWriter(f)

	j := map[string]bool{
		"Info:":         false,
		"Stream:":       false,
		"File Path:":    false,
		"Platform:":     false,
		"Author:":       false,
		"Sound:":        false,
		"OS:":           true,
		"Uploaded:":     true,
		"Title:":        true,
		"Description:":  true,
		"Duration:":     true,
		"Resolution:":   true,
		"Frame Rate:":   true,
		"Codec:":        true,
		"Bitrate:":      true,
		"Aspect Ratio:": true,
		"Language:":     true,
		"Subtitles:":    true,
		"Tags:":         true,
		"Category:":     true,
		"Release Date:": true,
		"License:":      true,
		"Views:":        true,
		"Likes:":        true,
		"Comments:":     true,
	}

	for h.Scan() {
		k := h.Text()
		l := strings.SplitN(k, ": ", 2)
		if len(l) == 2 && j[l[0]+":"] {
			k = l[0] + ": " + c
		}
		_, m := i.WriteString(k + "\n")
		if m != nil {
			fmt.Println("Ошибка при записи в файл:", m)
			return
		}
	}

	if n := h.Err(); n != nil {
		fmt.Println("Ошибка при чтении файла:", n)
		return
	}

	o := i.Flush()
	if o != nil {
		fmt.Println("Ошибка при сбросе буфера:", o)
		return
	}

	fmt.Println("Информация успешно скрыта. Выходные данные записаны в", b)
}
