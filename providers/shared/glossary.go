package shared

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

func BuildStatList(categorizedStats CategorizedStats) []PayloadStat {
    statList := []PayloadStat{}
    for statIndex := range categorizedStats.Offense {
        stat := categorizedStats.Offense[statIndex]
        payloadStat := buildPayloadStat("offense", stat.S, stat.V)
        statList = append(statList, payloadStat)
    }
    for statIndex := range categorizedStats.Defense {
        stat := categorizedStats.Defense[statIndex]
        payloadStat := buildPayloadStat("defense", stat.S, stat.V)
        statList = append(statList, payloadStat)
    }
    for statIndex := range categorizedStats.General {
        stat := categorizedStats.General[statIndex]
        payloadStat := buildPayloadStat("general", stat.S, stat.V)
        statList = append(statList, payloadStat)
    }
    return statList
}
