# dp-grafana-webhook

#### 启动方式  
    docker run -it -d \
        -e "SERVICEKEY=823f1602cb589bfc2ff9e9809199ecd5" \
        -e "SERVERID=9960" -e "SERVERURL=http://gaojing.baidu.com/event/create" \
        -e "ADVERTISEDURL=0.0.0.0:8888" \
        -p 8888:8888 dp-grafana-webhook

#### 环境变量解释
- SERVICEKEY  百度的token
- SERVERID    服务id
- SERVERURL   目前为定值
- ADVERTISEDURL  rest api 监听ip 以及端口
