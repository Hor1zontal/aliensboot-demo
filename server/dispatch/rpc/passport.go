// Code generated by aliensbot. DO NOT EDIT.
// source: module passport
package rpc

import (
	"github.com/KylinHe/aliensboot-server/protocol"
)

var Passport = &passportRPCHandler{&rpcHandler{name:"passport"}}


type passportRPCHandler struct {
	*rpcHandler
}


func (this *passportRPCHandler) UserRegister(node string, request *protocol.UserRegister) *protocol.UserRegisterRet {
	message := &protocol.Request{
		Passport:&protocol.Request_UserRegister{
			UserRegister:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetUserRegisterRet()
}

func (this *passportRPCHandler) UserLogin(node string, request *protocol.UserLogin) *protocol.UserLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_UserLogin{
			UserLogin:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetUserLoginRet()
}

func (this *passportRPCHandler) TokenLogin(node string, request *protocol.TokenLogin) *protocol.TokenLoginRet {
	message := &protocol.Request{
		Passport:&protocol.Request_TokenLogin{
			TokenLogin:request,
		},
	}
	messageRet := this.Request(node, message)
	return messageRet.GetTokenLoginRet()
}


