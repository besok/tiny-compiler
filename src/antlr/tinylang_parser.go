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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 361,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 7, 2, 62, 10, 2, 12, 2, 14, 2, 65, 11, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	78, 10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 88, 10,
	5, 12, 5, 14, 5, 91, 11, 5, 5, 5, 93, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 104, 10, 6, 3, 7, 3, 7, 5, 7, 108, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 118, 10, 9, 12,
	9, 14, 9, 121, 11, 9, 5, 9, 123, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 131, 10, 10, 3, 11, 3, 11, 5, 11, 135, 10, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 146, 10,
	11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 159, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 5, 14, 169, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 5, 15, 175,
	10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 18, 7, 18, 189, 10, 18, 12, 18, 14, 18, 192, 11, 18, 3,
	18, 3, 18, 3, 18, 7, 18, 197, 10, 18, 12, 18, 14, 18, 200, 11, 18, 3, 18,
	3, 18, 3, 18, 7, 18, 205, 10, 18, 12, 18, 14, 18, 208, 11, 18, 5, 18, 210,
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 216, 10, 19, 3, 19, 3, 19, 3,
	19, 5, 19, 221, 10, 19, 3, 19, 3, 19, 3, 19, 5, 19, 226, 10, 19, 3, 19,
	3, 19, 5, 19, 230, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 237,
	10, 20, 3, 21, 3, 21, 3, 21, 7, 21, 242, 10, 21, 12, 21, 14, 21, 245, 11,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 251, 10, 22, 3, 22, 3, 22, 3, 22,
	5, 22, 256, 10, 22, 3, 22, 3, 22, 3, 22, 5, 22, 261, 10, 22, 3, 22, 3,
	22, 5, 22, 265, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 271, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 281, 10,
	24, 12, 24, 14, 24, 284, 11, 24, 3, 25, 3, 25, 7, 25, 288, 10, 25, 12,
	25, 14, 25, 291, 11, 25, 3, 25, 5, 25, 294, 10, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 5, 26, 302, 10, 26, 3, 26, 3, 26, 3, 26, 5, 26, 307,
	10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	5, 27, 318, 10, 27, 3, 27, 3, 27, 3, 27, 5, 27, 323, 10, 27, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 28, 5, 28, 330, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 340, 10, 29, 3, 29, 3, 29, 3, 29, 5,
	29, 345, 10, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 2, 2, 31, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 2, 7, 4, 2, 25, 25, 31, 31, 5, 2, 26, 26, 28,
	29, 39, 39, 4, 2, 13, 13, 35, 35, 4, 2, 26, 26, 28, 29, 3, 2, 14, 15, 2,
	412, 2, 63, 3, 2, 2, 2, 4, 68, 3, 2, 2, 2, 6, 81, 3, 2, 2, 2, 8, 92, 3,
	2, 2, 2, 10, 94, 3, 2, 2, 2, 12, 107, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2,
	16, 122, 3, 2, 2, 2, 18, 130, 3, 2, 2, 2, 20, 134, 3, 2, 2, 2, 22, 147,
	3, 2, 2, 2, 24, 160, 3, 2, 2, 2, 26, 162, 3, 2, 2, 2, 28, 174, 3, 2, 2,
	2, 30, 176, 3, 2, 2, 2, 32, 181, 3, 2, 2, 2, 34, 209, 3, 2, 2, 2, 36, 229,
	3, 2, 2, 2, 38, 236, 3, 2, 2, 2, 40, 238, 3, 2, 2, 2, 42, 264, 3, 2, 2,
	2, 44, 270, 3, 2, 2, 2, 46, 282, 3, 2, 2, 2, 48, 285, 3, 2, 2, 2, 50, 295,
	3, 2, 2, 2, 52, 310, 3, 2, 2, 2, 54, 326, 3, 2, 2, 2, 56, 333, 3, 2, 2,
	2, 58, 348, 3, 2, 2, 2, 60, 62, 5, 4, 3, 2, 61, 60, 3, 2, 2, 2, 62, 65,
	3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66, 3, 2, 2, 2,
	65, 63, 3, 2, 2, 2, 66, 67, 7, 2, 2, 3, 67, 3, 3, 2, 2, 2, 68, 69, 7, 3,
	2, 2, 69, 70, 7, 31, 2, 2, 70, 71, 7, 4, 2, 2, 71, 72, 5, 8, 5, 2, 72,
	73, 7, 5, 2, 2, 73, 74, 5, 12, 7, 2, 74, 75, 7, 6, 2, 2, 75, 77, 5, 46,
	24, 2, 76, 78, 5, 10, 6, 2, 77, 76, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	79, 3, 2, 2, 2, 79, 80, 7, 7, 2, 2, 80, 5, 3, 2, 2, 2, 81, 82, 7, 31, 2,
	2, 82, 83, 5, 24, 13, 2, 83, 7, 3, 2, 2, 2, 84, 89, 5, 6, 4, 2, 85, 86,
	7, 8, 2, 2, 86, 88, 5, 6, 4, 2, 87, 85, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2,
	89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3,
	2, 2, 2, 92, 84, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 9, 3, 2, 2, 2, 94,
	103, 7, 19, 2, 2, 95, 104, 5, 26, 14, 2, 96, 104, 5, 14, 8, 2, 97, 104,
	7, 31, 2, 2, 98, 104, 7, 36, 2, 2, 99, 104, 7, 34, 2, 2, 100, 104, 7, 35,
	2, 2, 101, 104, 5, 36, 19, 2, 102, 104, 5, 40, 21, 2, 103, 95, 3, 2, 2,
	2, 103, 96, 3, 2, 2, 2, 103, 97, 3, 2, 2, 2, 103, 98, 3, 2, 2, 2, 103,
	99, 3, 2, 2, 2, 103, 100, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 102, 3,
	2, 2, 2, 104, 11, 3, 2, 2, 2, 105, 108, 5, 24, 13, 2, 106, 108, 7, 27,
	2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 13, 3, 2, 2, 2,
	109, 110, 9, 2, 2, 2, 110, 111, 7, 4, 2, 2, 111, 112, 5, 16, 9, 2, 112,
	113, 7, 5, 2, 2, 113, 15, 3, 2, 2, 2, 114, 119, 5, 18, 10, 2, 115, 116,
	7, 8, 2, 2, 116, 118, 5, 18, 10, 2, 117, 115, 3, 2, 2, 2, 118, 121, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 123, 3, 2, 2,
	2, 121, 119, 3, 2, 2, 2, 122, 114, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123,
	17, 3, 2, 2, 2, 124, 131, 7, 35, 2, 2, 125, 131, 7, 36, 2, 2, 126, 131,
	7, 34, 2, 2, 127, 131, 7, 31, 2, 2, 128, 131, 5, 14, 8, 2, 129, 131, 5,
	26, 14, 2, 130, 124, 3, 2, 2, 2, 130, 125, 3, 2, 2, 2, 130, 126, 3, 2,
	2, 2, 130, 127, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 129, 3, 2, 2, 2,
	131, 19, 3, 2, 2, 2, 132, 135, 7, 31, 2, 2, 133, 135, 5, 26, 14, 2, 134,
	132, 3, 2, 2, 2, 134, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136, 145,
	7, 9, 2, 2, 137, 146, 5, 36, 19, 2, 138, 146, 5, 14, 8, 2, 139, 146, 7,
	34, 2, 2, 140, 146, 7, 36, 2, 2, 141, 146, 7, 35, 2, 2, 142, 146, 7, 31,
	2, 2, 143, 146, 5, 28, 15, 2, 144, 146, 5, 26, 14, 2, 145, 137, 3, 2, 2,
	2, 145, 138, 3, 2, 2, 2, 145, 139, 3, 2, 2, 2, 145, 140, 3, 2, 2, 2, 145,
	141, 3, 2, 2, 2, 145, 142, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 144,
	3, 2, 2, 2, 146, 21, 3, 2, 2, 2, 147, 148, 7, 10, 2, 2, 148, 149, 7, 31,
	2, 2, 149, 150, 5, 24, 13, 2, 150, 158, 7, 9, 2, 2, 151, 159, 5, 36, 19,
	2, 152, 159, 7, 34, 2, 2, 153, 159, 7, 36, 2, 2, 154, 159, 7, 35, 2, 2,
	155, 159, 7, 31, 2, 2, 156, 159, 5, 28, 15, 2, 157, 159, 5, 14, 8, 2, 158,
	151, 3, 2, 2, 2, 158, 152, 3, 2, 2, 2, 158, 153, 3, 2, 2, 2, 158, 154,
	3, 2, 2, 2, 158, 155, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 158, 157, 3, 2,
	2, 2, 159, 23, 3, 2, 2, 2, 160, 161, 9, 3, 2, 2, 161, 25, 3, 2, 2, 2, 162,
	163, 7, 31, 2, 2, 163, 168, 7, 11, 2, 2, 164, 169, 7, 35, 2, 2, 165, 169,
	7, 31, 2, 2, 166, 169, 5, 14, 8, 2, 167, 169, 5, 36, 19, 2, 168, 164, 3,
	2, 2, 2, 168, 165, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168, 167, 3, 2, 2,
	2, 169, 170, 3, 2, 2, 2, 170, 171, 7, 12, 2, 2, 171, 27, 3, 2, 2, 2, 172,
	175, 5, 32, 17, 2, 173, 175, 5, 30, 16, 2, 174, 172, 3, 2, 2, 2, 174, 173,
	3, 2, 2, 2, 175, 29, 3, 2, 2, 2, 176, 177, 7, 11, 2, 2, 177, 178, 9, 4,
	2, 2, 178, 179, 7, 12, 2, 2, 179, 180, 9, 5, 2, 2, 180, 31, 3, 2, 2, 2,
	181, 182, 7, 6, 2, 2, 182, 183, 5, 34, 18, 2, 183, 184, 7, 7, 2, 2, 184,
	33, 3, 2, 2, 2, 185, 190, 7, 34, 2, 2, 186, 187, 7, 8, 2, 2, 187, 189,
	7, 34, 2, 2, 188, 186, 3, 2, 2, 2, 189, 192, 3, 2, 2, 2, 190, 188, 3, 2,
	2, 2, 190, 191, 3, 2, 2, 2, 191, 210, 3, 2, 2, 2, 192, 190, 3, 2, 2, 2,
	193, 198, 7, 36, 2, 2, 194, 195, 7, 8, 2, 2, 195, 197, 7, 36, 2, 2, 196,
	194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199,
	3, 2, 2, 2, 199, 210, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 201, 206, 7, 35,
	2, 2, 202, 203, 7, 8, 2, 2, 203, 205, 7, 35, 2, 2, 204, 202, 3, 2, 2, 2,
	205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207,
	210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 209, 185, 3, 2, 2, 2, 209, 193,
	3, 2, 2, 2, 209, 201, 3, 2, 2, 2, 210, 35, 3, 2, 2, 2, 211, 212, 5, 38,
	20, 2, 212, 215, 7, 40, 2, 2, 213, 216, 5, 38, 20, 2, 214, 216, 5, 36,
	19, 2, 215, 213, 3, 2, 2, 2, 215, 214, 3, 2, 2, 2, 216, 230, 3, 2, 2, 2,
	217, 220, 7, 4, 2, 2, 218, 221, 5, 38, 20, 2, 219, 221, 5, 36, 19, 2, 220,
	218, 3, 2, 2, 2, 220, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 225,
	7, 40, 2, 2, 223, 226, 5, 38, 20, 2, 224, 226, 5, 36, 19, 2, 225, 223,
	3, 2, 2, 2, 225, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 7, 5,
	2, 2, 228, 230, 3, 2, 2, 2, 229, 211, 3, 2, 2, 2, 229, 217, 3, 2, 2, 2,
	230, 37, 3, 2, 2, 2, 231, 237, 5, 14, 8, 2, 232, 237, 7, 35, 2, 2, 233,
	237, 7, 31, 2, 2, 234, 237, 7, 36, 2, 2, 235, 237, 5, 26, 14, 2, 236, 231,
	3, 2, 2, 2, 236, 232, 3, 2, 2, 2, 236, 233, 3, 2, 2, 2, 236, 234, 3, 2,
	2, 2, 236, 235, 3, 2, 2, 2, 237, 39, 3, 2, 2, 2, 238, 243, 5, 42, 22, 2,
	239, 240, 9, 6, 2, 2, 240, 242, 5, 42, 22, 2, 241, 239, 3, 2, 2, 2, 242,
	245, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 244, 3, 2, 2, 2, 244, 41, 3,
	2, 2, 2, 245, 243, 3, 2, 2, 2, 246, 247, 5, 44, 23, 2, 247, 250, 7, 41,
	2, 2, 248, 251, 5, 44, 23, 2, 249, 251, 5, 40, 21, 2, 250, 248, 3, 2, 2,
	2, 250, 249, 3, 2, 2, 2, 251, 265, 3, 2, 2, 2, 252, 255, 7, 4, 2, 2, 253,
	256, 5, 44, 23, 2, 254, 256, 5, 40, 21, 2, 255, 253, 3, 2, 2, 2, 255, 254,
	3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 260, 7, 41, 2, 2, 258, 261, 5, 44,
	23, 2, 259, 261, 5, 40, 21, 2, 260, 258, 3, 2, 2, 2, 260, 259, 3, 2, 2,
	2, 261, 262, 3, 2, 2, 2, 262, 263, 7, 5, 2, 2, 263, 265, 3, 2, 2, 2, 264,
	246, 3, 2, 2, 2, 264, 252, 3, 2, 2, 2, 265, 43, 3, 2, 2, 2, 266, 271, 5,
	36, 19, 2, 267, 271, 5, 38, 20, 2, 268, 271, 7, 34, 2, 2, 269, 271, 5,
	26, 14, 2, 270, 266, 3, 2, 2, 2, 270, 267, 3, 2, 2, 2, 270, 268, 3, 2,
	2, 2, 270, 269, 3, 2, 2, 2, 271, 45, 3, 2, 2, 2, 272, 281, 5, 22, 12, 2,
	273, 281, 5, 20, 11, 2, 274, 281, 5, 14, 8, 2, 275, 281, 5, 48, 25, 2,
	276, 281, 5, 56, 29, 2, 277, 281, 5, 58, 30, 2, 278, 281, 7, 20, 2, 2,
	279, 281, 7, 21, 2, 2, 280, 272, 3, 2, 2, 2, 280, 273, 3, 2, 2, 2, 280,
	274, 3, 2, 2, 2, 280, 275, 3, 2, 2, 2, 280, 276, 3, 2, 2, 2, 280, 277,
	3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 280, 279, 3, 2, 2, 2, 281, 284, 3, 2,
	2, 2, 282, 280, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 47, 3, 2, 2, 2,
	284, 282, 3, 2, 2, 2, 285, 289, 5, 50, 26, 2, 286, 288, 5, 52, 27, 2, 287,
	286, 3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290,
	3, 2, 2, 2, 290, 293, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 294, 5, 54,
	28, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294, 49, 3, 2, 2, 2,
	295, 296, 7, 22, 2, 2, 296, 301, 7, 4, 2, 2, 297, 302, 7, 34, 2, 2, 298,
	302, 5, 40, 21, 2, 299, 302, 7, 31, 2, 2, 300, 302, 5, 14, 8, 2, 301, 297,
	3, 2, 2, 2, 301, 298, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 300, 3, 2,
	2, 2, 302, 303, 3, 2, 2, 2, 303, 304, 7, 5, 2, 2, 304, 306, 7, 6, 2, 2,
	305, 307, 5, 46, 24, 2, 306, 305, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307,
	308, 3, 2, 2, 2, 308, 309, 7, 7, 2, 2, 309, 51, 3, 2, 2, 2, 310, 311, 7,
	24, 2, 2, 311, 312, 7, 22, 2, 2, 312, 317, 7, 4, 2, 2, 313, 318, 7, 34,
	2, 2, 314, 318, 5, 40, 21, 2, 315, 318, 7, 31, 2, 2, 316, 318, 5, 14, 8,
	2, 317, 313, 3, 2, 2, 2, 317, 314, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317,
	316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 320, 7, 5, 2, 2, 320, 322,
	7, 6, 2, 2, 321, 323, 5, 46, 24, 2, 322, 321, 3, 2, 2, 2, 322, 323, 3,
	2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 325, 7, 7, 2, 2, 325, 53, 3, 2, 2,
	2, 326, 327, 7, 24, 2, 2, 327, 329, 7, 6, 2, 2, 328, 330, 5, 46, 24, 2,
	329, 328, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331,
	332, 7, 7, 2, 2, 332, 55, 3, 2, 2, 2, 333, 334, 7, 18, 2, 2, 334, 339,
	7, 4, 2, 2, 335, 340, 7, 34, 2, 2, 336, 340, 5, 40, 21, 2, 337, 340, 7,
	31, 2, 2, 338, 340, 5, 14, 8, 2, 339, 335, 3, 2, 2, 2, 339, 336, 3, 2,
	2, 2, 339, 337, 3, 2, 2, 2, 339, 338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2,
	341, 342, 7, 5, 2, 2, 342, 344, 7, 6, 2, 2, 343, 345, 5, 46, 24, 2, 344,
	343, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 347,
	7, 7, 2, 2, 347, 57, 3, 2, 2, 2, 348, 349, 7, 23, 2, 2, 349, 350, 7, 4,
	2, 2, 350, 351, 5, 20, 11, 2, 351, 352, 7, 16, 2, 2, 352, 353, 5, 40, 21,
	2, 353, 354, 7, 16, 2, 2, 354, 355, 5, 20, 11, 2, 355, 356, 7, 5, 2, 2,
	356, 357, 7, 6, 2, 2, 357, 358, 5, 46, 24, 2, 358, 359, 7, 7, 2, 2, 359,
	59, 3, 2, 2, 2, 42, 63, 77, 89, 92, 103, 107, 119, 122, 130, 134, 145,
	158, 168, 174, 190, 198, 206, 209, 215, 220, 225, 229, 236, 243, 250, 255,
	260, 264, 270, 280, 282, 289, 293, 301, 306, 317, 322, 329, 339, 344,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'func'", "'('", "')'", "'{'", "'}'", "','", "'='", "'var'", "'['",
	"']'", "'..'", "'&&'", "'||'", "';'", "", "'while'", "'return'", "'break'",
	"'continue'", "'if'", "'for'", "'else'", "", "'string'", "'void'", "'num'",
	"'bool'", "'null'", "", "'true'", "'false'", "", "", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "WHILE",
	"RETURN", "BREAK", "CONTINUE", "IF", "FOR", "ELSE", "SYS_FUNC", "STRING",
	"VOID", "NUM", "BOOL", "NULL", "ITEM", "TRUE", "FALSE", "BOOL_VAL", "NUMBER",
	"STRING_RAW", "SQ", "COMMENT", "ARRAY", "NUM_SIGN", "BOOL_SIGN",
}

