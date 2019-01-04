// Code generated by aliensbot. DO NOT EDIT.
// source: scene_interface.proto
package service

import (
	"github.com/KylinHe/aliensboot-core/log"
	"github.com/KylinHe/aliensboot-core/mmo"
	"github.com/KylinHe/aliensboot-server/protocol"
)




//
func handleMigrateIn(authID int64, gateID string, request *protocol.MigrateIn) {
	err := mmo.MigrateIn(mmo.EntityID(request.GetSpaceID()), mmo.EntityID(request.GetEntityID()), request.GetData())
	if err != nil {
		log.Errorf("handle migrateIn error : %v", err)
	}
}
