// Generated from ../FoxySheep.g4 by ANTLR 4.7.

package parser // FoxySheep

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 188, 525,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 6, 2, 29, 10, 2, 13, 2, 14, 2, 30, 3, 2, 5, 2, 34, 10,
	2, 7, 2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 75, 10, 3, 3, 3, 3, 3, 5, 3, 79,
	10, 3, 3, 3, 5, 3, 82, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3,
	90, 10, 3, 12, 3, 14, 3, 93, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 113, 10, 3, 3, 3, 3, 3, 5, 3, 117, 10, 3, 7, 3, 119, 10, 3, 12, 3,
	14, 3, 122, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 142, 10, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 209, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	322, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 339, 10, 3, 13, 3, 14, 3, 340, 3, 3,
	3, 3, 3, 3, 5, 3, 346, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 361, 10, 3, 7, 3, 363, 10, 3,
	12, 3, 14, 3, 366, 11, 3, 3, 4, 5, 4, 369, 10, 4, 3, 4, 3, 4, 3, 5, 5,
	5, 374, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 382, 10, 5, 3,
	6, 3, 6, 3, 6, 5, 6, 387, 10, 6, 3, 6, 5, 6, 390, 10, 6, 3, 6, 3, 6, 5,
	6, 394, 10, 6, 3, 6, 5, 6, 397, 10, 6, 5, 6, 399, 10, 6, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 405, 10, 7, 5, 7, 407, 10, 7, 3, 8, 3, 8, 5, 8, 411, 10,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 5, 9, 417, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 424, 10, 10, 3, 11, 5, 11, 427, 10, 11, 3, 11, 3, 11, 5,
	11, 431, 10, 11, 7, 11, 433, 10, 11, 12, 11, 14, 11, 436, 11, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 448,
	10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 454, 10, 13, 12, 13, 14, 13,
	457, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 6, 13, 463, 10, 13, 13, 13, 14,
	13, 464, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	5, 13, 476, 10, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 7, 13, 520, 10, 13, 12, 13, 14, 13, 523, 11, 13, 3,
	13, 2, 4, 4, 24, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 31,
	4, 2, 79, 80, 82, 82, 3, 2, 114, 115, 3, 2, 160, 163, 4, 2, 56, 56, 124,
	124, 4, 2, 42, 42, 106, 106, 5, 2, 116, 117, 120, 120, 122, 122, 3, 2,
	176, 177, 3, 2, 184, 187, 4, 2, 41, 41, 43, 47, 3, 2, 48, 51, 3, 2, 39,
	40, 3, 2, 52, 55, 3, 2, 57, 59, 3, 2, 60, 61, 4, 2, 24, 24, 62, 63, 3,
	2, 65, 66, 3, 2, 67, 68, 3, 2, 69, 72, 3, 2, 89, 92, 3, 2, 93, 94, 3, 2,
	100, 103, 5, 2, 42, 42, 105, 107, 109, 109, 3, 2, 123, 124, 3, 2, 134,
	137, 3, 2, 76, 77, 3, 2, 37, 38, 4, 2, 4, 4, 6, 6, 3, 2, 160, 161, 4, 2,
	169, 169, 174, 174, 2, 658, 2, 26, 3, 2, 2, 2, 4, 141, 3, 2, 2, 2, 6, 368,
	3, 2, 2, 2, 8, 381, 3, 2, 2, 2, 10, 398, 3, 2, 2, 2, 12, 406, 3, 2, 2,
	2, 14, 408, 3, 2, 2, 2, 16, 416, 3, 2, 2, 2, 18, 423, 3, 2, 2, 2, 20, 426,
	3, 2, 2, 2, 22, 447, 3, 2, 2, 2, 24, 475, 3, 2, 2, 2, 26, 37, 5, 4, 3,
	2, 27, 29, 7, 181, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 30, 28,
	3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 34, 5, 4, 3, 2,
	33, 32, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2, 35, 28, 3,
	2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38,
	3, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 41, 8, 3, 1, 2, 41, 142, 5, 10,
	6, 2, 42, 142, 7, 7, 2, 2, 43, 44, 7, 9, 2, 2, 44, 45, 5, 4, 3, 2, 45,
	46, 7, 10, 2, 2, 46, 142, 3, 2, 2, 2, 47, 48, 7, 11, 2, 2, 48, 49, 5, 20,
	11, 2, 49, 50, 7, 12, 2, 2, 50, 142, 3, 2, 2, 2, 51, 52, 7, 18, 2, 2, 52,
	53, 5, 20, 11, 2, 53, 54, 7, 19, 2, 2, 54, 142, 3, 2, 2, 2, 55, 56, 7,
	20, 2, 2, 56, 57, 5, 4, 3, 2, 57, 58, 7, 21, 2, 2, 58, 142, 3, 2, 2, 2,
	59, 60, 7, 22, 2, 2, 60, 61, 5, 4, 3, 2, 61, 62, 7, 23, 2, 2, 62, 142,
	3, 2, 2, 2, 63, 64, 7, 28, 2, 2, 64, 65, 5, 20, 11, 2, 65, 66, 7, 29, 2,
	2, 66, 142, 3, 2, 2, 2, 67, 68, 7, 30, 2, 2, 68, 69, 5, 20, 11, 2, 69,
	70, 7, 31, 2, 2, 70, 142, 3, 2, 2, 2, 71, 142, 5, 18, 10, 2, 72, 142, 5,
	16, 9, 2, 73, 75, 5, 6, 4, 2, 74, 73, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75,
	76, 3, 2, 2, 2, 76, 78, 9, 2, 2, 2, 77, 79, 5, 4, 3, 2, 78, 77, 3, 2, 2,
	2, 78, 79, 3, 2, 2, 2, 79, 142, 3, 2, 2, 2, 80, 82, 5, 6, 4, 2, 81, 80,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 142, 7, 81, 2,
	2, 84, 142, 5, 6, 4, 2, 85, 86, 7, 36, 2, 2, 86, 142, 7, 7, 2, 2, 87, 91,
	7, 167, 2, 2, 88, 90, 5, 24, 13, 2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2,
	2, 91, 89, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 94, 142, 7, 170, 2, 2, 95, 96, 7, 175, 2, 2, 96, 142, 5, 4,
	3, 78, 97, 98, 9, 3, 2, 2, 98, 142, 5, 4, 3, 76, 99, 100, 7, 126, 2, 2,
	100, 101, 5, 4, 3, 2, 101, 102, 7, 127, 2, 2, 102, 103, 5, 4, 3, 65, 103,
	142, 3, 2, 2, 2, 104, 105, 7, 139, 2, 2, 105, 142, 5, 4, 3, 64, 106, 107,
	7, 140, 2, 2, 107, 142, 5, 4, 3, 63, 108, 109, 9, 4, 2, 2, 109, 142, 5,
	4, 3, 57, 110, 112, 7, 188, 2, 2, 111, 113, 5, 4, 3, 2, 112, 111, 3, 2,
	2, 2, 112, 113, 3, 2, 2, 2, 113, 120, 3, 2, 2, 2, 114, 116, 7, 132, 2,
	2, 115, 117, 5, 4, 3, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117,
	119, 3, 2, 2, 2, 118, 114, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118,
	3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 142, 3, 2, 2, 2, 122, 120, 3, 2,
	2, 2, 123, 124, 9, 5, 2, 2, 124, 142, 5, 4, 3, 32, 125, 126, 5, 6, 4, 2,
	126, 127, 7, 86, 2, 2, 127, 128, 5, 4, 3, 21, 128, 142, 3, 2, 2, 2, 129,
	130, 5, 6, 4, 2, 130, 131, 7, 108, 2, 2, 131, 132, 5, 4, 3, 2, 132, 133,
	9, 6, 2, 2, 133, 134, 5, 4, 3, 6, 134, 142, 3, 2, 2, 2, 135, 136, 5, 6,
	4, 2, 136, 137, 7, 108, 2, 2, 137, 138, 5, 4, 3, 2, 138, 139, 7, 42, 2,
	2, 139, 140, 7, 110, 2, 2, 140, 142, 3, 2, 2, 2, 141, 40, 3, 2, 2, 2, 141,
	42, 3, 2, 2, 2, 141, 43, 3, 2, 2, 2, 141, 47, 3, 2, 2, 2, 141, 51, 3, 2,
	2, 2, 141, 55, 3, 2, 2, 2, 141, 59, 3, 2, 2, 2, 141, 63, 3, 2, 2, 2, 141,
	67, 3, 2, 2, 2, 141, 71, 3, 2, 2, 2, 141, 72, 3, 2, 2, 2, 141, 74, 3, 2,
	2, 2, 141, 81, 3, 2, 2, 2, 141, 84, 3, 2, 2, 2, 141, 85, 3, 2, 2, 2, 141,
	87, 3, 2, 2, 2, 141, 95, 3, 2, 2, 2, 141, 97, 3, 2, 2, 2, 141, 99, 3, 2,
	2, 2, 141, 104, 3, 2, 2, 2, 141, 106, 3, 2, 2, 2, 141, 108, 3, 2, 2, 2,
	141, 110, 3, 2, 2, 2, 141, 123, 3, 2, 2, 2, 141, 125, 3, 2, 2, 2, 141,
	129, 3, 2, 2, 2, 141, 135, 3, 2, 2, 2, 142, 364, 3, 2, 2, 2, 143, 144,
	12, 81, 2, 2, 144, 145, 7, 78, 2, 2, 145, 363, 5, 4, 3, 82, 146, 147, 12,
	75, 2, 2, 147, 148, 7, 118, 2, 2, 148, 363, 5, 4, 3, 76, 149, 150, 12,
	74, 2, 2, 150, 151, 7, 121, 2, 2, 151, 363, 5, 4, 3, 75, 152, 153, 12,
	73, 2, 2, 153, 154, 7, 119, 2, 2, 154, 363, 5, 4, 3, 73, 155, 156, 12,
	72, 2, 2, 156, 157, 7, 138, 2, 2, 157, 158, 5, 4, 3, 2, 158, 159, 7, 138,
	2, 2, 159, 160, 5, 4, 3, 73, 160, 363, 3, 2, 2, 2, 161, 162, 12, 71, 2,
	2, 162, 163, 9, 7, 2, 2, 163, 363, 5, 4, 3, 71, 164, 165, 12, 67, 2, 2,
	165, 166, 7, 125, 2, 2, 166, 363, 5, 4, 3, 68, 167, 168, 12, 66, 2, 2,
	168, 169, 7, 112, 2, 2, 169, 363, 5, 4, 3, 66, 170, 171, 12, 62, 2, 2,
	171, 172, 7, 141, 2, 2, 172, 363, 5, 4, 3, 63, 173, 174, 12, 61, 2, 2,
	174, 175, 7, 142, 2, 2, 175, 363, 5, 4, 3, 62, 176, 177, 12, 60, 2, 2,
	177, 178, 7, 143, 2, 2, 178, 363, 5, 4, 3, 61, 179, 180, 12, 59, 2, 2,
	180, 181, 7, 128, 2, 2, 181, 363, 5, 4, 3, 60, 182, 183, 12, 58, 2, 2,
	183, 184, 7, 110, 2, 2, 184, 363, 5, 4, 3, 59, 185, 186, 12, 56, 2, 2,
	186, 187, 9, 8, 2, 2, 187, 363, 5, 4, 3, 57, 188, 189, 12, 55, 2, 2, 189,
	190, 7, 129, 2, 2, 190, 363, 5, 4, 3, 56, 191, 192, 12, 54, 2, 2, 192,
	193, 7, 145, 2, 2, 193, 363, 5, 4, 3, 55, 194, 195, 12, 53, 2, 2, 195,
	196, 7, 146, 2, 2, 196, 363, 5, 4, 3, 54, 197, 198, 12, 52, 2, 2, 198,
	199, 7, 147, 2, 2, 199, 363, 5, 4, 3, 53, 200, 201, 12, 51, 2, 2, 201,
	202, 7, 148, 2, 2, 202, 363, 5, 4, 3, 52, 203, 204, 12, 50, 2, 2, 204,
	205, 7, 149, 2, 2, 205, 363, 5, 4, 3, 51, 206, 208, 12, 49, 2, 2, 207,
	209, 7, 178, 2, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210,
	3, 2, 2, 2, 210, 363, 5, 4, 3, 50, 211, 212, 12, 48, 2, 2, 212, 213, 7,
	150, 2, 2, 213, 363, 5, 4, 3, 49, 214, 215, 12, 47, 2, 2, 215, 216, 7,
	151, 2, 2, 216, 363, 5, 4, 3, 48, 217, 218, 12, 46, 2, 2, 218, 219, 7,
	152, 2, 2, 219, 363, 5, 4, 3, 47, 220, 221, 12, 45, 2, 2, 221, 222, 7,
	153, 2, 2, 222, 363, 5, 4, 3, 46, 223, 224, 12, 44, 2, 2, 224, 225, 7,
	154, 2, 2, 225, 363, 5, 4, 3, 45, 226, 227, 12, 43, 2, 2, 227, 228, 7,
	155, 2, 2, 228, 363, 5, 4, 3, 44, 229, 230, 12, 42, 2, 2, 230, 231, 7,
	156, 2, 2, 231, 363, 5, 4, 3, 43, 232, 233, 12, 41, 2, 2, 233, 234, 9,
	9, 2, 2, 234, 363, 5, 4, 3, 42, 235, 236, 12, 40, 2, 2, 236, 237, 7, 130,
	2, 2, 237, 363, 5, 4, 3, 41, 238, 239, 12, 39, 2, 2, 239, 240, 7, 131,
	2, 2, 240, 363, 5, 4, 3, 40, 241, 242, 12, 36, 2, 2, 242, 243, 9, 10, 2,
	2, 243, 363, 5, 4, 3, 37, 244, 245, 12, 35, 2, 2, 245, 246, 9, 11, 2, 2,
	246, 363, 5, 4, 3, 36, 247, 248, 12, 34, 2, 2, 248, 249, 9, 12, 2, 2, 249,
	363, 5, 4, 3, 35, 250, 251, 12, 33, 2, 2, 251, 252, 9, 13, 2, 2, 252, 363,
	5, 4, 3, 34, 253, 254, 12, 31, 2, 2, 254, 255, 9, 14, 2, 2, 255, 363, 5,
	4, 3, 32, 256, 257, 12, 30, 2, 2, 257, 258, 9, 15, 2, 2, 258, 363, 5, 4,
	3, 31, 259, 260, 12, 29, 2, 2, 260, 261, 9, 16, 2, 2, 261, 363, 5, 4, 3,
	30, 262, 263, 12, 28, 2, 2, 263, 264, 7, 64, 2, 2, 264, 363, 5, 4, 3, 29,
	265, 266, 12, 27, 2, 2, 266, 267, 9, 17, 2, 2, 267, 363, 5, 4, 3, 27, 268,
	269, 12, 26, 2, 2, 269, 270, 9, 18, 2, 2, 270, 363, 5, 4, 3, 26, 271, 272,
	12, 25, 2, 2, 272, 273, 9, 19, 2, 2, 273, 363, 5, 4, 3, 26, 274, 275, 12,
	24, 2, 2, 275, 276, 7, 73, 2, 2, 276, 363, 5, 4, 3, 24, 277, 278, 12, 22,
	2, 2, 278, 279, 7, 25, 2, 2, 279, 363, 5, 4, 3, 23, 280, 281, 12, 20, 2,
	2, 281, 282, 7, 86, 2, 2, 282, 363, 5, 4, 3, 21, 283, 284, 12, 19, 2, 2,
	284, 285, 7, 87, 2, 2, 285, 363, 5, 4, 3, 20, 286, 287, 12, 18, 2, 2, 287,
	288, 7, 88, 2, 2, 288, 363, 5, 4, 3, 19, 289, 290, 12, 17, 2, 2, 290, 291,
	9, 20, 2, 2, 291, 363, 5, 4, 3, 17, 292, 293, 12, 16, 2, 2, 293, 294, 9,
	21, 2, 2, 294, 363, 5, 4, 3, 17, 295, 296, 12, 15, 2, 2, 296, 297, 9, 22,
	2, 2, 297, 363, 5, 4, 3, 15, 298, 299, 12, 13, 2, 2, 299, 300, 7, 157,
	2, 2, 300, 363, 5, 4, 3, 14, 301, 302, 12, 12, 2, 2, 302, 303, 7, 158,
	2, 2, 303, 363, 5, 4, 3, 13, 304, 305, 12, 11, 2, 2, 305, 306, 7, 159,
	2, 2, 306, 363, 5, 4, 3, 12, 307, 308, 12, 10, 2, 2, 308, 309, 7, 74, 2,
	2, 309, 363, 5, 4, 3, 10, 310, 311, 12, 9, 2, 2, 311, 312, 7, 75, 2, 2,
	312, 363, 5, 4, 3, 10, 313, 314, 12, 7, 2, 2, 314, 315, 9, 23, 2, 2, 315,
	363, 5, 4, 3, 7, 316, 317, 12, 89, 2, 2, 317, 318, 7, 85, 2, 2, 318, 321,
	7, 7, 2, 2, 319, 320, 7, 85, 2, 2, 320, 322, 7, 7, 2, 2, 321, 319, 3, 2,
	2, 2, 321, 322, 3, 2, 2, 2, 322, 363, 3, 2, 2, 2, 323, 324, 12, 80, 2,
	2, 324, 325, 7, 13, 2, 2, 325, 326, 5, 20, 11, 2, 326, 327, 7, 14, 2, 2,
	327, 363, 3, 2, 2, 2, 328, 329, 12, 79, 2, 2, 329, 363, 5, 22, 12, 2, 330,
	331, 12, 77, 2, 2, 331, 363, 9, 3, 2, 2, 332, 333, 12, 70, 2, 2, 333, 363,
	9, 24, 2, 2, 334, 335, 12, 69, 2, 2, 335, 363, 9, 25, 2, 2, 336, 338, 12,
	68, 2, 2, 337, 339, 7, 34, 2, 2, 338, 337, 3, 2, 2, 2, 339, 340, 3, 2,
	2, 2, 340, 338, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 363, 3, 2, 2, 2,
	342, 343, 12, 38, 2, 2, 343, 345, 7, 132, 2, 2, 344, 346, 5, 4, 3, 2, 345,
	344, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 363, 3, 2, 2, 2, 347, 348,
	12, 23, 2, 2, 348, 363, 9, 26, 2, 2, 349, 350, 12, 14, 2, 2, 350, 363,
	7, 104, 2, 2, 351, 352, 12, 8, 2, 2, 352, 353, 7, 42, 2, 2, 353, 363, 7,
	110, 2, 2, 354, 355, 12, 4, 2, 2, 355, 356, 9, 27, 2, 2, 356, 363, 7, 7,
	2, 2, 357, 358, 12, 3, 2, 2, 358, 360, 7, 133, 2, 2, 359, 361, 5, 4, 3,
	2, 360, 359, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 363, 3, 2, 2, 2, 362,
	143, 3, 2, 2, 2, 362, 146, 3, 2, 2, 2, 362, 149, 3, 2, 2, 2, 362, 152,
	3, 2, 2, 2, 362, 155, 3, 2, 2, 2, 362, 161, 3, 2, 2, 2, 362, 164, 3, 2,
	2, 2, 362, 167, 3, 2, 2, 2, 362, 170, 3, 2, 2, 2, 362, 173, 3, 2, 2, 2,
	362, 176, 3, 2, 2, 2, 362, 179, 3, 2, 2, 2, 362, 182, 3, 2, 2, 2, 362,
	185, 3, 2, 2, 2, 362, 188, 3, 2, 2, 2, 362, 191, 3, 2, 2, 2, 362, 194,
	3, 2, 2, 2, 362, 197, 3, 2, 2, 2, 362, 200, 3, 2, 2, 2, 362, 203, 3, 2,
	2, 2, 362, 206, 3, 2, 2, 2, 362, 211, 3, 2, 2, 2, 362, 214, 3, 2, 2, 2,
	362, 217, 3, 2, 2, 2, 362, 220, 3, 2, 2, 2, 362, 223, 3, 2, 2, 2, 362,
	226, 3, 2, 2, 2, 362, 229, 3, 2, 2, 2, 362, 232, 3, 2, 2, 2, 362, 235,
	3, 2, 2, 2, 362, 238, 3, 2, 2, 2, 362, 241, 3, 2, 2, 2, 362, 244, 3, 2,
	2, 2, 362, 247, 3, 2, 2, 2, 362, 250, 3, 2, 2, 2, 362, 253, 3, 2, 2, 2,
	362, 256, 3, 2, 2, 2, 362, 259, 3, 2, 2, 2, 362, 262, 3, 2, 2, 2, 362,
	265, 3, 2, 2, 2, 362, 268, 3, 2, 2, 2, 362, 271, 3, 2, 2, 2, 362, 274,
	3, 2, 2, 2, 362, 277, 3, 2, 2, 2, 362, 280, 3, 2, 2, 2, 362, 283, 3, 2,
	2, 2, 362, 286, 3, 2, 2, 2, 362, 289, 3, 2, 2, 2, 362, 292, 3, 2, 2, 2,
	362, 295, 3, 2, 2, 2, 362, 298, 3, 2, 2, 2, 362, 301, 3, 2, 2, 2, 362,
	304, 3, 2, 2, 2, 362, 307, 3, 2, 2, 2, 362, 310, 3, 2, 2, 2, 362, 313,
	3, 2, 2, 2, 362, 316, 3, 2, 2, 2, 362, 323, 3, 2, 2, 2, 362, 328, 3, 2,
	2, 2, 362, 330, 3, 2, 2, 2, 362, 332, 3, 2, 2, 2, 362, 334, 3, 2, 2, 2,
	362, 336, 3, 2, 2, 2, 362, 342, 3, 2, 2, 2, 362, 347, 3, 2, 2, 2, 362,
	349, 3, 2, 2, 2, 362, 351, 3, 2, 2, 2, 362, 354, 3, 2, 2, 2, 362, 357,
	3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2,
	2, 2, 365, 5, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 367, 369, 5, 8, 5, 2, 368,
	367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371,
	7, 3, 2, 2, 371, 7, 3, 2, 2, 2, 372, 374, 7, 33, 2, 2, 373, 372, 3, 2,
	2, 2, 373, 374, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376, 7, 3, 2, 2,
	376, 382, 7, 33, 2, 2, 377, 378, 7, 3, 2, 2, 378, 379, 7, 33, 2, 2, 379,
	380, 7, 3, 2, 2, 380, 382, 7, 33, 2, 2, 381, 373, 3, 2, 2, 2, 381, 377,
	3, 2, 2, 2, 382, 9, 3, 2, 2, 2, 383, 384, 7, 6, 2, 2, 384, 386, 7, 5, 2,
	2, 385, 387, 5, 12, 7, 2, 386, 385, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387,
	389, 3, 2, 2, 2, 388, 390, 5, 14, 8, 2, 389, 388, 3, 2, 2, 2, 389, 390,
	3, 2, 2, 2, 390, 399, 3, 2, 2, 2, 391, 393, 9, 28, 2, 2, 392, 394, 5, 12,
	7, 2, 393, 392, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2, 2,
	395, 397, 5, 14, 8, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397,
	399, 3, 2, 2, 2, 398, 383, 3, 2, 2, 2, 398, 391, 3, 2, 2, 2, 399, 11, 3,
	2, 2, 2, 400, 401, 7, 32, 2, 2, 401, 407, 9, 28, 2, 2, 402, 404, 7, 33,
	2, 2, 403, 405, 9, 28, 2, 2, 404, 403, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2,
	405, 407, 3, 2, 2, 2, 406, 400, 3, 2, 2, 2, 406, 402, 3, 2, 2, 2, 407,
	13, 3, 2, 2, 2, 408, 410, 7, 113, 2, 2, 409, 411, 9, 29, 2, 2, 410, 409,
	3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 7, 6,
	2, 2, 413, 15, 3, 2, 2, 2, 414, 417, 7, 83, 2, 2, 415, 417, 7, 84, 2, 2,
	416, 414, 3, 2, 2, 2, 416, 415, 3, 2, 2, 2, 417, 17, 3, 2, 2, 2, 418, 424,
	7, 95, 2, 2, 419, 424, 7, 96, 2, 2, 420, 424, 7, 97, 2, 2, 421, 424, 7,
	98, 2, 2, 422, 424, 7, 99, 2, 2, 423, 418, 3, 2, 2, 2, 423, 419, 3, 2,
	2, 2, 423, 420, 3, 2, 2, 2, 423, 421, 3, 2, 2, 2, 423, 422, 3, 2, 2, 2,
	424, 19, 3, 2, 2, 2, 425, 427, 5, 4, 3, 2, 426, 425, 3, 2, 2, 2, 426, 427,
	3, 2, 2, 2, 427, 434, 3, 2, 2, 2, 428, 430, 7, 15, 2, 2, 429, 431, 5, 4,
	3, 2, 430, 429, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 433, 3, 2, 2, 2,
	432, 428, 3, 2, 2, 2, 433, 436, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 434,
	435, 3, 2, 2, 2, 435, 21, 3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 437, 438, 7,
	13, 2, 2, 438, 439, 7, 13, 2, 2, 439, 440, 5, 20, 11, 2, 440, 441, 7, 14,
	2, 2, 441, 442, 7, 14, 2, 2, 442, 448, 3, 2, 2, 2, 443, 444, 7, 26, 2,
	2, 444, 445, 5, 20, 11, 2, 445, 446, 7, 27, 2, 2, 446, 448, 3, 2, 2, 2,
	447, 437, 3, 2, 2, 2, 447, 443, 3, 2, 2, 2, 448, 23, 3, 2, 2, 2, 449, 450,
	8, 13, 1, 2, 450, 476, 5, 4, 3, 2, 451, 455, 7, 167, 2, 2, 452, 454, 5,
	24, 13, 2, 453, 452, 3, 2, 2, 2, 454, 457, 3, 2, 2, 2, 455, 453, 3, 2,
	2, 2, 455, 456, 3, 2, 2, 2, 456, 458, 3, 2, 2, 2, 457, 455, 3, 2, 2, 2,
	458, 476, 7, 170, 2, 2, 459, 460, 7, 165, 2, 2, 460, 462, 7, 167, 2, 2,
	461, 463, 5, 24, 13, 2, 462, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464,
	462, 3, 2, 2, 2, 464, 465, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 467,
	7, 170, 2, 2, 467, 476, 3, 2, 2, 2, 468, 469, 7, 171, 2, 2, 469, 470, 5,
	24, 13, 2, 470, 471, 7, 168, 2, 2, 471, 472, 5, 24, 13, 5, 472, 476, 3,
	2, 2, 2, 473, 474, 7, 171, 2, 2, 474, 476, 5, 24, 13, 4, 475, 449, 3, 2,
	2, 2, 475, 451, 3, 2, 2, 2, 475, 459, 3, 2, 2, 2, 475, 468, 3, 2, 2, 2,
	475, 473, 3, 2, 2, 2, 476, 521, 3, 2, 2, 2, 477, 478, 12, 15, 2, 2, 478,
	479, 7, 174, 2, 2, 479, 480, 5, 24, 13, 2, 480, 481, 7, 168, 2, 2, 481,
	482, 5, 24, 13, 16, 482, 520, 3, 2, 2, 2, 483, 484, 12, 14, 2, 2, 484,
	485, 7, 169, 2, 2, 485, 486, 5, 24, 13, 2, 486, 487, 7, 168, 2, 2, 487,
	488, 5, 24, 13, 15, 488, 520, 3, 2, 2, 2, 489, 490, 12, 13, 2, 2, 490,
	491, 9, 30, 2, 2, 491, 520, 5, 24, 13, 13, 492, 493, 12, 12, 2, 2, 493,
	494, 7, 172, 2, 2, 494, 495, 5, 24, 13, 2, 495, 496, 7, 168, 2, 2, 496,
	497, 5, 24, 13, 13, 497, 520, 3, 2, 2, 2, 498, 499, 12, 11, 2, 2, 499,
	500, 7, 172, 2, 2, 500, 520, 5, 24, 13, 12, 501, 502, 12, 9, 2, 2, 502,
	503, 7, 169, 2, 2, 503, 520, 5, 24, 13, 10, 504, 505, 12, 8, 2, 2, 505,
	506, 7, 174, 2, 2, 506, 507, 5, 24, 13, 2, 507, 508, 7, 168, 2, 2, 508,
	509, 5, 24, 13, 9, 509, 520, 3, 2, 2, 2, 510, 511, 12, 7, 2, 2, 511, 512,
	7, 174, 2, 2, 512, 520, 5, 24, 13, 8, 513, 514, 12, 6, 2, 2, 514, 515,
	7, 166, 2, 2, 515, 520, 5, 24, 13, 7, 516, 517, 12, 3, 2, 2, 517, 518,
	7, 164, 2, 2, 518, 520, 5, 24, 13, 4, 519, 477, 3, 2, 2, 2, 519, 483, 3,
	2, 2, 2, 519, 489, 3, 2, 2, 2, 519, 492, 3, 2, 2, 2, 519, 498, 3, 2, 2,
	2, 519, 501, 3, 2, 2, 2, 519, 504, 3, 2, 2, 2, 519, 510, 3, 2, 2, 2, 519,
	513, 3, 2, 2, 2, 519, 516, 3, 2, 2, 2, 520, 523, 3, 2, 2, 2, 521, 519,
	3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522, 25, 3, 2, 2, 2, 523, 521, 3, 2,
	2, 2, 42, 30, 33, 37, 74, 78, 81, 91, 112, 116, 120, 141, 208, 321, 340,
	345, 360, 362, 364, 368, 373, 381, 386, 389, 393, 396, 398, 404, 406, 410,
	416, 423, 426, 430, 434, 447, 455, 464, 475, 519, 521,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "'('", "')'", "'{'", "'}'", "'['", "']'", "','",
	"'(*'", "'*)'", "'\u2329'", "'\u232A'", "'\u230A'", "'\u230B'", "'\u2308'",
	"'\u2309'", "'||'", "'|'", "'\u301A'", "'\u301B'", "'\uF603'", "'\uF604'",
	"'\uF605'", "'\uF606'", "'``'", "'`'", "'''", "'\"'", "'<<'", "'>>>'",
	"'>>'", "'==='", "'=!='", "", "'='", "", "", "", "'<'", "'>'", "'\uF3D0'",
	"'\uF3D1'", "'\u2225'", "'\u2226'", "'\u2208'", "'\u2209'", "'\u2282'",
	"'\u2283'", "'\u00AC'", "'&&'", "'\u2227'", "'\u22BC'", "'\u22BB'", "'\uF4A2'",
	"'\u2228'", "'\u22BD'", "'\u29E6'", "'\uF523'", "'\u2970'", "'\u22A2'",
	"'\u22A8'", "'\u22A3'", "'\u2AE4'", "'\u22A5'", "'\u22A4'", "'\u220D'",
	"'\u2234'", "'\u2235'", "'...'", "'..'", "'?'", "'___'", "'__'", "'_.'",
	"'_'", "", "", "'::'", "':'", "'~~'", "'/;'", "'->'", "'\uF522'", "':>'",
	"'\uF51F'", "'/.'", "'//.'", "", "", "", "'##'", "'#'", "'+='", "'-='",
	"'*='", "'/='", "'&'", "'^:='", "':='", "'^='", "'/:'", "'\uF4A1'", "'.'",
	"'^^'", "'^'", "'*^'", "'++'", "'--'", "'@@@'", "'@@'", "'@*'", "'@'",
	"'/@'", "'/*'", "'//@'", "'!!'", "'!'", "'<>'", "'\u222B'", "'\uF74C'",
	"'\uF4A0'", "'\\'", "'\u22C2'", "'\u22C3'", "';;'", "';'", "'\uF3C7'",
	"'\uF3C9'", "'\uF3CE'", "'\uF3C8'", "'~'", "'\u2207'", "'\uF520'", "'\u2218'",
	"'\u2299'", "'**'", "'\u2216'", "'\u22C4'", "'\u22C0'", "'\u22C1'", "'\u2297'",
	"'\u00B7'", "'\u22C6'", "'\u2240'", "'\u2210'", "'\u2322'", "'\u2323'",
	"'\u2295'", "'\u2296'", "'\u2236'", "'//'", "'\uF432'", "'+'", "'-'", "'\u00B1'",
	"'\u2213'", "'\\`'", "'\\!'", "'\\/'", "'\\('", "'\\%'", "'\\&'", "'\\)'",
	"'\\@'", "'\\_'", "'\\^'", "'\\+'", "'\\*'", "'/'", "'\u00F7'", "", "'*'",
	"'\u00D7'", "'\n'", "'\uF3B1'",
}
var symbolicNames = []string{
	"", "Name", "DecimalNumber", "NumberInBase", "DIGITS", "StringLiteral",
	"COMMENT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACKET", "RBRACKET",
	"COMMA", "LCOMMENT", "RCOMMENT", "LANGLE", "RANGLE", "LFLOOR", "RFLOOR",
	"LCEILING", "RCEILING", "DOUBLEBAR", "BAR", "LBARBRACKET", "RBARBRACKET",
	"LBRACKETINGBAR", "RBRACKETINGBAR", "LDOUBLEBRACKETINGBAR", "RDOUBLEBRACKETINGBAR",
	"DOUBLEBACKQUOTE", "BACKQUOTE", "SINGLEQUOTE", "QUOTE", "DOUBLELESS", "TRIPPLEGREATER",
	"DOUBLEGREATER", "TRIPPLEEQUAL", "EQUALBANGEQUAL", "EqualSymbol", "EQUAL",
	"NotEqualSymbol", "GreaterEqualSymbol", "LessEqualSymbol", "LESS", "GREATER",
	"VERTICALBAR", "NOTVERTICALBAR", "DOUBLEVERTICALBAR", "NOTDOUBLEVERTICALBAR",
	"ELEMENT", "NOTELEMENT", "SUBSET", "SUPERSET", "NOT", "DOUBLEAMP", "AND",
	"NAND", "XOR", "XNOR", "OR", "NOR", "LRDOUBLEARROW", "RDOUBLEARROW", "LCONTAINS",
	"RIGHTTEE", "DOUBLERIGHTTEE", "LEFTTEE", "DOUBLELEFTTEE", "UPTEE", "DOWNTEE",
	"SUCHTHAT", "THEREFORE", "BECAUSE", "TRIPPLEDOT", "DOUBLEDOT", "QUESTIONMARK",
	"TRIPPLEBLANK", "DOUBLEBLANK", "BLANKDOT", "BLANK", "PERCENTDIGITS", "PERCENTS",
	"DOUBLECOLON", "RAWCOLON", "DOUBLETILDE", "SLASHSEMI", "MINUSGREATER",
	"RARROW", "COLONGREATER", "COLONARROW", "SLASHDOT", "DOUBLESLASHDOT", "HASHDIGITS",
	"HASHStringLiteral", "DOUBLEHASHDIGITS", "DOUBLEHASH", "HASH", "PLUSEQUAL",
	"MINUSEQUAL", "ASTERISKEQUAL", "SLASHEQUAL", "AMP", "CARETCOLONEQUAL",
	"COLONEQUAL", "CARETEQUAL", "SLASHCOLON", "FUNCTIONARROW", "DOT", "DOUBLECARET",
	"CARET", "ASTERISKCARET", "DOUBLEPLUS", "DOUBLEMINUS", "TRIPPLEAT", "DOUBLEAT",
	"ATASTERISK", "AT", "MAP", "SLASHASTERISK", "MAPALL", "DOUBLEBANG", "BANG",
	"LESSGREATER", "INTEGRAL", "DIFFERENTIALD", "CROSS", "RAWBACKSLASH", "INTERSECTION",
	"UNION", "DOUBLESEMICOLON", "SEMICOLON", "TRANSPOSE", "CONJUGATETRANSPOSE",
	"HERMITIANCONJUGATE", "CONJUGATE", "TILDE", "DEL", "SQUARE", "SMALLCIRCLE",
	"CIRCLEDOT", "DOUBLEASTERISK", "BACKSLASH", "DIAMOND", "WEDGE", "VEE",
	"CIRCLETIMES", "CENTERDOT", "STAR", "VERTICALTILDE", "COPRODUCT", "CAP",
	"CUP", "CIRCLEPLUS", "CIRCLEMINUS", "COLON", "DOUBLESLASH", "VERTICALSEPARATOR",
	"PLUS", "MINUS", "PLUSMINUS", "MINUSPLUS", "FormBox", "InterpretedBox",
	"BoxFraction", "BoxLeftBoxParenthesis", "BoxOtherscript", "BoxOverscript",
	"BoxRightBoxParenthesis", "BoxSqrt", "BoxSubscript", "BoxSuperscript",
	"BoxUnderscript", "BoxConstructor", "SLASH", "DIVIDE", "MultiplicationSymbol",
	"ASTERISK", "TIMES", "NEWLINE", "CONTINUATION", "WHITESPACE", "BINARYPLUS",
	"BINARYMINUS", "BINARYMINUSPLUS", "BINARYPLUSMINUS", "SPANSEMICOLONS",
}

