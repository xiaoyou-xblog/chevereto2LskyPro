## chevereto图床数据库转LskyPro

首先我们需要在lsky文件里面新建`images`目录

![image](./images/07.png)



下载项目，然后打包这个两个文件

![image-20210125162837561](./images/01.png)

打包后，自己上传到服务器上面，然后我们 `chmod +x main` 让文件具有执行权限

![image](./images/02.png)



# 迁移数据库

然后到`configs/app.ini` 中进行配置

![image](./images/03.png)

然后我们启动

![image](./images/04.png)

选择操作，我们先转换数据库

![image](./images/05.png)

下面我们需要迁移一下网站的数据

![image](./images/06.png)

最后我们打开网站就可以看到数据了

![image](./images/08.png)

# 删除重复文件

默认图床一个图片会默认分成三种，这里我们用不到

![image](./images/09.png)

我们这里注意配置一下路径

![image](./images/10.png)

这个是我们lsky图床的图片存储路径

![image](./images/11.png)

然后我们直接启动应用，选择删除重复文件

![image](./images/12.png)

最后就全部迁移完毕了！