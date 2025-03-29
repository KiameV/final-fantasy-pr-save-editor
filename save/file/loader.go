package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/kiamev/ffpr-save-cypher/rijndael"
	oj "github.com/virtuald/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/config"
)

const bestiaryFile = "dp3fS2vqP7GDj8eF72YKqbT7FIAF=e7Shy2CsTITm2E="

func LoadSave(game global.Game, dir string, fileName string, saveType global.SaveFileType) (data *save.Data, err error) {
	var b []byte
	data = save.New(game)
	if b, data.Trimmed, err = loadFile(game, filepath.Join(dir, fileName), saveType, "save"); err != nil {
		return
	}
	err = oj.Unmarshal(b, data)
	// var t []byte
	// if b, t, err = loadFile(game, filepath.Join(dir, bestiaryFile), saveType, "bestiary"); err == nil {
	// 	data.BestiaryDataTrim = &t
	// 	data.BestiaryDataInternal = &b
	// }
	return
}

func loadFile(game global.Game, fromFile string, saveType global.SaveFileType, name string) (out []byte, trimmed []byte, err error) {
	var (
		b []byte
	)
	if b, err = os.ReadFile(fromFile); err != nil {
		return
	}
	if saveType == global.PS {
		return b, nil, nil
	}
	if len(b) < 10 {
		err = errors.New("unable to load file")
		return
	}
	// Format
	if b[0] == 239 && b[1] == 187 && b[2] == 191 {
		trimmed = []byte{239, 187, 191}
		b = b[3:]
	}
	for len(b)%4 != 0 {
		b = append(b, '=')
	}
	// Decode
	b, _ = base64.StdEncoding.DecodeString(string(b))
	if len(b) == 0 {
		err = errors.New("unable to load file")
		return
	}
	// Decrypt
	if b, err = rijndael.New().Decrypt(b); err != nil {
		return
	}

	// Flate
	println(fmt.Sprintf("%d", b[len(b)-1]))
	zr := flate.NewReader(bytes.NewReader(b))
	defer func() { _ = zr.Close() }()
	out, err = io.ReadAll(zr)
	if name == "bestiary" && strings.Contains(strings.ToLower(err.Error()), "eof") {
		err = nil
	}
	if err == nil {
		printFile(filepath.Join(config.Dir(game), "_"+name+"_loaded.json"), out)
	}
	return
}
