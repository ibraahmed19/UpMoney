package message

import (
	"gorm.io/gorm"
)

type Message struct {
	ID        uint   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Triggers  string `gorm:"column:triggers" json:"triggers"` //first msg:  from client > getMessageByTriggers
	Text      string `gorm:"column:text" json:"text"`
	Next      uint   `gorm:"column:next" json:"next"`
	CreatedBy uint   `gorm:"column:created_by" json:"created_by"`
	gorm.Model
}

func SearchTrigger(db *gorm.DB, message Message) (messages []Message, err error) {
	return messages, db.Where(&message).Find(&messages).Error
}

func GetMessageByID(db *gorm.DB, id uint) (message Message, err error) {
	return message, db.Find(&message, "id=?", id).Error
}

func CheckMessageExists(db *gorm.DB, id uint) bool {

	// init vars
	message := &Message{}

	// check if row exists
	check := db.Where("id=?", id).First(&message)
	if check.Error != nil {
		return false
	}

	if check.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

/*getMessageByTriggers : in order to get the firsr response to first question or message of the client --> return message

message --> variable next


get MessageById(next)
.+.+.+.+.+ etc

next = null stop

id : 1
triggers : ["Hello", "word"]
text: "how can i help you"
Next: 2

//
id: 2
triggers: ["transaction", "ttransfert", "payment", "virement"]
text: "sure! can you please give me the amount to be sended "
nex: 3

id: 3
triggers :
text "can you please now give the purpose of this trabsaction"
negorm

id 4
triggers
text can you please give me

id 5
triggers
text transaction has been seccesfully sended
next null

id: 6
triggers
text sorry i can't help you
next: null
*/