var ruleNames = []string{
	"prog", "expr", "symbol", "context", "numberLiteral", "numberLiteralPrecision",
	"numberLiteralExponent", "outExpression", "slotExpression", "expressionList",
	"accessExpression", "box",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FoxySheepParser struct {
	*antlr.BaseParser
}

func NewFoxySheepParser(input antlr.TokenStream) *FoxySheepParser {
	this := new(FoxySheepParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FoxySheep.g4"

	return this
}

// FoxySheepParser tokens.
const (
	FoxySheepParserEOF                    = antlr.TokenEOF
	FoxySheepParserName                   = 1
	FoxySheepParserDecimalNumber          = 2
	FoxySheepParserNumberInBase           = 3
	FoxySheepParserDIGITS                 = 4
	FoxySheepParserStringLiteral          = 5
	FoxySheepParserCOMMENT                = 6
	FoxySheepParserLPAREN                 = 7
	FoxySheepParserRPAREN                 = 8
	FoxySheepParserLBRACE                 = 9
	FoxySheepParserRBRACE                 = 10
	FoxySheepParserLBRACKET               = 11
	FoxySheepParserRBRACKET               = 12
	FoxySheepParserCOMMA                  = 13
	FoxySheepParserLCOMMENT               = 14
	FoxySheepParserRCOMMENT               = 15
	FoxySheepParserLANGLE                 = 16
	FoxySheepParserRANGLE                 = 17
	FoxySheepParserLFLOOR                 = 18
	FoxySheepParserRFLOOR                 = 19
	FoxySheepParserLCEILING               = 20
	FoxySheepParserRCEILING               = 21
	FoxySheepParserDOUBLEBAR              = 22
	FoxySheepParserBAR                    = 23
	FoxySheepParserLBARBRACKET            = 24
	FoxySheepParserRBARBRACKET            = 25
	FoxySheepParserLBRACKETINGBAR         = 26
	FoxySheepParserRBRACKETINGBAR         = 27
	FoxySheepParserLDOUBLEBRACKETINGBAR   = 28
	FoxySheepParserRDOUBLEBRACKETINGBAR   = 29
	FoxySheepParserDOUBLEBACKQUOTE        = 30
	FoxySheepParserBACKQUOTE              = 31
	FoxySheepParserSINGLEQUOTE            = 32
	FoxySheepParserQUOTE                  = 33
	FoxySheepParserDOUBLELESS             = 34
	FoxySheepParserTRIPPLEGREATER         = 35
	FoxySheepParserDOUBLEGREATER          = 36
	FoxySheepParserTRIPPLEEQUAL           = 37
	FoxySheepParserEQUALBANGEQUAL         = 38
	FoxySheepParserEqualSymbol            = 39
	FoxySheepParserEQUAL                  = 40
	FoxySheepParserNotEqualSymbol         = 41
	FoxySheepParserGreaterEqualSymbol     = 42
	FoxySheepParserLessEqualSymbol        = 43
	FoxySheepParserLESS                   = 44
	FoxySheepParserGREATER                = 45
	FoxySheepParserVERTICALBAR            = 46
	FoxySheepParserNOTVERTICALBAR         = 47
	FoxySheepParserDOUBLEVERTICALBAR      = 48
	FoxySheepParserNOTDOUBLEVERTICALBAR   = 49
	FoxySheepParserELEMENT                = 50
	FoxySheepParserNOTELEMENT             = 51
	FoxySheepParserSUBSET                 = 52
	FoxySheepParserSUPERSET               = 53
	FoxySheepParserNOT                    = 54
	FoxySheepParserDOUBLEAMP              = 55
	FoxySheepParserAND                    = 56
	FoxySheepParserNAND                   = 57
	FoxySheepParserXOR                    = 58
	FoxySheepParserXNOR                   = 59
	FoxySheepParserOR                     = 60
	FoxySheepParserNOR                    = 61
	FoxySheepParserLRDOUBLEARROW          = 62
	FoxySheepParserRDOUBLEARROW           = 63
	FoxySheepParserLCONTAINS              = 64
	FoxySheepParserRIGHTTEE               = 65
	FoxySheepParserDOUBLERIGHTTEE         = 66
	FoxySheepParserLEFTTEE                = 67
	FoxySheepParserDOUBLELEFTTEE          = 68
	FoxySheepParserUPTEE                  = 69
	FoxySheepParserDOWNTEE                = 70
	FoxySheepParserSUCHTHAT               = 71
	FoxySheepParserTHEREFORE              = 72
	FoxySheepParserBECAUSE                = 73
	FoxySheepParserTRIPPLEDOT             = 74
	FoxySheepParserDOUBLEDOT              = 75
	FoxySheepParserQUESTIONMARK           = 76
	FoxySheepParserTRIPPLEBLANK           = 77
	FoxySheepParserDOUBLEBLANK            = 78
	FoxySheepParserBLANKDOT               = 79
	FoxySheepParserBLANK                  = 80
	FoxySheepParserPERCENTDIGITS          = 81
	FoxySheepParserPERCENTS               = 82
	FoxySheepParserDOUBLECOLON            = 83
	FoxySheepParserRAWCOLON               = 84
	FoxySheepParserDOUBLETILDE            = 85
	FoxySheepParserSLASHSEMI              = 86
	FoxySheepParserMINUSGREATER           = 87
	FoxySheepParserRARROW                 = 88
	FoxySheepParserCOLONGREATER           = 89
	FoxySheepParserCOLONARROW             = 90
	FoxySheepParserSLASHDOT               = 91
	FoxySheepParserDOUBLESLASHDOT         = 92
	FoxySheepParserHASHDIGITS             = 93
	FoxySheepParserHASHStringLiteral      = 94
	FoxySheepParserDOUBLEHASHDIGITS       = 95
	FoxySheepParserDOUBLEHASH             = 96
	FoxySheepParserHASH                   = 97
	FoxySheepParserPLUSEQUAL              = 98
	FoxySheepParserMINUSEQUAL             = 99
	FoxySheepParserASTERISKEQUAL          = 100
	FoxySheepParserSLASHEQUAL             = 101
	FoxySheepParserAMP                    = 102
	FoxySheepParserCARETCOLONEQUAL        = 103
	FoxySheepParserCOLONEQUAL             = 104
	FoxySheepParserCARETEQUAL             = 105
	FoxySheepParserSLASHCOLON             = 106
	FoxySheepParserFUNCTIONARROW          = 107
	FoxySheepParserDOT                    = 108
	FoxySheepParserDOUBLECARET            = 109
	FoxySheepParserCARET                  = 110
	FoxySheepParserASTERISKCARET          = 111
	FoxySheepParserDOUBLEPLUS             = 112
	FoxySheepParserDOUBLEMINUS            = 113
	FoxySheepParserTRIPPLEAT              = 114
	FoxySheepParserDOUBLEAT               = 115
	FoxySheepParserATASTERISK             = 116
	FoxySheepParserAT                     = 117
	FoxySheepParserMAP                    = 118
	FoxySheepParserSLASHASTERISK          = 119
	FoxySheepParserMAPALL                 = 120
	FoxySheepParserDOUBLEBANG             = 121
	FoxySheepParserBANG                   = 122
	FoxySheepParserLESSGREATER            = 123
	FoxySheepParserINTEGRAL               = 124
	FoxySheepParserDIFFERENTIALD          = 125
	FoxySheepParserCROSS                  = 126
	FoxySheepParserRAWBACKSLASH           = 127
	FoxySheepParserINTERSECTION           = 128
	FoxySheepParserUNION                  = 129
	FoxySheepParserDOUBLESEMICOLON        = 130
	FoxySheepParserSEMICOLON              = 131
	FoxySheepParserTRANSPOSE              = 132
	FoxySheepParserCONJUGATETRANSPOSE     = 133
	FoxySheepParserHERMITIANCONJUGATE     = 134
	FoxySheepParserCONJUGATE              = 135
	FoxySheepParserTILDE                  = 136
	FoxySheepParserDEL                    = 137
	FoxySheepParserSQUARE                 = 138
	FoxySheepParserSMALLCIRCLE            = 139
	FoxySheepParserCIRCLEDOT              = 140
	FoxySheepParserDOUBLEASTERISK         = 141
	FoxySheepParserBACKSLASH              = 142
	FoxySheepParserDIAMOND                = 143
	FoxySheepParserWEDGE                  = 144
	FoxySheepParserVEE                    = 145
	FoxySheepParserCIRCLETIMES            = 146
	FoxySheepParserCENTERDOT              = 147
	FoxySheepParserSTAR                   = 148
	FoxySheepParserVERTICALTILDE          = 149
	FoxySheepParserCOPRODUCT              = 150
	FoxySheepParserCAP                    = 151
	FoxySheepParserCUP                    = 152
	FoxySheepParserCIRCLEPLUS             = 153
	FoxySheepParserCIRCLEMINUS            = 154
	FoxySheepParserCOLON                  = 155
	FoxySheepParserDOUBLESLASH            = 156
	FoxySheepParserVERTICALSEPARATOR      = 157
	FoxySheepParserPLUS                   = 158
	FoxySheepParserMINUS                  = 159
	FoxySheepParserPLUSMINUS              = 160
	FoxySheepParserMINUSPLUS              = 161
	FoxySheepParserFormBox                = 162
	FoxySheepParserInterpretedBox         = 163
	FoxySheepParserBoxFraction            = 164
	FoxySheepParserBoxLeftBoxParenthesis  = 165
	FoxySheepParserBoxOtherscript         = 166
	FoxySheepParserBoxOverscript          = 167
	FoxySheepParserBoxRightBoxParenthesis = 168
	FoxySheepParserBoxSqrt                = 169
	FoxySheepParserBoxSubscript           = 170
	FoxySheepParserBoxSuperscript         = 171
	FoxySheepParserBoxUnderscript         = 172
	FoxySheepParserBoxConstructor         = 173
	FoxySheepParserSLASH                  = 174
	FoxySheepParserDIVIDE                 = 175
	FoxySheepParserMultiplicationSymbol   = 176
	FoxySheepParserASTERISK               = 177
	FoxySheepParserTIMES                  = 178
	FoxySheepParserNEWLINE                = 179
	FoxySheepParserCONTINUATION           = 180
	FoxySheepParserWHITESPACE             = 181
	FoxySheepParserBINARYPLUS             = 182
	FoxySheepParserBINARYMINUS            = 183
	FoxySheepParserBINARYMINUSPLUS        = 184
	FoxySheepParserBINARYPLUSMINUS        = 185
	FoxySheepParserSPANSEMICOLONS         = 186
)

// FoxySheepParser rules.
const (
	FoxySheepParserRULE_prog                   = 0
	FoxySheepParserRULE_expr                   = 1
	FoxySheepParserRULE_symbol                 = 2
	FoxySheepParserRULE_context                = 3
	FoxySheepParserRULE_numberLiteral          = 4
	FoxySheepParserRULE_numberLiteralPrecision = 5
	FoxySheepParserRULE_numberLiteralExponent  = 6
	FoxySheepParserRULE_outExpression          = 7
	FoxySheepParserRULE_slotExpression         = 8
	FoxySheepParserRULE_expressionList         = 9
	FoxySheepParserRULE_accessExpression       = 10
	FoxySheepParserRULE_box                    = 11
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ProgContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserNEWLINE)
}

func (s *ProgContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNEWLINE, i)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FoxySheepParserRULE_prog)
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
		p.SetState(24)
		p.expr(0)
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FoxySheepParserNEWLINE {
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(25)
					p.Match(FoxySheepParserNEWLINE)
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(28)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
			{
				p.SetState(30)
				p.expr(0)
			}

		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = FoxySheepParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PatternExpContext struct {
	*ExprContext
}

func NewPatternExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternExpContext {
	var p = new(PatternExpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PatternExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternExpContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *PatternExpContext) RAWCOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRAWCOLON, 0)
}

