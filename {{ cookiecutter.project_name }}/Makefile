.PHONY: air
# run dev server
air:
	air

.PHONY: wire
# generate ic
wire:
	cd cmd && wire

.PHONY: ent
# generate validate proto
ent:
	cd data && ent generate --feature sql/execquery --feature sql/lock ./ent/schema
