package database

import "os"
import "gopkg.in/mgo.v2"

func Dial() *mgo.Session {
    session, err := mgo.Dial(os.Getenv("MONGO_URI"))
    if err != nil {
        panic(err)
    }
    return session
}

func Count(db string, collection string) int {
    session := Dial()

    // close session at end of function
    defer session.Close()
    c := session.DB(db).C(collection)
    n, err := c.Count()
    if err != nil {
        panic(err)
    }
    return n
}
