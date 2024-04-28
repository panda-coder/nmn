# Node Mapping Notation

# Examples

**example1.nmn**
```yaml
version: 0.1
add:
- "example2.nmn"
nodes:
- id: 'node1'
  name: 'NodeExample'
  edges:
    - type: 'standart'
      node: 'node2'
```

**example2.nmn**
```yaml
version: 0.1
add:
- "sub/example3.nmn"
nodes:
- id: 'node2'
  name: 'NodeExample'
- id: 'node3'
  name: 'NodeExample'
```


**sub/example3.nmn**
```yaml
version: 0.1
add:
- "example4.nmn"
nodes:
- id: 'node4'
  name: 'NodeExample'
```

**sub/example4.nmn**
```yaml
version: 0.1
nodes:
- id: 'node5'
  name: 'NodeExample'
- id: 'node6'
  name: 'NodeExample'
```

## Contributors

Thank you for contributing to the NodeMappingNotation!

[![Contributors](https://contrib.rocks/image?repo=panda-coder/nmn)](https://github.com/panda-coder/nmn/graphs/contributors)

## License

[MIT](./LICENSE)