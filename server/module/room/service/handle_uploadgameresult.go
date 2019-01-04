// Code generated by aliensbot. DO NOT EDIT.
// source: room_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-server/module/room/core"
	"github.com/KylinHe/aliensboot-server/protocol"
)

//
func handleUploadGameResult(authID int64, gateID string, request *protocol.UploadGameResult) {
	game := core.RoomManager.GetRoomByPlayerID(authID)
	game.UploadResult(authID, request.GetDetail())
}
