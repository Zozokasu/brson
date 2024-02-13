package golang_brson

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/andybalholm/brotli"
	"gopkg.in/mgo.v2/bson"
	"io"
)

// FrDT00003
var header = []byte{0x46, 0x72, 0x44, 0x54, 0x00, 0x00, 0x00, 0x00, 0x03}

func DecodeBrson(r []byte) ([]byte, error) {
	if !bytes.Equal(r[:9], header) {
		return nil, errors.New("Use brson")
	}
	reader := brotli.NewReader(bytes.NewReader(r[9:]))
	return io.ReadAll(reader)
}

func ReadBrson(r []byte) (interface{}, error) {
	data, err := DecodeBrson(r)
	if err != nil {
		return nil, err
	}
	var val interface{}
	err = bson.Unmarshal(data, &val)
	if err != nil {
		return nil, err
	}
	return val, err
}

func DecodeBrsonJson(r []byte) ([]byte, error) {
	val, err := ReadBrson(r)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(val)
	return res, err
}

func EncodeJsonBrson(j []byte) ([]byte, error) {
	var val interface{}
	err := json.Unmarshal(j, &val)
	if err != nil {
		return nil, err
	}
	res, err := bson.Marshal(val)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	brWriter := brotli.NewWriterLevel(&buf, 4)
	brWriter.Write(res)
	brWriter.Close()
	var brson []byte
	brson = append(header, buf.Bytes()...)
	return brson, nil

}
