package shared

import "math"
import "strconv"

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

type GlossaryEntry struct {
    Identifier  StatIdentifier  `json:"identifier"`
    Link        string          `json:"link"`
    Name        string          `json:"name"`
    Key         string          `json:"key"`
    Description string          `json:"description"`
}


func formatForIdentifier(identifier StatIdentifier) StatFormat {
    return StatFormat{RemoveZero: false, FormatType: "decimal", Precision: 0}
}

func buildPayloadStat(category string, name string, value float64) PayloadStat {
    payloadStat := PayloadStat{}
    payloadStat.Value = JsonFloat64(value)
    payloadStat.Identifier = StatIdentifier{Category: category, Key: name}
    payloadStat.Format = formatForIdentifier(payloadStat.Identifier)
    return payloadStat
}

func BuildStatList(categorizedStats CategorizedStats, statsRequested []StatIdentifier) []PayloadStat {
    statList := []PayloadStat{}
    for requestedStatIndex := range statsRequested {
        identifier := statsRequested[requestedStatIndex]
        value := categorizedStats.GetStatValue(identifier)
        payloadStat := buildPayloadStat(identifier.Category, identifier.Key, value)
        statList = append(statList, payloadStat)
    }
    return statList
}
