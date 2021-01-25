// @Description 
// @Author 小游
// @Date 2021/01/25
package main

import (
	"fmt"
	"img/common"
	"img/model"
	"img/sql"
	"os"
	"regexp"
)

// 转换数据库
func changeData()  {
	// 初始化cheverto
	sql.InitDb1()
	// 初始化lsky
	sql.InitDb2()
	prefix1 := common.GetConfigString("cheverto", "Prefix")
	prefix2 := common.GetConfigString("lsky", "Prefix")
	// 获取iamges文件夹的id
	folder,err:=sql.Db2Dql("SELECT id from lsky_folders WHERE name = 'images'")
	if err!=nil || folder[0][0]==""{
		fmt.Println("没有查询到images文件夹id，请先新建image文件夹!")
		return
	}
	var input int
	// 是否需要清空数据库
	fmt.Printf("是否需要清空数据库(1:是 0:否):")
	if _, err := fmt.Scan(&input); err == nil && input == 1 {
		sql.Db2Dml("truncate "+prefix2+"images")
	}
	// 查询所有的图片
	if data,err:=sql.Db1Dql("SELECT image_name,image_extension,image_size,image_date,image_date_gmt,image_uploader_ip,image_md5 FROM "+prefix1+"images");err==nil{
		fmt.Printf("总计%d张图片!开始转换\n",len(data))
		errNum :=0
		for k,v:=range  data{
			var lsky model.LsKy
			lsky.Strategy = "local"
			create:=common.Str2time(v[3])
			// 路径
			lsky.Path = "images/" + common.Time2String(create,false)
			// 图片名
			lsky.Name = v[0] + "." + v[1]
			// 总路径
			lsky.PathName = lsky.Path + "/" + lsky.Name
			// 大小
			lsky.Size = v[2]
			// 类型
			switch v[1] {
			case "png":
				lsky.Mime = "image/png"
			case "jpg":
				lsky.Mime = "image/jpeg"
			case "jpeg":
				lsky.Mime = "image/jpeg"
			case "gif":
				lsky.Mime = "image/gif"
			case "ico":
				lsky.Mime = "image/x-icon"
			default:
				lsky.Mime = "image/png"
			}
			// md5加密值
			lsky.Md5 = v[6]
			// 上传ip
			lsky.Ip = v[5]
			// 上传时间
			lsky.Create = common.Str2time(v[4]).Unix()
			// 开始插入
			if sql.Db2Dml(`INSERT delayed INTO `+prefix2+`images
			(user_id,folder_id,strategy,path,`+"`name`"+`,alias_name,pathname,size,mime,sha1,md5,ip,suspicious,create_time) 
			VALUES 
			(1,?,?,?,?,'',?,?,?,?,?,?,0,?)`,
				folder[0][0],lsky.Strategy,lsky.Path,lsky.Name,lsky.PathName,lsky.Size,lsky.Mime,lsky.Sha1,lsky.Md5,lsky.Ip,lsky.Create){
				fmt.Printf("第%d张图片转换成功!\n",k)
			}else{
				errNum++
			}
		}
		fmt.Printf("转换完成！！！有%d转换失败", errNum)
	}else{
		fmt.Println(err)
	}
}


// 删除多余图片
func deleteMoreImage()  {
	path:=common.GetConfigString("img","path")
	list:=new([]string)
	// 获取所有的文件
	if common.GetAllFile(path,list)!=nil{
		fmt.Println("获取文件失败")
	}
	re,_:=regexp.Compile(`\.[th|md].*?$`)
	total:=0
	errNUm:=0
	// 遍历删除文件夹
	for _,v:=range *list{
		if re.MatchString(v){
			if os.Remove(v)==nil{
				fmt.Println("删除"+v)
				total++
			}else {
				errNUm++
			}
		}
	}
	// 打印成功
	fmt.Printf("已为你删除%d个文件，删除失败%d个",total,errNUm)
}


// 转换函数
func main()  {
	var input int
	// 是否需要清空数据库
	fmt.Printf("欢迎使用图床转换工具\n请选择操作(1转换数据库 2删除重复文件):")
	if _, err := fmt.Scan(&input); err == nil && input == 1 {
		changeData()
	}else if err == nil && input == 2{
		deleteMoreImage()
	}
}