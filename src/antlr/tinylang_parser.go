// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // TinyLang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 348,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 3, 2,
	7, 2, 68, 10, 2, 12, 2, 14, 2, 71, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 7, 5, 91, 10, 5, 12, 5, 14, 5, 94, 11, 5, 5, 5, 96, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 104, 10, 6, 3, 7, 3, 7, 5, 7, 108,
	10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 118, 10, 9,
	12, 9, 14, 9, 121, 11, 9, 5, 9, 123, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10,
	128, 10, 10, 3, 11, 3, 11, 5, 11, 132, 10, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 140, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 150, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 162, 10, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 5, 16, 168, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 7, 19, 182, 10, 19, 12, 19,
	14, 19, 185, 11, 19, 3, 19, 3, 19, 3, 19, 7, 19, 190, 10, 19, 12, 19, 14,
	19, 193, 11, 19, 3, 19, 3, 19, 3, 19, 7, 19, 198, 10, 19, 12, 19, 14, 19,
	201, 11, 19, 5, 19, 203, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 209,
	10, 20, 3, 20, 3, 20, 3, 20, 5, 20, 214, 10, 20, 3, 20, 3, 20, 3, 20, 5,
	20, 219, 10, 20, 3, 20, 3, 20, 5, 20, 223, 10, 20, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 21, 5, 21, 230, 10, 21, 3, 22, 3, 22, 7, 22, 234, 10, 22, 12,
	22, 14, 22, 237, 11, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	5, 24, 246, 10, 24, 3, 24, 3, 24, 3, 24, 5, 24, 251, 10, 24, 3, 24, 3,
	24, 3, 24, 5, 24, 256, 10, 24, 3, 24, 3, 24, 5, 24, 260, 10, 24, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 267, 10, 25, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 7, 27, 279, 10, 27, 12, 27,
	14, 27, 282, 11, 27, 3, 28, 3, 28, 7, 28, 286, 10, 28, 12, 28, 14, 28,
	289, 11, 28, 3, 28, 5, 28, 292, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 301, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 306,
	10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31,
	316, 10, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 5, 32, 327, 10, 32, 3, 32, 3, 32, 3, 32, 5, 32, 332, 10, 32, 3, 32,
	3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 2, 2, 34, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 2, 8, 4, 2, 23, 23, 34, 34, 4, 2, 29, 31, 33, 34, 5, 2,
	24, 24, 26, 27, 37, 37, 4, 2, 13, 13, 29, 29, 4, 2, 24, 24, 26, 27, 3,
	2, 18, 19, 2, 382, 2, 69, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2, 6, 84, 3, 2, 2,
	2, 8, 95, 3, 2, 2, 2, 10, 97, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109,
	3, 2, 2, 2, 16, 122, 3, 2, 2, 2, 18, 127, 3, 2, 2, 2, 20, 131, 3, 2, 2,
	2, 22, 141, 3, 2, 2, 2, 24, 151, 3, 2, 2, 2, 26, 153, 3, 2, 2, 2, 28, 155,
	3, 2, 2, 2, 30, 167, 3, 2, 2, 2, 32, 169, 3, 2, 2, 2, 34, 174, 3, 2, 2,
	2, 36, 202, 3, 2, 2, 2, 38, 222, 3, 2, 2, 2, 40, 229, 3, 2, 2, 2, 42, 231,
	3, 2, 2, 2, 44, 238, 3, 2, 2, 2, 46, 259, 3, 2, 2, 2, 48, 266, 3, 2, 2,
	2, 50, 268, 3, 2, 2, 2, 52, 280, 3, 2, 2, 2, 54, 283, 3, 2, 2, 2, 56, 293,
	3, 2, 2, 2, 58, 309, 3, 2, 2, 2, 60, 312, 3, 2, 2, 2, 62, 319, 3, 2, 2,
	2, 64, 335, 3, 2, 2, 2, 66, 68, 5, 4, 3, 2, 67, 66, 3, 2, 2, 2, 68, 71,
	3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2,
	71, 69, 3, 2, 2, 2, 72, 73, 7, 2, 2, 3, 73, 3, 3, 2, 2, 2, 74, 75, 7, 3,
	2, 2, 75, 76, 7, 34, 2, 2, 76, 77, 7, 4, 2, 2, 77, 78, 5, 8, 5, 2, 78,
	79, 7, 5, 2, 2, 79, 80, 5, 12, 7, 2, 80, 81, 7, 6, 2, 2, 81, 82, 5, 52,
	27, 2, 82, 83, 7, 7, 2, 2, 83, 5, 3, 2, 2, 2, 84, 85, 7, 34, 2, 2, 85,
	86, 5, 26, 14, 2, 86, 7, 3, 2, 2, 2, 87, 92, 5, 6, 4, 2, 88, 89, 7, 8,
	2, 2, 89, 91, 5, 6, 4, 2, 90, 88, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2,
	95, 87, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 9, 3, 2, 2, 2, 97, 103, 7,
	17, 2, 2, 98, 104, 5, 28, 15, 2, 99, 104, 5, 14, 8, 2, 100, 104, 5, 24,
	13, 2, 101, 104, 5, 38, 20, 2, 102, 104, 5, 42, 22, 2, 103, 98, 3, 2, 2,
	2, 103, 99, 3, 2, 2, 2, 103, 100, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103,
	102, 3, 2, 2, 2, 104, 11, 3, 2, 2, 2, 105, 108, 5, 26, 14, 2, 106, 108,
	7, 25, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 13, 3, 2,
	2, 2, 109, 110, 9, 2, 2, 2, 110, 111, 7, 4, 2, 2, 111, 112, 5, 16, 9, 2,
	112, 113, 7, 5, 2, 2, 113, 15, 3, 2, 2, 2, 114, 119, 5, 18, 10, 2, 115,
	116, 7, 8, 2, 2, 116, 118, 5, 18, 10, 2, 117, 115, 3, 2, 2, 2, 118, 121,
	3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 123, 3, 2,
	2, 2, 121, 119, 3, 2, 2, 2, 122, 114, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2,
	123, 17, 3, 2, 2, 2, 124, 128, 5, 24, 13, 2, 125, 128, 5, 14, 8, 2, 126,
	128, 5, 28, 15, 2, 127, 124, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 126,
	3, 2, 2, 2, 128, 19, 3, 2, 2, 2, 129, 132, 7, 34, 2, 2, 130, 132, 5, 28,
	15, 2, 131, 129, 3, 2, 2, 2, 131, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2,
	133, 139, 7, 9, 2, 2, 134, 140, 5, 38, 20, 2, 135, 140, 5, 14, 8, 2, 136,
	140, 5, 24, 13, 2, 137, 140, 5, 30, 16, 2, 138, 140, 5, 28, 15, 2, 139,
	134, 3, 2, 2, 2, 139, 135, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 139, 137,
	3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 21, 3, 2, 2, 2, 141, 142, 7, 10,
	2, 2, 142, 143, 7, 34, 2, 2, 143, 144, 5, 26, 14, 2, 144, 149, 7, 9, 2,
	2, 145, 150, 5, 38, 20, 2, 146, 150, 5, 24, 13, 2, 147, 150, 5, 30, 16,
	2, 148, 150, 5, 14, 8, 2, 149, 145, 3, 2, 2, 2, 149, 146, 3, 2, 2, 2, 149,
	147, 3, 2, 2, 2, 149, 148, 3, 2, 2, 2, 150, 23, 3, 2, 2, 2, 151, 152, 9,
	3, 2, 2, 152, 25, 3, 2, 2, 2, 153, 154, 9, 4, 2, 2, 154, 27, 3, 2, 2, 2,
	155, 156, 7, 34, 2, 2, 156, 161, 7, 11, 2, 2, 157, 162, 7, 29, 2, 2, 158,
	162, 7, 34, 2, 2, 159, 162, 5, 14, 8, 2, 160, 162, 5, 38, 20, 2, 161, 157,
	3, 2, 2, 2, 161, 158, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 160, 3, 2,
	2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 7, 12, 2, 2, 164, 29, 3, 2, 2, 2,
	165, 168, 5, 34, 18, 2, 166, 168, 5, 32, 17, 2, 167, 165, 3, 2, 2, 2, 167,
	166, 3, 2, 2, 2, 168, 31, 3, 2, 2, 2, 169, 170, 7, 11, 2, 2, 170, 171,
	9, 5, 2, 2, 171, 172, 7, 12, 2, 2, 172, 173, 9, 6, 2, 2, 173, 33, 3, 2,
	2, 2, 174, 175, 7, 6, 2, 2, 175, 176, 5, 36, 19, 2, 176, 177, 7, 7, 2,
	2, 177, 35, 3, 2, 2, 2, 178, 183, 7, 32, 2, 2, 179, 180, 7, 8, 2, 2, 180,
	182, 7, 32, 2, 2, 181, 179, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181,
	3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 203, 3, 2, 2, 2, 185, 183, 3, 2,
	2, 2, 186, 191, 7, 33, 2, 2, 187, 188, 7, 8, 2, 2, 188, 190, 7, 33, 2,
	2, 189, 187, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191,
	192, 3, 2, 2, 2, 192, 203, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 199,
	7, 29, 2, 2, 195, 196, 7, 8, 2, 2, 196, 198, 7, 29, 2, 2, 197, 195, 3,
	2, 2, 2, 198, 201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2,
	2, 200, 203, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 178, 3, 2, 2, 2, 202,
	186, 3, 2, 2, 2, 202, 194, 3, 2, 2, 2, 203, 37, 3, 2, 2, 2, 204, 205, 5,
	40, 21, 2, 205, 208, 7, 38, 2, 2, 206, 209, 5, 40, 21, 2, 207, 209, 5,
	38, 20, 2, 208, 206, 3, 2, 2, 2, 208, 207, 3, 2, 2, 2, 209, 223, 3, 2,
	2, 2, 210, 213, 7, 4, 2, 2, 211, 214, 5, 40, 21, 2, 212, 214, 5, 38, 20,
	2, 213, 211, 3, 2, 2, 2, 213, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215,
	218, 7, 38, 2, 2, 216, 219, 5, 40, 21, 2, 217, 219, 5, 38, 20, 2, 218,
	216, 3, 2, 2, 2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221,
	7, 5, 2, 2, 221, 223, 3, 2, 2, 2, 222, 204, 3, 2, 2, 2, 222, 210, 3, 2,
	2, 2, 223, 39, 3, 2, 2, 2, 224, 230, 5, 14, 8, 2, 225, 230, 7, 29, 2, 2,
	226, 230, 7, 34, 2, 2, 227, 230, 7, 33, 2, 2, 228, 230, 5, 28, 15, 2, 229,
	224, 3, 2, 2, 2, 229, 225, 3, 2, 2, 2, 229, 226, 3, 2, 2, 2, 229, 227,
	3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230, 41, 3, 2, 2, 2, 231, 235, 5, 46,
	24, 2, 232, 234, 5, 44, 23, 2, 233, 232, 3, 2, 2, 2, 234, 237, 3, 2, 2,
	2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2, 236, 43, 3, 2, 2, 2, 237,
	235, 3, 2, 2, 2, 238, 239, 7, 40, 2, 2, 239, 240, 5, 46, 24, 2, 240, 45,
	3, 2, 2, 2, 241, 242, 5, 48, 25, 2, 242, 245, 7, 39, 2, 2, 243, 246, 5,
	48, 25, 2, 244, 246, 5, 42, 22, 2, 245, 243, 3, 2, 2, 2, 245, 244, 3, 2,
	2, 2, 246, 260, 3, 2, 2, 2, 247, 250, 7, 4, 2, 2, 248, 251, 5, 48, 25,
	2, 249, 251, 5, 42, 22, 2, 250, 248, 3, 2, 2, 2, 250, 249, 3, 2, 2, 2,
	251, 252, 3, 2, 2, 2, 252, 255, 7, 39, 2, 2, 253, 256, 5, 48, 25, 2, 254,
	256, 5, 42, 22, 2, 255, 253, 3, 2, 2, 2, 255, 254, 3, 2, 2, 2, 256, 257,
	3, 2, 2, 2, 257, 258, 7, 5, 2, 2, 258, 260, 3, 2, 2, 2, 259, 241, 3, 2,
	2, 2, 259, 247, 3, 2, 2, 2, 260, 47, 3, 2, 2, 2, 261, 267, 5, 38, 20, 2,
	262, 267, 5, 40, 21, 2, 263, 267, 7, 30, 2, 2, 264, 267, 7, 31, 2, 2, 265,
	267, 5, 28, 15, 2, 266, 261, 3, 2, 2, 2, 266, 262, 3, 2, 2, 2, 266, 263,
	3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266, 265, 3, 2, 2, 2, 267, 49, 3, 2,
	2, 2, 268, 269, 9, 7, 2, 2, 269, 51, 3, 2, 2, 2, 270, 279, 5, 22, 12, 2,
	271, 279, 5, 20, 11, 2, 272, 279, 5, 14, 8, 2, 273, 279, 5, 54, 28, 2,
	274, 279, 5, 62, 32, 2, 275, 279, 5, 64, 33, 2, 276, 279, 5, 50, 26, 2,
	277, 279, 5, 10, 6, 2, 278, 270, 3, 2, 2, 2, 278, 271, 3, 2, 2, 2, 278,
	272, 3, 2, 2, 2, 278, 273, 3, 2, 2, 2, 278, 274, 3, 2, 2, 2, 278, 275,
	3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 277, 3, 2, 2, 2, 279, 282, 3, 2,
	2, 2, 280, 278, 3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 53, 3, 2, 2, 2,
	282, 280, 3, 2, 2, 2, 283, 287, 5, 56, 29, 2, 284, 286, 5, 58, 30, 2, 285,
	284, 3, 2, 2, 2, 286, 289, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288,
	3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 290, 292, 5, 60,
	31, 2, 291, 290, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 55, 3, 2, 2, 2,
	293, 294, 7, 20, 2, 2, 294, 300, 7, 4, 2, 2, 295, 301, 7, 30, 2, 2, 296,
	301, 7, 31, 2, 2, 297, 301, 5, 42, 22, 2, 298, 301, 7, 34, 2, 2, 299, 301,
	5, 14, 8, 2, 300, 295, 3, 2, 2, 2, 300, 296, 3, 2, 2, 2, 300, 297, 3, 2,
	2, 2, 300, 298, 3, 2, 2, 2, 300, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2,
	302, 303, 7, 5, 2, 2, 303, 305, 7, 6, 2, 2, 304, 306, 5, 52, 27, 2, 305,
	304, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308,
	7, 7, 2, 2, 308, 57, 3, 2, 2, 2, 309, 310, 7, 22, 2, 2, 310, 311, 5, 56,
	29, 2, 311, 59, 3, 2, 2, 2, 312, 313, 7, 22, 2, 2, 313, 315, 7, 6, 2, 2,
	314, 316, 5, 52, 27, 2, 315, 314, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316,
	317, 3, 2, 2, 2, 317, 318, 7, 7, 2, 2, 318, 61, 3, 2, 2, 2, 319, 320, 7,
	16, 2, 2, 320, 326, 7, 4, 2, 2, 321, 327, 7, 30, 2, 2, 322, 327, 7, 31,
	2, 2, 323, 327, 5, 42, 22, 2, 324, 327, 7, 34, 2, 2, 325, 327, 5, 14, 8,
	2, 326, 321, 3, 2, 2, 2, 326, 322, 3, 2, 2, 2, 326, 323, 3, 2, 2, 2, 326,
	324, 3, 2, 2, 2, 326, 325, 3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 329,
	7, 5, 2, 2, 329, 331, 7, 6, 2, 2, 330, 332, 5, 52, 27, 2, 331, 330, 3,
	2, 2, 2, 331, 332, 3, 2, 2, 2, 332, 333, 3, 2, 2, 2, 333, 334, 7, 7, 2,
	2, 334, 63, 3, 2, 2, 2, 335, 336, 7, 21, 2, 2, 336, 337, 7, 4, 2, 2, 337,
	338, 5, 22, 12, 2, 338, 339, 7, 14, 2, 2, 339, 340, 5, 42, 22, 2, 340,
	341, 7, 14, 2, 2, 341, 342, 5, 20, 11, 2, 342, 343, 7, 5, 2, 2, 343, 344,
	7, 6, 2, 2, 344, 345, 5, 52, 27, 2, 345, 346, 7, 7, 2, 2, 346, 65, 3, 2,
	2, 2, 39, 69, 92, 95, 103, 107, 119, 122, 127, 131, 139, 149, 161, 167,
	183, 191, 199, 202, 208, 213, 218, 222, 229, 235, 245, 250, 255, 259, 266,
	278, 280, 287, 291, 300, 305, 315, 326, 331,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "')'", "'{'", "'}'", "','", "'='", "'var'", "'['",
	"']'", "'..'", "';'", "", "'while'", "'return'", "'break'", "'continue'",
	"'if'", "'for'", "'else'", "", "'string'", "'void'", "'num'", "'bool'",
	"'null'", "", "'true'", "'false'", "", "", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "WHILE", "RETURN",
	"BREAK", "CONTINUE", "IF", "FOR", "ELSE", "SYS_FUNC", "STRING", "VOID",
	"NUM", "BOOL", "NULL", "NUMBER", "TRUE", "FALSE", "BOOL_VAL", "STRING_RAW",
	"ITEM", "SQ", "COMMENT", "ARRAY", "NUM_SIGN", "BOOL_SIGN", "BOOL_PL",
}

