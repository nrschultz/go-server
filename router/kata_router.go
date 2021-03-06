package main

import (
    "io"
    "net/http"
    "strconv"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/nrschultz/go-server/database"
    "github.com/nrschultz/go-server/providers/game"
)


func count(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    dbName, _ := params["dbName"]
    collectionName, _ := params["collectionName"]
    io.WriteString(w, strconv.Itoa(database.Count(dbName, collectionName)))
}

func gameStats(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    gameId := params["gameId"]
    streamId := params["streamId"]
    teamId := params["teamId"]
    statsRequested := r.FormValue("stats_requested")
    qualifyingStat := r.FormValue("qualifying_stat")
    // m := make(map[string]string)
    payload := game.StatPayload(gameId, streamId, teamId, statsRequested, qualifyingStat)
    // io.WriteString(w, payload)
    json.NewEncoder(w).Encode(payload)
}

func main() {
    rtr := mux.NewRouter()
    rtr.HandleFunc("/count/{dbName}/{collectionName}", count)
    rtr.HandleFunc("/stats/game/{gameId}/stream/{streamId}/team/{teamId}/", gameStats)
    // rtr.HandleFunc("/stats/game/{gameId}/team/{teamId}/", gameStats)
    // TODO: Make it work without a streamId

    http.Handle("/", rtr)

    http.ListenAndServe(":80", nil)
}
