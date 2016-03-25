package game

import "github.com/nrschultz/go-server/database"
import "gopkg.in/mgo.v2/bson"

type Stat struct {
    S string        `bson:"s"`
    V float64       `bson:"v"`
}

type CategorizedStats struct {
    Id      bson.ObjectId   `bson:"id"`
    Offense []Stat          `bson:"offense"`
    Defense []Stat          `bson:"defense"`
    General []Stat          `bson:"general"`
}

type GameAccountStats struct {
    Player  []CategorizedStats `bson:"player"`
    Team    []CategorizedStats `bson:"team"`
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
