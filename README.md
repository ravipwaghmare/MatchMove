# MatchMove

# Sample Curl Commands

Create User:-
curl -X POST localhost:3000/auth/user/signup -H 'Content-Type: application/json' -d '{"username" :"xxxxx","name":"xxxxxx","email":"xxxxx@gmail.com","password":"xxxxx","acctype":"Admin"}'

Genertae Token For Client:-
curl -X POST localhost:3000/token/sendtoken -H 'Content-Type: application/json' -d '{"admin_email" :"xxxx@gmail.com","client_email":"xxxx@gmail.com"}'


Sign In:-
curl -X POST localhost:3000/auth/user/signin -H 'Content-Type: application/json' -d '{"email":"xxxxx@gmail.com", "password":"12345678", "token":"21fb7bd948879bed"}'
