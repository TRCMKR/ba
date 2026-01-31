# ba

BA stands for *Bibliotheca Alexandrina* (Latin, Library of Alexandria) – the most famous library of the ancient world.

> We can roam the bloated stacks of the Library of Alexandria, where all imagination and knowledge are assembled;   
> we can recognize in its destruction the warning that all we gather will be lost,   
> but also that much of it can be collected again...
> - Alberto Manguel

BA is a collection of packages need for fast development of Golang microservices

---

## ⚠️ Stability Notice

**No compatibility guarantees are provided**.  
Breaking changes may be introduced in minor or patch releases.

---

## Usage 

To add BA submodule to your project, run the following command:

```bash
go get -u github.com/trcmkr/ba/<module_name>
```

---

## Features

- [x] closer – releases resources on context done or any of signals
- [x] utils – lo-inspired collection of functions
- [ ] bapub – kafka publisher
- [ ] basup – kafka subscriber
- [ ] flusher – accumulates elements into a batch and on flush passes the batch to a handler 

--- 

## Release

Release new version only using this command:

```bash
./release <commit_name> <module_name> major/minor/patch
```