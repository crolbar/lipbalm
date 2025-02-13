-   [x] fix repeated work in JoinVertical
        in getLines when we construct thhe blocks we get the StringWidth
        of each line, and then again when we are writing the lines + padding
        we again get the string width of each line
        (10-20% improvement)
-   [x] better StringWidth (35-40% improvement, - support for icons and other multicell chars)
-   [x] JoinHorizontal
-   [x] color: combine bg + fg into one ansi code
-   [ ] expand{Vertically,Horizontally} funcs that add whitespaces/new lines to the string + alignment
-   [x] basic ansi color injection
