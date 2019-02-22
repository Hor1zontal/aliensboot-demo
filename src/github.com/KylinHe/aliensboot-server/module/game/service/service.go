// Code generated by aliensbot. DO NOT EDIT.
// source: module game
package service

import (
    "github.com/KylinHe/aliensboot-core/chanrpc"
    "github.com/KylinHe/aliensboot-core/cluster/center/service"
    "github.com/KylinHe/aliensboot-core/cluster/center"
    "github.com/KylinHe/aliensboot-core/exception"
    "github.com/KylinHe/aliensboot-core/protocol/base"
    "github.com/KylinHe/aliensboot-server/module/game/conf"
    "github.com/KylinHe/aliensboot-server/protocol"
    "github.com/gogo/protobuf/proto"
)

var instance service.IService = nil

var handlers = make(map[uint16]func(request *base.Any)*base.Any)

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}


//register self handler
func RegisteHandler(msgID uint16, handler func(request *base.Any)*base.Any) {
	handlers[msgID] = handler
}

func handleInternal(request *base.Any) (bool, *base.Any) {
	handler := handlers[request.Id]
	if handler == nil {
		return false, nil
	}
	response := handler(request)
	return true, response
}

func handle(request *base.Any) (response *base.Any) {
    ok, response := handleInternal(request)
	if ok {
		return response
	}
	requestProxy := &protocol.Request{}
	responseProxy := &protocol.Response{}
	response = &base.Any{}
	isResponse := false
	defer func() {
		//处理消息异常
		if err := recover(); err != nil {
			switch err.(type) {
			case protocol.Code:
				responseProxy.Code = err.(protocol.Code)
				break
			default:
				exception.PrintStackDetail(err)
				responseProxy.Code = protocol.Code_ServerException
			}
			isResponse = true
		}
		if !isResponse {
            return
        }
		data, _ := proto.Marshal(responseProxy)
		responseProxy.Session = requestProxy.GetSession()
		response.Value = data
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		exception.GameException(protocol.Code_InvalidRequest)
	}
	isResponse = handleRequest(request.GetAuthId(), request.GetGateId(), requestProxy, responseProxy)
	return
}

func handleRequest(authID int64, gateID string, request *protocol.Request, response *protocol.Response) bool {
	
	if request.GetGetRoleInfo() != nil {
		messageRet := &protocol.GetRoleInfoRet{}
		handleGetRoleInfo(authID, gateID, request.GetGetRoleInfo(), messageRet)
		response.Game = &protocol.Response_GetRoleInfoRet{messageRet}
		return true
	}
	
	if request.GetLoginRole() != nil {
		messageRet := &protocol.LoginRoleRet{}
		handleLoginRole(authID, gateID, request.GetLoginRole(), messageRet)
		response.Game = &protocol.Response_LoginRoleRet{messageRet}
		return true
	}
	
	if request.GetChangeNickname() != nil {
		messageRet := &protocol.ChangeNicknameRet{}
		handleChangeNickname(authID, gateID, request.GetChangeNickname(), messageRet)
		response.Game = &protocol.Response_ChangeNicknameRet{messageRet}
		return true
	}
	
	
	response.Code = protocol.Code_InvalidRequest
	return true
}

