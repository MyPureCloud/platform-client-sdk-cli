package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotasyncaggregatequeryresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotasyncaggregatequeryresponseDud struct { 
    


    

}

// Botasyncaggregatequeryresponse
type Botasyncaggregatequeryresponse struct { 
    // Results
    Results []Botaggregatedatacontainer `json:"results"`


    // Cursor - Cursor token to retrieve next page
    Cursor string `json:"cursor"`

}

// String returns a JSON representation of the model
func (o *Botasyncaggregatequeryresponse) String() string {
     o.Results = []Botaggregatedatacontainer{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botasyncaggregatequeryresponse) MarshalJSON() ([]byte, error) {
    type Alias Botasyncaggregatequeryresponse

    if BotasyncaggregatequeryresponseMarshalled {
        return []byte("{}"), nil
    }
    BotasyncaggregatequeryresponseMarshalled = true

    return json.Marshal(&struct {
        
        Results []Botaggregatedatacontainer `json:"results"`
        
        Cursor string `json:"cursor"`
        *Alias
    }{

        
        Results: []Botaggregatedatacontainer{{}},
        


        

        Alias: (*Alias)(u),
    })
}

