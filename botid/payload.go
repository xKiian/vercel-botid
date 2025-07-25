package botid

import (
	"encoding/json"
	"math/rand"
	"os"
)

type WebGL struct {
	WebGLUnmaskedVendor string `json:"webgl_unmasked_vendor"`
}
type GPUInfo struct {
	WebGLRenderer string  `json:"webgl_unmasked_renderer"`
	WebGL         []WebGL `json:"webgl"`
}

type EncryptedPayload struct {
	DomController bool    `json:"p"`
	RandomSeed    float64 `json:"S"`
	GpuVendor     struct {
		UnmaskedVendorWebgl   string `json:"v"`
		UnmaskedRendererWebgl string `json:"r"`
	} `json:"w"`
	Selenium  bool `json:"s"`
	Headless  bool `json:"h"`
	Devtools  bool `json:"b"`
	Devtools2 bool `json:"d"`
}

var gpus []GPUInfo

func init() {
	file, err := os.Open("webgl.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&gpus)
	if err != nil {
		panic(err)
	}
}

func BuildPayload(seed float64) EncryptedPayload {
	gpu := gpus[rand.Intn(len(gpus))]
	return EncryptedPayload{
		DomController: false,
		RandomSeed:    seed,
		GpuVendor: struct {
			UnmaskedVendorWebgl   string `json:"v"`
			UnmaskedRendererWebgl string `json:"r"`
		}{
			UnmaskedVendorWebgl:   gpu.WebGL[0].WebGLUnmaskedVendor,
			UnmaskedRendererWebgl: gpu.WebGLRenderer,
		},
		Selenium:  false,
		Headless:  false,
		Devtools:  false,
		Devtools2: false,
	}
}