func (s *PatternExpContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PatternExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPatternExp(s)
	}
}

func (s *PatternExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPatternExp(s)
	}
}

func (s *PatternExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPatternExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrContext struct {
	*ExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserOR, 0)
}

func (s *OrContext) NOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNOR, 0)
}

func (s *OrContext) DOUBLEBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEBAR, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConjugateContext struct {
	*ExprContext
}

func NewConjugateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConjugateContext {
	var p = new(ConjugateContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConjugateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConjugateContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConjugateContext) CONJUGATE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCONJUGATE, 0)
}

func (s *ConjugateContext) TRANSPOSE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRANSPOSE, 0)
}

func (s *ConjugateContext) CONJUGATETRANSPOSE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCONJUGATETRANSPOSE, 0)
}

func (s *ConjugateContext) HERMITIANCONJUGATE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserHERMITIANCONJUGATE, 0)
}

func (s *ConjugateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterConjugate(s)
	}
}

func (s *ConjugateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitConjugate(s)
	}
}

func (s *ConjugateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitConjugate(s)

	default:
		return t.VisitChildren(s)
	}
}

type CeilingContext struct {
	*ExprContext
}

func NewCeilingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilingContext {
	var p = new(CeilingContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CeilingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilingContext) LCEILING() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLCEILING, 0)
}

func (s *CeilingContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CeilingContext) RCEILING() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRCEILING, 0)
}

func (s *CeilingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCeiling(s)
	}
}

func (s *CeilingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCeiling(s)
	}
}

func (s *CeilingContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCeiling(s)

	default:
		return t.VisitChildren(s)
	}
}

type InfixContext struct {
	*ExprContext
}

func NewInfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InfixContext {
	var p = new(InfixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InfixContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InfixContext) AllTILDE() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserTILDE)
}

func (s *InfixContext) TILDE(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTILDE, i)
}

func (s *InfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterInfix(s)
	}
}

func (s *InfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitInfix(s)
	}
}

func (s *InfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitInfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type ThereforeContext struct {
	*ExprContext
}

func NewThereforeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThereforeContext {
	var p = new(ThereforeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ThereforeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThereforeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ThereforeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ThereforeContext) THEREFORE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTHEREFORE, 0)
}

func (s *ThereforeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterTherefore(s)
	}
}

func (s *ThereforeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitTherefore(s)
	}
}

func (s *ThereforeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitTherefore(s)

	default:
		return t.VisitChildren(s)
	}
}

type TagUnsetContext struct {
	*ExprContext
}

func NewTagUnsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagUnsetContext {
	var p = new(TagUnsetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TagUnsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagUnsetContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *TagUnsetContext) SLASHCOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHCOLON, 0)
}

func (s *TagUnsetContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TagUnsetContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEQUAL, 0)
}

func (s *TagUnsetContext) DOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOT, 0)
}

func (s *TagUnsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterTagUnset(s)
	}
}

func (s *TagUnsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitTagUnset(s)
	}
}

func (s *TagUnsetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitTagUnset(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccessorContext struct {
	*ExprContext
}

func NewAccessorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessorContext {
	var p = new(AccessorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AccessorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AccessorContext) AccessExpression() IAccessExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessExpressionContext)
}

func (s *AccessorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAccessor(s)
	}
}

func (s *AccessorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAccessor(s)
	}
}

func (s *AccessorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAccessor(s)

	default:
		return t.VisitChildren(s)
	}
}

type CircleMinusContext struct {
	*ExprContext
}

func NewCircleMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CircleMinusContext {
	var p = new(CircleMinusContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CircleMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircleMinusContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CircleMinusContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CircleMinusContext) CIRCLEMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCIRCLEMINUS, 0)
}

func (s *CircleMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCircleMinus(s)
	}
}

func (s *CircleMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCircleMinus(s)
	}
}

func (s *CircleMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCircleMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivideContext struct {
	*ExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivideContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivideContext) SLASH() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASH, 0)
}

func (s *DivideContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIVIDE, 0)
}

func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDivide(s)
	}
}

func (s *DivideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDivide(s)

	default:
		return t.VisitChildren(s)
	}
}

type ImpliesContext struct {
	*ExprContext
}

func NewImpliesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImpliesContext {
	var p = new(ImpliesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ImpliesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpliesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ImpliesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ImpliesContext) RDOUBLEARROW() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRDOUBLEARROW, 0)
}

func (s *ImpliesContext) LCONTAINS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLCONTAINS, 0)
}

func (s *ImpliesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterImplies(s)
	}
}

func (s *ImpliesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitImplies(s)
	}
}

func (s *ImpliesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitImplies(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusOpContext struct {
	*ExprContext
}

func NewPlusOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusOpContext {
	var p = new(PlusOpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PlusOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusOpContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PlusOpContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PlusOpContext) BINARYPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBINARYPLUS, 0)
}

func (s *PlusOpContext) BINARYMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBINARYMINUS, 0)
}

func (s *PlusOpContext) BINARYPLUSMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBINARYPLUSMINUS, 0)
}

func (s *PlusOpContext) BINARYMINUSPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBINARYMINUSPLUS, 0)
}

func (s *PlusOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPlusOp(s)
	}
}

func (s *PlusOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPlusOp(s)
	}
}

func (s *PlusOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPlusOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type RightCompositionContext struct {
	*ExprContext
}

func NewRightCompositionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightCompositionContext {
	var p = new(RightCompositionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RightCompositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightCompositionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RightCompositionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RightCompositionContext) SLASHASTERISK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHASTERISK, 0)
}

func (s *RightCompositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterRightComposition(s)
	}
}

func (s *RightCompositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitRightComposition(s)
	}
}

func (s *RightCompositionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitRightComposition(s)

	default:
		return t.VisitChildren(s)
	}
}

type NonCommutativeMultiplyContext struct {
	*ExprContext
}

func NewNonCommutativeMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NonCommutativeMultiplyContext {
	var p = new(NonCommutativeMultiplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NonCommutativeMultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonCommutativeMultiplyContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NonCommutativeMultiplyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NonCommutativeMultiplyContext) DOUBLEASTERISK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEASTERISK, 0)
}

func (s *NonCommutativeMultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNonCommutativeMultiply(s)
	}
}

func (s *NonCommutativeMultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNonCommutativeMultiply(s)
	}
}

func (s *NonCommutativeMultiplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNonCommutativeMultiply(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListContext struct {
	*ExprContext
}

func NewListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListContext {
	var p = new(ListContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLBRACE, 0)
}

func (s *ListContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *ListContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRBRACE, 0)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

type CupContext struct {
	*ExprContext
}

func NewCupContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CupContext {
	var p = new(CupContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CupContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CupContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CupContext) CUP() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCUP, 0)
}

func (s *CupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCup(s)
	}
}

func (s *CupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCup(s)
	}
}

func (s *CupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCup(s)

	default:
		return t.VisitChildren(s)
	}
}

type SameContext struct {
	*ExprContext
}

func NewSameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SameContext {
	var p = new(SameContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SameContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SameContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SameContext) TRIPPLEEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRIPPLEEQUAL, 0)
}

