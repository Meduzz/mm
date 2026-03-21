# mm
model manager

This is intended to be a cli companion to llama.cpp.

## TODO

### Commands
**serve**
This command will start the llama-server with a few defaults.

**ls**
This command will list local models available to llama.cpp.

**rm**
This command will take the name of a model as parameter and will remove the model from the disk.

**unload**
This command will optionally take the name of a model. If a model is specified, then it will unload that model via a POST call to the llama-server. If no model is specified, then all models currently loaded will be unloaded.