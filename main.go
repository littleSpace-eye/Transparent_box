package main

import (
	"database/sql"
	_ "database/sql"
	"demo/datasource"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron/v3"
	"log"
	"math/big"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func auth() *bind.TransactOpts {
	const key = `{"address":"c36d5b457614e0e9a4dcc3960d153ee520dcdd7f","crypto":{"cipher":"aes-128-ctr","ciphertext":"53b60a72e016c7d47615665434ee5c38e96192f5fa48f163cb3e7effd0a8233d","cipherparams":{"iv":"7bdd4f6cc89f17ce5baec70b6066ee60"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"2753882645e277476a92ab5d51e9832c18b007f61bb2e9d59c757ddf6cb392b6"},"mac":"2a5a9f2f864b3c7984b7e3d7d4d7f49e8830aded7b555d48cecb359c0d913fdc"},"id":"35ceec31-4ffc-4dd9-9ac6-765ca57fa187","version":3}`
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), "123", big.NewInt(666))
	if err != nil {
		fmt.Println("could not create auth")
	}
	auth.GasLimit = 3000000
	auth.GasPrice = big.NewInt(15000000)
	return auth
}

var getId int
var getAddress string
var getIP string
var getTimes int
var getToday_Account float32
var getAll_Account float32
var getAll_Times int
var getSumDay_Account float32
var getSumIP_Accoount float32

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		fmt.Println("数据库连接失败")
	}

	// 初始化数据库
	datasource.Init()

	engine := gin.Default()
	conn, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		panic("链接以太坊合约出错")
	}
	fmt.Println(conn)
	address := "0x68536aC63d72A29e54b8787617531cfd1F7eD4Ab"
	abi, err := NewABi(common.HexToAddress(address), conn)
	if err != nil {
		fmt.Println(err)
	}

	engine.LoadHTMLGlob("static/*")
	engine.Static("/img", "./img")

	engine.GET("/", func(context *gin.Context) {
		var member []datasource.Record
		rows, err := db.Query("select * from record;")
		if err != nil {
			fmt.Println(err)
		}
		for rows.Next() {
			var record datasource.Record
			err := rows.Scan(&record.ID, &record.ID, &record.Name, &record.Tel, &record.E_mail, &record.Sex, &record.Password)
			if err != nil {
				log.Fatalln(err.Error())
			}
			//record.Today_Account /=1000000000000000
			//record.All_Account /=1000000000000000
			member = append(member, record)
			fmt.Println(member)

		}
		a, _ := abi.GetBalance(nil)
		fmt.Println(a)
		context.HTML(http.StatusOK, "index.html", gin.H{"res": member, "getbalance": a})

	})
	print("hello world")
	engine.POST("/send", func(context *gin.Context) {
		account := context.PostForm("address")
		money := context.PostForm("money")
		intMoney, _ := strconv.Atoi(money)
		cash := int64(intMoney)
		ip, _ := externalIP()
		err1 := db.QueryRow("select *,sum(today_account) as totalSum from record where Address =? group by id ;", account).Scan(&getId, &getAddress, &getIP, &getTimes, &getToday_Account, &getAll_Account, &getAll_Times, &getSumDay_Account)
		//水龙头服务每天最多转出100个测试以太币
		if getSumDay_Account > 1000000000000000000 {
			context.HTML(http.StatusOK, "alert.html", gin.H{})
			fmt.Println("错误在143")
			return
		} else {
			//每个地址只能申请10次测试币否则就不再受理
			if getTimes > 10 {
				context.HTML(http.StatusOK, "alert.html", gin.H{})
				fmt.Println("错误在149")
				return
			} else {
				//每个地址累计只能申请10000次测试币
				if getAll_Times > 10000 {
					context.HTML(http.StatusOK, "alert.html", gin.H{})
					fmt.Println("错误在155")
					return
				} else {
					//每个IP地址每天只能申请100次测试币
					db.QueryRow("select sum(times)  from record where IP = ?;", ip).Scan(&getSumIP_Accoount)
					if getSumIP_Accoount >= 100 {
						fmt.Println("错误在161")
						context.HTML(http.StatusOK, "alert.html", gin.H{})
						return
					} else {
						if err1 == sql.ErrNoRows {
							db.Exec("insert into record(Address, IP, times, today_account, all_account,all_times) values(?,?,?,?,?,?);", account, ip.String(), 1, money, money, 1)
						} else {
							_, err := db.Exec("update record set times = times+1,today_account = today_account+?,all_account = all_account +?,all_times = all_times+1 where id=? ;", money, money, getId)
							if err != nil {
								fmt.Println(err)
							}
						}
					}
				}
			}
		}

		a, err := abi.TransferToAnother(auth(), common.HexToAddress(account), big.NewInt(cash))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(a.Hash())

		context.Redirect(http.StatusMovedPermanently, "http://localhost:8000/")

	})
	//定时任务 每当晚上零点的时候更新数据库数据
	cron2 := cron.New() //创建一个cron实例

	//执行定时任务（每晚上十二点执行一次）
	cron2.AddFunc("* * */0 * * *", func() { db.Exec("update record set times =0,today_account=0 where times<>0;") })

	//启动/关闭
	cron2.Start()
	defer db.Close()
	engine.Run(":8000")
}
