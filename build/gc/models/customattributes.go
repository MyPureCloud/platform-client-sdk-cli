package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CustomattributesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CustomattributesDud struct { 
    


    


    


    


    


    


    CustomAttributesTimestamps map[string]string `json:"customAttributesTimestamps"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Customattributes
type Customattributes struct { 
    // Id - The id of the Custom Attributes record.
    Id string `json:"id"`


    // Name
    Name string `json:"name"`


    // ConversationId - The id of the conversation.
    ConversationId string `json:"conversationId"`


    // Divisions - The list of division ids that the record is visible in. If [], the record is visible to all divisions (Unassigned Division).
    Divisions []string `json:"divisions"`


    // Schema - The schema that dictates which attributes can be included.
    Schema Conversationdataschema `json:"schema"`


    // CustomAttributes - The map of attribute values.
    CustomAttributes map[string]interface{} `json:"customAttributes"`


    


    // Version - The latest version of the Custom Attributes record.
    Version int `json:"version"`


    // DateCreated - The date the record was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The date the record was last updated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    

}

// String returns a JSON representation of the model
func (o *Customattributes) String() string {
    
    
    
     o.Divisions = []string{""} 
    
     o.CustomAttributes = map[string]interface{}{"": Interface{}} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Customattributes) MarshalJSON() ([]byte, error) {
    type Alias Customattributes

    if CustomattributesMarshalled {
        return []byte("{}"), nil
    }
    CustomattributesMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Name string `json:"name"`
        
        ConversationId string `json:"conversationId"`
        
        Divisions []string `json:"divisions"`
        
        Schema Conversationdataschema `json:"schema"`
        
        CustomAttributes map[string]interface{} `json:"customAttributes"`
        
        Version int `json:"version"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        
        Divisions: []string{""},
        


        


        
        CustomAttributes: map[string]interface{}{"": Interface{}},
        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

