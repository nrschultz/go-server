package game

import "gopkg.in/mgo.v2/bson"


func StatPayload(gameAccountId string, teamId string) GameJsonPayload {
    gameAccount := Lookup(gameAccountId)
    return gameAccount.TransformToPayload(bson.ObjectIdHex(teamId))
}

