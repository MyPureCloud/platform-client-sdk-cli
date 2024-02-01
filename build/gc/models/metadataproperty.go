package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadatapropertyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadatapropertyDud struct { 
    VarType string `json:"type"`


    Displayname string `json:"displayname"`


    Description string `json:"description"`


    Sensitive string `json:"sensitive"`


    Help string `json:"help"`


    VarDefault string `json:"default"`


    Enum []string `json:"enum"`

}

// Metadataproperty - Data property required as input for installing an accelerator
type Metadataproperty struct { 
    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Metadataproperty) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadataproperty) MarshalJSON() ([]byte, error) {
    type Alias Metadataproperty

    if MetadatapropertyMarshalled {
        return []byte("{}"), nil
    }
    MetadatapropertyMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

