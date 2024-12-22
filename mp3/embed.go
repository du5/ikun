package mp3

import (
	"bytes"
	_ "embed"
	"io"
	"log"

	"github.com/faiface/beep"
	bmp3 "github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

//go:embed 唱.mp3
var Chang []byte

//go:embed 跳.mp3
var Tiao []byte

//go:embed rap.mp3
var Rap []byte

//go:embed 篮球.mp3
var Lanqiu []byte

//go:embed 鸡.mp3
var Ji []byte

//go:embed 你.mp3
var Ni []byte

//go:embed 太.mp3
var Tai []byte

//go:embed 美.mp3
var Mei []byte

//go:embed 美2.mp3
var Mei2 []byte

var mixer = beep.Mixer{}

func init() {
	_ = speaker.Init(44100, 4410)
	speaker.Play(&mixer)
}

func Play(b []byte) {
	streamer, _, err := bmp3.Decode(io.NopCloser(bytes.NewReader(b)))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()
	mixer.Add(streamer)
}
