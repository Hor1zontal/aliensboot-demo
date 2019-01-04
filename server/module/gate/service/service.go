// Code generated by aliensbot. DO NOT EDIT.
// source: module gate
package service

import (
	"github.com/KylinHe/aliensboot-core/chanrpc"
	"github.com/KylinHe/aliensboot-core/cluster/center"
	"github.com/KylinHe/aliensboot-core/cluster/center/service"
	"github.com/KylinHe/aliensboot-core/exception"
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/protocol/base"
	"github.com/KylinHe/aliensboot-server/module/gate/conf"
	"github.com/KylinHe/aliensboot-server/protocol"
	"github.com/gogo/protobuf/proto"
)

var instance service.IService = nil

func Init(chanRpc *chanrpc.Server) {
	instance = center.PublicService(conf.Config.Service, service.NewRpcHandler(chanRpc, handle))
}

func Close() {
	center.ReleaseService(instance)
}

func handle(request *base.Any) *base.Any {
	requestProxy := &protocol.Request{}
	defer func() {
		exception.CatchStackDetail()
	}()
	error := proto.Unmarshal(request.Value, requestProxy)
	if error != nil {
		log.Debugf("un expect request data : %v", request)
		return nil
	}
	handleRequest(requestProxy)
	return nil
}

func handleRequest(request *protocol.Request) {

	if request.GetHeartBeat() != nil {
		handleHeartBeat(request.GetHeartBeat())
		return
	}

	if request.GetBindService() != nil {
		handleBindService(request.GetBindService())
		return
	}

	if request.GetKickOut() != nil {
		handleKickOut(request.GetKickOut())
		return
	}

	if request.GetPushMessage() != nil {
		handlePushMessage(request.GetPushMessage())
		return
	}

	exception.GameException(protocol.Code_InvalidRequest)
}
