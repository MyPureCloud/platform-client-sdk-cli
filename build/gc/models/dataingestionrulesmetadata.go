package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataingestionrulesmetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataingestionrulesmetadataDud struct { 
    


    


    

}

// Dataingestionrulesmetadata
type Dataingestionrulesmetadata struct { 
    // CountByStatus - Count of data ingestion rules by status
    CountByStatus map[string]int `json:"countByStatus"`


    // Platform - The platform for which the data ingestion rules are available
    Platform string `json:"platform"`


    // TotalCount - The total count of data ingestion rule
    TotalCount int `json:"totalCount"`

}

// String returns a JSON representation of the model
func (o *Dataingestionrulesmetadata) String() string {
     o.CountByStatus = map[string]int{"": 0} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataingestionrulesmetadata) MarshalJSON() ([]byte, error) {
    type Alias Dataingestionrulesmetadata

    if DataingestionrulesmetadataMarshalled {
        return []byte("{}"), nil
    }
    DataingestionrulesmetadataMarshalled = true

    return json.Marshal(&struct {
        
        CountByStatus map[string]int `json:"countByStatus"`
        
        Platform string `json:"platform"`
        
        TotalCount int `json:"totalCount"`
        *Alias
    }{

        
        CountByStatus: map[string]int{"": 0},
        


        


        

        Alias: (*Alias)(u),
    })
}

