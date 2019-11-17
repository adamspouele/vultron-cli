# vultron-cli
Cli of vultron


## configuration

### connect to DigitalOcean cloud provider

connect vultron to DigitalOcean by setting `VULTRON_DO_TOKEN` environment variable which is the DigitalOcean personal access token.

## Tag schemas

All tags start by 'vultron:' prefix which is the first layer of a vultron tag. This allow vultron to recognize his own tags.

'vultron:'

The second layer of a tag is 'kind:' or 'res:'
'kind:' identifies in what the node is working in, below are the possible values :

- consul
- nomad
- client

examples : 

* vultron:kind:cluster

'res:' identifies the type of resource the node is   of, below are the possible values :

- server
- client


example : 

* vultron:res:server  
* vultron:res:client  

There is a third layer 'res:', this is the 'prop:' sub-layer.
the sub-layer allow mean 'property', has the word mean this layer allow to bind properties to nodes.
After the 'prop:' layer there is a sub layer 'key:value' which are the key and values.

example : 

* vultron:prop:k:v
* vultron:prop:name:test




