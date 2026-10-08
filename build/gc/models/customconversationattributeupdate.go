package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomconversationattributeupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomconversationattributeupdateDud struct { 
    


    


    

}

// Customconversationattributeupdate - A single Conversation Custom Attribute value update made during a guide session turn.
type Customconversationattributeupdate struct { 
    // Name - The name of the conversation attribute that was updated.
    Name string `json:"name"`


    // Value - The value of the conversation attribute after the update. The JSON type depends on the attribute's definition in its schema, for example string, number, or boolean.
    Value interface{} `json:"value"`


    // DateModified - The date and time the attribute was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Customconversationattributeupdate) String() string {
    
     o.Value = Interface{} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customconversationattributeupdate) MarshalJSON() ([]byte, error) {
    type Alias Customconversationattributeupdate

    if CustomconversationattributeupdateMarshalled {
        return []byte("{}"), nil
    }
    CustomconversationattributeupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Value interface{} `json:"value"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        
        Value: Interface{},
        


        

        Alias: (*Alias)(u),
    })
}

