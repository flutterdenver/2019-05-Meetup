
%.blockdiag.svg: %.blockdiag
	@blockdiag -a -T svg -o "$@" "$(basename $@)"

%.seqdiag.svg: %.seqdiag
	@seqdiag -a -T svg -o "$@" "$(basename $@)"

.PHONY: html
html:
	@asciidoctor index.adoc

.PHONY: pdf
pdf:
	@asciidoctor-pdf index.adoc -o "Flutter-Denver-gRPC-Intro.pdf"

.PHONY: blockdiag
blockdiag:
	@find . -type f -name '*.blockdiag' | sed 's/$$/.svg/' | xargs make -s

.PHONY: seqdiag
seqdiag:
	@find . -type f -name '*.seqdiag' | sed 's/$$/.svg/' | xargs make -s

.PHONY: svg
svg: blockdiag seqdiag

.PHONY: serve
serve: html seqdiag
	@reflex -d none -c reflex.conf
