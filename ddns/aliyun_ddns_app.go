package main

import (
    alidns20150109 "github.com/alibabacloud-go/alidns-20150109/v2/client"
    openapi "github.com/alibabacloud-go/darabonba-openapi/client"
    "github.com/alibabacloud-go/tea/tea"
    "github.com/bitly/go-simplejson"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "strconv"
    "strings"
    "time"
)

/* 在这里添加自己的访问秘钥 */
//var accessKeyId string  = ""
//var accessSecret string = ""
//var recordId string     = "19966927349687296"
//var type_ string        = "AAAA"
//var rr string           = "nas"
var lang string         = "zh-CN"

var accessKeyId string
var accessSecret string
var recordId string
var type_ string
var rr_ string
var ttl_ int

// nas A 749428833990588416
// nas AAAA 19966927349687296
var userClientIp string = "101.88.247.188"

func main() {
    log.Println("启动参数列表2：", os.Args[3:])

    // ${AKID} ${AKSCT} ${RR} ${RECORDID} ${TYPE} ${TTL}
    if len(os.Args) != 7 {
        log.Fatal("请设置 1accessKeyId 2accessSecret 3rr 4recordId 5type 6ttl ")
    }
    accessKeyId     = os.Args[1]
    accessSecret    = os.Args[2]
    rr_             = os.Args[3]
    recordId        = os.Args[4]
    type_           = os.Args[5]
    ttl_,_          = strconv.Atoi(os.Args[6])
    UpdateAliyunDDNS()
}

func UpdateAliyunDDNS() {
    for true {
        UpdateAliyunDDNSSingle()
        time.Sleep(time.Duration(10 * time.Minute))
    }
}

func UpdateAliyunDDNSSingle() {
    log.Println()

    wanIpValue, _ := GetHostWanIp()
    log.Println("<---- 查得公网IP", wanIpValue)
    aliyunOriginIpValue, _ := GetDomainRecordInfo()
    log.Println("<---- 查得已配IP", aliyunOriginIpValue)

    if wanIpValue == "" {
        log.Println("公网IP为空", wanIpValue)
    } else if wanIpValue == aliyunOriginIpValue {
        log.Println("IP未发生变动，无需更新", wanIpValue)
    } else {
        SetDomainRecordInfo(wanIpValue)
        GetDomainRecordInfo()
    }

    log.Println()
}


/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient (accessKeyId *string, accessKeySecret *string) (_result *alidns20150109.Client, _err error) {
    config := &openapi.Config{
        // 您的AccessKey ID
        AccessKeyId: accessKeyId,
        // 您的AccessKey Secret
        AccessKeySecret: accessKeySecret,
    }
    // 访问的域名
    config.Endpoint = tea.String("alidns.cn-shanghai.aliyuncs.com")
    _result = &alidns20150109.Client{}
    _result, _err = alidns20150109.NewClient(config)
    return _result, _err
}

// nas.okzhang.com
func _main (args []*string) (_err error) {
    return nil
}

func SetDomainRecordInfo(ipWanValue string) (_err error) {
    log.Println(time.Now().Local(), "start 设置DNS配置")

    client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessSecret))
    if _err != nil {
        log.Println("初始化客户端失败...")
        return _err
    }

    updateDomainRecordRequest := &alidns20150109.UpdateDomainRecordRequest{}
    updateDomainRecordRequest.SetRecordId(recordId)
    updateDomainRecordRequest.SetRR(rr_)
    updateDomainRecordRequest.SetType(type_)
    updateDomainRecordRequest.SetValue(ipWanValue)
    updateDomainRecordRequest.SetTTL(int64(ttl_))

    // 复制代码运行请自行打印 API 的返回值
    resp, _err := client.UpdateDomainRecord(updateDomainRecordRequest)
    if _err != nil {
        log.Println(time.Now().Local(), "设置DNS失败", resp, _err)
        return _err
    } else {
        log.Println(time.Now().Local(), "设置DNS成功", resp.Body)
        return _err
    }
}

/**
 example:
 body": {
    "Status": "ENABLE",
    "RR": "nas",
    "RequestId": "FD51428C-1CC4-5F9E-B5B2-1248B909479D",
    "DomainName": "xxx.com",
    "TTL": 600,
    "Line": "default",
    "Locked": false,
    "Type": "AAAA",
    "Value": "240e:xxx:xxx:xxx:2233",
    "RecordId": "19966927349687296"
 }
 */
func GetDomainRecordInfo() (value string, _err error) {
    log.Println(time.Now().Local(), "start 获取DNS配置")

    client, _err := CreateClient(tea.String(accessKeyId), tea.String(accessSecret))
    if _err != nil {
        log.Println("初始化客户端失败...", _err)
        return "", _err
    }

    request := &alidns20150109.DescribeDomainRecordInfoRequest{Lang: tea.String(lang), UserClientIp: tea.String(userClientIp), RecordId: tea.String(recordId)}
    resp, _err := client.DescribeDomainRecordInfo(request)
    if _err != nil {
        log.Println("获取DNS配置失败", _err)
        return "", _err
    }  else {
        //fmt.Println(resp)
        //var status = resp.Body.Status
        var domainName  = tea.StringValue(resp.Body.DomainName)
        var rr          = tea.StringValue(resp.Body.RR)
        var type_       = tea.StringValue(resp.Body.Type)
        var value       = tea.StringValue(resp.Body.Value)
        var ttl_        = tea.Int64Value(resp.Body.TTL)
        log.Println(time.Now().Local(), "已获取DNS配置", rr + "." + domainName, " ----> ", value, "(" + type_ + ") ttl:", ttl_)
        return value, _err
    }
}

func GetHostWanIp() (value string, _err error){
    log.Println(time.Now().Local(), "start 获取公网IP")

    toolUrl := [] string {
        //"https://www.baidu.com",
        "https://api-ipv6.ip.sb/ip",
        "http://v6.ident.me/",
        "https://api6.ipify.org/",
        "https://ipv6.lookup.test-ipv6.com/ip/",
        //"http://checkipv6.dyndns.com/",
    }
    for _, url := range toolUrl {
        resp, _err := http.Get(url)

        if _err != nil || resp.StatusCode != 200{
            log.Println(_err)
            //return "", _err
        } else {
            defer resp.Body.Close()

            if url == toolUrl[len(toolUrl)-1] {
                json, _ := simplejson.NewFromReader(resp.Body)
                ip, _ := json.Get("ip").String()
                if ip != "" {
                    log.Println(time.Now().Local(), "查得IP ", url, " ----> " , ip)
                    return ip, nil
                }
            } else {
                body, _ := ioutil.ReadAll(resp.Body)
                resultBody := strings.ReplaceAll(string(body), "\n", "")

                log.Println(time.Now().Local(), "查得IP ", url, " ----> " , resultBody)
                return resultBody, nil
            }
        }
    }

    return "", _err
}