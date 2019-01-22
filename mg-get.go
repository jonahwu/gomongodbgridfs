package main

import (
	"fmt"
	"gopkg.in/mgo.v2"

	"io"
	"log"
	"os"
)

type fileinfo struct {
	//文件大小
	LENGTH int32
	//md5
	MD5 string
	//文件名
	FILENAME string
}

func check(err error) {
	log.Print(err)
}

func upload() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	defer session.Close()
	if err != nil {
		fmt.Println("can not connect to mongodb")
		fmt.Println(err)
		return
	}
	names, err := session.DatabaseNames()
	if err != nil {
		fmt.Println("未查询到数据库名字:", err)
	}
	fmt.Println(names)
	//通过文件名创建mp3
	file, err := session.DB("gridfs").GridFS("fs").Create("my.mp3")
	if err != nil {
		fmt.Println(err)
		return
	}
	out, _ := os.OpenFile("/go/src/github.com/mongodbgridfs/a.png", os.O_RDWR, 0666)
	_, err = io.Copy(file, out)
	err = file.Close()
	err = out.Close()
}

/*
func delete() {
	//直接利用名字移除
	err := session.DB("gridfs").GridFS("fs").Remove("my.mp3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("刪除成功")
	}
}
*/

func ReadAll() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	defer session.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	names, err := session.DatabaseNames()
	if err != nil {
		fmt.Println("未查询到数据库名字:", err)
	}
	fmt.Println(names)
	//通过文件名获取mp3
	gfs := session.DB("gridfs").GridFS("fs")
	iter := gfs.Find(nil).Iter()

	result := new(fileinfo)
	for iter.Next(&result) {
		fmt.Println("一个一个输出：", result)
	}

}

func GetImage() {
	session, err := mgo.Dial("mongodb://127.0.0.1:27017")
	defer session.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	names, err := session.DatabaseNames()
	if err != nil {
		fmt.Println("未查询到数据库名字:", err)
	}
	fmt.Println(names)
	//通过文件名获取mp3
	file, err := session.DB("gridfs").GridFS("fs").Open("a.png")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(file)
	}

	//out, _ := os.OpenFile("/go/src/github.com/mongodbgridfs/o.txt", os.O_CREATE, 0666)
	//should need bindata to write
	out, _ := os.OpenFile("/go/src/github.com/mongodbgridfs/o.png", os.O_CREATE|os.O_RDWR, 0666)
	//out, _ := os.Open("/go/src/github.com/mongodbgridfs/o.png")
	_, err = io.Copy(out, file)
	// modify it to echo respoonse as following
	/*
		func (c *context) Stream(code int, contentType string, r io.Reader) (err error) {
			c.writeContentType(contentType)
			c.response.WriteHeader(code)
			_, err = io.Copy(c.response, r)
			return
		}
	*/
	check(err)
	err = file.Close()
	check(err)
}

func main() {
	//upload()
	ReadAll()
	GetImage()
}
