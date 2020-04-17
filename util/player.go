package util

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

func LoadSound(filename string) ([][]byte, error) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file :", err)
		return nil, err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file.
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		// If this is the end of the file, just return.
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				return nil, err
			}
			return nil, nil
		}

		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		// Read encoded pcm from dca file.
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)

		// Should not be any end of file errors
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return nil, err
		}

		var buffer = make([][]byte, 0)

		// Append encoded pcm data to the buffer.
		buffer = append(buffer, InBuf)

		return buffer, nil
	}
}

// PlaySound plays the current buffer to the provided channel.
func PlaySound(s *discordgo.Session, guildID, channelID string, buffer [][]byte) (err error) {

	// Join the provided voice channel.
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}

	// Sleep for a specified amount of time before playing the sound
	time.Sleep(250 * time.Millisecond)

	// Start speaking.
	vc.Speaking(true)

	// Send the buffer data.
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// Stop speaking
	vc.Speaking(false)

	// Sleep for a specificed amount of time before ending.
	time.Sleep(250 * time.Millisecond)

	// Disconnect from the provided voice channel.
	vc.Disconnect()

	return nil
}

func PlayFile(s *discordgo.Session, m *discordgo.MessageCreate, filename string) {
	// Find the channel that the message came from.
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		// Could not find channel.
		return
	}

	// Find the guild for that channel.
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		// Could not find guild.
		return
	}

	// Look for the message sender in that guild's current voice states.
	for _, vs := range g.VoiceStates {
		if vs.UserID == m.Author.ID {
			buffer, err := LoadSound(filename)
			if err != nil {
				fmt.Println("Error loading sound:", err)
			}

			err = PlaySound(s, g.ID, vs.ChannelID, buffer)
			if err != nil {
				fmt.Println("Error playing sound:", err)
			}

			return
		}
	}
}
