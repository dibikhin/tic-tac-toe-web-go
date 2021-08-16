# tic-tac-toe-web-go
3x3 Tic-tac-toe for 2 friends in web

Wikipedia: Tic-tac-toe (noughts and crosses) is a game for two players, X and O, who take turns marking the spaces in a 3×3 grid. The player who succeeds in placing three of their marks in a diagonal, horizontal, or vertical row is the winner.

# Specs DRAFT

At once:
- 1 game
- 2 players
- 3x3 board hardcoded

Other:
- PvP only (for 2 friends, cannot play vs. computer yet)
- Turn's timeout? TODO:
- No saves
- Dummy logins

Architecture

General:
- 1 server + (0 to max 2) clients

Server:
- Data: in-memory db (survives clients' disconnects)
- Sessions: ? TODO:
- Caching: no
- Timeouts: ? TODO:
- Passive (IoC, client pulls server)
  - BUT server rules client
  - fat
- Stateless
- https://12factor.net/ TODO:
- Longrunning (not a FaaS)

Techs:
- Go (1.11+)
- gRPC (v1.37+)
- Protobuf (v3)

Client:
- Go (1.11+)
- gRPC (v1.37+)
- Protobuf (v3)
- Active
  - BUT server rules client (to keep client simple)
  - thin, w/o logic, status loop only
- CLI
