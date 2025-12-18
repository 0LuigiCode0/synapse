package sys

/*
#include <stdio.h>
#include <stdint.h>

void c9(uint64_t a1, uint64_t a2, uint64_t a3, uint64_t a4,uint64_t a5, uint64_t a6, uint64_t a7, uint64_t a8, uint64_t a9) {
        printf("Args: %lu %lu %lu %lu %lu %lu %lu %lu %lu\n",a1, a2, a3, a4, a5, a6, a7, a8, a9);
        fflush(stdout);
}
void c12(uint64_t a1, uint64_t a2, uint64_t a3, uint64_t a4,uint64_t a5, uint64_t a6, uint64_t a7, uint64_t a8, uint64_t a9,uint64_t a10,uint64_t a11,uint64_t a12) {
        printf("Args: %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu\n",a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
        fflush(stdout);
}
void c15(uint64_t a1, uint64_t a2, uint64_t a3, uint64_t a4,uint64_t a5, uint64_t a6, uint64_t a7, uint64_t a8, uint64_t a9,uint64_t a10,uint64_t a11,uint64_t a12,uint64_t a13,uint64_t a14,uint64_t a15) {
        printf("Args: %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu %lu\n",a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13 ,a14 ,a15);
        fflush(stdout);
}
*/
import "C"
import "unsafe"

var (
	addrC9  = uintptr(unsafe.Pointer(C.c9))
	addrC12 = uintptr(unsafe.Pointer(C.c12))
	addrC15 = uintptr(unsafe.Pointer(C.c15))
)
