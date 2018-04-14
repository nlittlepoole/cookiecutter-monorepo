# Research

This is not a service. It is just a repostitory to host proof of concept and research work via Jupyter Notebooks. 

## Using

```sh
$ make setup
$ make build
$ make run
```
You will be presented with a prompt that looks like the following:

```
Copy/paste this URL into your browser when you connect for the first time,
    to login with a token:
        http://localhost:8888/?token=<SOME_HASH>
```
Just copy the URL with the hash to use the

## Notes

- A .env_file has been created in the research/ directory. You can add any environment variables you want to pass to the Containers there. I recommend using a prefix schema. Like setting all environment variables on your host to `PYRESEARCH_<KEY>=<VALUE>`. Then you can run `env | grep PYRESEARCH > .env_file` to automatically push all the variables
- Add necessary requirements to requirements.txt to get them installed in the container when you `make build`
- You can override the notebook directory by passing a `WORKINGDIR=<PATH>` argument to the make command
- You can override the docker tag used (default is `research`) by passing a `TAG=<tag>` argument to the make command
- You can override the notebook port by passing `PORT=<PORT>` to the make command
