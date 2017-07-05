# jump
A small CLI to move about quickly in the filesystem.

## Demo
![Demo](https://raw.githubusercontent.com/reficul31/jump/master/assets/demo.gif)

## About
Inspired from [Teleport](https://github.com/bollu/teleport).  
If you are like me, you must hate the golang project structure. I had to write a lot of ```cd``` commands just to change project directories.  
Gone are those days! This is a small tool that can help navigate the filesystem faster.

## Installation
Before you go about installing the tool make sure that
- You have golang installed
- You have set the PATH variable

Jump is build with golang and can be easily installed with the following command
> go get -u github.com/reficul31/jump

To use the jump you will also require to add the following line to your shell script.
> source $GOPATH/src/github.com/reficul31/jump/jump.sh

## Usage
The CLI can be used in the place of cd as well. This means that commands such as given below will work just the same as a normal cd command.
> jp /home/user/path/to/dir

##### Add a checkpoint
To add a checkpoint, you can use the following command.
> jp add <name>

##### Remove a checkpoint
To remove a checkpoint, you can use the following command.
> jp rm <name>

To remove all the checkpoints set the ```--all`` flag to true.
> jp --all rm

##### Show Checkpoints
To see all the active checkpoints, you can use the following command.
> jp show

## Contributing
Pull Requests and Feature Requests are welcome.

## Status
Some work still left to be done. Any help would be deeply appreciated.
- [ ] Add tests
- [ ] Travis configuration
- [ ] Add color to console writes
