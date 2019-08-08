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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 41, 353,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 7, 2, 58, 10, 2, 12, 2, 14, 2, 61, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 72, 10, 3, 3, 3, 3, 3, 3, 3, 5, 3, 77,
	10, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 86, 10, 4, 12, 4,
	14, 4, 89, 11, 4, 5, 4, 91, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 102, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 7, 7, 112, 10, 7, 12, 7, 14, 7, 115, 11, 7, 5, 7, 117, 10,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 125, 10, 8, 3, 9, 3, 9, 5,
	9, 129, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 140, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 5, 10, 153, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 163, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	5, 13, 169, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 7, 16, 183, 10, 16, 12, 16, 14, 16, 186,
	11, 16, 3, 16, 3, 16, 3, 16, 7, 16, 191, 10, 16, 12, 16, 14, 16, 194, 11,
	16, 3, 16, 3, 16, 3, 16, 7, 16, 199, 10, 16, 12, 16, 14, 16, 202, 11, 16,
	5, 16, 204, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 210, 10, 17, 3,
	17, 3, 17, 3, 17, 5, 17, 215, 10, 17, 3, 17, 3, 17, 3, 17, 5, 17, 220,
	10, 17, 3, 17, 3, 17, 5, 17, 224, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 5, 18, 231, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 236, 10, 19, 12, 19,
	14, 19, 239, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 245, 10, 20, 3,
	20, 3, 20, 3, 20, 5, 20, 250, 10, 20, 3, 20, 3, 20, 3, 20, 5, 20, 255,
	10, 20, 3, 20, 3, 20, 5, 20, 259, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 5,
	21, 265, 10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 273,
	10, 22, 12, 22, 14, 22, 276, 11, 22, 3, 23, 3, 23, 7, 23, 280, 10, 23,
	12, 23, 14, 23, 283, 11, 23, 3, 23, 5, 23, 286, 10, 23, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 24, 5, 24, 294, 10, 24, 3, 24, 3, 24, 3, 24, 5, 24,
	299, 10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 5, 25, 310, 10, 25, 3, 25, 3, 25, 3, 25, 5, 25, 315, 10, 25, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 322, 10, 26, 3, 26, 3, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 332, 10, 27, 3, 27, 3, 27, 3, 27,
	5, 27, 337, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 2, 2, 29, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 2, 7, 4, 2, 25, 25, 31, 31, 5, 2, 26, 26, 28, 29,
	39, 39, 4, 2, 13, 13, 35, 35, 4, 2, 26, 26, 28, 29, 3, 2, 14, 15, 2, 404,
	2, 59, 3, 2, 2, 2, 4, 64, 3, 2, 2, 2, 6, 90, 3, 2, 2, 2, 8, 92, 3, 2, 2,
	2, 10, 103, 3, 2, 2, 2, 12, 116, 3, 2, 2, 2, 14, 124, 3, 2, 2, 2, 16, 128,
	3, 2, 2, 2, 18, 141, 3, 2, 2, 2, 20, 154, 3, 2, 2, 2, 22, 156, 3, 2, 2,
	2, 24, 168, 3, 2, 2, 2, 26, 170, 3, 2, 2, 2, 28, 175, 3, 2, 2, 2, 30, 203,
	3, 2, 2, 2, 32, 223, 3, 2, 2, 2, 34, 230, 3, 2, 2, 2, 36, 232, 3, 2, 2,
	2, 38, 258, 3, 2, 2, 2, 40, 264, 3, 2, 2, 2, 42, 274, 3, 2, 2, 2, 44, 277,
	3, 2, 2, 2, 46, 287, 3, 2, 2, 2, 48, 302, 3, 2, 2, 2, 50, 318, 3, 2, 2,
	2, 52, 325, 3, 2, 2, 2, 54, 340, 3, 2, 2, 2, 56, 58, 5, 4, 3, 2, 57, 56,
	3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2,
	60, 62, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 2, 2, 3, 63, 3, 3, 2,
	2, 2, 64, 65, 7, 3, 2, 2, 65, 66, 7, 31, 2, 2, 66, 67, 7, 4, 2, 2, 67,
	68, 5, 6, 4, 2, 68, 71, 7, 5, 2, 2, 69, 72, 5, 20, 11, 2, 70, 72, 7, 27,
	2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74,
	7, 6, 2, 2, 74, 76, 5, 42, 22, 2, 75, 77, 5, 8, 5, 2, 76, 75, 3, 2, 2,
	2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 7, 7, 2, 2, 79, 5, 3,
	2, 2, 2, 80, 81, 7, 31, 2, 2, 81, 87, 5, 20, 11, 2, 82, 83, 7, 8, 2, 2,
	83, 84, 7, 31, 2, 2, 84, 86, 5, 20, 11, 2, 85, 82, 3, 2, 2, 2, 86, 89,
	3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2,
	89, 87, 3, 2, 2, 2, 90, 80, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 7, 3, 2,
	2, 2, 92, 101, 7, 19, 2, 2, 93, 102, 5, 22, 12, 2, 94, 102, 5, 10, 6, 2,
	95, 102, 7, 31, 2, 2, 96, 102, 7, 36, 2, 2, 97, 102, 7, 34, 2, 2, 98, 102,
	7, 35, 2, 2, 99, 102, 5, 32, 17, 2, 100, 102, 5, 36, 19, 2, 101, 93, 3,
	2, 2, 2, 101, 94, 3, 2, 2, 2, 101, 95, 3, 2, 2, 2, 101, 96, 3, 2, 2, 2,
	101, 97, 3, 2, 2, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101, 100,
	3, 2, 2, 2, 102, 9, 3, 2, 2, 2, 103, 104, 9, 2, 2, 2, 104, 105, 7, 4, 2,
	2, 105, 106, 5, 12, 7, 2, 106, 107, 7, 5, 2, 2, 107, 11, 3, 2, 2, 2, 108,
	113, 5, 14, 8, 2, 109, 110, 7, 8, 2, 2, 110, 112, 5, 14, 8, 2, 111, 109,
	3, 2, 2, 2, 112, 115, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2,
	2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 116, 108, 3, 2, 2, 2,
	116, 117, 3, 2, 2, 2, 117, 13, 3, 2, 2, 2, 118, 125, 7, 35, 2, 2, 119,
	125, 7, 36, 2, 2, 120, 125, 7, 34, 2, 2, 121, 125, 7, 31, 2, 2, 122, 125,
	5, 10, 6, 2, 123, 125, 5, 22, 12, 2, 124, 118, 3, 2, 2, 2, 124, 119, 3,
	2, 2, 2, 124, 120, 3, 2, 2, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2, 2,
	2, 124, 123, 3, 2, 2, 2, 125, 15, 3, 2, 2, 2, 126, 129, 7, 31, 2, 2, 127,
	129, 5, 22, 12, 2, 128, 126, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 139, 7, 9, 2, 2, 131, 140, 5, 32, 17, 2, 132, 140, 5,
	10, 6, 2, 133, 140, 7, 34, 2, 2, 134, 140, 7, 36, 2, 2, 135, 140, 7, 35,
	2, 2, 136, 140, 7, 31, 2, 2, 137, 140, 5, 24, 13, 2, 138, 140, 5, 22, 12,
	2, 139, 131, 3, 2, 2, 2, 139, 132, 3, 2, 2, 2, 139, 133, 3, 2, 2, 2, 139,
	134, 3, 2, 2, 2, 139, 135, 3, 2, 2, 2, 139, 136, 3, 2, 2, 2, 139, 137,
	3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 17, 3, 2, 2, 2, 141, 142, 7, 10,
	2, 2, 142, 143, 7, 31, 2, 2, 143, 144, 5, 20, 11, 2, 144, 152, 7, 9, 2,
	2, 145, 153, 5, 32, 17, 2, 146, 153, 7, 34, 2, 2, 147, 153, 7, 36, 2, 2,
	148, 153, 7, 35, 2, 2, 149, 153, 7, 31, 2, 2, 150, 153, 5, 24, 13, 2, 151,
	153, 5, 10, 6, 2, 152, 145, 3, 2, 2, 2, 152, 146, 3, 2, 2, 2, 152, 147,
	3, 2, 2, 2, 152, 148, 3, 2, 2, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3, 2,
	2, 2, 152, 151, 3, 2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 155, 9, 3, 2, 2,
	155, 21, 3, 2, 2, 2, 156, 157, 7, 31, 2, 2, 157, 162, 7, 11, 2, 2, 158,
	163, 7, 35, 2, 2, 159, 163, 7, 31, 2, 2, 160, 163, 5, 10, 6, 2, 161, 163,
	5, 32, 17, 2, 162, 158, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 162, 160, 3,
	2, 2, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 7, 12, 2,
	2, 165, 23, 3, 2, 2, 2, 166, 169, 5, 28, 15, 2, 167, 169, 5, 26, 14, 2,
	168, 166, 3, 2, 2, 2, 168, 167, 3, 2, 2, 2, 169, 25, 3, 2, 2, 2, 170, 171,
	7, 11, 2, 2, 171, 172, 9, 4, 2, 2, 172, 173, 7, 12, 2, 2, 173, 174, 9,
	5, 2, 2, 174, 27, 3, 2, 2, 2, 175, 176, 7, 6, 2, 2, 176, 177, 5, 30, 16,
	2, 177, 178, 7, 7, 2, 2, 178, 29, 3, 2, 2, 2, 179, 184, 7, 34, 2, 2, 180,
	181, 7, 8, 2, 2, 181, 183, 7, 34, 2, 2, 182, 180, 3, 2, 2, 2, 183, 186,
	3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 204, 3, 2,
	2, 2, 186, 184, 3, 2, 2, 2, 187, 192, 7, 36, 2, 2, 188, 189, 7, 8, 2, 2,
	189, 191, 7, 36, 2, 2, 190, 188, 3, 2, 2, 2, 191, 194, 3, 2, 2, 2, 192,
	190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 204, 3, 2, 2, 2, 194, 192,
	3, 2, 2, 2, 195, 200, 7, 35, 2, 2, 196, 197, 7, 8, 2, 2, 197, 199, 7, 35,
	2, 2, 198, 196, 3, 2, 2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2,
	200, 201, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203,
	179, 3, 2, 2, 2, 203, 187, 3, 2, 2, 2, 203, 195, 3, 2, 2, 2, 204, 31, 3,
	2, 2, 2, 205, 206, 5, 34, 18, 2, 206, 209, 7, 40, 2, 2, 207, 210, 5, 34,
	18, 2, 208, 210, 5, 32, 17, 2, 209, 207, 3, 2, 2, 2, 209, 208, 3, 2, 2,
	2, 210, 224, 3, 2, 2, 2, 211, 214, 7, 4, 2, 2, 212, 215, 5, 34, 18, 2,
	213, 215, 5, 32, 17, 2, 214, 212, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215,
	216, 3, 2, 2, 2, 216, 219, 7, 40, 2, 2, 217, 220, 5, 34, 18, 2, 218, 220,
	5, 32, 17, 2, 219, 217, 3, 2, 2, 2, 219, 218, 3, 2, 2, 2, 220, 221, 3,
	2, 2, 2, 221, 222, 7, 5, 2, 2, 222, 224, 3, 2, 2, 2, 223, 205, 3, 2, 2,
	2, 223, 211, 3, 2, 2, 2, 224, 33, 3, 2, 2, 2, 225, 231, 5, 10, 6, 2, 226,
	231, 7, 35, 2, 2, 227, 231, 7, 31, 2, 2, 228, 231, 7, 36, 2, 2, 229, 231,
	5, 22, 12, 2, 230, 225, 3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 230, 227, 3,
	2, 2, 2, 230, 228, 3, 2, 2, 2, 230, 229, 3, 2, 2, 2, 231, 35, 3, 2, 2,
	2, 232, 237, 5, 38, 20, 2, 233, 234, 9, 6, 2, 2, 234, 236, 5, 38, 20, 2,
	235, 233, 3, 2, 2, 2, 236, 239, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237,
	238, 3, 2, 2, 2, 238, 37, 3, 2, 2, 2, 239, 237, 3, 2, 2, 2, 240, 241, 5,
	40, 21, 2, 241, 244, 7, 41, 2, 2, 242, 245, 5, 40, 21, 2, 243, 245, 5,
	36, 19, 2, 244, 242, 3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 259, 3, 2,
	2, 2, 246, 249, 7, 4, 2, 2, 247, 250, 5, 40, 21, 2, 248, 250, 5, 36, 19,
	2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251,
	254, 7, 41, 2, 2, 252, 255, 5, 40, 21, 2, 253, 255, 5, 36, 19, 2, 254,
	252, 3, 2, 2, 2, 254, 253, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 257,
	7, 5, 2, 2, 257, 259, 3, 2, 2, 2, 258, 240, 3, 2, 2, 2, 258, 246, 3, 2,
	2, 2, 259, 39, 3, 2, 2, 2, 260, 265, 5, 32, 17, 2, 261, 265, 5, 34, 18,
	2, 262, 265, 7, 34, 2, 2, 263, 265, 5, 22, 12, 2, 264, 260, 3, 2, 2, 2,
	264, 261, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 263, 3, 2, 2, 2, 265,
	41, 3, 2, 2, 2, 266, 273, 5, 18, 10, 2, 267, 273, 5, 16, 9, 2, 268, 273,
	5, 10, 6, 2, 269, 273, 5, 44, 23, 2, 270, 273, 5, 52, 27, 2, 271, 273,
	5, 54, 28, 2, 272, 266, 3, 2, 2, 2, 272, 267, 3, 2, 2, 2, 272, 268, 3,
	2, 2, 2, 272, 269, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 2, 2,
	2, 273, 276, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275,
	43, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 277, 281, 5, 46, 24, 2, 278, 280,
	5, 48, 25, 2, 279, 278, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281, 279, 3,
	2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2,
	2, 284, 286, 5, 50, 26, 2, 285, 284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2,
	286, 45, 3, 2, 2, 2, 287, 288, 7, 22, 2, 2, 288, 293, 7, 4, 2, 2, 289,
	294, 7, 34, 2, 2, 290, 294, 5, 36, 19, 2, 291, 294, 7, 31, 2, 2, 292, 294,
	5, 10, 6, 2, 293, 289, 3, 2, 2, 2, 293, 290, 3, 2, 2, 2, 293, 291, 3, 2,
	2, 2, 293, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 296, 7, 5, 2, 2,
	296, 298, 7, 6, 2, 2, 297, 299, 5, 42, 22, 2, 298, 297, 3, 2, 2, 2, 298,
	299, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 301, 7, 7, 2, 2, 301, 47, 3,
	2, 2, 2, 302, 303, 7, 24, 2, 2, 303, 304, 7, 22, 2, 2, 304, 309, 7, 4,
	2, 2, 305, 310, 7, 34, 2, 2, 306, 310, 5, 36, 19, 2, 307, 310, 7, 31, 2,
	2, 308, 310, 5, 10, 6, 2, 309, 305, 3, 2, 2, 2, 309, 306, 3, 2, 2, 2, 309,
	307, 3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312,
	7, 5, 2, 2, 312, 314, 7, 6, 2, 2, 313, 315, 5, 42, 22, 2, 314, 313, 3,
	2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 317, 7, 7, 2,
	2, 317, 49, 3, 2, 2, 2, 318, 319, 7, 24, 2, 2, 319, 321, 7, 6, 2, 2, 320,
	322, 5, 42, 22, 2, 321, 320, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322, 323,
	3, 2, 2, 2, 323, 324, 7, 7, 2, 2, 324, 51, 3, 2, 2, 2, 325, 326, 7, 18,
	2, 2, 326, 331, 7, 4, 2, 2, 327, 332, 7, 34, 2, 2, 328, 332, 5, 36, 19,
	2, 329, 332, 7, 31, 2, 2, 330, 332, 5, 10, 6, 2, 331, 327, 3, 2, 2, 2,
	331, 328, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331, 330, 3, 2, 2, 2, 332,
	333, 3, 2, 2, 2, 333, 334, 7, 5, 2, 2, 334, 336, 7, 6, 2, 2, 335, 337,
	5, 42, 22, 2, 336, 335, 3, 2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338, 3,
	2, 2, 2, 338, 339, 7, 7, 2, 2, 339, 53, 3, 2, 2, 2, 340, 341, 7, 23, 2,
	2, 341, 342, 7, 4, 2, 2, 342, 343, 5, 16, 9, 2, 343, 344, 7, 16, 2, 2,
	344, 345, 5, 36, 19, 2, 345, 346, 7, 16, 2, 2, 346, 347, 5, 16, 9, 2, 347,
	348, 7, 5, 2, 2, 348, 349, 7, 6, 2, 2, 349, 350, 5, 42, 22, 2, 350, 351,
	7, 7, 2, 2, 351, 55, 3, 2, 2, 2, 42, 59, 71, 76, 87, 90, 101, 113, 116,
	124, 128, 139, 152, 162, 168, 184, 192, 200, 203, 209, 214, 219, 223, 230,
	237, 244, 249, 254, 258, 264, 272, 274, 281, 285, 293, 298, 309, 314, 321,
	331, 336,
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
	"file", "funcInit", "funcArgs", "funcReturn", "funcInvoc", "funcArgsInvoc",
	"funcInvocArgs", "updVariable", "newVariable", "variableType", "arrayElem",
	"arrayInit", "arrayInitEmpty", "arrayInitValue", "arrayInitElems", "expr",
	"exprOperand", "boolExpr", "boolExprSingle", "boolExprOperand", "statementBody",
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
	TinyLangParserRULE_funcArgs        = 2
	TinyLangParserRULE_funcReturn      = 3
	TinyLangParserRULE_funcInvoc       = 4
	TinyLangParserRULE_funcArgsInvoc   = 5
	TinyLangParserRULE_funcInvocArgs   = 6
	TinyLangParserRULE_updVariable     = 7
	TinyLangParserRULE_newVariable     = 8
	TinyLangParserRULE_variableType    = 9
	TinyLangParserRULE_arrayElem       = 10
	TinyLangParserRULE_arrayInit       = 11
	TinyLangParserRULE_arrayInitEmpty  = 12
	TinyLangParserRULE_arrayInitValue  = 13
	TinyLangParserRULE_arrayInitElems  = 14
	TinyLangParserRULE_expr            = 15
	TinyLangParserRULE_exprOperand     = 16
	TinyLangParserRULE_boolExpr        = 17
	TinyLangParserRULE_boolExprSingle  = 18
	TinyLangParserRULE_boolExprOperand = 19
	TinyLangParserRULE_statementBody   = 20
	TinyLangParserRULE_ifElseSt        = 21
	TinyLangParserRULE_ifSt            = 22
	TinyLangParserRULE_elseIfSt        = 23
	TinyLangParserRULE_elseSt          = 24
	TinyLangParserRULE_whileSt         = 25
	TinyLangParserRULE_forSt           = 26
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TinyLangParserT__0 {
		{
			p.SetState(54)
			p.FuncInit()
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(60)
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

func (s *FuncInitContext) StatementBody() IStatementBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementBodyContext)
}

