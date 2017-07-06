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

func initImageMime() {
	mime.AddExtensionType(".arw", "image/x-sony-arw")
	mime.AddExtensionType(".bmp", "image/bmp")
	mime.AddExtensionType(".cmx", "image/x-cmx")
	mime.AddExtensionType(".cod", "image/cis-cod")
	mime.AddExtensionType(".cr2", "image/x-canon-cr2")
	mime.AddExtensionType(".crw", "image/x-canon-crw")
	mime.AddExtensionType(".dds", "image/x-dds")
	mime.AddExtensionType(".dib", "image/bmp")
	mime.AddExtensionType(".djv", "image/vnd.djvu")
	mime.AddExtensionType(".djvu", "image/vnd.djvu")
	mime.AddExtensionType(".dng", "image/x-adobe-dng")
	mime.AddExtensionType(".dwg", "image/x-dwg")
	mime.AddExtensionType(".fits", "image/fits")
	mime.AddExtensionType(".g3", "image/fax-g3")
	mime.AddExtensionType(".icns", "image/x-icns")
	mime.AddExtensionType(".ico", "image/x-ico")
	mime.AddExtensionType(".ief", "image/ief")
	mime.AddExtensionType(".jfif", "image/jpeg")
	mime.AddExtensionType(".jng", "image/x-jng")
	mime.AddExtensionType(".jp2", "image/jp2")
	mime.AddExtensionType(".jpeg", "image/jpeg")
	mime.AddExtensionType(".jpe", "image/jpeg")
	mime.AddExtensionType(".k25", "image/x-kodak-k25")
	mime.AddExtensionType(".kdc", "image/x-kodak-kdc")
	mime.AddExtensionType(".msod", "image/x-msod")
	mime.AddExtensionType(".nef", "image/x-nikon-nef")
	mime.AddExtensionType(".orf", "image/x-olympus-orf")
	mime.AddExtensionType(".pbm", "image/x-portable-bitmap")
	mime.AddExtensionType(".pcd", "image/x-photo-cd")
	mime.AddExtensionType(".pct", "image/x-pict")
	mime.AddExtensionType(".pcx", "image/vnd.zbrush.pcx")
	mime.AddExtensionType(".pef", "image/x-pentax-pef")
	mime.AddExtensionType(".pgm", "image/x-portable-graymap")
	mime.AddExtensionType(".pic", "image/x-pict")
	mime.AddExtensionType(".pict", "image/x-pict")
	mime.AddExtensionType(".pjpeg", "image/jpeg")
	mime.AddExtensionType(".pnm", "image/x-portable-anymap")
	mime.AddExtensionType(".ppm", "image/x-portable-pixmap")
	mime.AddExtensionType(".psd", "image/psd")
	mime.AddExtensionType(".ras", "image/x-cmu-raster")
	mime.AddExtensionType(".raw", "image/x-panasonic-raw")
	mime.AddExtensionType(".rgb", "image/x-rgb")
	mime.AddExtensionType(".sr2", "image/x-sony-sr2")
	mime.AddExtensionType(".srf", "image/x-sony-srf")
	mime.AddExtensionType(".svgz", "image/svg+xml-compressed")
	mime.AddExtensionType(".tga", "image/x-targa")
	mime.AddExtensionType(".tiff", "image/tiff")
	mime.AddExtensionType(".tif", "image/tiff")
	mime.AddExtensionType(".wmf", "image/x-wmf")
	mime.AddExtensionType(".x3f", "image/x-sigma-x3f")
	mime.AddExtensionType(".xbm", "image/x-xbitmap")
	mime.AddExtensionType(".xcf", "image/x-xcf")
	mime.AddExtensionType(".xpm", "image/x-xpixmap")
	mime.AddExtensionType(".xwd", "image/x-xwindowdump")
}
