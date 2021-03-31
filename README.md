## colorize

Simple library for applying color to terminal output
![colorized output](example/example_output.png)

## Use
```golang
colorizer := colorize.New(colornames.Firebrick, colornames.Mediumaquamarine)
colorizer.Println("I've got a firebrick foreground and medium aquamarine background!")
```
![colorized output](example/readme_example_output.png)

You can get back a color formatted string or byte slice directly as well.
A couple of the major print methods are supported, but open to suggestions
about additional functionality.
