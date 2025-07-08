# template-go-api

## References:
- https://github.com/golang-standards/project-layout
- https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Option 1: Layer-First Structure (chosen)
- /domain/<br>
  - /user/<br>
- /repository/<br>
  - /user/<br>
- /api/<br>
  - /user/<br>


## Option 2: Feature-First Structure
- /internal/<br>
  - /user/<br>
  - /order/<br>
  - /auth/<br>

## tl;dr
- Choose layer-first if you want strict architecture enforcement, plan for growth, or have multiple teams.

- Choose feature-first if you want fast development, simplicity, or are a small team.