package grcode

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import "image"
import "image/draw"

import "unsafe"
import "runtime"

type ZbarImage struct {
	image *C.zbar_image_t
}

// VX_DF_IMAGE_Y800 is grayscale.
// Y800, GRAY is required
const VX_DF_IMAGE_Y800 = 0x30303859

// NewSymbol new the zbar image object
func NewZbarImage(img image.Image) *ZbarImage {
	bounds := img.Bounds()
	w := bounds.Max.X - bounds.Min.X
	h := bounds.Max.Y - bounds.Min.Y

	zImg := &ZbarImage{image: C.zbar_image_create()}
	// draw image to zbar image with garyscale
	gray := image.NewGray(bounds)
	draw.Draw(gray, bounds, img, image.ZP, draw.Over)

	C.zbar_image_set_format(zImg.image, C.ulong(VX_DF_IMAGE_Y800))
	C.zbar_image_set_size(zImg.image, C.uint(w), C.uint(h))
	C.zbar_image_set_data(zImg.image, unsafe.Pointer(&gray.Pix[0]), C.ulong(len(gray.Pix)), nil)

	// finalizer
	runtime.SetFinalizer(zImg, (*ZbarImage).Destroy)
	return zImg
}

// GetSymbol returns the first symbol of result
// use symbol.Next() to retrive next
func (z *ZbarImage) GetSymbol() *Symbol {
	s := C.zbar_image_first_symbol(z.image)
	if s == nil {
		return nil
	}
	return NewSymbol(s)
}

// Destroy suicides
func (z *ZbarImage) Destroy() {
	C.zbar_image_destroy(z.image)
}
