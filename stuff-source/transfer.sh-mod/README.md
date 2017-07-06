# Transfer.sh
Easy file sharing from the command line

## Uploading files

```
$ curl --upload-file ./hello.txt https://transfer.sh/
$ http -vv --form https://transfer.sh/ file@./hello.txt
$ wget --method PUT --body-file=./hello.txt https://transfer.sh/ -O - -nv
PS H:\> invoke-webrequest -method put -infile .\hello.txt https://transfer.sh/hello.txt
```

For each uploaded file you will receive a link.
Example link to hello.txt: https://transfer.sh/66nb8/hello.txt

You can upload or download multiple files at once:
```
$ curl -i -F filedata=@./1.txt -F filedata=@./2.txt http://transfer.sh/
```

Download multiple files as tar.gz or zip archive:
```
$ curl 'http://transfer.sh/(15HKz/1.txt,Vkngk/2.txt).tar.gz'
$ curl 'http://transfer.sh/(15HKz/1.txt,Vkngk/2.txt).zip'
```

By default files are stored for 14 days and amount
of downloads is unlimited. You can override this with
the following HTTP headers:

 * Max-Days – to set maximum of days for storing file
 * Max-Downloads – to set maximum of downloads

Example:
```
$ curl -H "Max-Downloads: 1" -H "Max-Days: 5" --upload-file ./hello.txt http://transfer.sh/
```

## Creators

**Remco Verhoef**
- <https://twitter.com/remco_verhoef>
- <https://twitter.com/dutchcoders>

**Uvis Grinfelds**

## Copyright and license
Code and documentation copyright 2011-2014 Remco Verhoef. 
Code released under [the MIT license](LICENSE). 