var ruleNames = []string{
	"file", "funcInit", "funcArg", "funcArgs", "funcReturn", "funcReturnType",
	"funcInvoc", "funcArgsInvoc", "funcInvocArgs", "updVariable", "newVariable",
	"variableType", "arrayElem", "arrayInit", "arrayInitEmpty", "arrayInitValue",
	"arrayInitElems", "expr", "exprOperand", "boolExpr", "boolExprSingle",
	"boolExprOperand", "statementBody", "ifElseSt", "ifSt", "elseIfSt", "elseSt",
	"whileSt", "forSt",
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
	TinyLangParserT__12      = 13
	TinyLangParserT__13      = 14
	TinyLangParserWS         = 15
	TinyLangParserWHILE      = 16
	TinyLangParserRETURN     = 17
	TinyLangParserBREAK      = 18
	TinyLangParserCONTINUE   = 19
	TinyLangParserIF         = 20
	TinyLangParserFOR        = 21
	TinyLangParserELSE       = 22
	TinyLangParserSYS_FUNC   = 23
	TinyLangParserSTRING     = 24
	TinyLangParserVOID       = 25
	TinyLangParserNUM        = 26
	TinyLangParserBOOL       = 27
	TinyLangParserNULL       = 28
	TinyLangParserITEM       = 29
	TinyLangParserTRUE       = 30
	TinyLangParserFALSE      = 31
	TinyLangParserBOOL_VAL   = 32
	TinyLangParserNUMBER     = 33
	TinyLangParserSTRING_RAW = 34
	TinyLangParserSQ         = 35
	TinyLangParserCOMMENT    = 36
	TinyLangParserARRAY      = 37
	TinyLangParserNUM_SIGN   = 38
	TinyLangParserBOOL_SIGN  = 39
)