func (s *SameContext) EQUALBANGEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEQUALBANGEQUAL, 0)
}

func (s *SameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSame(s)
	}
}

func (s *SameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSame(s)
	}
}

func (s *SameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSame(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoxParenContext struct {
	*ExprContext
}

func NewBoxParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoxParenContext {
	var p = new(BoxParenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoxParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoxParenContext) BoxLeftBoxParenthesis() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxLeftBoxParenthesis, 0)
}

func (s *BoxParenContext) BoxRightBoxParenthesis() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxRightBoxParenthesis, 0)
}

func (s *BoxParenContext) AllBox() []IBoxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoxContext)(nil)).Elem())
	var tst = make([]IBoxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoxContext)
		}
	}

	return tst
}

func (s *BoxParenContext) Box(i int) IBoxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoxContext)
}

func (s *BoxParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBoxParen(s)
	}
}

func (s *BoxParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBoxParen(s)
	}
}

func (s *BoxParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBoxParen(s)

	default:
		return t.VisitChildren(s)
	}
}

type OptionalContext struct {
	*ExprContext
}

func NewOptionalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OptionalContext {
	var p = new(OptionalContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OptionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OptionalContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OptionalContext) RAWCOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRAWCOLON, 0)
}

func (s *OptionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOptional(s)
	}
}

func (s *OptionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOptional(s)
	}
}

func (s *OptionalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOptional(s)

	default:
		return t.VisitChildren(s)
	}
}

type SuchThatContext struct {
	*ExprContext
}

func NewSuchThatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SuchThatContext {
	var p = new(SuchThatContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SuchThatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SuchThatContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SuchThatContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SuchThatContext) SUCHTHAT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSUCHTHAT, 0)
}

func (s *SuchThatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSuchThat(s)
	}
}

func (s *SuchThatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSuchThat(s)
	}
}

func (s *SuchThatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSuchThat(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleBracketingBarContext struct {
	*ExprContext
}

func NewDoubleBracketingBarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleBracketingBarContext {
	var p = new(DoubleBracketingBarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DoubleBracketingBarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleBracketingBarContext) LDOUBLEBRACKETINGBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLDOUBLEBRACKETINGBAR, 0)
}

func (s *DoubleBracketingBarContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *DoubleBracketingBarContext) RDOUBLEBRACKETINGBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRDOUBLEBRACKETINGBAR, 0)
}

func (s *DoubleBracketingBarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDoubleBracketingBar(s)
	}
}

func (s *DoubleBracketingBarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDoubleBracketingBar(s)
	}
}

func (s *DoubleBracketingBarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDoubleBracketingBar(s)

	default:
		return t.VisitChildren(s)
	}
}

type PatternBlankDotContext struct {
	*ExprContext
}

func NewPatternBlankDotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternBlankDotContext {
	var p = new(PatternBlankDotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PatternBlankDotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternBlankDotContext) BLANKDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBLANKDOT, 0)
}

func (s *PatternBlankDotContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *PatternBlankDotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPatternBlankDot(s)
	}
}

func (s *PatternBlankDotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPatternBlankDot(s)
	}
}

func (s *PatternBlankDotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPatternBlankDot(s)

	default:
		return t.VisitChildren(s)
	}
}

type DotContext struct {
	*ExprContext
}

func NewDotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DotContext {
	var p = new(DotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DotContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DotContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DotContext) DOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOT, 0)
}

func (s *DotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDot(s)
	}
}

func (s *DotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDot(s)
	}
}

func (s *DotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDot(s)

	default:
		return t.VisitChildren(s)
	}
}

type VerticalBarContext struct {
	*ExprContext
}

func NewVerticalBarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VerticalBarContext {
	var p = new(VerticalBarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VerticalBarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalBarContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *VerticalBarContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VerticalBarContext) VERTICALBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserVERTICALBAR, 0)
}

func (s *VerticalBarContext) NOTVERTICALBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNOTVERTICALBAR, 0)
}

func (s *VerticalBarContext) DOUBLEVERTICALBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEVERTICALBAR, 0)
}

func (s *VerticalBarContext) NOTDOUBLEVERTICALBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNOTDOUBLEVERTICALBAR, 0)
}

func (s *VerticalBarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterVerticalBar(s)
	}
}

func (s *VerticalBarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitVerticalBar(s)
	}
}

func (s *VerticalBarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitVerticalBar(s)

	default:
		return t.VisitChildren(s)
	}
}

type SquareContext struct {
	*ExprContext
}

func NewSquareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SquareContext {
	var p = new(SquareContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SquareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SquareContext) SQUARE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSQUARE, 0)
}

func (s *SquareContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SquareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSquare(s)
	}
}

func (s *SquareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSquare(s)
	}
}

func (s *SquareContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSquare(s)

	default:
		return t.VisitChildren(s)
	}
}

type AlternativesContext struct {
	*ExprContext
}

func NewAlternativesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AlternativesContext {
	var p = new(AlternativesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AlternativesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlternativesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AlternativesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AlternativesContext) BAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBAR, 0)
}

func (s *AlternativesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAlternatives(s)
	}
}

func (s *AlternativesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAlternatives(s)
	}
}

func (s *AlternativesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAlternatives(s)

	default:
		return t.VisitChildren(s)
	}
}

type OutContext struct {
	*ExprContext
}

func NewOutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OutContext {
	var p = new(OutContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutContext) OutExpression() IOutExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOutExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOutExpressionContext)
}

func (s *OutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOut(s)
	}
}

func (s *OutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOut(s)
	}
}

func (s *OutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOut(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoxConstructorContext struct {
	*ExprContext
}

func NewBoxConstructorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoxConstructorContext {
	var p = new(BoxConstructorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoxConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoxConstructorContext) BoxConstructor() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxConstructor, 0)
}

func (s *BoxConstructorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BoxConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBoxConstructor(s)
	}
}

func (s *BoxConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBoxConstructor(s)
	}
}

func (s *BoxConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBoxConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotContext struct {
	*ExprContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotContext) BANG() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBANG, 0)
}

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNOT, 0)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNot(s)
	}
}

func (s *NotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNot(s)

	default:
		return t.VisitChildren(s)
	}
}

type PostfixContext struct {
	*ExprContext
}

func NewPostfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostfixContext {
	var p = new(PostfixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PostfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PostfixContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PostfixContext) DOUBLESLASH() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLESLASH, 0)
}

func (s *PostfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPostfix(s)
	}
}

func (s *PostfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPostfix(s)
	}
}

func (s *PostfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPostfix(s)

	default:
		return t.VisitChildren(s)
	}
}

type PatternBlanksContext struct {
	*ExprContext
}

func NewPatternBlanksContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternBlanksContext {
	var p = new(PatternBlanksContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PatternBlanksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternBlanksContext) TRIPPLEBLANK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRIPPLEBLANK, 0)
}

func (s *PatternBlanksContext) DOUBLEBLANK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEBLANK, 0)
}

func (s *PatternBlanksContext) BLANK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBLANK, 0)
}

func (s *PatternBlanksContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *PatternBlanksContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PatternBlanksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPatternBlanks(s)
	}
}

func (s *PatternBlanksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPatternBlanks(s)
	}
}

func (s *PatternBlanksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPatternBlanks(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryPlusMinusContext struct {
	*ExprContext
}

func NewUnaryPlusMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusMinusContext {
	var p = new(UnaryPlusMinusContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryPlusMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusMinusContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryPlusMinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMINUS, 0)
}

func (s *UnaryPlusMinusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPLUS, 0)
}

func (s *UnaryPlusMinusContext) PLUSMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPLUSMINUS, 0)
}

func (s *UnaryPlusMinusContext) MINUSPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMINUSPLUS, 0)
}

func (s *UnaryPlusMinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterUnaryPlusMinus(s)
	}
}

func (s *UnaryPlusMinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitUnaryPlusMinus(s)
	}
}

func (s *UnaryPlusMinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitUnaryPlusMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

type CapContext struct {
	*ExprContext
}

func NewCapContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CapContext {
	var p = new(CapContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CapContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CapContext) CAP() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCAP, 0)
}

func (s *CapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCap(s)
	}
}

func (s *CapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCap(s)
	}
}

func (s *CapContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCap(s)

	default:
		return t.VisitChildren(s)
	}
}

type CirclePlusContext struct {
	*ExprContext
}

func NewCirclePlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CirclePlusContext {
	var p = new(CirclePlusContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CirclePlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CirclePlusContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CirclePlusContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CirclePlusContext) CIRCLEPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCIRCLEPLUS, 0)
}

func (s *CirclePlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCirclePlus(s)
	}
}

func (s *CirclePlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCirclePlus(s)
	}
}

func (s *CirclePlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCirclePlus(s)

	default:
		return t.VisitChildren(s)
	}
}

type BecauseContext struct {
	*ExprContext
}

func NewBecauseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BecauseContext {
	var p = new(BecauseContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BecauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BecauseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BecauseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BecauseContext) BECAUSE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBECAUSE, 0)
}

func (s *BecauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBecause(s)
	}
}

func (s *BecauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBecause(s)
	}
}

func (s *BecauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBecause(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringLiteralContext struct {
	*ExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserStringLiteral, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndContext struct {
	*ExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) DOUBLEAMP() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEAMP, 0)
}

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserAND, 0)
}

func (s *AndContext) NAND() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNAND, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

type GetContext struct {
	*ExprContext
}

func NewGetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetContext {
	var p = new(GetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetContext) DOUBLELESS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLELESS, 0)
}

func (s *GetContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserStringLiteral, 0)
}

func (s *GetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterGet(s)
	}
}

func (s *GetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitGet(s)
	}
}

func (s *GetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitGet(s)

	default:
		return t.VisitChildren(s)
	}
}

type EquivalentContext struct {
	*ExprContext
}

func NewEquivalentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EquivalentContext {
	var p = new(EquivalentContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EquivalentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EquivalentContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EquivalentContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EquivalentContext) LRDOUBLEARROW() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLRDOUBLEARROW, 0)
}

func (s *EquivalentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterEquivalent(s)
	}
}

func (s *EquivalentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitEquivalent(s)
	}
}

func (s *EquivalentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitEquivalent(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompoundExpressionContext struct {
	*ExprContext
}

func NewCompoundExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompoundExpressionContext {
	var p = new(CompoundExpressionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompoundExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundExpressionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompoundExpressionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompoundExpressionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSEMICOLON, 0)
}

func (s *CompoundExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCompoundExpression(s)
	}
}

func (s *CompoundExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCompoundExpression(s)
	}
}

func (s *CompoundExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCompoundExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type DerivativeContext struct {
	*ExprContext
}

func NewDerivativeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DerivativeContext {
	var p = new(DerivativeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DerivativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DerivativeContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DerivativeContext) AllSINGLEQUOTE() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserSINGLEQUOTE)
}

func (s *DerivativeContext) SINGLEQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSINGLEQUOTE, i)
}

func (s *DerivativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDerivative(s)
	}
}

func (s *DerivativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDerivative(s)
	}
}

func (s *DerivativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDerivative(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlotContext struct {
	*ExprContext
}

func NewSlotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotContext {
	var p = new(SlotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SlotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotContext) SlotExpression() ISlotExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISlotExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISlotExpressionContext)
}

func (s *SlotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlot(s)
	}
}

func (s *SlotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlot(s)
	}
}

func (s *SlotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlot(s)

	default:
		return t.VisitChildren(s)
	}
}

type RightTeeContext struct {
	*ExprContext
}

func NewRightTeeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightTeeContext {
	var p = new(RightTeeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RightTeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightTeeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RightTeeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RightTeeContext) RIGHTTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRIGHTTEE, 0)
}

func (s *RightTeeContext) DOUBLERIGHTTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLERIGHTTEE, 0)
}

func (s *RightTeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterRightTee(s)
	}
}

func (s *RightTeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitRightTee(s)
	}
}

func (s *RightTeeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitRightTee(s)

	default:
		return t.VisitChildren(s)
	}
}

type XorContext struct {
	*ExprContext
}

func NewXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XorContext {
	var p = new(XorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *XorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *XorContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *XorContext) XOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserXOR, 0)
}

func (s *XorContext) XNOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserXNOR, 0)
}

func (s *XorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterXor(s)
	}
}

func (s *XorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitXor(s)
	}
}

func (s *XorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitXor(s)

	default:
		return t.VisitChildren(s)
	}
}

type RuleContext struct {
	*ExprContext
}

func NewRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RuleContext {
	var p = new(RuleContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RuleContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RuleContext) MINUSGREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMINUSGREATER, 0)
}

func (s *RuleContext) RARROW() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRARROW, 0)
}

func (s *RuleContext) COLONGREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOLONGREATER, 0)
}

func (s *RuleContext) COLONARROW() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOLONARROW, 0)
}

func (s *RuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitRule(s)
	}
}

func (s *RuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitRule(s)

	default:
		return t.VisitChildren(s)
	}
}

type HeadExpressionContext struct {
	*ExprContext
}

func NewHeadExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HeadExpressionContext {
	var p = new(HeadExpressionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *HeadExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeadExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *HeadExpressionContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLBRACKET, 0)
}

func (s *HeadExpressionContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *HeadExpressionContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRBRACKET, 0)
}

func (s *HeadExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterHeadExpression(s)
	}
}

func (s *HeadExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitHeadExpression(s)
	}
}

func (s *HeadExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitHeadExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReplaceAllContext struct {
	*ExprContext
}

func NewReplaceAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReplaceAllContext {
	var p = new(ReplaceAllContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ReplaceAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceAllContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ReplaceAllContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReplaceAllContext) SLASHDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHDOT, 0)
}

func (s *ReplaceAllContext) DOUBLESLASHDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLESLASHDOT, 0)
}

func (s *ReplaceAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterReplaceAll(s)
	}
}

func (s *ReplaceAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitReplaceAll(s)
	}
}

func (s *ReplaceAllContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitReplaceAll(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntersectionContext struct {
	*ExprContext
}

func NewIntersectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntersectionContext {
	var p = new(IntersectionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntersectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntersectionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IntersectionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IntersectionContext) INTERSECTION() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserINTERSECTION, 0)
}

func (s *IntersectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterIntersection(s)
	}
}

func (s *IntersectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitIntersection(s)
	}
}

func (s *IntersectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitIntersection(s)

	default:
		return t.VisitChildren(s)
	}
}

type PreIncrementContext struct {
	*ExprContext
}

func NewPreIncrementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreIncrementContext {
	var p = new(PreIncrementContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PreIncrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreIncrementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PreIncrementContext) DOUBLEPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEPLUS, 0)
}

func (s *PreIncrementContext) DOUBLEMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEMINUS, 0)
}

func (s *PreIncrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPreIncrement(s)
	}
}

func (s *PreIncrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPreIncrement(s)
	}
}

func (s *PreIncrementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPreIncrement(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntegrateContext struct {
	*ExprContext
}

func NewIntegrateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegrateContext {
	var p = new(IntegrateContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IntegrateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegrateContext) INTEGRAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserINTEGRAL, 0)
}

func (s *IntegrateContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IntegrateContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IntegrateContext) DIFFERENTIALD() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIFFERENTIALD, 0)
}

func (s *IntegrateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterIntegrate(s)
	}
}

func (s *IntegrateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitIntegrate(s)
	}
}

func (s *IntegrateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitIntegrate(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetContext struct {
	*ExprContext
}

func NewSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetContext {
	var p = new(SetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SetContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SetContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEQUAL, 0)
}

func (s *SetContext) COLONEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOLONEQUAL, 0)
}

func (s *SetContext) CARETEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCARETEQUAL, 0)
}

func (s *SetContext) CARETCOLONEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCARETCOLONEQUAL, 0)
}

func (s *SetContext) FUNCTIONARROW() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserFUNCTIONARROW, 0)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSet(s)
	}
}

func (s *SetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSet(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpEqualsContext struct {
	*ExprContext
}

func NewOpEqualsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpEqualsContext {
	var p = new(OpEqualsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OpEqualsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpEqualsContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OpEqualsContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpEqualsContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPLUSEQUAL, 0)
}

func (s *OpEqualsContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMINUSEQUAL, 0)
}

func (s *OpEqualsContext) ASTERISKEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserASTERISKEQUAL, 0)
}

func (s *OpEqualsContext) SLASHEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHEQUAL, 0)
}

func (s *OpEqualsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOpEquals(s)
	}
}

func (s *OpEqualsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOpEquals(s)
	}
}

func (s *OpEqualsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOpEquals(s)

	default:
		return t.VisitChildren(s)
	}
}

type MessageContext struct {
	*ExprContext
}

func NewMessageContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MessageContext {
	var p = new(MessageContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MessageContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MessageContext) AllDOUBLECOLON() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserDOUBLECOLON)
}

func (s *MessageContext) DOUBLECOLON(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLECOLON, i)
}

func (s *MessageContext) AllStringLiteral() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserStringLiteral)
}

func (s *MessageContext) StringLiteral(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserStringLiteral, i)
}

func (s *MessageContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterMessage(s)
	}
}

func (s *MessageContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitMessage(s)
	}
}

func (s *MessageContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitMessage(s)

	default:
		return t.VisitChildren(s)
	}
}

type CrossContext struct {
	*ExprContext
}

func NewCrossContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CrossContext {
	var p = new(CrossContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CrossContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CrossContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CrossContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CrossContext) CROSS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCROSS, 0)
}

func (s *CrossContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCross(s)
	}
}

func (s *CrossContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCross(s)
	}
}

func (s *CrossContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCross(s)

	default:
		return t.VisitChildren(s)
	}
}

type PatternTestContext struct {
	*ExprContext
}

func NewPatternTestContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PatternTestContext {
	var p = new(PatternTestContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PatternTestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternTestContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PatternTestContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PatternTestContext) QUESTIONMARK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserQUESTIONMARK, 0)
}

func (s *PatternTestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPatternTest(s)
	}
}

func (s *PatternTestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPatternTest(s)
	}
}

func (s *PatternTestContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPatternTest(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixContext struct {
	*ExprContext
}

func NewPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixContext {
	var p = new(PrefixContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PrefixContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PrefixContext) AT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserAT, 0)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (s *PrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type BackslashContext struct {
	*ExprContext
}

func NewBackslashContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BackslashContext {
	var p = new(BackslashContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BackslashContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BackslashContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BackslashContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BackslashContext) RAWBACKSLASH() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRAWBACKSLASH, 0)
}

func (s *BackslashContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBackslash(s)
	}
}

func (s *BackslashContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBackslash(s)
	}
}

