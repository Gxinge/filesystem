package peizhi

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

var Juicefs *juicefs

type conf1 struct {
	Jfs juicefs `yaml:"juicefs"`
}

type juicefs struct {
	Path       string `yaml:"path"`
	Cachesize  int    `yaml:"cachesize"`
	Log        string `yaml:"log"`
	Cachedir   string `yaml:"cachedir"`
	TestOrDemo string `yaml:"testOrDemo"`
}

func Init(dataFile string) {
	yamlFile, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(conf1)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	log.Printf("load conf success\n %v", c)
	// 绑定到外部可以访问的变量中
	Juicefs = &c.Jfs

	//1.判断路径是否合法
	path := Juicefs.Path
	fmt.Println(path)
	_, err1 := os.Stat(path)
	if err1 == nil {
		fmt.Println("路径不合法")
		os.Exit(3)
	}

	//2.判断缓存空间是否足够
	cachesize := Juicefs.Cachesize
	cachedir := Juicefs.Cachedir
	if cachesize > 0 {
		usage := diskUsage(cachedir)
		if uint64(cachesize*1024) > usage {
			cachesize = int(usage / 1024)
		}
	}
	//3.判断是测试环境还是演示环境
	var url string
	if Juicefs.TestOrDemo == "test" {
		url = "tikv://10.101.12.93:2379/test "
	} else {
		url = "tikv://10.101.12.93:2379/jfsdemo "
	}

	//4.挂载客户端
	var commond = "juicefs mount " + "--log " + Juicefs.Log + " --cache-dir " + cachedir + " --cache-size " + strconv.Itoa(cachesize) + " " + url + path
	commands := exec.Command(commond)
	output, err := commands.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(output))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Println(commond)
}

func diskUsage(path string) uint64 {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		fmt.Println(err)
	}
	size := fs.Bavail * uint64(fs.Bsize)
	return size
}
