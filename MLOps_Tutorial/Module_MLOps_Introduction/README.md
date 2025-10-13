# Module – MLOps Introduction

## Objectives

- Understand what MLOps is and why it’s essential
- Learn the ML model lifecycle: training → versioning → deployment → monitoring
- Create a simple ML training script and track it with MLflow
- Expose the model via an API
- Understand how orchestration and observability fit in production systems

---

## Prerequisites

- Basic Python knowledge (for ML model training)
- Installed `mlflow`, `fastapi`, and `scikit-learn`
- Optional: Docker knowledge for containerization

---

## Key Concepts

1. **MLOps (Machine Learning Operations)**
   A set of practices that unify ML system development (Dev) and ML system operations (Ops).

2. **Core Stages:**
   - **Experimentation**: Train models and log metrics
   - **Versioning**: Track models and data using MLflow or DVC
   - **Deployment**: Serve models through APIs
   - **Monitoring**: Track drift and performance in production

---

## Tasks

1. Train a simple ML model (Linear Regression)
2. Log the parameters and metrics using MLflow
3. Serve the trained model with FastAPI
4. Simulate basic orchestration and monitoring logic in Go and Rust