func (s *BackslashContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBackslash(s)

	default:
		return t.VisitChildren(s)
	}
}

type RepeatedContext struct {
	*ExprContext
}

func NewRepeatedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatedContext {
	var p = new(RepeatedContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RepeatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatedContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RepeatedContext) DOUBLEDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEDOT, 0)
}

func (s *RepeatedContext) TRIPPLEDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRIPPLEDOT, 0)
}

func (s *RepeatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterRepeated(s)
	}
}

func (s *RepeatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitRepeated(s)
	}
}

func (s *RepeatedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitRepeated(s)

	default:
		return t.VisitChildren(s)
	}
}

type MapApplyContext struct {
	*ExprContext
}

func NewMapApplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MapApplyContext {
	var p = new(MapApplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MapApplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapApplyContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MapApplyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MapApplyContext) MAP() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMAP, 0)
}

func (s *MapApplyContext) MAPALL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMAPALL, 0)
}

func (s *MapApplyContext) DOUBLEAT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEAT, 0)
}

func (s *MapApplyContext) TRIPPLEAT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRIPPLEAT, 0)
}

func (s *MapApplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterMapApply(s)
	}
}

func (s *MapApplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitMapApply(s)
	}
}

func (s *MapApplyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitMapApply(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnionContext struct {
	*ExprContext
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *UnionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnionContext) UNION() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserUNION, 0)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitUnion(s)

	default:
		return t.VisitChildren(s)
	}
}

type VerticalSeparatorContext struct {
	*ExprContext
}

func NewVerticalSeparatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VerticalSeparatorContext {
	var p = new(VerticalSeparatorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VerticalSeparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalSeparatorContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *VerticalSeparatorContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VerticalSeparatorContext) VERTICALSEPARATOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserVERTICALSEPARATOR, 0)
}

func (s *VerticalSeparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterVerticalSeparator(s)
	}
}

func (s *VerticalSeparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitVerticalSeparator(s)
	}
}

func (s *VerticalSeparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitVerticalSeparator(s)

	default:
		return t.VisitChildren(s)
	}
}

type FactorialContext struct {
	*ExprContext
}

func NewFactorialContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FactorialContext {
	var p = new(FactorialContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FactorialContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorialContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FactorialContext) BANG() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBANG, 0)
}

func (s *FactorialContext) DOUBLEBANG() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEBANG, 0)
}

func (s *FactorialContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterFactorial(s)
	}
}

func (s *FactorialContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitFactorial(s)
	}
}

func (s *FactorialContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitFactorial(s)

	default:
		return t.VisitChildren(s)
	}
}

type SpanAContext struct {
	*ExprContext
}

func NewSpanAContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpanAContext {
	var p = new(SpanAContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SpanAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanAContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SpanAContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SpanAContext) DOUBLESEMICOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLESEMICOLON, 0)
}

func (s *SpanAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSpanA(s)
	}
}

func (s *SpanAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSpanA(s)
	}
}

func (s *SpanAContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSpanA(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionContext struct {
	*ExprContext
}

func NewFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionContext {
	var p = new(FunctionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunctionContext) AMP() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserAMP, 0)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberContext struct {
	*ExprContext
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (s *NumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

type StarContext struct {
	*ExprContext
}

func NewStarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StarContext {
	var p = new(StarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StarContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StarContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StarContext) STAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSTAR, 0)
}

func (s *StarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterStar(s)
	}
}

func (s *StarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitStar(s)
	}
}

func (s *StarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitStar(s)

	default:
		return t.VisitChildren(s)
	}
}

type ComparisonContext struct {
	*ExprContext
}

func NewComparisonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonContext {
	var p = new(ComparisonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ComparisonContext) EqualSymbol() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEqualSymbol, 0)
}

func (s *ComparisonContext) NotEqualSymbol() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNotEqualSymbol, 0)
}

func (s *ComparisonContext) GreaterEqualSymbol() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserGreaterEqualSymbol, 0)
}

func (s *ComparisonContext) LessEqualSymbol() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLessEqualSymbol, 0)
}

func (s *ComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserGREATER, 0)
}

func (s *ComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLESS, 0)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

type TagSetContext struct {
	*ExprContext
}

func NewTagSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TagSetContext {
	var p = new(TagSetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TagSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TagSetContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *TagSetContext) SLASHCOLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHCOLON, 0)
}

func (s *TagSetContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TagSetContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TagSetContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEQUAL, 0)
}

func (s *TagSetContext) COLONEQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOLONEQUAL, 0)
}

func (s *TagSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterTagSet(s)
	}
}

func (s *TagSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitTagSet(s)
	}
}

func (s *TagSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitTagSet(s)

	default:
		return t.VisitChildren(s)
	}
}

type IncrementContext struct {
	*ExprContext
}

func NewIncrementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IncrementContext {
	var p = new(IncrementContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IncrementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncrementContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IncrementContext) DOUBLEPLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEPLUS, 0)
}

func (s *IncrementContext) DOUBLEMINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEMINUS, 0)
}

func (s *IncrementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterIncrement(s)
	}
}

func (s *IncrementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitIncrement(s)
	}
}

func (s *IncrementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitIncrement(s)

	default:
		return t.VisitChildren(s)
	}
}

type VerticalTildeContext struct {
	*ExprContext
}

func NewVerticalTildeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VerticalTildeContext {
	var p = new(VerticalTildeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VerticalTildeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerticalTildeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *VerticalTildeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VerticalTildeContext) VERTICALTILDE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserVERTICALTILDE, 0)
}

func (s *VerticalTildeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterVerticalTilde(s)
	}
}

func (s *VerticalTildeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitVerticalTilde(s)
	}
}

func (s *VerticalTildeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitVerticalTilde(s)

	default:
		return t.VisitChildren(s)
	}
}

type ColonContext struct {
	*ExprContext
}

func NewColonContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColonContext {
	var p = new(ColonContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ColonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColonContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ColonContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ColonContext) COLON() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOLON, 0)
}

func (s *ColonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterColon(s)
	}
}

func (s *ColonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitColon(s)
	}
}

func (s *ColonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitColon(s)

	default:
		return t.VisitChildren(s)
	}
}

type SmallCircleContext struct {
	*ExprContext
}

func NewSmallCircleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SmallCircleContext {
	var p = new(SmallCircleContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SmallCircleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SmallCircleContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SmallCircleContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SmallCircleContext) SMALLCIRCLE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSMALLCIRCLE, 0)
}

func (s *SmallCircleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSmallCircle(s)
	}
}

func (s *SmallCircleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSmallCircle(s)
	}
}

func (s *SmallCircleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSmallCircle(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesesContext struct {
	*ExprContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLPAREN, 0)
}

func (s *ParenthesesContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRPAREN, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitParentheses(s)
	}
}

func (s *ParenthesesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitParentheses(s)

	default:
		return t.VisitChildren(s)
	}
}

type SpanBContext struct {
	*ExprContext
}

func NewSpanBContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpanBContext {
	var p = new(SpanBContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SpanBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanBContext) SPANSEMICOLONS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSPANSEMICOLONS, 0)
}

func (s *SpanBContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SpanBContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SpanBContext) AllDOUBLESEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserDOUBLESEMICOLON)
}

func (s *SpanBContext) DOUBLESEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLESEMICOLON, i)
}

func (s *SpanBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSpanB(s)
	}
}

func (s *SpanBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSpanB(s)
	}
}

func (s *SpanBContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSpanB(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionContext struct {
	*ExprContext
}

func NewConditionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionContext {
	var p = new(ConditionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ConditionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ConditionContext) SLASHSEMI() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSLASHSEMI, 0)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

type FloorContext struct {
	*ExprContext
}

func NewFloorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloorContext {
	var p = new(FloorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorContext) LFLOOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLFLOOR, 0)
}

func (s *FloorContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloorContext) RFLOOR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRFLOOR, 0)
}

func (s *FloorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterFloor(s)
	}
}

func (s *FloorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitFloor(s)
	}
}

func (s *FloorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitFloor(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompositionContext struct {
	*ExprContext
}

func NewCompositionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompositionContext {
	var p = new(CompositionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompositionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompositionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompositionContext) ATASTERISK() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserATASTERISK, 0)
}

func (s *CompositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterComposition(s)
	}
}

func (s *CompositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitComposition(s)
	}
}

func (s *CompositionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitComposition(s)

	default:
		return t.VisitChildren(s)
	}
}

type CircleDotContext struct {
	*ExprContext
}

func NewCircleDotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CircleDotContext {
	var p = new(CircleDotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CircleDotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircleDotContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CircleDotContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CircleDotContext) CIRCLEDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCIRCLEDOT, 0)
}

func (s *CircleDotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCircleDot(s)
	}
}

func (s *CircleDotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCircleDot(s)
	}
}

func (s *CircleDotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCircleDot(s)

	default:
		return t.VisitChildren(s)
	}
}

type SymbolLiteralContext struct {
	*ExprContext
}

func NewSymbolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SymbolLiteralContext {
	var p = new(SymbolLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SymbolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolLiteralContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *SymbolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSymbolLiteral(s)
	}
}

func (s *SymbolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSymbolLiteral(s)
	}
}

func (s *SymbolLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSymbolLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

type CircleTimesContext struct {
	*ExprContext
}

func NewCircleTimesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CircleTimesContext {
	var p = new(CircleTimesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CircleTimesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircleTimesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CircleTimesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CircleTimesContext) CIRCLETIMES() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCIRCLETIMES, 0)
}

func (s *CircleTimesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCircleTimes(s)
	}
}

func (s *CircleTimesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCircleTimes(s)
	}
}

func (s *CircleTimesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCircleTimes(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnsetContext struct {
	*ExprContext
}

func NewUnsetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnsetContext {
	var p = new(UnsetContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnsetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnsetContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnsetContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserEQUAL, 0)
}

func (s *UnsetContext) DOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOT, 0)
}

func (s *UnsetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterUnset(s)
	}
}

func (s *UnsetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitUnset(s)
	}
}

func (s *UnsetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitUnset(s)

	default:
		return t.VisitChildren(s)
	}
}

type DelContext struct {
	*ExprContext
}

func NewDelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DelContext {
	var p = new(DelContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DelContext) DEL() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDEL, 0)
}

func (s *DelContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDel(s)
	}
}

func (s *DelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDel(s)
	}
}

func (s *DelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDel(s)

	default:
		return t.VisitChildren(s)
	}
}

type DiamondContext struct {
	*ExprContext
}

func NewDiamondContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DiamondContext {
	var p = new(DiamondContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DiamondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiamondContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DiamondContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DiamondContext) DIAMOND() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIAMOND, 0)
}

func (s *DiamondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterDiamond(s)
	}
}

func (s *DiamondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitDiamond(s)
	}
}

func (s *DiamondContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitDiamond(s)

	default:
		return t.VisitChildren(s)
	}
}

type WedgeContext struct {
	*ExprContext
}

func NewWedgeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WedgeContext {
	var p = new(WedgeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *WedgeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WedgeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *WedgeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WedgeContext) WEDGE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserWEDGE, 0)
}

func (s *WedgeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterWedge(s)
	}
}

func (s *WedgeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitWedge(s)
	}
}

func (s *WedgeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitWedge(s)

	default:
		return t.VisitChildren(s)
	}
}

type PutContext struct {
	*ExprContext
}

func NewPutContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PutContext {
	var p = new(PutContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PutContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PutContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserStringLiteral, 0)
}

func (s *PutContext) DOUBLEGREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEGREATER, 0)
}

func (s *PutContext) TRIPPLEGREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserTRIPPLEGREATER, 0)
}

func (s *PutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPut(s)
	}
}

func (s *PutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPut(s)
	}
}

func (s *PutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPut(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringJoinContext struct {
	*ExprContext
}

func NewStringJoinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringJoinContext {
	var p = new(StringJoinContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringJoinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringJoinContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StringJoinContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StringJoinContext) LESSGREATER() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLESSGREATER, 0)
}

func (s *StringJoinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterStringJoin(s)
	}
}

func (s *StringJoinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitStringJoin(s)
	}
}

func (s *StringJoinContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitStringJoin(s)

	default:
		return t.VisitChildren(s)
	}
}

type TeeContext struct {
	*ExprContext
}

func NewTeeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TeeContext {
	var p = new(TeeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TeeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TeeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TeeContext) LEFTTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLEFTTEE, 0)
}

func (s *TeeContext) DOUBLELEFTTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLELEFTTEE, 0)
}

func (s *TeeContext) UPTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserUPTEE, 0)
}

func (s *TeeContext) DOWNTEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOWNTEE, 0)
}

func (s *TeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterTee(s)
	}
}

func (s *TeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitTee(s)
	}
}

func (s *TeeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitTee(s)

	default:
		return t.VisitChildren(s)
	}
}

type SetContainmentContext struct {
	*ExprContext
}

func NewSetContainmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetContainmentContext {
	var p = new(SetContainmentContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SetContainmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContainmentContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SetContainmentContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SetContainmentContext) ELEMENT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserELEMENT, 0)
}

func (s *SetContainmentContext) NOTELEMENT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNOTELEMENT, 0)
}

func (s *SetContainmentContext) SUBSET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSUBSET, 0)
}

func (s *SetContainmentContext) SUPERSET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserSUPERSET, 0)
}

func (s *SetContainmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSetContainment(s)
	}
}

func (s *SetContainmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSetContainment(s)
	}
}

func (s *SetContainmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSetContainment(s)

	default:
		return t.VisitChildren(s)
	}
}

type VeeContext struct {
	*ExprContext
}

func NewVeeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VeeContext {
	var p = new(VeeContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *VeeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VeeContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *VeeContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *VeeContext) VEE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserVEE, 0)
}

func (s *VeeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterVee(s)
	}
}

func (s *VeeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitVee(s)
	}
}

func (s *VeeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitVee(s)

	default:
		return t.VisitChildren(s)
	}
}

type CenterDotContext struct {
	*ExprContext
}

func NewCenterDotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CenterDotContext {
	var p = new(CenterDotContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CenterDotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CenterDotContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CenterDotContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CenterDotContext) CENTERDOT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCENTERDOT, 0)
}

func (s *CenterDotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCenterDot(s)
	}
}

func (s *CenterDotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCenterDot(s)
	}
}

func (s *CenterDotContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCenterDot(s)

	default:
		return t.VisitChildren(s)
	}
}

type TimesContext struct {
	*ExprContext
}

func NewTimesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimesContext {
	var p = new(TimesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TimesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TimesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TimesContext) MultiplicationSymbol() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMultiplicationSymbol, 0)
}

func (s *TimesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterTimes(s)
	}
}

func (s *TimesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitTimes(s)
	}
}

func (s *TimesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitTimes(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringExpressionContext struct {
	*ExprContext
}

func NewStringExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringExpressionContext {
	var p = new(StringExpressionContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringExpressionContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *StringExpressionContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *StringExpressionContext) DOUBLETILDE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLETILDE, 0)
}

func (s *StringExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterStringExpression(s)
	}
}

func (s *StringExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitStringExpression(s)
	}
}

func (s *StringExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitStringExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketingBarContext struct {
	*ExprContext
}

func NewBracketingBarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketingBarContext {
	var p = new(BracketingBarContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BracketingBarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketingBarContext) LBRACKETINGBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLBRACKETINGBAR, 0)
}

func (s *BracketingBarContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *BracketingBarContext) RBRACKETINGBAR() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRBRACKETINGBAR, 0)
}

func (s *BracketingBarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBracketingBar(s)
	}
}

func (s *BracketingBarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBracketingBar(s)
	}
}

func (s *BracketingBarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBracketingBar(s)

	default:
		return t.VisitChildren(s)
	}
}

type CoproductContext struct {
	*ExprContext
}

func NewCoproductContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoproductContext {
	var p = new(CoproductContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CoproductContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoproductContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CoproductContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CoproductContext) COPRODUCT() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOPRODUCT, 0)
}

func (s *CoproductContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCoproduct(s)
	}
}

func (s *CoproductContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCoproduct(s)
	}
}

func (s *CoproductContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCoproduct(s)

	default:
		return t.VisitChildren(s)
	}
}

type AngleBracketContext struct {
	*ExprContext
}

func NewAngleBracketContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AngleBracketContext {
	var p = new(AngleBracketContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AngleBracketContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AngleBracketContext) LANGLE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLANGLE, 0)
}

func (s *AngleBracketContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AngleBracketContext) RANGLE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRANGLE, 0)
}

func (s *AngleBracketContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAngleBracket(s)
	}
}

func (s *AngleBracketContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAngleBracket(s)
	}
}

func (s *AngleBracketContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAngleBracket(s)

	default:
		return t.VisitChildren(s)
	}
}

type PowerContext struct {
	*ExprContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowerContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowerContext) CARET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCARET, 0)
}

func (s *PowerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterPower(s)
	}
}

