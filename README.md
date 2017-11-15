# markdown-toc

<p align="center">
    <strong>Generating your Markdown Table of Contents</strong>
</p>

<p align="center">
    <a href="https://travis-ci.org/sebdah/markdown-toc"><img src="https://img.shields.io/travis/sebdah/markdown-toc.svg" /></a>
    <a href="https://github.com/sebdah/markdown-toc/issues"><img src="https://img.shields.io/github/issues/sebdah/markdown-toc.svg" /></a>
    <a href="https://github.com/sebdah/markdown-toc/blob/master/LICENSE"><img src="https://img.shields.io/github/license/sebdah/markdown-toc.svg" /></a>
</p>

`markdown-toc` is a small application written in Go that helps you generate a
Table of Contents (ToC) for your Markdown file. There are already a number of
scripts etc doing this, but I failed to find one that suited my needs.

In short the features of `markdown-toc` are:

- Cross platform (OS X, Linux, Windows)
- Replacement of an existing ToC
  - The new file can be written to `stdout` or overwritten on disk
- Configurable header

**Please star the project if you like it!**

<!-- ToC start -->
# Table of Contents

- [markdown-toc](#markdown-toc)
- [Table of Contents](#table-of-contents)
- [Example usage](#example-usage)
  - [Generating a ToC to `stdout`](#generating-a-toc-to-stdout)
  - [Set a custom header](#set-a-custom-header)
  - [Print the full Markdown file, not only the ToC](#print-the-full-markdown-file-not-only-the-toc)
  - [Inject the ToC into a file on disk](#inject-the-toc-into-a-file-on-disk)
- [Helping out!](#helping-out)
- [License](#license)
<!-- ToC end -->




# Example usage

## Generating a ToC to `stdout`

Command:

    markdown-toc README.md

Output:

    <!-- ToC start -->
    # Table of Contents

    - [`markdown-toc` - Generate your Table of Contents](#`markdown-toc`---generate-your-table-of-contents)
    - [Example usage](#example-usage)
      - [Generating a ToC to `stdout`](#generating-a-toc-to-`stdout`)
    - [License](#license)
    <!-- ToC end -->

## Set a custom header

Command:

    markdown-toc --header "# ToC" README.md

Output:

    <!-- ToC start -->
    # ToC

    - [`markdown-toc` - Generate your Table of Contents](#`markdown-toc`---generate-your-table-of-contents)
    - [Example usage](#example-usage)
      - [Generating a ToC to `stdout`](#generating-a-toc-to-`stdout`)
    - [License](#license)
    <!-- ToC end -->

## Print the full Markdown file, not only the ToC

    markdown-toc --replace README.md

This will print the full Markdown of `README.md` and a table of contents section
will be injected into the Markdown based on the following rules:

- If no ToC was found, the ToC will be injected on top of the file
- If a section starting with `<!-- ToC start -->` and ending with
  `<!-- ToC end -->` is found, it will be replaced with the new ToC.

## Inject the ToC into a file on disk

    markdown-toc --replace --inline README.md

This will overwrite the `README.md` file on disk with the full Markdown of
`README.md` and a table of contents section will be injected into the Markdown
based on the following rules:

- If no ToC was found, the ToC will be injected on top of the file
- If a section starting with `<!-- ToC start -->` and ending with
  `<!-- ToC end -->` is found, it will be replaced with the new ToC.

# Helping out!

There are many ways to help out with this project. Here are a few:

- Answer questions [here](https://github.com/sebdah/markdown-toc/issues)
- Enhance the documentation
- Spread the good word on [Twitter](https://twitter.com) or similar places
- Implement awesome features. Some of the suggested features can be found
  [here](https://github.com/sebdah/markdown-toc/issues)

# License

MIT license