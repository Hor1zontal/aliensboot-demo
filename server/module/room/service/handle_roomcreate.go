// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)


//
func handleRoomCreate(authID int64, gateID string, request *protocol.RoomCreate, response *protocol.RoomCreateRet) {
	room := core.RoomManager.CreateRoom(request.GetAppID(), authID, request.GetRoomID(), request.GetForce(), request.GetMaxSeat())
	response.RoomID = room.GetID()
	response.Players = room.GetAllPlayerData()
}