func (s *PowerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitPower(s)
	}
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *FoxySheepParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, FoxySheepParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(39)
			p.NumberLiteral()
		}

	case 2:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(40)
			p.Match(FoxySheepParserStringLiteral)
		}

	case 3:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.Match(FoxySheepParserLPAREN)
		}
		{
			p.SetState(42)
			p.expr(0)
		}
		{
			p.SetState(43)
			p.Match(FoxySheepParserRPAREN)
		}

	case 4:
		localctx = NewListContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(45)
			p.Match(FoxySheepParserLBRACE)
		}
		{
			p.SetState(46)
			p.ExpressionList()
		}
		{
			p.SetState(47)
			p.Match(FoxySheepParserRBRACE)
		}

	case 5:
		localctx = NewAngleBracketContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(FoxySheepParserLANGLE)
		}
		{
			p.SetState(50)
			p.ExpressionList()
		}
		{
			p.SetState(51)
			p.Match(FoxySheepParserRANGLE)
		}

	case 6:
		localctx = NewFloorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(53)
			p.Match(FoxySheepParserLFLOOR)
		}
		{
			p.SetState(54)
			p.expr(0)
		}
		{
			p.SetState(55)
			p.Match(FoxySheepParserRFLOOR)
		}

	case 7:
		localctx = NewCeilingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(57)
			p.Match(FoxySheepParserLCEILING)
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		{
			p.SetState(59)
			p.Match(FoxySheepParserRCEILING)
		}

	case 8:
		localctx = NewBracketingBarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(FoxySheepParserLBRACKETINGBAR)
		}
		{
			p.SetState(62)
			p.ExpressionList()
		}
		{
			p.SetState(63)
			p.Match(FoxySheepParserRBRACKETINGBAR)
		}

	case 9:
		localctx = NewDoubleBracketingBarContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)
			p.Match(FoxySheepParserLDOUBLEBRACKETINGBAR)
		}
		{
			p.SetState(66)
			p.ExpressionList()
		}
		{
			p.SetState(67)
			p.Match(FoxySheepParserRDOUBLEBRACKETINGBAR)
		}

	case 10:
		localctx = NewSlotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.SlotExpression()
		}

	case 11:
		localctx = NewOutContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)
			p.OutExpression()
		}

	case 12:
		localctx = NewPatternBlanksContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(72)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FoxySheepParserName || _la == FoxySheepParserBACKQUOTE {
			{
				p.SetState(71)
				p.Symbol()
			}

		}
		p.SetState(74)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANK-77)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(75)
				p.expr(0)
			}

		}

	case 13:
		localctx = NewPatternBlankDotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FoxySheepParserName || _la == FoxySheepParserBACKQUOTE {
			{
				p.SetState(78)
				p.Symbol()
			}

		}
		{
			p.SetState(81)
			p.Match(FoxySheepParserBLANKDOT)
		}

	case 14:
		localctx = NewSymbolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Symbol()
		}

	case 15:
		localctx = NewGetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(FoxySheepParserDOUBLELESS)
		}
		{
			p.SetState(84)
			p.Match(FoxySheepParserStringLiteral)
		}

	case 16:
		localctx = NewBoxParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(85)
			p.Match(FoxySheepParserBoxLeftBoxParenthesis)
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserInterpretedBox-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxSqrt-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
			{
				p.SetState(86)
				p.box(0)
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			p.Match(FoxySheepParserBoxRightBoxParenthesis)
		}

	case 17:
		localctx = NewBoxConstructorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(93)
			p.Match(FoxySheepParserBoxConstructor)
		}
		{
			p.SetState(94)
			p.expr(76)
		}

	case 18:
		localctx = NewPreIncrementContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(95)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserDOUBLEPLUS || _la == FoxySheepParserDOUBLEMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(96)
			p.expr(74)
		}

	case 19:
		localctx = NewIntegrateContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(97)
			p.Match(FoxySheepParserINTEGRAL)
		}
		{
			p.SetState(98)
			p.expr(0)
		}
		{
			p.SetState(99)
			p.Match(FoxySheepParserDIFFERENTIALD)
		}
		{
			p.SetState(100)
			p.expr(63)
		}

	case 20:
		localctx = NewDelContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Match(FoxySheepParserDEL)
		}
		{
			p.SetState(103)
			p.expr(62)
		}

	case 21:
		localctx = NewSquareContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.Match(FoxySheepParserSQUARE)
		}
		{
			p.SetState(105)
			p.expr(61)
		}

	case 22:
		localctx = NewUnaryPlusMinusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(106)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(107)
			p.expr(55)
		}

	case 23:
		localctx = NewSpanBContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(108)
			p.Match(FoxySheepParserSPANSEMICOLONS)
		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(109)
				p.expr(0)
			}

		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(112)
					p.Match(FoxySheepParserDOUBLESEMICOLON)
				}
				p.SetState(114)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(113)
						p.expr(0)
					}

				}

			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}

	case 24:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserNOT || _la == FoxySheepParserBANG) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(122)
			p.expr(30)
		}

	case 25:
		localctx = NewPatternExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(123)
			p.Symbol()
		}
		{
			p.SetState(124)
			p.Match(FoxySheepParserRAWCOLON)
		}
		{
			p.SetState(125)
			p.expr(19)
		}

	case 26:
		localctx = NewTagSetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(127)
			p.Symbol()
		}
		{
			p.SetState(128)
			p.Match(FoxySheepParserSLASHCOLON)
		}
		{
			p.SetState(129)
			p.expr(0)
		}
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserEQUAL || _la == FoxySheepParserCOLONEQUAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(131)
			p.expr(4)
		}

	case 27:
		localctx = NewTagUnsetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Symbol()
		}
		{
			p.SetState(134)
			p.Match(FoxySheepParserSLASHCOLON)
		}
		{
			p.SetState(135)
			p.expr(0)
		}
		{
			p.SetState(136)
			p.Match(FoxySheepParserEQUAL)
		}
		{
			p.SetState(137)
			p.Match(FoxySheepParserDOT)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPatternTestContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(141)

				if !(p.Precpred(p.GetParserRuleContext(), 79)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 79)", ""))
				}
				{
					p.SetState(142)
					p.Match(FoxySheepParserQUESTIONMARK)
				}
				{
					p.SetState(143)
					p.expr(80)
				}

			case 2:
				localctx = NewCompositionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(144)

				if !(p.Precpred(p.GetParserRuleContext(), 73)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 73)", ""))
				}
				{
					p.SetState(145)
					p.Match(FoxySheepParserATASTERISK)
				}
				{
					p.SetState(146)
					p.expr(74)
				}

			case 3:
				localctx = NewRightCompositionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 72)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 72)", ""))
				}
				{
					p.SetState(148)
					p.Match(FoxySheepParserSLASHASTERISK)
				}
				{
					p.SetState(149)
					p.expr(73)
				}

			case 4:
				localctx = NewPrefixContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(150)

				if !(p.Precpred(p.GetParserRuleContext(), 71)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 71)", ""))
				}
				{
					p.SetState(151)
					p.Match(FoxySheepParserAT)
				}
				{
					p.SetState(152)
					p.expr(71)
				}

			case 5:
				localctx = NewInfixContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 70)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 70)", ""))
				}
				{
					p.SetState(154)
					p.Match(FoxySheepParserTILDE)
				}
				{
					p.SetState(155)
					p.expr(0)
				}
				{
					p.SetState(156)
					p.Match(FoxySheepParserTILDE)
				}
				{
					p.SetState(157)
					p.expr(71)
				}

			case 6:
				localctx = NewMapApplyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(159)

				if !(p.Precpred(p.GetParserRuleContext(), 69)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 69)", ""))
				}
				p.SetState(160)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-114)&-(0x1f+1)) == 0 && ((1<<uint((_la-114)))&((1<<(FoxySheepParserTRIPPLEAT-114))|(1<<(FoxySheepParserDOUBLEAT-114))|(1<<(FoxySheepParserMAP-114))|(1<<(FoxySheepParserMAPALL-114)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(161)
					p.expr(69)
				}

			case 7:
				localctx = NewStringJoinContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 65)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 65)", ""))
				}
				{
					p.SetState(163)
					p.Match(FoxySheepParserLESSGREATER)
				}
				{
					p.SetState(164)
					p.expr(66)
				}

			case 8:
				localctx = NewPowerContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 64)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 64)", ""))
				}
				{
					p.SetState(166)
					p.Match(FoxySheepParserCARET)
				}
				{
					p.SetState(167)
					p.expr(64)
				}

			case 9:
				localctx = NewSmallCircleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 60)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 60)", ""))
				}
				{
					p.SetState(169)
					p.Match(FoxySheepParserSMALLCIRCLE)
				}
				{
					p.SetState(170)
					p.expr(61)
				}

			case 10:
				localctx = NewCircleDotContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 59)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 59)", ""))
				}
				{
					p.SetState(172)
					p.Match(FoxySheepParserCIRCLEDOT)
				}
				{
					p.SetState(173)
					p.expr(60)
				}

			case 11:
				localctx = NewNonCommutativeMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 58)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 58)", ""))
				}
				{
					p.SetState(175)
					p.Match(FoxySheepParserDOUBLEASTERISK)
				}
				{
					p.SetState(176)
					p.expr(59)
				}

			case 12:
				localctx = NewCrossContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 57)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 57)", ""))
				}
				{
					p.SetState(178)
					p.Match(FoxySheepParserCROSS)
				}
				{
					p.SetState(179)
					p.expr(58)
				}

			case 13:
				localctx = NewDotContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 56)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 56)", ""))
				}
				{
					p.SetState(181)
					p.Match(FoxySheepParserDOT)
				}
				{
					p.SetState(182)
					p.expr(57)
				}

			case 14:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 54)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 54)", ""))
				}
				p.SetState(184)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserSLASH || _la == FoxySheepParserDIVIDE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(185)
					p.expr(55)
				}

			case 15:
				localctx = NewBackslashContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 53)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 53)", ""))
				}
				{
					p.SetState(187)
					p.Match(FoxySheepParserRAWBACKSLASH)
				}
				{
					p.SetState(188)
					p.expr(54)
				}

			case 16:
				localctx = NewDiamondContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 52)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 52)", ""))
				}
				{
					p.SetState(190)
					p.Match(FoxySheepParserDIAMOND)
				}
				{
					p.SetState(191)
					p.expr(53)
				}

			case 17:
				localctx = NewWedgeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 51)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 51)", ""))
				}
				{
					p.SetState(193)
					p.Match(FoxySheepParserWEDGE)
				}
				{
					p.SetState(194)
					p.expr(52)
				}

			case 18:
				localctx = NewVeeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 50)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 50)", ""))
				}
				{
					p.SetState(196)
					p.Match(FoxySheepParserVEE)
				}
				{
					p.SetState(197)
					p.expr(51)
				}

			case 19:
				localctx = NewCircleTimesContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 49)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 49)", ""))
				}
				{
					p.SetState(199)
					p.Match(FoxySheepParserCIRCLETIMES)
				}
				{
					p.SetState(200)
					p.expr(50)
				}

			case 20:
				localctx = NewCenterDotContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 48)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 48)", ""))
				}
				{
					p.SetState(202)
					p.Match(FoxySheepParserCENTERDOT)
				}
				{
					p.SetState(203)
					p.expr(49)
				}

			case 21:
				localctx = NewTimesContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 47)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 47)", ""))
				}
				p.SetState(206)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == FoxySheepParserMultiplicationSymbol {
					{
						p.SetState(205)
						p.Match(FoxySheepParserMultiplicationSymbol)
					}

				}
				{
					p.SetState(208)
					p.expr(48)
				}

			case 22:
				localctx = NewStarContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(209)

				if !(p.Precpred(p.GetParserRuleContext(), 46)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 46)", ""))
				}
				{
					p.SetState(210)
					p.Match(FoxySheepParserSTAR)
				}
				{
					p.SetState(211)
					p.expr(47)
				}

			case 23:
				localctx = NewVerticalTildeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(212)

				if !(p.Precpred(p.GetParserRuleContext(), 45)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 45)", ""))
				}
				{
					p.SetState(213)
					p.Match(FoxySheepParserVERTICALTILDE)
				}
				{
					p.SetState(214)
					p.expr(46)
				}

			case 24:
				localctx = NewCoproductContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(215)

				if !(p.Precpred(p.GetParserRuleContext(), 44)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 44)", ""))
				}
				{
					p.SetState(216)
					p.Match(FoxySheepParserCOPRODUCT)
				}
				{
					p.SetState(217)
					p.expr(45)
				}

			case 25:
				localctx = NewCapContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(218)

				if !(p.Precpred(p.GetParserRuleContext(), 43)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 43)", ""))
				}
				{
					p.SetState(219)
					p.Match(FoxySheepParserCAP)
				}
				{
					p.SetState(220)
					p.expr(44)
				}

			case 26:
				localctx = NewCupContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 42)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 42)", ""))
				}
				{
					p.SetState(222)
					p.Match(FoxySheepParserCUP)
				}
				{
					p.SetState(223)
					p.expr(43)
				}

			case 27:
				localctx = NewCirclePlusContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 41)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 41)", ""))
				}
				{
					p.SetState(225)
					p.Match(FoxySheepParserCIRCLEPLUS)
				}
				{
					p.SetState(226)
					p.expr(42)
				}

			case 28:
				localctx = NewCircleMinusContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(227)

				if !(p.Precpred(p.GetParserRuleContext(), 40)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 40)", ""))
				}
				{
					p.SetState(228)
					p.Match(FoxySheepParserCIRCLEMINUS)
				}
				{
					p.SetState(229)
					p.expr(41)
				}

			case 29:
				localctx = NewPlusOpContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(230)

				if !(p.Precpred(p.GetParserRuleContext(), 39)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 39)", ""))
				}
				p.SetState(231)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-182)&-(0x1f+1)) == 0 && ((1<<uint((_la-182)))&((1<<(FoxySheepParserBINARYPLUS-182))|(1<<(FoxySheepParserBINARYMINUS-182))|(1<<(FoxySheepParserBINARYMINUSPLUS-182))|(1<<(FoxySheepParserBINARYPLUSMINUS-182)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(232)
					p.expr(40)
				}

			case 30:
				localctx = NewIntersectionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(233)

				if !(p.Precpred(p.GetParserRuleContext(), 38)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 38)", ""))
				}
				{
					p.SetState(234)
					p.Match(FoxySheepParserINTERSECTION)
				}
				{
					p.SetState(235)
					p.expr(39)
				}

			case 31:
				localctx = NewUnionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(236)

				if !(p.Precpred(p.GetParserRuleContext(), 37)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 37)", ""))
				}
				{
					p.SetState(237)
					p.Match(FoxySheepParserUNION)
				}
				{
					p.SetState(238)
					p.expr(38)
				}

			case 32:
				localctx = NewComparisonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(239)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
				}
				p.SetState(240)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(FoxySheepParserEqualSymbol-39))|(1<<(FoxySheepParserNotEqualSymbol-39))|(1<<(FoxySheepParserGreaterEqualSymbol-39))|(1<<(FoxySheepParserLessEqualSymbol-39))|(1<<(FoxySheepParserLESS-39))|(1<<(FoxySheepParserGREATER-39)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(241)
					p.expr(35)
				}

			case 33:
				localctx = NewVerticalBarContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 33)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 33)", ""))
				}
				p.SetState(243)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-46)&-(0x1f+1)) == 0 && ((1<<uint((_la-46)))&((1<<(FoxySheepParserVERTICALBAR-46))|(1<<(FoxySheepParserNOTVERTICALBAR-46))|(1<<(FoxySheepParserDOUBLEVERTICALBAR-46))|(1<<(FoxySheepParserNOTDOUBLEVERTICALBAR-46)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(244)
					p.expr(34)
				}

			case 34:
				localctx = NewSameContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(245)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
				}
				p.SetState(246)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserTRIPPLEEQUAL || _la == FoxySheepParserEQUALBANGEQUAL) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(247)
					p.expr(33)
				}

			case 35:
				localctx = NewSetContainmentContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(248)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
				}
				p.SetState(249)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(FoxySheepParserELEMENT-50))|(1<<(FoxySheepParserNOTELEMENT-50))|(1<<(FoxySheepParserSUBSET-50))|(1<<(FoxySheepParserSUPERSET-50)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(250)
					p.expr(32)
				}

			case 36:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(251)

				if !(p.Precpred(p.GetParserRuleContext(), 29)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 29)", ""))
				}
				p.SetState(252)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-55)&-(0x1f+1)) == 0 && ((1<<uint((_la-55)))&((1<<(FoxySheepParserDOUBLEAMP-55))|(1<<(FoxySheepParserAND-55))|(1<<(FoxySheepParserNAND-55)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(253)
					p.expr(30)
				}

			case 37:
				localctx = NewXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(254)

				if !(p.Precpred(p.GetParserRuleContext(), 28)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 28)", ""))
				}
				p.SetState(255)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserXOR || _la == FoxySheepParserXNOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(256)
					p.expr(29)
				}

			case 38:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 27)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 27)", ""))
				}
				p.SetState(258)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserDOUBLEBAR || _la == FoxySheepParserOR || _la == FoxySheepParserNOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(259)
					p.expr(28)
				}

			case 39:
				localctx = NewEquivalentContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(260)

				if !(p.Precpred(p.GetParserRuleContext(), 26)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 26)", ""))
				}
				{
					p.SetState(261)
					p.Match(FoxySheepParserLRDOUBLEARROW)
				}
				{
					p.SetState(262)
					p.expr(27)
				}

			case 40:
				localctx = NewImpliesContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(263)

				if !(p.Precpred(p.GetParserRuleContext(), 25)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 25)", ""))
				}
				p.SetState(264)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserRDOUBLEARROW || _la == FoxySheepParserLCONTAINS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(265)
					p.expr(25)
				}

			case 41:
				localctx = NewRightTeeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(266)

				if !(p.Precpred(p.GetParserRuleContext(), 24)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 24)", ""))
				}
				p.SetState(267)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserRIGHTTEE || _la == FoxySheepParserDOUBLERIGHTTEE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(268)
					p.expr(24)
				}

			case 42:
				localctx = NewTeeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(269)

				if !(p.Precpred(p.GetParserRuleContext(), 23)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 23)", ""))
				}
				p.SetState(270)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-67)&-(0x1f+1)) == 0 && ((1<<uint((_la-67)))&((1<<(FoxySheepParserLEFTTEE-67))|(1<<(FoxySheepParserDOUBLELEFTTEE-67))|(1<<(FoxySheepParserUPTEE-67))|(1<<(FoxySheepParserDOWNTEE-67)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(271)
					p.expr(24)
				}

			case 43:
				localctx = NewSuchThatContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(272)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(273)
					p.Match(FoxySheepParserSUCHTHAT)
				}
				{
					p.SetState(274)
					p.expr(22)
				}

			case 44:
				localctx = NewAlternativesContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(275)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				{
					p.SetState(276)
					p.Match(FoxySheepParserBAR)
				}
				{
					p.SetState(277)
					p.expr(21)
				}

			case 45:
				localctx = NewOptionalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(278)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(279)
					p.Match(FoxySheepParserRAWCOLON)
				}
				{
					p.SetState(280)
					p.expr(19)
				}

			case 46:
				localctx = NewStringExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(281)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(282)
					p.Match(FoxySheepParserDOUBLETILDE)
				}
				{
					p.SetState(283)
					p.expr(18)
				}

			case 47:
				localctx = NewConditionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(284)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(285)
					p.Match(FoxySheepParserSLASHSEMI)
				}
				{
					p.SetState(286)
					p.expr(17)
				}

			case 48:
				localctx = NewRuleContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(287)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				p.SetState(288)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-87)&-(0x1f+1)) == 0 && ((1<<uint((_la-87)))&((1<<(FoxySheepParserMINUSGREATER-87))|(1<<(FoxySheepParserRARROW-87))|(1<<(FoxySheepParserCOLONGREATER-87))|(1<<(FoxySheepParserCOLONARROW-87)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(289)
					p.expr(15)
				}

			case 49:
				localctx = NewReplaceAllContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(290)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				p.SetState(291)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserSLASHDOT || _la == FoxySheepParserDOUBLESLASHDOT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(292)
					p.expr(15)
				}

			case 50:
				localctx = NewOpEqualsContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(293)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				p.SetState(294)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(FoxySheepParserPLUSEQUAL-98))|(1<<(FoxySheepParserMINUSEQUAL-98))|(1<<(FoxySheepParserASTERISKEQUAL-98))|(1<<(FoxySheepParserSLASHEQUAL-98)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(295)
					p.expr(13)
				}

			case 51:
				localctx = NewColonContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(296)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(297)
					p.Match(FoxySheepParserCOLON)
				}
				{
					p.SetState(298)
					p.expr(12)
				}

			case 52:
				localctx = NewPostfixContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(299)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(300)
					p.Match(FoxySheepParserDOUBLESLASH)
				}
				{
					p.SetState(301)
					p.expr(11)
				}

			case 53:
				localctx = NewVerticalSeparatorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(302)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(303)
					p.Match(FoxySheepParserVERTICALSEPARATOR)
				}
				{
					p.SetState(304)
					p.expr(10)
				}

			case 54:
				localctx = NewThereforeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(305)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(306)
					p.Match(FoxySheepParserTHEREFORE)
				}
				{
					p.SetState(307)
					p.expr(8)
				}

			case 55:
				localctx = NewBecauseContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(308)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(309)
					p.Match(FoxySheepParserBECAUSE)
				}
				{
					p.SetState(310)
					p.expr(8)
				}

			case 56:
				localctx = NewSetContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(311)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				p.SetState(312)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserEQUAL || (((_la-103)&-(0x1f+1)) == 0 && ((1<<uint((_la-103)))&((1<<(FoxySheepParserCARETCOLONEQUAL-103))|(1<<(FoxySheepParserCOLONEQUAL-103))|(1<<(FoxySheepParserCARETEQUAL-103))|(1<<(FoxySheepParserFUNCTIONARROW-103)))) != 0)) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(313)
					p.expr(5)
				}

			case 57:
				localctx = NewMessageContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(314)

				if !(p.Precpred(p.GetParserRuleContext(), 87)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 87)", ""))
				}
				{
					p.SetState(315)
					p.Match(FoxySheepParserDOUBLECOLON)
				}
				{
					p.SetState(316)
					p.Match(FoxySheepParserStringLiteral)
				}
				p.SetState(319)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(317)
						p.Match(FoxySheepParserDOUBLECOLON)
					}
					{
						p.SetState(318)
						p.Match(FoxySheepParserStringLiteral)
					}

				}

			case 58:
				localctx = NewHeadExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(321)

				if !(p.Precpred(p.GetParserRuleContext(), 78)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 78)", ""))
				}
				{
					p.SetState(322)
					p.Match(FoxySheepParserLBRACKET)
				}
				{
					p.SetState(323)
					p.ExpressionList()
				}
				{
					p.SetState(324)
					p.Match(FoxySheepParserRBRACKET)
				}

			case 59:
				localctx = NewAccessorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(326)

				if !(p.Precpred(p.GetParserRuleContext(), 77)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 77)", ""))
				}
				{
					p.SetState(327)
					p.AccessExpression()
				}

			case 60:
				localctx = NewIncrementContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(328)

				if !(p.Precpred(p.GetParserRuleContext(), 75)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 75)", ""))
				}
				p.SetState(329)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserDOUBLEPLUS || _la == FoxySheepParserDOUBLEMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}

			case 61:
				localctx = NewFactorialContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(330)

				if !(p.Precpred(p.GetParserRuleContext(), 68)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 68)", ""))
				}
				p.SetState(331)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserDOUBLEBANG || _la == FoxySheepParserBANG) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}

			case 62:
				localctx = NewConjugateContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(332)

				if !(p.Precpred(p.GetParserRuleContext(), 67)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 67)", ""))
				}
				p.SetState(333)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-132)&-(0x1f+1)) == 0 && ((1<<uint((_la-132)))&((1<<(FoxySheepParserTRANSPOSE-132))|(1<<(FoxySheepParserCONJUGATETRANSPOSE-132))|(1<<(FoxySheepParserHERMITIANCONJUGATE-132))|(1<<(FoxySheepParserCONJUGATE-132)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}

			case 63:
				localctx = NewDerivativeContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(334)

				if !(p.Precpred(p.GetParserRuleContext(), 66)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 66)", ""))
				}
				p.SetState(336)
				p.GetErrorHandler().Sync(p)
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(335)
							p.Match(FoxySheepParserSINGLEQUOTE)
						}

					default:
						panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					}

					p.SetState(338)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
				}

			case 64:
				localctx = NewSpanAContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(340)

				if !(p.Precpred(p.GetParserRuleContext(), 36)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 36)", ""))
				}
				{
					p.SetState(341)
					p.Match(FoxySheepParserDOUBLESEMICOLON)
				}
				p.SetState(343)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(342)
						p.expr(0)
					}

				}

			case 65:
				localctx = NewRepeatedContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				p.SetState(346)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserTRIPPLEDOT || _la == FoxySheepParserDOUBLEDOT) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}

			case 66:
				localctx = NewFunctionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(347)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(348)
					p.Match(FoxySheepParserAMP)
				}

			case 67:
				localctx = NewUnsetContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(349)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(350)
					p.Match(FoxySheepParserEQUAL)
				}
				{
					p.SetState(351)
					p.Match(FoxySheepParserDOT)
				}

			case 68:
				localctx = NewPutContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(352)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				p.SetState(353)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserTRIPPLEGREATER || _la == FoxySheepParserDOUBLEGREATER) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(354)
					p.Match(FoxySheepParserStringLiteral)
				}

			case 69:
				localctx = NewCompoundExpressionContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_expr)
				p.SetState(355)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(356)
					p.Match(FoxySheepParserSEMICOLON)
				}
				p.SetState(358)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(357)
						p.expr(0)
					}

				}

			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) CopyFrom(ctx *SymbolContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ContextNameContext struct {
	*SymbolContext
}

func NewContextNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContextNameContext {
	var p = new(ContextNameContext)

	p.SymbolContext = NewEmptySymbolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SymbolContext))

	return p
}

