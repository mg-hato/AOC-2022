# Terminal Output Reader

Input consists of a sequence of lines describing `cd` and `ls` commands. Each `ls` command will be followed by the list of filesystem items (directories / files). For input to be correct, following should hold:

1. First command is `cd /`
2. Any subsequent `cd` command can be of the following three forms:
   - `cd /`
   - `cd ..`
   - `cd <directory_name>`
3. Directory names are strings consisting of only lower case characters
4. File names are strings consisting of lower case characters and optionally they can contain dots. If they do contain dot(s) the following should hold:

   - The file name starts and ends with a lower case character
   - Only lower case characters can be next to a dot i.e. no two adjacent dots
5. In a `ls`-command listing of filesystem items, each item has a unique name within that listing

If the terminal output reader returns without an error, one can assume that the above guarantees hold.