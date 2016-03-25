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
    Players             []PlayerRowPayload              `json:"players"`
    AdditionalPlayers   []PlayerRowPayload              `json:"additional_players"`
    Totals              Totals                          `json:"totals"`
    Glossary            []shared.PayloadGlossaryEntry   `json:"glossary"`
}

func buildPlayerPayload(rosterPlayer shared.RosterPlayer, playerStats shared.CategorizedStats, statsRequested []shared.StatIdentifier, glossary []shared.PayloadGlossaryEntry) PlayerRowPayload {
    playerPayload := PlayerRowPayload{}
    playerPayload.RowInfo = PlayerRowInfo{JerseyNumber: rosterPlayer.Number, Id: playerStats.Id, Name: rosterPlayer.FullName()}
    playerPayload.Stats = shared.BuildStatList(playerStats, statsRequested, glossary)
    return playerPayload
}

func (account GameAccount) TransformToPayload(teamId bson.ObjectId, qualifyingStat shared.StatIdentifier, statsRequested []shared.StatIdentifier) GameJsonPayload {
    payload := GameJsonPayload{}
    payload.Players = []PlayerRowPayload{}
    payload.Glossary = shared.GetGlossary(statsRequested)
    team := shared.LookupTeam(teamId)

    for _, playerStats := range account.Stats.Player {
        if team.HasPlayer(playerStats.Id) && playerStats.GetStatValue(qualifyingStat) > 0 {
            rosterPlayer, _ := team.GetPlayer(playerStats.Id)
            playerPayload := buildPlayerPayload(rosterPlayer, playerStats, statsRequested, payload.Glossary)
            payload.Players = append(payload.Players, playerPayload)
        }
    }
    for _, team := range account.Stats.Team {
        if team.Id == teamId {
            payload.Totals = Totals{Stats: shared.BuildStatList(team, statsRequested, payload.Glossary)}
            break
        }
    }


    return payload
}
