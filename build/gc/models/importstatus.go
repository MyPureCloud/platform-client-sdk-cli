package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ImportstatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ImportstatusDud struct { 
    State string `json:"state"`


    TotalRecords int `json:"totalRecords"`


    CompletedRecords int `json:"completedRecords"`


    PercentComplete int `json:"percentComplete"`


    FailureReason string `json:"failureReason"`


    TargetContactListIds []string `json:"targetContactListIds"`


    ListNamePrefix string `json:"listNamePrefix"`

}

// Importstatus
type Importstatus struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Importstatus) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Importstatus) MarshalJSON() ([]byte, error) {
    type Alias Importstatus

    if ImportstatusMarshalled {
        return []byte("{}"), nil
    }
    ImportstatusMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

