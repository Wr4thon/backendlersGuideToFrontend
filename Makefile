presentation: pandoc pdfpc

pandoc:
	pandoc -t beamer --pdf-engine=xelatex --self-contained -o ./presentation.pdf ./slides/presentation.md
	
pdfpc:
	pdfpc ./presentation.pdf