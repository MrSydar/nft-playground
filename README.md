# svg-combiner
This tool allows you to create all possible permutations of your svg assets. For the example results, look into the `example/output` directory

# How to use

## Prepare assets

Create a main directory with all the needed resources. This directory should contain `template.svg` file and multiple 
subdirectories for each asset you want to combine.

Template file there should contain placeholders. Example of template file:
```
<svg width="800" height="600" xmlns="http://www.w3.org/2000/svg">
    <!-- body -->
    <!-- candy -->
    <!-- handle -->
</svg>
```

For each placeholder, there must be a corresponding assets directory with the same name.
For example, for `<!-- body -->` placeholder, there must be a `body` folder.

Next, put an arbitrary number of `svg` assets into each asset directory.
Assets should be enumerated from zero like `0.svg`, `1.svg`, etc.

### Example structrue:

    .
    ├── body                   # Directory with first type assets
    │   ├── 0.svg              # Actual assets
    │   ├── 1.svg
    │   └── ...
    ├── handle                 # Directory with second type assets
    │   ├── 0.svg
    │   ├── 1.svg
    │   └── ...
    ├── ...
    └── template.svg           # Template file
    
There is an example available in the `example/assets` directory of the current project.
 
## Run
Run the tool with the next parameters: 
- `a` : path to the asset directory
- `o` : path to the output directory where results will be generated

Example:
```
go run . -a ./example/assets/ice-cream -o ./example/output
``` 
