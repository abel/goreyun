package goreyun

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeParsse(t *testing.T) {
	var start int32
	ti, err := time.Parse("2006-01-02 15:04:05", "2016-05-20 06:00:00")
	if err == nil {
		start = int32(ti.Unix()) - (8 * 3600)
	} else {
		fmt.Printf("%v", err.Error())
	}

	Convey("time test", t, func() {
		Convey("TimeParsse", func() {
			So(start, ShouldEqual, 1463695200)
		})
	})
}

const (
	Success     = "{\"status\":0}"
	NameHead    = "test00"
	DevicedHead = "adsfewrwe233434dsfa"
	IPHead      = "192.168.188,1"
	UserCount   = 10
)

func createEquipContext(i int) *EquipContext {
	return &EquipContext{
		Deviceid:  DevicedHead + fmt.Sprint(i),
		Channelid: "qq",
		Ip:        IPHead + fmt.Sprint(i),
		Idfa:      "",
		Idfv:      "",
	}
}

func createUserContext(i int) *UserContext {
	u := &UserContext{
		Serverid: "1",
	}
	u.Deviceid = DevicedHead + fmt.Sprint(i)
	u.Channelid = "qq"
	u.Ip = IPHead + fmt.Sprint(i)
	return u
}

func createDevicetype() string {
	devicetype := []string{
		//"iphone4s",
		//"iphone5s",
		"sansung-GT9300",
		"xiaomi2",
		"xiaomi2s",
		"xiaomi3",
		"xiaomi4",
		"xiaomi4s",
		"xiaomi5",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]
}

func createOP() string {
	devicetype := []string{
		"中国移动",
		"中国联通",
		"中国电信",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]
}

func createNetwork() string {
	devicetype := []string{
		"2G",
		"3G",
		"WIFI",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]
}

func createResolution() string {
	devicetype := []string{
		"800*600",
		"1280*720",
		"1920*1080",
		"2048×1536",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]
}

func createOS() string {
	devicetype := []string{
		"android4.0",
		"android5.0",
		"android6.0",
		//"ios7.0",
		//"ios8.0",
		//"ios9.0",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]
}

func createPaymenttype() string {
	devicetype := []string{
		"alipay",
		"银联",
		"apple",
		"google",
		"weixin",
	}
	r := rand.Intn(len(devicetype))
	return devicetype[r]

}

//func TestInstall(t *testing.T) {
//	var resp string
//	for i := 0; i < UserCount; i++ {
//		req := InstallRest{
//			Appid:   Appid,
//			Context: createEquipContext(i),
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Install", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}

//func TestStartup(t *testing.T) {
//	var resp string
//	for i := 0; i < UserCount; i++ {
//		context := StartupContext{
//			EquipContext: createEquipContext(i),
//			Tz:           "+8",
//			Devicetype:   createDevicetype(),
//			Op:           createOP(),
//			Network:      createNetwork(),
//			Os:           createOS(),
//			Resolution:   createResolution(),
//		}
//		req := StartupRest{
//			Appid:   Appid,
//			Context: context,
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Startup", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}

//func TestRegister(t *testing.T) {
//	var resp string
//	//男：m女：f
//	gender := []string{"m", "f"}
//	for i := 0; i < UserCount; i++ {
//		context := RegisterContext{
//			EquipContext: createEquipContext(i),
//			Accounttype:  "test",
//			Gender:       gender[i%2],
//			Age:          fmt.Sprint(18 + i),
//			Serverid:     "1",
//		}
//		req := RegisterRest{
//			Appid:   Appid,
//			Who:     NameHead + fmt.Sprint(i),
//			Context: context,
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Register", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}

//func TestLoggedin(t *testing.T) {
//	var resp string
//	for i := 0; i < UserCount; i++ {
//		context := LoggedinContext{
//			UserContext: createUserContext(i),
//		}
//		req := LoggedinRest{
//			Appid:   Appid,
//			Who:     NameHead + fmt.Sprint(i),
//			Context: context,
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Loggedin", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}

//func createTransactionid(i int) string {
//	t := time.Now().UnixNano()
//	return fmt.Sprint(t) + fmt.Sprint(i)
//}

//func TestPayment(t *testing.T) {
//	var resp string
//	for i := 0; i < UserCount; i++ {
//		context := PaymentContext{
//			UserContext:       createUserContext(i),
//			Transactionid:     createTransactionid(i),     //交易的流水号
//			Paymenttype:       createPaymenttype(),        //支付类型，例如支付宝，银联，苹果、谷歌官方等,如果是系统赠送的，paymentType为：free
//			Currencytype:      "CNY",                      //货币类型，按照国际标准组织ISO 4217中规范的3位字母，例如CNY人民币、USD美金等
//			Currencyamount:    fmt.Sprint(10 * (i + 1)),   //支付的真实货币的金额
//			Virtualcoinamount: fmt.Sprint(1000 * (i + 1)), //通过充值获得的游戏内货币的数量
//			Iapname:           "Vip" + fmt.Sprint(i),      //游戏内购买道具的名称
//			Iapamount:         "1",                        //游戏内购买道具的数量
//		}
//		req := PaymentRest{
//			Appid:   Appid,
//			Who:     NameHead + fmt.Sprint(i),
//			Context: context,
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Payment", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}

//func TestEconomy(t *testing.T) {
//	var resp string
//	for i := 0; i < UserCount; i++ {
//		context := EconomyContext{
//			UserContext:    createUserContext(i),
//			Itemname:       "超级武器" + fmt.Sprint(i),             //游戏内虚拟物品的名称/ID
//			Itemamount:     fmt.Sprint(10 * (i + 1)),           //交易的数量
//			Itemtotalprice: fmt.Sprint(10 * (i + 1) * (i * 5)), //交易的总价
//		}
//		req := EconomyRest{
//			Appid:   Appid,
//			Who:     NameHead + fmt.Sprint(i),
//			Context: context,
//		}
//		resp = PostLog(&req)
//	}
//	Convey("reyun test", t, func() {
//		Convey("Economy", func() {
//			So(resp, ShouldEqual, Success)
//		})
//	})
//}
