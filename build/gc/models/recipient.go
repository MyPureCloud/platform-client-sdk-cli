package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecipientMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecipientDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Recipient
type Recipient struct { 
    


    // Name
    Name string `json:"name"`


    // Flow - An automate flow object which defines the set of actions to be taken, when a message is received by this provisioned phone number.
    Flow Flow `json:"flow"`


    // DateCreated - Date this recipient was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - Date this recipient was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    // CreatedBy - User that created this recipient
    CreatedBy User `json:"createdBy"`


    // ModifiedBy - User that modified this recipient
    ModifiedBy User `json:"modifiedBy"`


    // MessengerType - The messenger type for this recipient
    MessengerType string `json:"messengerType"`


    

}

// String returns a JSON representation of the model
func (o *Recipient) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recipient) MarshalJSON() ([]byte, error) {
    type Alias Recipient

    if RecipientMarshalled {
        return []byte("{}"), nil
    }
    RecipientMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Flow Flow `json:"flow"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        CreatedBy User `json:"createdBy"`
        
        ModifiedBy User `json:"modifiedBy"`
        
        MessengerType string `json:"messengerType"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

