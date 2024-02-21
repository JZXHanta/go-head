# go-head

#### A clone of the ```head``` command used in Unix/Linux command lines
---

- With no extra arguments other than file-name (or a space separated list of file-names) it will return the first 10 lines of the filename(s) specified
  - example: ```go-head example.txt``` 

- If ran with the argument '-' and no file-name it will read the first 10 lines of the stardard input
  - example: ```cat test.txt | go-head -```

- If ran with argument '-n' followed by a number it will return the specified number of lines from the beginning of the input
  - example: ```go-head -n 30 test.txt```

- Arguments can be used in conjuntion
  - example: ```cat test.txt | go-head -n 30 -```
