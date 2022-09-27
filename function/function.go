package function

import (
	"bufio"
	"fmt"
	"os"
)

//写入新文件
func Write(path string , p []byte){

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err!= nil {
		fmt.Println("写入文件失败")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write(p)

	writer.Flush()

}

//追加写文件
func WriteAppend(path string , p []byte){
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err!= nil {
		fmt.Println("写入文件失败")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write(p)

	writer.Flush()
}



//读取文件
func Read(in string) []byte {
	file, err := os.OpenFile(in, os.O_RDWR|os.O_APPEND, 0666)
	if err!= nil {
		fmt.Println("读文件失败")
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	//读取 Reader 对象中的内容到 []byte 类型的 buf 中
	reader.Read(buf)
	return buf
}

//创建文件夹
func Mkdir(path string){
	os.MkdirAll(path,os.ModePerm)
}

//删除文件(夹）
func Delete(path string){
	err := os.RemoveAll(path)
	if err != nil{
		fmt.Println("文件删除失败",err)
	}
}

//创建数据集
func DataSet(path string){
	path = "/dataset"+ path
	os.MkdirAll(path,os.ModePerm)
}
