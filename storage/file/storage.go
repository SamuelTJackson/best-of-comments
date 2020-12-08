package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SamuelTJackson/best-of-comments/youtube"
	"io"
	"io/ioutil"
	"os"
)
type Mode int
const (
	OVERWRITE Mode = iota
	APPEND
	MERGE
)
type storage struct {
	outputFile string
	inputFile string
	mode Mode
}
func NewStorage(outputFile string, inputFile string, mode Mode) *storage {
	return &storage{outputFile: outputFile, inputFile: inputFile, mode: mode}
}

func readFileToStruct(fileName string, value interface{}) error {
	file, err := os.Open(fileName)
	if err != nil {
		if _, ok := err.(*os.PathError); ok {
			return  nil
		}
		return err
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&value); err != nil {
		if err == io.EOF {
			return  nil
		}
		return  err
	}
	return  nil
}

func (s storage) SaveVideos(v ...youtube.Video) error {
	var localVideos []youtube.Video
	if err := readFileToStruct(s.inputFile, &localVideos); err != nil {
		return err
	}
	if s.mode == OVERWRITE {
		localVideos = v
	} else if s.mode == APPEND {
		localVideos = append(localVideos, v...)
	} else if s.mode == MERGE {
		var uniqueVideos []youtube.Video
		for _, newVideo := range v {
			insert := true
			for _, localVideo := range localVideos {
				if newVideo.GetID() == localVideo.ID	{
					insert = false
					break
				}
			}
			if insert {
				uniqueVideos = append(uniqueVideos, newVideo)
			}
		}
		localVideos = append(localVideos, uniqueVideos...)
	} else {
		return errors.New(fmt.Sprintf("Mode %v not recognized", s.mode))
	}

	data, err := json.MarshalIndent(localVideos, "", " ")
	if err != nil {
		return err
	}
	fmt.Printf("saving %d videos to file %s\n", len(localVideos), s.outputFile)
	return ioutil.WriteFile(s.outputFile,data, 0644)
}
func (s storage) GetVideo(id string) (*youtube.Video, error)  {
	//videos, err := readFileToStruct(s.inputFile)
	//if err != nil {
	//	return nil, err
	//}
	//for _, video := range videos {
	//	if video.ID == id {
	//		return video, nil
	//	}
	//}

	return nil, errors.New(fmt.Sprintf("there is no video with id %s", id))
}
