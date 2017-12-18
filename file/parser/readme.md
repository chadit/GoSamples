Rework of first attempt, to see older version, see parser.old


To run application : 
    - docker-compose up
    
* reads one or more keyword files, one keyword per line
* reads one or more text files, parses, creates a results file with the following format

tab-separated text file with the following output:
```
num dupes   d
med length  lm
std length  ls
med tokens  tm
std length  ts
keyword_a   ka
...
keyword_n   kn
```
where `d` represents the number of exactly duplicated lines seen, `lm` is the median length of lines
(number of unicode characters), `ls` is the standard deviation of line length, `tm` is the median
number of tokens (via whitespace tokenization) per line, `ts` is the standard deviation of tokens
per line, and `ka` ... `kn` represent the total number of lines the keyword was seen on, sorted
alphabetically.


CPU Profiler on 
go tool pprof -text ./cpu.pprof