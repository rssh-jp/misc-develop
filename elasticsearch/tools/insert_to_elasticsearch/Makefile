setup:
	wget https://dumps.wikimedia.org/jawiki/latest/jawiki-latest-pages-articles.xml.bz2
	git@github.com:zaemyung/wikiextractor.git
	python wikiextractor/WikiExtractor.py -o output -b 80M jawiki-latest-pages-articles.xml.bz2 --json
	ls ./output/AA/* -d | xargs -L 1 -P 10 bash -c './reformat_to_ndjson.py $0'

insert:
	ls ./output/AA/*_new.ndjson -d | xargs -L 1 -P 3 bash -c 'echo $0 ; cat $0 | curl -s -X POST -H '\''Content-Type: application/x-ndjson'\'' '\''http://127.0.0.1:9200/wiki/_bulk?pretty'\'' --data-binary @-;'

