package shared

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
