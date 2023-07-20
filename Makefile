# TODO: ヘルプを実装すること
# TODO: GoコンテナからのAPI実装
.PHONY: setup up build solr go down ps rm core rmcore reload add update delete nocommit rollback

SOLR_DIR = /opt/solr

setup: core reload add
up: # Start Solr
	@docker-compose up -d
build: # Rebuild Solr
	@docker-compose up --build
solr: # login Solr Container
	@docker-compose exec solr bash
go: # login Solr Container
	@docker-compose exec go bash
down: # Stop Solr
	@docker-compose down
ps: # Status Check Solr Container
	@docker-compose ps
rm: # Remove Solr volume
	@docker volume rm solr_data
core: # Create solrbook Core & SetUp Schema
	@docker-compose exec solr bash -c "\
		bin/solr create_core -c solrbook -d basic_configs && \
		cp $(SOLR_DIR)/server/solr/solrbook_setup/schema.xml $(SOLR_DIR)/server/solr/solrbook_setup/solrconfig.xml $(SOLR_DIR)/server/solr/solrbook/conf/"
rmcore: # Remove solrbook Core
	@docker-compose exec solr bash -c "bin/solr delete -c solrbook"
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
