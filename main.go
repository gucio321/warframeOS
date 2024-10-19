package main

import (
	"embed"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

//go:embed data
var data embed.FS

func getPath(cmd []string) string {
	c := strings.Join(cmd, " ")
	// command : file
	cmds := map[string]string{
		"yay":        "yay.mp3",
		"git add":    "commit.mp3",
		"git commit": "upgrade.mp3",
		"git push":   "pushing.mp3",
		"rm":         "eliminated.mp3",
		"man":        "a_little_help.mp3",
		"info":       "a_little_help.mp3",
		"neofetch":   "weak_machine.mp3",
		"fastfetch":  "weak_machine.mp3",
	}

	for k, v := range cmds {
		if strings.HasPrefix(c, k) {
			return filepath.Join("data", v)
		}
	}

	return ""
}

func main() {
	cmdParts := os.Args[1:]
	if len(cmdParts) == 0 {
		return
	}

	path := getPath(cmdParts)
	if path == "" {
		return
	}

	file, err := data.Open(path)
	if err != nil {
		panic(err)
	}

	r, err := mp3.NewDecoder(file)
	if err != nil {
		panic(err)
	}

	op := &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 48000

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	player := otoCtx.NewPlayer(r)
	player.Play()
	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
