package acc

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

//go:embed 唱.wav
var Chang []byte

//go:embed 跳.wav
var Tiao []byte

//go:embed rap.wav
var Rap []byte

//go:embed 篮球.wav
var Lanqiu []byte

//go:embed 鸡.wav
var Ji []byte

//go:embed 你.wav
var Ni []byte

//go:embed 太.wav
var Tai []byte

//go:embed 美.wav
var Mei []byte

//go:embed 美2.wav
var Mei2 []byte

func PlayAcc(b []byte) {
	streamer, format, err := wav.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	_ = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	speaker.Play(beep.Seq(streamer))
}
