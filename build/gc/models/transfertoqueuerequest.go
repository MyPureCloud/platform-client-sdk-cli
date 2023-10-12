package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransfertoqueuerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransfertoqueuerequestDud struct { 
    


    


    

}

// Transfertoqueuerequest
type Transfertoqueuerequest struct { 
    // TransferType
    TransferType string `json:"transferType"`


    // QueueId - The id of the queue.
    QueueId string `json:"queueId"`


    // QueueName - The name of the queue.
    QueueName string `json:"queueName"`

}

// String returns a JSON representation of the model
func (o *Transfertoqueuerequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transfertoqueuerequest) MarshalJSON() ([]byte, error) {
    type Alias Transfertoqueuerequest

    if TransfertoqueuerequestMarshalled {
        return []byte("{}"), nil
    }
    TransfertoqueuerequestMarshalled = true

    return json.Marshal(&struct {
        
        TransferType string `json:"transferType"`
        
        QueueId string `json:"queueId"`
        
        QueueName string `json:"queueName"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

