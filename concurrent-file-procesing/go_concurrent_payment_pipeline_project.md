# Go Learning Project: Concurrent Payment Transaction Processing Pipeline

## Purpose
This project is designed to deepen understanding of advanced Go concepts including:

- Parallel programming (goroutines, channels, worker pools)
- Context cancellation and graceful shutdown
- Structured error handling and error classification
- Generics used in realistic architectures
- Pipeline design and backpressure
- Deterministic testing of concurrent programs

The project will simulate a real-world backend ingestion system that processes payment transaction files.

---

# High-Level Architecture

The system will ultimately consist of two CLI programs:

## 1. Data Generator
Command:

`cmd/gendata`

Responsibilities:

- Generate synthetic payment transaction datasets
- Produce CSV files
- Create mostly valid data with configurable bad records
- Allow deterministic output via random seed

Example usage:

```
gendata --out ./data --files 5 --rows 1000 --bad-percent 7 --seed 42
```

Output: multiple CSV files containing payment transactions.

---

## 2. Ingestion Pipeline
Command:

`cmd/ingest`

Responsibilities:

- Discover input files
- Parse records
- Validate records
- Enrich records
- Route valid records to success output
- Route invalid records to rejection output
- Run processing concurrently using worker pipelines

Pipeline concept:

File discovery → parsing → validation → enrichment → output writers

Each stage communicates via channels.

---

# Domain Model

The core domain entity is a **Payment Transaction**.

Planned fields:

- transaction_id
- customer_id
- merchant_id
- amount
- currency
- timestamp
- payment_method
- status
- country

These represent realistic attributes commonly found in payment processing systems.

---

# Key Design Decisions

## Transaction IDs
Chosen approach: **UUID-based IDs**

Reasoning:

- globally unique
- realistic for distributed systems
- avoids collision problems

---

## Input Format
Chosen format: **CSV**

Reasons:

- common in data ingestion pipelines
- easy to inspect manually
- easy to corrupt intentionally for testing
- realistic for batch integrations

---

## Project Structure

```
cmd/
  gendata/
  ingest/

internal/
  domain/
  generator/
  pipeline/
  validation/
  enrichment/
  output/
```

Guiding principles:

- `cmd` contains only CLI entrypoints
- `internal` contains application logic
- domain types remain independent of IO

---

# Generator Design

The generator will:

1. Produce domain `Transaction` objects
2. Convert those objects into CSV rows

This separation allows reuse of the domain model across generator and ingestion pipeline.

---

## Randomness Strategy

The generator will **receive a random source as a dependency** rather than relying on global randomness.

Conceptual structure:

```
Generator
  └── rng
```

Advantages:

- deterministic datasets using seeds
- reproducible bugs
- easier testing
- avoids hidden global state

Example deterministic run:

```
gendata --rows 1000 --seed 42
```

---

# Development Milestones

## Milestone 1
Minimal generator

- CLI flags
- generate valid transactions
- write one CSV file

## Milestone 2
Multiple files

- configurable file counts

## Milestone 3
Bad data injection

Examples:

- malformed timestamps
- negative amounts
- missing fields
- duplicate transaction IDs

## Milestone 4
Start ingestion pipeline

Sequential processing of generated files.

## Milestone 5
Concurrent pipeline

- worker pools
- fan-out / fan-in

## Milestone 6
Context cancellation and shutdown rules

## Milestone 7
Error classification

Three categories:

- record errors
- file errors
- system errors

## Milestone 8
Reusable generic pipeline components

Possible candidates:

- worker pools
- pipeline stages
- result wrappers

## Milestone 9
Concurrency testing

Using modern Go tooling including synctest where appropriate.

---

# Learning Goals

By completing this project the developer should understand:

- when to use goroutines vs worker pools
- channel ownership rules
- how backpressure naturally arises in pipelines
- structured error handling in Go
- when generics improve architecture
- deterministic testing of concurrent systems

---

# CLI Framework Decision

The project uses **Cobra** for building command-line interfaces.

Reasons:

- Industry-standard CLI framework in Go
- Used by major tools such as Kubernetes, Helm, and Hugo
- Built-in support for subcommands, flags, and help generation

Design rule:

- Cobra commands only parse CLI flags and call application logic.
- Business logic must live in internal packages.

Conceptual flow:

CLI (Cobra command) → generator package → domain types → CSV output

Project CLI layout:

```
cmd/
  gendata/
    main.go
    root.go
    generate.go
```

---

# Generator Configuration Decision

The generator will use a **configuration struct** instead of long parameter lists.

Reasons:

- cleaner APIs
- easier extension as more CLI flags are added
- easier validation and defaults

Example conceptual structure:

```
GeneratorConfig
  ├── Rows
  ├── OutputFile
  └── Seed
```

---

# Initial Output Strategy

For the first milestone the generator will:

- create **one CSV file**
- write **N rows of transactions** into that file

This keeps the first implementation simple while establishing the data model and generation logic.

Future milestones will extend this to multiple files.

---

# Current Status

We are currently implementing:

**Milestone 1 — Transaction data generator CLI**.

