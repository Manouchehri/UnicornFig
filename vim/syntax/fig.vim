syntax keyword figKeywords
    \ if
    \ function
    \ define
    \ true
    \ false

syntax keyword figBuiltins
    \ concat
    \ substr
    \ index
    \ length
    \ upcase
    \ downcase
    \ split
    \ at
    \ not
    \ zero?
    \ and
    \ or
    \ list
    \ first
    \ tail
    \ append
    \ size
    \ mapping
    \ assoc
    \ get
    \ keys
    \ print
    \ env
    \ ignored

syntax match figNumber "\v<\d+>"
syntax match figNumber "\v<\d+\.\d+>"

syntax match figOperator "\v\*"
syntax match figOperator "\v/"
syntax match figOperator "\v\+"
syntax match figOperator "\v-"
syntax match figOperator "\v\="
syntax match figOperator "\v\<"
syntax match figOperator "\v\>"
syntax match figOperator "\v\>"
syntax match figOperator "\v\%"
syntax match figOperator "\v\>\="
syntax match figOperator "\v\<\="

syntax region figString start=/"/ end=/"/ oneline
syntax region figString start=/'/ end=/'/ oneline

syntax region figComment start=/;/ end=/$/ oneline

highlight default link figComment Comment
highlight default link figNumber Number
highlight default link figString String
highlight default link figKeywords Keyword
highlight figBuiltins ctermfg=Blue
highlight figOperator ctermfg=LightBlue
