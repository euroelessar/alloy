package filequeue

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Record) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Meta":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			if z.Meta == nil {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 string
				za0001, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				za0002, err = dc.ReadString()
				if err != nil {
					err = msgp.WrapError(err, "Meta", za0001)
					return
				}
				z.Meta[za0001] = za0002
			}
		case "Data":
			z.Data, err = dc.ReadBytes(z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Record) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 2
	// write "Meta"
	err = en.Append(0x82, 0xa4, 0x4d, 0x65, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.Meta)))
	if err != nil {
		err = msgp.WrapError(err, "Meta")
		return
	}
	for za0001, za0002 := range z.Meta {
		err = en.WriteString(za0001)
		if err != nil {
			err = msgp.WrapError(err, "Meta")
			return
		}
		err = en.WriteString(za0002)
		if err != nil {
			err = msgp.WrapError(err, "Meta", za0001)
			return
		}
	}
	// write "Data"
	err = en.Append(0xa4, 0x44, 0x61, 0x74, 0x61)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Data)
	if err != nil {
		err = msgp.WrapError(err, "Data")
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Record) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 2
	// string "Meta"
	o = append(o, 0x82, 0xa4, 0x4d, 0x65, 0x74, 0x61)
	o = msgp.AppendMapHeader(o, uint32(len(z.Meta)))
	for za0001, za0002 := range z.Meta {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendString(o, za0002)
	}
	// string "Data"
	o = append(o, 0xa4, 0x44, 0x61, 0x74, 0x61)
	o = msgp.AppendBytes(o, z.Data)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Record) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		err = msgp.WrapError(err)
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			err = msgp.WrapError(err)
			return
		}
		switch msgp.UnsafeString(field) {
		case "Meta":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				err = msgp.WrapError(err, "Meta")
				return
			}
			if z.Meta == nil {
				z.Meta = make(map[string]string, zb0002)
			} else if len(z.Meta) > 0 {
				for key := range z.Meta {
					delete(z.Meta, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 string
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta")
					return
				}
				za0002, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					err = msgp.WrapError(err, "Meta", za0001)
					return
				}
				z.Meta[za0001] = za0002
			}
		case "Data":
			z.Data, bts, err = msgp.ReadBytesBytes(bts, z.Data)
			if err != nil {
				err = msgp.WrapError(err, "Data")
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				err = msgp.WrapError(err)
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Record) Msgsize() (s int) {
	s = 1 + 5 + msgp.MapHeaderSize
	if z.Meta != nil {
		for za0001, za0002 := range z.Meta {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.StringPrefixSize + len(za0002)
		}
	}
	s += 5 + msgp.BytesPrefixSize + len(z.Data)
	return
}