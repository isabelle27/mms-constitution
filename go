mkdir build
latexmk -jobname=build/constitution constitution.tex
pdflatex -jobname=build/constitution constitution.tex
