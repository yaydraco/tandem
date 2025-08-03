# Tandem Media Agent

This directory contains media generation tools and templates for creating promotional content using the Tandem Media Agent.

## Overview

The Media Agent uses two powerful CLI tools to generate visual content:

- **VHS**: Creates terminal recordings and GIFs to showcase Tandem's interface and capabilities
- **Freeze**: Generates beautiful code screenshots with syntax highlighting and styling

## VHS Tape Files

VHS tape files (`.tape`) contain scripts that define terminal recording sessions. They specify:
- Commands to run
- Timing and pauses
- Output settings (resolution, format, theme)
- Styling options

## Freeze Configuration

Freeze configurations define the styling for code screenshots including:
- Color themes
- Font settings
- Padding and shadows
- Border styling

## Usage

The Media Agent can be invoked through:

1. **GitHub Actions**: Automatically triggered on pushes to generate content
2. **Manual CLI**: Using the Tandem CLI with specific prompts for the media agent
3. **Interactive Mode**: Through the TUI interface

## Examples

See the example files in this directory for ready-to-use templates that showcase Tandem's capabilities.