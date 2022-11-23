/*	Программа для уникализации строк, которые подаются в стандартный ввод. На вход подаются
отсортированные по кол-ву байт строки. Если новая строка больше текущей, то она новая, напечатаем
её, если совпадает, то мы её уже выводили, если меньше, то это значит, что файл не отсортирован */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt
		fmt.Fprintln(output, txt)
	}
	return nil
}
func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