var ruleNames = []string{
	"file", "funcInit", "funcArg", "funcArgs", "funcReturn", "funcReturnType",
	"funcInvoc", "funcArgsInvoc", "funcInvocArgs", "updVariable", "newVariable",
	"val", "variableType", "arrayElem", "arrayInit", "arrayInitEmpty", "arrayInitValue",
	"arrayInitElems", "expr", "exprOperand", "boolExpr", "boolExprSingleExtra",
	"boolExprSingle", "boolExprOperand", "breakOrContinue", "statementBody",
	"ifElseSt", "ifSt", "elseIfSt", "elseSt", "whileSt", "forSt",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TinyLangParser struct {
	*antlr.BaseParser
}

func NewTinyLangParser(input antlr.TokenStream) *TinyLangParser {
	this := new(TinyLangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "TinyLang.g4"

	return this
}

// TinyLangParser tokens.
const (
	TinyLangParserEOF        = antlr.TokenEOF
	TinyLangParserT__0       = 1
	TinyLangParserT__1       = 2
	TinyLangParserT__2       = 3
	TinyLangParserT__3       = 4
	TinyLangParserT__4       = 5
	TinyLangParserT__5       = 6
	TinyLangParserT__6       = 7
	TinyLangParserT__7       = 8
	TinyLangParserT__8       = 9
	TinyLangParserT__9       = 10
	TinyLangParserT__10      = 11
	TinyLangParserT__11      = 12
	TinyLangParserWS         = 13
	TinyLangParserWHILE      = 14
	TinyLangParserRETURN     = 15
	TinyLangParserBREAK      = 16
	TinyLangParserCONTINUE   = 17
	TinyLangParserIF         = 18
	TinyLangParserFOR        = 19
	TinyLangParserELSE       = 20
	TinyLangParserSYS_FUNC   = 21
	TinyLangParserSTRING     = 22
	TinyLangParserVOID       = 23
	TinyLangParserNUM        = 24
	TinyLangParserBOOL       = 25
	TinyLangParserNULL       = 26
	TinyLangParserNUMBER     = 27
	TinyLangParserTRUE       = 28
	TinyLangParserFALSE      = 29
	TinyLangParserBOOL_VAL   = 30
	TinyLangParserSTRING_RAW = 31
	TinyLangParserITEM       = 32
	TinyLangParserSQ         = 33
	TinyLangParserCOMMENT    = 34
	TinyLangParserARRAY      = 35
	TinyLangParserNUM_SIGN   = 36
	TinyLangParserBOOL_SIGN  = 37
	TinyLangParserBOOL_PL    = 38
)

// TinyLangParser rules.
const (
	TinyLangParserRULE_file                = 0
	TinyLangParserRULE_funcInit            = 1
	TinyLangParserRULE_funcArg             = 2
	TinyLangParserRULE_funcArgs            = 3
	TinyLangParserRULE_funcReturn          = 4
	TinyLangParserRULE_funcReturnType      = 5
	TinyLangParserRULE_funcInvoc           = 6
	TinyLangParserRULE_funcArgsInvoc       = 7
	TinyLangParserRULE_funcInvocArgs       = 8
	TinyLangParserRULE_updVariable         = 9
	TinyLangParserRULE_newVariable         = 10
	TinyLangParserRULE_val                 = 11
	TinyLangParserRULE_variableType        = 12
	TinyLangParserRULE_arrayElem           = 13
	TinyLangParserRULE_arrayInit           = 14
	TinyLangParserRULE_arrayInitEmpty      = 15
	TinyLangParserRULE_arrayInitValue      = 16
	TinyLangParserRULE_arrayInitElems      = 17
	TinyLangParserRULE_expr                = 18
	TinyLangParserRULE_exprOperand         = 19
	TinyLangParserRULE_boolExpr            = 20
	TinyLangParserRULE_boolExprSingleExtra = 21
	TinyLangParserRULE_boolExprSingle      = 22
	TinyLangParserRULE_boolExprOperand     = 23
	TinyLangParserRULE_breakOrContinue     = 24
	TinyLangParserRULE_statementBody       = 25
	TinyLangParserRULE_ifElseSt            = 26
	TinyLangParserRULE_ifSt                = 27
	TinyLangParserRULE_elseIfSt            = 28
	TinyLangParserRULE_elseSt              = 29
	TinyLangParserRULE_whileSt             = 30
	TinyLangParserRULE_forSt               = 31
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_file
	return p
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(TinyLangParserEOF, 0)
}

func (s *FileContext) AllFuncInit() []IFuncInitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncInitContext)(nil)).Elem())
	var tst = make([]IFuncInitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncInitContext)
		}
	}

	return tst
}

