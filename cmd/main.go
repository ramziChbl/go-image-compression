package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
)

func main() {
	// Override flag.Usage function to print the argument position and format
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [--quality QUALITY] IMAGE_PATH\n", os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}

	imageQuality := flag.Int("quality", 50, "quality of the compressed JPEG image")
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
