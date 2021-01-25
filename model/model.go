// @Description lsky数据库模型
// @Author 小游
// @Date 2021/01/25
package model

type LsKy struct {
	Strategy string  // 存储策略
	Path     string  // 保存路径
	Name     string  // 保存名称
	PathName string  // 总名称
	Size     string  // 图片大小
	Mime     string  // 图片类型
	Sha1     string  // 图片 hash sha1
	Md5      string  // 图片 hash md5加密
	Ip       string  // 上传者ip
	Create   int64  // 创建时间
}