func (s *FileContext) FuncInit(i int) IFuncInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncInitContext)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFile(s)
	}
}

func (s *FileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TinyLangParserRULE_file)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLangParserT__0 {
		{
			p.SetState(64)
			p.FuncInit()
		}

		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(70)
		p.Match(TinyLangParserEOF)
	}

	return localctx
}

// IFuncInitContext is an interface to support dynamic dispatch.
type IFuncInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncInitContext differentiates from other interfaces.
	IsFuncInitContext()
}

type FuncInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncInitContext() *FuncInitContext {
	var p = new(FuncInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcInit
	return p
}

func (*FuncInitContext) IsFuncInitContext() {}

func NewFuncInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncInitContext {
	var p = new(FuncInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcInit

	return p
}

func (s *FuncInitContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncInitContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *FuncInitContext) FuncArgs() IFuncArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncArgsContext)
}

func (s *FuncInitContext) FuncReturnType() IFuncReturnTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnTypeContext)
}

func (s *FuncInitContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *FuncInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncInit(s)
	}
}

func (s *FuncInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncInit(s)
	}
}

func (s *FuncInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncInit() (localctx IFuncInitContext) {
	localctx = NewFuncInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TinyLangParserRULE_funcInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(72)
		p.Match(TinyLangParserT__0)
	}
	{
		p.SetState(73)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(74)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(75)
		p.FuncArgs()
	}
	{
		p.SetState(76)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(77)
		p.FuncReturnType()
	}
	{
		p.SetState(78)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(79)
		p.StatementBody()
	}
	{
		p.SetState(80)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}

// IFuncArgContext is an interface to support dynamic dispatch.
type IFuncArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgContext differentiates from other interfaces.
	IsFuncArgContext()
}

type FuncArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgContext() *FuncArgContext {
	var p = new(FuncArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcArg
	return p
}

func (*FuncArgContext) IsFuncArgContext() {}

func NewFuncArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgContext {
	var p = new(FuncArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcArg

	return p
}

func (s *FuncArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *FuncArgContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FuncArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncArg(s)
	}
}

func (s *FuncArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncArg(s)
	}
}

func (s *FuncArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncArg() (localctx IFuncArgContext) {
	localctx = NewFuncArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TinyLangParserRULE_funcArg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(83)
		p.VariableType()
	}

	return localctx
}

// IFuncArgsContext is an interface to support dynamic dispatch.
type IFuncArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgsContext differentiates from other interfaces.
	IsFuncArgsContext()
}

type FuncArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsContext() *FuncArgsContext {
	var p = new(FuncArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcArgs
	return p
}

func (*FuncArgsContext) IsFuncArgsContext() {}

func NewFuncArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsContext {
	var p = new(FuncArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcArgs

	return p
}

func (s *FuncArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsContext) AllFuncArg() []IFuncArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncArgContext)(nil)).Elem())
	var tst = make([]IFuncArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncArgContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) FuncArg(i int) IFuncArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncArgContext)
}

func (s *FuncArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncArgs(s)
	}
}

func (s *FuncArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncArgs(s)
	}
}

func (s *FuncArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncArgs() (localctx IFuncArgsContext) {
	localctx = NewFuncArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TinyLangParserRULE_funcArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserITEM {
		{
			p.SetState(85)
			p.FuncArg()
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(86)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(87)
				p.FuncArg()
			}

			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFuncReturnContext is an interface to support dynamic dispatch.
type IFuncReturnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncReturnContext differentiates from other interfaces.
	IsFuncReturnContext()
}

type FuncReturnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncReturnContext() *FuncReturnContext {
	var p = new(FuncReturnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcReturn
	return p
}

func (*FuncReturnContext) IsFuncReturnContext() {}

func NewFuncReturnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnContext {
	var p = new(FuncReturnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcReturn

	return p
}

func (s *FuncReturnContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TinyLangParserRETURN, 0)
}

func (s *FuncReturnContext) ArrayElem() IArrayElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayElemContext)
}

func (s *FuncReturnContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *FuncReturnContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *FuncReturnContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FuncReturnContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *FuncReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncReturn(s)
	}
}

func (s *FuncReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncReturn(s)
	}
}

func (s *FuncReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncReturn() (localctx IFuncReturnContext) {
	localctx = NewFuncReturnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TinyLangParserRULE_funcReturn)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(TinyLangParserRETURN)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(96)
			p.ArrayElem()
		}

	case 2:
		{
			p.SetState(97)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(98)
			p.Val()
		}

	case 4:
		{
			p.SetState(99)
			p.Expr()
		}

	case 5:
		{
			p.SetState(100)
			p.BoolExpr()
		}

	}

	return localctx
}

// IFuncReturnTypeContext is an interface to support dynamic dispatch.
type IFuncReturnTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncReturnTypeContext differentiates from other interfaces.
	IsFuncReturnTypeContext()
}

type FuncReturnTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncReturnTypeContext() *FuncReturnTypeContext {
	var p = new(FuncReturnTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcReturnType
	return p
}

func (*FuncReturnTypeContext) IsFuncReturnTypeContext() {}

func NewFuncReturnTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncReturnTypeContext {
	var p = new(FuncReturnTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcReturnType

	return p
}

func (s *FuncReturnTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncReturnTypeContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FuncReturnTypeContext) VOID() antlr.TerminalNode {
	return s.GetToken(TinyLangParserVOID, 0)
}

func (s *FuncReturnTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncReturnTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncReturnTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncReturnType(s)
	}
}

func (s *FuncReturnTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncReturnType(s)
	}
}

func (s *FuncReturnTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncReturnType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncReturnType() (localctx IFuncReturnTypeContext) {
	localctx = NewFuncReturnTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TinyLangParserRULE_funcReturnType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(105)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserSTRING, TinyLangParserNUM, TinyLangParserBOOL, TinyLangParserARRAY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.VariableType()
		}

	case TinyLangParserVOID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(TinyLangParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFuncInvocContext is an interface to support dynamic dispatch.
type IFuncInvocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncInvocContext differentiates from other interfaces.
	IsFuncInvocContext()
}

type FuncInvocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncInvocContext() *FuncInvocContext {
	var p = new(FuncInvocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcInvoc
	return p
}

func (*FuncInvocContext) IsFuncInvocContext() {}

func NewFuncInvocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncInvocContext {
	var p = new(FuncInvocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcInvoc

	return p
}

func (s *FuncInvocContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncInvocContext) FuncArgsInvoc() IFuncArgsInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncArgsInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncArgsInvocContext)
}

func (s *FuncInvocContext) SYS_FUNC() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSYS_FUNC, 0)
}

func (s *FuncInvocContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *FuncInvocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncInvocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncInvocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncInvoc(s)
	}
}

func (s *FuncInvocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncInvoc(s)
	}
}

func (s *FuncInvocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncInvoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncInvoc() (localctx IFuncInvocContext) {
	localctx = NewFuncInvocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TinyLangParserRULE_funcInvoc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserSYS_FUNC || _la == TinyLangParserITEM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(108)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(109)
		p.FuncArgsInvoc()
	}
	{
		p.SetState(110)
		p.Match(TinyLangParserT__2)
	}

	return localctx
}

// IFuncArgsInvocContext is an interface to support dynamic dispatch.
type IFuncArgsInvocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncArgsInvocContext differentiates from other interfaces.
	IsFuncArgsInvocContext()
}

type FuncArgsInvocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncArgsInvocContext() *FuncArgsInvocContext {
	var p = new(FuncArgsInvocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcArgsInvoc
	return p
}

func (*FuncArgsInvocContext) IsFuncArgsInvocContext() {}

func NewFuncArgsInvocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncArgsInvocContext {
	var p = new(FuncArgsInvocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcArgsInvoc

	return p
}

func (s *FuncArgsInvocContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncArgsInvocContext) AllFuncInvocArgs() []IFuncInvocArgsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncInvocArgsContext)(nil)).Elem())
	var tst = make([]IFuncInvocArgsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncInvocArgsContext)
		}
	}

	return tst
}

func (s *FuncArgsInvocContext) FuncInvocArgs(i int) IFuncInvocArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocArgsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocArgsContext)
}

func (s *FuncArgsInvocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncArgsInvocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncArgsInvocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncArgsInvoc(s)
	}
}

func (s *FuncArgsInvocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncArgsInvoc(s)
	}
}

func (s *FuncArgsInvocContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncArgsInvoc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncArgsInvoc() (localctx IFuncArgsInvocContext) {
	localctx = NewFuncArgsInvocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TinyLangParserRULE_funcArgsInvoc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-21)&-(0x1f+1)) == 0 && ((1<<uint((_la-21)))&((1<<(TinyLangParserSYS_FUNC-21))|(1<<(TinyLangParserNUMBER-21))|(1<<(TinyLangParserTRUE-21))|(1<<(TinyLangParserFALSE-21))|(1<<(TinyLangParserSTRING_RAW-21))|(1<<(TinyLangParserITEM-21)))) != 0 {
		{
			p.SetState(112)
			p.FuncInvocArgs()
		}
		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(113)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(114)
				p.FuncInvocArgs()
			}

			p.SetState(119)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IFuncInvocArgsContext is an interface to support dynamic dispatch.
type IFuncInvocArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncInvocArgsContext differentiates from other interfaces.
	IsFuncInvocArgsContext()
}

type FuncInvocArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncInvocArgsContext() *FuncInvocArgsContext {
	var p = new(FuncInvocArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_funcInvocArgs
	return p
}

func (*FuncInvocArgsContext) IsFuncInvocArgsContext() {}

func NewFuncInvocArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncInvocArgsContext {
	var p = new(FuncInvocArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_funcInvocArgs

	return p
}

func (s *FuncInvocArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncInvocArgsContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *FuncInvocArgsContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *FuncInvocArgsContext) ArrayElem() IArrayElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayElemContext)
}

func (s *FuncInvocArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncInvocArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncInvocArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterFuncInvocArgs(s)
	}
}

func (s *FuncInvocArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitFuncInvocArgs(s)
	}
}

func (s *FuncInvocArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitFuncInvocArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) FuncInvocArgs() (localctx IFuncInvocArgsContext) {
	localctx = NewFuncInvocArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TinyLangParserRULE_funcInvocArgs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Val()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.FuncInvoc()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.ArrayElem()
		}

	}

	return localctx
}

// IUpdVariableContext is an interface to support dynamic dispatch.
type IUpdVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdVariableContext differentiates from other interfaces.
	IsUpdVariableContext()
}

type UpdVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdVariableContext() *UpdVariableContext {
	var p = new(UpdVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_updVariable
	return p
}

func (*UpdVariableContext) IsUpdVariableContext() {}

func NewUpdVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdVariableContext {
	var p = new(UpdVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_updVariable

	return p
}

func (s *UpdVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdVariableContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *UpdVariableContext) AllArrayElem() []IArrayElemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrayElemContext)(nil)).Elem())
	var tst = make([]IArrayElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrayElemContext)
		}
	}

	return tst
}

func (s *UpdVariableContext) ArrayElem(i int) IArrayElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrayElemContext)
}

func (s *UpdVariableContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UpdVariableContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *UpdVariableContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *UpdVariableContext) ArrayInit() IArrayInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitContext)
}

func (s *UpdVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterUpdVariable(s)
	}
}

func (s *UpdVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitUpdVariable(s)
	}
}

func (s *UpdVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitUpdVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) UpdVariable() (localctx IUpdVariableContext) {
	localctx = NewUpdVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TinyLangParserRULE_updVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(127)
			p.Match(TinyLangParserITEM)
		}

	case 2:
		{
			p.SetState(128)
			p.ArrayElem()
		}

	}
	{
		p.SetState(131)
		p.Match(TinyLangParserT__6)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(132)
			p.Expr()
		}

	case 2:
		{
			p.SetState(133)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(134)
			p.Val()
		}

	case 4:
		{
			p.SetState(135)
			p.ArrayInit()
		}

	case 5:
		{
			p.SetState(136)
			p.ArrayElem()
		}

	}

	return localctx
}

// INewVariableContext is an interface to support dynamic dispatch.
type INewVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNewVariableContext differentiates from other interfaces.
	IsNewVariableContext()
}

type NewVariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNewVariableContext() *NewVariableContext {
	var p = new(NewVariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_newVariable
	return p
}

func (*NewVariableContext) IsNewVariableContext() {}

func NewNewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewVariableContext {
	var p = new(NewVariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_newVariable

	return p
}

func (s *NewVariableContext) GetParser() antlr.Parser { return s.parser }

func (s *NewVariableContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *NewVariableContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *NewVariableContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NewVariableContext) Val() IValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *NewVariableContext) ArrayInit() IArrayInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitContext)
}

func (s *NewVariableContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *NewVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewVariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NewVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterNewVariable(s)
	}
}

func (s *NewVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitNewVariable(s)
	}
}

func (s *NewVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitNewVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) NewVariable() (localctx INewVariableContext) {
	localctx = NewNewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TinyLangParserRULE_newVariable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(139)
		p.Match(TinyLangParserT__7)
	}
	{
		p.SetState(140)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(141)
		p.VariableType()
	}
	{
		p.SetState(142)
		p.Match(TinyLangParserT__6)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(143)
			p.Expr()
		}

	case 2:
		{
			p.SetState(144)
			p.Val()
		}

	case 3:
		{
			p.SetState(145)
			p.ArrayInit()
		}

	case 4:
		{
			p.SetState(146)
			p.FuncInvoc()
		}

	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserTRUE, 0)
}

func (s *ValContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserFALSE, 0)
}

func (s *ValContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *ValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
}

func (s *ValContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitVal(s)
	}
}

func (s *ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitVal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TinyLangParserRULE_val)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(TinyLangParserNUMBER-27))|(1<<(TinyLangParserTRUE-27))|(1<<(TinyLangParserFALSE-27))|(1<<(TinyLangParserSTRING_RAW-27))|(1<<(TinyLangParserITEM-27)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IVariableTypeContext is an interface to support dynamic dispatch.
type IVariableTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableTypeContext differentiates from other interfaces.
	IsVariableTypeContext()
}

type VariableTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableTypeContext() *VariableTypeContext {
	var p = new(VariableTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_variableType
	return p
}

func (*VariableTypeContext) IsVariableTypeContext() {}

func NewVariableTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableTypeContext {
	var p = new(VariableTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_variableType

	return p
}

func (s *VariableTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableTypeContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(TinyLangParserARRAY, 0)
}

func (s *VariableTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING, 0)
}

func (s *VariableTypeContext) NUM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUM, 0)
}

func (s *VariableTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL, 0)
}

func (s *VariableTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterVariableType(s)
	}
}

func (s *VariableTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitVariableType(s)
	}
}

func (s *VariableTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitVariableType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) VariableType() (localctx IVariableTypeContext) {
	localctx = NewVariableTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TinyLangParserRULE_variableType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(TinyLangParserSTRING-22))|(1<<(TinyLangParserNUM-22))|(1<<(TinyLangParserBOOL-22))|(1<<(TinyLangParserARRAY-22)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayElemContext is an interface to support dynamic dispatch.
type IArrayElemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayElemContext differentiates from other interfaces.
	IsArrayElemContext()
}

type ArrayElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayElemContext() *ArrayElemContext {
	var p = new(ArrayElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_arrayElem
	return p
}

func (*ArrayElemContext) IsArrayElemContext() {}

func NewArrayElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayElemContext {
	var p = new(ArrayElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_arrayElem

	return p
}

func (s *ArrayElemContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayElemContext) AllITEM() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserITEM)
}

func (s *ArrayElemContext) ITEM(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, i)
}

func (s *ArrayElemContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
}

func (s *ArrayElemContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *ArrayElemContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ArrayElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterArrayElem(s)
	}
}

func (s *ArrayElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitArrayElem(s)
	}
}

func (s *ArrayElemContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitArrayElem(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ArrayElem() (localctx IArrayElemContext) {
	localctx = NewArrayElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TinyLangParserRULE_arrayElem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(154)
		p.Match(TinyLangParserT__8)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(155)
			p.Match(TinyLangParserNUMBER)
		}

	case 2:
		{
			p.SetState(156)
			p.Match(TinyLangParserITEM)
		}

	case 3:
		{
			p.SetState(157)
			p.FuncInvoc()
		}

	case 4:
		{
			p.SetState(158)
			p.Expr()
		}

	}
	{
		p.SetState(161)
		p.Match(TinyLangParserT__9)
	}

	return localctx
}

// IArrayInitContext is an interface to support dynamic dispatch.
type IArrayInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitContext differentiates from other interfaces.
	IsArrayInitContext()
}

type ArrayInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitContext() *ArrayInitContext {
	var p = new(ArrayInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_arrayInit
	return p
}

func (*ArrayInitContext) IsArrayInitContext() {}

func NewArrayInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitContext {
	var p = new(ArrayInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_arrayInit

	return p
}

func (s *ArrayInitContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitContext) ArrayInitValue() IArrayInitValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitValueContext)
}

func (s *ArrayInitContext) ArrayInitEmpty() IArrayInitEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitEmptyContext)
}

func (s *ArrayInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterArrayInit(s)
	}
}

func (s *ArrayInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitArrayInit(s)
	}
}

func (s *ArrayInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitArrayInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ArrayInit() (localctx IArrayInitContext) {
	localctx = NewArrayInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TinyLangParserRULE_arrayInit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.ArrayInitValue()
		}

	case TinyLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.ArrayInitEmpty()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArrayInitEmptyContext is an interface to support dynamic dispatch.
type IArrayInitEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitEmptyContext differentiates from other interfaces.
	IsArrayInitEmptyContext()
}

type ArrayInitEmptyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitEmptyContext() *ArrayInitEmptyContext {
	var p = new(ArrayInitEmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_arrayInitEmpty
	return p
}

func (*ArrayInitEmptyContext) IsArrayInitEmptyContext() {}

func NewArrayInitEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitEmptyContext {
	var p = new(ArrayInitEmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_arrayInitEmpty

	return p
}

func (s *ArrayInitEmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitEmptyContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
}

func (s *ArrayInitEmptyContext) NUM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUM, 0)
}

func (s *ArrayInitEmptyContext) STRING() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING, 0)
}

func (s *ArrayInitEmptyContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL, 0)
}

func (s *ArrayInitEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitEmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterArrayInitEmpty(s)
	}
}

func (s *ArrayInitEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitArrayInitEmpty(s)
	}
}

