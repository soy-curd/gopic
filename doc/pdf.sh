rm book.pdf
rm -r book-pdf

md2review src/chap00.md > chap00.re
md2review src/chap01.md > chap01.re
md2review src/chap02.md > chap02.re
md2review src/chap03.md > chap03.re
md2review src/chap04.md > chap04.re
md2review src/chap05.md > chap05.re

review-pdfmaker config.yml
open book.pdf
