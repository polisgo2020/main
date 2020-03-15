package main

import (
	"fmt"
	"unsafe"
)

func main() {
	orgSlice := []string{
		"Apple",
		"Orange",
		"Banana",
		"Grape",
		"Plum",
	}

	// capacity is 5
	InspectSlice(orgSlice)
	fmt.Printf("\n\n")

	// capacity is 10, new addresses of the elements
	orgSlice = append(orgSlice, "New el")
	InspectSlice(orgSlice)
	fmt.Printf("\n\n")

	// capacity is 20, new addresses of the elements
	orgSlice = append(orgSlice, "New el 2", "New el 3", "New el 4", "New el 5", "New el 6")
	InspectSlice(orgSlice)
}

// correct returning argument
func test() ([]string, error) {
	return nil, fmt.Errorf("")
}

// INCORRECT returning argument
func test2() (*[]string, error) {
	return nil, fmt.Errorf("")
}

func InspectSlice(slice []string) {
	// Capture the address to the slice structure
	address := unsafe.Pointer(&slice)
	addrSize := unsafe.Sizeof(address)

	// Capture the address where the length and cap size is stored
	lenAddr := uintptr(address) + addrSize
	capAddr := uintptr(address) + (addrSize * 2)

	// Create pointers to the length and cap size
	lenPtr := (*int)(unsafe.Pointer(lenAddr))
	capPtr := (*int)(unsafe.Pointer(capAddr))

	// Create a pointer to the underlying array
	addPtr := (*[11]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)
	fmt.Printf("len(slice)[%d] cap(slice)[%d]\n",
		len(slice),
		cap(slice))

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}
}
