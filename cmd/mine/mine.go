
package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	version = "v1"
	pathContainFolders = "resources/" + version
	standardBounds image.Rectangle
	fromNumber = flag.Int("from", 0, "int type, from index to start")
	toNumber = flag.Int("to", 22, "int type, to index to end")

)
func main() {
	flag.Parse()
	fmt.Printf("from: %d, to: %d\n", *fromNumber, *toNumber)
	res := iterateFiles(pathContainFolders)
	if res != nil {
		for k, v := range res {
			fmt.Printf("res: k = %d, v = %s \n", k, v.FilePath)
		}
	}

	// load file and image.Image
	var resultCount int
	var from []image.Image
	var to [][]image.Image
	for i, v := range res {
		res[i].Imgs = make([]image.Image, len(v.FilePath))
		for j, k := range v.FilePath {
			file,err := os.Open(k)
			if err != nil {
				log.Fatalf("failed to open: %s", err)
			}
			res[i].Imgs[j],err = png.Decode(file)
			if err != nil {
				log.Fatalf("failed to decode: %s", err)
			}
			file.Close()
		}
		if len(from) == 0 {
			from = res[i].Imgs
		} else {
			to = append(to, res[i].Imgs)
		}
		if resultCount == 0 {
			resultCount = len(v.FilePath)
		} else {
			resultCount *= len(v.FilePath)
		}
	}

	standardBounds = res[0].Imgs[0].Bounds()
	resultImgs := make([]image.Image, resultCount)
	fmt.Printf("result number: %d", len(resultImgs))
	if *toNumber == 22 || *toNumber > len(resultImgs) {
		*toNumber = len(resultImgs)
	}

	//// method one combine image together
	//generateMethod1(from, to, resultImgs)

	// method two combine image together
	generateMethod2(from, to)
}


func generateMethod2(from []image.Image, to [][]image.Image) {
	var count int
	var preTime float64 = float64(time.Now().Unix())
	for _, v1 := range from {
		i1 := image.NewRGBA(standardBounds)
		draw.Draw(i1, standardBounds, v1, image.Point{}, draw.Over)
		for _, v2 := range to[0] {
			i2 := image.NewRGBA(standardBounds)
			draw.Draw(i2, standardBounds, v2, image.Point{}, draw.Over)
			for _, v3 := range to[1] {
				i3 := image.NewRGBA(standardBounds)
				draw.Draw(i3, standardBounds, v3, image.Point{}, draw.Over)
				for _, v4 := range to[2] {
					i4 := image.NewRGBA(standardBounds)
					draw.Draw(i4, standardBounds, v4, image.Point{}, draw.Over)
					for _, v5 := range to[3] {
						i5 := image.NewRGBA(standardBounds)
						draw.Draw(i5, standardBounds, v5, image.Point{}, draw.Over)
						for _, v6 := range to[4] {
							i6 := image.NewRGBA(standardBounds)
							draw.Draw(i6, standardBounds, v6, image.Point{}, draw.Over)
							for _, v7 := range to[5] {
								i7 := image.NewRGBA(standardBounds)
								draw.Draw(i7, standardBounds, v7, image.Point{}, draw.Over)
								for _, v8 := range to[6] {
									i8 := image.NewRGBA(standardBounds)
									draw.Draw(i8, standardBounds, v8, image.Point{}, draw.Over)
									{
										if count < *fromNumber {											
											fmt.Printf("neglect index: %d, continuing...\n", count)
											count++
											continue
										}
										if count > *toNumber {											
											return
										}
										final := image.NewRGBA(standardBounds)
										draw.Draw(final, standardBounds, i1, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i2, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i3, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i4, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i5, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i6, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i7, image.Point{}, draw.Over)
										draw.Draw(final, standardBounds, i8, image.Point{}, draw.Over)
										path := "./output/" + version + "/" + strconv.Itoa(count) + ".png"
										result,err := os.Create(fmt.Sprintf(path))
										if err != nil {
											log.Fatalf("failed to create: %s", err)
											return
										}
										png.Encode(result, final)
										result.Close()
										fmt.Printf("created number: %d png, saved to %s\n", count, path)
										if count % 10 == 0 {
											now := float64(time.Now().Unix())
											fmt.Printf("\tlast 10 png generating time: %.2f\n", now - preTime)
											preTime = now
										}
										count++
									}

								}
							}
						}
					}
				}
			}
		}
	}
}

func generateMethod1(from []image.Image, to [][]image.Image, resultImgs []image.Image) {
	ret := fuckIterate(from, to)
	for i, v := range ret {
		third,err := os.Create(fmt.Sprintf("./output/" + version + "/" + strconv.Itoa(i) + ".png"))
		if err != nil {
			log.Fatalf("failed to create: %s", err)
			return
		}
		png.Encode(third, v)
		third.Close()
	}
	fmt.Printf("%d", len(resultImgs))
}

func fuckIterate(imgs []image.Image, imgsArray [][]image.Image) []image.Image {

	res := make([]image.Image, len(imgs) * len(imgsArray[0]))
	i := 0
	for _, k := range imgs {
		for _, v := range imgsArray[0] {
			//res[i] = k + v
			//i++

			x := image.NewRGBA(standardBounds)
			draw.Draw(x, standardBounds, k, image.Point{}, draw.Over)
			draw.Draw(x, standardBounds, v, image.Point{}, draw.Over)
			res[i] = x
			i++
		}
	}
	if len(imgsArray) > 1 {
		return fuckIterate(res, imgsArray[1:][:])
		//return nil
	} else {
		return res
	}
}



type WantedPath struct {
	SubFolder string
	FilePath []string
	Imgs []image.Image
}

func iterateFiles(path string) []*WantedPath {
	wp := make([]*WantedPath, 0)

	//if err := filepath.Walk("/tmp/", func(path string, info os.FileInfo, err error) error {
	if err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s \n", info.IsDir(), p)
		if info.IsDir() {
			pSubPath := strings.Replace(p, path, "", 1)
			if pSubPath != "" {
				wp = append(wp, &WantedPath{
					SubFolder: pSubPath,
					FilePath:  make([]string, 0),
				})
			}
		} else {
			if strings.Contains(p, wp[len(wp) - 1].SubFolder) {
				// put files to previously created []string
				wp[len(wp)-1].FilePath = append(wp[len(wp)-1].FilePath, p)
			}
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
	return wp
}
