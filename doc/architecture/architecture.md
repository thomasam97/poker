# Architecture

## Level 1 - Context

The Scrum-Poker app is stand alone web application.
It is self contained and has no external dependencies.

There are two types of users:
- **Admins** can change settings in a room. The first user is always the admin
- **Player** can vote

```s
          O
         /|\          ┌─────────────────┐
         / \  ───────▶│                 │
   O    ADMIN         │   Scrum-Poker   │
  /|\                 │    [Web App]    │
  / \  ──────────────▶│                 │
PLAYER                └─────────────────┘
```

## Level 2 - Container

The users use their browser to open the app ( [https://scrum-poker.sprinteins.com](https://scrum-poker.sprinteins.com) ).
The router directly serves the frontend files and forwards api requests to the API service.

```s
                              O
                             /|\
                             / \
                            USER
                              │
                              │
                              ▼
                  ╭───────────────────────╮
                  │ ◎ ○ ○ ░░░░░░░░░░░░░░░░│
                  ├───────────────────────┤
                  │        Browser        │
           ─ ─ ─ ▶│   [Firefox, Safari,   │─ ─ ─ ─ ─
          │       │        Chrome]        │         │
                  │                       │      <use>
          │       └───────────────────────┘         │
                              │
          │                   │                     │
             https://scrum-poker.sprinteins.com
          │                   │                     │
┌─────────────────────────────┼────────────────────────────┐
│         │                   │                     │      │
│                             ▼                            │
│         │             ┌───────────┐               │      │
│                       │  Router   │                      │
│         │             │  [nginx]  │               │      │
│                       │           │                      │
│         │             └───────────┘               │      │
│                             │                            │
│         │           ┌───────┴─────┐               │      │
│    <deliver>   / (default)      /api                     │
│         │           │             │               │      │
│                     │             ●                      │
│         │           │             ◡               │      │
│                     │   http://localhost:7788            │
│         │           ▼             │               ▼      │
│      ┌─────────────────────┐  ┌─────────────────────┐    │
│      │      Frontend       │  │       Backend       │    │
│      │[TypeScript, Svelte] │  │        [Go]         │    │
│      └─────────────────────┘  └─────────────────────┘    │
│                                                          │
│                                                          │
│Server [Ubuntu, Jiffybox]                                 │
└──────────────────────────────────────────────────────────┘
```

