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

type FpPayload struct {
	DomController bool    `json:"p"`
	Seed          float64 `json:"S"`
	GpuVendor     struct {
		UnmaskedVendorWebgl   string `json:"v"`
		UnmaskedRendererWebgl string `json:"r"`
	} `json:"w"`
	Selenium  bool `json:"s"`
	Headless  bool `json:"h"`
	Devtools  bool `json:"b"`
	Devtools2 bool `json:"d"`
}

type Payload struct {
	Arg1      float64 `json:"b"`
	Rand      float64 `json:"v"`
	Signature string  `json:"e"`
	Fp        string  `json:"s"`
	Arg2      float64 `json:"d"`
	Version   string  `json:"vr"`
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

func buildFp(key string, seed float64) (string, error) {
	gpu := gpus[rand.Intn(len(gpus))]
	fp := FpPayload{
		DomController: false,
		Seed:          seed,
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

	return Encrypt(key, fp)
}

func BuildPayload(ctx *ScriptCtx) (*Payload, error) {
	fp, err := buildFp(ctx.key, ctx.seed)
	if err != nil {
		return nil, err
	}
	return &Payload{
		Arg1:      ctx.arg1,
		Rand:      ctx.rand,
		Signature: ctx.signature,
		Fp:        fp,
		Arg2:      ctx.arg2,
		Version:   ctx.version,
	}, nil
}
