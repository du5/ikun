package mp3

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"sync"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	bmp3 "github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var (
	//go:embed 唱.mp3
	Chang []byte
	//go:embed 跳.mp3
	Tiao []byte
	//go:embed rap.mp3
	Rap []byte
	//go:embed 篮球.mp3
	Lanqiu []byte
	//go:embed 鸡.mp3
	Ji []byte
	//go:embed 你.mp3
	Ni []byte
	//go:embed 太.mp3
	Tai []byte
	//go:embed 美.mp3
	Mei []byte
	//go:embed 美2.mp3
	Mei2 []byte

	mixer *beep.Mixer
	ikunM ikunMusic
)

type ikunMusic struct {
	Mp3   [][]byte       `json:"mp3"`
	Index map[string]int `json:"index"`
	sync.Mutex
}

func (i *ikunMusic) getIOR(k string) io.ReadCloser {
	i.Lock()
	defer i.Unlock()
	return io.NopCloser(bytes.NewReader(i.Mp3[i.Index[k]]))
}

func init() {
	mixer = &beep.Mixer{}
	_ = speaker.Init(44100, 4410)
	speaker.Play(&effects.Volume{
		Streamer: mixer,
		Base:     2,
		Volume:   4,
		Silent:   false,
	})
	ikunM = ikunMusic{
		Mp3: [][]byte{Chang, Tiao, Rap, Lanqiu, Ji, Ni, Tai, Mei, Mei2},
		Index: map[string]int{
			"c":       0,
			"tiao":    1,
			"r":       2,
			"l":       3,
			"j":       4,
			"n":       5,
			"tai":     6,
			"m_short": 7,
			"m_long":  8,
		},
	}

}

func Play(k string) {
	streamer, _, err := bmp3.Decode(ikunM.getIOR(k))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	mixer.Add(streamer)
}
