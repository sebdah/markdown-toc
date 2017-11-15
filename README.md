# markdown-toc

<p align="center">
    <strong>Generating your Markdown Table of Contents</strong>
</p>

<p align="center">
    <a href="https://travis-ci.org/sebdah/markdown-toc"><img src="https://img.shields.io/travis/sebdah/markdown-toc.svg" /></a>
    <a href="https://hub.docker.com/r/sebdah/markdown-toc/"><img src="https://img.shields.io/badge/docker-latest-blue.svg" /></a>
    <a href="https://godoc.org/github.com/sebdah/markdown-toc"><img src="https://img.shields.io/badge/godoc-documentation-blue.svg" /></a>
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
- Able to skip `n` number of initial headers
- Configurable depth of headers to include
- Customizable header for the table of contents

**Please star the project if you like it!**

<!-- ToC start -->
# Table of Contents

- [Table of Contents](#table-of-contents)
- [Installation](#installation)
  - [Running it in Docker](#running-it-in-docker)
  - [Building it yourself](#building-it-yourself)
- [Example usage](#example-usage)
  - [Generating a ToC to `stdout`](#generating-a-toc-to-stdout)
  - [Controlling depth of headers](#controlling-depth-of-headers)
  - [Set a custom header](#set-a-custom-header)
  - [Render the ToC without a header](#render-the-toc-without-a-header)
  - [Skip `n` headers](#skip-n-headers)
  - [Print the full Markdown file, not only the ToC](#print-the-full-markdown-file-not-only-the-toc)
  - [Inject the ToC into a file on disk](#inject-the-toc-into-a-file-on-disk)
- [Helping out!](#helping-out)
  - [Running the test suite](#running-the-test-suite)
  - [Build locally with build flags](#build-locally-with-build-flags)
- [Releasing new versions](#releasing-new-versions)
- [License](#license)
<!-- ToC end -->

# Installation

## Running it in Docker

The project has a Docker image that you can use easily:

    docker run -v $(pwd)":/app" -w /app --rm -it sebdah/markdown-toc README.md

The above will mount your current directory into the Docker container. Just
modify the `-v` flag according to your needs if you need to modify some other
folder.

`markdown-toc` is the `ENTRYPOINT` in the `Dockerfile`, which means that it's
the default command.

## Building it yourself

Currently the easiest way is to clone the repository and run:

    make install

You will end up having a binary called `markdown-toc` in your system afterwards.

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

## Controlling depth of headers

Using the `--depth` flag you can control how many labels of headers to include
in the output. If the `--depth` is set to `1`, only level 1 headers are
included. Set this value to `0` (default) to include any depth of headers.

Command:

    markdown-toc --depth 1 README.md

Output:

    <!-- ToC start -->
    # Table of Contents

    - [`markdown-toc` - Generate your Table of Contents](#`markdown-toc`---generate-your-table-of-contents)
    - [Example usage](#example-usage)
    - [License](#license)
    <!-- ToC end -->

## Set a custom header

By default we print a header like `# Table of Contents` above the table of
contents. You can change the header to suit your project using the `--header`
flag.

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

## Render the ToC without a header

By default we print a header like `# Table of Contents` above the table of
contents. You can remove the header by providing `--no-header` to the command.

Command:

    markdown-toc --no-header README.md

Output:

    <!-- ToC start -->
    - [`markdown-toc` - Generate your Table of Contents](#`markdown-toc`---generate-your-table-of-contents)
    - [Example usage](#example-usage)
      - [Generating a ToC to `stdout`](#generating-a-toc-to-`stdout`)
    - [License](#license)
    <!-- ToC end -->

## Skip `n` headers

If you do not want to include `n` number of initial headers in your ToC, you can
use the `--skip-headers=1` flag. This is useful if you have your project name as
the first header and you don't really want that in the ToC for example.

Command:

    markdown-toc --skip-headers=1 README.md

Output:

    <!-- ToC start -->
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

## Running the test suite

    make test

## Build locally with build flags

    make build

# Releasing new versions

From the `master` branch run `make release`. You will need to have access to
pushing to the GitHub project as well as Docker Hub.

# License

MIT license