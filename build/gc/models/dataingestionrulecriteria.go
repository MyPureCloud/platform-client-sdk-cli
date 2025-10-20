package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataingestionrulecriteriaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataingestionrulecriteriaDud struct { 
    


    


    Name string `json:"name"`

}

// Dataingestionrulecriteria
type Dataingestionrulecriteria struct { 
    // EffectivePlatform - The effective platform for the data ingestion rule.
    EffectivePlatform string `json:"effectivePlatform"`


    // Id - The ID of the data ingestion rule.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Dataingestionrulecriteria) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataingestionrulecriteria) MarshalJSON() ([]byte, error) {
    type Alias Dataingestionrulecriteria

    if DataingestionrulecriteriaMarshalled {
        return []byte("{}"), nil
    }
    DataingestionrulecriteriaMarshalled = true

    return json.Marshal(&struct {
        
        EffectivePlatform string `json:"effectivePlatform"`
        
        Id string `json:"id"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