func (s *FuncInitContext) VariableType() IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
}

func (s *FuncInitContext) VOID() antlr.TerminalNode {
	return s.GetToken(TinyLangParserVOID, 0)
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
		p.SetState(62)
		p.Match(TinyLangParserT__0)
	}
	{
		p.SetState(63)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(64)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(65)
		p.FuncArgs()
	}
	{
		p.SetState(66)
		p.Match(TinyLangParserT__2)
	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserSTRING, TinyLangParserNUM, TinyLangParserBOOL, TinyLangParserARRAY:
		{
			p.SetState(67)
			p.VariableType()
		}

	case TinyLangParserVOID:
		{
			p.SetState(68)
			p.Match(TinyLangParserVOID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(71)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(72)
		p.StatementBody()
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserRETURN {
		{
			p.SetState(73)
			p.FuncReturn()
		}

	}
	{
		p.SetState(76)
		p.Match(TinyLangParserT__4)
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

func (s *FuncArgsContext) AllITEM() []antlr.TerminalNode {
	return s.GetTokens(TinyLangParserITEM)
}

func (s *FuncArgsContext) ITEM(i int) antlr.TerminalNode {
	return s.GetToken(TinyLangParserITEM, i)
}

func (s *FuncArgsContext) AllVariableType() []IVariableTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem())
	var tst = make([]IVariableTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableTypeContext)
		}
	}

	return tst
}

