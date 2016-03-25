package shared

import "math"
import "strconv"
import "encoding/json"
import "io/ioutil"

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
    JoinWith   string      `json:"join_with"`
}

type PayloadStat struct {
    Identifier  StatIdentifier  `json:"identifier"`
    Value       JsonFloat64     `json:"value"`
    Format      StatFormat      `json:"format"`
}

type PayloadGlossaryEntry struct {
    Identifier  StatIdentifier  `json:"identifier"`
    Link        string          `json:"link"`
    Name        string          `json:"name"`
    Key         string          `json:"key"`
    Description string          `json:"description"`
    Format      string          `json:"format"`
}

func (entry PayloadGlossaryEntry) GetStatFormat() StatFormat {
    statFormat := StatFormat{}
    statFormat.RemoveZero = entry.Format == "triple-dec"

    statFormat.FormatType = "decimal"
    statFormat.Precision = 0
    if entry.Format == "int" || entry.Format == "dec" || entry.Format == "triple-dec" {
        if entry.Format == "dec" {
            statFormat.Precision = 1
        } else if entry.Format == "triple-dec" {
            statFormat.Precision = 3
        }
    } else if entry.Format == "percentage" {
        statFormat.FormatType = "percentage"
        statFormat.Precision = 2
    } else if entry.Format == "join" {
        statFormat.FormatType = "join"
    }

    return statFormat
}


type GlossaryEntry struct {
    Link        string          `json:"link"`
    Name        string          `json:"abbrev"`
    Key         string          `json:"key"`
    Description string          `json:"description"`
    Format      string          `json:"format"`
}

type GlossarySport struct {
    Offense []GlossaryEntry     `json:"offense"`
    Defense []GlossaryEntry     `json:"defense"`
    General []GlossaryEntry     `json:"general"`
}

type Glossary struct {
    Baseball    GlossarySport   `json:"baseball"`
    Basketball  GlossarySport   `json:"basketball"`
}


func (g GlossarySport) BuildPayloadEntry(identifier StatIdentifier) PayloadGlossaryEntry {
    payload := PayloadGlossaryEntry{Identifier: identifier}
    category := []GlossaryEntry{}

    switch {
    case identifier.Category == "offense":
        category = g.Offense
        break
    case identifier.Category == "defense":
        category = g.Defense
        break
    case identifier.Category == "general":
        category = g.General
        break
    }
    chosenEntry := GlossaryEntry{}
    for _, entry := range category {
        if entry.Key == identifier.Key{
            chosenEntry = entry
            break
        }
    }
    payload.Key = chosenEntry.Key
    payload.Link = chosenEntry.Link
    payload.Name = chosenEntry.Name
    payload.Description = chosenEntry.Description
    payload.Format = chosenEntry.Format
    return payload

}

func GetGlossary(statsRequired []StatIdentifier) []PayloadGlossaryEntry {
    bytes, fileErr := ioutil.ReadFile("/go/src/github.com/nrschultz/go-server/data.json")
    if fileErr != nil {
        panic(fileErr)
    }
    glossary := Glossary{}
    jsonErr := json.Unmarshal(bytes, &glossary)
    if jsonErr != nil {
        panic(jsonErr)
    }

    sportGlossary := glossary.Baseball
    payloadGlossary := []PayloadGlossaryEntry{}
    for _, identifier := range statsRequired {
        payloadGlossary = append(payloadGlossary, sportGlossary.BuildPayloadEntry(identifier))
    }
    return payloadGlossary
}

func formatForIdentifier(identifier StatIdentifier, glossary []PayloadGlossaryEntry) StatFormat {
    chosenEntry := PayloadGlossaryEntry{}
    for _, entry := range glossary {
        if entry.Identifier.Category == identifier.Category && entry.Identifier.Key == identifier.Key {
            chosenEntry = entry
            break
        }
    }
    return chosenEntry.GetStatFormat()
}

func buildPayloadStat(category string, name string, value float64, glossary []PayloadGlossaryEntry) PayloadStat {
    payloadStat := PayloadStat{}
    payloadStat.Value = JsonFloat64(value)
    payloadStat.Identifier = StatIdentifier{Category: category, Key: name}
    payloadStat.Format = formatForIdentifier(payloadStat.Identifier, glossary)
    return payloadStat
}

func BuildStatList(categorizedStats CategorizedStats, statsRequested []StatIdentifier, glossary []PayloadGlossaryEntry) []PayloadStat {
    statList := []PayloadStat{}
    for _, identifier := range statsRequested {
        value := categorizedStats.GetStatValue(identifier)
        payloadStat := buildPayloadStat(identifier.Category, identifier.Key, value, glossary)
        statList = append(statList, payloadStat)
    }
    return statList
}
