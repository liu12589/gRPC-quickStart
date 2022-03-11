package sample

import (
	"github.com/google/uuid"
	"time"
	"useCpu/pb"
)
import "math/rand"

// init()函数，在执行包中任何其他代码之前都将被调用一次
func init(){
	rand.Seed(time.Now().UnixNano())
}
func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_AZERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_QWERTY
	}
}

func randomCPUBrand() string {
	return randomStringFormSet("Intel", "AMD")
}

func randomGPUBrand() string {
	return randomStringFormSet("NVIDIA", "AMD")
}

func randomStringFormSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFormSet("酷睿i9-12900", "G7400", "i5 10400F")
	} else {
		return randomStringFormSet("5950X", "5900X", "5600X")
	}
}

func randomGPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFormSet(
			"RTX 2060",
			"RTX 2070",
			"RTX 1660-Ti",
			"RTX 1070",
		)
	} else {
		return randomStringFormSet(
			"RX 590",
			"RX 580",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomId() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFormSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFormSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFormSet("Latitude", "Vostro", "XPS", "Alienware")
	default:
		return randomStringFormSet("Thinkpad x1", "Thinkpad p1", "Thinkpad p53")
	}
}
