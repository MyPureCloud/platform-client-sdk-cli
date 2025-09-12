package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DataextractionfileschemalistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DataextractionfileschemalistingDud struct { 
    


    


    


    


    


    

}

// Dataextractionfileschemalisting
type Dataextractionfileschemalisting struct { 
    // Entities
    Entities []Dataextractionfileschema `json:"entities"`


    // NextUri
    NextUri string `json:"nextUri"`


    // SelfUri
    SelfUri string `json:"selfUri"`


    // PreviousUri
    PreviousUri string `json:"previousUri"`


    // EnabledDataSchemas - The data schemas that are enabled for extraction
    EnabledDataSchemas []string `json:"enabledDataSchemas"`


    // Errors
    Errors Errorbody `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Dataextractionfileschemalisting) String() string {
     o.Entities = []Dataextractionfileschema{{}} 
    
    
    
     o.EnabledDataSchemas = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dataextractionfileschemalisting) MarshalJSON() ([]byte, error) {
    type Alias Dataextractionfileschemalisting

    if DataextractionfileschemalistingMarshalled {
        return []byte("{}"), nil
    }
    DataextractionfileschemalistingMarshalled = true

    return json.Marshal(&struct {
        
        Entities []Dataextractionfileschema `json:"entities"`
        
        NextUri string `json:"nextUri"`
        
        SelfUri string `json:"selfUri"`
        
        PreviousUri string `json:"previousUri"`
        
        EnabledDataSchemas []string `json:"enabledDataSchemas"`
        
        Errors Errorbody `json:"errors"`
        *Alias
    }{

        
        Entities: []Dataextractionfileschema{{}},
        


        


        


        


        
        EnabledDataSchemas: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

