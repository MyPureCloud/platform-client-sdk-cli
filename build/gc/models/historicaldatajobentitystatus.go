package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricaldatajobentitystatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricaldatajobentitystatusDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Historicaldatajobentitystatus
type Historicaldatajobentitystatus struct { 
    


    // State - Property denoting the status of the delete job
    State string `json:"state"`


    

}

// String returns a JSON representation of the model
func (o *Historicaldatajobentitystatus) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicaldatajobentitystatus) MarshalJSON() ([]byte, error) {
    type Alias Historicaldatajobentitystatus

    if HistoricaldatajobentitystatusMarshalled {
        return []byte("{}"), nil
    }
    HistoricaldatajobentitystatusMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

