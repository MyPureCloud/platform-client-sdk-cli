package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NoteMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NoteDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    ExternalDataSources []Externaldatasource `json:"externalDataSources"`


    SelfUri string `json:"selfUri"`

}

// Note
type Note struct { 
    


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    // EntityId - The id of the contact or organization to which this note refers. This only needs to be set for input when using the Bulk APIs.
    EntityId string `json:"entityId"`


    // EntityType - This is only need to be set when using Bulk API. Using any other value than contact or organization will result in null being used.
    EntityType string `json:"entityType"`


    // NoteText - Between 1 and 32,000 characters.
    NoteText string `json:"noteText"`


    // ModifyDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    ModifyDate time.Time `json:"modifyDate"`


    // CreateDate - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    CreateDate time.Time `json:"createDate"`


    // CreatedBy - When creating or updating a note, only User.id is required. User object is fully populated when expanding a note.
    CreatedBy User `json:"createdBy"`


    


    

}

// String returns a JSON representation of the model
func (o *Note) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Note) MarshalJSON() ([]byte, error) {
    type Alias Note

    if NoteMarshalled {
        return []byte("{}"), nil
    }
    NoteMarshalled = true

    return json.Marshal(&struct {
        
        Division Writablestarrabledivision `json:"division"`
        
        EntityId string `json:"entityId"`
        
        EntityType string `json:"entityType"`
        
        NoteText string `json:"noteText"`
        
        ModifyDate time.Time `json:"modifyDate"`
        
        CreateDate time.Time `json:"createDate"`
        
        CreatedBy User `json:"createdBy"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

