package telegram_tree

import (
	"fmt"
)

func ReplaceSymbols(symToNum map[string]int) error {
	if len(symToNum) == 0 {
		return fmt.Errorf("empty map")
	}
	numToSym := make(map[int]string, len(symToNum))
	for sym, num := range symToNum {
		numToSym[num] = sym
	}
	if len(symToNum) != len(numToSym) {
		return fmt.Errorf("invalid map")
	}
	symbolToNum = symToNum
	numToSymbol = numToSym
	return nil
}

var symbolToNum = map[string]int{
	"È": 0,
	"É": 1,
	"Ê": 2,
	"Ë": 3,
	"Ì": 4,
	"Í": 5,
	"Î": 6,
	"Ï": 7,
	"Ð": 8,
	"Ñ": 9,
	"Ò": 10,
	"Ó": 11,
	"Ô": 12,
	"Õ": 13,
	"Ö": 14,
	"×": 15,
	"Ø": 16,
	"Ù": 17,
	"Ú": 18,
	"Û": 19,
	"Ü": 20,
	"Ý": 21,
	"Þ": 22,
	"ß": 23,
	"à": 24,
	"á": 25,
	"â": 26,
	"ã": 27,
	"ä": 28,
	"å": 29,
	"æ": 30,
	"ç": 31,
	"è": 32,
	"é": 33,
	"ê": 34,
	"ë": 35,
	"ì": 36,
	"í": 37,
	"î": 38,
	"ï": 39,
	"ð": 40,
	"ñ": 41,
	"ò": 42,
	"ó": 43,
	"ô": 44,
	"õ": 45,
	"ö": 46,
	"÷": 47,
	"ø": 48,
	"ù": 49,
	"ú": 50,
	"û": 51,
	"ü": 52,
	"ý": 53,
	"þ": 54,
	"ÿ": 55,
	"Ā": 56,
	"ā": 57,
	"Ă": 58,
	"ă": 59,
	"Ą": 60,
	"ą": 61,
	"Ć": 62,
	"ć": 63,
	"Ĉ": 64,
	"ĉ": 65,
	"Ċ": 66,
	"ċ": 67,
	"Č": 68,
	"č": 69,
	"Ď": 70,
	"ď": 71,
	"Đ": 72,
	"đ": 73,
	"Ē": 74,
	"ē": 75,
	"Ĕ": 76,
	"ĕ": 77,
	"Ė": 78,
	"ė": 79,
	"Ę": 80,
	"ę": 81,
	"Ě": 82,
	"ě": 83,
	"Ĝ": 84,
	"ĝ": 85,
	"Ğ": 86,
	"ğ": 87,
	"Ġ": 88,
	"ġ": 89,
	"Ģ": 90,
	"ģ": 91,
	"Ĥ": 92,
	"ĥ": 93,
	"Ħ": 94,
	"ħ": 95,
	"Ĩ": 96,
	"ĩ": 97,
	"Ī": 98,
	"ī": 99,
	"Ĭ": 100,
	"ĭ": 101,
	"Į": 102,
	"į": 103,
	"İ": 104,
	"ı": 105,
	"Ĳ": 106,
	"ĳ": 107,
	"Ĵ": 108,
	"ĵ": 109,
	"Ķ": 110,
	"ķ": 111,
	"ĸ": 112,
	"Ĺ": 113,
	"ĺ": 114,
	"Ļ": 115,
	"ļ": 116,
	"Ľ": 117,
	"ľ": 118,
	"Ŀ": 119,
	"ŀ": 120,
	"Ł": 121,
	"ł": 122,
	"Ń": 123,
	"ń": 124,
	"Ņ": 125,
	"ņ": 126,
	"Ň": 127,
	"ň": 128,
	"ŉ": 129,
	"Ŋ": 130,
	"ŋ": 131,
	"Ō": 132,
	"ō": 133,
	"Ŏ": 134,
	"ŏ": 135,
	"Ő": 136,
	"ő": 137,
	"Œ": 138,
	"œ": 139,
	"Ŕ": 140,
	"ŕ": 141,
	"Ŗ": 142,
	"ŗ": 143,
	"Ř": 144,
	"ř": 145,
	"Ś": 146,
	"ś": 147,
	"Ŝ": 148,
	"ŝ": 149,
	"Ş": 150,
	"ş": 151,
	"Š": 152,
	"š": 153,
	"Ţ": 154,
	"ţ": 155,
	"Ť": 156,
	"ť": 157,
	"Ŧ": 158,
	"ŧ": 159,
	"Ũ": 160,
	"ũ": 161,
	"Ū": 162,
	"ū": 163,
	"Ŭ": 164,
	"ŭ": 165,
	"Ů": 166,
	"ů": 167,
	"Ű": 168,
	"ű": 169,
	"Ų": 170,
	"ų": 171,
	"Ŵ": 172,
	"ŵ": 173,
	"Ŷ": 174,
	"ŷ": 175,
	"Ÿ": 176,
	"Ź": 177,
	"ź": 178,
	"Ż": 179,
	"ż": 180,
	"Ž": 181,
	"ž": 182,
	"ſ": 183,
	"ƀ": 184,
	"Ɓ": 185,
	"Ƃ": 186,
	"ƃ": 187,
	"Ƅ": 188,
	"ƅ": 189,
	"Ɔ": 190,
	"Ƈ": 191,
	"ƈ": 192,
	"Ɖ": 193,
	"Ɗ": 194,
	"Ƌ": 195,
	"ƌ": 196,
	"ƍ": 197,
	"Ǝ": 198,
	"Ə": 199,
	"Ɛ": 200,
	"Ƒ": 201,
	"ƒ": 202,
	"Ɠ": 203,
	"Ɣ": 204,
	"ƕ": 205,
	"Ɩ": 206,
	"Ɨ": 207,
	"Ƙ": 208,
	"ƙ": 209,
	"ƚ": 210,
	"ƛ": 211,
	"Ɯ": 212,
	"Ɲ": 213,
	"ƞ": 214,
	"Ɵ": 215,
	"Ơ": 216,
	"ơ": 217,
	"Ƣ": 218,
	"ƣ": 219,
	"Ƥ": 220,
	"ƥ": 221,
	"Ʀ": 222,
	"Ƨ": 223,
	"ƨ": 224,
	"Ʃ": 225,
	"ƪ": 226,
	"ƫ": 227,
	"Ƭ": 228,
	"ƭ": 229,
	"Ʈ": 230,
	"Ư": 231,
	"ư": 232,
	"Ʊ": 233,
	"Ʋ": 234,
	"Ƴ": 235,
	"ƴ": 236,
	"Ƶ": 237,
	"ƶ": 238,
	"Ʒ": 239,
	"Ƹ": 240,
	"ƹ": 241,
	"ƺ": 242,
	"ƻ": 243,
	"Ƽ": 244,
	"ƽ": 245,
	"ƾ": 246,
	"ƿ": 247,
	"ǀ": 248,
	"ǁ": 249,
	"ǂ": 250,
	"ǃ": 251,
	"Ǆ": 252,
	"ǅ": 253,
	"ǆ": 254,
	"Ǉ": 255,
	"ǈ": 256,
	"ǉ": 257,
	"Ǌ": 258,
	"ǋ": 259,
	"ǌ": 260,
	"Ǎ": 261,
	"ǎ": 262,
	"Ǐ": 263,
	"ǐ": 264,
	"Ǒ": 265,
	"ǒ": 266,
	"Ǔ": 267,
	"ǔ": 268,
	"Ǖ": 269,
	"ǖ": 270,
	"Ǘ": 271,
	"ǘ": 272,
	"Ǚ": 273,
	"ǚ": 274,
	"Ǜ": 275,
	"ǜ": 276,
	"ǝ": 277,
	"Ǟ": 278,
	"ǟ": 279,
	"Ǡ": 280,
	"ǡ": 281,
	"Ǣ": 282,
	"ǣ": 283,
	"Ǥ": 284,
	"ǥ": 285,
	"Ǧ": 286,
	"ǧ": 287,
	"Ǩ": 288,
	"ǩ": 289,
	"Ǫ": 290,
	"ǫ": 291,
	"Ǭ": 292,
	"ǭ": 293,
	"Ǯ": 294,
	"ǯ": 295,
	"ǰ": 296,
	"Ǳ": 297,
	"ǲ": 298,
	"ǳ": 299,
	"Ǵ": 300,
}
var numToSymbol = map[int]string{
	0:   "È",
	1:   "É",
	2:   "Ê",
	3:   "Ë",
	4:   "Ì",
	5:   "Í",
	6:   "Î",
	7:   "Ï",
	8:   "Ð",
	9:   "Ñ",
	10:  "Ò",
	11:  "Ó",
	12:  "Ô",
	13:  "Õ",
	14:  "Ö",
	15:  "×",
	16:  "Ø",
	17:  "Ù",
	18:  "Ú",
	19:  "Û",
	20:  "Ü",
	21:  "Ý",
	22:  "Þ",
	23:  "ß",
	24:  "à",
	25:  "á",
	26:  "â",
	27:  "ã",
	28:  "ä",
	29:  "å",
	30:  "æ",
	31:  "ç",
	32:  "è",
	33:  "é",
	34:  "ê",
	35:  "ë",
	36:  "ì",
	37:  "í",
	38:  "î",
	39:  "ï",
	40:  "ð",
	41:  "ñ",
	42:  "ò",
	43:  "ó",
	44:  "ô",
	45:  "õ",
	46:  "ö",
	47:  "÷",
	48:  "ø",
	49:  "ù",
	50:  "ú",
	51:  "û",
	52:  "ü",
	53:  "ý",
	54:  "þ",
	55:  "ÿ",
	56:  "Ā",
	57:  "ā",
	58:  "Ă",
	59:  "ă",
	60:  "Ą",
	61:  "ą",
	62:  "Ć",
	63:  "ć",
	64:  "Ĉ",
	65:  "ĉ",
	66:  "Ċ",
	67:  "ċ",
	68:  "Č",
	69:  "č",
	70:  "Ď",
	71:  "ď",
	72:  "Đ",
	73:  "đ",
	74:  "Ē",
	75:  "ē",
	76:  "Ĕ",
	77:  "ĕ",
	78:  "Ė",
	79:  "ė",
	80:  "Ę",
	81:  "ę",
	82:  "Ě",
	83:  "ě",
	84:  "Ĝ",
	85:  "ĝ",
	86:  "Ğ",
	87:  "ğ",
	88:  "Ġ",
	89:  "ġ",
	90:  "Ģ",
	91:  "ģ",
	92:  "Ĥ",
	93:  "ĥ",
	94:  "Ħ",
	95:  "ħ",
	96:  "Ĩ",
	97:  "ĩ",
	98:  "Ī",
	99:  "ī",
	100: "Ĭ",
	101: "ĭ",
	102: "Į",
	103: "į",
	104: "İ",
	105: "ı",
	106: "Ĳ",
	107: "ĳ",
	108: "Ĵ",
	109: "ĵ",
	110: "Ķ",
	111: "ķ",
	112: "ĸ",
	113: "Ĺ",
	114: "ĺ",
	115: "Ļ",
	116: "ļ",
	117: "Ľ",
	118: "ľ",
	119: "Ŀ",
	120: "ŀ",
	121: "Ł",
	122: "ł",
	123: "Ń",
	124: "ń",
	125: "Ņ",
	126: "ņ",
	127: "Ň",
	128: "ň",
	129: "ŉ",
	130: "Ŋ",
	131: "ŋ",
	132: "Ō",
	133: "ō",
	134: "Ŏ",
	135: "ŏ",
	136: "Ő",
	137: "ő",
	138: "Œ",
	139: "œ",
	140: "Ŕ",
	141: "ŕ",
	142: "Ŗ",
	143: "ŗ",
	144: "Ř",
	145: "ř",
	146: "Ś",
	147: "ś",
	148: "Ŝ",
	149: "ŝ",
	150: "Ş",
	151: "ş",
	152: "Š",
	153: "š",
	154: "Ţ",
	155: "ţ",
	156: "Ť",
	157: "ť",
	158: "Ŧ",
	159: "ŧ",
	160: "Ũ",
	161: "ũ",
	162: "Ū",
	163: "ū",
	164: "Ŭ",
	165: "ŭ",
	166: "Ů",
	167: "ů",
	168: "Ű",
	169: "ű",
	170: "Ų",
	171: "ų",
	172: "Ŵ",
	173: "ŵ",
	174: "Ŷ",
	175: "ŷ",
	176: "Ÿ",
	177: "Ź",
	178: "ź",
	179: "Ż",
	180: "ż",
	181: "Ž",
	182: "ž",
	183: "ſ",
	184: "ƀ",
	185: "Ɓ",
	186: "Ƃ",
	187: "ƃ",
	188: "Ƅ",
	189: "ƅ",
	190: "Ɔ",
	191: "Ƈ",
	192: "ƈ",
	193: "Ɖ",
	194: "Ɗ",
	195: "Ƌ",
	196: "ƌ",
	197: "ƍ",
	198: "Ǝ",
	199: "Ə",
	200: "Ɛ",
	201: "Ƒ",
	202: "ƒ",
	203: "Ɠ",
	204: "Ɣ",
	205: "ƕ",
	206: "Ɩ",
	207: "Ɨ",
	208: "Ƙ",
	209: "ƙ",
	210: "ƚ",
	211: "ƛ",
	212: "Ɯ",
	213: "Ɲ",
	214: "ƞ",
	215: "Ɵ",
	216: "Ơ",
	217: "ơ",
	218: "Ƣ",
	219: "ƣ",
	220: "Ƥ",
	221: "ƥ",
	222: "Ʀ",
	223: "Ƨ",
	224: "ƨ",
	225: "Ʃ",
	226: "ƪ",
	227: "ƫ",
	228: "Ƭ",
	229: "ƭ",
	230: "Ʈ",
	231: "Ư",
	232: "ư",
	233: "Ʊ",
	234: "Ʋ",
	235: "Ƴ",
	236: "ƴ",
	237: "Ƶ",
	238: "ƶ",
	239: "Ʒ",
	240: "Ƹ",
	241: "ƹ",
	242: "ƺ",
	243: "ƻ",
	244: "Ƽ",
	245: "ƽ",
	246: "ƾ",
	247: "ƿ",
	248: "ǀ",
	249: "ǁ",
	250: "ǂ",
	251: "ǃ",
	252: "Ǆ",
	253: "ǅ",
	254: "ǆ",
	255: "Ǉ",
	256: "ǈ",
	257: "ǉ",
	258: "Ǌ",
	259: "ǋ",
	260: "ǌ",
	261: "Ǎ",
	262: "ǎ",
	263: "Ǐ",
	264: "ǐ",
	265: "Ǒ",
	266: "ǒ",
	267: "Ǔ",
	268: "ǔ",
	269: "Ǖ",
	270: "ǖ",
	271: "Ǘ",
	272: "ǘ",
	273: "Ǚ",
	274: "ǚ",
	275: "Ǜ",
	276: "ǜ",
	277: "ǝ",
	278: "Ǟ",
	279: "ǟ",
	280: "Ǡ",
	281: "ǡ",
	282: "Ǣ",
	283: "ǣ",
	284: "Ǥ",
	285: "ǥ",
	286: "Ǧ",
	287: "ǧ",
	288: "Ǩ",
	289: "ǩ",
	290: "Ǫ",
	291: "ǫ",
	292: "Ǭ",
	293: "ǭ",
	294: "Ǯ",
	295: "ǯ",
	296: "ǰ",
	297: "Ǳ",
	298: "ǲ",
	299: "ǳ",
	300: "Ǵ",
}

func ConvertNumberToSymbol(in int) (string, error) {
	val, ok := numToSymbol[in]
	if ok {
		return val, nil
	}
	return "", fmt.Errorf("unsupported number")
}

func convertSymbolToNum(in string) (int, error) {
	val, ok := symbolToNum[in]
	if ok {
		return val, nil
	}
	return 0, fmt.Errorf("unsupported symbol")
}
