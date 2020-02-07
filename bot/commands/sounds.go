package commands

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var buffer = make([][]byte, 0)

func PlayBora(s *discordgo.Session, guildID, channelID string) (err error) {
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	time.Sleep(250 * time.Millisecond)

	vc.Speaking(true)

	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	vc.Speaking(false)

	time.Sleep(250 * time.Millisecond)

	vc.Disconnect()

	return nil
}

func loadBora() error {
	file, err := os.Open("./assets/souds/bora-cumpadi-ajuda-o-doente.mp3")
	if err != nil {
		fmt.Println("Error opening file :", err)
		return err
	}

	var opuslen int16

	for {
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}

		if err != nil {
			fmt.Println("Error reading file :", err)
			return err
		}

		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		if err != nil {
			fmt.Println("Error reading file :", err)
			return err
		}

		buffer = append(buffer, InBuf)
	}
}
