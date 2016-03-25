package game

import "log"
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



func LookupGameAccount(gameId bson.ObjectId, streamId bson.ObjectId) GameAccount {
    session := database.Dial()
    c := session.DB("data").C("game_account")

    gameAccount := GameAccount{}

    findErr := c.Find(bson.M{"_id": streamId, "game_id": gameId}).One(&gameAccount)
    if findErr != nil {
        log.Print("Could not find Game Account")
        panic(findErr)
    }

    return gameAccount
}

