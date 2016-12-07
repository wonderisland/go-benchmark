// AUTOGENERATED FILE: easyjson marshaller/unmarshallers.

package go_benchmark

import (
	json "encoding/json"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ = json.RawMessage{}
	_ = jlexer.Lexer{}
	_ = jwriter.Writer{}
)

func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark(in *jlexer.Lexer, out *LargePayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				if out.Topics == nil {
					out.Topics = new(DSTopicsList)
				}
				(*out.Topics).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark(out *jwriter.Writer, in LargePayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"topics\":")
	if in.Topics == nil {
		out.RawString("null")
	} else {
		(*in.Topics).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LargePayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LargePayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LargePayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LargePayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark1(in *jlexer.Lexer, out *DSTopicsList) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "topics":
			if in.IsNull() {
				in.Skip()
				out.Topics = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Topics = make([]*DSTopic, 0, 8)
				} else {
					out.Topics = []*DSTopic{}
				}
				for !in.IsDelim(']') {
					var v1 *DSTopic
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(DSTopic)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Topics = append(out.Topics, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark1(out *jwriter.Writer, in DSTopicsList) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"topics\":")
	if in.Topics == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in.Topics {
			if v2 > 0 {
				out.RawByte(',')
			}
			if v3 == nil {
				out.RawString("null")
			} else {
				(*v3).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DSTopicsList) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DSTopicsList) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DSTopicsList) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DSTopicsList) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark1(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark2(in *jlexer.Lexer, out *DSTopic) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark2(out *jwriter.Writer, in DSTopic) {
	out.RawByte('{')
	first := true
	_ = first
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DSTopic) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DSTopic) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DSTopic) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DSTopic) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark2(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark3(in *jlexer.Lexer, out *DSUser) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Username":
			out.Username = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark3(out *jwriter.Writer, in DSUser) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"Username\":")
	out.String(string(in.Username))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DSUser) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DSUser) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DSUser) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DSUser) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark3(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark4(in *jlexer.Lexer, out *MediumPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "person":
			if in.IsNull() {
				in.Skip()
				out.Person = nil
			} else {
				if out.Person == nil {
					out.Person = new(CBPerson)
				}
				(*out.Person).UnmarshalEasyJSON(in)
			}
		case "compnay":
			out.Company = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark4(out *jwriter.Writer, in MediumPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"person\":")
	if in.Person == nil {
		out.RawString("null")
	} else {
		(*in.Person).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"compnay\":")
	out.String(string(in.Company))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v MediumPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MediumPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *MediumPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MediumPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark4(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark5(in *jlexer.Lexer, out *CBPerson) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			if in.IsNull() {
				in.Skip()
				out.Name = nil
			} else {
				if out.Name == nil {
					out.Name = new(CBName)
				}
				(*out.Name).UnmarshalEasyJSON(in)
			}
		case "github":
			if in.IsNull() {
				in.Skip()
				out.Github = nil
			} else {
				if out.Github == nil {
					out.Github = new(CBGithub)
				}
				(*out.Github).UnmarshalEasyJSON(in)
			}
		case "gravatar":
			if in.IsNull() {
				in.Skip()
				out.Gravatar = nil
			} else {
				if out.Gravatar == nil {
					out.Gravatar = new(CBGravatar)
				}
				(*out.Gravatar).UnmarshalEasyJSON(in)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark5(out *jwriter.Writer, in CBPerson) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"name\":")
	if in.Name == nil {
		out.RawString("null")
	} else {
		(*in.Name).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"github\":")
	if in.Github == nil {
		out.RawString("null")
	} else {
		(*in.Github).MarshalEasyJSON(out)
	}
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"gravatar\":")
	if in.Gravatar == nil {
		out.RawString("null")
	} else {
		(*in.Gravatar).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBPerson) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBPerson) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBPerson) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBPerson) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark5(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark6(in *jlexer.Lexer, out *CBName) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "fullName":
			out.FullName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark6(out *jwriter.Writer, in CBName) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"fullName\":")
	out.String(string(in.FullName))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBName) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBName) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBName) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBName) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark6(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark7(in *jlexer.Lexer, out *CBGithub) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "followers":
			out.Followers = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark7(out *jwriter.Writer, in CBGithub) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"followers\":")
	out.Int(int(in.Followers))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBGithub) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBGithub) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBGithub) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBGithub) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark7(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark8(in *jlexer.Lexer, out *CBGravatar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "avatars":
			if in.IsNull() {
				in.Skip()
				out.Avatars = nil
			} else {
				in.Delim('[')
				if !in.IsDelim(']') {
					out.Avatars = make([]*CBAvatar, 0, 8)
				} else {
					out.Avatars = []*CBAvatar{}
				}
				for !in.IsDelim(']') {
					var v4 *CBAvatar
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(CBAvatar)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Avatars = append(out.Avatars, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark8(out *jwriter.Writer, in CBGravatar) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"avatars\":")
	if in.Avatars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v5, v6 := range in.Avatars {
			if v5 > 0 {
				out.RawByte(',')
			}
			if v6 == nil {
				out.RawString("null")
			} else {
				(*v6).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBGravatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBGravatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBGravatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBGravatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark8(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark9(in *jlexer.Lexer, out *CBAvatar) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.Url = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark9(out *jwriter.Writer, in CBAvatar) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"url\":")
	out.String(string(in.Url))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CBAvatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CBAvatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CBAvatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CBAvatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark9(l, v)
}
func easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark10(in *jlexer.Lexer, out *SmallPayload) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "st":
			out.St = int(in.Int())
		case "uuid":
			out.Uuid = string(in.String())
		case "ua":
			out.Ua = string(in.String())
		case "tz":
			out.Tz = int(in.Int())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark10(out *jwriter.Writer, in SmallPayload) {
	out.RawByte('{')
	first := true
	_ = first
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"st\":")
	out.Int(int(in.St))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"uuid\":")
	out.String(string(in.Uuid))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"ua\":")
	out.String(string(in.Ua))
	if !first {
		out.RawByte(',')
	}
	first = false
	out.RawString("\"tz\":")
	out.Int(int(in.Tz))
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SmallPayload) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SmallPayload) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeGithubComJsonIteratorGoBenchmark10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SmallPayload) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SmallPayload) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeGithubComJsonIteratorGoBenchmark10(l, v)
}