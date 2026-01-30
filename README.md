# Report Orchestrator
## Overview
Report Orchestrator is a backend system that handles report generation that consumes time as a background service without blocking client requests.

This project focuses on backend architecture, job orchestration, and state managerment rather than the actual business logic.

## Problem Statement
In many real-world systems such as analystics platforms, generating reports takes significant time due to data aggregation and computation. Executing such workloads synchronously leads to poor user experience, request timeouts, and limited scalability.

Report Orchestrator addresses this by:
- Decoupling request handling from execution
- Processing reports asynchronously
- Providing reliable status tracking and failure visibility

## High-Level Architecture
The system is composed of the following logical components:
1) **API Layer**:
Accepts report requests and exposes report status to clients
2) **Scheduler**:
Periodically scans for pending reports and assigns them for execution
3) **Worker Pool**:
Executes report generation tasks asynchronously
4) **Storage Layer**:
Persists report jobs and their lifecycle states

## Report Lifecycle
Each report request is modeled as a job with the following lifecycle:

    REQUESTED → RUNNING → COMPLETED       
                        → FAILED

- **REQUESTED**: Job accepted and persisted
- **RUNNING**: Job is being processed by a worker
- **COMPLETED**: Report generated successfully
- **FAILED**: Processing terminated due to an error

## What this Project Demonstrates?
- Asynchronous job orchestration
- Background worker execution
- Deterministic state transitions
- Failure handling and isolation
- Backend-focused system design (UI complexity is intentionally kept minimal to emphasize backend correctness)

## Documentation
Detailed design and domain documentation is available in the /docs directory.

#### Author’s Note
This project was built to explore realistic backend engineering concerns such as orchestration, durability, and observability, rather than focusing on UI or domain-specific business logic.