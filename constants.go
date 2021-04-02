package main

type videoMediaMap map[string]string
type audioMediaMap map[int]string
type dataMediaMap map[string]string
type muxMediaMap map[string]string
type interlaceModes map[string]string
type colorSpaces map[string]string
type transferCharacteristics map[string]string
type rawVideoCompName map[string]string
type sourceChannelSymbols map[string]string

var VideoMediaTypes = videoMediaMap{
	"raw":  "video/raw",
	"h264": "video/H264",
	"vc2":  "video/vc2",
}
var AudioMediaTypes = audioMediaMap{
	24: "audio/L24",
	20: "audio/L20",
	16: "audio/L16",
	8:  "audio/L8",
}
var DataMediaTypes = dataMediaMap{
	"smpte291": "video/smpte291",
	"json":     "application/json",
}
var MuxMediaTypes = muxMediaMap{
	"2022-6": "video/SMPTE-2022-6",
}

const VideoFormat = "urn:x-nmos:format:video"
const AudioFormat = "urn:x-nmos:format:audio"
const DataFormat = "urn:x-nmos:format:data"
const MuxFormat = "urn:x-nmos:format:mux"
const SenderTransport = "urn:x-nmos:transport:rtp.mcast"
const ReceiverTransport = "urn:x-nmos:transport:rtp"

var InterlaceModes = interlaceModes{
	"progressive":    "progressive",
	"interlaced_tff": "interlaced_tff",
	"interlaced_bff": "interlaced_bff",
	"interlaced_psf": "interlaced_psf",
}
var ColorSpaces = colorSpaces{
	"BT601":  "BT601",
	"BT709":  "BT709",
	"BT2020": "BT2020",
	"BT2100": "BT2100",
}
var TransferCharacteristics = transferCharacteristics{
	"SDR": "SDR",
	"HLG": "HLG",
	"PQ":  "PQ",
}
var RawVideoCompName = rawVideoCompName{
	"Y":        "Y",
	"Cb":       "Cb",
	"Cr":       "Cr",
	"I":        "I",
	"Ct":       "Ct",
	"Cp":       "Cp",
	"A":        "A",
	"R":        "R",
	"G":        "G",
	"B":        "B",
	"DepthMap": "DepthMap",
}
var SourceChannelSymbols = sourceChannelSymbols{
	"L":   "L",
	"R":   "R",
	"C":   "C",
	"LFE": "LFE",
	"Ls":  "Ls",
	"Rs":  "Rs",
	"Lss": "Lss",
	"Rss": "Rss",
	"Lrs": "Lrs",
	"Rrs": "Rrs",
	"Lc":  "Lc",
	"Rc":  "Rc",
	"Cs":  "Cs",
	"HI":  "HI",
	"VIN": "VIN",
	"M1":  "M1",
	"M2":  "M2",
	"Lt":  "Lt",
	"Rt":  "Rt",
	"Lst": "Lst",
	"Rst": "Rst",
	"S":   "S",
}
