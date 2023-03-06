package main

import (
	"bufio"
	"fmt"
	"strings"
)

type TextCount struct {
	words int
	rows  int
}

func (t *TextCount) LineCount(scanner *bufio.Scanner) {

}

func (t *TextCount) WordCount(scanner *bufio.Scanner) {

}

func main() {
	//var str string = "Hello World Where are you"
	//var str = "jack700001\tb17c1098bbbaffa12a8b596a27790e77\njack700002\t6e59cf7586f2a394ba0ad0d6ce63426f\njack700003\ta93debabd2cefa8a90ce6bfa79473b2c\njack700004\t2c22ad6af59c34a49f2c4900fc7cf36e\njack700005\t63eb298f982ecbd43939f4e9628344ad\njack700006\t135aebfcfae1aaacd0a888e18f80fd50\njack700007\t30e1cb1b3c57e6518e03ac952f31c063\njack700008\t7890e9981d4c27a1a9b2d0e8401e2c9d\njack700009\t847835113083c8a67625c340c8c4a32c\njack700010\td9e72c6fb851cea21b046f75f9f79ba1\njack700011\td2072f82d90aa69cf3d1f0c16b3788ef\njack700012\t2a36db0089fdbd31aff24960a774dad1\njack700013\t2eeb86b64df5d3c151b3c798d33b8fa1\njack700014\t7856cde4a2b9b6f770bdb7647c8f2b40\njack700015\tabd5497a0e4906f8965ebc537df68083\njack700016\t941845d62a17b22172399289f3bdd170\njack700017\t125c8254c1e5e36c8d0d427c2ea86fbd\njack700018\t4a8108730896a4d1a0ef53bb32dec0f8\njack700019\tc1b2fc6e8f57ed57f90dc4ea5883c6c0\njack700992\tc5f2ae9a1759a16d5e5d280895495eaa\njack700993\t378c324a8ba742501c665114578fac40\njack700994\tdb6f9a70c5c5e3ed21949ae5ffc9c059\njack700995\td035d79abe85c26bd48419c31ee2086b\njack700996\tb3f71b753614679de8f509b8c6bcf5df\njack700997\t15d984aa53d440be0df14643540aac7d\njack700998\tba741a4ce24b33ca813a7a2dfaf8ede1\njack700999\t0eebb86b5144bb3dfc72caff5af660c8\njack701000\tf46c70c53e0fd84e478dda8b3ed798f9"
	var str = "jack700001 b17c1098bbbaffa12a8b596a27790e77\njack700002 6e59cf7586f2a394ba0ad0d6ce63426f\njack700003 a93debabd2cefa8a90ce6bfa79473b2c\njack700004 2c22ad6af59c34a49f2c4900fc7cf36e\njack700005 63eb298f982ecbd43939f4e9628344ad\njack700006 135aebfcfae1aaacd0a888e18f80fd50\njack700007 30e1cb1b3c57e6518e03ac952f31c063\njack700008 7890e9981d4c27a1a9b2d0e8401e2c9d\njack700009 847835113083c8a67625c340c8c4a32c\njack700010 d9e72c6fb851cea21b046f75f9f79ba1\njack700011 d2072f82d90aa69cf3d1f0c16b3788ef\njack700012 2a36db0089fdbd31aff24960a774dad1\njack700013 2eeb86b64df5d3c151b3c798d33b8fa1\njack700014 7856cde4a2b9b6f770bdb7647c8f2b40\njack700015 abd5497a0e4906f8965ebc537df68083\njack700016 941845d62a17b22172399289f3bdd170\njack700017 125c8254c1e5e36c8d0d427c2ea86fbd\njack700018 4a8108730896a4d1a0ef53bb32dec0f8\njack700019 c1b2fc6e8f57ed57f90dc4ea5883c6c0\njack700992 c5f2ae9a1759a16d5e5d280895495eaa\njack700993 378c324a8ba742501c665114578fac40\njack700994 db6f9a70c5c5e3ed21949ae5ffc9c059\njack700995 d035d79abe85c26bd48419c31ee2086b\njack700996 b3f71b753614679de8f509b8c6bcf5df\njack700997 15d984aa53d440be0df14643540aac7d\njack700998 ba741a4ce24b33ca813a7a2dfaf8ede1\njack700999 0eebb86b5144bb3dfc72caff5af660c8\njack701000 f46c70c53e0fd84e478dda8b3ed798f9"
	lineScanner := bufio.NewScanner(strings.NewReader(str))
	lineScanner.Split(bufio.ScanLines)
	//bufio.li
	var lineCount, wordCount int
	for lineScanner.Scan() {
		//fmt.Println(lineScanner.Text())
		lineCount++

		//wordScanner := bufio.NewScanner(strings.NewReader(lineScanner.Text()))
	}

	//wordScanner := bufio.NewScanner(strings.NewReader(str))
	lineScanner.Split(bufio.ScanWords)
	for lineScanner.Scan() {
		wordCount++
	}

	fmt.Printf("行数：%d, 单词数：%d\n", lineCount, wordCount)

	//var str string = "Hello World Where are you"
	//scaner := bufio.NewScanner(strings.NewReader(str))
	//scaner.Split(bufio.ScanWords)
	//for scaner.Scan() {
	//	fmt.Println(scaner.Text())
	//}

}
