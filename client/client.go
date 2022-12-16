package client

import (
	"github.com/jinzhu/gorm"
)

var Global = struct {
	UserConf CostVar `confs:"read=rainbow;format=json"`
	Mini            gorm.DB  `confs:"name=Mini;read=rainbow"`
	PartServiceConf PartServiceConf `confs:"read=rainbow;format=json"`
}{}

//CostVar 定义变量
type CostVar struct {
	JwtKey string
}


type PartServiceConf struct {
	HttpPort string
	HttpIp   string
	RpcPort  string
	RpcIp    string
}