func (s *ContextNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextNameContext) Name() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserName, 0)
}

func (s *ContextNameContext) Context() IContextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContextContext)
}

func (s *ContextNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterContextName(s)
	}
}

func (s *ContextNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitContextName(s)
	}
}

func (s *ContextNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitContextName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FoxySheepParserRULE_symbol)

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

	localctx = NewContextNameContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.Context()
		}

	}
	{
		p.SetState(368)
		p.Match(FoxySheepParserName)
	}

	return localctx
}

// IContextContext is an interface to support dynamic dispatch.
type IContextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContextContext differentiates from other interfaces.
	IsContextContext()
}

type ContextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContextContext() *ContextContext {
	var p = new(ContextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_context
	return p
}

func (*ContextContext) IsContextContext() {}

func NewContextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContextContext {
	var p = new(ContextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_context

	return p
}

func (s *ContextContext) GetParser() antlr.Parser { return s.parser }

func (s *ContextContext) CopyFrom(ctx *ContextContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SimpleContextContext struct {
	*ContextContext
}

func NewSimpleContextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SimpleContextContext {
	var p = new(SimpleContextContext)

	p.ContextContext = NewEmptyContextContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ContextContext))

	return p
}

func (s *SimpleContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleContextContext) Name() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserName, 0)
}

func (s *SimpleContextContext) AllBACKQUOTE() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserBACKQUOTE)
}

func (s *SimpleContextContext) BACKQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBACKQUOTE, i)
}

func (s *SimpleContextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSimpleContext(s)
	}
}

func (s *SimpleContextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSimpleContext(s)
	}
}

func (s *SimpleContextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSimpleContext(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompoundContextContext struct {
	*ContextContext
}

func NewCompoundContextContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompoundContextContext {
	var p = new(CompoundContextContext)

	p.ContextContext = NewEmptyContextContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ContextContext))

	return p
}

func (s *CompoundContextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundContextContext) AllName() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserName)
}

func (s *CompoundContextContext) Name(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserName, i)
}

func (s *CompoundContextContext) AllBACKQUOTE() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserBACKQUOTE)
}

func (s *CompoundContextContext) BACKQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBACKQUOTE, i)
}

func (s *CompoundContextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterCompoundContext(s)
	}
}

func (s *CompoundContextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitCompoundContext(s)
	}
}

func (s *CompoundContextContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitCompoundContext(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) Context() (localctx IContextContext) {
	localctx = NewContextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FoxySheepParserRULE_context)
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

	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSimpleContextContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FoxySheepParserBACKQUOTE {
			{
				p.SetState(370)
				p.Match(FoxySheepParserBACKQUOTE)
			}

		}
		{
			p.SetState(373)
			p.Match(FoxySheepParserName)
		}
		{
			p.SetState(374)
			p.Match(FoxySheepParserBACKQUOTE)
		}

	case 2:
		localctx = NewCompoundContextContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(FoxySheepParserName)
		}
		{
			p.SetState(376)
			p.Match(FoxySheepParserBACKQUOTE)
		}
		{
			p.SetState(377)
			p.Match(FoxySheepParserName)
		}
		{
			p.SetState(378)
			p.Match(FoxySheepParserBACKQUOTE)
		}

	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) CopyFrom(ctx *NumberLiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NumberBaseNContext struct {
	*NumberLiteralContext
}

func NewNumberBaseNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberBaseNContext {
	var p = new(NumberBaseNContext)

	p.NumberLiteralContext = NewEmptyNumberLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberLiteralContext))

	return p
}

func (s *NumberBaseNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberBaseNContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIGITS, 0)
}

func (s *NumberBaseNContext) NumberInBase() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserNumberInBase, 0)
}

func (s *NumberBaseNContext) NumberLiteralPrecision() INumberLiteralPrecisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralPrecisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralPrecisionContext)
}

func (s *NumberBaseNContext) NumberLiteralExponent() INumberLiteralExponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralExponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralExponentContext)
}

func (s *NumberBaseNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNumberBaseN(s)
	}
}

func (s *NumberBaseNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNumberBaseN(s)
	}
}

func (s *NumberBaseNContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNumberBaseN(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberBaseTenContext struct {
	*NumberLiteralContext
}

func NewNumberBaseTenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberBaseTenContext {
	var p = new(NumberBaseTenContext)

	p.NumberLiteralContext = NewEmptyNumberLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NumberLiteralContext))

	return p
}

func (s *NumberBaseTenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberBaseTenContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIGITS, 0)
}

func (s *NumberBaseTenContext) DecimalNumber() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDecimalNumber, 0)
}

func (s *NumberBaseTenContext) NumberLiteralPrecision() INumberLiteralPrecisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralPrecisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralPrecisionContext)
}

func (s *NumberBaseTenContext) NumberLiteralExponent() INumberLiteralExponentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralExponentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralExponentContext)
}

func (s *NumberBaseTenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNumberBaseTen(s)
	}
}

func (s *NumberBaseTenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNumberBaseTen(s)
	}
}

func (s *NumberBaseTenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNumberBaseTen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) NumberLiteral() (localctx INumberLiteralContext) {
	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FoxySheepParserRULE_numberLiteral)
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

	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNumberBaseNContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(381)
			p.Match(FoxySheepParserDIGITS)
		}
		{
			p.SetState(382)
			p.Match(FoxySheepParserNumberInBase)
		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(383)
				p.NumberLiteralPrecision()
			}

		}
		p.SetState(387)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(386)
				p.NumberLiteralExponent()
			}

		}

	case 2:
		localctx = NewNumberBaseTenContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(389)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserDecimalNumber || _la == FoxySheepParserDIGITS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(390)
				p.NumberLiteralPrecision()
			}

		}
		p.SetState(394)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(393)
				p.NumberLiteralExponent()
			}

		}

	}

	return localctx
}

// INumberLiteralPrecisionContext is an interface to support dynamic dispatch.
type INumberLiteralPrecisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralPrecisionContext differentiates from other interfaces.
	IsNumberLiteralPrecisionContext()
}

type NumberLiteralPrecisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralPrecisionContext() *NumberLiteralPrecisionContext {
	var p = new(NumberLiteralPrecisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_numberLiteralPrecision
	return p
}

func (*NumberLiteralPrecisionContext) IsNumberLiteralPrecisionContext() {}

func NewNumberLiteralPrecisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralPrecisionContext {
	var p = new(NumberLiteralPrecisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_numberLiteralPrecision

	return p
}

func (s *NumberLiteralPrecisionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralPrecisionContext) DOUBLEBACKQUOTE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEBACKQUOTE, 0)
}

func (s *NumberLiteralPrecisionContext) DecimalNumber() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDecimalNumber, 0)
}

func (s *NumberLiteralPrecisionContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIGITS, 0)
}

func (s *NumberLiteralPrecisionContext) BACKQUOTE() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBACKQUOTE, 0)
}

func (s *NumberLiteralPrecisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralPrecisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralPrecisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNumberLiteralPrecision(s)
	}
}

func (s *NumberLiteralPrecisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNumberLiteralPrecision(s)
	}
}

func (s *NumberLiteralPrecisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNumberLiteralPrecision(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) NumberLiteralPrecision() (localctx INumberLiteralPrecisionContext) {
	localctx = NewNumberLiteralPrecisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FoxySheepParserRULE_numberLiteralPrecision)
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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FoxySheepParserDOUBLEBACKQUOTE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)
			p.Match(FoxySheepParserDOUBLEBACKQUOTE)
		}
		p.SetState(399)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserDecimalNumber || _la == FoxySheepParserDIGITS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case FoxySheepParserBACKQUOTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(400)
			p.Match(FoxySheepParserBACKQUOTE)
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
			p.SetState(401)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FoxySheepParserDecimalNumber || _la == FoxySheepParserDIGITS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberLiteralExponentContext is an interface to support dynamic dispatch.
type INumberLiteralExponentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralExponentContext differentiates from other interfaces.
	IsNumberLiteralExponentContext()
}

type NumberLiteralExponentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralExponentContext() *NumberLiteralExponentContext {
	var p = new(NumberLiteralExponentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_numberLiteralExponent
	return p
}

func (*NumberLiteralExponentContext) IsNumberLiteralExponentContext() {}

func NewNumberLiteralExponentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralExponentContext {
	var p = new(NumberLiteralExponentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_numberLiteralExponent

	return p
}

func (s *NumberLiteralExponentContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralExponentContext) ASTERISKCARET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserASTERISKCARET, 0)
}

func (s *NumberLiteralExponentContext) DIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDIGITS, 0)
}

func (s *NumberLiteralExponentContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPLUS, 0)
}

func (s *NumberLiteralExponentContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserMINUS, 0)
}

func (s *NumberLiteralExponentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralExponentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralExponentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterNumberLiteralExponent(s)
	}
}

func (s *NumberLiteralExponentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitNumberLiteralExponent(s)
	}
}

func (s *NumberLiteralExponentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitNumberLiteralExponent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) NumberLiteralExponent() (localctx INumberLiteralExponentContext) {
	localctx = NewNumberLiteralExponentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FoxySheepParserRULE_numberLiteralExponent)
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
		p.SetState(406)
		p.Match(FoxySheepParserASTERISKCARET)
	}
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FoxySheepParserPLUS || _la == FoxySheepParserMINUS {
		p.SetState(407)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FoxySheepParserPLUS || _la == FoxySheepParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(410)
		p.Match(FoxySheepParserDIGITS)
	}

	return localctx
}

// IOutExpressionContext is an interface to support dynamic dispatch.
type IOutExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOutExpressionContext differentiates from other interfaces.
	IsOutExpressionContext()
}

type OutExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOutExpressionContext() *OutExpressionContext {
	var p = new(OutExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_outExpression
	return p
}

func (*OutExpressionContext) IsOutExpressionContext() {}

func NewOutExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OutExpressionContext {
	var p = new(OutExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_outExpression

	return p
}

func (s *OutExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OutExpressionContext) CopyFrom(ctx *OutExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OutExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OutUnnumberedContext struct {
	*OutExpressionContext
}

func NewOutUnnumberedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OutUnnumberedContext {
	var p = new(OutUnnumberedContext)

	p.OutExpressionContext = NewEmptyOutExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OutExpressionContext))

	return p
}

func (s *OutUnnumberedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutUnnumberedContext) PERCENTS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPERCENTS, 0)
}

func (s *OutUnnumberedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOutUnnumbered(s)
	}
}

func (s *OutUnnumberedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOutUnnumbered(s)
	}
}

func (s *OutUnnumberedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOutUnnumbered(s)

	default:
		return t.VisitChildren(s)
	}
}

type OutNumberedContext struct {
	*OutExpressionContext
}

func NewOutNumberedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OutNumberedContext {
	var p = new(OutNumberedContext)

	p.OutExpressionContext = NewEmptyOutExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OutExpressionContext))

	return p
}

func (s *OutNumberedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OutNumberedContext) PERCENTDIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserPERCENTDIGITS, 0)
}

func (s *OutNumberedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterOutNumbered(s)
	}
}

func (s *OutNumberedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitOutNumbered(s)
	}
}

func (s *OutNumberedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitOutNumbered(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) OutExpression() (localctx IOutExpressionContext) {
	localctx = NewOutExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FoxySheepParserRULE_outExpression)

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

	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FoxySheepParserPERCENTDIGITS:
		localctx = NewOutNumberedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(412)
			p.Match(FoxySheepParserPERCENTDIGITS)
		}

	case FoxySheepParserPERCENTS:
		localctx = NewOutUnnumberedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(413)
			p.Match(FoxySheepParserPERCENTS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISlotExpressionContext is an interface to support dynamic dispatch.
type ISlotExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlotExpressionContext differentiates from other interfaces.
	IsSlotExpressionContext()
}

type SlotExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlotExpressionContext() *SlotExpressionContext {
	var p = new(SlotExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_slotExpression
	return p
}

func (*SlotExpressionContext) IsSlotExpressionContext() {}

func NewSlotExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlotExpressionContext {
	var p = new(SlotExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_slotExpression

	return p
}

func (s *SlotExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SlotExpressionContext) CopyFrom(ctx *SlotExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SlotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SlotDigitsContext struct {
	*SlotExpressionContext
}

func NewSlotDigitsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotDigitsContext {
	var p = new(SlotDigitsContext)

	p.SlotExpressionContext = NewEmptySlotExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SlotExpressionContext))

	return p
}

func (s *SlotDigitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotDigitsContext) HASHDIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserHASHDIGITS, 0)
}

func (s *SlotDigitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlotDigits(s)
	}
}

func (s *SlotDigitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlotDigits(s)
	}
}

