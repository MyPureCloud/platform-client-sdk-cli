package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomconversationattributeinputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomconversationattributeinputDud struct { 
    


    


    

}

// Customconversationattributeinput - A Conversation Custom Attributes schema and its associated records made available to a guide session turn.
type Customconversationattributeinput struct { 
    // SchemaId - The ID of the Conversation Custom Attributes schema.
    SchemaId string `json:"schemaId"`


    // DivisionIds - The division IDs associated with this schema.
    DivisionIds []string `json:"divisionIds"`


    // RecordIds - The record IDs associated with this schema.
    RecordIds []string `json:"recordIds"`

}

// String returns a JSON representation of the model
func (o *Customconversationattributeinput) String() string {
    
     o.DivisionIds = []string{""} 
     o.RecordIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customconversationattributeinput) MarshalJSON() ([]byte, error) {
    type Alias Customconversationattributeinput

    if CustomconversationattributeinputMarshalled {
        return []byte("{}"), nil
    }
    CustomconversationattributeinputMarshalled = true

    return json.Marshal(&struct {
        
        SchemaId string `json:"schemaId"`
        
        DivisionIds []string `json:"divisionIds"`
        
        RecordIds []string `json:"recordIds"`
        *Alias
    }{

        


        
        DivisionIds: []string{""},
        


        
        RecordIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

