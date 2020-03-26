## 当前开放系统
    Ubuntu 18.04.3 LTS
    
### 镜像打包
    <p align="center"><img src="/readme.png" alt="Godog logo" style="width:250px;" /></p>
    
    先打包main,再进行镜像打包
    
    GOOS=linux go build -o main main.go
    docker build -t go-mail:1.0.1 .
    eg：
        root@chen:/home/chen/go/src/github.com/go-mail# GOOS=linux go build -o main main.go
        root@chen:/home/chen/go/src/github.com/go-mail# docker build -t go-mail:1.0.1 .
        
### 本地运行服务
    docker run --rm -ti go-mail:1.0.1 
    root@chen:/home/chen/go/src/github.com/go-mail# docker run --rm -ti go-mail:1.0.1  
    
### 镜像上传
    docker images -> 查询[ImageId] 15f7d3c8a063
        sudo docker tag [ImageId] registry.cn-hangzhou.aliyuncs.com/sonicmoving/go-mail:[镜像版本号]
            sudo docker tag eb52921a3088 registry.cn-hangzhou.aliyuncs.com/sonicmoving/go-mail:1.0.1
        sudo docker push registry.cn-hangzhou.aliyuncs.com/sonicmoving/go-mail:1.0.1