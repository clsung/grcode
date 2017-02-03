package grcode

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

type Symbol struct {
	symbol *C.zbar_symbol_t
}

// NewSymbol new the symbol object
func NewSymbol(z *C.zbar_symbol_t) *Symbol {
	return &Symbol{symbol: z}
}

// Next returns next symbol
func (s *Symbol) Next() *Symbol {
	n := C.zbar_symbol_next(s.symbol)
	if n != nil {
		return NewSymbol(n)
	}
	return nil
}

// Data returns the string data of the symbol
func (s *Symbol) Data() string {
	data := C.zbar_symbol_get_data(s.symbol)
	if data != nil {
		return C.GoString(data)
	}
	return ""
}

// Type returns the symbol type
func (s *Symbol) Type() SymbolType {
	return SymbolType{t: C.zbar_symbol_get_type(s.symbol)}
}
