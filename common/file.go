// @Description  读取项目的配置文件
// @Author 小游
// @Date 2021/01/25
package common

/**
这里主要是写和文件相关的函数
*/
import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"os"
)

/*读取配置文件*/
func GetConfig(module string) map[string]string {
	config, err := goconfig.LoadConfigFile("configs/app.ini") //加载配置文件
	if err != nil {
		return nil
	}
	glob, _ := config.GetSection(module) //读取全部mysql配置
	return glob
}

/*设置配置文件*/
func SetConfig(module string, key string, value string) bool {
	const file = "configs/app.ini"
	config, err := goconfig.LoadConfigFile(file) //加载配置文件
	if err != nil {
		return false
	}
	if !config.SetValue(module, key, value) {
		if err := goconfig.SaveConfigFile(config, file); err == nil {
			return true
		} else {
			fmt.Println(err.Error())
		}
	}
	return false
}

/*文件夹是否存在*/
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	}
	return true
}

//获取配置文件
func GetConfigString(module string, key string) string {
	config, err := goconfig.LoadConfigFile("configs/app.ini") //加载配置文件
	if err != nil {
		return ""
	}
	glob, err := config.GetSection(module) //读取全部mysql配置
	if err == nil {
		return glob[key]
	}
	return ""
}

// 获取某个路径下所有的路径
func GetAllFile(pathname string,list *[]string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			GetAllFile(pathname + fi.Name() + "/", list)
		} else {
			*list = append(*list,pathname+fi.Name())
		}
	}
	return err
}