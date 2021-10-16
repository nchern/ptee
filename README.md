# ptee

Pronounced `/piːtiː/`. Process `tee`.

This utility starts a process(1st argument specifies the command to run), copies stdin to this process' stdin AND to stdout as well.

## Use cases


### Shell

Prints  out `file` to console and copies it to clipboard on Mac

```bash
$ cat file | ptee pbcopy
```

### Vim integrations

Can be used as a convenient util to write `vim` command filters. With its help you can send selected lines to an external command without modifying a text in the current buffer. 
