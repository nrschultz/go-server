package game

import "gopkg.in/mgo.v2/bson"
import "github.com/nrschultz/go-server/providers/shared"

type PlayerRowInfo struct {
    JerseyNumber    string          `json:"jersey_number"`
    Id              bson.ObjectId   `json:"id"`
    Name            string          `json:"name"`
}

type PlayerRowPayload struct {
    RowInfo PlayerRowInfo       `json:"row_info"`
    Stats []shared.PayloadStat  `json:"stats"`
}

type Totals struct {
    Stats []shared.PayloadStat `json:"stats"`
}

type GameJsonPayload struct {
    Players             []PlayerRowPayload      `json:"players"`
    AdditionalPlayers   []PlayerRowPayload      `json:"additional_players"`
    Totals              Totals                  `json:"totals"`
    Glossary            []shared.GlossaryEntry  `json:"glossary"`
}

func buildPlayerPayload(player shared.CategorizedStats, statsRequested []shared.StatIdentifier) PlayerRowPayload {
    playerPayload := PlayerRowPayload{}
    playerPayload.RowInfo = PlayerRowInfo{JerseyNumber: "8", Id: player.Id, Name: "Nick Schultz"}
    playerPayload.Stats = shared.BuildStatList(player, statsRequested)
    return playerPayload
}

func (account GameAccount) TransformToPayload(teamId bson.ObjectId, qualifyingStat shared.StatIdentifier, statsRequested []shared.StatIdentifier) GameJsonPayload {
    payload := GameJsonPayload{}
    payload.Players = []PlayerRowPayload{}
    for playerIndex := range account.Stats.Player {
        player := account.Stats.Player[playerIndex]
        if player.GetStatValue(qualifyingStat) > 0 {
            playerPayload := buildPlayerPayload(player, statsRequested)
            payload.Players = append(payload.Players, playerPayload)
        }
    }
    for teamIndex := range account.Stats.Team {
        team := account.Stats.Team[teamIndex]
        if team.Id == teamId {
            payload.Totals = Totals{Stats: shared.BuildStatList(team, statsRequested)}
            break
        }
    }
    return payload
}
