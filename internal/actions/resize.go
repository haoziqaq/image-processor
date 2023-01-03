package actions

import (
	"github.com/disintegration/imaging"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"path"
)

type ResizeActionOption struct {
	Input  string
	Output string
	Width  int
	Height int
}

func prepareResizeAction(context *cli.Context) *ResizeActionOption {
	cwd, _ := os.Getwd()
	width := context.Int("width")
	height := context.Int("height")
	input := cwd
	output := path.Join(cwd, "output")

	outputPath, err := os.Stat(output)

	if err == nil && outputPath.IsDir() {
		log.Fatalln("Output path is existed")
	}

	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(output, os.ModePerm)
			if err != nil {
				log.Fatalln("Output mkdir error")
			}
		} else {
			log.Fatalln(err)
		}
	}

	return &ResizeActionOption{
		Width:  width,
		Height: height,
		Input:  input,
		Output: output,
	}
}

func Resize(context *cli.Context) {
	option := prepareResizeAction(context)
	files, err := os.ReadDir(option.Input)
	if err != nil {
		log.Fatalf("Failed to read dir: %v", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filename := path.Join(option.Input, file.Name())
		newFilename := path.Join(option.Output, file.Name())
		image, err := imaging.Open(filename)

		if err != nil {
			log.Fatalf("Failed to open image: %v", err)
			return
		}

		image = imaging.Resize(image, option.Width, option.Height, imaging.Lanczos)
		err = imaging.Save(image, newFilename)

		if err != nil {
			log.Fatalf("Failed to open image: %v", err)
		}

		log.Printf("Process image success %v -> %v", filename, newFilename)
	}

	log.Println("All image process success!")
}
