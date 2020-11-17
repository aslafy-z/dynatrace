package customservices

import "encoding/json"

// QueueEntryPointType has no documentation
type QueueEntryPointType string

// QueueEntryPointTypes offers the known enum values
var QueueEntryPointTypes = struct {
	IBMMQ    QueueEntryPointType
	Jms      QueueEntryPointType
	Kafka    QueueEntryPointType
	MSMQ     QueueEntryPointType
	RabbitMQ QueueEntryPointType
	Values   func() []string
}{
	"IBM_MQ",
	"JMS",
	"KAFKA",
	"MSMQ",
	"RABBIT_MQ",
	func() []string {
		return []string{"IBM_MQ", "JMS", "KAFKA", "MSMQ", "RABBIT_MQ"}
	},
}

// UnmarshalJSON performs custom unmarshalling of this enum type
func (t *QueueEntryPointType) UnmarshalJSON(data []byte) error {
	var name string
	if err := json.Unmarshal(data, &name); err != nil {
		return err
	}
	*t = QueueEntryPointType(name)
	return nil
}