// TinyLangParser rules.
const (
	TinyLangParserRULE_file            = 0
	TinyLangParserRULE_funcInit        = 1
	TinyLangParserRULE_funcArg         = 2
	TinyLangParserRULE_funcArgs        = 3
	TinyLangParserRULE_funcReturn      = 4
	TinyLangParserRULE_funcReturnType  = 5
	TinyLangParserRULE_funcInvoc       = 6
	TinyLangParserRULE_funcArgsInvoc   = 7
	TinyLangParserRULE_funcInvocArgs   = 8
	TinyLangParserRULE_updVariable     = 9
	TinyLangParserRULE_newVariable     = 10
	TinyLangParserRULE_variableType    = 11
	TinyLangParserRULE_arrayElem       = 12
	TinyLangParserRULE_arrayInit       = 13
	TinyLangParserRULE_arrayInitEmpty  = 14
	TinyLangParserRULE_arrayInitValue  = 15
	TinyLangParserRULE_arrayInitElems  = 16
	TinyLangParserRULE_expr            = 17
	TinyLangParserRULE_exprOperand     = 18
	TinyLangParserRULE_boolExpr        = 19
	TinyLangParserRULE_boolExprSingle  = 20
	TinyLangParserRULE_boolExprOperand = 21
	TinyLangParserRULE_statementBody   = 22
	TinyLangParserRULE_ifElseSt        = 23
	TinyLangParserRULE_ifSt            = 24
	TinyLangParserRULE_elseIfSt        = 25
	TinyLangParserRULE_elseSt          = 26
	TinyLangParserRULE_whileSt         = 27
	TinyLangParserRULE_forSt           = 28
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
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLangParserT__0 {
		{
			p.SetState(58)
			p.FuncInit()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
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

func (s *FuncInitContext) FuncReturn() IFuncReturnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncReturnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncReturnContext)
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
		p.SetState(66)
		p.Match(TinyLangParserT__0)
	}
	{
		p.SetState(67)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(68)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(69)
		p.FuncArgs()
	}
	{
		p.SetState(70)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(71)
		p.FuncReturnType()
	}
	{
		p.SetState(72)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(73)
		p.StatementBody()
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserRETURN {
		{
			p.SetState(74)
			p.FuncReturn()
		}

	}
	{
		p.SetState(77)
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
		p.SetState(79)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(80)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserITEM {
		{
			p.SetState(82)
			p.FuncArg()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(83)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(84)
				p.FuncArg()
			}

			p.SetState(89)
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

func (s *FuncReturnContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *FuncReturnContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *FuncReturnContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
}

func (s *FuncReturnContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
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
		p.SetState(92)
		p.Match(TinyLangParserRETURN)
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(93)
			p.ArrayElem()
		}

	case 2:
		{
			p.SetState(94)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(95)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(96)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		{
			p.SetState(97)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 6:
		{
			p.SetState(98)
			p.Match(TinyLangParserNUMBER)
		}

	case 7:
		{
			p.SetState(99)
			p.Expr()
		}

	case 8:
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

	if ((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(TinyLangParserSYS_FUNC-23))|(1<<(TinyLangParserITEM-23))|(1<<(TinyLangParserBOOL_VAL-23))|(1<<(TinyLangParserNUMBER-23))|(1<<(TinyLangParserSTRING_RAW-23)))) != 0 {
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

func (s *FuncInvocArgsContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
}

func (s *FuncInvocArgsContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *FuncInvocArgsContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
}

func (s *FuncInvocArgsContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
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

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(TinyLangParserNUMBER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(123)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(124)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(125)
			p.Match(TinyLangParserITEM)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(126)
			p.FuncInvoc()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(127)
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

func (s *UpdVariableContext) AllITEM() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserITEM)
}

func (s *UpdVariableContext) ITEM(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, i)
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

func (s *UpdVariableContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
}

func (s *UpdVariableContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *UpdVariableContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(130)
			p.Match(TinyLangParserITEM)
		}

	case 2:
		{
			p.SetState(131)
			p.ArrayElem()
		}

	}
	{
		p.SetState(134)
		p.Match(TinyLangParserT__6)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(135)
			p.Expr()
		}

	case 2:
		{
			p.SetState(136)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(137)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		{
			p.SetState(138)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		{
			p.SetState(139)
			p.Match(TinyLangParserNUMBER)
		}

	case 6:
		{
			p.SetState(140)
			p.Match(TinyLangParserITEM)
		}

	case 7:
		{
			p.SetState(141)
			p.ArrayInit()
		}

	case 8:
		{
			p.SetState(142)
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

func (s *NewVariableContext) AllITEM() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserITEM)
}

func (s *NewVariableContext) ITEM(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, i)
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

func (s *NewVariableContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
}

func (s *NewVariableContext) STRING_RAW() antlr.TerminalNode {
	return s.GetToken(TinyLangParserSTRING_RAW, 0)
}

func (s *NewVariableContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinyLangParserNUMBER, 0)
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
		p.SetState(145)
		p.Match(TinyLangParserT__7)
	}
	{
		p.SetState(146)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(147)
		p.VariableType()
	}
	{
		p.SetState(148)
		p.Match(TinyLangParserT__6)
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(149)
			p.Expr()
		}

	case 2:
		{
			p.SetState(150)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 3:
		{
			p.SetState(151)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 4:
		{
			p.SetState(152)
			p.Match(TinyLangParserNUMBER)
		}

	case 5:
		{
			p.SetState(153)
			p.Match(TinyLangParserITEM)
		}

	case 6:
		{
			p.SetState(154)
			p.ArrayInit()
		}

	case 7:
		{
			p.SetState(155)
			p.FuncInvoc()
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
	p.EnterRule(localctx, 22, TinyLangParserRULE_variableType)
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
		p.SetState(158)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(TinyLangParserSTRING-24))|(1<<(TinyLangParserNUM-24))|(1<<(TinyLangParserBOOL-24))|(1<<(TinyLangParserARRAY-24)))) != 0) {
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
	p.EnterRule(localctx, 24, TinyLangParserRULE_arrayElem)

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
		p.SetState(160)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(161)
		p.Match(TinyLangParserT__8)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(162)
			p.Match(TinyLangParserNUMBER)
		}

	case 2:
		{
			p.SetState(163)
			p.Match(TinyLangParserITEM)
		}

	case 3:
		{
			p.SetState(164)
			p.FuncInvoc()
		}

	case 4:
		{
			p.SetState(165)
			p.Expr()
		}

	}
	{
		p.SetState(168)
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
	p.EnterRule(localctx, 26, TinyLangParserRULE_arrayInit)

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

	p.SetState(172)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(170)
			p.ArrayInitValue()
		}

	case TinyLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(171)
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
	p.EnterRule(localctx, 28, TinyLangParserRULE_arrayInitEmpty)
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
		p.SetState(174)
		p.Match(TinyLangParserT__8)
	}
	{
		p.SetState(175)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserT__10 || _la == TinyLangParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(176)
		p.Match(TinyLangParserT__9)
	}
	{
		p.SetState(177)
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
	p.EnterRule(localctx, 30, TinyLangParserRULE_arrayInitValue)

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
		p.SetState(179)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(180)
		p.ArrayInitElems()
	}
	{
		p.SetState(181)
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
	p.EnterRule(localctx, 32, TinyLangParserRULE_arrayInitElems)
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

	p.SetState(207)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserBOOL_VAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(183)
			p.Match(TinyLangParserBOOL_VAL)
		}
		p.SetState(188)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(184)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(185)
				p.Match(TinyLangParserBOOL_VAL)
			}

			p.SetState(190)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserSTRING_RAW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Match(TinyLangParserSTRING_RAW)
		}
		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(192)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(193)
				p.Match(TinyLangParserSTRING_RAW)
			}

			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)
			p.Match(TinyLangParserNUMBER)
		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(200)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(201)
				p.Match(TinyLangParserNUMBER)
			}

			p.SetState(206)
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
	p.EnterRule(localctx, 34, TinyLangParserRULE_expr)

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

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserSYS_FUNC, TinyLangParserITEM, TinyLangParserNUMBER, TinyLangParserSTRING_RAW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.ExprOperand()
		}

		{
			p.SetState(210)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(211)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(212)
				p.Expr()
			}

		}

	case TinyLangParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(TinyLangParserT__1)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(216)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(217)
				p.Expr()
			}

		}
		{
			p.SetState(220)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(221)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(222)
				p.Expr()
			}

		}
		{
			p.SetState(225)
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
	p.EnterRule(localctx, 36, TinyLangParserRULE_exprOperand)

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

	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.FuncInvoc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(TinyLangParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(233)
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

func (s *BoolExprContext) AllBoolExprSingle() []IBoolExprSingleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExprSingleContext)(nil)).Elem())
	var tst = make([]IBoolExprSingleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExprSingleContext)
		}
	}

	return tst
}

