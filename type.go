package grcode

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

// following is reference from zbar.h
const (
	ZBAR_NONE    = 0      /**< no symbol decoded */
	ZBAR_PARTIAL = 1      /**< intermediate status */
	ZBAR_EAN8    = 8      /**< EAN-8 */
	ZBAR_UPCE    = 9      /**< UPC-E */
	ZBAR_ISBN10  = 10     /**< ISBN-10 (from EAN-13). @since 0.4 */
	ZBAR_UPCA    = 12     /**< UPC-A */
	ZBAR_EAN13   = 13     /**< EAN-13 */
	ZBAR_ISBN13  = 14     /**< ISBN-13 (from EAN-13). @since 0.4 */
	ZBAR_I25     = 25     /**< Interleaved 2 of 5. @since 0.4 */
	ZBAR_CODE39  = 39     /**< Code 39. @since 0.4 */
	ZBAR_PDF417  = 57     /**< PDF417. @since 0.6 */
	ZBAR_QRCODE  = 64     /**< QR Code. @since 0.10 */
	ZBAR_CODE128 = 128    /**< Code 128 */
	ZBAR_SYMBOL  = 0x00ff /**< mask for base symbol type */
	ZBAR_ADDON2  = 0x0200 /**< 2-digit add-on flag */
	ZBAR_ADDON5  = 0x0500 /**< 5-digit add-on flag */
	ZBAR_ADDON   = 0x0700 /**< add-on flag mask */
	NONE         = C.ZBAR_NONE
	PARTIAL      = C.ZBAR_PARTIAL
	EAN8         = C.ZBAR_EAN8
	UPCE         = C.ZBAR_UPCE
	ISBN10       = C.ZBAR_ISBN10
	UPCA         = C.ZBAR_UPCA
	EAN13        = C.ZBAR_EAN13
	ISBN13       = C.ZBAR_ISBN13
	I25          = C.ZBAR_I25
	CODE39       = C.ZBAR_CODE39
	PDF414       = C.ZBAR_PDF417
	QRCODE       = C.ZBAR_QRCODE
	CODE128      = C.ZBAR_CODE128
	SYMBOL       = C.ZBAR_SYMBOL
	ADDON2       = C.ZBAR_ADDON2
	ADDON5       = C.ZBAR_ADDON5
	ADDON        = C.ZBAR_ADDON
)

type SymbolType struct {
	t C.zbar_symbol_type_t
}

func (t SymbolType) IsQRCODE() bool {
	if t.t == QRCODE {
		return true
	}
	return false
}
