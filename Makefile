.PHONY: generate update_spec

generate:
	swagger generate client --spec=swagger.json --name=logface --target=.

update_spec:
	wget -q -O swagger.json https://api.logface.io/swagger.json

