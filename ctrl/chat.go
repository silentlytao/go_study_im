package ctrl

import (
	"../model"
	"../service"
	"encoding/json"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
)

const (
	//点对点单聊 dstid 用户id
	CMD_SINGLE_MSG = 10
	//群聊 dstid 群id
	CMD_ROOM_MSG = 11
	//心跳消息，不处理
	CMD_HEART = 0
)

//本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

//映射关系表
var clientMap = make(map[int64]*Node, 0)

//读写锁
var rwLocker sync.RWMutex

var contactService service.ContactService

var userService service.UserService

func init()  {
	go udpSendProc()
	go udpRecvProc()
}

func Chat(writer http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	id := query.Get("id")
	token := query.Get("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	isValidate := checkToken(userId, token)

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isValidate
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//todo 获得conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	//todo 获取用户全部群id
	comIds := contactService.SearchCommunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}

	//todo userid和node形成绑定关系
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()
	//todo 完成发送逻辑,con
	go sendProc(node)
	//todo 完成接收逻辑
	go recvProc(node)
	log.Printf("<-%d\n", userId)
	sendMsg(userId, []byte("hello,world!"))
}

//ws发送协程
func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//ws接收协程
func recvProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//dispatch(data)
		broadMsg(data)
		log.Printf("[ws]<=%s\n", data)
	}
}
//发送通道
var udpSendChan chan []byte = make(chan[] byte,1024)
//广播
func broadMsg(data []byte)  {
	udpSendChan <- data
}

//udp数据的发送协程
func udpSendProc()  {
	log.Println("start udpSendProc")
	con,err := net.DialUDP("udp",nil,&net.UDPAddr{
		IP: net.IPv4(192,168,10,255),
		Port: 3000,
	})
	defer con.Close()
	if err != nil{
		log.Println(err.Error())
		return
	}
	//通过con 发送消息
	for  {
		select {
			case data := <- udpSendChan:
				log.Println("udp send data",data)
				_,err = con.Write(data)
				if err != nil{
					log.Println(err.Error())
					return
				}
		}
	}
}

//udp数据的接收协程
func udpRecvProc()  {
	log.Println("start udpRecvProc")
	//监听udp广播端口
	con,err := net.ListenUDP("udp",&net.UDPAddr{
		IP: net.IPv4zero,
		Port: 3000,
	})
	defer con.Close()
	if err != nil{
		log.Println(err.Error())
		return
	}
	//处理端口发送过来的数据
	for  {
		var buf [512]byte
		n,err := con.Read(buf[0:])
		if err != nil{
			log.Println(err.Error())
			return
		}
		log.Println("udp recv data",buf[0:n])
		dispatch(buf[0:n])
	}
	log.Println("stop udpRecvProc")
}

//分发
func dispatch(data []byte) {
	//todo 解析data为message
	msg := model.Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//todo 根据cmd对逻辑进行处理
	switch msg.Cmd {
	case CMD_SINGLE_MSG:
		sendMsg(msg.Dstid, data)
	case CMD_ROOM_MSG:
		//todo 群聊转发逻辑
		for userId, v := range clientMap {
			//排除当前发送用户的所有群id
			if v.GroupSets.Has(msg.Dstid) && userId != msg.Userid {
				v.DataQueue <- data
			}
		}
	case CMD_HEART:
		//todo 一般啥都不做
	}
}

//todo 发送消息
func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

//验证token是否有效
func checkToken(userId int64, token string) bool {
	user := UserService.Find(userId)
	return user.Token == token
}