func (s *ArrayInitEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitArrayInitEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ArrayInitEmpty() (localctx IArrayInitEmptyContext) {
	localctx = NewArrayInitEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TinyLangParserRULE_arrayInitEmpty)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(TinyLangParserT__8)
	}
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserT__10 || _la == TinyLangParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(169)
		p.Match(TinyLangParserT__9)
	}
	{
		p.SetState(170)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLangParserSTRING)|(1<<TinyLangParserNUM)|(1<<TinyLangParserBOOL))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrayInitValueContext is an interface to support dynamic dispatch.
type IArrayInitValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitValueContext differentiates from other interfaces.
	IsArrayInitValueContext()
}

type ArrayInitValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitValueContext() *ArrayInitValueContext {
	var p = new(ArrayInitValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_arrayInitValue
	return p
}

func (*ArrayInitValueContext) IsArrayInitValueContext() {}

func NewArrayInitValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitValueContext {
	var p = new(ArrayInitValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_arrayInitValue

	return p
}

func (s *ArrayInitValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitValueContext) ArrayInitElems() IArrayInitElemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayInitElemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayInitElemsContext)
}

func (s *ArrayInitValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterArrayInitValue(s)
	}
}

func (s *ArrayInitValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitArrayInitValue(s)
	}
}

func (s *ArrayInitValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitArrayInitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ArrayInitValue() (localctx IArrayInitValueContext) {
	localctx = NewArrayInitValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TinyLangParserRULE_arrayInitValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(173)
		p.ArrayInitElems()
	}
	{
		p.SetState(174)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}

// IArrayInitElemsContext is an interface to support dynamic dispatch.
type IArrayInitElemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayInitElemsContext differentiates from other interfaces.
	IsArrayInitElemsContext()
}

type ArrayInitElemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayInitElemsContext() *ArrayInitElemsContext {
	var p = new(ArrayInitElemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_arrayInitElems
	return p
}

func (*ArrayInitElemsContext) IsArrayInitElemsContext() {}

func NewArrayInitElemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayInitElemsContext {
	var p = new(ArrayInitElemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_arrayInitElems

	return p
}

func (s *ArrayInitElemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayInitElemsContext) AllBOOL_VAL() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserBOOL_VAL)
}

func (s *ArrayInitElemsContext) BOOL_VAL(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, i)
}

func (s *ArrayInitElemsContext) AllSTRING_RAW() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserSTRING_RAW)
}

func (s *ArrayInitElemsContext) STRING_RAW(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, i)
}

func (s *ArrayInitElemsContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserNUMBER)
}

func (s *ArrayInitElemsContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, i)
}

func (s *ArrayInitElemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayInitElemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayInitElemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterArrayInitElems(s)
	}
}

func (s *ArrayInitElemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitArrayInitElems(s)
	}
}

func (s *ArrayInitElemsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitArrayInitElems(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ArrayInitElems() (localctx IArrayInitElemsContext) {
	localctx = NewArrayInitElemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TinyLangParserRULE_arrayInitElems)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserBOOL_VAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Match(TinyLangParserBOOL_VAL)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(177)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(178)
				p.Match(TinyLangParserBOOL_VAL)
			}

			p.SetState(183)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserSTRING_RAW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(184)
			p.Match(TinyLangParserSTRING_RAW)
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(185)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(186)
				p.Match(TinyLangParserSTRING_RAW)
			}

			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(192)
			p.Match(TinyLangParserNUMBER)
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(193)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(194)
				p.Match(TinyLangParserNUMBER)
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) NUM_SIGN() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUM_SIGN, 0)
}

func (s *ExprContext) AllExprOperand() []IExprOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprOperandContext)(nil)).Elem())
	var tst = make([]IExprOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprOperandContext)
		}
	}

	return tst
}

func (s *ExprContext) ExprOperand(i int) IExprOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprOperandContext)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TinyLangParserRULE_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(220)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserSYS_FUNC, TinyLangParserNUMBER, TinyLangParserSTRING_RAW, TinyLangParserITEM:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.ExprOperand()
		}

		{
			p.SetState(203)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(204)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(205)
				p.Expr()
			}

		}

	case TinyLangParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(TinyLangParserT__1)
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(209)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(210)
				p.Expr()
			}

		}
		{
			p.SetState(213)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(214)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(215)
				p.Expr()
			}

		}
		{
			p.SetState(218)
			p.Match(TinyLangParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExprOperandContext is an interface to support dynamic dispatch.
type IExprOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprOperandContext differentiates from other interfaces.
	IsExprOperandContext()
}

type ExprOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprOperandContext() *ExprOperandContext {
	var p = new(ExprOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_exprOperand
	return p
}

func (*ExprOperandContext) IsExprOperandContext() {}

func NewExprOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprOperandContext {
	var p = new(ExprOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_exprOperand

	return p
}

func (s *ExprOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprOperandContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *ExprOperandContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
}

func (s *ExprOperandContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *ExprOperandContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *ExprOperandContext) ArrayElem() IArrayElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayElemContext)
}

func (s *ExprOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterExprOperand(s)
	}
}

func (s *ExprOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitExprOperand(s)
	}
}

func (s *ExprOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitExprOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ExprOperand() (localctx IExprOperandContext) {
	localctx = NewExprOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TinyLangParserRULE_exprOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)
			p.FuncInvoc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(223)
			p.Match(TinyLangParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(224)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(225)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(226)
			p.ArrayElem()
		}

	}

	return localctx
}

// IBoolExprContext is an interface to support dynamic dispatch.
type IBoolExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprContext differentiates from other interfaces.
	IsBoolExprContext()
}

type BoolExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprContext() *BoolExprContext {
	var p = new(BoolExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_boolExpr
	return p
}

func (*BoolExprContext) IsBoolExprContext() {}

func NewBoolExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprContext {
	var p = new(BoolExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_boolExpr

	return p
}

func (s *BoolExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprContext) BoolExprSingle() IBoolExprSingleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprSingleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprSingleContext)
}

func (s *BoolExprContext) AllBoolExprSingleExtra() []IBoolExprSingleExtraContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprSingleExtraContext)(nil)).Elem())
	var tst = make([]IBoolExprSingleExtraContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprSingleExtraContext)
		}
	}

	return tst
}

func (s *BoolExprContext) BoolExprSingleExtra(i int) IBoolExprSingleExtraContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprSingleExtraContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprSingleExtraContext)
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) BoolExpr() (localctx IBoolExprContext) {
	localctx = NewBoolExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TinyLangParserRULE_boolExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.BoolExprSingle()
	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(230)
				p.BoolExprSingleExtra()
			}

		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolExprSingleExtraContext is an interface to support dynamic dispatch.
type IBoolExprSingleExtraContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprSingleExtraContext differentiates from other interfaces.
	IsBoolExprSingleExtraContext()
}

type BoolExprSingleExtraContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprSingleExtraContext() *BoolExprSingleExtraContext {
	var p = new(BoolExprSingleExtraContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_boolExprSingleExtra
	return p
}

func (*BoolExprSingleExtraContext) IsBoolExprSingleExtraContext() {}

func NewBoolExprSingleExtraContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprSingleExtraContext {
	var p = new(BoolExprSingleExtraContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_boolExprSingleExtra

	return p
}

func (s *BoolExprSingleExtraContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprSingleExtraContext) BOOL_PL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_PL, 0)
}

func (s *BoolExprSingleExtraContext) BoolExprSingle() IBoolExprSingleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprSingleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprSingleContext)
}

func (s *BoolExprSingleExtraContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprSingleExtraContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExprSingleExtraContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterBoolExprSingleExtra(s)
	}
}

func (s *BoolExprSingleExtraContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitBoolExprSingleExtra(s)
	}
}

func (s *BoolExprSingleExtraContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitBoolExprSingleExtra(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) BoolExprSingleExtra() (localctx IBoolExprSingleExtraContext) {
	localctx = NewBoolExprSingleExtraContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TinyLangParserRULE_boolExprSingleExtra)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(TinyLangParserBOOL_PL)
	}
	{
		p.SetState(237)
		p.BoolExprSingle()
	}

	return localctx
}

// IBoolExprSingleContext is an interface to support dynamic dispatch.
type IBoolExprSingleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprSingleContext differentiates from other interfaces.
	IsBoolExprSingleContext()
}

type BoolExprSingleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprSingleContext() *BoolExprSingleContext {
	var p = new(BoolExprSingleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_boolExprSingle
	return p
}

func (*BoolExprSingleContext) IsBoolExprSingleContext() {}

func NewBoolExprSingleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprSingleContext {
	var p = new(BoolExprSingleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_boolExprSingle

	return p
}

func (s *BoolExprSingleContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprSingleContext) BOOL_SIGN() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_SIGN, 0)
}

func (s *BoolExprSingleContext) AllBoolExprOperand() []IBoolExprOperandContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprOperandContext)(nil)).Elem())
	var tst = make([]IBoolExprOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprOperandContext)
		}
	}

	return tst
}

func (s *BoolExprSingleContext) BoolExprOperand(i int) IBoolExprOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprOperandContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprOperandContext)
}

func (s *BoolExprSingleContext) AllBoolExpr() []IBoolExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprContext)(nil)).Elem())
	var tst = make([]IBoolExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprContext)
		}
	}

	return tst
}

func (s *BoolExprSingleContext) BoolExpr(i int) IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *BoolExprSingleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprSingleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExprSingleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterBoolExprSingle(s)
	}
}

func (s *BoolExprSingleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitBoolExprSingle(s)
	}
}

func (s *BoolExprSingleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitBoolExprSingle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) BoolExprSingle() (localctx IBoolExprSingleContext) {
	localctx = NewBoolExprSingleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TinyLangParserRULE_boolExprSingle)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(239)
			p.BoolExprOperand()
		}

		{
			p.SetState(240)
			p.Match(TinyLangParserBOOL_SIGN)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(241)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(242)
				p.BoolExpr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
			p.Match(TinyLangParserT__1)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(246)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(247)
				p.BoolExpr()
			}

		}
		{
			p.SetState(250)
			p.Match(TinyLangParserBOOL_SIGN)
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(251)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(252)
				p.BoolExpr()
			}

		}
		{
			p.SetState(255)
			p.Match(TinyLangParserT__2)
		}

	}

	return localctx
}

// IBoolExprOperandContext is an interface to support dynamic dispatch.
type IBoolExprOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolExprOperandContext differentiates from other interfaces.
	IsBoolExprOperandContext()
}

type BoolExprOperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolExprOperandContext() *BoolExprOperandContext {
	var p = new(BoolExprOperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_boolExprOperand
	return p
}

func (*BoolExprOperandContext) IsBoolExprOperandContext() {}

func NewBoolExprOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExprOperandContext {
	var p = new(BoolExprOperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_boolExprOperand

	return p
}

func (s *BoolExprOperandContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExprOperandContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BoolExprOperandContext) ExprOperand() IExprOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprOperandContext)
}

func (s *BoolExprOperandContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserTRUE, 0)
}

func (s *BoolExprOperandContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserFALSE, 0)
}

func (s *BoolExprOperandContext) ArrayElem() IArrayElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayElemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayElemContext)
}

func (s *BoolExprOperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprOperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExprOperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterBoolExprOperand(s)
	}
}

func (s *BoolExprOperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitBoolExprOperand(s)
	}
}

func (s *BoolExprOperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitBoolExprOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) BoolExprOperand() (localctx IBoolExprOperandContext) {
	localctx = NewBoolExprOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TinyLangParserRULE_boolExprOperand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.ExprOperand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Match(TinyLangParserTRUE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Match(TinyLangParserFALSE)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(263)
			p.ArrayElem()
		}

	}

	return localctx
}

// IBreakOrContinueContext is an interface to support dynamic dispatch.
type IBreakOrContinueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakOrContinueContext differentiates from other interfaces.
	IsBreakOrContinueContext()
}

type BreakOrContinueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakOrContinueContext() *BreakOrContinueContext {
	var p = new(BreakOrContinueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_breakOrContinue
	return p
}

func (*BreakOrContinueContext) IsBreakOrContinueContext() {}

func NewBreakOrContinueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakOrContinueContext {
	var p = new(BreakOrContinueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_breakOrContinue

	return p
}

func (s *BreakOrContinueContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakOrContinueContext) BREAK() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBREAK, 0)
}

func (s *BreakOrContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserCONTINUE, 0)
}

func (s *BreakOrContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakOrContinueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakOrContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterBreakOrContinue(s)
	}
}

func (s *BreakOrContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitBreakOrContinue(s)
	}
}

func (s *BreakOrContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitBreakOrContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) BreakOrContinue() (localctx IBreakOrContinueContext) {
	localctx = NewBreakOrContinueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TinyLangParserRULE_breakOrContinue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserBREAK || _la == TinyLangParserCONTINUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementBodyContext is an interface to support dynamic dispatch.
type IStatementBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementBodyContext differentiates from other interfaces.
	IsStatementBodyContext()
}

type StatementBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBodyContext() *StatementBodyContext {
	var p = new(StatementBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_statementBody
	return p
}

func (*StatementBodyContext) IsStatementBodyContext() {}

func NewStatementBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBodyContext {
	var p = new(StatementBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_statementBody

	return p
}

func (s *StatementBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBodyContext) AllNewVariable() []INewVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INewVariableContext)(nil)).Elem())
	var tst = make([]INewVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INewVariableContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) NewVariable(i int) INewVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INewVariableContext)
}

func (s *StatementBodyContext) AllUpdVariable() []IUpdVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdVariableContext)(nil)).Elem())
	var tst = make([]IUpdVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdVariableContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) UpdVariable(i int) IUpdVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdVariableContext)
}

func (s *StatementBodyContext) AllFuncInvoc() []IFuncInvocContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem())
	var tst = make([]IFuncInvocContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncInvocContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) FuncInvoc(i int) IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *StatementBodyContext) AllIfElseSt() []IIfElseStContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfElseStContext)(nil)).Elem())
	var tst = make([]IIfElseStContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfElseStContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) IfElseSt(i int) IIfElseStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfElseStContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfElseStContext)
}

func (s *StatementBodyContext) AllWhileSt() []IWhileStContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhileStContext)(nil)).Elem())
	var tst = make([]IWhileStContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhileStContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) WhileSt(i int) IWhileStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileStContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhileStContext)
}

func (s *StatementBodyContext) AllForSt() []IForStContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IForStContext)(nil)).Elem())
	var tst = make([]IForStContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IForStContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) ForSt(i int) IForStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IForStContext)
}

func (s *StatementBodyContext) AllBreakOrContinue() []IBreakOrContinueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBreakOrContinueContext)(nil)).Elem())
	var tst = make([]IBreakOrContinueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBreakOrContinueContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) BreakOrContinue(i int) IBreakOrContinueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakOrContinueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBreakOrContinueContext)
}

func (s *StatementBodyContext) AllFuncReturn() []IFuncReturnContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFuncReturnContext)(nil)).Elem())
	var tst = make([]IFuncReturnContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFuncReturnContext)
		}
	}

	return tst
}

func (s *StatementBodyContext) FuncReturn(i int) IFuncReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnContext)
}

func (s *StatementBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterStatementBody(s)
	}
}

func (s *StatementBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitStatementBody(s)
	}
}

func (s *StatementBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitStatementBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) StatementBody() (localctx IStatementBodyContext) {
	localctx = NewStatementBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TinyLangParserRULE_statementBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-8)&-(0x1f+1)) == 0 && ((1<<uint((_la-8)))&((1<<(TinyLangParserT__7-8))|(1<<(TinyLangParserWHILE-8))|(1<<(TinyLangParserRETURN-8))|(1<<(TinyLangParserBREAK-8))|(1<<(TinyLangParserCONTINUE-8))|(1<<(TinyLangParserIF-8))|(1<<(TinyLangParserFOR-8))|(1<<(TinyLangParserSYS_FUNC-8))|(1<<(TinyLangParserITEM-8)))) != 0 {
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(268)
				p.NewVariable()
			}

		case 2:
			{
				p.SetState(269)
				p.UpdVariable()
			}

		case 3:
			{
				p.SetState(270)
				p.FuncInvoc()
			}

		case 4:
			{
				p.SetState(271)
				p.IfElseSt()
			}

		case 5:
			{
				p.SetState(272)
				p.WhileSt()
			}

		case 6:
			{
				p.SetState(273)
				p.ForSt()
			}

		case 7:
			{
				p.SetState(274)
				p.BreakOrContinue()
			}

		case 8:
			{
				p.SetState(275)
				p.FuncReturn()
			}

		}

		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIfElseStContext is an interface to support dynamic dispatch.
type IIfElseStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfElseStContext differentiates from other interfaces.
	IsIfElseStContext()
}

type IfElseStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfElseStContext() *IfElseStContext {
	var p = new(IfElseStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_ifElseSt
	return p
}

func (*IfElseStContext) IsIfElseStContext() {}

func NewIfElseStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfElseStContext {
	var p = new(IfElseStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_ifElseSt

	return p
}

func (s *IfElseStContext) GetParser() antlr.Parser { return s.parser }

func (s *IfElseStContext) IfSt() IIfStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStContext)
}

func (s *IfElseStContext) AllElseIfSt() []IElseIfStContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStContext)(nil)).Elem())
	var tst = make([]IElseIfStContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStContext)
		}
	}

	return tst
}

func (s *IfElseStContext) ElseIfSt(i int) IElseIfStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStContext)
}

func (s *IfElseStContext) ElseSt() IElseStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStContext)
}

func (s *IfElseStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfElseStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterIfElseSt(s)
	}
}

func (s *IfElseStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitIfElseSt(s)
	}
}

func (s *IfElseStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitIfElseSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) IfElseSt() (localctx IIfElseStContext) {
	localctx = NewIfElseStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TinyLangParserRULE_ifElseSt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.IfSt()
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(282)
				p.ElseIfSt()
			}

		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserELSE {
		{
			p.SetState(288)
			p.ElseSt()
		}

	}

	return localctx
}

// IIfStContext is an interface to support dynamic dispatch.
type IIfStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStContext differentiates from other interfaces.
	IsIfStContext()
}

type IfStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStContext() *IfStContext {
	var p = new(IfStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_ifSt
	return p
}

func (*IfStContext) IsIfStContext() {}

func NewIfStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStContext {
	var p = new(IfStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_ifSt

	return p
}

func (s *IfStContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStContext) IF() antlr.TerminalNode {
	return s.GetToken(TinyLangParserIF, 0)
}

func (s *IfStContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserTRUE, 0)
}

func (s *IfStContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserFALSE, 0)
}

func (s *IfStContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *IfStContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *IfStContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *IfStContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *IfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterIfSt(s)
	}
}

func (s *IfStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitIfSt(s)
	}
}

func (s *IfStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitIfSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) IfSt() (localctx IIfStContext) {
	localctx = NewIfStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TinyLangParserRULE_ifSt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(TinyLangParserIF)
	}
	{
		p.SetState(292)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(293)
			p.Match(TinyLangParserTRUE)
		}

	case 2:
		{
			p.SetState(294)
			p.Match(TinyLangParserFALSE)
		}

	case 3:
		{
			p.SetState(295)
			p.BoolExpr()
		}

	case 4:
		{
			p.SetState(296)
			p.Match(TinyLangParserITEM)
		}

	case 5:
		{
			p.SetState(297)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(300)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(301)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(302)
			p.StatementBody()
		}

	}
	{
		p.SetState(305)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}

// IElseIfStContext is an interface to support dynamic dispatch.
type IElseIfStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStContext differentiates from other interfaces.
	IsElseIfStContext()
}

type ElseIfStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStContext() *ElseIfStContext {
	var p = new(ElseIfStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_elseIfSt
	return p
}

func (*ElseIfStContext) IsElseIfStContext() {}

func NewElseIfStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStContext {
	var p = new(ElseIfStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_elseIfSt

	return p
}

func (s *ElseIfStContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserELSE, 0)
}

func (s *ElseIfStContext) IfSt() IIfStContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStContext)
}

func (s *ElseIfStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterElseIfSt(s)
	}
}

func (s *ElseIfStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitElseIfSt(s)
	}
}

func (s *ElseIfStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitElseIfSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ElseIfSt() (localctx IElseIfStContext) {
	localctx = NewElseIfStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, TinyLangParserRULE_elseIfSt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(308)
		p.IfSt()
	}

	return localctx
}

// IElseStContext is an interface to support dynamic dispatch.
type IElseStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStContext differentiates from other interfaces.
	IsElseStContext()
}

type ElseStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStContext() *ElseStContext {
	var p = new(ElseStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_elseSt
	return p
}

func (*ElseStContext) IsElseStContext() {}

func NewElseStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStContext {
	var p = new(ElseStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_elseSt

	return p
}

func (s *ElseStContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserELSE, 0)
}

func (s *ElseStContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *ElseStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterElseSt(s)
	}
}

func (s *ElseStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitElseSt(s)
	}
}

func (s *ElseStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitElseSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ElseSt() (localctx IElseStContext) {
	localctx = NewElseStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TinyLangParserRULE_elseSt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(311)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(312)
			p.StatementBody()
		}

	}
	{
		p.SetState(315)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}

// IWhileStContext is an interface to support dynamic dispatch.
type IWhileStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileStContext differentiates from other interfaces.
	IsWhileStContext()
}

type WhileStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStContext() *WhileStContext {
	var p = new(WhileStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_whileSt
	return p
}

func (*WhileStContext) IsWhileStContext() {}

func NewWhileStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStContext {
	var p = new(WhileStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_whileSt

	return p
}

func (s *WhileStContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStContext) WHILE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserWHILE, 0)
}

func (s *WhileStContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserTRUE, 0)
}

func (s *WhileStContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TinyLangParserFALSE, 0)
}

func (s *WhileStContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *WhileStContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *WhileStContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *WhileStContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *WhileStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterWhileSt(s)
	}
}

func (s *WhileStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitWhileSt(s)
	}
}

func (s *WhileStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitWhileSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) WhileSt() (localctx IWhileStContext) {
	localctx = NewWhileStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TinyLangParserRULE_whileSt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(TinyLangParserWHILE)
	}
	{
		p.SetState(318)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(319)
			p.Match(TinyLangParserTRUE)
		}

	case 2:
		{
			p.SetState(320)
			p.Match(TinyLangParserFALSE)
		}

	case 3:
		{
			p.SetState(321)
			p.BoolExpr()
		}

	case 4:
		{
			p.SetState(322)
			p.Match(TinyLangParserITEM)
		}

	case 5:
		{
			p.SetState(323)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(326)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(327)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(328)
			p.StatementBody()
		}

	}
	{
		p.SetState(331)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}

// IForStContext is an interface to support dynamic dispatch.
type IForStContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStContext differentiates from other interfaces.
	IsForStContext()
}

type ForStContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStContext() *ForStContext {
	var p = new(ForStContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinyLangParserRULE_forSt
	return p
}

func (*ForStContext) IsForStContext() {}

func NewForStContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStContext {
	var p = new(ForStContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinyLangParserRULE_forSt

	return p
}

func (s *ForStContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStContext) FOR() antlr.TerminalNode {
	return s.GetToken(TinyLangParserFOR, 0)
}

func (s *ForStContext) NewVariable() INewVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INewVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewVariableContext)
}

func (s *ForStContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *ForStContext) UpdVariable() IUpdVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdVariableContext)
}

func (s *ForStContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *ForStContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.EnterForSt(s)
	}
}

func (s *ForStContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinyLangListener); ok {
		listenerT.ExitForSt(s)
	}
}

func (s *ForStContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case TinyLangVisitor:
		return t.VisitForSt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *TinyLangParser) ForSt() (localctx IForStContext) {
	localctx = NewForStContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TinyLangParserRULE_forSt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(TinyLangParserFOR)
	}
	{
		p.SetState(334)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(335)
		p.NewVariable()
	}
	{
		p.SetState(336)
		p.Match(TinyLangParserT__11)
	}
	{
		p.SetState(337)
		p.BoolExpr()
	}
	{
		p.SetState(338)
		p.Match(TinyLangParserT__11)
	}
	{
		p.SetState(339)
		p.UpdVariable()
	}
	{
		p.SetState(340)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(341)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(342)
		p.StatementBody()
	}
	{
		p.SetState(343)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}
