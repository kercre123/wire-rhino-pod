package wirepod

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/digital-dream-labs/chipper/pkg/vtt"

	leopard "github.com/Picovoice/leopard/binding/go"
	rhino "github.com/Picovoice/rhino/binding/go/v2"
	opus "github.com/digital-dream-labs/opus-go/opus"
)

var leopardSTT leopard.Leopard

var debugLogging bool

var botNum int = 0
var disableLiveTranscription bool = false

func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
	TODO:
	1. Implement jdocs. These are files which are stored on the bot which contain the bot's
	default location, unit settings, etc. Helpful for weather.
		- current workaround: setup specific bots with botSetup.sh
*/

func (s *Server) ProcessIntent(req *vtt.IntentRequest) (*vtt.IntentResponse, error) {
	var voiceTimer int = 0
	var transcription1 string = ""
	var transcription2 string = ""
	var transcribedText string
	var isOpus bool
	var micDataLeopard []int16
	var micDataRhino [][]int16
	var die bool = false
	var doSTT bool = true
	var sayStarting = true
	var leopardSTT leopard.Leopard
	var rhinoSTI rhino.Rhino
	var numInRange int = 0
	var oldDataLength int = 0
	var transcribedIntent string = ""
	var transcribedSlots map[string]string
	var rhinoDone bool = false
	var leopardFallback bool = false
	var successMatched bool = false
	if os.Getenv("DEBUG_LOGGING") != "true" && os.Getenv("DEBUG_LOGGING") != "false" {
		fmt.Println("No valid value for DEBUG_LOGGING, setting to true")
		debugLogging = true
	} else {
		if os.Getenv("DEBUG_LOGGING") == "true" {
			debugLogging = true
		} else {
			debugLogging = false
		}
	}
	botNum = botNum + 1
	justThisBotNum := botNum
	if botNum > 1 {
		if debugLogging == true {
			fmt.Println("Multiple bots are streaming, live transcription disabled")
		}
		disableLiveTranscription = true
	} else {
		disableLiveTranscription = false
	}
	if os.Getenv("DISABLE_LIVE_TRANSCRIPTION") == "true" {
		if debugLogging == true {
			fmt.Println("DISABLE_LIVE_TRANSCRIPTION is true, live transcription disabled")
		}
		disableLiveTranscription = true
	}
	if justThisBotNum == 1 {
		leopardSTT = leopardSTT1
		rhinoSTI = rhinoSTI1
	} else if justThisBotNum == 2 {
		leopardSTT = leopardSTT2
		rhinoSTI = rhinoSTI2
	} else if justThisBotNum == 3 {
		leopardSTT = leopardSTT3
		rhinoSTI = rhinoSTI3
	} else if justThisBotNum == 4 {
		leopardSTT = leopardSTT4
		rhinoSTI = rhinoSTI4
	} else if justThisBotNum == 5 {
		leopardSTT = leopardSTT5
		rhinoSTI = rhinoSTI5
	} else {
		fmt.Println("Too many bots are connected, sending error to bot " + strconv.Itoa(justThisBotNum))
		IntentPass(req, "intent_system_noaudio", "Too many bots, max is 5", map[string]string{"error": "EOF"}, true, justThisBotNum)
		botNum = botNum - 1
		return nil, nil
	}
	if debugLogging == true {
		fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " ESN: " + req.Device)
		fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Session: " + req.Session)
		fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Language: " + req.LangString)
		fmt.Println("Stream " + strconv.Itoa(justThisBotNum) + " opened.")
	}
	data := []byte{}
	data = append(data, req.FirstReq.InputAudio...)
	if len(data) > 0 {
		if data[0] == 0x4f {
			isOpus = true
			if debugLogging == true {
				fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Stream Type: Opus")
			}
		} else {
			isOpus = false
			if debugLogging == true {
				fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Stream Type: PCM")
			}
		}
	}
	stream := opus.OggStream{}
	go func() {
		if isOpus == true {
			time.Sleep(time.Millisecond * 500)
		} else {
			time.Sleep(time.Millisecond * 1100)
		}
		for voiceTimer < 7 {
			voiceTimer = voiceTimer + 1
			time.Sleep(time.Millisecond * 750)
		}
	}()
	go func() {
		if botNum > 1 {
			disableLiveTranscription = true
		}
		for doSTT == true {
			if micDataLeopard != nil && die == false && voiceTimer > 0 {
				if leopardFallback == true {
					if sayStarting == true {
						if debugLogging == true {
							fmt.Println("No intent was transcribed, falling back to leopard")
							fmt.Printf("Transcribing stream %d...\n", justThisBotNum)
						}
						sayStarting = false
					}
					processOneData := micDataLeopard
					transcription1Raw, err := leopardSTT.Process(processOneData)
					if err != nil {
						log.Println(err)
					}
					transcription1 = strings.ToLower(transcription1Raw)
					if debugLogging == true {
						if disableLiveTranscription == false {
							fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcription1)
						}
					}
					if transcription1 != "" && transcription2 != "" && transcription1 == transcription2 {
						transcribedText = transcription1
						if debugLogging == true {
							if disableLiveTranscription == true {
								fmt.Printf("\n")
							} else {
								fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcribedText)
							}
						}
						die = true
						break
					} else if voiceTimer == 7 {
						transcribedText = transcription2
						if debugLogging == true {
							if disableLiveTranscription == true {
								fmt.Printf("\n")
							} else {
								fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcribedText)
							}
						}
						die = true
						break
					}
					time.Sleep(time.Millisecond * 300)
					processTwoData := micDataLeopard
					transcription2Raw, err := leopardSTT.Process(processTwoData)
					if err != nil {
						log.Println(err)
					}
					transcription2 = strings.ToLower(transcription2Raw)
					if debugLogging == true {
						if disableLiveTranscription == false {
							fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcription2)
						}
					}
					if transcription1 != "" && transcription2 != "" && transcription1 == transcription2 {
						transcribedText = transcription1
						if debugLogging == true {
							if disableLiveTranscription == true {
								fmt.Printf("\n")
							} else {
								fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcribedText)
							}
						}
						die = true
						break
					} else if voiceTimer == 7 {
						transcribedText = transcription2
						if debugLogging == true {
							if disableLiveTranscription == true {
								fmt.Printf("\n")
							} else {
								fmt.Println("Bot " + strconv.Itoa(justThisBotNum) + " Transcription: " + transcribedText)
							}
						}
						die = true
						break
					}
					time.Sleep(time.Millisecond * 300)
				}
			}
		}
	}()
	for {
		chunk, err := req.Stream.Recv()
		if err != nil {
			if err == io.EOF {
				IntentPass(req, "intent_system_noaudio", "EOF error", map[string]string{"error": "EOF"}, true, justThisBotNum)
				break
			}
		}
		if die == true {
			break
		}
		data = append(data, chunk.InputAudio...)
		// returns []int16, framesize unknown
		micDataLeopard = bytesToIntLeopard(stream, data, die, isOpus)
		// returns [][]int16, 512 framesize
		micDataRhino = bytesToIntRhino(stream, data, die, isOpus)
		numInRange = 0
		for _, sample := range micDataRhino {
			if rhinoDone == false {
				if numInRange >= oldDataLength {
					isFinalized, err := rhinoSTI.Process(sample)
					if isFinalized {
						inference, err := rhinoSTI.GetInference()
						if err != nil {
							fmt.Println("Error getting inference: " + err.Error())
						}
						if inference.IsUnderstood {
							transcribedIntent = inference.Intent
							transcribedSlots = inference.Slots
							die = true
							leopardFallback = false
							rhinoDone = true
							break
						} else {
							fmt.Println("No intent was matched.")
							leopardFallback = true
							rhinoDone = true
							break
						}
					}
					if err != nil {
						fmt.Println("Error: " + err.Error())
						break
					}
					numInRange = numInRange + 1
				} else {
					numInRange = numInRange + 1
				}
			}
		}
		oldDataLength = len(micDataRhino)
		if die == true {
			break
		}
	}
	if leopardFallback == true {
		successMatched = processTextAll(req, transcribedText, matchListList, intentsList, isOpus, justThisBotNum)
	} else {
		successMatched = true
		paramCheckerSlots(req, transcribedIntent, transcribedSlots, isOpus, justThisBotNum)
	}
	if successMatched == false {
		if debugLogging == true {
			fmt.Println("No intent was matched.")
		}
		IntentPass(req, "intent_system_noaudio", transcribedText, map[string]string{"": ""}, false, justThisBotNum)
	}
	botNum = botNum - 1
	return nil, nil
}
