.PHONY: models generate

# ==============================================================================
# Go migrate postgresql

# ==============================================================================
# Swagger Models
models:
	$(call print-target)
	find ./api/models -type f -not -name '*_test.go' -delete
	swagger generate model -m api -f api/swagger/pet.yml -t api/models

generate: models
