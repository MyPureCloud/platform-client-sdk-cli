package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AcceleratorspecificationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AcceleratorspecificationDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    Description string `json:"description"`


    Origin string `json:"origin"`


    VarType string `json:"type"`


    Classification string `json:"classification"`


    Tags []string `json:"tags"`


    Permissions []string `json:"permissions"`


    Products []string `json:"products"`


    Documentation []Metadatadocumentation `json:"documentation"`


    Presentation []Metadatapresentation `json:"presentation"`


    Results Metadataresults `json:"results"`


    SelfUri string `json:"selfUri"`

}

// Acceleratorspecification - Metadata for a CX infrastructure as code accelerator
type Acceleratorspecification struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    

}

// String returns a JSON representation of the model
func (o *Acceleratorspecification) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Acceleratorspecification) MarshalJSON() ([]byte, error) {
    type Alias Acceleratorspecification

    if AcceleratorspecificationMarshalled {
        return []byte("{}"), nil
    }
    AcceleratorspecificationMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

