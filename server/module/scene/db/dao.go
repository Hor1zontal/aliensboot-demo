/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/4/12
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package db

import "github.com/KylinHe/aliensboot-core/mmo"

type Entity struct {
	ID 		mmo.EntityID  `bson:"_id"`
	Type 	mmo.EntityType  `bson:"type"`
	Data 	[]byte  `bson:"data"`
}