func (s *SlotDigitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlotDigits(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlotSequenceDigitsContext struct {
	*SlotExpressionContext
}

func NewSlotSequenceDigitsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotSequenceDigitsContext {
	var p = new(SlotSequenceDigitsContext)

	p.SlotExpressionContext = NewEmptySlotExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SlotExpressionContext))

	return p
}

func (s *SlotSequenceDigitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotSequenceDigitsContext) DOUBLEHASHDIGITS() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEHASHDIGITS, 0)
}

func (s *SlotSequenceDigitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlotSequenceDigits(s)
	}
}

func (s *SlotSequenceDigitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlotSequenceDigits(s)
	}
}

func (s *SlotSequenceDigitsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlotSequenceDigits(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlotNamedContext struct {
	*SlotExpressionContext
}

func NewSlotNamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotNamedContext {
	var p = new(SlotNamedContext)

	p.SlotExpressionContext = NewEmptySlotExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SlotExpressionContext))

	return p
}

func (s *SlotNamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotNamedContext) HASHStringLiteral() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserHASHStringLiteral, 0)
}

func (s *SlotNamedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlotNamed(s)
	}
}

func (s *SlotNamedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlotNamed(s)
	}
}

func (s *SlotNamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlotNamed(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlotSequenceContext struct {
	*SlotExpressionContext
}

func NewSlotSequenceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotSequenceContext {
	var p = new(SlotSequenceContext)

	p.SlotExpressionContext = NewEmptySlotExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SlotExpressionContext))

	return p
}

func (s *SlotSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotSequenceContext) DOUBLEHASH() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserDOUBLEHASH, 0)
}

func (s *SlotSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlotSequence(s)
	}
}

func (s *SlotSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlotSequence(s)
	}
}

func (s *SlotSequenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlotSequence(s)

	default:
		return t.VisitChildren(s)
	}
}

type SlotUnnamedContext struct {
	*SlotExpressionContext
}

func NewSlotUnnamedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SlotUnnamedContext {
	var p = new(SlotUnnamedContext)

	p.SlotExpressionContext = NewEmptySlotExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SlotExpressionContext))

	return p
}

func (s *SlotUnnamedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlotUnnamedContext) HASH() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserHASH, 0)
}

func (s *SlotUnnamedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterSlotUnnamed(s)
	}
}

func (s *SlotUnnamedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitSlotUnnamed(s)
	}
}

func (s *SlotUnnamedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitSlotUnnamed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) SlotExpression() (localctx ISlotExpressionContext) {
	localctx = NewSlotExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FoxySheepParserRULE_slotExpression)

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

	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FoxySheepParserHASHDIGITS:
		localctx = NewSlotDigitsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(FoxySheepParserHASHDIGITS)
		}

	case FoxySheepParserHASHStringLiteral:
		localctx = NewSlotNamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Match(FoxySheepParserHASHStringLiteral)
		}

	case FoxySheepParserDOUBLEHASHDIGITS:
		localctx = NewSlotSequenceDigitsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(418)
			p.Match(FoxySheepParserDOUBLEHASHDIGITS)
		}

	case FoxySheepParserDOUBLEHASH:
		localctx = NewSlotSequenceContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(419)
			p.Match(FoxySheepParserDOUBLEHASH)
		}

	case FoxySheepParserHASH:
		localctx = NewSlotUnnamedContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(420)
			p.Match(FoxySheepParserHASH)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) CopyFrom(ctx *ExpressionListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionListedContext struct {
	*ExpressionListContext
}

func NewExpressionListedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionListedContext {
	var p = new(ExpressionListedContext)

	p.ExpressionListContext = NewEmptyExpressionListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionListContext))

	return p
}

func (s *ExpressionListedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListedContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExpressionListedContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionListedContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserCOMMA)
}

func (s *ExpressionListedContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserCOMMA, i)
}

func (s *ExpressionListedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterExpressionListed(s)
	}
}

func (s *ExpressionListedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitExpressionListed(s)
	}
}

func (s *ExpressionListedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitExpressionListed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FoxySheepParserRULE_expressionList)
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

	localctx = NewExpressionListedContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
		{
			p.SetState(423)
			p.expr(0)
		}

	}
	p.SetState(432)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FoxySheepParserCOMMA {
		{
			p.SetState(426)
			p.Match(FoxySheepParserCOMMA)
		}
		p.SetState(428)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
			{
				p.SetState(427)
				p.expr(0)
			}

		}

		p.SetState(434)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAccessExpressionContext is an interface to support dynamic dispatch.
type IAccessExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessExpressionContext differentiates from other interfaces.
	IsAccessExpressionContext()
}

type AccessExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessExpressionContext() *AccessExpressionContext {
	var p = new(AccessExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_accessExpression
	return p
}

func (*AccessExpressionContext) IsAccessExpressionContext() {}

func NewAccessExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessExpressionContext {
	var p = new(AccessExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_accessExpression

	return p
}

func (s *AccessExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessExpressionContext) CopyFrom(ctx *AccessExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccessExpressionBContext struct {
	*AccessExpressionContext
}

func NewAccessExpressionBContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessExpressionBContext {
	var p = new(AccessExpressionBContext)

	p.AccessExpressionContext = NewEmptyAccessExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AccessExpressionContext))

	return p
}

func (s *AccessExpressionBContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessExpressionBContext) LBARBRACKET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLBARBRACKET, 0)
}

func (s *AccessExpressionBContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AccessExpressionBContext) RBARBRACKET() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRBARBRACKET, 0)
}

func (s *AccessExpressionBContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAccessExpressionB(s)
	}
}

func (s *AccessExpressionBContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAccessExpressionB(s)
	}
}

func (s *AccessExpressionBContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAccessExpressionB(s)

	default:
		return t.VisitChildren(s)
	}
}

type AccessExpressionAContext struct {
	*AccessExpressionContext
}

func NewAccessExpressionAContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessExpressionAContext {
	var p = new(AccessExpressionAContext)

	p.AccessExpressionContext = NewEmptyAccessExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AccessExpressionContext))

	return p
}

func (s *AccessExpressionAContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessExpressionAContext) AllLBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserLBRACKET)
}

func (s *AccessExpressionAContext) LBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserLBRACKET, i)
}

func (s *AccessExpressionAContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *AccessExpressionAContext) AllRBRACKET() []antlr.TerminalNode {
	return s.GetTokens(FoxySheepParserRBRACKET)
}

func (s *AccessExpressionAContext) RBRACKET(i int) antlr.TerminalNode {
	return s.GetToken(FoxySheepParserRBRACKET, i)
}

func (s *AccessExpressionAContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterAccessExpressionA(s)
	}
}

func (s *AccessExpressionAContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitAccessExpressionA(s)
	}
}

func (s *AccessExpressionAContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitAccessExpressionA(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) AccessExpression() (localctx IAccessExpressionContext) {
	localctx = NewAccessExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FoxySheepParserRULE_accessExpression)

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

	p.SetState(445)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FoxySheepParserLBRACKET:
		localctx = NewAccessExpressionAContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(FoxySheepParserLBRACKET)
		}
		{
			p.SetState(436)
			p.Match(FoxySheepParserLBRACKET)
		}
		{
			p.SetState(437)
			p.ExpressionList()
		}
		{
			p.SetState(438)
			p.Match(FoxySheepParserRBRACKET)
		}
		{
			p.SetState(439)
			p.Match(FoxySheepParserRBRACKET)
		}

	case FoxySheepParserLBARBRACKET:
		localctx = NewAccessExpressionBContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(441)
			p.Match(FoxySheepParserLBARBRACKET)
		}
		{
			p.SetState(442)
			p.ExpressionList()
		}
		{
			p.SetState(443)
			p.Match(FoxySheepParserRBARBRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoxContext is an interface to support dynamic dispatch.
type IBoxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoxContext differentiates from other interfaces.
	IsBoxContext()
}

type BoxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoxContext() *BoxContext {
	var p = new(BoxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FoxySheepParserRULE_box
	return p
}

func (*BoxContext) IsBoxContext() {}

func NewBoxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoxContext {
	var p = new(BoxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FoxySheepParserRULE_box

	return p
}

func (s *BoxContext) GetParser() antlr.Parser { return s.parser }

func (s *BoxContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BoxContext) BoxLeftBoxParenthesis() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxLeftBoxParenthesis, 0)
}

func (s *BoxContext) BoxRightBoxParenthesis() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxRightBoxParenthesis, 0)
}

func (s *BoxContext) AllBox() []IBoxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoxContext)(nil)).Elem())
	var tst = make([]IBoxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoxContext)
		}
	}

	return tst
}

func (s *BoxContext) Box(i int) IBoxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoxContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoxContext)
}

func (s *BoxContext) InterpretedBox() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserInterpretedBox, 0)
}

func (s *BoxContext) BoxSqrt() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxSqrt, 0)
}

func (s *BoxContext) BoxOtherscript() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxOtherscript, 0)
}

func (s *BoxContext) BoxUnderscript() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxUnderscript, 0)
}

func (s *BoxContext) BoxOverscript() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxOverscript, 0)
}

func (s *BoxContext) BoxSubscript() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxSubscript, 0)
}

func (s *BoxContext) BoxFraction() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserBoxFraction, 0)
}

func (s *BoxContext) FormBox() antlr.TerminalNode {
	return s.GetToken(FoxySheepParserFormBox, 0)
}

func (s *BoxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.EnterBox(s)
	}
}

func (s *BoxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FoxySheepListener); ok {
		listenerT.ExitBox(s)
	}
}

func (s *BoxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FoxySheepVisitor:
		return t.VisitBox(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FoxySheepParser) Box() (localctx IBoxContext) {
	return p.box(0)
}

func (p *FoxySheepParser) box(_p int) (localctx IBoxContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoxContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoxContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, FoxySheepParserRULE_box, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(448)
			p.expr(0)
		}

	case 2:
		{
			p.SetState(449)
			p.Match(FoxySheepParserBoxLeftBoxParenthesis)
		}
		p.SetState(453)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserInterpretedBox-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxSqrt-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
			{
				p.SetState(450)
				p.box(0)
			}

			p.SetState(455)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(456)
			p.Match(FoxySheepParserBoxRightBoxParenthesis)
		}

	case 3:
		{
			p.SetState(457)
			p.Match(FoxySheepParserInterpretedBox)
		}
		{
			p.SetState(458)
			p.Match(FoxySheepParserBoxLeftBoxParenthesis)
		}
		p.SetState(460)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FoxySheepParserName)|(1<<FoxySheepParserDecimalNumber)|(1<<FoxySheepParserDIGITS)|(1<<FoxySheepParserStringLiteral)|(1<<FoxySheepParserLPAREN)|(1<<FoxySheepParserLBRACE)|(1<<FoxySheepParserLANGLE)|(1<<FoxySheepParserLFLOOR)|(1<<FoxySheepParserLCEILING)|(1<<FoxySheepParserLBRACKETINGBAR)|(1<<FoxySheepParserLDOUBLEBRACKETINGBAR)|(1<<FoxySheepParserBACKQUOTE))) != 0) || _la == FoxySheepParserDOUBLELESS || _la == FoxySheepParserNOT || (((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(FoxySheepParserTRIPPLEBLANK-77))|(1<<(FoxySheepParserDOUBLEBLANK-77))|(1<<(FoxySheepParserBLANKDOT-77))|(1<<(FoxySheepParserBLANK-77))|(1<<(FoxySheepParserPERCENTDIGITS-77))|(1<<(FoxySheepParserPERCENTS-77))|(1<<(FoxySheepParserHASHDIGITS-77))|(1<<(FoxySheepParserHASHStringLiteral-77))|(1<<(FoxySheepParserDOUBLEHASHDIGITS-77))|(1<<(FoxySheepParserDOUBLEHASH-77))|(1<<(FoxySheepParserHASH-77)))) != 0) || (((_la-112)&-(0x1f+1)) == 0 && ((1<<uint((_la-112)))&((1<<(FoxySheepParserDOUBLEPLUS-112))|(1<<(FoxySheepParserDOUBLEMINUS-112))|(1<<(FoxySheepParserBANG-112))|(1<<(FoxySheepParserINTEGRAL-112))|(1<<(FoxySheepParserDEL-112))|(1<<(FoxySheepParserSQUARE-112)))) != 0) || (((_la-158)&-(0x1f+1)) == 0 && ((1<<uint((_la-158)))&((1<<(FoxySheepParserPLUS-158))|(1<<(FoxySheepParserMINUS-158))|(1<<(FoxySheepParserPLUSMINUS-158))|(1<<(FoxySheepParserMINUSPLUS-158))|(1<<(FoxySheepParserInterpretedBox-158))|(1<<(FoxySheepParserBoxLeftBoxParenthesis-158))|(1<<(FoxySheepParserBoxSqrt-158))|(1<<(FoxySheepParserBoxConstructor-158))|(1<<(FoxySheepParserSPANSEMICOLONS-158)))) != 0) {
			{
				p.SetState(459)
				p.box(0)
			}

			p.SetState(462)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(464)
			p.Match(FoxySheepParserBoxRightBoxParenthesis)
		}

	case 4:
		{
			p.SetState(466)
			p.Match(FoxySheepParserBoxSqrt)
		}
		{
			p.SetState(467)
			p.box(0)
		}
		{
			p.SetState(468)
			p.Match(FoxySheepParserBoxOtherscript)
		}
		{
			p.SetState(469)
			p.box(3)
		}

	case 5:
		{
			p.SetState(471)
			p.Match(FoxySheepParserBoxSqrt)
		}
		{
			p.SetState(472)
			p.box(2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(517)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(475)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(476)
					p.Match(FoxySheepParserBoxUnderscript)
				}
				{
					p.SetState(477)
					p.box(0)
				}
				{
					p.SetState(478)
					p.Match(FoxySheepParserBoxOtherscript)
				}
				{
					p.SetState(479)
					p.box(14)
				}

			case 2:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(481)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(482)
					p.Match(FoxySheepParserBoxOverscript)
				}
				{
					p.SetState(483)
					p.box(0)
				}
				{
					p.SetState(484)
					p.Match(FoxySheepParserBoxOtherscript)
				}
				{
					p.SetState(485)
					p.box(13)
				}

			case 3:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(487)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				p.SetState(488)
				_la = p.GetTokenStream().LA(1)

				if !(_la == FoxySheepParserBoxOverscript || _la == FoxySheepParserBoxUnderscript) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(489)
					p.box(11)
				}

			case 4:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(490)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(491)
					p.Match(FoxySheepParserBoxSubscript)
				}
				{
					p.SetState(492)
					p.box(0)
				}
				{
					p.SetState(493)
					p.Match(FoxySheepParserBoxOtherscript)
				}
				{
					p.SetState(494)
					p.box(11)
				}

			case 5:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(496)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(497)
					p.Match(FoxySheepParserBoxSubscript)
				}
				{
					p.SetState(498)
					p.box(10)
				}

			case 6:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(499)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(500)
					p.Match(FoxySheepParserBoxOverscript)
				}
				{
					p.SetState(501)
					p.box(8)
				}

			case 7:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(502)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(503)
					p.Match(FoxySheepParserBoxUnderscript)
				}
				{
					p.SetState(504)
					p.box(0)
				}
				{
					p.SetState(505)
					p.Match(FoxySheepParserBoxOtherscript)
				}
				{
					p.SetState(506)
					p.box(7)
				}

			case 8:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(508)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(509)
					p.Match(FoxySheepParserBoxUnderscript)
				}
				{
					p.SetState(510)
					p.box(6)
				}

			case 9:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(511)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(512)
					p.Match(FoxySheepParserBoxFraction)
				}
				{
					p.SetState(513)
					p.box(5)
				}

			case 10:
				localctx = NewBoxContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, FoxySheepParserRULE_box)
				p.SetState(514)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(515)
					p.Match(FoxySheepParserFormBox)
				}
				{
					p.SetState(516)
					p.box(2)
				}

			}

		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

func (p *FoxySheepParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	case 11:
		var t *BoxContext = nil
		if localctx != nil {
			t = localctx.(*BoxContext)
		}
		return p.Box_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FoxySheepParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 79)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 73)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 72)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 71)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 70)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 69)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 65)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 64)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 60)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 59)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 58)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 57)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 56)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 54)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 53)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 52)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 51)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 50)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 49)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 48)

	case 20:
		return p.Precpred(p.GetParserRuleContext(), 47)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 46)

	case 22:
		return p.Precpred(p.GetParserRuleContext(), 45)

	case 23:
		return p.Precpred(p.GetParserRuleContext(), 44)

	case 24:
		return p.Precpred(p.GetParserRuleContext(), 43)

	case 25:
		return p.Precpred(p.GetParserRuleContext(), 42)

	case 26:
		return p.Precpred(p.GetParserRuleContext(), 41)

	case 27:
		return p.Precpred(p.GetParserRuleContext(), 40)

	case 28:
		return p.Precpred(p.GetParserRuleContext(), 39)

	case 29:
		return p.Precpred(p.GetParserRuleContext(), 38)

	case 30:
		return p.Precpred(p.GetParserRuleContext(), 37)

	case 31:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 32:
		return p.Precpred(p.GetParserRuleContext(), 33)

	case 33:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 34:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 35:
		return p.Precpred(p.GetParserRuleContext(), 29)

	case 36:
		return p.Precpred(p.GetParserRuleContext(), 28)

	case 37:
		return p.Precpred(p.GetParserRuleContext(), 27)

	case 38:
		return p.Precpred(p.GetParserRuleContext(), 26)

	case 39:
		return p.Precpred(p.GetParserRuleContext(), 25)

	case 40:
		return p.Precpred(p.GetParserRuleContext(), 24)

	case 41:
		return p.Precpred(p.GetParserRuleContext(), 23)

	case 42:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 43:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 44:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 45:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 46:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 47:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 48:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 49:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 50:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 51:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 52:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 53:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 54:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 55:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 56:
		return p.Precpred(p.GetParserRuleContext(), 87)

	case 57:
		return p.Precpred(p.GetParserRuleContext(), 78)

	case 58:
		return p.Precpred(p.GetParserRuleContext(), 77)

	case 59:
		return p.Precpred(p.GetParserRuleContext(), 75)

	case 60:
		return p.Precpred(p.GetParserRuleContext(), 68)

	case 61:
		return p.Precpred(p.GetParserRuleContext(), 67)

	case 62:
		return p.Precpred(p.GetParserRuleContext(), 66)

	case 63:
		return p.Precpred(p.GetParserRuleContext(), 36)

	case 64:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 65:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 66:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 67:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 68:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FoxySheepParser) Box_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 69:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 70:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 71:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 72:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 73:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 74:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 75:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 76:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 77:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 78:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
