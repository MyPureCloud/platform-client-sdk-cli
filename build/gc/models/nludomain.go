package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainDud struct { 
    Id string `json:"id"`


    


    


    DraftVersion Nludomainversion `json:"draftVersion"`


    LastPublishedVersion Nludomainversion `json:"lastPublishedVersion"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    


    SelfUri string `json:"selfUri"`

}

// Nludomain
type Nludomain struct { 
    


    // Name - The name of the NLU domain.
    Name string `json:"name"`


    // Language - The language culture of the NLU domain, e.g. `en-us`, `de-de`.
    Language string `json:"language"`


    


    


    


    


    // EngineVersion - The version of the NLU engine to use.
    EngineVersion string `json:"engineVersion"`


    

}

// String returns a JSON representation of the model
func (o *Nludomain) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomain) MarshalJSON() ([]byte, error) {
    type Alias Nludomain

    if NludomainMarshalled {
        return []byte("{}"), nil
    }
    NludomainMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Language string `json:"language"`
        
        EngineVersion string `json:"engineVersion"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

