package shared

import "log"
import "errors"
import "github.com/nrschultz/go-server/database"
import "gopkg.in/mgo.v2/bson"

type Stat struct {
    Key string        `bson:"s"`
    Value float64     `bson:"v"`
}

type CategorizedStats struct {
    Id      bson.ObjectId   `bson:"id"`
    Offense []Stat          `bson:"offense"`
    Defense []Stat          `bson:"defense"`
    General []Stat          `bson:"general"`
}

func (categorizedStats CategorizedStats) GetStatValue(identifier StatIdentifier) float64 {
    statList := []Stat{}
    switch {
    case identifier.Category == "offense":
        statList = categorizedStats.Offense
        break
    case identifier.Category == "defense":
        statList = categorizedStats.Defense
        break
    case identifier.Category == "general":
        statList = categorizedStats.General
        break
    }
    for statIndex := range statList {
        stat := statList[statIndex]
        if stat.Key == identifier.Key {
            return stat.Value
        }
    }
    return float64(0)
}


type RosterPlayer struct {
    Id          bson.ObjectId   `bson:"player_id"`
    FirstName   string          `bson:"fname"`
    LastName    string          `bson:"lname"`
    Number      string          `bson:"num"`
}

func (player RosterPlayer) FullName() string {
    return player.FirstName + " " + player.LastName
}

type Team struct {
    Id      bson.ObjectId   `bson:"_id"`
    Roster  []RosterPlayer  `bson:"roster"`
    Dropped []RosterPlayer  `bson:"dropped"`
}

func (team Team) HasPlayer(playerId bson.ObjectId) bool {
    _, err := team.GetPlayer(playerId)
    if err == nil {
        return true
    } else {
        return false
    }
}

func (team Team) GetPlayer(playerId bson.ObjectId) (RosterPlayer, error) {
    for _, player := range team.Roster {
        if player.Id == playerId {
            return player, nil
        }
    }
    for _, player := range team.Dropped {
        if player.Id == playerId {
            return player, nil
        }
    }
    return RosterPlayer{FirstName:"hey"}, errors.New("could not find player on team")
}

func LookupTeam(teamId bson.ObjectId) Team {
    session := database.Dial()
    c := session.DB("data").C("team")

    team := Team{}

    findErr := c.Find(bson.M{"_id": teamId}).One(&team)
    if findErr != nil {
        log.Print("Could not find Team")
        panic(findErr)
    }

    return team
}
