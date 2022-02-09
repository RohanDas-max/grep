# Grep Cli
   ### to run use `make build` then you can use `./mygrep [string] [filename]`. This command will check if the string is available in the file if yes then the line will be Printed otherwise it will abort

 ```  eg: ./mygrep "hello" test.txt```

  ###  To search in a directory `./mygrep [string_to_search] [directory]`
 ``` eg: ./mygrep "hello" dir ```

  ### To search in standard input values `./mygrep [string_to_search]` then hit enter and the prompt will wait for your input if the input contains the parameter it will be printed. 
  ``` eg: ./mygrep "Hi" ```