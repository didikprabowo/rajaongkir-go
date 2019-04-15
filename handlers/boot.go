package handlers

import (
	"os"
)

// OS
var TYPERJ string = os.Getenv("TYPERJ")
var URLProvince string = os.Getenv("URLProvince")
var URLCity string = os.Getenv("URLCity")
var URLCost string = os.Getenv("URLCost")
var key string = os.Getenv("key")

//HARDCODE
// const TYPERJ string = "starter"
// const URLProvince string = "https://api.rajaongkir.com/" + TYPERJ + "/province"
// const URLCity string = "https://api.rajaongkir.com/" + TYPERJ + "/city"
// const URLCost string = "https://api.rajaongkir.com/" + TYPERJ + "/cost"
// const key string = "6141096f7c8fed1a3226cc03ac3dab0b"
