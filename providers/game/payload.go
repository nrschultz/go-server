package game

import "math"
import "gopkg.in/mgo.v2/bson"
import "strconv"
import "github.com/nrschultz/go-server/providers/shared"

type JsonFloat64 float64

func (value JsonFloat64) MarshalJSON() ([]byte, error) {
    floatValue := float64(value)
    vStr := ""
    if floatValue == math.Inf(1) {
        vStr = "inf"
    } else if floatValue == math.Inf(-1) {
        vStr = "-inf"
    } else {
        vStr = strconv.FormatFloat(floatValue, 'f', -1, 64)
    }
    return []byte(`"`+ vStr +`"`), nil
}

type StatIdentifier struct {
    Category    string  `json:"category"`
    Key         string  `json:"key"`
}

type StatFormat struct {
    RemoveZero bool     `json:"remove_zero"`
    FormatType string   `json:"type"`
    Precision  int      `json:"precision"`
}

type PayloadStat struct {
    Identifier  StatIdentifier  `json:"identifier"`
    Value       JsonFloat64     `json:"value"`
    Format      StatFormat      `json:"format"`
}

type PlayerRowInfo struct {
    JerseyNumber    string          `json:"jersey_number"`
    Id              bson.ObjectId   `json:"id"`
    Name            string          `json:"name"`
}

type PlayerRowPayload struct {
    RowInfo PlayerRowInfo   `json:"row_info"`
    Stats []PayloadStat     `json:"stats"`
}

type GlossaryEntry struct {
    Identifier  StatIdentifier  `json:"identifier"`
    Link        string          `json:"link"`
    Name        string          `json:"name"`
    Key         string          `json:"key"`
    Description string          `json:"description"`
}

type Totals struct {
    Stats []PayloadStat `json:"stats"`
}

type GameJsonPayload struct {
    Players             []PlayerRowPayload  `json:"players"`
    AdditionalPlayers   []PlayerRowPayload  `json:"additional_players"`
    Totals              Totals              `json:"totals"`
    Glossary            []GlossaryEntry     `json:"glossary"`
}



func buildPlayerPayload(player CategorizedStats) PlayerRowPayload {
    playerPayload := PlayerRowPayload{}
    playerPayload.RowInfo = PlayerRowInfo{JerseyNumber: "8", Id: player.Id, Name: "Nick Schultz"}
    playerPayload.Stats = shared.BuildStatList(player)
    return playerPayload
}

func (account GameAccount) TransformToPayload(teamId bson.ObjectId) GameJsonPayload {
    payload := GameJsonPayload{}
    payload.Players = []PlayerRowPayload{}
    for playerIndex := range account.Stats.Player {
        playerPayload := buildPlayerPayload(account.Stats.Player[playerIndex])
        payload.Players = append(payload.Players, playerPayload)
    }
    for teamIndex := range account.Stats.Team {
        team := account.Stats.Team[teamIndex]
        if team.Id == teamId {
            payload.Totals = Totals{Stats: shared.BuildStatList(team)}
            break
        }
    }
    return payload
}
