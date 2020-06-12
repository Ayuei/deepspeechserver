package speech

import (
	"fmt"
	"github.com/ayuei/deepspeechserver/server"
	"github.com/ayuei/deepspeechserver/utils"
	"log"

	deepspeech "github.com/asticode/go-astideepspeech"
)

type Speech struct {
	*deepspeech.Model
	IsProcessing bool
}

func New(modelPath string, scorer string) *Speech{
	model, err := deepspeech.New(modelPath)
	utils.Chkerror(err)
	err = model.EnableExternalScorer(scorer)
	utils.Chkerror(err)

	return &Speech{
		model,
		false,
	}
}

func (model *Speech) Start(client *server.Client) {
	defer client.Close()

	stream, err := model.NewStream()
	utils.Chkerror(err)

	for {
		select{
		case input := <- client.ReadBuffer:
			if len(input.PCM) > 0 {
				fmt.Println("Received message in Model")
				stream.FeedAudioContent(input.PCM)
			} else {
				fmt.Println("Finishing up")
				pred, err := stream.Finish()
				if utils.Chkerror(err){
					log.Fatal(err)
				}

				stream, err = model.NewStream()
				if utils.Chkerror(err){
					log.Fatal(err)
				}

				client.WriteBuffer <- pred
			}
		}
	}
}
