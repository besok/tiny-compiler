// Code generated from C:/projects/tiny-compiler/src/antlr\TinyLang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 277,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 111,
	10, 14, 13, 14, 14, 14, 112, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 5, 22, 174, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 6, 28,
	203, 10, 28, 13, 28, 14, 28, 204, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 220, 10, 31,
	3, 32, 3, 32, 3, 32, 3, 32, 7, 32, 226, 10, 32, 12, 32, 14, 32, 229, 11,
	32, 3, 32, 3, 32, 3, 33, 3, 33, 7, 33, 235, 10, 33, 12, 33, 14, 33, 238,
	11, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 247, 10,
	35, 12, 35, 14, 35, 250, 11, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	36, 3, 36, 3, 36, 3, 36, 5, 36, 261, 10, 36, 3, 37, 3, 37, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 5, 38, 270, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 5,
	39, 276, 10, 39, 4, 227, 248, 2, 40, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 3, 2, 8, 5, 2, 11, 12, 15, 15,
	34, 34, 3, 2, 50, 59, 3, 2, 99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 6,
	2, 39, 39, 44, 45, 47, 47, 49, 49, 4, 2, 62, 62, 64, 64, 2, 291, 2, 3,
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11,
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2,
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2,
	2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2,
	2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2,
	2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3,
	2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57,
	3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2,
	65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2,
	2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2,
	2, 5, 84, 3, 2, 2, 2, 7, 86, 3, 2, 2, 2, 9, 88, 3, 2, 2, 2, 11, 90, 3,
	2, 2, 2, 13, 92, 3, 2, 2, 2, 15, 94, 3, 2, 2, 2, 17, 96, 3, 2, 2, 2, 19,
	100, 3, 2, 2, 2, 21, 102, 3, 2, 2, 2, 23, 104, 3, 2, 2, 2, 25, 107, 3,
	2, 2, 2, 27, 110, 3, 2, 2, 2, 29, 116, 3, 2, 2, 2, 31, 122, 3, 2, 2, 2,
	33, 129, 3, 2, 2, 2, 35, 135, 3, 2, 2, 2, 37, 144, 3, 2, 2, 2, 39, 147,
	3, 2, 2, 2, 41, 151, 3, 2, 2, 2, 43, 173, 3, 2, 2, 2, 45, 175, 3, 2, 2,
	2, 47, 182, 3, 2, 2, 2, 49, 187, 3, 2, 2, 2, 51, 191, 3, 2, 2, 2, 53, 196,
	3, 2, 2, 2, 55, 202, 3, 2, 2, 2, 57, 206, 3, 2, 2, 2, 59, 211, 3, 2, 2,
	2, 61, 219, 3, 2, 2, 2, 63, 221, 3, 2, 2, 2, 65, 232, 3, 2, 2, 2, 67, 239,
	3, 2, 2, 2, 69, 242, 3, 2, 2, 2, 71, 256, 3, 2, 2, 2, 73, 262, 3, 2, 2,
	2, 75, 269, 3, 2, 2, 2, 77, 275, 3, 2, 2, 2, 79, 80, 7, 104, 2, 2, 80,
	81, 7, 119, 2, 2, 81, 82, 7, 112, 2, 2, 82, 83, 7, 101, 2, 2, 83, 4, 3,
	2, 2, 2, 84, 85, 7, 42, 2, 2, 85, 6, 3, 2, 2, 2, 86, 87, 7, 43, 2, 2, 87,
	8, 3, 2, 2, 2, 88, 89, 7, 125, 2, 2, 89, 10, 3, 2, 2, 2, 90, 91, 7, 127,
	2, 2, 91, 12, 3, 2, 2, 2, 92, 93, 7, 46, 2, 2, 93, 14, 3, 2, 2, 2, 94,
	95, 7, 63, 2, 2, 95, 16, 3, 2, 2, 2, 96, 97, 7, 120, 2, 2, 97, 98, 7, 99,
	2, 2, 98, 99, 7, 116, 2, 2, 99, 18, 3, 2, 2, 2, 100, 101, 7, 93, 2, 2,
	101, 20, 3, 2, 2, 2, 102, 103, 7, 95, 2, 2, 103, 22, 3, 2, 2, 2, 104, 105,
	7, 48, 2, 2, 105, 106, 7, 48, 2, 2, 106, 24, 3, 2, 2, 2, 107, 108, 7, 61,
	2, 2, 108, 26, 3, 2, 2, 2, 109, 111, 9, 2, 2, 2, 110, 109, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113,
	114, 3, 2, 2, 2, 114, 115, 8, 14, 2, 2, 115, 28, 3, 2, 2, 2, 116, 117,
	7, 121, 2, 2, 117, 118, 7, 106, 2, 2, 118, 119, 7, 107, 2, 2, 119, 120,
	7, 110, 2, 2, 120, 121, 7, 103, 2, 2, 121, 30, 3, 2, 2, 2, 122, 123, 7,
	116, 2, 2, 123, 124, 7, 103, 2, 2, 124, 125, 7, 118, 2, 2, 125, 126, 7,
	119, 2, 2, 126, 127, 7, 116, 2, 2, 127, 128, 7, 112, 2, 2, 128, 32, 3,
	2, 2, 2, 129, 130, 7, 100, 2, 2, 130, 131, 7, 116, 2, 2, 131, 132, 7, 103,
	2, 2, 132, 133, 7, 99, 2, 2, 133, 134, 7, 109, 2, 2, 134, 34, 3, 2, 2,
	2, 135, 136, 7, 101, 2, 2, 136, 137, 7, 113, 2, 2, 137, 138, 7, 112, 2,
	2, 138, 139, 7, 118, 2, 2, 139, 140, 7, 107, 2, 2, 140, 141, 7, 112, 2,
	2, 141, 142, 7, 119, 2, 2, 142, 143, 7, 103, 2, 2, 143, 36, 3, 2, 2, 2,
	144, 145, 7, 107, 2, 2, 145, 146, 7, 104, 2, 2, 146, 38, 3, 2, 2, 2, 147,
	148, 7, 104, 2, 2, 148, 149, 7, 113, 2, 2, 149, 150, 7, 116, 2, 2, 150,
	40, 3, 2, 2, 2, 151, 152, 7, 103, 2, 2, 152, 153, 7, 110, 2, 2, 153, 154,
	7, 117, 2, 2, 154, 155, 7, 103, 2, 2, 155, 42, 3, 2, 2, 2, 156, 157, 7,
	67, 2, 2, 157, 158, 7, 102, 2, 2, 158, 174, 7, 102, 2, 2, 159, 160, 7,
	78, 2, 2, 160, 161, 7, 103, 2, 2, 161, 174, 7, 112, 2, 2, 162, 163, 7,
	81, 2, 2, 163, 164, 7, 119, 2, 2, 164, 165, 7, 118, 2, 2, 165, 166, 7,
	114, 2, 2, 166, 167, 7, 119, 2, 2, 167, 174, 7, 118, 2, 2, 168, 169, 7,
	75, 2, 2, 169, 170, 7, 112, 2, 2, 170, 171, 7, 114, 2, 2, 171, 172, 7,
	119, 2, 2, 172, 174, 7, 118, 2, 2, 173, 156, 3, 2, 2, 2, 173, 159, 3, 2,
	2, 2, 173, 162, 3, 2, 2, 2, 173, 168, 3, 2, 2, 2, 174, 44, 3, 2, 2, 2,
	175, 176, 7, 117, 2, 2, 176, 177, 7, 118, 2, 2, 177, 178, 7, 116, 2, 2,
	178, 179, 7, 107, 2, 2, 179, 180, 7, 112, 2, 2, 180, 181, 7, 105, 2, 2,
	181, 46, 3, 2, 2, 2, 182, 183, 7, 120, 2, 2, 183, 184, 7, 113, 2, 2, 184,
	185, 7, 107, 2, 2, 185, 186, 7, 102, 2, 2, 186, 48, 3, 2, 2, 2, 187, 188,
	7, 112, 2, 2, 188, 189, 7, 119, 2, 2, 189, 190, 7, 111, 2, 2, 190, 50,
	3, 2, 2, 2, 191, 192, 7, 100, 2, 2, 192, 193, 7, 113, 2, 2, 193, 194, 7,
	113, 2, 2, 194, 195, 7, 110, 2, 2, 195, 52, 3, 2, 2, 2, 196, 197, 7, 112,
	2, 2, 197, 198, 7, 119, 2, 2, 198, 199, 7, 110, 2, 2, 199, 200, 7, 110,
	2, 2, 200, 54, 3, 2, 2, 2, 201, 203, 9, 3, 2, 2, 202, 201, 3, 2, 2, 2,
	203, 204, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205,
	56, 3, 2, 2, 2, 206, 207, 7, 118, 2, 2, 207, 208, 7, 116, 2, 2, 208, 209,
	7, 119, 2, 2, 209, 210, 7, 103, 2, 2, 210, 58, 3, 2, 2, 2, 211, 212, 7,
	104, 2, 2, 212, 213, 7, 99, 2, 2, 213, 214, 7, 110, 2, 2, 214, 215, 7,
	117, 2, 2, 215, 216, 7, 103, 2, 2, 216, 60, 3, 2, 2, 2, 217, 220, 5, 57,
	29, 2, 218, 220, 5, 59, 30, 2, 219, 217, 3, 2, 2, 2, 219, 218, 3, 2, 2,
	2, 220, 62, 3, 2, 2, 2, 221, 227, 7, 36, 2, 2, 222, 223, 7, 94, 2, 2, 223,
	226, 7, 36, 2, 2, 224, 226, 11, 2, 2, 2, 225, 222, 3, 2, 2, 2, 225, 224,
	3, 2, 2, 2, 226, 229, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 227, 225, 3, 2,
	2, 2, 228, 230, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 230, 231, 7, 36, 2, 2,
	231, 64, 3, 2, 2, 2, 232, 236, 9, 4, 2, 2, 233, 235, 9, 5, 2, 2, 234, 233,
	3, 2, 2, 2, 235, 238, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2,
	2, 2, 237, 66, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 239, 240, 7, 93, 2, 2,
	240, 241, 7, 95, 2, 2, 241, 68, 3, 2, 2, 2, 242, 243, 7, 49, 2, 2, 243,
	244, 7, 44, 2, 2, 244, 248, 3, 2, 2, 2, 245, 247, 11, 2, 2, 2, 246, 245,
	3, 2, 2, 2, 247, 250, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 248, 246, 3, 2,
	2, 2, 249, 251, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251, 252, 7, 44, 2, 2,
	252, 253, 7, 49, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255, 8, 35, 2, 2, 255,
	70, 3, 2, 2, 2, 256, 260, 5, 67, 34, 2, 257, 261, 5, 45, 23, 2, 258, 261,
	5, 49, 25, 2, 259, 261, 5, 51, 26, 2, 260, 257, 3, 2, 2, 2, 260, 258, 3,
	2, 2, 2, 260, 259, 3, 2, 2, 2, 261, 72, 3, 2, 2, 2, 262, 263, 9, 6, 2,
	2, 263, 74, 3, 2, 2, 2, 264, 265, 7, 63, 2, 2, 265, 270, 7, 63, 2, 2, 266,
	267, 7, 35, 2, 2, 267, 270, 7, 63, 2, 2, 268, 270, 9, 7, 2, 2, 269, 264,
	3, 2, 2, 2, 269, 266, 3, 2, 2, 2, 269, 268, 3, 2, 2, 2, 270, 76, 3, 2,
	2, 2, 271, 272, 7, 40, 2, 2, 272, 276, 7, 40, 2, 2, 273, 274, 7, 126, 2,
	2, 274, 276, 7, 126, 2, 2, 275, 271, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2,
	276, 78, 3, 2, 2, 2, 14, 2, 112, 173, 204, 219, 225, 227, 236, 248, 260,
	269, 275, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'func'", "'('", "')'", "'{'", "'}'", "','", "'='", "'var'", "'['",
	"']'", "'..'", "';'", "", "'while'", "'return'", "'break'", "'continue'",
	"'if'", "'for'", "'else'", "", "'string'", "'void'", "'num'", "'bool'",
	"'null'", "", "'true'", "'false'", "", "", "", "'[]'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "WHILE", "RETURN",
	"BREAK", "CONTINUE", "IF", "FOR", "ELSE", "SYS_FUNC", "STRING", "VOID",
	"NUM", "BOOL", "NULL", "NUMBER", "TRUE", "FALSE", "BOOL_VAL", "STRING_RAW",
	"ITEM", "SQ", "COMMENT", "ARRAY", "NUM_SIGN", "BOOL_SIGN", "BOOL_PL",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "WS", "WHILE", "RETURN", "BREAK", "CONTINUE",
	"IF", "FOR", "ELSE", "SYS_FUNC", "STRING", "VOID", "NUM", "BOOL", "NULL",
	"NUMBER", "TRUE", "FALSE", "BOOL_VAL", "STRING_RAW", "ITEM", "SQ", "COMMENT",
	"ARRAY", "NUM_SIGN", "BOOL_SIGN", "BOOL_PL",
}

type TinyLangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewTinyLangLexer(input antlr.CharStream) *TinyLangLexer {

	l := new(TinyLangLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "TinyLang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TinyLangLexer tokens.
const (
	TinyLangLexerT__0       = 1
	TinyLangLexerT__1       = 2
	TinyLangLexerT__2       = 3
	TinyLangLexerT__3       = 4
	TinyLangLexerT__4       = 5
	TinyLangLexerT__5       = 6
	TinyLangLexerT__6       = 7
	TinyLangLexerT__7       = 8
	TinyLangLexerT__8       = 9
	TinyLangLexerT__9       = 10
	TinyLangLexerT__10      = 11
	TinyLangLexerT__11      = 12
	TinyLangLexerWS         = 13
	TinyLangLexerWHILE      = 14
	TinyLangLexerRETURN     = 15
	TinyLangLexerBREAK      = 16
	TinyLangLexerCONTINUE   = 17
	TinyLangLexerIF         = 18
	TinyLangLexerFOR        = 19
	TinyLangLexerELSE       = 20
	TinyLangLexerSYS_FUNC   = 21
	TinyLangLexerSTRING     = 22
	TinyLangLexerVOID       = 23
	TinyLangLexerNUM        = 24
	TinyLangLexerBOOL       = 25
	TinyLangLexerNULL       = 26
	TinyLangLexerNUMBER     = 27
	TinyLangLexerTRUE       = 28
	TinyLangLexerFALSE      = 29
	TinyLangLexerBOOL_VAL   = 30
	TinyLangLexerSTRING_RAW = 31
	TinyLangLexerITEM       = 32
	TinyLangLexerSQ         = 33
	TinyLangLexerCOMMENT    = 34
	TinyLangLexerARRAY      = 35
	TinyLangLexerNUM_SIGN   = 36
	TinyLangLexerBOOL_SIGN  = 37
	TinyLangLexerBOOL_PL    = 38
)
