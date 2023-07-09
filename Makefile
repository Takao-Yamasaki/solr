up: # Start Solr
	docker-compose up -d
build: # Rebuild Solr
	docker-compose up --build
exec: # login Solr Container
	docker-compose exec solr bash
down: # Stop Solr
	docker-compose down
ps: # Status Check Solr Container
	docker-compose ps
rm: # Remove Solr volume
	docker volume rm solr_data
core: # Create solrbook Core
	bin/solr create_core -c solrbook -d basic_configs
reload: # Reload Core
	curl "http://localhost:8983/solr/admin/cores?action=RELOAD&core=solrbook"
index: # Add Index
	curl "http://localhost:8983/solr/solrbook/update?commit=true" --data-binary @sample-books.json -H 'Content-Type: application/json; charset=utf8' -T "sample-books.json" -X POST
