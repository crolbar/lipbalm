- [x] fix repeated work in JoinVertical
  in getLines when we construct thhe blocks we get the StringWidth
  of each line, and then again when we are writing the lines + padding
  we again get the string width of each line
  (10-20% improvement)
- [ ] StringWidth
- [ ] JoinHorizontal
- [ ] basic ansi color injection
