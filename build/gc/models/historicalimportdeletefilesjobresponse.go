package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoricalimportdeletefilesjobresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoricalimportdeletefilesjobresponseDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Historicalimportdeletefilesjobresponse
type Historicalimportdeletefilesjobresponse struct { 
    


    // State - Property denoting the state of the remove request
    State string `json:"state"`


    // Entities - The request entities that got deleted
    Entities []Historicaldatadeleteentity `json:"entities"`


    // DisallowedEntities - The request entities that were disallowed to be deleted
    DisallowedEntities []Historicaldatadisalloweddeleteentity `json:"disallowedEntities"`


    

}

// String returns a JSON representation of the model
func (o *Historicalimportdeletefilesjobresponse) String() string {
    
     o.Entities = []Historicaldatadeleteentity{{}} 
     o.DisallowedEntities = []Historicaldatadisalloweddeleteentity{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historicalimportdeletefilesjobresponse) MarshalJSON() ([]byte, error) {
    type Alias Historicalimportdeletefilesjobresponse

    if HistoricalimportdeletefilesjobresponseMarshalled {
        return []byte("{}"), nil
    }
    HistoricalimportdeletefilesjobresponseMarshalled = true

    return json.Marshal(&struct {
        
        State string `json:"state"`
        
        Entities []Historicaldatadeleteentity `json:"entities"`
        
        DisallowedEntities []Historicaldatadisalloweddeleteentity `json:"disallowedEntities"`
        *Alias
    }{

        


        


        
        Entities: []Historicaldatadeleteentity{{}},
        


        
        DisallowedEntities: []Historicaldatadisalloweddeleteentity{{}},
        


        

        Alias: (*Alias)(u),
    })
}

