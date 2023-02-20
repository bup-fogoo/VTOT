# VTOT
audio to text
# 项目地址
```
git clone https://github.com/bup-fogoo/VTOT.git
```
## 项目功能
对音频和视频文件进行智能语音识别成文字。
使用时请注意，文件名只允许字母和数字，格式只允许mp3、mp4、wav文件上传,
单个文件最大限制1G、时长120分钟。上传文件只保留一小时。
## 配置
配置路径 /config/application.yaml

需要注意的是这里的vps必须是公网，如需更改对外开放18888端口，docker-compose.yaml里的对外映射端口也要一并更改。

aliCloudSdk配置获取需要自行去 <em>阿里云智能语音识别服务</em> 获取以下三个appKey、accessKeyID、accessKeySecret。

```
vps: http://192.168.1.1:18888/

aliCloudSdk:
  appKey: xxx
  accessKeyID: xxx
  accessKeySecret: xxx
```
## 启动
ps:此项目需要已经安装好docker
```
docker-compose build
docker-compose up -d
```
