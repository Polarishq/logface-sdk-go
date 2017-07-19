.PHONY: generate update_spec

generate:
	swagger generate client --spec=swagger.json --name=logface --target=src/generated/

update_spec:
	wget -q -O swagger.json https://api.logface.io/swagger.json