func (s *FuncArgsContext) VariableType(i int) IVariableTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableTypeContext)
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
	p.EnterRule(localctx, 4, TinyLangParserRULE_funcArgs)
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
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserITEM {
		{
			p.SetState(78)
			p.Match(TinyLangParserITEM)
		}
		{
			p.SetState(79)
			p.VariableType()
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(80)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(81)
				p.Match(TinyLangParserITEM)
			}
			{
				p.SetState(82)
				p.VariableType()
			}

			p.SetState(87)
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
	p.EnterRule(localctx, 6, TinyLangParserRULE_funcReturn)

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
		p.SetState(90)
		p.Match(TinyLangParserRETURN)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(91)
			p.ArrayElem()
		}

	case 2:
		{
			p.SetState(92)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(93)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(94)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		{
			p.SetState(95)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 6:
		{
			p.SetState(96)
			p.Match(TinyLangParserNUMBER)
		}

	case 7:
		{
			p.SetState(97)
			p.Expr()
		}

	case 8:
		{
			p.SetState(98)
			p.BoolExpr()
		}

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
	p.EnterRule(localctx, 8, TinyLangParserRULE_funcInvoc)
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
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserSYS_FUNC || _la == TinyLangParserITEM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(102)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(103)
		p.FuncArgsInvoc()
	}
	{
		p.SetState(104)
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
	p.EnterRule(localctx, 10, TinyLangParserRULE_funcArgsInvoc)
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
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-23)&-(0x1f+1)) == 0 && ((1<<uint((_la-23)))&((1<<(TinyLangParserSYS_FUNC-23))|(1<<(TinyLangParserITEM-23))|(1<<(TinyLangParserBOOL_VAL-23))|(1<<(TinyLangParserNUMBER-23))|(1<<(TinyLangParserSTRING_RAW-23)))) != 0 {
		{
			p.SetState(106)
			p.FuncInvocArgs()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(107)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(108)
				p.FuncInvocArgs()
			}

			p.SetState(113)
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
	p.EnterRule(localctx, 12, TinyLangParserRULE_funcInvocArgs)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.Match(TinyLangParserNUMBER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.Match(TinyLangParserITEM)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.FuncInvoc()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(121)
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
	p.EnterRule(localctx, 14, TinyLangParserRULE_updVariable)

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
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(124)
			p.Match(TinyLangParserITEM)
		}

	case 2:
		{
			p.SetState(125)
			p.ArrayElem()
		}

	}
	{
		p.SetState(128)
		p.Match(TinyLangParserT__6)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(129)
			p.Expr()
		}

	case 2:
		{
			p.SetState(130)
			p.FuncInvoc()
		}

	case 3:
		{
			p.SetState(131)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		{
			p.SetState(132)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		{
			p.SetState(133)
			p.Match(TinyLangParserNUMBER)
		}

	case 6:
		{
			p.SetState(134)
			p.Match(TinyLangParserITEM)
		}

	case 7:
		{
			p.SetState(135)
			p.ArrayInit()
		}

	case 8:
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
	p.EnterRule(localctx, 16, TinyLangParserRULE_newVariable)

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
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(143)
			p.Expr()
		}

	case 2:
		{
			p.SetState(144)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 3:
		{
			p.SetState(145)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 4:
		{
			p.SetState(146)
			p.Match(TinyLangParserNUMBER)
		}

	case 5:
		{
			p.SetState(147)
			p.Match(TinyLangParserITEM)
		}

	case 6:
		{
			p.SetState(148)
			p.ArrayInit()
		}

	case 7:
		{
			p.SetState(149)
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
	p.EnterRule(localctx, 18, TinyLangParserRULE_variableType)
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
		p.SetState(152)
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
	p.EnterRule(localctx, 20, TinyLangParserRULE_arrayElem)

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
		p.SetState(154)
		p.Match(TinyLangParserITEM)
	}
	{
		p.SetState(155)
		p.Match(TinyLangParserT__8)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(156)
			p.Match(TinyLangParserNUMBER)
		}

	case 2:
		{
			p.SetState(157)
			p.Match(TinyLangParserITEM)
		}

	case 3:
		{
			p.SetState(158)
			p.FuncInvoc()
		}

	case 4:
		{
			p.SetState(159)
			p.Expr()
		}

	}
	{
		p.SetState(162)
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
	p.EnterRule(localctx, 22, TinyLangParserRULE_arrayInit)

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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.ArrayInitValue()
		}

	case TinyLangParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
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
	p.EnterRule(localctx, 24, TinyLangParserRULE_arrayInitEmpty)
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
		p.SetState(168)
		p.Match(TinyLangParserT__8)
	}
	{
		p.SetState(169)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TinyLangParserT__10 || _la == TinyLangParserNUMBER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(170)
		p.Match(TinyLangParserT__9)
	}
	{
		p.SetState(171)
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
	p.EnterRule(localctx, 26, TinyLangParserRULE_arrayInitValue)

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
		p.SetState(173)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(174)
		p.ArrayInitElems()
	}
	{
		p.SetState(175)
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
	p.EnterRule(localctx, 28, TinyLangParserRULE_arrayInitElems)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserBOOL_VAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(TinyLangParserBOOL_VAL)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(178)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(179)
				p.Match(TinyLangParserBOOL_VAL)
			}

			p.SetState(184)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserSTRING_RAW:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Match(TinyLangParserSTRING_RAW)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(186)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(187)
				p.Match(TinyLangParserSTRING_RAW)
			}

			p.SetState(192)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case TinyLangParserNUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Match(TinyLangParserNUMBER)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == TinyLangParserT__5 {
			{
				p.SetState(194)
				p.Match(TinyLangParserT__5)
			}
			{
				p.SetState(195)
				p.Match(TinyLangParserNUMBER)
			}

			p.SetState(200)
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
	p.EnterRule(localctx, 30, TinyLangParserRULE_expr)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinyLangParserSYS_FUNC, TinyLangParserITEM, TinyLangParserNUMBER, TinyLangParserSTRING_RAW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.ExprOperand()
		}

		{
			p.SetState(204)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(205)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(206)
				p.Expr()
			}

		}

	case TinyLangParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(209)
			p.Match(TinyLangParserT__1)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(210)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(211)
				p.Expr()
			}

		}
		{
			p.SetState(214)
			p.Match(TinyLangParserNUM_SIGN)
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(215)
				p.ExprOperand()
			}

		case 2:
			{
				p.SetState(216)
				p.Expr()
			}

		}
		{
			p.SetState(219)
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
	p.EnterRule(localctx, 32, TinyLangParserRULE_exprOperand)

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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(223)
			p.FuncInvoc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(224)
			p.Match(TinyLangParserNUMBER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(225)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(226)
			p.Match(TinyLangParserSTRING_RAW)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(227)
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
	p.EnterRule(localctx, 34, TinyLangParserRULE_boolExpr)
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
		p.SetState(230)
		p.BoolExprSingle()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(231)
				_la = p.GetTokenStream().LA(1)

				if !(_la == TinyLangParserT__11 || _la == TinyLangParserT__12) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(232)
				p.BoolExprSingle()
			}

		}
		p.SetState(237)
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
	p.EnterRule(localctx, 36, TinyLangParserRULE_boolExprSingle)

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

	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.BoolExprOperand()
		}

		{
			p.SetState(239)
			p.Match(TinyLangParserBOOL_SIGN)
		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(240)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(241)
				p.BoolExpr()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(244)
			p.Match(TinyLangParserT__1)
		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(245)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(246)
				p.BoolExpr()
			}

		}
		{
			p.SetState(249)
			p.Match(TinyLangParserBOOL_SIGN)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(250)
				p.BoolExprOperand()
			}

		case 2:
			{
				p.SetState(251)
				p.BoolExpr()
			}

		}
		{
			p.SetState(254)
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
	p.EnterRule(localctx, 38, TinyLangParserRULE_boolExprOperand)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Expr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(259)
			p.ExprOperand()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(261)
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
	p.EnterRule(localctx, 40, TinyLangParserRULE_statementBody)
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
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinyLangParserT__7)|(1<<TinyLangParserWHILE)|(1<<TinyLangParserIF)|(1<<TinyLangParserFOR)|(1<<TinyLangParserSYS_FUNC)|(1<<TinyLangParserITEM))) != 0 {
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(264)
				p.NewVariable()
			}

		case 2:
			{
				p.SetState(265)
				p.UpdVariable()
			}

		case 3:
			{
				p.SetState(266)
				p.FuncInvoc()
			}

		case 4:
			{
				p.SetState(267)
				p.IfElseSt()
			}

		case 5:
			{
				p.SetState(268)
				p.WhileSt()
			}

		case 6:
			{
				p.SetState(269)
				p.ForSt()
			}

		}

		p.SetState(274)
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
	p.EnterRule(localctx, 42, TinyLangParserRULE_ifElseSt)
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
		p.SetState(275)
		p.IfSt()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(276)
				p.ElseIfSt()
			}

		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == TinyLangParserELSE {
		{
			p.SetState(282)
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
	p.EnterRule(localctx, 44, TinyLangParserRULE_ifSt)

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
		p.SetState(285)
		p.Match(TinyLangParserIF)
	}
	{
		p.SetState(286)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(287)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(288)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(289)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(290)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(293)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(294)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(295)
			p.StatementBody()
		}

	}
	{
		p.SetState(298)
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
	p.EnterRule(localctx, 46, TinyLangParserRULE_elseIfSt)

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
		p.SetState(300)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(301)
		p.Match(TinyLangParserIF)
	}
	{
		p.SetState(302)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(303)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(304)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(305)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(306)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(309)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(310)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(311)
			p.StatementBody()
		}

	}
	{
		p.SetState(314)
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
	p.EnterRule(localctx, 48, TinyLangParserRULE_elseSt)

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
		p.SetState(316)
		p.Match(TinyLangParserELSE)
	}
	{
		p.SetState(317)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(318)
			p.StatementBody()
		}

	}
	{
		p.SetState(321)
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
	p.EnterRule(localctx, 50, TinyLangParserRULE_whileSt)

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
		p.SetState(323)
		p.Match(TinyLangParserWHILE)
	}
	{
		p.SetState(324)
		p.Match(TinyLangParserT__1)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(325)
			p.Match(TinyLangParserBOOL_VAL)
		}

	case 2:
		{
			p.SetState(326)
			p.BoolExpr()
		}

	case 3:
		{
			p.SetState(327)
			p.Match(TinyLangParserITEM)
		}

	case 4:
		{
			p.SetState(328)
			p.FuncInvoc()
		}

	}
	{
		p.SetState(331)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(332)
		p.Match(TinyLangParserT__3)
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(333)
			p.StatementBody()
		}

	}
	{
		p.SetState(336)
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
	p.EnterRule(localctx, 52, TinyLangParserRULE_forSt)

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
		p.SetState(338)
		p.Match(TinyLangParserFOR)
	}
	{
		p.SetState(339)
		p.Match(TinyLangParserT__1)
	}
	{
		p.SetState(340)
		p.UpdVariable()
	}
	{
		p.SetState(341)
		p.Match(TinyLangParserT__13)
	}
	{
		p.SetState(342)
		p.BoolExpr()
	}
	{
		p.SetState(343)
		p.Match(TinyLangParserT__13)
	}
	{
		p.SetState(344)
		p.UpdVariable()
	}
	{
		p.SetState(345)
		p.Match(TinyLangParserT__2)
	}
	{
		p.SetState(346)
		p.Match(TinyLangParserT__3)
	}
	{
		p.SetState(347)
		p.StatementBody()
	}
	{
		p.SetState(348)
		p.Match(TinyLangParserT__4)
	}

	return localctx
}
