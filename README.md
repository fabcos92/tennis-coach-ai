# 🎾 Tennis Coach AI — Tennis Performance Analyzer (LLM-powered API)

A backend system demonstrating practical LLM integration with clean architecture and production-oriented design.

---

## 🚀 Project Goal

This project showcases:
- designing a backend service around an unreliable external dependency (LLM)
- clean separation of concerns (domain / application / infrastructure)
- structured output generation from unstructured input
- pragmatic system design (no overengineering)

---

## ✨ Features

- Tennis performance analysis from:
  - structured match stats (JSON)
  - free-form match descriptions (text)
- Structured LLM output:
  - issues (with severity)
  - actionable recommendations
  - focus area
- Pluggable LLM providers:
  - mock (local development)
  - Groq (active)
  - OpenAI-compatible
- Reliability layer:
  - retries & timeouts for external calls
  - JSON parsing & validation
  - fallback responses for invalid LLM output

---

## 🧠 Architecture

The system follows a clean layered design:

- HTTP layer (request handling)
- Application layer (use cases / command handlers)
- Domain layer (business rules and validation)
- Infrastructure layer (LLM, external integrations)

LLM is treated as an unreliable external dependency and fully decoupled from business logic.

---

## 🧪 Example

### Request

```json
{
  "type": "text",
  "text": "I struggle with long rallies and lose focus under pressure"
}
```

### Response

```json
{
  "issues": [
    {
      "text": "Poor endurance during long rallies",
      "severity": "high" 
    },
    {
      "text": "Loss of focus under pressure",
      "severity": "medium"
    }
  ],
  "recommendations": [
    "Endurance-focused rally drills",
    "Pre-point mental routine"
  ],
  "focus_area": "Endurance and Mental Toughness"
}
```

---

## 🛠 Tech Stack

- Go
- REST API (net/http)
- LLM APIs (Groq / OpenAI-compatible)

---

## 🔌 LLM Design

The system uses a provider-agnostic interface:

- Provider-agnostic interface (Groq / OpenAI / Mock)
- Runtime provider switching via configuration
- Prompt builder separated from business logic
- Response mapping isolated in infrastructure layer

---

## ⚙️ Reliability Design

- retry logic for external API calls
- request timeouts
- structured output validation
- fallback responses for malformed LLM output

---

## 📈 Status
MVP — production-oriented foundation, ready for extension
(e.g. async processing, prompt versioning, observability)

---

## 📄 License
MIT