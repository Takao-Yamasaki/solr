.PHONY: go solr
up: # Start Solr
	@docker-compose up -d
build: # Rebuild Solr
	@docker-compose up --build
solr: # login Solr Container
	docker-compose exec solr bash
go: # login Solr Container
	@docker-compose exec go bash
down: # Stop Solr
	@docker-compose down
ps: # Status Check Solr Container
	@docker-compose ps
rm: # Remove Solr volume
	@docker volume rm solr_data
core: # Create solrbook Core
	@bin/solr create_core -c solrbook -d basic_configs
reload: # Reload Core
	@curl "http://localhost:8983/solr/admin/cores?action=RELOAD&core=solrbook"
add: # Add Index
	@curl "http://localhost:8983/solr/solrbook/update?commit=true" --data-binary @sample-books.json -H 'Content-Type: application/json; charset=utf8' -T "sample-books.json" -X POST
update: # Update Index
	@curl "http://localhost:8983/solr/solrbook/update?commit=true" --data-binary @update.json -H 'Content-Type: application/json; charset=utf8' -T "update.json" -X POST
delete: # Delete Index
	@curl "http://localhost:8983/solr/solrbook/update?commit=true" --data-binary @delete.json -H 'Content-Type: application/json; charset=utf8' -T "delete.json" -X POST
nocommit: # Register Document without Commit
	@curl "http://localhost:8983/solr/solrbook/update" --data-binary @update.json -H 'Content-Type: application/json; charset=utf8' -T "update.json" -X POST
rollback: # Rollback
	@curl "http://localhost:8983/solr/solrbook/update?rollback=true"
