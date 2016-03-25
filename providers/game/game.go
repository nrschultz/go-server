package game

import "gopkg.in/mgo.v2/bson"
import "encoding/json"
import "github.com/nrschultz/go-server/providers/shared"


func StatPayload(gameId string, streamId string, teamId string, statsRequested string, qualifyingStat string) GameJsonPayload {
    gameAccount := LookupGameAccount(bson.ObjectIdHex(gameId), bson.ObjectIdHex(streamId))
    qualifyingIdentifier := shared.StatIdentifier{}
    requestedIdentifiers := []shared.StatIdentifier{}
    qualErr := json.Unmarshal([]byte(qualifyingStat), &qualifyingIdentifier)
    if qualErr != nil {
        panic(qualErr)
    }
    rqErr := json.Unmarshal([]byte(statsRequested), &requestedIdentifiers)
    if rqErr != nil {
        panic(rqErr)
    }

    return gameAccount.TransformToPayload(bson.ObjectIdHex(teamId), qualifyingIdentifier, requestedIdentifiers)
}