func (s *BoolExprContext) BoolExprSingle(i int) IBoolExprSingleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprSingleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExprSingleContext)
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
	p.EnterRule(localctx, 38, TinyLangParserRULE_boolExpr)
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
		p.SetState(236)
		p.BoolExprSingle()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(237)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TinyLangParserT__11 || _la == TinyLangParserT__12) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(238)
				p.BoolExprSingle()
			}

		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 40, TinyLangParserRULE_boolExprSingle)

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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.BoolExprOperand()
		}

		{
			p.SetState(245)
			p.Match(TinyLangParserBOOL_SIGN)
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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(TinyLangParserT__1)
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
			p.Match(TinyLangParserBOOL_SIGN)
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(256)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(257)
				p.BoolExpr()
			}

		}
		{
			p.SetState(260)
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

func (s *BoolExprOperandContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
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
	p.EnterRule(localctx, 42, TinyLangParserRULE_boolExprOperand)

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

	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.ExprOperand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(266)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(267)
			p.ArrayElem()
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

func (s *StatementBodyContext) AllBREAK() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserBREAK)
}

func (s *StatementBodyContext) BREAK(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserBREAK, i)
}

func (s *StatementBodyContext) AllCONTINUE() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserCONTINUE)
}

func (s *StatementBodyContext) CONTINUE(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserCONTINUE, i)
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
	p.EnterRule(localctx, 44, TinyLangParserRULE_statementBody)
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
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLangParserT__7)|(1<<TinyLangParserWHILE)|(1<<TinyLangParserBREAK)|(1<<TinyLangParserCONTINUE)|(1<<TinyLangParserIF)|(1<<TinyLangParserFOR)|(1<<TinyLangParserSYS_FUNC)|(1<<TinyLangParserITEM))) != 0 {
		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(270)
				p.NewVariable()
			}

		case 2:
			{
				p.SetState(271)
				p.UpdVariable()
			}

		case 3:
			{
				p.SetState(272)
				p.FuncInvoc()
			}

		case 4:
			{
				p.SetState(273)
				p.IfElseSt()
			}

		case 5:
			{
				p.SetState(274)
				p.WhileSt()
			}

		case 6:
			{
				p.SetState(275)
				p.ForSt()
			}

		case 7:
			{
				p.SetState(276)
				p.Match(TinyLangParserBREAK)
			}

		case 8:
			{
				p.SetState(277)
				p.Match(TinyLangParserCONTINUE)
			}

		}

		p.SetState(282)
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
	p.EnterRule(localctx, 46, TinyLangParserRULE_ifElseSt)
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
		p.SetState(283)
		p.IfSt()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(284)
				p.ElseIfSt()
			}

		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserELSE {
		{
			p.SetState(290)
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

func (s *IfStContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
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
	p.EnterRule(localctx, 48, TinyLangParserRULE_ifSt)

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
		p.SetState(293)
		p.Match(TinyLangParserIF)
	}
	{
		p.SetState(294)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(295)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(296)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(297)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(298)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(301)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(302)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(303)
			p.StatementBody()
		}

	}
	{
		p.SetState(306)
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

func (s *ElseIfStContext) IF() antlr.TerminalNode {
	return s.GetToken(TinyLangParserIF, 0)
}

func (s *ElseIfStContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
}

func (s *ElseIfStContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
}

func (s *ElseIfStContext) ITEM() antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, 0)
}

