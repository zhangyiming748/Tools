# 定期清理git历史
1. Checkout

```shell
$ git checkout --orphan latest_branch  // 基于当前分支创建新分支
```

2. Add all the files

```shell
$ git add -A // 追踪全部文件
```

3. Commit the changes

```shell
$ git commit -am "commit message" // 添加提交信息
```

4. Delete the branch

```shell
$ git branch -D master // 删除主分支
```

5. Rename the current branch to master

```shell
$ git branch -m master // 命名当前分支为主分支
```

6. Finally, force update your repository

```shell
$ git push -f --set-upstream origin master // 强制推送并设置与远程分支的对应关系
```

# 目前已有功能

- [x] 旋转视频
- [x] 转换视频为mp4格式
- [x] 图片转换为webP
- [x] webP转换为其他格式
- [x] 暴力破解压缩文件密码 切记**解压目录一定要新建的空目录**
- [x] 查询天气
- [x] 探测文件类型
- [x] 批量转换文本编码
- [x] 仅提取音频
- [x] 转换视频为webm
- [x] 使用youtube-dl批量下载
- [x] 调用hey命令压测api
- [ ] 使用curl查询天气并保存为图片
- [x] 计算文件哈希值(sha1/sha256/md5)
- [x] 遍历UOS网站下载全部系统镜像
- [x] 批量剪切片头
- [x] 多线程下载功能

## conf.ini
![新版配置文件](https://s1.ax1x.com/2022/09/12/vXJLkD.png)
```log
[main]
# function = rotate
# function = ToMp4
# function = ToWebm
# function = ToWebp
# function = ToGif
# function = WebpTo
# function = Unzip
# function = Weather
# function = Detect
# function = ExtractAudio
# function = GetTime
# function = hey
# function = youtube-dl
# function = WeatherPNG
# function = HASH
# function = wget
# function = duplicate
# function = decode
# function = cut
# function = multi
# function = ToFlac
# function = ToMp3
# function = bilibili
# function = frame
# function = ToH265
# function = MediaInfo
# function = resolution
# function = panic
# function = printLog
function = default
# delAfterDone = true
[rotate]
direction = ToRight
# direction = ToLeft
[location]
src = /Users/zen/Downloads/Downie
pattern = webm;mp4;MP4;mov:MOV
dst = /Users/zen/Downloads/Downie/h265
webpto = jpg
passwd = /Users/zen/Downloads/passwd.txt
[youtube-dl]
goroutine = 2
fp = /Users/zen/Gitlab/Tools/link.txt
addr = 127.0.0.1
port = 8889
target = /Users/zen/Gitlab/Tools/storage
isproxy = true
[hey]
URL = http://127.0.0.1:3306/HappyCount
Requests = 10000
Concurrent = 50
[unzip]
passwd = passwdDict.txt
[download]
wget = /home/zen/go/src/Tools/test/balenaetcher.sh
multi = /home/zen/Github/Tools/link.txt
[title]
start = 00:01:40.079
[frame]
fps = 30
[resolution]
p = 1080
```

# 关于shell

> deb系默认没有安装dos2unix工具,而且也没有一个叫这个名字的工具(我在solaris里用过dos2unix,不知道为啥Ubuntu没有)但是有一个替代工具tofrodos,下面就说一下它的安装和使用

## 安装tofrodos

```shell
$ sudo apt-get install tofrodos
```

### 用法

```shell
todos Hello.txt (即unix2dos Hello.txt) 
fromdos Hello.txt (即dos2unix Hello.txt)
```

## 做一些优化

```shell
ln -s /usr/bin/todos /usr/bin/unix2dos 
ln -s /usr/bin/fromdos /usr/bin/dos2unix 
# 或者
echo "alias unix2dos=todos" >> ~/.bash_profile
echo "alias dos2unix=fromdos"" >> ~/.bash_profile
```

# rotate功能说明

`旋转符合条件的视频`

```
src = 输入目录
dst = 输出目录
pattern = 扩展名
**输入输出不能为相同目录**
可选参数:
delAfterDone = 处理结束后删除源文件
```

# ToMp4功能说明

`转换符合条件的视频文件为mp4格式`

```
src = 输入目录
dst = 输出目录
pattern = 扩展名
可选参数:
delAfterDone = 处理结束后删除源文件
```

# ToWebm功能说明

`转换符合条件的视频文件为webm格式`

```
src = 输入目录
dst = 输出目录
pattern = 扩展名
可选参数:
delAfterDone = 处理结束后删除源文件
```

# ToWebp功能说明

`转换符合条件的图片文件为webp格式`

```
src = 输入目录
pattern = 扩展名
可选参数:
delAfterDone = 处理结束后删除源文件
```

# WebpTo功能说明

`转换webp图片文件为指定格式`

```
src = 输入目录
pattern = 扩展名
可选参数:
delAfterDone = 处理结束后删除源文件
```

# Unzip功能说明

`使用字典破解zip压缩文件`

```
src = 输入文件
passwd = 字典位置
```

# Weather功能说明

`使用高德开放api获取指定地点天气`

# ExtractAudio功能说明

`提取视频中音频`

```
src = 输入目录
pattern = 扩展名
```

# GetTime功能说明

`输入开始和结束时间获取ffmpeg可识别的命令`

# hey

`使用hey压测网址`

```
URL = 网址
Requests = 请求次数
Concurrent = 并发数
```

# youtube-dl功能说明

`使用youtube-dl命令下载列表中网址的视频`

```
goroutine = 并发下载数
fp = 列表文件位置
addr = 代理ip
port = 代理端口
target = 视频保存位置
isproxy = 是(1)否(0)使用代理
```

# WeatherPNG功能说明

`生成当前位置的天气预报图片`

# 哈希功能说明

`如果只给一个目录(src,pattern生效),打印所有文件的哈希值`
`如果给两个文件(src,dst生效),判断两个文件是否相等`

# WGET功能说明

```
wget = 文件列表
dst = 下载位置
```

# decode功能说明

```
转换GBK/GB2312/GB18030编码的文本文档到UTF8
src = 要转换的文件夹路径
pattern = 文本文件的扩展名
```