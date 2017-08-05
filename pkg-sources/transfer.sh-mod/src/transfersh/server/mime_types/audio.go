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

package mime_types

import (
	"mime"
)

func initAudioMime() {
	mime.AddExtensionType(".aac", "audio/aac")
	mime.AddExtensionType(".ac3", "audio/ac3")
	mime.AddExtensionType(".aif", "audio/aiff")
	mime.AddExtensionType(".aifc", "audio/aiff")
	mime.AddExtensionType(".aiff", "audio/aiff")
	mime.AddExtensionType(".asx", "audio/x-ms-asx")
	mime.AddExtensionType(".au", "audio/basic")
	mime.AddExtensionType(".axa", "audio/annodex")
	mime.AddExtensionType(".f4b", "audio/x-m4b")
	mime.AddExtensionType(".flac", "audio/flac")
	mime.AddExtensionType(".it", "audio/x-it")
	mime.AddExtensionType(".m3u8", "audio/x-mpegurl")
	mime.AddExtensionType(".m3u", "audio/x-mpegurl")
	mime.AddExtensionType(".m4a", "audio/m4a")
	mime.AddExtensionType(".mid", "audio/midi")
	mime.AddExtensionType(".midi", "audio/midi")
	mime.AddExtensionType(".mp3", "audio/mpeg")
	mime.AddExtensionType(".mp+", "audio/x-musepack")
	mime.AddExtensionType(".oga", "audio/x-ogg")
	mime.AddExtensionType(".ogg", "audio/x-ogg")
	mime.AddExtensionType(".pls", "audio/x-mpegurl")
	mime.AddExtensionType(".psflib", "audio/x-psflib")
	mime.AddExtensionType(".ra", "audio/vnd.rn-realaudio")
	mime.AddExtensionType(".ram", "audio/vnd.rn-realaudio")
	mime.AddExtensionType(".rmi", "audio/mid")
	mime.AddExtensionType(".snd", "audio/basic")
	mime.AddExtensionType(".spx", "audio/x-speex")
	mime.AddExtensionType(".stm", "audio/x-stm")
	mime.AddExtensionType(".tta", "audio/x-tta")
	mime.AddExtensionType(".voc", "audio/x-voc")
	mime.AddExtensionType(".wav", "audio/wav")
	mime.AddExtensionType(".wma", "audio/x-ms-wma")
	mime.AddExtensionType(".wv", "audio/x-wavpack")
	mime.AddExtensionType(".wvc", "audio/x-wavpack-correction")
	mime.AddExtensionType(".wvp", "audio/x-wavpack")
}
