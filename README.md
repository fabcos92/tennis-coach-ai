# 🎾 Tennis Coach AI — Tennis Performance Analyzer (LLM-powered API)

A backend system demonstrating practical LLM integration with structured output and reliability patterns.

---

## 🚀 Project Goal

The goal of this project is to demonstrate:
- backend system design
- integration with LLM APIs
- handling of unreliable external dependencies
- clean and modular architecture

---

## ✨ Features

- Tennis performance analysis from:
  - match descriptions (text input)
  - structured stats (JSON input)
- Structured LLM output:
  - issues
  - recommendations
  - focus area
- Pluggable LLM providers:
  - mock (local development)
  - Groq (active)
  - OpenAI-compatible (ready)
- Reliability layer:
  - retries for external calls
  - request timeouts
  - JSON validation & parsing
  - fallback responses

---

## 🧠 Architecture

The system follows a clean layered design:

- HTTP layer (handlers)
- service layer (business logic)
- LLM abstraction layer (provider interface)

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
    "Poor endurance during long rallies",
    "Loss of focus under pressure"
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
- Llama models via Groq

---

## 🔌 LLM Design

The system uses a provider-agnostic interface:

- Easy to swap LLM providers via configuration
- Supports mock mode for local development
- Designed to avoid coupling business logic with external APIs

---

## ⚙️ Reliability Design

- retry logic for external API calls
- request timeouts
- structured output validation
- fallback responses for malformed LLM output

---

## 📈 Status
MVP — functional, tested, and ready for extension.

---

## 📄 License
MIT