package main

import (
    "encoding/csv"
    "fmt"
    "net/http"
    "os"
    "strings"
    "makves/struct"
)

func getRecordByID(filePath, id string) (*structure.Record, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)

    _, err = reader.Read()
    if err != nil {
        return nil, err
    }

    for {
        record, err := reader.Read()
        if err != nil {
            break
        }

        if record[0] == id {
            return &structure.Record{
                ID:                        record[0],
                UID:                       record[1],
                Domain:                    record[2],
                CN:                        record[3],
                Department:                record[4],
                Title:                     record[5],
                Who:                       record[6],
                LogonCount:                record[7],
                NumLogons7:                record[8],
                NumShare7:                 record[9],
                NumFile7:                  record[10],
                NumAd7:                    record[11],
                NumN7:                     record[12],
                NumLogons14:               record[13],
                NumShare14:                record[14],
                NumFile14:                 record[15],
                NumAd14:                   record[16],
                NumN14:                    record[17],
                NumLogons30:               record[18],
                NumShare30:                record[19],
                NumFile30:                 record[20],
                NumAd30:                   record[21],
                NumN30:                    record[22],
                NumLogons150:              record[23],
                NumShare150:               record[24],
                NumFile150:                record[25],
                NumAd150:                  record[26],
                NumN150:                   record[27],
                NumLogons365:              record[28],
                NumShare365:               record[29],
                NumFile365:                record[30],
                NumAd365:                  record[31],
                NumN365:                   record[32],
                HasUserPrincipalName:      record[33],
                HasMail:                   record[34],
                HasPhone:                  record[35],
                FlagDisabled:              record[36],
                FlagLockout:               record[37],
                FlagPasswordNotRequired:   record[38],
                FlagPasswordCantChange:    record[39],
                FlagDontExpirePassword:    record[40],
                OwnedFiles:                record[41],
                NumMailboxes:              record[42],
                NumMemberOfGroups:         record[43],
                NumMemberOfIndirectGroups: record[44],
                MemberOfIndirectGroupsIDs: record[45],
                MemberOfGroupsIDs:         record[46],
                IsAdmin:                   record[47],
                IsService:                 record[48],
            }, nil
        }
    }

    return nil, fmt.Errorf("record with ID %s not found", id)
}

func handler(w http.ResponseWriter, r *http.Request) {
    idsParam := r.URL.Query().Get("ids")
    if idsParam == "" {
        http.Error(w, "Missing 'ids' parameter", http.StatusBadRequest)
        return
    }

    ids := strings.Split(idsParam, ",")
    filePath := "ueba.csv"

    var records []structure.Record
    var notFound []string

    for _, id := range ids {
        id = strings.TrimSpace(id)
        record, err := getRecordByID(filePath, id)
        if err != nil {
            notFound = append(notFound, id)
        } else {
            records = append(records, *record)
        }
    }

    if len(records) > 0 {
        for _, record := range records {
            fmt.Fprintf(w, "Record for ID %s: %+v\n", record.ID, record)
        }
    }

    if len(notFound) > 0 {
        fmt.Fprintf(w, "\nRecords not found for IDs: %s\n", strings.Join(notFound, ", "))
    }
}

func main() {
    http.HandleFunc("/get-items", handler)
    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}
