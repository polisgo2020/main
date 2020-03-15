package main

import (
	"fmt"
	"unsafe"
)

func main() {
	orgSlice := make([]string, 5, 8)
	orgSlice[0] = "Apple"
	orgSlice[1] = "Orange"
	orgSlice[2] = "Banana"
	orgSlice[3] = "Grape"
	orgSlice[4] = "Plum"

	InspectSlice(orgSlice)
	fmt.Printf("\n\n")
	//capacity will not be enlarged
	InspectSlice(append(orgSlice, "New element"))
	fmt.Printf("\n\n")
	// new slice pointer will be equal to `Banana` address.
	slice2 := orgSlice[2:4]
	InspectSlice(slice2)

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
	addPtr := (*[8]string)(unsafe.Pointer(*(*uintptr)(address)))

	fmt.Printf("Slice Addr[%p] Len Addr[0x%x] Cap Addr[0x%x]\n",
		address,
		lenAddr,
		capAddr)

	fmt.Printf("Slice Length[%d] Cap[%d]\n",
		*lenPtr,
		*capPtr)

	for index := 0; index < *lenPtr; index++ {
		fmt.Printf("[%d] %p %s\n",
			index,
			&(*addPtr)[index],
			(*addPtr)[index])
	}
}
