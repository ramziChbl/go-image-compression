package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	imageQuality := flag.Int("quality", 50, "quality of the compressed JPEG image")
	if len(os.Args) <= 1 {
		fmt.Fprintf(os.Stderr, "No image path was given as an argument\n")
		os.Exit(1)
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	inputPath := flag.Arg(0)

	outputPath, err := compressJPEG(inputPath, *imageQuality)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Fprintf(os.Stdout, "%s\n", outputPath)
}

func compressJPEG(inputPath string, imageQuality int) (string, error) {
	outputPath := "out.jpeg"
	inReader, err := os.Open(inputPath)
	if err != nil {
		return "", err
	}
	defer inReader.Close()

	jpegImage, err := jpeg.Decode(inReader)
	if err != nil {
		return "", err
	}

	outWriter, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer outWriter.Close()

	err = jpeg.Encode(outWriter, jpegImage, &jpeg.Options{Quality: imageQuality})
	if err != nil {
		return "", err
	}
	return outputPath, nil
}