func (s *ElseIfStContext) FuncInvoc() IFuncInvocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncInvocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncInvocContext)
}

func (s *ElseIfStContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
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
	p.EnterRule(localctx, 50, TinyLangParserRULE_elseIfSt)

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
		p.SetState(308)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(309)
		p.Match(TinyLangParserIF)
	}
	{
		p.SetState(310)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(311)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(312)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(313)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(314)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(317)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(318)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(319)
			p.StatementBody()
		}

	}
	{
		p.SetState(322)
		p.Match(TinyLangParserT__4)
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
	p.EnterRule(localctx, 52, TinyLangParserRULE_elseSt)

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
		p.SetState(324)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(325)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(326)
			p.StatementBody()
		}

	}
	{
		p.SetState(329)
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

func (s *WhileStContext) BOOL_VAL() antlr.TerminalNode {
	return s.GetToken(TinyLangParserBOOL_VAL, 0)
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
	p.EnterRule(localctx, 54, TinyLangParserRULE_whileSt)

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
		p.SetState(331)
		p.Match(TinyLangParserWHILE)
	}
	{
		p.SetState(332)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(337)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(333)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(334)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(335)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(336)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(339)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(340)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(341)
			p.StatementBody()
		}

	}
	{
		p.SetState(344)
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

func (s *ForStContext) AllUpdVariable() []IUpdVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdVariableContext)(nil)).Elem())
	var tst = make([]IUpdVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdVariableContext)
		}
	}

	return tst
}

func (s *ForStContext) UpdVariable(i int) IUpdVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdVariableContext)
}

func (s *ForStContext) BoolExpr() IBoolExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExprContext)
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
	p.EnterRule(localctx, 56, TinyLangParserRULE_forSt)

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
		p.SetState(346)
		p.Match(TinyLangParserFOR)
	}
	{
		p.SetState(347)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(348)
		p.UpdVariable()
	}
	{
		p.SetState(349)
		p.Match(TinyLangParserT__13)
	}
	{
		p.SetState(350)
		p.BoolExpr()
	}
	{
		p.SetState(351)
		p.Match(TinyLangParserT__13)
	}
	{
		p.SetState(352)
		p.UpdVariable()
	}
	{
		p.SetState(353)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(354)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(355)
		p.StatementBody()
	}
	{
		p.SetState(356)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}
