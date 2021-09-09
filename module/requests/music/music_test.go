package music

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMusicInfoByKugou(t *testing.T) {
	a := assert.New(t)

	song, err := GetMusicInfoByKugou("句号")
	if err != nil {
		t.Error(err)
	}
	a.EqualValues("句号", song.Title)
	a.EqualValues("摩天动物园", song.Album)
	a.EqualValues(235, song.Time)

	song, err = GetMusicInfoByKugou("素颜")
	if err != nil {
		t.Error(err)
	}
	a.EqualValues("素颜", song.Title)
	a.EqualValues("素颜", song.Album)
	a.EqualValues(241, song.Time)

	song, err = GetMusicInfoByKugou("一世心上长安")
	if err != nil {
		t.Error(err)
	}
	a.EqualValues("一世心上长安", song.Title)
	a.EqualValues("剑网3·乘梦江湖", song.Album)
	a.EqualValues(251, song.Time)
}
