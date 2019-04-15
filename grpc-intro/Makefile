
clean:
	@rm -f *.html
	@rm -f *.svg

%.blockdiag.svg: %.blockdiag
	@blockdiag -a -T svg -o "$@" "$(basename $@)"

%.seqdiag.svg: %.seqdiag
	@seqdiag -a -T svg -o "$@" "$(basename $@)"

blockdiag:
	@find . -type f -name '*.blockdiag' | sed 's/$$/.svg/' | xargs make -s

seqdiag:
	@find . -type f -name '*.seqdiag' | sed 's/$$/.svg/' | xargs make -s

images: blockdiag seqdiag

html: images
	@asciidoctor index.adoc

pdf:
	@asciidoctor-pdf index.adoc

watch: html
	@reflex -d none -c reflex.conf
