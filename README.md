palette
=======

Tiny HTTP API returning the RGB values used in an image. 

Works with GIF, JPEG and PNG.

## Usage

```
Usage of ./pal:
  -listen string
    	host:port to use (default ":8080")
  -max_size int
    	maximum size of image to process in bytes (default 1000000)
  -r_timeout duration
    	read timeout in seconds (default 1m0s)
  -w_timeout duration
    	write timeout in seconds (default 1m0s)
```

## The API

Return all colours used in an image:

```bash
curl --data-binary "@/path/to/image.png" "localhost:8080/"
```

Filter out colours that take up less than n pixels:

```bash
curl --data-binary "@/path/to/image.png" "localhost:8080/?min=100"
```

Output:

```json
[{"red":150,"green":206,"blue":180},{"red":255,"green":238,"blue":173},{"red":255,"green":111,"blue":105},{"red":255,"green":204,"blue":92},{"red":136,"green":216,"blue":176}]
```

Pretty print output:

```bash
curl --data-binary "@/path/to/image.png" "localhost:8080/?pretty=true"
```

Output:

```json
[
	{
		"red": 136,
		"green": 216,
		"blue": 176
	},
	{
		"red": 150,
		"green": 206,
		"blue": 180
	},
	{
		"red": 255,
		"green": 238,
		"blue": 173
	},
	{
		"red": 255,
		"green": 111,
		"blue": 105
	},
	{
		"red": 255,
		"green": 204,
		"blue": 92
	}
]
```
