package pubsub

import (
	shell "github.com/ipfs/go-ipfs-api"
	pubsub "github.com/zacharyestep/go-orbitdb/pubsub"
  "github.com/libp2p/go-libp2p-peer"
)

var sh *shell.Shell = shell.NewShell("http://localhost:5001")

type Subscription struct {
	sub *shell.PubSubSubscription
}

/*type Record interface {
				   From() peer.ID
					      Data() []byte
				}
				
			// Message is a pubsub message.
			type Message struct {
								From     peer.ID
									Data     []byte
										Seqno    []byte
											TopicIDs []string
							}
				*/

type RecordFromMessage struct   {
			msg * shell.Message
}

func (rm RecordFromMessage) From() peer.ID {
				return rm.msg.From
}

func (rm RecordFromMessage) Data() []byte {
				return rm.msg.Data
}

func (s *Subscription) Next() (pubsub.Record, error) {
				m, e := s.sub.Next()
				return RecordFromMessage{m},e
}

func (s *Subscription) Cancel() error {
	return s.sub.Cancel()
}

type PubSub struct {
	sh *shell.Shell
}

func (ps PubSub) Subscribe(topic string) (pubsub.Subscription, error) {
	sub, err := ps.sh.PubSubSubscribe(topic)

	return &Subscription{
		sub,
	}, err
}

func (ps PubSub) Publish(topic, data string) error {
	return ps.sh.PubSubPublish(topic, data)
}

func New() PubSub {
	return PubSub{sh}
}
