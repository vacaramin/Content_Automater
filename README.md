# Content Automater

Content Automater is a tool designed to automate the generation of SEO content based on keyword data stored in an XLSX file. This tool utilizes an XLSX file where each subsheet represents a product name, and within each sheet, keywords are listed along with their difficulty and monthly search volume. By leveraging the power of OpenAI's GPT (Generative Pre-trained Transformer) model, Content Automater transforms this data into rich and relevant content.

## Implementation Details

The implementation of Content Automater is done in Go (Golang). It includes an Excelize wrapper called GoSheet, which provides functions for interfacing with XLSX files. These functions are designed for easy access within the main package, allowing seamless integration of keyword data extraction and content generation.

## Features

- <b>Automated Content Generation:</b> Content Automater automates the process of generating SEO content based on keyword data.
- <b>XLSX Data Integration:</b> Keywords, along with their difficulty and monthly search volume, are extracted from an XLSX file.
- <b>OpenAI GPT Integration:</b> OpenAI's GPT model is used to transform keyword data into rich and relevant content.
- <b>Excelize Wrapper:</b> **GoSheet** The tool includes a wrapper for Excelize, simplifying the process of working with XLSX files in Go.