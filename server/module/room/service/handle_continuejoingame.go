// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)


//
func handleContinueJoinGame(authID int64, gateID string, request *protocol.ContinueJoinGame) {
	room := core.RoomManager.GetRoomByPlayerID(authID)
	code := request.GetCode()
	//同意
	if code == 0 {
		room.AcceptJoinGame(authID, request.GetPlayerID())
	}
}
