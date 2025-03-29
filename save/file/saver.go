package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kiamev/ffpr-save-cypher/rijndael"
	oj "github.com/virtuald/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/config"
)

func SaveSave(game global.Game, data *save.Data, slot int, dir string, file string, saveType global.SaveFileType) (err error) {
	var (
		out []byte
	)

	if out, err = oj.Marshal(data); err != nil {
		return
	}

	if err = saveFile(game, out, dir, file, data.Trimmed, saveType); err == nil {
		// if data.BestiaryDataInternal != nil && len(*data.BestiaryDataInternal) > 0 {
		// 	var trim []byte
		// 	if data.BestiaryDataTrim != nil && len(*data.BestiaryDataTrim) > 0 {
		// 		trim = *data.BestiaryDataTrim
		// 	}
		// 	_ = saveFile(game, *data.BestiaryDataInternal, dir, bestiaryFile, trim, saveType)
		// }
		err = nil
	}
	return
}

func saveFile(game global.Game, data []byte, dir string, file string, trimmed []byte, saveType global.SaveFileType) (err error) {
	var (
		b  bytes.Buffer
		zw *flate.Writer
	)
	printFile(filepath.Join(config.Dir(game), "_save.json"), data)
	if saveType == global.PC {
		// Flate
		if zw, err = flate.NewWriter(&b, 6); err != nil {
			return
		}
		if _, err = zw.Write(data); err != nil {
			return
		}
		_ = zw.Flush()
		_ = zw.Close()

		// Encrypt
		if data, err = rijndael.New().Encrypt(b.Bytes()); err != nil {
			return
		}

		// Encode
		s := base64.StdEncoding.EncodeToString(data)

		// Format
		if len(trimmed) > 0 {
			data = make([]byte, 0, len(trimmed)+len(s))
			data = append(data, trimmed...)
			data = append(data, []byte(s)...)
		} else {
			data = []byte(s)
		}
	}

	// Write to file
	toFile := filepath.Join(dir, file)
	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	if err = os.WriteFile(toFile, data, 0777); err != nil {
		return fmt.Errorf("failed to write save file %s: %v", toFile, err)
	}
	return
}
