# SVG to JSON CLI Tool

This Go command-line tool allows you to convert SVG files to JSON format while processing all `.svg` files in a specified directory and its subdirectories.

## Usage

To use the tool, provide a directory path as a command-line argument. The tool will recursively search for `.svg` files in the specified directory and convert them to JSON format. If a JSON file already exists with the same name, it will be overwritten. The tool will print log messages indicating the conversion process.

```bash
svg2json directory_name
```

Replace directory_name with the path to the directory containing the .svg files you want to convert.
## Example
Suppose you have a directory named svg_files that contains multiple SVG files. To convert all the SVG files in this directory and its subdirectories to JSON, run the following command:

```bash
svg2json svg_files
```
The tool will process all .svg files and create corresponding .json files in the same location.

## Error Handling
If there are errors during the conversion process, the tool will print error messages in red.
