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


#### 配置方式
1. 打开http://staging.datapipeline.com:3000/alerting/notifications
2. New Channel  -->  Type (webhook) 
3. url 填写dp-grafana-webhook http://ip:port  method post
4. 这是每个dashbord的报警


#### 处理报警
1. 登陆http://tonggao.baidu.com/console 查看报警描述，定位报警来源；
2. 登陆 http://staging.datapipeline.com:3000/alerting/list 查看报警情况；
3. 修复报警引起的问题；
4. 登陆http://tonggao.baidu.com/console 将报警标注已解决。


#### 报警流程
    ONCALL  -->  CEO
###### 流程解释
    当平台出现问题，第一时间打电话给ONCALL的人，如果ONCALL的人没有接通(包括挂断)，会打电话给CEO

#### 源码编译
    make build
