package mime_types

/*
The MIT License (MIT)

Copyright (c) 2017 Leonid Plyushch <leonid.plyushch@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

import (
	"mime"
)

func initVideoMime() {
	mime.AddExtensionType(".3ga", "video/3gpp")
	mime.AddExtensionType(".3gpp", "video/3gpp")
	mime.AddExtensionType(".3gp", "video/3gpp")
	mime.AddExtensionType(".asf", "video/x-ms-asf")
	mime.AddExtensionType(".asr", "video/x-ms-asf")
	mime.AddExtensionType(".asx", "video/x-ms-asf")
	mime.AddExtensionType(".avi", "video/avi")
	mime.AddExtensionType(".divx", "video/x-msvideo")
	mime.AddExtensionType(".dv", "video/x-dv")
	mime.AddExtensionType(".f4v", "video/mp4")
	mime.AddExtensionType(".lsf", "video/x-la-asf")
	mime.AddExtensionType(".lsx", "video/x-la-asf")
	mime.AddExtensionType(".m1v", "video/mpeg")
	mime.AddExtensionType(".m2v", "video/mpeg")
	mime.AddExtensionType(".m4v", "video/mp4")
	mime.AddExtensionType(".mov", "video/quicktime")
	mime.AddExtensionType(".mp2", "video/mpeg")
	mime.AddExtensionType(".mp4", "video/mp4")
	mime.AddExtensionType(".mpa", "video/mpeg")
	mime.AddExtensionType(".mpeg", "video/mpeg")
	mime.AddExtensionType(".mpe", "video/mpeg")
	mime.AddExtensionType(".mpg", "video/mpeg")
	mime.AddExtensionType(".mpv2", "video/mpeg")
	mime.AddExtensionType(".mqv", "video/quicktime")
	mime.AddExtensionType(".ogv", "video/ogg")
	mime.AddExtensionType(".qt", "video/quicktime")
	mime.AddExtensionType(".rv", "video/vnd.rn-realvideo")
	mime.AddExtensionType(".webm", "video/webm")
	mime.AddExtensionType(".wmv", "video/x-ms-wmv")
}
