![solflower](https://github.com/realphant0m/SOLFLOWER/blob/master/solflower.png)

The SOLFLOWER CLI provides a simple way to generate a bird’s eye view of the entirety of a codebase. 

> SOLFLOWER is meant to mature into a tool that helps developers conceive and grasp the mental model of a smart contract project hastily, by providing a interactive tree with integrated human-centered control-flow graph analysis (step 2/4 development).


## Getting Started

1. Start the CLI: 
2. Specify codebase directory.
3. Specify output directory (or leave blank to use default)
4. Specify whether you want to open the generated flower in the browser instantly or not.
3. Enjoy!

```
"insert path"./SOLFLOWER.go
```

```
Available Flags:
  --path=<project-path>             Local path of the codebase.

  --url=<project-github-url>        URL path of the codebase. i.e. --url=github.com/janforys/30-in-30-solidity-challenge

  --output=<output-path>            Local path of the output visualization.
  
  --open-in-browser                 Optional choice of opening the output in the browser. 
```

### Prerequisites

To use this tool a few prerequisites are required.
1. Make sure Golang is installed.
2. Download this repository.
3. Install the required cloc executable. 
4. Build the binaries. 

```
git clone https://github.com/realphant0m/SOLFLOWER.git
npm install -g cloc
go build 
```


## Sample Visualization

![solflower](https://github.com/realphant0m/SOLFLOWER/blob/master/exampleflower.png)


## Built With

* [survey] (https://github.com/AlecAivazis/survey) - A golang library for building interactive prompts
* [pflag] (https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags
* [go-sh] (https://github.com/codeskyblue/go-sh) - like python-sh, for easy call shell with golang
* [mousetrap] (https://github.com/inconshreveable/mousetrap) - Detect starting from Windows explorer
* [cobra] (https://github.com/spf13/cobra) - A Commander for modern Go CLI interactions
* [inject] (https://github.com/codegangsta/inject) - Dependency injection for go
* [codeflower] (https://github.com/fzaninotto/CodeFlower) - Heavily inspired by this JS tool for step 1/4 of development.


## Contributing

Please feel free to open issues, help improve the tool or similar! 

## Future

In the future, this tool is planned to become a complete human-centered control-graph integrated visualization tool, different from standard control-graph analysis this tool will emphasize and focus on **easy digestibility/rapid mental modelling for smart contract developers.**

