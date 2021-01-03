package main
// Huge Shout out to: https://www.sanarias.com/blog/1214PlayingwithimagesinHTTPresponseingolang
// For demonstrating how to hook images up to http.HandleFunc rather concisely
import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"strconv"
	//
	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	http.HandleFunc("/", imageHandler)
	log.Println("Listening on 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

// writeImage encodes an image 'img' in png format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func drawImage() image.Image {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 297, 210.0))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.BeginPath() // Initialize a new path
	gc.MoveTo(10, 10) // Move to a position to start the new path
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	return dest
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	m := drawImage()
	var img image.Image = m
	writeImage(w, &img)
}
