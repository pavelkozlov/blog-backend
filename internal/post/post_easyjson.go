// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package post

import (
	json "encoding/json"

	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson5a72dc82DecodeBlogInternalPost(in *jlexer.Lexer, out *NewBlog) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "is_draft":
			out.IsDraft = bool(in.Bool())
		case "description":
			if in.IsNull() {
				in.Skip()
				out.Description = nil
			} else {
				if out.Description == nil {
					out.Description = new(string)
				}
				*out.Description = string(in.String())
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
func easyjson5a72dc82EncodeBlogInternalPost(out *jwriter.Writer, in NewBlog) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"is_draft\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDraft))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		if in.Description == nil {
			out.RawString("null")
		} else {
			out.String(string(*in.Description))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v NewBlog) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeBlogInternalPost(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v NewBlog) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeBlogInternalPost(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *NewBlog) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeBlogInternalPost(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *NewBlog) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeBlogInternalPost(l, v)
}
func easyjson5a72dc82DecodeBlogInternalPost1(in *jlexer.Lexer, out *Blogs) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Blogs, 0, 8)
			} else {
				*out = Blogs{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 *Blog
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Blog)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson5a72dc82EncodeBlogInternalPost1(out *jwriter.Writer, in Blogs) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
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
}

// MarshalJSON supports json.Marshaler interface
func (v Blogs) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeBlogInternalPost1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Blogs) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeBlogInternalPost1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Blogs) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeBlogInternalPost1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Blogs) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeBlogInternalPost1(l, v)
}
func easyjson5a72dc82DecodeBlogInternalPost2(in *jlexer.Lexer, out *Blog) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "title":
			out.Title = string(in.String())
		case "slug":
			out.Slug = string(in.String())
		case "body":
			out.Body = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "is_draft":
			out.IsDraft = bool(in.Bool())
		case "author_id":
			out.AuthorId = string(in.String())
		case "created_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated_at":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UpdatedAt).UnmarshalJSON(data))
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
func easyjson5a72dc82EncodeBlogInternalPost2(out *jwriter.Writer, in Blog) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix[1:])
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"slug\":"
		out.RawString(prefix)
		out.String(string(in.Slug))
	}
	{
		const prefix string = ",\"body\":"
		out.RawString(prefix)
		out.String(string(in.Body))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"is_draft\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsDraft))
	}
	{
		const prefix string = ",\"author_id\":"
		out.RawString(prefix)
		out.String(string(in.AuthorId))
	}
	{
		const prefix string = ",\"created_at\":"
		out.RawString(prefix)
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated_at\":"
		out.RawString(prefix)
		out.Raw((in.UpdatedAt).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Blog) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson5a72dc82EncodeBlogInternalPost2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Blog) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson5a72dc82EncodeBlogInternalPost2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Blog) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson5a72dc82DecodeBlogInternalPost2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Blog) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson5a72dc82DecodeBlogInternalPost2(l, v)
}
