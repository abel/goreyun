package goreyun

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

//REST SDK URL
//http://game.reyun.com/restdoc

var (
	host = "http://log.reyun.com/receive/rest/"
	//需要改为自己申请的Appid
	Appid = "34fa5a932c970b5cd2344b36addbb615"
)

type ReyunLog interface {
	Name() string
}

//统计玩家第一次打开应用方法
type EquipContext struct {
	Deviceid  string `json:"deviceid"`  //设备ID(128)
	Channelid string `json:"channelid"` //渠道ID(最长16)
	Idfa      string `json:"idfa"`      //广告标识符
	Idfv      string `json:"idfv"`      //Vindor标示符
	Ip        string `json:"ip"`        //IP
}

type UserContext struct {
	EquipContext `json:",inline"`
	Serverid     string `json:"serverid"` //服务器ID
	Level        string `json:"level"`    //账户等级
}

type InstallRest struct {
	Appid   string        `json:"appid"`   //应用appid
	Context *EquipContext `json:"context"` //上下文信息
}

func (self *InstallRest) Name() string {
	return "install"
}

//统计玩家打开应用方法
type StartupContext struct {
	*EquipContext `json:",inline"`
	Tz            string `json:"tz"`         //时区,默认：+8
	Devicetype    string `json:"devicetype"` //设备类型
	Op            string `json:"op"`         //运营商
	Network       string `json:"network"`    //网络制式2G,3G,WIFI
	Os            string `json:"os"`         //操作系统
	Resolution    string `json:"resolution"` //分辨率
}

type StartupRest struct {
	Appid   string         `json:"appid"`   //应用appid
	Context StartupContext `json:"context"` //上下文信息
}

func (self *StartupRest) Name() string {
	return "startup"
}

//4.统计玩家账号注册方法
type RegisterContext struct {
	*EquipContext `json:",inline"`
	Accounttype   string `json:"accounttype"` //账户类型
	Gender        string `json:"gender"`      //账户性别(男：m女：f 其他：o)
	Age           string `json:"accounttype"` //账户年龄(年龄范围：0-120 默认：-1)
	Serverid      string `json:"serverid"`    //服务器ID
}

type RegisterRest struct {
	Appid   string          `json:"appid"`   //应用appid
	Who     string          `json:"who"`     //账户ID
	Context RegisterContext `json:"context"` //上下文信息
}

func (self *RegisterRest) Name() string {
	return "register"
}

//5.统计玩家登陆方法
type LoggedinRest struct {
	Appid   string       `json:"appid"`   //应用appid
	Who     string       `json:"who"`     //账户ID
	Context *UserContext `json:"context"` //上下文信息
}

func (self *LoggedinRest) Name() string {
	return "loggedin"
}

//6.统计玩家充值方法
type PaymentContext struct {
	*UserContext      `json:",inline"`
	Transactionid     string `json:"transactionid"`     //交易的流水号
	Paymenttype       string `json:"paymenttype"`       //支付类型，例如支付宝，银联，苹果、谷歌官方等,如果是系统赠送的，paymentType为：free
	Currencytype      string `json:"currencytype"`      //货币类型，按照国际标准组织ISO 4217中规范的3位字母，例如CNY人民币、USD美金等
	Currencyamount    string `json:"currencyamount"`    //支付的真实货币的金额
	Virtualcoinamount string `json:"virtualcoinamount"` //通过充值获得的游戏内货币的数量
	Iapname           string `json:"iapname"`           //游戏内购买道具的名称
	Iapamount         string `json:"iapamount"`         //游戏内购买道具的数量
}

type PaymentRest struct {
	Appid   string         `json:"appid"`   //应用appid
	Who     string         `json:"who"`     //账户ID
	Context PaymentContext `json:"context"` //上下文信息
}

func (self *PaymentRest) Name() string {
	return "payment"
}

//7.统计玩家在游戏内虚拟交易的方法
type EconomyContext struct {
	*UserContext
	Itemname       string `json:"itemname"`       //游戏内虚拟物品的名称/ID
	Itemamount     string `json:"itemamount"`     //交易的数量
	Itemtotalprice string `json:"itemtotalprice"` //交易的总价
}

type EconomyRest struct {
	Appid   string         `json:"appid"`   //应用appid
	Who     string         `json:"who"`     //账户ID
	Context EconomyContext `json:"context"` //上下文信息
}

func (self *EconomyRest) Name() string {
	return "economy"
}

//8.统计玩家的任务/关卡/副本方法
type QuestContext struct {
	*UserContext
	Questid     string `json:"questid"`     //当前任务/关卡/副本的编号或名称
	Queststatus string `json:"queststatus"` //当前任务/关卡/副本的状态，有如下三种类型：开始：a 完成：c 失败：f
	Questtype   string `json:"questtype"`   //当前任务/关卡/副本的类型，例如： 新手任务：new 主线任务：main 支线任务：sub 开发者也可以根据自己游戏的特点自定义类型
}

type QuestRest struct {
	Appid   string       `json:"appid"`   //应用appid
	Who     string       `json:"who"`     //账户ID
	Context QuestContext `json:"context"` //上下文信息
}

func (self *QuestRest) Name() string {
	return "quest"
}

//9.统计玩家自定义事件方法
type EventContext struct {
	*UserContext
	User_define1 string `json:"user-define1"` //用户自定义
	User_define2 string `json:"user-define2"` //用户自定义
}

type EventRest struct {
	Appid   string       `json:"appid"`   //应用appid
	Who     string       `json:"who"`     //账户ID
	What    string       `json:"what"`    //自定义事件的名称
	Context EventContext `json:"context"` //上下文信息
}

func (self *EventRest) Name() string {
	return "event"
}

//10.统计玩家在线的方法
type HeartbeatRest struct {
	Appid   string       `json:"appid"`   //应用appid
	Who     string       `json:"who"`     //账户ID
	Context *UserContext `json:"context"` //上下文信息
}

func (self *HeartbeatRest) Name() string {
	return "heartbeat"
}

type Replay struct {
	Status int               `json:"status"`
	Result map[string]string `json:"result"`
}

func PostUrl(addr string, body io.Reader) string {
	client := http.Client{}
	reqest, err := http.NewRequest("POST", addr, body)
	if err != nil {
		return err.Error()
	}
	reqest.Header.Add("Content-Type", "application/json")
	reqest.Header.Add("Accept-Encoding", "deflate")
	//reqest.Header.Add("Connection", "close")
	response, err := client.Do(reqest)
	if err != nil {
		return err.Error()
	}
	var resp string
	defer response.Body.Close()
	bodyByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		resp = err.Error()
	} else {
		resp = string(bodyByte)
	}
	return resp
}

func PostLog(log ReyunLog) string {
	bin, err := json.Marshal(log)
	if err != nil {
		return err.Error()
	}
	buffer := bytes.NewReader(bin)
	return PostUrl(host+log.Name(), buffer)
}
