package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)
var (
	ImageFilePath1 =   "./resources/v1/background/渐变背景1.png"
	ImageFilePath2 =   "./resources/v1/body/身体填充1.png"
)
func TestSimple(t *testing.T) {
	image1,err := os.Open(ImageFilePath1)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	image2,err := os.Open(ImageFilePath2)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	second,err := png.Decode(image2)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image2.Close()

	offset := image.Pt(300, 200)
	b := first.Bounds()
	image3 := image.NewRGBA(b)
	draw.Draw(image3, b, first, image.ZP, draw.Src)
	draw.Draw(image3, second.Bounds().Add(offset), second, image.ZP, draw.Over)

	third,err := os.Create("./output/test_result1.jpg")
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{jpeg.DefaultQuality})
	defer third.Close()
}


func Test_Iterate(t *testing.T) {
	str1 := []string {"1", "2", "3"}
	str2 := []string {"a", "b", "c"}
	str3 := []string {"x", "y", "z"}
	res := iterates(str1, [][]string{str2, str3})
	for i, v := range res {
		fmt.Printf("index: %d, value: %s\n", i, v)
	}
}

func iterates(ints []string, intsArray [][]string) []string {

	res := make([]string, len(ints) * len(intsArray[0]))
	i := 0
	for _, k := range ints {
		for _, v := range intsArray[0] {
			res[i] = k + v
			i++
		}
	}
	if len(intsArray) > 1 {
		return iterates(res, intsArray[1:][:])
		//return nil
	} else {
		return res
	}

}