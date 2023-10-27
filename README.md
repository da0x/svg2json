# SVG to JSON CLI Tool

This command-line tool allows you to convert an SVG file to JSON format.

## Installation

To install the tool, make sure you have Go installed on your system. Then, run the following command:

```bash
go install github.com/da0x/svg2json
```

## Usage
After installation, you can convert an SVG file to JSON using the following command:

```bash
svg2json -input input.svg -output output.json
```
Replace input.svg with the path to your SVG file.
Replace output.json with the desired name for the output JSON file. If not specified, it will default to the input filename with the extension changed to .json.
The tool will parse the SVG file and save the JSON representation to the output file.

## Example
Suppose you have an SVG file named example.svg that you want to convert to JSON:

```bash
svg2json -input example.svg
```
This will generate a JSON file named example.json containing the SVG data in the following JSON format.
```json
{
  "api_version": "v2",
  "kind": "document",
  "spec": {
    "attributes": {
      "height": "800px",
      "viewBox": "0 0 24 24",
      "width": "800px",
      "xmlns": "http://www.w3.org/2000/svg"
    },
    "children": [
      {
        "attributes": {
          "d": "M1 3v18h22V3zm1 1h20v16H2zm17 6H5V9h14zm-6 4H5v-1h8z"
        },
        "content": "",
        "name": "path"
      },
      {
        "attributes": {
          "d": "M0 0h24v24H0z",
          "fill": "none"
        },
        "content": "",
        "name": "path"
      }
    ],
    "name": "svg"
  },
  "type": "svg"
}
```

## Contribution
Feel free to contribute directly by opening an issue or a pull request.

