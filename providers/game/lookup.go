package game

import "github.com/nrschultz/go-server/database"
import "github.com/nrschultz/go-server/providers/shared"
import "gopkg.in/mgo.v2/bson"


type GameAccountStats struct {
    Player  []shared.CategorizedStats `bson:"player"`
    Team    []shared.CategorizedStats `bson:"team"`
}

type GameAccount struct {
    Id      bson.ObjectId       `bson:"_id"`
    Stats   GameAccountStats    `bson:"stats"`
}


func Lookup(gameAccountId string) GameAccount {
    session := database.Dial()
    c := session.DB("data").C("game_account")

    gameAccount := GameAccount{}

    findErr := c.Find(bson.M{"_id": bson.ObjectIdHex(gameAccountId)}).One(&gameAccount)
    if findErr != nil {
        panic(findErr)
    }

    return gameAccount
}
