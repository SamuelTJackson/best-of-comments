package file

import (
	"fmt"
	"github.com/SamuelTJackson/best-of-comments/youtube"
	"testing"
	"time"
)

func TestReadFileToStruct(t *testing.T) {
	var videos []youtube.Video
	if err := readFileToStruct("./test-input.json", &videos); err != nil {
		t.Error(err)
	}
	fmt.Println(videos)
	if len(videos) != 2 {
		t.Errorf("length should be 2 but was %d", len(videos))
	}
	if videos[0].ID != "1" {
		t.Errorf("id should be 1 but was %s", videos[0].ID)
	}
	if videos[1].ID != "2" {
		t.Errorf("id should be 2 but was %s", videos[1].ID)
	}
}

func TestStorage_SaveVideos(t *testing.T) {
	outputFile := "./output.json"
	inputFile := "./input.json"
	storage := NewStorage(outputFile, inputFile, OVERWRITE)
	videos := []youtube.Video{{
		ID:        "1",
		Published: time.Now(),
	}, {
		ID:        "2",
		Published: time.Now(),
	}}
	if err := storage.SaveVideos(videos...); err != nil {
		t.Errorf(err.Error())
	}
	var readInVideos []youtube.Video
	if err := readFileToStruct(outputFile, &readInVideos); err != nil {
		t.Error(err)
	}
	if len(readInVideos) != 2 {
		t.Errorf("should be only 2 videos but was %d", len(readInVideos))
	}
	if !readInVideos[0].Published.Equal(videos[0].Published) {
		t.Errorf("%v time should be equal to %v", readInVideos[0].Published, videos[0].Published)
	}
	if !readInVideos[1].Published.Equal(videos[1].Published) {
		t.Errorf("%v time should be equal to %v", readInVideos[1].Published, videos[1].Published)
	}
	storage = NewStorage(outputFile, inputFile, APPEND)
	if err := storage.SaveVideos(videos...); err != nil {
		t.Error(err)
	}
	readInVideos = []youtube.Video{}
	if err := readFileToStruct(outputFile, &readInVideos); err != nil {
		t.Error(err)
	}
	if len(readInVideos) != 4 {
		t.Errorf("should be 4 videos but was %d", len(readInVideos))
	}
	storage = NewStorage(outputFile, inputFile, MERGE)
	if err := storage.SaveVideos(videos...); err != nil {
		t.Error(err)
	}
	if err := readFileToStruct(outputFile, &readInVideos); err != nil {
		t.Error(err)
	}
	if len(readInVideos) != 3 {
		t.Errorf("should be 3 videos but was %d", len(readInVideos))
	}
}
